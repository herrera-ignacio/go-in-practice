package main

/**
This program uses a goroutine to run the echoing behavior
in the background, while a timer runs in the foreground.
*/

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	go echo(os.Stdin, os.Stdout) // goroutine
	time.Sleep(15 * time.Second) // sleeps 30 sec
	fmt.Println("Timed out")     // Done sleeping
	os.Exit(0)
}

func echo(in io.Reader, out io.Writer) {
	io.Copy(out, in) // Copies data to an os.Writer from an os.Reader
}
