package main

import (
	"github.com/Sanmoo/go-api-lambda-boilerplate/adapters/inbound/aws"
	"github.com/Sanmoo/go-api-lambda-boilerplate/adapters/inbound/http"
	"github.com/Sanmoo/go-api-lambda-boilerplate/adapters/outbound/storage/memory"
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/usecases"
)

func main() {
	aws.ListenAndServe(&aws.BooksHandler{
		BooksHandler: *http.NewBooksHandler(usecases.NewBooksUsecases(memory.NewBooksRepository())),
	}, "books")
}
