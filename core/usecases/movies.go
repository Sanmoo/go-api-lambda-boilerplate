package usecases

import (
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/model"
)

type MoviesUsecases struct {
	repository Repository[model.Movie]
}

func NewMoviesUsecases(repository Repository[model.Movie]) *MoviesUsecases {
	return &MoviesUsecases{
		repository: repository,
	}
}

func (u *MoviesUsecases) ListMovies() ([]model.Movie, error) {
	return u.repository.GetAll()
}

func (u *MoviesUsecases) CreateMovie(movie model.Movie) (model.Movie, error) {
	return u.repository.Create(movie)
}

func (u *MoviesUsecases) UpdateMovie(movie model.Movie) (model.Movie, error) {
	return u.repository.Update(movie)
}

func (u *MoviesUsecases) DeleteMovie(id string) error {
	return u.repository.Delete(id)
}

func (u *MoviesUsecases) GetMovieByID(id string) (model.Movie, error) {
	return u.repository.GetByID(id)
}
