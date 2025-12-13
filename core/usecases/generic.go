package usecases

type GenericUsecases[T any] struct {
	repo Repository[T]
}

func NewGenericUsecases[T any](repo Repository[T]) *GenericUsecases[T] {
	return &GenericUsecases[T]{repo: repo}
}

func (u *GenericUsecases[T]) List() ([]T, error) {
	return u.repo.GetAll()
}

func (u *GenericUsecases[T]) Create(entity T) (T, error) {
	return u.repo.Create(entity)
}

func (u *GenericUsecases[T]) Update(entity T) (T, error) {
	return u.repo.Update(entity)
}

func (u *GenericUsecases[T]) Delete(id string) error {
	return u.repo.Delete(id)
}

func (u *GenericUsecases[T]) GetByID(id string) (T, error) {
	return u.repo.GetByID(id)
}
