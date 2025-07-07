package usecases

import (
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/model"
)

type BooksUsecases struct {
	repository Repository[model.Book]
}

func NewBooksUsecases(repository Repository[model.Book]) *BooksUsecases {
	return &BooksUsecases{
		repository: repository,
	}
}

func (b *BooksUsecases) ListBooks() ([]model.Book, error) {
	bs, _ := b.repository.GetAll()
	return bs, nil
}

func (b *BooksUsecases) CreateBook(book model.Book) (model.Book, error) {
	return b.repository.Create(book)
}

func (b *BooksUsecases) UpdateBook(book model.Book) (model.Book, error) {
	return b.repository.Update(book)
}

func (b *BooksUsecases) DeleteBook(id string) error {
	return b.repository.Delete(id)
}

func (b *BooksUsecases) GetBookByID(id string) (model.Book, error) {
	return b.repository.GetByID(id)
}
