package http

import (
	"encoding/json"
	"net/http"

	"github.com/Sanmoo/go-api-lambda-boilerplate/core/usecases"
)

type MoviesListHandler struct {
}

func (MoviesListHandler) MoviesList(w http.ResponseWriter, r *http.Request, params MoviesListParams) {
	movies, _ := usecases.ListMovies()
	response := make([]Movie, len(movies))

	for i := range movies {
		response[i] = Movie{
			Director: &movies[i].Director,
			Genre:    &movies[i].Genre,
			Id:       movies[i].ID,
			Title:    movies[i].Title,
		}

		if movies[i].Rating != nil {
			rating := int32(*movies[i].Rating)
			response[i].Rating = &rating
		}
	}

	jsonRes, _ := json.Marshal(response)
	w.Write([]byte(jsonRes))
}
