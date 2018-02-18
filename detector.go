package realip

import (
	"net/http"
	"strings"
)

// Detector contains logic to detect real IP address.
type Detector struct {
	Blacklist Matcher
}

// FromRequest returns client's real IP address from an http.Request.
// An error is returned if real IP address can't be determined.
func (d *Detector) FromRequest(r *http.Request) (string, error) {
	// Fetch header value.
	xForwardedFor := r.Header.Get("X-Forwarded-For")

	if xForwardedFor == "" {
		ip, err := ipFromRemoteAddr(r.RemoteAddr)
		if err != nil {
			return "", err
		}
		unallowed, err := d.Blacklist.Match(ip)
		if err != nil {
			return "", err
		}
		if !unallowed {
			return ip, nil
		}
		return "", ErrUnableToDetermineRealIP
	}

	return d.FromXForwardedFor(xForwardedFor)
}

// FromXForwardedFor returns client's real IP address based on
// an X-Forwareded-For header value.
// An error is returned if real IP address can't be determined.
func (d *Detector) FromXForwardedFor(xForwardedFor string) (string, error) {
	// Check list of IP in X-Forwarded-For and return the *last*
	// allowed address.
	// We use last address because proxies append the addresses to
	// the end of the header (and *not* prepend them to the beginning).
	addresses := strings.Split(xForwardedFor, ",")
	l := len(addresses) - 1
	for i := range addresses {
		address := addresses[l-i]
		address = strings.TrimSpace(address)
		unallowed, err := d.Blacklist.Match(address)
		if err != nil {
			return "", err
		}
		if !unallowed {
			return address, nil
		}
	}
	return "", ErrUnableToDetermineRealIP
}
