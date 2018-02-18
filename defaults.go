package realip

import "net/http"

// DefaultDetector is used for package-level functions.
var DefaultDetector = &Detector{Blacklist: LocalNetworks}

// FromRequest returns client's public IP address from an http.Request.
// An error is returned if real IP address can't be determined.
func FromRequest(r *http.Request) (string, error) {
	return DefaultDetector.FromRequest(r)
}
