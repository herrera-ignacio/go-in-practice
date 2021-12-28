package main

import (
	"fmt"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello, my name is Nacho Herrera")
}

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("Listening on localhost:4000")
	http.ListenAndServe("localhost:4000", nil)
}
