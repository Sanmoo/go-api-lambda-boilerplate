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

func (u *TVSeriesUsecases) ListTVSeries() ([]model.TVSeries, error) {
	return u.repository.GetAll()
}

func (u *TVSeriesUsecases) CreateTVSeries(tvSeries model.TVSeries) (model.TVSeries, error) {
	return u.repository.Create(tvSeries)
}

func (u *TVSeriesUsecases) UpdateTVSeries(tvSeries model.TVSeries) (model.TVSeries, error) {
	return u.repository.Update(tvSeries)
}

func (u *TVSeriesUsecases) DeleteTVSeries(id string) error {
	return u.repository.Delete(id)
}

func (u *TVSeriesUsecases) GetTVSeriesByID(id string) (model.TVSeries, error) {
	return u.repository.GetByID(id)
}
