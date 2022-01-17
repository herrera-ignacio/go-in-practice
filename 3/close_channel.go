package main

import (
	"fmt"
	"time"
)

/**
The best way to close a channel is to use one additional channel to
indicate that you're done with the channel.
*/

func main() {
	msg := make(chan string)
	done := make(chan bool)
	until := time.After(5 * time.Second)

	go send(msg, done)

	for {
		select {
		case m := <-msg:
			fmt.Println(m)
		case <-until:
			done <- true
			time.Sleep(500 * time.Millisecond)
			return
		}
	}
}

/**
ch is a receiving channel, while done is a sending channel.
*/
func send(ch chan<- string, done <-chan bool) {
	for {
		select {
		case <-done:
			println("Done")
			close(ch)
			return
		default:
			ch <- "hello"
			time.Sleep(500 * time.Millisecond)
		}
	}
}
