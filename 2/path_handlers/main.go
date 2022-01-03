package main

import (
	"fmt"
	"net/http"
	"path"
	"strings"
)

func main() {
	// 1. Gets an instance of a path-based router
	pr := newPathResolver()

	// 2. Maps functions to paths
	pr.Add("GET /hello", hello)
	pr.Add("* /goodbye/*", goodbye)

	// 3. Sets the HTTP server to use your router
	http.ListenAndServe(":8080", pr)

}

func newPathResolver() *pathResolver {
	return &pathResolver{make(map[string]http.HandlerFunc)}
}

type pathResolver struct {
	handlers map[string]http.HandlerFunc
}

func (p *pathResolver) Add(path string, handler http.HandlerFunc) {
	p.handlers[path] = handler
}

func (p *pathResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	check := req.Method + " " + req.URL.Path

	// 4. Iterates over registered path
	for pattern, handlerFunc := range p.handlers {
		// 5. Check whether current path matches a registered one
		if ok, err := path.Match(pattern, check); ok && err == nil {
			// Executes the handler function for a matched path
			handlerFunc(res, req)
			return
		} else if err != nil {
			fmt.Fprint(res, err)
		}
	}

	// 7. If no path matches, the page wasn't found
	http.NotFound(res, req)
}

func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")

	if name == "" {
		name = "Nacho Herrera"
	}

	fmt.Fprint(res, "Hello, my name is ", name)
}

func goodbye(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	name := parts[2]

	if name == "" {
		name = "Nacho Herrera"
	}

	fmt.Fprint(res, "Goodbye ", name)
}
