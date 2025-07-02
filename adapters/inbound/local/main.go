package main

import (
	"log"
	"net/http"

	adpt "github.com/Sanmoo/go-api-lambda-boilerplate/adapters/inbound/http"
)

type LocalBaseHandler struct {
	adpt.BlankHandler
}

type LocalHandler struct {
	LocalBaseHandler
	adpt.BooksHandler
	adpt.MoviesListHandler
}

func newLocalHandler() LocalHandler {
	return LocalHandler{}
}

func main() {
	var server adpt.ServerInterface = newLocalHandler()
	h := adpt.HandlerWithOptions(server, adpt.StdHTTPServerOptions{
		BaseURL: "/default",
	})

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(s.ListenAndServe())
}
