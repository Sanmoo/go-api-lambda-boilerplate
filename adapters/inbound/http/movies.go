package http

import (
	"encoding/json"
	"net/http"

	"github.com/Sanmoo/go-api-lambda-boilerplate/core/usecases"
)

type MoviesListHandler struct {
	BaseHandler
}

func (MoviesListHandler) MoviesList(w http.ResponseWriter, r *http.Request, params MoviesListParams) {
	movies, _ := usecases.ListMovies()
	response := make([]Movie, len(movies))

	for i := range movies {
		rating := int32(*movies[i].Rating)
		response[i] = Movie{
			Title:    movies[i].Title,
			Id:       &movies[i].ID,
			Rating:   &rating,
			Director: movies[i].Director,
		}
	}

	jsonRes, _ := json.Marshal(response)
	w.Write([]byte(jsonRes))
}
