package main

import (
	"fmt"
	"time"
)

/**
This example will cause the program to panic
panic: send on closed channel
*/

func main() {
	msg := make(chan string)
	until := time.After(5 * time.Second)

	// starts a send goroutine with a sending channel
	go send(msg)

	// Loops over a select that watches for messages from send, or for a time-out
	for {
		select {
		case m := <-msg:
			fmt.Println(m)
		case <-until:
			// When the timeout occurs, shut things down.
			close(msg)
			// You pause to ensure that you see the failure before the main goroutine exits.
			time.Sleep(500 * time.Millisecond)
			return
		}
	}
}

// Sends "hello" to the channel every half-second
func send(ch chan string) {
	for {
		ch <- "hello"
		time.Sleep(500 * time.Millisecond)
	}
}
