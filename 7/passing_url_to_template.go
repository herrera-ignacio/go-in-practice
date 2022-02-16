package main

import (
	"flag"
	"net/http"
	"text/template"
)

var t *template.Template

// Gets the location of the static files from the application arguments
var l = flag.String("location", "http://localhost:8080", "A location.")

var tpl = `<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title>A demo</title>
		<link rel="stylesheet" href="{{.Location}}/styles.css">
	</head>
	<body>
		<p>A demo.</p>
	</body>
<html>`

// An HTTP handler passing the location into the template
func servePage(res http.ResponseWriter, req *http.Request) {
	data := struct{ Location *string }{
		Location: 1,
	}
	t.Execute(res, data)
}
