package http

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Sanmoo/go-api-lambda-boilerplate/core/model"
	"github.com/Sanmoo/go-api-lambda-boilerplate/core/usecases"
)

type TvSeriesHandler struct {
	uc GenericUseCase[model.TVSeries]
}

func NewTvSeriesHandler(uc *usecases.TVSeriesUsecases) *TvSeriesHandler {
	return &TvSeriesHandler{
		uc: uc,
	}
}

func (h *TvSeriesHandler) TVSeriesList(w http.ResponseWriter, r *http.Request, params TVSeriesListParams) {
	series, _ := h.uc.List()
	response := make([]TVSeries, len(series))

	for i, serie := range series {
		response[i] = *TVSeriesFromModel(&serie)
	}

	jsonRes, _ := json.Marshal(response)
	w.Write([]byte(jsonRes))
}

func (h *TvSeriesHandler) TVSeriesCreate(w http.ResponseWriter, r *http.Request) {
	var requestTVSeries TVSeries
	requestPayload, _ := io.ReadAll(r.Body)
	json.Unmarshal(requestPayload, &requestTVSeries)

	createdTVSeries, _ := h.uc.Create(*requestTVSeries.ToModel())
	jsonRes, _ := json.Marshal(TVSeriesFromModel(&createdTVSeries))
	w.Write([]byte(jsonRes))
}

func (h *TvSeriesHandler) TVSeriesRead(w http.ResponseWriter, r *http.Request, id string) {
	tvSeries, err := h.uc.GetByID(id)

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

func (h *TvSeriesHandler) TVSeriesPut(w http.ResponseWriter, r *http.Request, id string) {
	requestTVSeries, _ := unmarshalFromReq[TVSeries](r)
	requestTVSeries.Id = &id
	modelTVSeries := requestTVSeries.ToModel()
	updatedTVSeries, err := h.uc.Update(*modelTVSeries)

	respondWithJSON(responseData{
		data:       TVSeriesFromModel(&updatedTVSeries),
		usecaseErr: err,
		w:          w,
		statusCode: http.StatusOK,
	})
}

func (h *TvSeriesHandler) TVSeriesDelete(w http.ResponseWriter, r *http.Request, id string) {
	err := h.uc.Delete(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
