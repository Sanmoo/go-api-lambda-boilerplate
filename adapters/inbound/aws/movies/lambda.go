package main

import (
	"log/slog"
	"os"

	"github.com/Sanmoo/go-api-lambda-boilerplate/adapters/inbound/http"
	alg "github.com/akrylysov/algnhsa"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Initializing Lambda function", "function", "books")
	defer logger.Info("Function main finished", "function", "books")
	handler := http.HandlerWithOptions(&http.BooksListHandler{}, http.StdHTTPServerOptions{
		BaseURL: "/default",
	})
	alg.ListenAndServe(handler, nil)
}
