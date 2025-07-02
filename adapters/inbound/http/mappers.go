package http

import "github.com/Sanmoo/go-api-lambda-boilerplate/core/model"

func (b *Book) ToModel() model.Book {
	return model.Book{
		Media: model.Media{
			Title:  b.Title,
			ID:     b.Id,
			Rating: b.Rating,
		},
		Author: *b.Author,
		Genre:  *b.Genre,
		Status: (*model.BookStatus)(b.Status),
	}
}

func (b *TVSeries) ToModel() model.TVSeries {
	return model.TVSeries{
		Media: model.Media{
			Title:  b.Title,
			ID:     b.Id,
			Rating: b.Rating,
		},
		Seasons:  b.Seasons,
		Finished: b.Finished,
		Status:   (*model.TVSeriesStatus)(b.Status),
	}
}

func TVSeriesFromModel(tv model.TVSeries) *TVSeries {
	return &TVSeries{
		Finished: tv.Finished,
		Id:       tv.ID,
		Rating:   tv.Rating,
		Seasons:  tv.Seasons,
		Status:   (*TVSeriesStatus)(tv.Status),
		Title:    tv.Title,
	}
}
