package http

import (
	"errors"

	"github.com/Sanmoo/go-api-lambda-boilerplate/core/model"
)

func (b *Book) ToModel() *model.Book {
	return &model.Book{
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

func BookFromModel(book *model.Book) *Book {
	return &Book{
		Id:     book.ID,
		Title:  book.Title,
		Rating: book.Rating,
		Author: &book.Author,
		Genre:  &book.Genre,
		Status: (*BookStatus)(book.Status),
	}
}

func (b *TVSeries) ToModel() *model.TVSeries {
	return &model.TVSeries{
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

func TVSeriesFromModel(tv *model.TVSeries) *TVSeries {
	return &TVSeries{
		Finished: tv.Finished,
		Id:       tv.ID,
		Rating:   tv.Rating,
		Seasons:  tv.Seasons,
		Status:   (*TVSeriesStatus)(tv.Status),
		Title:    tv.Title,
	}
}

func (m *Movie) ToModel() (*model.Movie, error) {
	if m.Director == nil {
		return nil, errors.New("director is required")
	}

	if m.Genre == nil {
		return nil, errors.New("genre is required")
	}

	return &model.Movie{
		Media: model.Media{
			Title:  m.Title,
			ID:     m.Id,
			Rating: m.Rating,
		},
		Director: *m.Director,
		Genre:    *m.Genre,
		Status:   (*model.MovieStatus)(m.Status),
	}, nil
}

func MovieFromModel(movie *model.Movie) *Movie {
	return &Movie{
		Id:       movie.ID,
		Title:    movie.Title,
		Rating:   movie.Rating,
		Director: &movie.Director,
		Genre:    &movie.Genre,
		Status:   (*MovieStatus)(movie.Status),
	}
}

func (e *ElectronicGame) ToModel() (*model.ElectronicGame, error) {
	if (e.Platform == nil) || (e.Genre == nil) {
		return nil, errors.New("platform and genre are required")
	}

	return &model.ElectronicGame{
		Media: model.Media{
			Title:  e.Title,
			ID:     e.Id,
			Rating: e.Rating,
		},
		Platform: *e.Platform,
		Genre:    *e.Genre,
		Status:   (*model.ElectronicGameStatus)(e.Status),
	}, nil
}

func ElectronicGameFromModel(game *model.ElectronicGame) *ElectronicGame {
	return &ElectronicGame{
		Id:       game.ID,
		Title:    game.Title,
		Rating:   game.Rating,
		Platform: &game.Platform,
		Genre:    &game.Genre,
		Status:   (*ElectronicGameStatus)(game.Status),
	}
}

func (n *NonElectronicGame) ToModel() (*model.NonElectronicGame, error) {
	if n.Type == nil || n.Status == nil {
		return nil, errors.New("type and status are required")
	}

	return &model.NonElectronicGame{
		Media: model.Media{
			Title:  n.Title,
			ID:     n.Id,
			Rating: n.Rating,
		},
		Type:   model.NonElectronicGameType(*n.Type),
		Status: model.NonElectronicGameStatus(*n.Status),
	}, nil
}

func NonElectronicGameFromModel(game *model.NonElectronicGame) *NonElectronicGame {
	return &NonElectronicGame{
		Id:     game.ID,
		Title:  game.Title,
		Rating: game.Rating,
		Type:   (*NonElectronicGameType)(&game.Type),
		Status: (*NonElectronicGameStatus)(&game.Status),
	}
}

