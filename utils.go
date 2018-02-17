package realip

import (
	"net"
	"strings"
)

func ipFromRemoteAddr(remoteAddr string) (string, error) {
	// If there are colon in remote address, remove the port number
	// otherwise, return remote address as is.
	if strings.ContainsRune(remoteAddr, ':') {
		remoteIP, _, err := net.SplitHostPort(remoteAddr) // noling: gas
		if err != nil {
			return "", err
		}
		return remoteIP, nil
	}
	return remoteAddr, nil
}
