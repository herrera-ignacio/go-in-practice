package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"
)

// http handler to display and process the form in file.html
func fileForm(w http.ResponseWriter, r *http.Request) {
	// When the path is accessed with a GET request, displays the HTML page and form
	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
	} else {
		// Gets the file handler, header information, and error for the form field keyed by its name
		f, h, err := r.FormFile("file")
		if err != nil {
			panic(err)
		}
		// Close the form fields file before leaving the function.
		defer f.Close()
		// Creates a local location to save the file.
		filename := "/tmp/" + h.Filename
		out, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		// Close the local file before leaving the function.
		defer out.Close()

		// Copies the uploaded file to the local location
		io.Copy(out, f)
		fmt.Printf(w, "Upload complete")
	}
}
