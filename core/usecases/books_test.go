package usecases

import (
	"reflect"
	"testing"

	"github.com/Sanmoo/go-api-lambda-boilerplate/core/mocks"
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/model"
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/usecases/testhelpers"
	"go.uber.org/mock/gomock"
)

func buildBooksSut(t *testing.T) (*BooksUsecases, *mocks.MockRepository[model.Book]) {
	ctrl := gomock.NewController(t)
	fakeRepo := mocks.NewMockRepository[model.Book](ctrl)
	usecases := NewBooksUsecases(fakeRepo)
	return usecases, fakeRepo
}

func TestCreatesABooksUsecases(t *testing.T) {
	usecases, _ := buildBooksSut(t)

	if usecases == nil {
		t.Error("Expected BooksUsecases to be created, but got nil")
	}
}

func TestCreatesABook(t *testing.T) {
	// Arrange
	book := testhelpers.BookFactory.Build()
	sut, repo := buildBooksSut(t)
	repo.EXPECT().
		Create(gomock.Eq(book)).
		Return(book, nil).
		Times(1)

	// Act
	created, err := sut.CreateBook(book)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if created != book {
		t.Errorf("Expected created book to be %v, but got %v", book, created)
	}
}

func TestUpdatesABook(t *testing.T) {
	// Arrange
	book := testhelpers.BookFactory.Build()
	sut, repo := buildBooksSut(t)
	repo.EXPECT().
		Update(gomock.Eq(book)).
		Return(book, nil).
		Times(1)

	// Act
	updated, err := sut.UpdateBook(book)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if updated != book {
		t.Errorf("Expected updated book to be %v, but got %v", book, updated)
	}
}

func TestListBooks(t *testing.T) {
	// Arrange
	books := testhelpers.BookFactory.Batch(2)
	sut, repo := buildBooksSut(t)
	repo.EXPECT().
		GetAll().
		Return(books, nil).
		Times(1)

	// Act
	list, err := sut.ListBooks()

	// Assert
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(list, books) {
		t.Errorf("Expected list of books to be %v, but got %v", books, list)
	}
}

func TestDeleteBook(t *testing.T) {
	// Arrange
	book := testhelpers.BookFactory.Build()
	sut, repo := buildBooksSut(t)
	repo.EXPECT().
		Delete(gomock.Eq(*book.ID)).
		Return(nil).
		Times(1)

	// Act
	err := sut.DeleteBook(*book.ID)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

func TestGetBookByID(t *testing.T) {
	// Arrange
	book := testhelpers.BookFactory.Build()
	sut, repo := buildBooksSut(t)
	repo.EXPECT().
		GetByID(gomock.Eq(*book.ID)).
		Return(book, nil).
		Times(1)

	// Act
	foundBook, err := sut.GetBookByID(*book.ID)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if foundBook != book {
		t.Errorf("Expected found book to be %v, but got %v", book, foundBook)
	}
}
