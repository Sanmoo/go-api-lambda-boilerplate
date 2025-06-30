package http

import (
	"encoding/json"
	"net/http"

	"github.com/Sanmoo/go-api-lambda-boilerplate/core/usecases"
)

type BooksListHandler struct {
	BaseHandler
}

func (BooksListHandler) BooksList(w http.ResponseWriter, r *http.Request, params BooksListParams) {
	books, _ := usecases.ListBooks()
	response := make([]Book, len(books))

	for i := range books {
		response[i] = Book{
			Title: books[i].Title,
			Id:    &books[i].ID,
		}
	}

	jsonRes, _ := json.Marshal(response)
	w.Write([]byte(jsonRes))
}
