package realip

import "net/http"

// DefaultDetector is used for package-level functions.
var DefaultDetector = &Detector{Blacklist: LocalNetworks}

// FromRequest returns client's public IP address from an http.Request.
// In case the real IP address cannot be determined, and error is returned.
func FromRequest(r *http.Request) (string, error) {
	return DefaultDetector.FromRequest(r)
}
