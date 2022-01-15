package main

import "time"

func main() {
	// Blocks for five secondes
	time.Sleep(5 * time.Second)

	// Creates a channel that will get notified in five seconds,
	sleep := timeAfter(5 * time.Second)

	// Then block until that channel receives a notification.
	<-sleep
}
