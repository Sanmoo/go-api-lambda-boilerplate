package aws

import (
	"github.com/Sanmoo/go-api-lambda-boilerplate/adapters/inbound/http"
)

type BaseHandler struct {
	http.BlankHandler
}

type BooksHandler struct {
	BaseHandler
	http.BooksHandler
}

type ElectronicGamesHandler struct {
	BaseHandler
	http.ElectronicGamesHandler
}

type NonElectronicGamesHandler struct {
	BaseHandler
	http.NonElectronicGamesHandler
}

type MoviesHandler struct {
	BaseHandler
	http.MoviesHandler
}

type TVSeriesHandler struct {
	BaseHandler
	http.TvSeriesHandler
}
