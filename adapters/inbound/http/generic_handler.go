package http

import (
	"encoding/json"
	"io"
	"net/http"

	_ "embed"
)

type GenericHandler[T any] struct {
	usecases GenericUseCase[T]
}

func NewGenericHandler[T any](usecases GenericUseCase[T]) *GenericHandler[T] {
	return &GenericHandler[T]{
		usecases: usecases,
	}
}

func (h *GenericHandler[T]) BooksList(w http.ResponseWriter, r *http.Request, params BooksListParams) {
	books, _ := h.usecases.List()
	response := make([]T, len(books))

	for i := range books {
		response[i] = *BookFromModel(&books[i])
	}

	jsonRes, _ := json.Marshal(response)
	w.Write([]byte(jsonRes))
}

func (h *GenericHandler) BooksCreate(w http.ResponseWriter, r *http.Request) {
	var requestBook Book
	requestPayload, _ := io.ReadAll(r.Body)
	json.Unmarshal(requestPayload, &requestBook)

	createdBook, _ := h.usecases.Create(*requestBook.ToModel())
	jsonRes, _ := json.Marshal(BookFromModel(&createdBook))
	w.Write([]byte(jsonRes))
}

func (h *GenericHandler) BooksRead(w http.ResponseWriter, r *http.Request, id string) {
	book, err := h.usecases.GetByID(id)

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

func (h *GenericHandler) BooksPut(w http.ResponseWriter, r *http.Request, id string) {
	requestBook, _ := unmarshalFromReq[Book](r)
	requestBook.Id = &id
	modelBook := requestBook.ToModel()
	updatedBook, err := h.usecases.Update(*modelBook)

	respondWithJSON(responseData{
		data:       BookFromModel(&updatedBook),
		usecaseErr: err,
		w:          w,
		statusCode: http.StatusOK,
	})
}

func (h *GenericHandler) BooksDelete(w http.ResponseWriter, r *http.Request, id string) {
	err := h.usecases.Delete(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
