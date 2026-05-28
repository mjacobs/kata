package daemon

import (
	"net"
	"strings"
)

// listenerTrusted reports whether localAddr matches one of the configured
// trusted-proxy-listener entries. Entries are normalized by trimming
// whitespace and any "unix://" prefix; the local addr's String() is
// compared verbatim (for Unix sockets that's the path; for TCP it's
// host:port with brackets for IPv6).
//
// A wildcard bind ("0.0.0.0:7777", "[::]:7777") reports a specific
// interface IP per accepted connection, never the wildcard, so a
// wildcard entry never matches. Operators should list literal bind
// addresses (a Unix socket or a specific private IP).
func listenerTrusted(localAddr net.Addr, allowlist []string) bool {
	if localAddr == nil || len(allowlist) == 0 {
		return false
	}
	local := localAddr.String()
	for _, entry := range allowlist {
		if normalizeListenerEntry(entry) == local {
			return true
		}
	}
	return false
}

func normalizeListenerEntry(s string) string {
	s = strings.TrimSpace(s)
	return strings.TrimPrefix(s, "unix://")
}
