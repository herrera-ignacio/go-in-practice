package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("can't divide by zero")

func main() {
	fmt.Println("Divide 1 by 0 with precheck")
	_, err := precheckDivide(1, 0) // Error handled
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Println("Divide 2 by 0 without precheck")
	divide(2, 0) // PANIC! runtime error
}

func precheckDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return divide(a, b), nil
}

func divide(a, b int) int {
	return a / b
}
