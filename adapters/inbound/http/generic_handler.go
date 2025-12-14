package http

import (
	"net/http"
	"reflect"

	"github.com/Sanmoo/go-api-lambda-boilerplate/core/usecases"
)

// GenericCRUDHandler provides generic CRUD HTTP handlers for any domain.
// T is the domain model type (e.g., model.Book).
// U is the HTTP DTO type (e.g., Book).
// P is the list parameters type (e.g., BooksListParams).
type GenericCRUDHandler[T any, U any, P any] struct {
	usecases  *usecases.GenericUsecases[T]
	toModel   func(*U) (*T, error)
	fromModel func(*T) *U
}

// NewGenericCRUDHandler creates a new generic CRUD handler.
func NewGenericCRUDHandler[T any, U any, P any](
	usecases *usecases.GenericUsecases[T],
	toModel func(*U) (*T, error),
	fromModel func(*T) *U,
) *GenericCRUDHandler[T, U, P] {
	return &GenericCRUDHandler[T, U, P]{
		usecases:  usecases,
		toModel:   toModel,
		fromModel: fromModel,
	}
}

// List handles GET requests for listing entities.
func (h *GenericCRUDHandler[T, U, P]) List(w http.ResponseWriter, r *http.Request, params P) {
	entities, err := h.usecases.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := make([]U, len(entities))
	for i := range entities {
		response[i] = *h.fromModel(&entities[i])
	}

	respondWithJSON(responseData{
		data:       response,
		usecaseErr: nil,
		w:          w,
		statusCode: http.StatusOK,
	})
}

// Create handles POST requests for creating entities.
func (h *GenericCRUDHandler[T, U, P]) Create(w http.ResponseWriter, r *http.Request) {
	reqDTO, err := unmarshalFromReq[U](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	model, err := h.toModel(reqDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	created, err := h.usecases.Create(*model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(responseData{
		data:       h.fromModel(&created),
		usecaseErr: nil,
		w:          w,
		statusCode: http.StatusCreated,
	})
}

// Read handles GET requests for reading a single entity by ID.
func (h *GenericCRUDHandler[T, U, P]) Read(w http.ResponseWriter, r *http.Request, id string) {
	entity, err := h.usecases.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(responseData{
		data:       h.fromModel(&entity),
		usecaseErr: nil,
		w:          w,
		statusCode: http.StatusOK,
	})
}

// Update handles PUT requests for updating entities.
func (h *GenericCRUDHandler[T, U, P]) Update(w http.ResponseWriter, r *http.Request, id string) {
	reqDTO, err := unmarshalFromReq[U](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set ID from path parameter using reflection
	setID(reqDTO, id)

	model, err := h.toModel(reqDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updated, err := h.usecases.Update(*model)

	respondWithJSON(responseData{
		data:       h.fromModel(&updated),
		usecaseErr: err,
		w:          w,
		statusCode: http.StatusOK,
	})
}

// Delete handles DELETE requests for deleting entities.
func (h *GenericCRUDHandler[T, U, P]) Delete(w http.ResponseWriter, r *http.Request, id string) {
	err := h.usecases.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// setID uses reflection to set the Id field on a DTO.
func setID(dto interface{}, id string) {
	val := reflect.ValueOf(dto)
	if val.Kind() != reflect.Ptr {
		return
	}

	elem := val.Elem()
	if elem.Kind() != reflect.Struct {
		return
	}

	field := elem.FieldByName("Id")
	if !field.IsValid() || !field.CanSet() {
		return
	}

	if field.Type().String() == "*string" {
		field.Set(reflect.ValueOf(&id))
	}
}

// ToModelAdapter adapts a ToModel function that doesn't return error.
func ToModelAdapter[T any, U any](f func(*U) *T) func(*U) (*T, error) {
	return func(u *U) (*T, error) {
		return f(u), nil
	}
}
