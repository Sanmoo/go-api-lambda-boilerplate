package memory

import (
	"errors"

	h "github.com/Sanmoo/go-api-lambda-boilerplate/adapters/outbound/storage"
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/model"
)

type NonElectronicGamesRepository struct {
	games []model.NonElectronicGame
}

func NewNonElectronicGamesRepository() *NonElectronicGamesRepository {
	return &NonElectronicGamesRepository{
		games: []model.NonElectronicGame{
			{Media: model.Media{Title: "Chess", ID: h.Ptr("1")}},
			{Media: model.Media{Title: "Monopoly", ID: h.Ptr("2")}},
			{Media: model.Media{Title: "Scrabble", ID: h.Ptr("3")}},
			{Media: model.Media{Title: "Settlers of Catan", ID: h.Ptr("4")}},
		},
	}
}

func (r *NonElectronicGamesRepository) GetAll() ([]model.NonElectronicGame, error) {
	return r.games, nil
}
func (r *NonElectronicGamesRepository) Create(game model.NonElectronicGame) (model.NonElectronicGame, error) {
	r.games = append(r.games, game)
	return game, nil
}
func (r *NonElectronicGamesRepository) Update(game model.NonElectronicGame) (model.NonElectronicGame, error) {
	for i, g := range r.games {
		if *g.Media.ID == *game.Media.ID {
			r.games[i] = game
			return game, nil
		}
	}
	return model.NonElectronicGame{}, errors.New("non-electronic game not found")
}
func (r *NonElectronicGamesRepository) Delete(id string) error {
	for i, g := range r.games {
		if *g.Media.ID == id {
			r.games = append(r.games[:i], r.games[i+1:]...)
			return nil
		}
	}
	return errors.New("non-electronic game not found")
}
func (r *NonElectronicGamesRepository) GetByID(id string) (model.NonElectronicGame, error) {
	for _, g := range r.games {
		if *g.Media.ID == id {
			return g, nil
		}
	}
	return model.NonElectronicGame{}, errors.New("non-electronic game not found")
}
