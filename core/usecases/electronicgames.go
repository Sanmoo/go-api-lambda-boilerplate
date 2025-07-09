package usecases

import (
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/model"
)

type ElectronicGamesUsecases struct {
	repository Repository[model.ElectronicGame]
}

func NewElectronicGamesUsecases(repository Repository[model.ElectronicGame]) *ElectronicGamesUsecases {
	return &ElectronicGamesUsecases{
		repository: repository,
	}
}

func (u *ElectronicGamesUsecases) List() ([]model.ElectronicGame, error) {
	return u.repository.GetAll()
}

func (u *ElectronicGamesUsecases) Create(game model.ElectronicGame) (model.ElectronicGame, error) {
	return u.repository.Create(game)
}

func (u *ElectronicGamesUsecases) Update(game model.ElectronicGame) (model.ElectronicGame, error) {
	return u.repository.Update(game)
}

func (u *ElectronicGamesUsecases) Delete(id string) error {
	return u.repository.Delete(id)
}

func (u *ElectronicGamesUsecases) GetByID(id string) (model.ElectronicGame, error) {
	return u.repository.GetByID(id)
}
