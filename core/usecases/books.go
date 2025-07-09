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

func (b *BooksUsecases) List() ([]model.Book, error) {
	bs, _ := b.repository.GetAll()
	return bs, nil
}

func (b *BooksUsecases) Create(book model.Book) (model.Book, error) {
	return b.repository.Create(book)
}

func (b *BooksUsecases) Update(book model.Book) (model.Book, error) {
	return b.repository.Update(book)
}

func (b *BooksUsecases) Delete(id string) error {
	return b.repository.Delete(id)
}

func (b *BooksUsecases) GetByID(id string) (model.Book, error) {
	return b.repository.GetByID(id)
}
