package memory

import (
	"errors"

	h "github.com/Sanmoo/go-api-lambda-boilerplate/adapters/outbound/storage"
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/model"
)

type BooksRepository struct {
	books []model.Book
}

func NewBooksRepository() *BooksRepository {
	return &BooksRepository{
		books: []model.Book{
			{Media: model.Media{Title: "The Great Gatsby", ID: h.Ptr("1")}, Author: "F. Scott Fitzgerald"},
			{Media: model.Media{Title: "1984", ID: h.Ptr("2")}, Author: "George Orwell"},
			{Media: model.Media{Title: "To Kill a Mockingbird", ID: h.Ptr("3")}, Author: "Harper Lee"},
			{Media: model.Media{Title: "Pride and Prejudice", ID: h.Ptr("4")}, Author: "Jane Austen"},
		},
	}
}

func (r *BooksRepository) GetAll() ([]model.Book, error) {
	return r.books, nil
}

func (r *BooksRepository) Create(book model.Book) (model.Book, error) {
	r.books = append(r.books, book)
	return book, nil
}

func (r *BooksRepository) Update(book model.Book) (model.Book, error) {
	for i, b := range r.books {
		if *b.Media.ID == *book.Media.ID {
			r.books[i] = book
			return book, nil
		}
	}
	return model.Book{}, errors.New("book not found")
}

func (r *BooksRepository) Delete(id string) error {
	for i, b := range r.books {
		if *b.Media.ID == id {
			r.books = append(r.books[:i], r.books[i+1:]...)
			return nil
		}
	}
	return errors.New("book not found")
}

func (r *BooksRepository) GetByID(id string) (model.Book, error) {
	for _, b := range r.books {
		if *b.Media.ID == id {
			return b, nil
		}
	}
	return model.Book{}, errors.New("book not found")
}
