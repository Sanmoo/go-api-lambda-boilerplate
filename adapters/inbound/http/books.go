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
		response[i] = *BookFromModel(&books[i])
	}

	jsonRes, _ := json.Marshal(response)
	w.Write([]byte(jsonRes))
}

func (BooksHandler) BooksCreate(w http.ResponseWriter, r *http.Request) {
	var requestBook Book
	requestPayload, _ := io.ReadAll(r.Body)
	json.Unmarshal(requestPayload, &requestBook)

	createdBook, _ := usecases.CreateBook(*requestBook.ToModel())
	jsonRes, _ := json.Marshal(BookFromModel(&createdBook))
	w.Write([]byte(jsonRes))
}

func (BooksHandler) BooksGet(w http.ResponseWriter, r *http.Request, id string) {
	book, err := usecases.GetBookByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(responseData{
		data:       BookFromModel(&book),
		usecaseErr: nil,
		w:          w,
		statusCode: http.StatusOK,
	})
}

func (BooksHandler) BooksUpdate(w http.ResponseWriter, r *http.Request, id string) {
	requestBook, _ := unmarshalFromReq[Book](r)
	requestBook.Id = &id
	modelBook := requestBook.ToModel()
	updatedBook, err := usecases.UpdateBook(*modelBook)

	respondWithJSON(responseData{
		data:       BookFromModel(&updatedBook),
		usecaseErr: err,
		w:          w,
		statusCode: http.StatusOK,
	})
}

func (BooksHandler) BooksDelete(w http.ResponseWriter, r *http.Request, id string) {
	err := usecases.DeleteBook(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
