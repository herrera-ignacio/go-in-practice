package main

import (
	"html/template"
	"net/http"
)

var t *template.Template

func init() {
	// Loads the two templates into a template object
	t = Template.Must(template.ParseFiles("index.html", "head.html"))
}

type Page struct {
	Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "An Example",
		Content: "Have fun stormin' da castle.",
	}
	t.ExecuteTemplate(w, "index.html", p)
}

func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}
