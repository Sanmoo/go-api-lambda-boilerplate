package usecases

import (
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/model"
)

func ptr[T any](v T) *T {
	return &v
}

func ListTVSeries() ([]model.TVSeries, error) {
	tvSeries := []model.TVSeries{
		{Media: model.Media{Title: "The Great Gatsby", ID: ptr("1")}},
		{Media: model.Media{Title: "1984", ID: ptr("2")}},
		{Media: model.Media{Title: "To Kill a Mockingbird", ID: ptr("3")}},
		{Media: model.Media{Title: "Pride and Prejudice", ID: ptr("4")}},
	}

	return tvSeries, nil
}

func CreateTVSeries(tvSeries model.TVSeries) (model.TVSeries, error) {
	// This function would typically interact with a database or an external service
	// to create a new TV series. For now, we will return the TV series as is.

	return tvSeries, nil
}
