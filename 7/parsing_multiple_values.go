package main

import (
	"fmt"
	"net/http"
)

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	// Maximum memory to store file parts, where rest is stored to disk
	maxMemory := 16 << 20

	// Parses a multipart form
	err := r.ParseMultipartForm(maxMemory)

	// Handles any error parsing the form
	if err != nil {
		fmt.Println(err)
	}

	// Iterates over all the POST values of the names from field
	for k, v := range r.PostForm["names"] {
		fmt.Println(v)
	}
}
