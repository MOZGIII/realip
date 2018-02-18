package realip

import (
	"net"
	"strings"
)

func ipFromRemoteAddr(remoteAddr string) (string, error) {
	if strings.ContainsRune(remoteAddr, ':') {
		remoteIP, _, err := net.SplitHostPort(remoteAddr) // noling: gas
		if err != nil {
			return "", err
		}
		return remoteIP, nil
	}
	return remoteAddr, nil
}
