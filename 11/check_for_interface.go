package main

import (
	"fmt"
	"io"
)

type Name struct {
	First, Last string
}

// Name implements Stringer interface
func (n *Name) String() string {
	return n.First + " " + n.Last
}

func main() {
	n := &Name{First: "Nacho", Last: "Herrera"}
	// Creates a nil pointer of type fmt.Stringer
	stringer := (*fmt.Stringer)(nil)

	// Tests whether n is a fmt.Stringer (has a String() method)
	implements(n, stringer)

	// Creates a nil pointer of type io.Writer
	writer := (*io.Writer)(nil)

	// Tests whether n is an io.Writer (has a Write() method)
	implements(n, writer)
}

func implements(concrete interface{}, target interface{}) bool {
	// Gets a reflect.Type that describes the target of the pointer
	iface := reflect.Typeof(target).Elem()

	// Gets the reflect.Type of the concrete type passed in
	v := reflect.ValueOf(concrete)
	t := v.Type()

	if t.Implements(iface) {
		fmt.Printf("%T is a %s\n", concrete, iface.Name())
		return true
	}

	fmt.Printf("%T is not a %s\n", concrete, iface.Name())
	return false
}
