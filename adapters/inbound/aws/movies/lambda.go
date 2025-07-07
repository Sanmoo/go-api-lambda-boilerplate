package main

import (
	"github.com/Sanmoo/go-api-lambda-boilerplate/adapters/inbound/aws"
	"github.com/Sanmoo/go-api-lambda-boilerplate/adapters/inbound/http"
	"github.com/Sanmoo/go-api-lambda-boilerplate/adapters/outbound/storage/memory"
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/usecases"
)

func main() {
	aws.ListenAndServe(&aws.MoviesHandler{
		MoviesHandler: *http.NewMoviesHandler(usecases.NewMoviesUsecases(memory.NewMoviesRepository())),
	}, "movies")
}
