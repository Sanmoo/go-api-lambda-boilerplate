package model

// Defines values for BookStatus.
const (
	BookStatusGaveUp  BookStatus = "gave_up"
	BookStatusPaused  BookStatus = "paused"
	BookStatusRead    BookStatus = "read"
	BookStatusReading BookStatus = "reading"
	BookStatusToRead  BookStatus = "to_read"
)

// Defines values for ElectronicGameStatus.
const (
	ElectronicGameStatusBeated     ElectronicGameStatus = "beated"
	ElectronicGameStatusGaveUp     ElectronicGameStatus = "gave_up"
	ElectronicGameStatusPlatinated ElectronicGameStatus = "platinated"
	ElectronicGameStatusToBeat     ElectronicGameStatus = "to_beat"
	ElectronicGameStatusToTry      ElectronicGameStatus = "to_try"
	ElectronicGameStatusTried      ElectronicGameStatus = "tried"
)

// Defines values for MovieStatus.
const (
	MovieStatusGaveUpIncomplete MovieStatus = "gave_up_incomplete"
	MovieStatusToWatch          MovieStatus = "to_watch"
	MovieStatusWatched          MovieStatus = "watched"
)

// Defines values for NonElectronicGameStatus.
const (
	ToTry NonElectronicGameStatus = "to_try"
	Tried NonElectronicGameStatus = "tried"
)

// Defines values for NonElectronicGameType.
const (
	Board NonElectronicGameType = "board"
	Card  NonElectronicGameType = "card"
	Dice  NonElectronicGameType = "dice"
	Other NonElectronicGameType = "other"
)

// Defines values for TVSeriesStatus.
const (
	TVSeriesStatusGaveUpIncomplete TVSeriesStatus = "gave_up_incomplete"
	TVSeriesStatusToWatch          TVSeriesStatus = "to_watch"
	TVSeriesStatusWatched          TVSeriesStatus = "watched"
	TVSeriesStatusWatching         TVSeriesStatus = "watching"
)

// Book defines model for Book.
type Book struct {
	Media
	Author string
	Genre  string
	Status *BookStatus
}

// BookStatus defines model for BookStatus.
type BookStatus string

// ElectronicGame defines model for ElectronicGame.
type ElectronicGame struct {
	Media
	Platform string
	Genre    string
	Status   *ElectronicGameStatus
}

// ElectronicGameStatus defines model for ElectronicGameStatus.
type ElectronicGameStatus string

// Media defines model for Media.
type Media struct {
	ID     *string
	Rating *int32
	Title  string
}

// Movie defines model for Movie.
type Movie struct {
	Media
	Director string
	Genre    string
	Status   *MovieStatus
}

// MovieStatus defines model for MovieStatus.
type MovieStatus string

// NonElectronicGame defines model for NonElectronicGame.
type NonElectronicGame struct {
	Media
	Type   NonElectronicGameType
	Status NonElectronicGameStatus
}

// NonElectronicGameStatus defines model for NonElectronicGameStatus.
type NonElectronicGameStatus string

// NonElectronicGameType defines model for NonElectronicGameType.
type NonElectronicGameType string

// TVSeries defines model for TVSeries.
type TVSeries struct {
	Media
	Seasons  *int32
	Finished *bool
	Status   *TVSeriesStatus
}

// TVSeriesStatus defines model for TVSeriesStatus.
type TVSeriesStatus string
