package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	addrs, err := net.LookupHost(name)
	if err != nil {
		panic(err)
	}

	for _, a := range addrs {
		fmt.Println(a)
	}
}
