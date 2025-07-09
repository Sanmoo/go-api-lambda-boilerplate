package usecases

import (
	"reflect"
	"testing"

	"github.com/Sanmoo/go-api-lambda-boilerplate/core/mocks"
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/model"
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/usecases/testhelpers"
	"go.uber.org/mock/gomock"
)

func buildElectronicGamesSut(t *testing.T) (*ElectronicGamesUsecases, *mocks.MockRepository[model.ElectronicGame]) {
	ctrl := gomock.NewController(t)
	fakeRepo := mocks.NewMockRepository[model.ElectronicGame](ctrl)
	usecases := NewElectronicGamesUsecases(fakeRepo)
	return usecases, fakeRepo
}

func TestCreatesAElectronicGamesUsecases(t *testing.T) {
	usecases, _ := buildElectronicGamesSut(t)

	if usecases == nil {
		t.Error("Expected ElectronicGamesUsecases to be created, but got nil")
	}
}

func TestCreatesAElectronicGame(t *testing.T) {
	// Arrange
	game := testhelpers.ElectronicGameFactory.Build()
	sut, repo := buildElectronicGamesSut(t)
	repo.EXPECT().
		Create(gomock.Eq(game)).
		Return(game, nil).
		Times(1)

	// Act
	created, err := sut.CreateElectronicGame(game)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if created != game {
		t.Errorf("Expected created game to be %v, but got %v", game, created)
	}
}

func TestUpdatesAElectronicGame(t *testing.T) {
	// Arrange
	game := testhelpers.ElectronicGameFactory.Build()
	sut, repo := buildElectronicGamesSut(t)
	repo.EXPECT().
		Update(gomock.Eq(game)).
		Return(game, nil).
		Times(1)

	// Act
	updated, err := sut.UpdateElectronicGame(game)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if updated != game {
		t.Errorf("Expected updated game to be %v, but got %v", game, updated)
	}
}

func TestListElectronicGames(t *testing.T) {
	// Arrange
	games := testhelpers.ElectronicGameFactory.Batch(2)
	sut, repo := buildElectronicGamesSut(t)
	repo.EXPECT().
		GetAll().
		Return(games, nil).
		Times(1)

	// Act
	list, err := sut.ListElectronicGames()

	// Assert
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(list, games) {
		t.Errorf("Expected list of games to be %v, but got %v", games, list)
	}
}

func TestDeleteElectronicGame(t *testing.T) {
	// Arrange
	game := testhelpers.ElectronicGameFactory.Build()
	sut, repo := buildElectronicGamesSut(t)
	repo.EXPECT().
		Delete(gomock.Eq(*game.ID)).
		Return(nil).
		Times(1)

	// Act
	err := sut.DeleteElectronicGame(*game.ID)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

func TestGetElectronicGameByID(t *testing.T) {
	// Arrange
	game := testhelpers.ElectronicGameFactory.Build()
	sut, repo := buildElectronicGamesSut(t)
	repo.EXPECT().
		GetByID(gomock.Eq(*game.ID)).
		Return(game, nil).
		Times(1)

	// Act
	foundElectronicGame, err := sut.GetElectronicGameByID(*game.ID)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if foundElectronicGame != game {
		t.Errorf("Expected found game to be %v, but got %v", game, foundElectronicGame)
	}
}
