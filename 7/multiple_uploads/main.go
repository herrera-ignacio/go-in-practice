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
		// Parses the form and handle errors
		err := r.ParseMultipartForm(16 << 20)
		if err != nil {
			fmt.FPrint(w, err)
			return
		}

		// Retrieves a slice containing the files from the Multipartform
		data := r.MultipartForm
		files := data.File["files"]

		// Iterates over the file uplodaed to the files field
		for _, file := range files {
			// Opens a file handler for one of the uploaded files
			f, err := file.Open()
			defer f.Close()
			if err != nil {
				fmt.Fprintf(w, err)
				return
			}

			// Creates a local file to store the contents of the uploaded file
			out, err := os.Create("/tmp/" + file.Filename)
			defer out.Close()
			if err != nil {
				fmt.Fprint(w, err)
				return
			}

			// Copies the uploaded file to the location on the filesystem
			_, err = io.Copy(out, f)
			if err != nil {
				fmt.Fprintln(w, err)
				return
			}
		}
		fmt.FPrint(w, "Upload completed")
	}
}
