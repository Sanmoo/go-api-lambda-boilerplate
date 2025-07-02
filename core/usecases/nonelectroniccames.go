package usecases

import (
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/model"
)

func NonElectronicCames() ([]model.NonElectronicGame, error) {
	nonelectronicGames := []model.NonElectronicGame{
		{Media: model.Media{Title: "Chess", ID: ptr("1")}},
		{Media: model.Media{Title: "Checkers", ID: ptr("2")}},
		{Media: model.Media{Title: "Backgammon", ID: ptr("3")}},
		{Media: model.Media{Title: "Go", ID: ptr("4")}},
	}

	return nonelectronicGames, nil
}

func CreateNonElectronicGame(game model.NonElectronicGame) (model.NonElectronicGame, error) {
	// This function would typically interact with a database or an external service
	// to create a new non-electronic game. For now, we will return the game as is.

	return game, nil
}

func UpdateNonElectronicGame(game model.NonElectronicGame) (model.NonElectronicGame, error) {
	// This function would typically interact with a database or an external service
	// to update an existing non-electronic game. For now, we will return the game as is.

	return game, nil
}

func DeleteNonElectronicGame(id string) error {
	// This function would typically interact with a database or an external service
	// to delete a non-electronic game by its ID. For now, we will return nil to indicate success.

	return nil
}

func GetNonElectronicGameByID(id string) (model.NonElectronicGame, error) {
	// This function would typically interact with a database or an external service
	// to retrieve a non-electronic game by its ID. For now, we will return a static game.

	game := model.NonElectronicGame{
		Media: model.Media{Title: "Chess", ID: ptr(id)},
	}

	return game, nil
}
