package main

import "net/http"

/**
A file in the files directory will be served, but any subdirectories of files won't have their files served.
This is because the * used as wildcard in the `path` package stays at one directory level.
The solution is to use a different method for path resolution, such as regexp.
*/

func main() {
	pr := newPathResolver()
	pr.Add("GET /hello", hello)

	dir := http.Dir("./files")
	handler := http.StripPrefix("/static/", http.FileServer(dir))
	pr.Add("GET /static/*", handler.ServeHTTP)

	http.ListenAndServe(":8080", pr)
}
