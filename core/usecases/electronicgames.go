package usecases

import (
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/model"
)

func ListElectronicGames() ([]model.ElectronicGame, error) {
	// This function would typically interact with a database or an external service
	// to retrieve a list of electronic games. For now, we will return a static list.

	id1 := "1"
	id2 := "2"
	id3 := "3"
	id4 := "4"
	games := []model.ElectronicGame{
		{Media: model.Media{Title: "The Legend of Zelda: Breath of the Wild", ID: &id1}},
		{Media: model.Media{Title: "Super Mario Odyssey", ID: &id2}},
		{Media: model.Media{Title: "God of War", ID: &id3}},
		{Media: model.Media{Title: "The Witcher 3: Wild Hunt", ID: &id4}},
	}

	return games, nil
}

func CreateElectronicGame(game model.ElectronicGame) (model.ElectronicGame, error) {
	// This function would typically interact with a database or an external service
	// to create a new electronic game. For now, we will return the game as is.

	return game, nil
}

func UpdateElectronicGame(game model.ElectronicGame) (model.ElectronicGame, error) {
	// This function would typically interact with a database or an external service
	// to update an existing electronic game. For now, we will return the game as is.

	return game, nil
}

func DeleteElectronicGame(id string) error {
	// This function would typically interact with a database or an external service
	// to delete an electronic game by its ID. For now, we will return nil to indicate success.

	return nil
}

func GetElectronicGameByID(id string) (model.ElectronicGame, error) {
	// This function would typically interact with a database or an external service
	// to retrieve an electronic game by its ID. For now, we will return a static game.

	game := model.ElectronicGame{
		Media: model.Media{Title: "The Legend of Zelda: Breath of the Wild", ID: &id},
	}

	return game, nil
}
