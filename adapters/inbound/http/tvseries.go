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
		response[i] = *TVSeriesFromModel(&serie)
	}

	jsonRes, _ := json.Marshal(response)
	w.Write([]byte(jsonRes))
}

func (TvSeriesHandler) TVSeriesCreate(w http.ResponseWriter, r *http.Request) {
	var requestTVSeries TVSeries
	requestPayload, _ := io.ReadAll(r.Body)
	json.Unmarshal(requestPayload, &requestTVSeries)

	createdTVSeries, _ := usecases.CreateTVSeries(*requestTVSeries.ToModel())
	jsonRes, _ := json.Marshal(TVSeriesFromModel(&createdTVSeries))
	w.Write([]byte(jsonRes))
}

func (TvSeriesHandler) TVSeriesGet(w http.ResponseWriter, r *http.Request, id string) {
	tvSeries, err := usecases.GetTVSeriesByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(responseData{
		data:       TVSeriesFromModel(&tvSeries),
		usecaseErr: nil,
		w:          w,
		statusCode: http.StatusOK,
	})
}

func (TvSeriesHandler) TVSeriesUpdate(w http.ResponseWriter, r *http.Request, id string) {
	requestTVSeries, _ := unmarshalFromReq[TVSeries](r)
	requestTVSeries.Id = &id
	modelTVSeries := requestTVSeries.ToModel()
	updatedTVSeries, err := usecases.UpdateTVSeries(*modelTVSeries)

	respondWithJSON(responseData{
		data:       TVSeriesFromModel(&updatedTVSeries),
		usecaseErr: err,
		w:          w,
		statusCode: http.StatusOK,
	})
}

func (TvSeriesHandler) TVSeriesDelete(w http.ResponseWriter, r *http.Request, id string) {
	err := usecases.DeleteTVSeries(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
