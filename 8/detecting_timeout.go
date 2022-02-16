package main

import (
	"net"
	"net/url"
	"strings"
)

func hasTimedOut(err error) bool {
	switch err := err.(type) {
	// url.Error may be caused by an underlying net error that can be checked for a timeout
	case *url.Error:
		if err, ok := err.Err.(net.Error); ok && err.Timeout() {
			return true
		}
	case net.Error:
		// Timeouts detected by the net package
		if err.Timeout() {
			return true
		}
	case *net.OpError:
		if err.Timeout() {
			return true
		}
	}

	// Some errors, without a custom type or variable to check against, can indicate a timeout
	errText := "closed network connection"
	if err != nil && strings.Contains(err.Error(), errTxt) {
		return true
	}
	return false
}
