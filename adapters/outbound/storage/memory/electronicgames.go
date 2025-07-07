package memory

import (
	"errors"

	h "github.com/Sanmoo/go-api-lambda-boilerplate/adapters/outbound/storage"
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/model"
)

type ElectronicGamesRepository struct {
	games []model.ElectronicGame
}

func NewElectronicGamesRepository() *ElectronicGamesRepository {
	return &ElectronicGamesRepository{
		games: []model.ElectronicGame{
			{Media: model.Media{Title: "The Legend of Zelda: Breath of the Wild", ID: h.Ptr("1")}},
			{Media: model.Media{Title: "Super Mario Odyssey", ID: h.Ptr("2")}},
			{Media: model.Media{Title: "God of War", ID: h.Ptr("3")}},
			{Media: model.Media{Title: "The Witcher 3: Wild Hunt", ID: h.Ptr("4")}},
		},
	}
}

func (r *ElectronicGamesRepository) GetAll() ([]model.ElectronicGame, error) {
	return r.games, nil
}

func (r *ElectronicGamesRepository) Create(game model.ElectronicGame) (model.ElectronicGame, error) {
	r.games = append(r.games, game)
	return game, nil
}

func (r *ElectronicGamesRepository) Update(game model.ElectronicGame) (model.ElectronicGame, error) {
	for i, g := range r.games {
		if *g.Media.ID == *game.Media.ID {
			r.games[i] = game
			return game, nil
		}
	}
	return model.ElectronicGame{}, errors.New("electronic game not found")
}

func (r *ElectronicGamesRepository) Delete(id string) error {
	for i, g := range r.games {
		if *g.Media.ID == id {
			r.games = append(r.games[:i], r.games[i+1:]...)
			return nil
		}
	}
	return errors.New("electronic game not found")
}

func (r *ElectronicGamesRepository) GetByID(id string) (model.ElectronicGame, error) {
	for _, g := range r.games {
		if *g.Media.ID == id {
			return g, nil
		}
	}
	return model.ElectronicGame{}, errors.New("electronic game not found")
}
