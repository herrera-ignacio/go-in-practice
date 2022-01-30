package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Concat returns a string and an error
func Concat(parts ...string) (string, error) {
	// Returns an error if nothing was passed in
	if len(parts) == 0 {
		return "", errors.New("no strings supplied")
	}

	return strings.Join(parts, " "), nil
}

func main() {
	args := os.Args[1:]

	if result, err := Concat(args...); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Concatenated string: '%s'\n", result)
	}
}
