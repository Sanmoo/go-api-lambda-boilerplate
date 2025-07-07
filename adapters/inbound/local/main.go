package main

import (
	"log"
	"net/http"

	adpt "github.com/Sanmoo/go-api-lambda-boilerplate/adapters/inbound/http"
	"github.com/Sanmoo/go-api-lambda-boilerplate/adapters/outbound/storage/memory"
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/usecases"
	"github.com/getkin/kin-openapi/openapi3"
	middleware "github.com/oapi-codegen/nethttp-middleware"
)

type LocalBaseHandler struct {
	adpt.BlankHandler
}

type LocalHandler struct {
	adpt.BooksHandler
	adpt.MoviesHandler
	adpt.TvSeriesHandler
	adpt.ElectronicGamesHandler
	adpt.NonElectronicGamesHandler
}

func newLocalHandler() *LocalHandler {
	return &LocalHandler{
		BooksHandler:              *adpt.NewBooksHandler(usecases.NewBooksUsecases(memory.NewBooksRepository())),
		MoviesHandler:             *adpt.NewMoviesHandler(usecases.NewMoviesUsecases(memory.NewMoviesRepository())),
		TvSeriesHandler:           *adpt.NewTvSeriesHandler(usecases.NewTVSeriesUsecases(memory.NewTvSeriesRepository())),
		ElectronicGamesHandler:    *adpt.NewElectronicGamesHandler(usecases.NewElectronicGamesUsecases(memory.NewElectronicGamesRepository())),
		NonElectronicGamesHandler: *adpt.NewNonElectronicGamesHandler(usecases.NewNonElectronicGamesUsecases(memory.NewNonElectronicGamesRepository())),
	}
}

func main() {
	var server adpt.ServerInterface = newLocalHandler()

	spec, err := openapi3.NewLoader().LoadFromData([]byte(adpt.OpenAPI))
	must(err)
	spec.Servers = nil

	mw := middleware.OapiRequestValidatorWithOptions(spec, &middleware.Options{})

	h := adpt.HandlerWithOptions(server, adpt.StdHTTPServerOptions{
		Middlewares: []adpt.MiddlewareFunc{mw},
	})

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(s.ListenAndServe())
}
