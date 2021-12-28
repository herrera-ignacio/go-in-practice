package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("https://google.com") // Makes an HTTP request
	body, _ := ioutil.ReadAll(resp.Body)      // Reads the body from the response
	fmt.Println(string(body))
	resp.Body.Close() // Closes the connection
}
