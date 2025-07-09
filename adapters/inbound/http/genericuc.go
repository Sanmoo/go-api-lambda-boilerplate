package http

type GenericUseCase[T any] interface {
	List() ([]T, error)
	Create(entity T) (T, error)
	Update(entity T) (T, error)
	Delete(id string) error
	GetByID(id string) (T, error)
}
