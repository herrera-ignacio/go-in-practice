package main

import (
	"fmt"
	"net/http"
)

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	// Parses a simple form containing only text-based fields
	err := r.ParseForm()
	// Handles any errors that ocurred parsing the form
	if err != nil {
		fmt.Println(err)
	}
	// Gets the first value for the name field from the form
	name := r.FormValue("name")
}
