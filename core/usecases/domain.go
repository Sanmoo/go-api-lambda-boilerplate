package usecases

import (
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/model"
)

type BooksUsecases struct {
	*GenericUsecases[model.Book]
}

func NewBooksUsecases(repo Repository[model.Book]) *BooksUsecases {
	return &BooksUsecases{GenericUsecases: NewGenericUsecases(repo)}
}

func (b *BooksUsecases) ListBooks() ([]model.Book, error) {
	return b.List()
}

func (b *BooksUsecases) CreateBook(book model.Book) (model.Book, error) {
	return b.Create(book)
}

func (b *BooksUsecases) UpdateBook(book model.Book) (model.Book, error) {
	return b.Update(book)
}

func (b *BooksUsecases) DeleteBook(id string) error {
	return b.Delete(id)
}

func (b *BooksUsecases) GetBookByID(id string) (model.Book, error) {
	return b.GetByID(id)
}

type MoviesUsecases struct {
	*GenericUsecases[model.Movie]
}

func NewMoviesUsecases(repo Repository[model.Movie]) *MoviesUsecases {
	return &MoviesUsecases{GenericUsecases: NewGenericUsecases(repo)}
}

func (m *MoviesUsecases) ListMovies() ([]model.Movie, error) {
	return m.List()
}

func (m *MoviesUsecases) CreateMovie(movie model.Movie) (model.Movie, error) {
	return m.Create(movie)
}

func (m *MoviesUsecases) UpdateMovie(movie model.Movie) (model.Movie, error) {
	return m.Update(movie)
}

func (m *MoviesUsecases) DeleteMovie(id string) error {
	return m.Delete(id)
}

func (m *MoviesUsecases) GetMovieByID(id string) (model.Movie, error) {
	return m.GetByID(id)
}

type TVSeriesUsecases struct {
	*GenericUsecases[model.TVSeries]
}

func NewTVSeriesUsecases(repo Repository[model.TVSeries]) *TVSeriesUsecases {
	return &TVSeriesUsecases{GenericUsecases: NewGenericUsecases(repo)}
}

func (t *TVSeriesUsecases) ListTVSeries() ([]model.TVSeries, error) {
	return t.List()
}

func (t *TVSeriesUsecases) CreateTVSeries(tv model.TVSeries) (model.TVSeries, error) {
	return t.Create(tv)
}

func (t *TVSeriesUsecases) UpdateTVSeries(tv model.TVSeries) (model.TVSeries, error) {
	return t.Update(tv)
}

func (t *TVSeriesUsecases) DeleteTVSeries(id string) error {
	return t.Delete(id)
}

func (t *TVSeriesUsecases) GetTVSeriesByID(id string) (model.TVSeries, error) {
	return t.GetByID(id)
}

type ElectronicGamesUsecases struct {
	*GenericUsecases[model.ElectronicGame]
}

func NewElectronicGamesUsecases(repo Repository[model.ElectronicGame]) *ElectronicGamesUsecases {
	return &ElectronicGamesUsecases{GenericUsecases: NewGenericUsecases(repo)}
}

func (e *ElectronicGamesUsecases) ListElectronicGames() ([]model.ElectronicGame, error) {
	return e.List()
}

func (e *ElectronicGamesUsecases) CreateElectronicGame(game model.ElectronicGame) (model.ElectronicGame, error) {
	return e.Create(game)
}

func (e *ElectronicGamesUsecases) UpdateElectronicGame(game model.ElectronicGame) (model.ElectronicGame, error) {
	return e.Update(game)
}

func (e *ElectronicGamesUsecases) DeleteElectronicGame(id string) error {
	return e.Delete(id)
}

func (e *ElectronicGamesUsecases) GetElectronicGameByID(id string) (model.ElectronicGame, error) {
	return e.GetByID(id)
}

type NonElectronicGamesUsecases struct {
	*GenericUsecases[model.NonElectronicGame]
}

func NewNonElectronicGamesUsecases(repo Repository[model.NonElectronicGame]) *NonElectronicGamesUsecases {
	return &NonElectronicGamesUsecases{GenericUsecases: NewGenericUsecases(repo)}
}

func (n *NonElectronicGamesUsecases) ListNonElectronicGames() ([]model.NonElectronicGame, error) {
	return n.List()
}

func (n *NonElectronicGamesUsecases) CreateNonElectronicGame(game model.NonElectronicGame) (model.NonElectronicGame, error) {
	return n.Create(game)
}

func (n *NonElectronicGamesUsecases) UpdateNonElectronicGame(game model.NonElectronicGame) (model.NonElectronicGame, error) {
	return n.Update(game)
}

func (n *NonElectronicGamesUsecases) DeleteNonElectronicGame(id string) error {
	return n.Delete(id)
}

func (n *NonElectronicGamesUsecases) GetNonElectronicGameByID(id string) (model.NonElectronicGame, error) {
	return n.GetByID(id)
}
