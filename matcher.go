package realip

import "net"

// Matcher provides a way to match the IP address.
type Matcher interface {
	Match(address string) (bool, error)
}

// NetMatcher contains logic to perform matching of the IP addresses
// agains a list of networks.
type NetMatcher []net.IPNet

// NetMatcherFromCIDR constructs a NetMatcher from a list of networks
// in CIDR notation.
// See net.ParseCIDR.
func NetMatcherFromCIDR(strings []string) (*NetMatcher, error) {
	nets := make([]net.IPNet, len(strings))
	for i, string := range strings {
		_, Net, err := net.ParseCIDR(string)
		if err != nil {
			return nil, err
		}
		nets[i] = *Net
	}
	return (*NetMatcher)(&nets), nil
}

var _ Matcher = (*NetMatcher)(nil)

// Match determines whether any of the networks contain the specified
// address.
func (m *NetMatcher) Match(address string) (bool, error) {
	ip := net.ParseIP(address)
	if ip == nil {
		return false, ErrInvalidIP
	}

	list := ([]net.IPNet)(*m)
	for i := range list {
		if list[i].Contains(ip) {
			return true, nil
		}
	}
	return false, nil
}
