package http

import (
	"encoding/json"
	"net/http"

	_ "embed"

	cases "github.com/Sanmoo/go-api-lambda-boilerplate/core/usecases"
)

type BooksHandler struct {
	usecases *cases.BooksUsecases
}

func NewBooksHandler(usecases *cases.BooksUsecases) *BooksHandler {
	return &BooksHandler{
		usecases: usecases,
	}
}

func (b *BooksHandler) BooksList(w http.ResponseWriter, r *http.Request, params BooksListParams) {
	books, err := b.usecases.ListBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := make([]Book, len(books))
	for i := range books {
		response[i] = *BookFromModel(&books[i])
	}
	jsonRes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

func (b *BooksHandler) BooksCreate(w http.ResponseWriter, r *http.Request) {
	requestBook, err := unmarshalFromReq[Book](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdBook, err := b.usecases.CreateBook(*requestBook.ToModel())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonRes, err := json.Marshal(BookFromModel(&createdBook))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

func (b *BooksHandler) BooksRead(w http.ResponseWriter, r *http.Request, id string) {
	book, err := b.usecases.GetBookByID(id)

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

func (b *BooksHandler) BooksPut(w http.ResponseWriter, r *http.Request, id string) {
	requestBook, _ := unmarshalFromReq[Book](r)
	requestBook.Id = &id
	modelBook := requestBook.ToModel()
	updatedBook, err := b.usecases.UpdateBook(*modelBook)

	respondWithJSON(responseData{
		data:       BookFromModel(&updatedBook),
		usecaseErr: err,
		w:          w,
		statusCode: http.StatusOK,
	})
}

func (b *BooksHandler) BooksDelete(w http.ResponseWriter, r *http.Request, id string) {
	err := b.usecases.DeleteBook(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
