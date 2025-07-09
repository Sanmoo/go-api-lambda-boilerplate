package usecases

import (
	"reflect"
	"testing"

	"github.com/Sanmoo/go-api-lambda-boilerplate/core/mocks"
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/model"
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/usecases/testhelpers"
	"go.uber.org/mock/gomock"
)

func buildMoviesSut(t *testing.T) (*MoviesUsecases, *mocks.MockRepository[model.Movie]) {
	ctrl := gomock.NewController(t)
	fakeRepo := mocks.NewMockRepository[model.Movie](ctrl)
	usecases := NewMoviesUsecases(fakeRepo)
	return usecases, fakeRepo
}

func TestCreatesAMoviesUsecases(t *testing.T) {
	usecases, _ := buildMoviesSut(t)

	if usecases == nil {
		t.Error("Expected MoviesUsecases to be created, but got nil")
	}
}

func TestCreatesAMovie(t *testing.T) {
	// Arrange
	movie := testhelpers.MovieFactory.Build()
	sut, repo := buildMoviesSut(t)
	repo.EXPECT().
		Create(gomock.Eq(movie)).
		Return(movie, nil).
		Times(1)

	// Act
	created, err := sut.CreateMovie(movie)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if created != movie {
		t.Errorf("Expected created movie to be %v, but got %v", movie, created)
	}
}

func TestUpdatesAMovie(t *testing.T) {
	// Arrange
	movie := testhelpers.MovieFactory.Build()
	sut, repo := buildMoviesSut(t)
	repo.EXPECT().
		Update(gomock.Eq(movie)).
		Return(movie, nil).
		Times(1)

	// Act
	updated, err := sut.UpdateMovie(movie)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if updated != movie {
		t.Errorf("Expected updated movie to be %v, but got %v", movie, updated)
	}
}

func TestListMovies(t *testing.T) {
	// Arrange
	movies := testhelpers.MovieFactory.Batch(2)
	sut, repo := buildMoviesSut(t)
	repo.EXPECT().
		GetAll().
		Return(movies, nil).
		Times(1)

	// Act
	list, err := sut.ListMovies()

	// Assert
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(list, movies) {
		t.Errorf("Expected list of movies to be %v, but got %v", movies, list)
	}
}

func TestDeleteMovie(t *testing.T) {
	// Arrange
	movie := testhelpers.MovieFactory.Build()
	sut, repo := buildMoviesSut(t)
	repo.EXPECT().
		Delete(gomock.Eq(*movie.ID)).
		Return(nil).
		Times(1)

	// Act
	err := sut.DeleteMovie(*movie.ID)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

func TestGetMovieByID(t *testing.T) {
	// Arrange
	movie := testhelpers.MovieFactory.Build()
	sut, repo := buildMoviesSut(t)
	repo.EXPECT().
		GetByID(gomock.Eq(*movie.ID)).
		Return(movie, nil).
		Times(1)

	// Act
	foundMovie, err := sut.GetMovieByID(*movie.ID)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if foundMovie != movie {
		t.Errorf("Expected found movie to be %v, but got %v", movie, foundMovie)
	}
}
