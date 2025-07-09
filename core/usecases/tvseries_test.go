package usecases

import (
	"reflect"
	"testing"

	"github.com/Sanmoo/go-api-lambda-boilerplate/core/mocks"
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/model"
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/usecases/testhelpers"
	"go.uber.org/mock/gomock"
)

func buildTVSeriesSut(t *testing.T) (*TVSeriesUsecases, *mocks.MockRepository[model.TVSeries]) {
	ctrl := gomock.NewController(t)
	fakeRepo := mocks.NewMockRepository[model.TVSeries](ctrl)
	usecases := NewTVSeriesUsecases(fakeRepo)
	return usecases, fakeRepo
}

func TestCreatesATVSeriesUsecases(t *testing.T) {
	usecases, _ := buildTVSeriesSut(t)

	if usecases == nil {
		t.Error("Expected TVSeriesUsecases to be created, but got nil")
	}
}

func TestCreatesATVSeries(t *testing.T) {
	// Arrange
	tvSeries := testhelpers.TVSeriesFactory.Build()
	sut, repo := buildTVSeriesSut(t)
	repo.EXPECT().
		Create(gomock.Eq(tvSeries)).
		Return(tvSeries, nil).
		Times(1)

	// Act
	created, err := sut.Create(tvSeries)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if created != tvSeries {
		t.Errorf("Expected created tvSeries to be %v, but got %v", tvSeries, created)
	}
}

func TestUpdatesATVSeries(t *testing.T) {
	// Arrange
	tvSeries := testhelpers.TVSeriesFactory.Build()
	sut, repo := buildTVSeriesSut(t)
	repo.EXPECT().
		Update(gomock.Eq(tvSeries)).
		Return(tvSeries, nil).
		Times(1)

	// Act
	updated, err := sut.Update(tvSeries)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if updated != tvSeries {
		t.Errorf("Expected updated tvSeries to be %v, but got %v", tvSeries, updated)
	}
}

func TestListTVSeries(t *testing.T) {
	// Arrange
	tvSeries := testhelpers.TVSeriesFactory.Batch(2)
	sut, repo := buildTVSeriesSut(t)
	repo.EXPECT().
		GetAll().
		Return(tvSeries, nil).
		Times(1)

	// Act
	list, err := sut.List()

	// Assert
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(list, tvSeries) {
		t.Errorf("Expected list of tvSeries to be %v, but got %v", tvSeries, list)
	}
}

func TestDeleteTVSeries(t *testing.T) {
	// Arrange
	tvSeries := testhelpers.TVSeriesFactory.Build()
	sut, repo := buildTVSeriesSut(t)
	repo.EXPECT().
		Delete(gomock.Eq(*tvSeries.ID)).
		Return(nil).
		Times(1)

	// Act
	err := sut.Delete(*tvSeries.ID)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

func TestGetTVSeriesByID(t *testing.T) {
	// Arrange
	tvSeries := testhelpers.TVSeriesFactory.Build()
	sut, repo := buildTVSeriesSut(t)
	repo.EXPECT().
		GetByID(gomock.Eq(*tvSeries.ID)).
		Return(tvSeries, nil).
		Times(1)

	// Act
	foundTVSeries, err := sut.GetByID(*tvSeries.ID)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if foundTVSeries != tvSeries {
		t.Errorf("Expected found tvSerie to be %v, but got %v", tvSeries, foundTVSeries)
	}
}
