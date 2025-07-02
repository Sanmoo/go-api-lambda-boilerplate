package http

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Sanmoo/go-api-lambda-boilerplate/core/usecases"
)

type BooksHandler struct {
}

func (BooksHandler) BooksList(w http.ResponseWriter, r *http.Request, params BooksListParams) {
	books, _ := usecases.ListBooks()
	response := make([]Book, len(books))

	for i := range books {
		response[i] = Book{
			Author: &books[i].Author,
			Genre:  &books[i].Genre,
			Id:     books[i].ID,
			Rating: nil,
			Status: nil,
			Title:  books[i].Title,
		}
	}

	jsonRes, _ := json.Marshal(response)
	w.Write([]byte(jsonRes))
}

func (BooksHandler) BooksCreate(w http.ResponseWriter, r *http.Request) {
	var requestBook Book
	requestPayload, _ := io.ReadAll(r.Body)
	json.Unmarshal(requestPayload, &requestBook)

	createdBook, _ := usecases.CreateBook(requestBook.ToModel())
	jsonRes, _ := json.Marshal(createdBook)
	w.Write([]byte(jsonRes))
}
