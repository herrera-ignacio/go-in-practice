package main

import (
	"os"
	"time"
)

func main() {
	/**
	Creates a channel that will receive a message when 30 seconds have elapsed
	*/
	done := time.After(30 * time.Second)

	/**
	Makes a new channel por passing bytes from Stdin to Stdout.
	Because you haven't specified a size, this channel can hold
	only one message at a time.
	*/
	echo := make(chan []byte)

	go readStdin(echo)

	for {
		/**
		Uses a select statement to pass data from Stdin to Stdout when received,
		or to shut down when time-out event occurs.
		*/
		select {
		case buf := <-echo:
			os.Stdout.Write(buf)
		case <-done:
			os.Exit(0)
		}
	}

}

/**
Takes a write-only channel (chan<-) and
sends any received input to that channel.
*/
func readStdin(out chan<- []byte) {
	for {
		/**
		Copies some data from Stdin into data.
		Note that File.Read blocks until it receives data.
		*/
		data := make([]byte, 1024)
		l, _ := os.Stdin.Read(data)
		if l > 0 {
			/**
			Sends the buffered data over the channel.
			*/
			out <- data
		}
	}
}
