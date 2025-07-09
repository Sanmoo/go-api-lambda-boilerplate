package usecases

import "testing"

func CreatesABooksUsecases(t *testing.T) {
	usecases := NewBooksUsecases()

	if usecases == nil {
		t.Error("Expected BooksUsecases to be created, but got nil")
	}

	if usecases.repository != repo {
		t.Errorf("Expected repository to be %v, but got %v", repo, usecases.repository)
	}
}
