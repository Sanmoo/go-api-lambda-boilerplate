package usecases

import (
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/model"
)

type TVSeriesUsecases struct {
	repository Repository[model.TVSeries]
}

func NewTVSeriesUsecases(repository Repository[model.TVSeries]) *TVSeriesUsecases {
	return &TVSeriesUsecases{
		repository: repository,
	}
}

func (u *TVSeriesUsecases) List() ([]model.TVSeries, error) {
	return u.repository.GetAll()
}

func (u *TVSeriesUsecases) Create(tvSeries model.TVSeries) (model.TVSeries, error) {
	return u.repository.Create(tvSeries)
}

func (u *TVSeriesUsecases) Update(tvSeries model.TVSeries) (model.TVSeries, error) {
	return u.repository.Update(tvSeries)
}

func (u *TVSeriesUsecases) Delete(id string) error {
	return u.repository.Delete(id)
}

func (u *TVSeriesUsecases) GetByID(id string) (model.TVSeries, error) {
	return u.repository.GetByID(id)
}
