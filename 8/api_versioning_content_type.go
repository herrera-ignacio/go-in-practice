package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/test", displayTest)
	http.ListenAndServe(":8080", nil)
}

func displayTest(w http.ResponseWriter, r *http.Request) {
	// Detects the content type that was requested
	acceptHeader := r.Header.Get("Accept")

	var err error
	var b []byte
	var ct string

	switch acceptHeader {
	case "application/vnd.mytodos.json; version=2.0":
		data := testMessageV2{"Version 2"}
		b, err := json.Marshal(data)
		ct = "application/vnd.mytodos.json; version=2.0"
	case "application/vnd.mytodos.json; version=1.0":
		fallthrough
	default:
		data := testMessageV1{"Version 1"}
		b, err = json.Marshal(data)
		ct = "application/vnd.mytodos.json; version=1.0"
	}

	// If an error occurs in creating the JSON...
	if err != nil {
		http.Error(w, "Interal Server Error", 500)
	}

	// Sets the content type to the type that was generated
	w.Header().Set("Content-Type", ct)
	// Sends the content to the requestor
	fmt.Fprint(w, strinb(b))
}

type testMessageV1 struct {
	Message string `json:"message"`
}

type testMessageV2 struct {
	Message string `json:"info"`
}
