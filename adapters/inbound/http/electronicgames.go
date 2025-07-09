package http

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Sanmoo/go-api-lambda-boilerplate/core/model"
)

type ElectronicGamesHandler struct {
	usecases GenericUseCase[model.ElectronicGame]
}

func NewElectronicGamesHandler(usecases GenericUseCase[model.ElectronicGame]) *ElectronicGamesHandler {
	return &ElectronicGamesHandler{
		usecases: usecases,
	}
}

func (b *ElectronicGamesHandler) ElectronicGamesList(w http.ResponseWriter, r *http.Request, params ElectronicGamesListParams) {
	games, _ := b.usecases.List()
	response := make([]ElectronicGame, len(games))

	for i, game := range games {
		response[i] = *ElectronicGameFromModel(&game)
	}

	jsonRes, _ := json.Marshal(response)
	w.Write([]byte(jsonRes))
}

func (b *ElectronicGamesHandler) ElectronicGamesCreate(w http.ResponseWriter, r *http.Request) {
	var requestGame ElectronicGame
	requestPayload, _ := io.ReadAll(r.Body)
	json.Unmarshal(requestPayload, &requestGame)

	modelGame, err := requestGame.ToModel()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdGame, _ := b.usecases.Create(*modelGame)
	jsonRes, _ := json.Marshal(ElectronicGameFromModel(&createdGame))
	w.Write([]byte(jsonRes))
}

func (b *ElectronicGamesHandler) ElectronicGamesRead(w http.ResponseWriter, r *http.Request, id string) {
	game, err := b.usecases.GetByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(responseData{
		data:       ElectronicGameFromModel(&game),
		usecaseErr: nil,
		w:          w,
		statusCode: http.StatusOK,
	})
}

func (b *ElectronicGamesHandler) ElectronicGamesPut(w http.ResponseWriter, r *http.Request, id string) {
	requestGame, _ := unmarshalFromReq[ElectronicGame](r)
	requestGame.Id = &id
	modelGame, _ := requestGame.ToModel()
	updatedGame, err := b.usecases.Update(*modelGame)

	respondWithJSON(responseData{
		data:       ElectronicGameFromModel(&updatedGame),
		usecaseErr: err,
		w:          w,
		statusCode: http.StatusOK,
	})
}

func (b *ElectronicGamesHandler) ElectronicGamesDelete(w http.ResponseWriter, r *http.Request, id string) {
	err := b.usecases.Delete(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
