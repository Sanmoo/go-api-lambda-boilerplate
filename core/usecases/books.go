package usecases

import (
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/model"
)

func ListBooks() ([]model.Book, error) {
	// This function would typically interact with a database or an external service
	// to retrieve a list of books. For now, we will return a static list.

	books := []model.Book{
		{Media: model.Media{Title: "The Great Gatsby", ID: "1"}},
		{Media: model.Media{Title: "1984", ID: "2"}},
		{Media: model.Media{Title: "To Kill a Mockingbird", ID: "3"}},
		{Media: model.Media{Title: "Pride and Prejudice", ID: "4"}},
	}

	return books, nil
}
