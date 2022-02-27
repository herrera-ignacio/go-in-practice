package main

import (
	"fmt"
	"os/exec"
)

func checkDep(name string) error {
	// Checks the whether the passed-in dependency is in one of the PATHs.
	// Returns an error when the dependency isn't found
	if _, err := exec.LookPath(name); err != nil {
		es := "Could not find '%s' in PATH: %s"
		return fmt.Errorf(es, name, err)
	}

	return nil
}
