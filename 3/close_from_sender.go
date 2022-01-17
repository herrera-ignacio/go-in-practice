package main

import "time"

/**
By closing the channel from the sender, no panic will happen
but something weird will happen, you will print thousands of "Received"
because a closed channel always returns the channel's `nil` value (i.e., false in this case).
*/

func main() {
	ch := make(chan bool)
	timeout := time.After(600 * time.Millisecond)

	go send(ch)

	for {
		select {
		case <-ch:
			println("Received")
		case <-timeout:
			println("Time out")
			return // exit goroutine!
		default:
			println("Waiting...")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

/**
Sends a single message over the channel and then closes the channel.
*/
func send(ch chan bool) {
	time.Sleep(120 * time.Millisecond)
	ch <- true
	close(ch)
	println("Sent and closed")
}
