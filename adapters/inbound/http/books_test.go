package http

import (
	"testing"

	"github.com/Sanmoo/go-api-lambda-boilerplate/adapters/inbound/http/testhelpers/mocks"
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/model"
	"go.uber.org/mock/gomock"
)

type GUC = mocks.MockGenericUseCase[model.Book]

func buildTestBooksHandler(t *testing.T) (*BooksHandler, *GUC) {
	ctrl := gomock.NewController(t)
	fakeUc := mocks.NewMockGenericUseCase[model.Book](ctrl)
	sut := NewBooksHandler(fakeUc)
	return sut, fakeUc
}

func TestBooksHandlerCreate(t *testing.T) {
	var sut *BooksHandler
	sut, _ = buildTestBooksHandler(t)

	if sut == nil {
		t.Error("Expected BooksHandler to be created, but got nil")
	}
}
