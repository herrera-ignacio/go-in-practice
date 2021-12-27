package main

import "fmt"

func Names() (first string, second string) {
	/**
	Values assigned to named return variables
	**/
	first = "Foo"
	second = "Bar"
	return
}

func main() {
	n1, n2 := Names()
	fmt.Println(n1, n2)
}
