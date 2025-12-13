package memory

import (
	h "github.com/Sanmoo/go-api-lambda-boilerplate/adapters/outbound/storage"
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/model"
)

func NewBooksRepository() *MemoryRepository[model.Book] {
	return NewMemoryRepository(
		func(book model.Book) string {
			if book.Media.ID == nil {
				return ""
			}
			return *book.Media.ID
		},
		[]model.Book{
			{Media: model.Media{Title: "The Great Gatsby", ID: h.Ptr("1")}, Author: "F. Scott Fitzgerald"},
			{Media: model.Media{Title: "1984", ID: h.Ptr("2")}, Author: "George Orwell"},
			{Media: model.Media{Title: "To Kill a Mockingbird", ID: h.Ptr("3")}, Author: "Harper Lee"},
			{Media: model.Media{Title: "Pride and Prejudice", ID: h.Ptr("4")}, Author: "Jane Austen"},
		}...,
	)
}

func NewMoviesRepository() *MemoryRepository[model.Movie] {
	return NewMemoryRepository(
		func(movie model.Movie) string {
			if movie.Media.ID == nil {
				return ""
			}
			return *movie.Media.ID
		},
		[]model.Movie{
			{Media: model.Media{Title: "Inception", ID: h.Ptr("1")}, Director: "Christopher Nolan", Genre: "Sci-Fi"},
			{Media: model.Media{Title: "The Shawshank Redemption", ID: h.Ptr("2")}, Director: "Frank Darabont", Genre: "Drama"},
			{Media: model.Media{Title: "The Godfather", ID: h.Ptr("3")}, Director: "Francis Ford Coppola", Genre: "Crime"},
			{Media: model.Media{Title: "Pulp Fiction", ID: h.Ptr("4")}, Director: "Quentin Tarantino", Genre: "Crime"},
		}...,
	)
}

func NewTVSeriesRepository() *MemoryRepository[model.TVSeries] {
	return NewMemoryRepository(
		func(tv model.TVSeries) string {
			if tv.Media.ID == nil {
				return ""
			}
			return *tv.Media.ID
		},
		[]model.TVSeries{
			{Media: model.Media{Title: "Breaking Bad", ID: h.Ptr("1")}, Seasons: h.Ptr[int32](5), Finished: h.Ptr(true)},
			{Media: model.Media{Title: "Game of Thrones", ID: h.Ptr("2")}, Seasons: h.Ptr[int32](8), Finished: h.Ptr(true)},
			{Media: model.Media{Title: "Stranger Things", ID: h.Ptr("3")}, Seasons: h.Ptr[int32](4), Finished: h.Ptr(false)},
			{Media: model.Media{Title: "The Office", ID: h.Ptr("4")}, Seasons: h.Ptr[int32](9), Finished: h.Ptr(true)},
		}...,
	)
}

func NewElectronicGamesRepository() *MemoryRepository[model.ElectronicGame] {
	return NewMemoryRepository(
		func(game model.ElectronicGame) string {
			if game.Media.ID == nil {
				return ""
			}
			return *game.Media.ID
		},
		[]model.ElectronicGame{
			{Media: model.Media{Title: "The Legend of Zelda: Breath of the Wild", ID: h.Ptr("1")}, Platform: "Nintendo Switch", Genre: "Adventure"},
			{Media: model.Media{Title: "Red Dead Redemption 2", ID: h.Ptr("2")}, Platform: "PlayStation 4", Genre: "Action-Adventure"},
			{Media: model.Media{Title: "Minecraft", ID: h.Ptr("3")}, Platform: "Multi-platform", Genre: "Sandbox"},
			{Media: model.Media{Title: "The Witcher 3: Wild Hunt", ID: h.Ptr("4")}, Platform: "PC", Genre: "RPG"},
		}...,
	)
}

func NewNonElectronicGamesRepository() *MemoryRepository[model.NonElectronicGame] {
	return NewMemoryRepository(
		func(game model.NonElectronicGame) string {
			if game.Media.ID == nil {
				return ""
			}
			return *game.Media.ID
		},
		[]model.NonElectronicGame{
			{Media: model.Media{Title: "Chess", ID: h.Ptr("1")}, Type: model.Board},
			{Media: model.Media{Title: "Poker", ID: h.Ptr("2")}, Type: model.Card},
			{Media: model.Media{Title: "Monopoly", ID: h.Ptr("3")}, Type: model.Board},
			{Media: model.Media{Title: "Dungeons & Dragons", ID: h.Ptr("4")}, Type: model.Other},
		}...,
	)
}
