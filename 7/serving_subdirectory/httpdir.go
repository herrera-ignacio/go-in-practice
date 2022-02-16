package main

import "net/http"

func main() {
	// A directory and its subdirectories on the filesystem are chosen to serve.
	dir := http.Dir("./files/")
	// The /static/ path serves the directory and needs to be removed before looking up file path.
	// Thie `StripPrefix` is used to remove any prefix in the URL before passing the path to the file server to find.
	// When serving a subpath in your application, this is needed to find the right files.
	handler := http.StripPrefix("/static/", http.FileServer(dir))
	http.Handle("/static/", handler)

	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}
