package main

/**
# Output may change from run to run
Outside a goroutine
Outside again
Inside a goroutine
*/

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Outside a goroutine")

	go func() { // declares an anonymous function and executes it as a goroutine
		fmt.Println("Inside a goroutine")
	}()

	fmt.Println("Outside again")

	runtime.Gosched() // yields to the scheduler
}
