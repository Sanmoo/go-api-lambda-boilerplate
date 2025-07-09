package usecases

import (
	"reflect"
	"testing"

	"github.com/Sanmoo/go-api-lambda-boilerplate/core/mocks"
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/model"
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/usecases/testhelpers"
	"go.uber.org/mock/gomock"
)

func buildNonElectronicGamesSut(t *testing.T) (*NonElectronicGamesUsecases, *mocks.MockRepository[model.NonElectronicGame]) {
	ctrl := gomock.NewController(t)
	fakeRepo := mocks.NewMockRepository[model.NonElectronicGame](ctrl)
	usecases := NewNonElectronicGamesUsecases(fakeRepo)
	return usecases, fakeRepo
}

func TestCreatesANonElectronicGamesUsecases(t *testing.T) {
	usecases, _ := buildNonElectronicGamesSut(t)

	if usecases == nil {
		t.Error("Expected NonElectronicGamesUsecases to be created, but got nil")
	}
}

func TestCreatesANonElectronicGame(t *testing.T) {
	// Arrange
	game := testhelpers.NonElectronicGameFactory.Build()
	sut, repo := buildNonElectronicGamesSut(t)
	repo.EXPECT().
		Create(gomock.Eq(game)).
		Return(game, nil).
		Times(1)

	// Act
	created, err := sut.CreateNonElectronicGame(game)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if created != game {
		t.Errorf("Expected created game to be %v, but got %v", game, created)
	}
}

func TestUpdatesANonElectronicGame(t *testing.T) {
	// Arrange
	game := testhelpers.NonElectronicGameFactory.Build()
	sut, repo := buildNonElectronicGamesSut(t)
	repo.EXPECT().
		Update(gomock.Eq(game)).
		Return(game, nil).
		Times(1)

	// Act
	updated, err := sut.UpdateNonElectronicGame(game)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if updated != game {
		t.Errorf("Expected updated game to be %v, but got %v", game, updated)
	}
}

func TestListNonElectronicGames(t *testing.T) {
	// Arrange
	games := testhelpers.NonElectronicGameFactory.Batch(2)
	sut, repo := buildNonElectronicGamesSut(t)
	repo.EXPECT().
		GetAll().
		Return(games, nil).
		Times(1)

	// Act
	list, err := sut.ListNonElectronicGames()

	// Assert
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(list, games) {
		t.Errorf("Expected list of games to be %v, but got %v", games, list)
	}
}

func TestDeleteNonElectronicGame(t *testing.T) {
	// Arrange
	game := testhelpers.NonElectronicGameFactory.Build()
	sut, repo := buildNonElectronicGamesSut(t)
	repo.EXPECT().
		Delete(gomock.Eq(*game.ID)).
		Return(nil).
		Times(1)

	// Act
	err := sut.DeleteNonElectronicGame(*game.ID)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

func TestGetNonElectronicGameByID(t *testing.T) {
	// Arrange
	game := testhelpers.NonElectronicGameFactory.Build()
	sut, repo := buildNonElectronicGamesSut(t)
	repo.EXPECT().
		GetByID(gomock.Eq(*game.ID)).
		Return(game, nil).
		Times(1)

	// Act
	foundNonElectronicGame, err := sut.GetNonElectronicGameByID(*game.ID)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if foundNonElectronicGame != game {
		t.Errorf("Expected found game to be %v, but got %v", game, foundNonElectronicGame)
	}
}
