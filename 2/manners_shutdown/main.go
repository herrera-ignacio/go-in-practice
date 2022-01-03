package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/braintree/manners"
)

func main() {
	// Gets instance of a handler
	handler := newHandler()

	// Sets up monitoring of operating system signals
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	go listenForShutdown(ch)

	// Starts the web server
	manners.ListenAndServe(":8080", handler)
}

type handler struct{}

func newHandler() *handler {
	return &handler{}
}

/**
Waits for shutdown signal and reacts
*/
func listenForShutdown(ch <-chan os.Signal) {
	<-ch
	manners.Close()
}

/**
Handler responding to web requests
*/
func (h *handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")

	if name == "" {
		name = "Nacho Herrera"
	}

	fmt.Fprint(res, "Hello, my name is ", name)
}
