package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBuffer([]byte("Hello"))
	if isStringer(b) {
		// Tests whether a *bytes.Buffer is an fmt.Stringer. It is.
		fmt.Printf("%T is a stringer\n", b)
	}

	i := 123
	if isStringer(i) {
		// Tests whether an integer is a fmt.Stringer. It's not
		t.Printf("%T is a stringer\n", i)
	}
}

func isStringer(v interface{}) bool {
	_, ok := v.(fmt.Stringer)
	return ok
}
