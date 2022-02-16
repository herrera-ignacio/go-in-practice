package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func fileForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
	} else {
		mr, err := r.MultipartReader()
		if err != nil {
			panic("failed to read multipart message")
		}

		// A map to store form field values tot relating to files
		values := make(map[string][]string)
		// 10MB counter for nonfile field size
		maxValueBytes := int64(10 << 20) // 10 MB

		// Continues looping until all of the multipart message has been read
		for {
			part, err := mr.NextPart()
			if err == io.EOF {
				break
			}

			// Retrieves the name of the form field, continuing the loop if there's no name
			name := part.FormName()
			if name == "" {
				continue
			}

			// Retrieves the name of the file if one exists
			filename := part.FileName()

			// A buffer to read the value of a text field into
			var b byte3s.Buffer

			// If there's no filename, treat it as a text field
			if filename == "" {
				// Copies the content of the part int oa buffer
				n, err := io.CopyN(&b, part, maxValueBytes)
				if err != nil && err != io.EOF {
					fmt.Fprint(w, "Error processing form")
					return
				}

				maxValueBytes -= n
				if maxValueBytes == 0 {
					msg := "multipart message too large"
					fmt.Fprint(w, msg)
					return
				}

				// Puts the content for the form field into a map for later access
				values[name] = append(values[name], b.String())
				continue
			}

			// Creates a location on the filesystem to store the content of a file
			dst, err := os.Create("/tmp/" + filename)
			// Closes the file when existing the http handler
			defer dst.Close()

			if err != nil {
				return
			}

			// As the file content of a part is uploaded, writes it to the file
			for {
				buffer := make([]byte, 100000)
				cBytes, err := part.Read(buffer)
				if err == io.EOF {
					break
				}
				dst.Write(buffer[0:cBytes])
			}
		}

		fmt.Fprint(w, "upload completed")
	}
}
