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

func (u *ElectronicGamesUsecases) ListElectronicGames() ([]model.ElectronicGame, error) {
	return u.repository.GetAll()
}

func (u *ElectronicGamesUsecases) CreateElectronicGame(game model.ElectronicGame) (model.ElectronicGame, error) {
	return u.repository.Create(game)
}

func (u *ElectronicGamesUsecases) UpdateElectronicGame(game model.ElectronicGame) (model.ElectronicGame, error) {
	return u.repository.Update(game)
}

func (u *ElectronicGamesUsecases) DeleteElectronicGame(id string) error {
	return u.repository.Delete(id)
}

func (u *ElectronicGamesUsecases) GetElectronicGameByID(id string) (model.ElectronicGame, error) {
	return u.repository.GetByID(id)
}
