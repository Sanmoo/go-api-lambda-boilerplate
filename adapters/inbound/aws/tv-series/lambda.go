package main

import (
	"github.com/Sanmoo/go-api-lambda-boilerplate/adapters/inbound/aws"
)

func main() {
	aws.ListenAndServe(aws.TVSeriesHandler{}, "tv-series")
}
