package daemon_test

import (
	"encoding/json"
	"net"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.kenn.io/kata/internal/config"
	"go.kenn.io/kata/internal/daemon"
)

// startTrustedProxyTestServer pre-binds a loopback TCP listener, then builds a
// daemon.Server whose trusted-proxy allowlist names that listener's address,
// and finally swaps the listener into an httptest.Server. The two-step dance
// exists because withTrustedProxyActor captures the allowlist at NewServer
// time; httptest.NewServer would otherwise pick its own random port and the
// allowlist entry would never match.
func startTrustedProxyTestServer(t *testing.T, headerName string) *httptest.Server {
	t.Helper()
	l, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)

	d := openTestDB(t)
	cfg := daemon.ServerConfig{
		DB:        d.db,
		StartedAt: d.now,
		Auth: config.AuthConfig{
			Proxy: config.ProxyConfig{
				TrustedActorHeader:    headerName,
				TrustedProxyListeners: []string{l.Addr().String()},
			},
		},
	}
	srv := daemon.NewServer(cfg)
	t.Cleanup(func() { _ = srv.Close() })
	ts := httptest.NewUnstartedServer(srv.Handler())
	require.NoError(t, ts.Listener.Close())
	ts.Listener = l
	ts.Start()
	t.Cleanup(ts.Close)
	return ts
}

// trustedProxyCreateProject posts to /api/v1/projects with a path-free init
// (name only) so the test does not need a git workspace on disk. Returns the
// new project's rowid. The headers map is passed through to doReq so callers
// can carry an X-Kata-Actor header on the setup call when token-identity mode
// or trusted-proxy mode requires attribution.
func trustedProxyCreateProject(t *testing.T, ts *httptest.Server, headers map[string]string) int64 {
	t.Helper()
	resp, raw := doReq(t, ts, "POST", "/api/v1/projects",
		map[string]any{"name": "trusted-proxy-e2e"}, headers)
	require.Equalf(t, http.StatusOK, resp.StatusCode, "init project: %s", raw)
	var body struct {
		Project struct {
			ID int64 `json:"id"`
		} `json:"project"`
	}
	require.NoError(t, json.Unmarshal(raw, &body), string(raw))
	require.NotZero(t, body.Project.ID, "init returned zero project id: %s", raw)
	return body.Project.ID
}

// trustedProxyCreateIssue posts to /api/v1/projects/{id}/issues and returns
// the new issue's short_id, which is the wire-format ref the close action
// expects. Headers are forwarded to doReq so the call can carry the trusted
// actor header along with the body actor.
func trustedProxyCreateIssue(t *testing.T, ts *httptest.Server, projectID int64, headers map[string]string) string {
	t.Helper()
	resp, raw := doReq(t, ts, "POST",
		"/api/v1/projects/"+strconv.FormatInt(projectID, 10)+"/issues",
		map[string]any{"actor": "setup", "title": "trusted proxy header e2e"}, headers)
	require.Equalf(t, http.StatusOK, resp.StatusCode, "create issue: %s", raw)
	var body struct {
		Issue struct {
			ShortID string `json:"short_id"`
		} `json:"issue"`
	}
	require.NoError(t, json.Unmarshal(raw, &body), string(raw))
	require.NotEmpty(t, body.Issue.ShortID, "create returned no short_id: %s", raw)
	return body.Issue.ShortID
}

// TestTrustedProxyHeader_CreditsHeaderActor exercises the full pipeline:
// withTrustedProxyActor turns the X-Kata-Actor header into a PrincipalTrustedProxy
// on the trusted listener, attributedActor prefers the principal's actor over
// the request body's actor, and the close event row records the header value
// rather than the body value.
func TestTrustedProxyHeader_CreditsHeaderActor(t *testing.T) {
	ts := startTrustedProxyTestServer(t, "X-Kata-Actor")

	// Seed via the HTTP surface so the trusted-proxy principal is the
	// only writer in scope. initProject (POST /api/v1/projects) does
	// not call attributedActor, but the header is still applied by the
	// middleware so passing it here is harmless and keeps the test
	// uniform. createIssue does call attributedActor, so the header is
	// required for the seed issue to materialize.
	projectID := trustedProxyCreateProject(t, ts, map[string]string{"X-Kata-Actor": "setup"})
	issueRef := trustedProxyCreateIssue(t, ts, projectID,
		map[string]string{"X-Kata-Actor": "setup"})

	// The body carries actor="client-claim"; the trusted header carries
	// "proxy-user". A passing test proves attributedActor preferred the
	// principal's actor over the body's.
	body := map[string]any{
		"actor":   "client-claim",
		"reason":  "done",
		"message": "verified end-to-end after the fix landed.",
		"source":  "tui",
	}
	resp, raw := doReq(t, ts, "POST",
		"/api/v1/projects/"+strconv.FormatInt(projectID, 10)+"/issues/"+issueRef+"/actions/close",
		body, map[string]string{"X-Kata-Actor": "proxy-user"})
	require.Equalf(t, http.StatusOK, resp.StatusCode, "close: %s", raw)

	var payload struct {
		Event *struct {
			Actor string `json:"actor"`
		} `json:"event"`
	}
	require.NoError(t, json.Unmarshal(raw, &payload), string(raw))
	require.NotNil(t, payload.Event, "close returned no event row: %s", raw)
	assert.Equal(t, "proxy-user", payload.Event.Actor,
		"event.actor must reflect the trusted header value, not the body actor")
}

// TestTrustedProxyHeader_MissingHeaderRejects verifies the negative path:
// when the trusted listener is configured but the client omits the actor
// header, withTrustedProxyActor installs PrincipalTrustedProxyAbsent and
// ensureAttributedWriteAllowed rejects the write with actor_header_required.
// Without this guard a misconfigured proxy could silently let body-supplied
// actor values through on a listener that callers expect to be header-attributed.
func TestTrustedProxyHeader_MissingHeaderRejects(t *testing.T) {
	ts := startTrustedProxyTestServer(t, "X-Kata-Actor")

	projectID := trustedProxyCreateProject(t, ts, map[string]string{"X-Kata-Actor": "setup"})
	issueRef := trustedProxyCreateIssue(t, ts, projectID,
		map[string]string{"X-Kata-Actor": "setup"})

	body := map[string]any{
		"actor":   "client-claim",
		"reason":  "done",
		"message": "should be rejected because the trusted header is missing.",
		"source":  "tui",
	}
	resp, raw := doReq(t, ts, "POST",
		"/api/v1/projects/"+strconv.FormatInt(projectID, 10)+"/issues/"+issueRef+"/actions/close",
		body, nil) // no X-Kata-Actor on the close call
	assertAPIError(t, resp.StatusCode, raw, http.StatusBadRequest, "actor_header_required")
}

// TestTrustedProxyHeader_ReadsNotBlocked verifies that a header-less GET on a
// trusted listener still returns 2xx. The middleware stores PrincipalTrustedProxyAbsent
// for the read, but read handlers never call attributedActor, so the missing
// header must not reject the request.
func TestTrustedProxyHeader_ReadsNotBlocked(t *testing.T) {
	ts := startTrustedProxyTestServer(t, "X-Kata-Actor")

	// GET /api/v1/health is a read with no actor and no required setup.
	// On a trusted listener with no header, the middleware stores a
	// trusted-but-absent principal; the health handler never calls
	// attributedActor, so it must still return 200.
	resp, raw := doReq(t, ts, "GET", "/api/v1/health", nil, nil)
	require.Equal(t, http.StatusOK, resp.StatusCode, "body: %s", raw)
}

// TestTrustedProxyHeader_ModeOffUnchanged verifies that with mode off (empty
// TrustedActorHeader) the header is ignored and the body actor flows through
// as it did before this feature landed. The middleware short-circuits before
// checking the allowlist, so even a request that would otherwise look like a
// trusted-proxy call is attributed to the body actor.
func TestTrustedProxyHeader_ModeOffUnchanged(t *testing.T) {
	// Mode off (empty TrustedActorHeader). The header is ignored, and the
	// body actor is used as today.
	ts := startTrustedProxyTestServer(t, "" /* mode off */)

	projectID := trustedProxyCreateProject(t, ts, nil)
	issueRef := trustedProxyCreateIssue(t, ts, projectID, nil)

	body := map[string]any{
		"actor":   "client-claim",
		"reason":  "done",
		"message": "mode off; body actor must be used.",
		"source":  "tui",
	}
	resp, raw := doReq(t, ts, "POST",
		"/api/v1/projects/"+strconv.FormatInt(projectID, 10)+"/issues/"+issueRef+"/actions/close",
		body, map[string]string{"X-Kata-Actor": "should-be-ignored"})
	require.Equal(t, http.StatusOK, resp.StatusCode, "body: %s", raw)

	var payload struct {
		Event *struct {
			Actor string `json:"actor"`
		} `json:"event"`
	}
	require.NoError(t, json.Unmarshal(raw, &payload))
	require.NotNil(t, payload.Event)
	assert.Equal(t, "client-claim", payload.Event.Actor,
		"mode off: header is ignored, body actor wins")
}
