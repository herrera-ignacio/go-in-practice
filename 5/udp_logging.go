package main

import (
	"log"
	"net"
	"time"
)

func main() {
	timeout := 30 * time.Second
	conn, err = net.DialTimeout("udp", "localhost:1902", timeout)
	if err != nil {
		panic("Failed to connect.")
	}
	defer conn.Close()

	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "example", f)
	logger.Println("This is a regular message.")
	logger.Panicln("This is a panic.")
}
