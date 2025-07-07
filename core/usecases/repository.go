package usecases

type Repository[T any] interface {
	GetAll() ([]T, error)
	GetByID(id string) (T, error)
	Create(entity T) (T, error)
	Update(entity T) (T, error)
	Delete(id string) error
}
