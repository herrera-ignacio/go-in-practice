package pad

import (
	"log"
	"strings"
)

func Pad(s string, max uint) string {
	log.Printf("Testing Len: %d, Str: %s\n", max, s)
	ln := uint(len(s))

	if ln > max {
		// If the string is longer than the max, truncates it.
		return s[:max-1] // this should be s[:max], test will fail
	}

	// Pads the string until it's the max length
	s += strings.Repeat(" ", int(max-ln))
	return s
}
