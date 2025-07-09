package usecases

import (
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/model"
)

type NonElectronicGamesUsecases struct {
	repository Repository[model.NonElectronicGame]
}

func NewNonElectronicGamesUsecases(repository Repository[model.NonElectronicGame]) *NonElectronicGamesUsecases {
	return &NonElectronicGamesUsecases{
		repository: repository,
	}
}

func (u *NonElectronicGamesUsecases) List() ([]model.NonElectronicGame, error) {
	return u.repository.GetAll()
}

func (u *NonElectronicGamesUsecases) Create(game model.NonElectronicGame) (model.NonElectronicGame, error) {
	return u.repository.Create(game)
}

func (u *NonElectronicGamesUsecases) Update(game model.NonElectronicGame) (model.NonElectronicGame, error) {
	return u.repository.Update(game)
}

func (u *NonElectronicGamesUsecases) Delete(id string) error {
	return u.repository.Delete(id)
}

func (u *NonElectronicGamesUsecases) GetByID(id string) (model.NonElectronicGame, error) {
	return u.repository.GetByID(id)
}
