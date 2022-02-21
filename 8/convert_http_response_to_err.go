package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Error struct {
	HTTPCode int    `json:"-"`
	Code     int    `json:"code, omitempty"`
	Message  string `json:"message"`
}

// The Error method implements the error interface on the Error struct
// This allows instances of `Error` to be passed between functions as an `error`.
func (e Error) Error() string {
	fs := "HTTP: %d, Code: %d, Message %s"
	return fmt.Sprintf(fs, e.HTTPCode, e.Code, e.Message)
}

func get(u string) (*http.Response, error) {
	res, err := http.Get(u)
	if err != nil {
		return res, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		if res.Header.Get("Content-Type") != "application/json" {
			sm := "Uknown error. HTTP status: %s"
			return res, fmt.Errorf(sm, res.Status)
		}

		// Reads the body of the response into a buffer
		b, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()

		// Parses the JSON response and responds to any errors
		err = json.Unmarshal(b, &data)
		if err != nil {
			sm := "Unable to parse json: %s, HTTP status: %s"
			return res, fmt.Errorf(sm, err, res.Status)
		}
		// Ads the HTTP status code to the Error instance
		data.Err.HTTPCode = res.statusCode

		// Returns the custom error and the response
		return res, data.Err
	}

	// When there's no error, returns the response as expected
	return res, nil
}

func main() {
	res, err := get("http://localhost:8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s, b")
}
