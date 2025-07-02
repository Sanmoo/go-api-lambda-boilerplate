package main

import (
	"log/slog"
	"os"

	"github.com/Sanmoo/go-api-lambda-boilerplate/adapters/inbound/aws"
	"github.com/Sanmoo/go-api-lambda-boilerplate/adapters/inbound/http"
	alg "github.com/akrylysov/algnhsa"
)

type MoviesHandler struct {
	aws.BaseHandler
	http.MoviesListHandler
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Debug("Initializing Lambda function", "function", "books")
	defer logger.Debug("Function main finished", "function", "books")
	handler := http.HandlerWithOptions(&MoviesHandler{}, http.StdHTTPServerOptions{
		BaseURL: "/default",
	})
	alg.ListenAndServe(handler, nil)
}
