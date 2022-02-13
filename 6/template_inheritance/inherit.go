package main

import (
	"net/http"
	"text/template"
)

// A map to store templates in a map of named templates
var t map[string]*template.Template

func init() {
	// Sets up the template map
	t = make(map[string]*template.Template)
	temp := template.Must(template.ParseFiles("base.html", "user.html"))
	t["user.html"] = temp
	temp = template.Must(template.ParseFiles("base.html", "page.html"))
	t["page.html"] = temp
}

type Page struct {
	Title, Content string
}

type User struct {
	Username, Name string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "An Example",
		Content: "Have fun stormin' da castle.",
	}
	t["page.html"].ExecuteTemplate(w, "base", p)
}

func displayUser(w http.ResponseWriter, r *http.Request) {
	u := &User{
		Username: "iherrera",
		Name:     "Ignacio Herrera",
	}
	t["user.html"].ExecuteTemplate(w, "base", u)
}

func main() {
	http.HandleFunc("/user", displayUser)
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}
