package usecases

import (
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/model"
)

func ListBooks() ([]model.Book, error) {
	// This function would typically interact with a database or an external service
	// to retrieve a list of books. For now, we will return a static list.

	id1 := "1"
	id2 := "2"
	id3 := "3"
	id4 := "4"
	books := []model.Book{
		{Media: model.Media{Title: "The Great Gatsby", ID: &id1}},
		{Media: model.Media{Title: "1984", ID: &id2}},
		{Media: model.Media{Title: "To Kill a Mockingbird", ID: &id3}},
		{Media: model.Media{Title: "Pride and Prejudice", ID: &id4}},
	}

	return books, nil
}

func CreateBook(book model.Book) (model.Book, error) {
	// This function would typically interact with a database or an external service
	// to create a new book. For now, we will return the book as is.

	return book, nil
}
