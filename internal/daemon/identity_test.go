package daemon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPrincipalKind_TrustedProxyConstants locks the values of the two
// trusted-proxy principal kinds. They must be distinct from each other
// and from the kinds defined by the identity layer so logs and audit
// reads can tell them apart.
func TestPrincipalKind_TrustedProxyConstants(t *testing.T) {
	assert.NotEqual(t, PrincipalTrustedProxy, PrincipalTrustedProxyAbsent)
	assert.NotEqual(t, PrincipalTrustedProxy, PrincipalDBToken)
	assert.NotEqual(t, PrincipalTrustedProxy, PrincipalStaticToken)
	assert.NotEqual(t, PrincipalTrustedProxy, PrincipalBootstrap)
	assert.NotEqual(t, PrincipalTrustedProxyAbsent, PrincipalDBToken)
	assert.NotEqual(t, PrincipalTrustedProxyAbsent, PrincipalStaticToken)
	assert.NotEqual(t, PrincipalTrustedProxyAbsent, PrincipalBootstrap)

	// PrincipalKind is a string named type. Lock the snake_case string
	// values so they match the existing convention ("db_token",
	// "bootstrap", "static_token") and stay stable for log/audit
	// consumers.
	assert.Equal(t, PrincipalKind("trusted_proxy"), PrincipalTrustedProxy)
	assert.Equal(t, PrincipalKind("trusted_proxy_absent"), PrincipalTrustedProxyAbsent)
}
