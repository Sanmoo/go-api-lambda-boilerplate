package http

import (
	"net/http"
)

type BaseHandler struct {
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func notImplemented(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (BaseHandler) BooksList(w http.ResponseWriter, r *http.Request, params BooksListParams) {
	notImplemented(w, r)
}

// (POST /books)
func (BaseHandler) BooksCreate(w http.ResponseWriter, r *http.Request) {
	notImplemented(w, r)
}

// (DELETE /books/{id})
func (BaseHandler) BooksDelete(w http.ResponseWriter, r *http.Request, id string) {
	notImplemented(w, r)
}

// (GET /books/{id})
func (BaseHandler) BooksRead(w http.ResponseWriter, r *http.Request, id string) {
	notImplemented(w, r)
}

// (PUT /books/{id})
func (BaseHandler) BooksPut(w http.ResponseWriter, r *http.Request, id string) {
	notImplemented(w, r)
}

// (GET /electronic-games)
func (BaseHandler) ElectronicGamesList(w http.ResponseWriter, r *http.Request, params ElectronicGamesListParams) {
	notImplemented(w, r)
}

// (POST /electronic-games)
func (BaseHandler) ElectronicGamesCreate(w http.ResponseWriter, r *http.Request) {
	notImplemented(w, r)
}

// (DELETE /electronic-games/{id})
func (BaseHandler) ElectronicGamesDelete(w http.ResponseWriter, r *http.Request, id string) {
	notImplemented(w, r)
}

// (GET /electronic-games/{id})
func (BaseHandler) ElectronicGamesRead(w http.ResponseWriter, r *http.Request, id string) {
	notImplemented(w, r)
}

// (PUT /electronic-games/{id})
func (BaseHandler) ElectronicGamesPut(w http.ResponseWriter, r *http.Request, id string) {
	notImplemented(w, r)
}

// (GET /movies)
func (BaseHandler) MoviesList(w http.ResponseWriter, r *http.Request, params MoviesListParams) {
	notImplemented(w, r)
}

// (POST /movies)
func (BaseHandler) MoviesCreate(w http.ResponseWriter, r *http.Request) {
	notImplemented(w, r)
}

// (DELETE /movies/{id})
func (BaseHandler) MoviesDelete(w http.ResponseWriter, r *http.Request, id string) {
	notImplemented(w, r)
}

// (GET /movies/{id})
func (BaseHandler) MoviesRead(w http.ResponseWriter, r *http.Request, id string) {
	notImplemented(w, r)
}

// (PUT /movies/{id})
func (BaseHandler) MoviesPut(w http.ResponseWriter, r *http.Request, id string) {
	notImplemented(w, r)
}

// (GET /non-electronic-games)
func (BaseHandler) NonElectronicGamesList(w http.ResponseWriter, r *http.Request, params NonElectronicGamesListParams) {
	notImplemented(w, r)
}

// (POST /non-electronic-games)
func (BaseHandler) NonElectronicGamesCreate(w http.ResponseWriter, r *http.Request) {
	notImplemented(w, r)
}

// (DELETE /non-electronic-games/{id})
func (BaseHandler) NonElectronicGamesDelete(w http.ResponseWriter, r *http.Request, id string) {
	notImplemented(w, r)
}

// (GET /non-electronic-games/{id})
func (BaseHandler) NonElectronicGamesRead(w http.ResponseWriter, r *http.Request, id string) {
	notImplemented(w, r)
}

// (PUT /non-electronic-games/{id})
func (BaseHandler) NonElectronicGamesPut(w http.ResponseWriter, r *http.Request, id string) {
	notImplemented(w, r)
}

// (GET /tv-series)
func (BaseHandler) TVSeriesList(w http.ResponseWriter, r *http.Request, params TVSeriesListParams) {
	notImplemented(w, r)
}

// (POST /tv-series)
func (BaseHandler) TVSeriesCreate(w http.ResponseWriter, r *http.Request) {
	notImplemented(w, r)
}

// (DELETE /tv-series/{id})
func (BaseHandler) TVSeriesDelete(w http.ResponseWriter, r *http.Request, id string) {
	notImplemented(w, r)
}

// (GET /tv-series/{id})
func (BaseHandler) TVSeriesRead(w http.ResponseWriter, r *http.Request, id string) {
	notImplemented(w, r)
}

// (PUT /tv-series/{id})
func (BaseHandler) TVSeriesPut(w http.ResponseWriter, r *http.Request, id string) {
	notImplemented(w, r)
}
