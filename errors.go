package realip

import "errors"

// ErrInvalidIP is returned when an IP address that
// is considered to be real is found to be invalid.
var ErrInvalidIP = errors.New("realip: IP address is not valid")

// ErrUnableToDetermineRealIP occurs when the X-Forwarded-For header
// is non-empty, but contains only unacceptable addresses.
var ErrUnableToDetermineRealIP = errors.New("realip: unable to determine real IP address")
