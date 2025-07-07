package memory

import (
	"errors"

	h "github.com/Sanmoo/go-api-lambda-boilerplate/adapters/outbound/storage"
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/model"
)

type TvSeriesRepository struct {
	tvSeries []model.TVSeries
}

func NewTvSeriesRepository() *TvSeriesRepository {
	return &TvSeriesRepository{
		tvSeries: []model.TVSeries{
			{Media: model.Media{Title: "Breaking Bad", ID: h.Ptr("1")}},
			{Media: model.Media{Title: "Game of Thrones", ID: h.Ptr("2")}},
			{Media: model.Media{Title: "Stranger Things", ID: h.Ptr("3")}},
			{Media: model.Media{Title: "The Crown", ID: h.Ptr("4")}},
		},
	}
}

func (r *TvSeriesRepository) GetAll() ([]model.TVSeries, error) {
	return r.tvSeries, nil
}
func (r *TvSeriesRepository) Create(tvSeries model.TVSeries) (model.TVSeries, error) {
	r.tvSeries = append(r.tvSeries, tvSeries)
	return tvSeries, nil
}
func (r *TvSeriesRepository) Update(tvSeries model.TVSeries) (model.TVSeries, error) {
	for i, ts := range r.tvSeries {
		if *ts.Media.ID == *tvSeries.Media.ID {
			r.tvSeries[i] = tvSeries
			return tvSeries, nil
		}
	}
	return model.TVSeries{}, errors.New("tv series not found")
}

func (r *TvSeriesRepository) Delete(id string) error {
	for i, ts := range r.tvSeries {
		if *ts.Media.ID == id {
			r.tvSeries = append(r.tvSeries[:i], r.tvSeries[i+1:]...)
			return nil
		}
	}
	return errors.New("tv series not found")
}

func (r *TvSeriesRepository) GetByID(id string) (model.TVSeries, error) {
	for _, ts := range r.tvSeries {
		if *ts.Media.ID == id {
			return ts, nil
		}
	}
	return model.TVSeries{}, errors.New("tv series not found")
}
