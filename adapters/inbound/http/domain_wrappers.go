package http

import (
	"net/http"

	"github.com/Sanmoo/go-api-lambda-boilerplate/core/model"
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/usecases"
)

// MoviesHandler handles HTTP requests for movies.
type MoviesHandler struct {
	*GenericCRUDHandler[model.Movie, Movie, MoviesListParams]
}

// NewMoviesHandler creates a new movies handler.
func NewMoviesHandler(usecases *usecases.MoviesUsecases) *MoviesHandler {
	return &MoviesHandler{
		GenericCRUDHandler: NewGenericCRUDHandler[model.Movie, Movie, MoviesListParams](
			usecases.GenericUsecases,
			(*Movie).ToModel,
			MovieFromModel,
		),
	}
}

// MoviesList handles GET /movies.
func (h *MoviesHandler) MoviesList(w http.ResponseWriter, r *http.Request, params MoviesListParams) {
	h.List(w, r, params)
}

// MoviesCreate handles POST /movies.
func (h *MoviesHandler) MoviesCreate(w http.ResponseWriter, r *http.Request) {
	h.Create(w, r)
}

// MoviesRead handles GET /movies/{id}.
func (h *MoviesHandler) MoviesRead(w http.ResponseWriter, r *http.Request, id string) {
	h.Read(w, r, id)
}

// MoviesPut handles PUT /movies/{id}.
func (h *MoviesHandler) MoviesPut(w http.ResponseWriter, r *http.Request, id string) {
	h.Update(w, r, id)
}

// MoviesDelete handles DELETE /movies/{id}.
func (h *MoviesHandler) MoviesDelete(w http.ResponseWriter, r *http.Request, id string) {
	h.Delete(w, r, id)
}

// TvSeriesHandler handles HTTP requests for TV series.
type TvSeriesHandler struct {
	*GenericCRUDHandler[model.TVSeries, TVSeries, TVSeriesListParams]
}

// NewTvSeriesHandler creates a new TV series handler.
func NewTvSeriesHandler(usecases *usecases.TVSeriesUsecases) *TvSeriesHandler {
	return &TvSeriesHandler{
		GenericCRUDHandler: NewGenericCRUDHandler[model.TVSeries, TVSeries, TVSeriesListParams](
			usecases.GenericUsecases,
			ToModelAdapter((*TVSeries).ToModel),
			TVSeriesFromModel,
		),
	}
}

// TVSeriesList handles GET /tv-series.
func (h *TvSeriesHandler) TVSeriesList(w http.ResponseWriter, r *http.Request, params TVSeriesListParams) {
	h.List(w, r, params)
}

// TVSeriesCreate handles POST /tv-series.
func (h *TvSeriesHandler) TVSeriesCreate(w http.ResponseWriter, r *http.Request) {
	h.Create(w, r)
}

// TVSeriesRead handles GET /tv-series/{id}.
func (h *TvSeriesHandler) TVSeriesRead(w http.ResponseWriter, r *http.Request, id string) {
	h.Read(w, r, id)
}

// TVSeriesPut handles PUT /tv-series/{id}.
func (h *TvSeriesHandler) TVSeriesPut(w http.ResponseWriter, r *http.Request, id string) {
	h.Update(w, r, id)
}

// TVSeriesDelete handles DELETE /tv-series/{id}.
func (h *TvSeriesHandler) TVSeriesDelete(w http.ResponseWriter, r *http.Request, id string) {
	h.Delete(w, r, id)
}

// ElectronicGamesHandler handles HTTP requests for electronic games.
type ElectronicGamesHandler struct {
	*GenericCRUDHandler[model.ElectronicGame, ElectronicGame, ElectronicGamesListParams]
}

// NewElectronicGamesHandler creates a new electronic games handler.
func NewElectronicGamesHandler(usecases *usecases.ElectronicGamesUsecases) *ElectronicGamesHandler {
	return &ElectronicGamesHandler{
		GenericCRUDHandler: NewGenericCRUDHandler[model.ElectronicGame, ElectronicGame, ElectronicGamesListParams](
			usecases.GenericUsecases,
			(*ElectronicGame).ToModel,
			ElectronicGameFromModel,
		),
	}
}

// ElectronicGamesList handles GET /electronic-games.
func (h *ElectronicGamesHandler) ElectronicGamesList(w http.ResponseWriter, r *http.Request, params ElectronicGamesListParams) {
	h.List(w, r, params)
}

// ElectronicGamesCreate handles POST /electronic-games.
func (h *ElectronicGamesHandler) ElectronicGamesCreate(w http.ResponseWriter, r *http.Request) {
	h.Create(w, r)
}

// ElectronicGamesRead handles GET /electronic-games/{id}.
func (h *ElectronicGamesHandler) ElectronicGamesRead(w http.ResponseWriter, r *http.Request, id string) {
	h.Read(w, r, id)
}

// ElectronicGamesPut handles PUT /electronic-games/{id}.
func (h *ElectronicGamesHandler) ElectronicGamesPut(w http.ResponseWriter, r *http.Request, id string) {
	h.Update(w, r, id)
}

// ElectronicGamesDelete handles DELETE /electronic-games/{id}.
func (h *ElectronicGamesHandler) ElectronicGamesDelete(w http.ResponseWriter, r *http.Request, id string) {
	h.Delete(w, r, id)
}

// NonElectronicGamesHandler handles HTTP requests for non-electronic games.
type NonElectronicGamesHandler struct {
	*GenericCRUDHandler[model.NonElectronicGame, NonElectronicGame, NonElectronicGamesListParams]
}

// NewNonElectronicGamesHandler creates a new non-electronic games handler.
func NewNonElectronicGamesHandler(usecases *usecases.NonElectronicGamesUsecases) *NonElectronicGamesHandler {
	return &NonElectronicGamesHandler{
		GenericCRUDHandler: NewGenericCRUDHandler[model.NonElectronicGame, NonElectronicGame, NonElectronicGamesListParams](
			usecases.GenericUsecases,
			(*NonElectronicGame).ToModel,
			NonElectronicGameFromModel,
		),
	}
}

// NonElectronicGamesList handles GET /non-electronic-games.
func (h *NonElectronicGamesHandler) NonElectronicGamesList(w http.ResponseWriter, r *http.Request, params NonElectronicGamesListParams) {
	h.List(w, r, params)
}

// NonElectronicGamesCreate handles POST /non-electronic-games.
func (h *NonElectronicGamesHandler) NonElectronicGamesCreate(w http.ResponseWriter, r *http.Request) {
	h.Create(w, r)
}

// NonElectronicGamesRead handles GET /non-electronic-games/{id}.
func (h *NonElectronicGamesHandler) NonElectronicGamesRead(w http.ResponseWriter, r *http.Request, id string) {
	h.Read(w, r, id)
}

// NonElectronicGamesPut handles PUT /non-electronic-games/{id}.
func (h *NonElectronicGamesHandler) NonElectronicGamesPut(w http.ResponseWriter, r *http.Request, id string) {
	h.Update(w, r, id)
}

// NonElectronicGamesDelete handles DELETE /non-electronic-games/{id}.
func (h *NonElectronicGamesHandler) NonElectronicGamesDelete(w http.ResponseWriter, r *http.Request, id string) {
	h.Delete(w, r, id)
}
