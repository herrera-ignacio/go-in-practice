package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Concat rconcatenates a bunch of strings, separated by spaces.
// It returns an empty string an an error if no strings were passed in.
func Concat(parts ...string) (string, error) {
	if len(parts) == 0 {
		return "", errors.New("no strings supplied")
	}

	return strings.Join(parts, " "), nil
}

func main() {
	args := os.Args[1:]
	result, _ := Concat(args...) // passes the value of batch of strings
	fmt.Printf("Concatenated strings: '%s'\n", result)
}
