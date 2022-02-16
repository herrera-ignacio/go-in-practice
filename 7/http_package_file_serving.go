package main

import "net/http"

func main() {
	// Uses a directory on the filesystem
	dir := http.Dir("./files")
	// Serves the filesystem directory
	http.ListenAndServe(":8080", http.FileServer(dir))
}
