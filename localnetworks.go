package realip

// LocalNetworks is a built-in matcher that provides
// matching against well-known local networks.
var LocalNetworks *NetMatcher

func init() {
	networks := []string{
		"127.0.0.1/8",    // localhost
		"10.0.0.0/8",     // 24-bit block
		"172.16.0.0/12",  // 20-bit block
		"192.168.0.0/16", // 16-bit block
		"169.254.0.0/16", // link local address
		"::1/128",        // localhost IPv6
		"fc00::/7",       // unique local address IPv6
		"fe80::/10",      // link local address IPv6
	}

	matcher, err := NetMatcherFromCIDR(networks)
	if err != nil {
		panic(err)
	}
	LocalNetworks = matcher
}
