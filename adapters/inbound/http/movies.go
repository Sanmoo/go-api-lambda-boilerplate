package http

import (
	"encoding/json"
	"net/http"

	"github.com/Sanmoo/go-api-lambda-boilerplate/core/usecases"
)

type MoviesHandler struct {
	usecases *usecases.MoviesUsecases
}

func NewMoviesHandler(usecases *usecases.MoviesUsecases) *MoviesHandler {
	return &MoviesHandler{
		usecases: usecases,
	}
}

func (h *MoviesHandler) MoviesList(w http.ResponseWriter, r *http.Request, params MoviesListParams) {
	movies, _ := h.usecases.ListMovies()
	response := make([]Movie, len(movies))

	for i, movie := range movies {
		response[i] = *MovieFromModel(&movie)
	}

	jsonRes, _ := json.Marshal(response)
	w.Write([]byte(jsonRes))
}

func (h *MoviesHandler) MoviesCreate(w http.ResponseWriter, r *http.Request) {
	requestMovie, _ := unmarshalFromReq[Movie](r)
	modelMovie, err := requestMovie.ToModel()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdMovie, err := h.usecases.CreateMovie(*modelMovie)

	respondWithJSON(responseData{
		data:       MovieFromModel(&createdMovie),
		usecaseErr: err,
		w:          w,
		statusCode: http.StatusCreated,
	})
}

func (h *MoviesHandler) MoviesRead(w http.ResponseWriter, r *http.Request, id string) {
	movie, err := h.usecases.GetMovieByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(responseData{
		data:       MovieFromModel(&movie),
		usecaseErr: nil,
		w:          w,
		statusCode: http.StatusOK,
	})
}

func (h *MoviesHandler) MoviesPut(w http.ResponseWriter, r *http.Request, id string) {
	requestMovie, _ := unmarshalFromReq[Movie](r)
	requestMovie.Id = &id
	modelMovie, _ := requestMovie.ToModel()
	updatedMovie, err := h.usecases.UpdateMovie(*modelMovie)

	respondWithJSON(responseData{
		data:       MovieFromModel(&updatedMovie),
		usecaseErr: err,
		w:          w,
		statusCode: http.StatusOK,
	})
}

func (h *MoviesHandler) MoviesDelete(w http.ResponseWriter, r *http.Request, id string) {
	err := h.usecases.DeleteMovie(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
