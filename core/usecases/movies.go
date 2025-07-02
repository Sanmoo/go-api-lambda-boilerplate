package usecases

import (
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/model"
)

func ListMovies() ([]model.Movie, error) {
	var rating5 = int32(5)
	id1 := "1"
	id2 := "2"
	id3 := "3"
	id4 := "4"
	movies := []model.Movie{
		{Media: model.Media{Title: "The Great Gatsby", ID: &id1, Rating: &rating5}, Director: "Baz Luhrmann", Genre: "Drama", Status: ptr(model.MovieStatusToWatch)},
		{Media: model.Media{Title: "1984", ID: &id2, Rating: nil}, Director: "Michael Radford", Genre: "Dystopian", Status: ptr(model.MovieStatusWatched)},
		{Media: model.Media{Title: "To Kill a Mockingbird", ID: &id3, Rating: nil}, Director: "Robert Mulligan", Genre: "Drama", Status: ptr(model.MovieStatusToWatch)},
		{Media: model.Media{Title: "Pride and Prejudice", ID: &id4, Rating: nil}, Director: "Joe Wright", Genre: "Romance", Status: ptr(model.MovieStatusWatched)},
	}

	return movies, nil
}

func CreateMovie(movie model.Movie) (model.Movie, error) {
	// This function would typically interact with a database or an external service
	// to create a new movie. For now, we will return the movie as is.

	return movie, nil
}

func UpdateMovie(movie model.Movie) (model.Movie, error) {
	// This function would typically interact with a database or an external service
	// to update an existing movie. For now, we will return the movie as is.

	return movie, nil
}

func DeleteMovie(id string) error {
	// This function would typically interact with a database or an external service
	// to delete a movie by its ID. For now, we will return nil to indicate success.

	return nil
}

func GetMovieByID(id string) (model.Movie, error) {
	// This function would typically interact with a database or an external service
	// to retrieve a movie by its ID. For now, we will return a static movie.

	movie := model.Movie{
		Media: model.Media{Title: "The Great Gatsby", ID: &id},
	}

	return movie, nil
}
