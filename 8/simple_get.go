package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	res, _ := Get("http://goinpracticebook.com")
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", b)
}
