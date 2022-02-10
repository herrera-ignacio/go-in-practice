package main

import (
	"net/http"
	"text/template"
)

// Parses the template when the package is initialized
var t = template.Must(template.ParseFiles("templates/simple.html"))

type Page struct {
	Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "An Example",
		Content: "Have fun stormin' da castle.",
	}
	// Only executes the template in the http handler function
	t.Execute(w, p)
}

func main() {
	http.HandlerFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}
