package aws

import (
	"net/http"

	adpt "github.com/Sanmoo/go-api-lambda-boilerplate/adapters/inbound/http"
)

type BlankHandler struct {
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func notImplemented(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (BlankHandler) BooksList(w http.ResponseWriter, r *http.Request, params adpt.BooksListParams) {
	notImplemented(w, r)
}

// (POST /books)
func (BlankHandler) BooksCreate(w http.ResponseWriter, r *http.Request) {
	notImplemented(w, r)
}

// (DELETE /books/{id})
func (BlankHandler) BooksDelete(w http.ResponseWriter, r *http.Request, id string) {
	notImplemented(w, r)
}

// (GET /books/{id})
func (BlankHandler) BooksRead(w http.ResponseWriter, r *http.Request, id string) {
	notImplemented(w, r)
}

// (PUT /books/{id})
func (BlankHandler) BooksPut(w http.ResponseWriter, r *http.Request, id string) {
	notImplemented(w, r)
}

// (GET /electronic-games)
func (BlankHandler) ElectronicGamesList(w http.ResponseWriter, r *http.Request, params adpt.ElectronicGamesListParams) {
	notImplemented(w, r)
}

// (POST /electronic-games)
func (BlankHandler) ElectronicGamesCreate(w http.ResponseWriter, r *http.Request) {
	notImplemented(w, r)
}

// (DELETE /electronic-games/{id})
func (BlankHandler) ElectronicGamesDelete(w http.ResponseWriter, r *http.Request, id string) {
	notImplemented(w, r)
}

// (GET /electronic-games/{id})
func (BlankHandler) ElectronicGamesRead(w http.ResponseWriter, r *http.Request, id string) {
	notImplemented(w, r)
}

// (PUT /electronic-games/{id})
func (BlankHandler) ElectronicGamesPut(w http.ResponseWriter, r *http.Request, id string) {
	notImplemented(w, r)
}

// (GET /movies)
func (BlankHandler) MoviesList(w http.ResponseWriter, r *http.Request, params adpt.MoviesListParams) {
	notImplemented(w, r)
}

// (POST /movies)
func (BlankHandler) MoviesCreate(w http.ResponseWriter, r *http.Request) {
	notImplemented(w, r)
}

// (DELETE /movies/{id})
func (BlankHandler) MoviesDelete(w http.ResponseWriter, r *http.Request, id string) {
	notImplemented(w, r)
}

// (GET /movies/{id})
func (BlankHandler) MoviesRead(w http.ResponseWriter, r *http.Request, id string) {
	notImplemented(w, r)
}

// (PUT /movies/{id})
func (BlankHandler) MoviesPut(w http.ResponseWriter, r *http.Request, id string) {
	notImplemented(w, r)
}

// (GET /non-electronic-games)
func (BlankHandler) NonElectronicGamesList(w http.ResponseWriter, r *http.Request, params adpt.NonElectronicGamesListParams) {
	notImplemented(w, r)
}

// (POST /non-electronic-games)
func (BlankHandler) NonElectronicGamesCreate(w http.ResponseWriter, r *http.Request) {
	notImplemented(w, r)
}

// (DELETE /non-electronic-games/{id})
func (BlankHandler) NonElectronicGamesDelete(w http.ResponseWriter, r *http.Request, id string) {
	notImplemented(w, r)
}

// (GET /non-electronic-games/{id})
func (BlankHandler) NonElectronicGamesRead(w http.ResponseWriter, r *http.Request, id string) {
	notImplemented(w, r)
}

// (PUT /non-electronic-games/{id})
func (BlankHandler) NonElectronicGamesPut(w http.ResponseWriter, r *http.Request, id string) {
	notImplemented(w, r)
}

// (GET /tv-series)
func (BlankHandler) TVSeriesList(w http.ResponseWriter, r *http.Request, params adpt.TVSeriesListParams) {
	notImplemented(w, r)
}

// (POST /tv-series)
func (BlankHandler) TVSeriesCreate(w http.ResponseWriter, r *http.Request) {
	notImplemented(w, r)
}

// (DELETE /tv-series/{id})
func (BlankHandler) TVSeriesDelete(w http.ResponseWriter, r *http.Request, id string) {
	notImplemented(w, r)
}

// (GET /tv-series/{id})
func (BlankHandler) TVSeriesRead(w http.ResponseWriter, r *http.Request, id string) {
	notImplemented(w, r)
}

// (PUT /tv-series/{id})
func (BlankHandler) TVSeriesPut(w http.ResponseWriter, r *http.Request, id string) {
	notImplemented(w, r)
}

type BaseHandler struct {
	BlankHandler
}

type BooksHandler struct {
	BaseHandler
	adpt.BooksHandler
}

type ElectronicGamesHandler struct {
	BaseHandler
	adpt.ElectronicGamesHandler
}

type NonElectronicGamesHandler struct {
	BaseHandler
	adpt.NonElectronicGamesHandler
}

type MoviesHandler struct {
	BaseHandler
	adpt.MoviesHandler
}

type TVSeriesHandler struct {
	BaseHandler
	adpt.TvSeriesHandler
}
