package main

/**
Although existing services such as *Logstash* and *Heka* are available, you'll avail yourself of a simple UNIX tool called *Netcat* (`nc`).

You can start a simple TCP server that listenes (`-l`) continuously (`-k`) on port 1902: `nc -lk 1902`.
*/

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1902")
	if err != nil {
		panic("Failed to connect to localhost:1902")
	}
	defer conn.Close()

	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "example ", f)
	logger.Println("This is a regular message.")
	logger.Panicln("This is a panic.")
}
