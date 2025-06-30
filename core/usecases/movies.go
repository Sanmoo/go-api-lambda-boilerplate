package usecases

import (
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/model"
)

func ListMovies() ([]model.Movie, error) {
	// This function would typically interact with a database or an external service
	// to retrieve a list of books. For now, we will return a static list.

	movies := []model.Movie{
		{Media: model.Media{Title: "The Great Gatsby", ID: "1", Rating: nil}, Director: "Baz Luhrmann", Genre: "Drama", Status: model.MovieStatusToWatch},
		{Media: model.Media{Title: "1984", ID: "2", Rating: nil}, Director: "Michael Radford", Genre: "Dystopian", Status: model.MovieStatusWatched},
		{Media: model.Media{Title: "To Kill a Mockingbird", ID: "3", Rating: nil}, Director: "Robert Mulligan", Genre: "Drama", Status: model.MovieStatusToWatch},
		{Media: model.Media{Title: "Pride and Prejudice", ID: "4", Rating: nil}, Director: "Joe Wright", Genre: "Romance", Status: model.MovieStatusWatched},
	}

	return movies, nil
}
