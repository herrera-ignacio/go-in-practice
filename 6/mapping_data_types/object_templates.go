package main

import (
	"bytes"
	"html/template"
	"net/http"
)

// Variables to hold persistent data shared between requests.
var t *template.Template
var qc template.HTML // Holds quote data type to be used from a higher-level template

func init() {
	// Loads the two templates files for later use.
	t = template.Must(template.ParseFiles("index.html", "quote.html"))
}

type Page struct {
	Title   string
	Content template.HTML
}

type Quote struct {
	Quote, Name string
}

func main() {
	// Populates a dataset to supply to template
	q := &Quote{
		Quote:  `You keep using that word. I do not thing it means what you think it means.`,
		Person: "Nacho Herrera",
	}

	var b bytes.Buffer
	// Writes templates and data
	t.ExecuteTemplate(&b, "quote.html", q)
	// Stores quote as HTML text in a global variable
	qc = template.HTML(b.String())

	// Serves handler using built-in web server
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "A User",
		Content: qc,
	}
	t.ExecuteTemplate(w, "index.html", p)
}
