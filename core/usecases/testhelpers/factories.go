package testhelpers

import (
	"github.com/Goldziher/fabricator"
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/model"
)

var BookFactory = fabricator.New(model.Book{})
var MovieFactory = fabricator.New(model.Movie{})
var ElectronicGameFactory = fabricator.New(model.ElectronicGame{})
var NonElectronicGameFactory = fabricator.New(model.NonElectronicGame{})
var TVSeriesFactory = fabricator.New(model.TVSeries{})
