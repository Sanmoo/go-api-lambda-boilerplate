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

func (u *MoviesUsecases) List() ([]model.Movie, error) {
	return u.repository.GetAll()
}

func (u *MoviesUsecases) Create(movie model.Movie) (model.Movie, error) {
	return u.repository.Create(movie)
}

func (u *MoviesUsecases) Update(movie model.Movie) (model.Movie, error) {
	return u.repository.Update(movie)
}

func (u *MoviesUsecases) Delete(id string) error {
	return u.repository.Delete(id)
}

func (u *MoviesUsecases) GetByID(id string) (model.Movie, error) {
	return u.repository.GetByID(id)
}
