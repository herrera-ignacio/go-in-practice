package main

import (
	"errors"
	"fmt"
	"math/rand"
)

var ErrTimeout = errors.New("request timeout")
var ErrRejected = errors.New("request rejected")

var random = rand.New(rand.NewSource(35))

func main() {
	response, err := SendRequest("Hello") // tubbed-out

	// Handles the time-out condition with retries
	for err == ErrTimeout {
		fmt.Println("Timeout. Retrying.")
		response, err = SendRequest("Hello")
	}

	if err != nil {
		// Handles any other error as a failure
		fmt.Println(err)
	} else {
		// If there's no error, prints the result
		fmt.Println(response)
	}
}

// Stubs out a message sender
func SendRequest(req string) (string, error) {
	switch random.Int() % 3 {
	case 0:
		return "Success", nil
	case 1:
		return "", ErrRejected
	default:
		return "", ErrTimeout
	}
}
