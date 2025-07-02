package http

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Sanmoo/go-api-lambda-boilerplate/core/usecases"
)

type TvSeriesHandler struct {
}

func (TvSeriesHandler) TVSeriesList(w http.ResponseWriter, r *http.Request, params TVSeriesListParams) {
	series, _ := usecases.ListTVSeries()
	response := make([]TVSeries, len(series))

	for i, serie := range series {
		response[i] = *TVSeriesFromModel(serie)
	}

	jsonRes, _ := json.Marshal(response)
	w.Write([]byte(jsonRes))
}

func (TvSeriesHandler) TVSeriesCreate(w http.ResponseWriter, r *http.Request) {
	var requestTVSeries TVSeries
	requestPayload, _ := io.ReadAll(r.Body)
	json.Unmarshal(requestPayload, &requestTVSeries)

	createdTVSeries, _ := usecases.CreateTVSeries(requestTVSeries.ToModel())
	jsonRes, _ := json.Marshal(TVSeriesFromModel(createdTVSeries))
	w.Write([]byte(jsonRes))
}
