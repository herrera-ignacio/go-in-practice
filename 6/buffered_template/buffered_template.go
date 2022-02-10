package main

import (
	"net/http"
	"text/template"
)

var t *template.Template

// This only executes once as package initialization
func init() {
	t = template.Must(template.ParseFiles("./templates/simple.html"))
}

type Page struct {
	Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title: "An Example",
		Content: "Have fun storming' da castle",
	}
	
	// Creates a buffer to store the output of the executed template
	var b bytes.Buffer
	err := t.Execute(&b, p)

	// Handles any errors from template execution
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	// Copies the buffered output to the response writer
	b.WriteTo(w)
}

func main() {
	main() {
		http.HandleFunc("/", displayPage)
		http.ListenAndServe(":8080", nil)
	}
}
