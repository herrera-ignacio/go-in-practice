package main

import (
	"fmt"
	"time"
)

func printCount(c chan int) {
	num := 0
	for num >= 0 {
		num = <-c // wait for value to come in
		fmt.Print(num, " ")
	}
}

func main() {
	c := make(chan int)
	a := []int{8, 6, 7, 5, 3, 0, 9, -1}

	go printCount(c)

	// Passes ints into channel
	for _, v := range a {
		c <- v
	}

	time.Sleep(time.Millisecond * 1) // main pauses before ending

	fmt.Println("End of main")
}
