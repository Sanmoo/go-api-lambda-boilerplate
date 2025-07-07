package memory

import (
	"errors"

	h "github.com/Sanmoo/go-api-lambda-boilerplate/adapters/outbound/storage"
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/model"
)

type MoviesRepository struct {
	movies []model.Movie
}

func NewMoviesRepository() *MoviesRepository {
	return &MoviesRepository{
		movies: []model.Movie{
			{Media: model.Media{Title: "Inception", ID: h.Ptr("1")}},
			{Media: model.Media{Title: "The Matrix", ID: h.Ptr("2")}},
			{Media: model.Media{Title: "Interstellar", ID: h.Ptr("3")}},
			{Media: model.Media{Title: "The Shawshank Redemption", ID: h.Ptr("4")}},
		},
	}
}
func (r *MoviesRepository) GetAll() ([]model.Movie, error) {
	return r.movies, nil
}
func (r *MoviesRepository) Create(movie model.Movie) (model.Movie, error) {
	r.movies = append(r.movies, movie)
	return movie, nil
}
func (r *MoviesRepository) Update(movie model.Movie) (model.Movie, error) {
	for i, m := range r.movies {
		if *m.Media.ID == *movie.Media.ID {
			r.movies[i] = movie
			return movie, nil
		}
	}
	return model.Movie{}, errors.New("movie not found")
}

func (r *MoviesRepository) Delete(id string) error {
	for i, m := range r.movies {
		if *m.Media.ID == id {
			r.movies = append(r.movies[:i], r.movies[i+1:]...)
			return nil
		}
	}
	return errors.New("movie not found")
}

func (r *MoviesRepository) GetByID(id string) (model.Movie, error) {
	for _, m := range r.movies {
		if *m.Media.ID == id {
			return m, nil
		}
	}
	return model.Movie{}, errors.New("movie not found")
}
