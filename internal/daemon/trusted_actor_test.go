package daemon

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListenerTrusted(t *testing.T) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "100.64.0.5:7777")
	if err != nil {
		t.Fatal(err)
	}
	unixAddr := &net.UnixAddr{Name: "/run/kata/proxy.sock", Net: "unix"}
	otherTCP, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")

	cases := []struct {
		name      string
		local     net.Addr
		allowlist []string
		want      bool
	}{
		{"empty allowlist", tcpAddr, nil, false},
		{"nil addr", nil, []string{"100.64.0.5:7777"}, false},
		{"tcp match", tcpAddr, []string{"100.64.0.5:7777"}, true},
		{"tcp no match", otherTCP, []string{"100.64.0.5:7777"}, false},
		{"unix match with prefix", unixAddr, []string{"unix:///run/kata/proxy.sock"}, true},
		{"unix match plain path", unixAddr, []string{"/run/kata/proxy.sock"}, true},
		{"unix no match", unixAddr, []string{"/different/path"}, false},
		{"whitespace in entry trimmed", tcpAddr, []string{"  100.64.0.5:7777 "}, true},
		{"wildcard 0.0.0.0 never matches", tcpAddr, []string{"0.0.0.0:7777"}, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := listenerTrusted(tc.local, tc.allowlist)
			assert.Equal(t, tc.want, got)
		})
	}
}
