package http

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Sanmoo/go-api-lambda-boilerplate/core/usecases"
)

type ElectronicGamesHandler struct {
}

func (ElectronicGamesHandler) ElectronicGamesList(w http.ResponseWriter, r *http.Request, params ElectronicGamesListParams) {
	games, _ := usecases.ListElectronicGames()
	response := make([]ElectronicGame, len(games))

	for i, game := range games {
		response[i] = *ElectronicGameFromModel(&game)
	}

	jsonRes, _ := json.Marshal(response)
	w.Write([]byte(jsonRes))
}

func (ElectronicGamesHandler) ElectronicGamesCreate(w http.ResponseWriter, r *http.Request) {
	var requestGame ElectronicGame
	requestPayload, _ := io.ReadAll(r.Body)
	json.Unmarshal(requestPayload, &requestGame)

	modelGame, err := requestGame.ToModel()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdGame, _ := usecases.CreateElectronicGame(*modelGame)
	jsonRes, _ := json.Marshal(ElectronicGameFromModel(&createdGame))
	w.Write([]byte(jsonRes))
}

func (ElectronicGamesHandler) ElectronicGamesGet(w http.ResponseWriter, r *http.Request, id string) {
	game, err := usecases.GetElectronicGameByID(id)

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

func (ElectronicGamesHandler) ElectronicGamesUpdate(w http.ResponseWriter, r *http.Request, id string) {
	requestGame, _ := unmarshalFromReq[ElectronicGame](r)
	requestGame.Id = &id
	modelGame, _ := requestGame.ToModel()
	updatedGame, err := usecases.UpdateElectronicGame(*modelGame)

	respondWithJSON(responseData{
		data:       ElectronicGameFromModel(&updatedGame),
		usecaseErr: err,
		w:          w,
		statusCode: http.StatusOK,
	})
}

func (ElectronicGamesHandler) ElectronicGamesDelete(w http.ResponseWriter, r *http.Request, id string) {
	err := usecases.DeleteElectronicGame(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
