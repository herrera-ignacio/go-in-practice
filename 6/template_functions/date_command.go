package main

import (
	"net/http"
	"text/template"
	"time"
)

// HTML template as a string
var tpl = `<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title>Date Example</title>
	</head>
	<body>
		<!-- Pipes Date through the dateFormat command -->
		<p>{{.Date | dateFormat "Jan 2, 2006"}}</p>
	</body> 
</html>`

// Maps Go functions to template functions
var funcMap = template.FuncMap{
	"dateFormat": dateFormat,
}

// Function to convert a time to a formatted string
func dateFormat(layout string, d time.Time) string {
	return d.Format(layout)
}

func serveTemplate(res http.ResponseWriter, req *http.Request) {
	t := template.New("date")
	t.Funcs(funcMap)
	t.Parse(tpl)
	data := struct{ Date time.Time }{
		Date: time.Now(),
	}
	t.Execute(res, data)
}

func main() {
	http.HandleFunc("/", serveTemplate)
	http.ListenAndServe(":8080", nil)
}
