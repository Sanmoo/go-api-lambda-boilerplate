package http

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Sanmoo/go-api-lambda-boilerplate/core/usecases"
)

type NonElectronicGamesHandler struct {
}

func (NonElectronicGamesHandler) NonElectronicGamesList(w http.ResponseWriter, r *http.Request, params NonElectronicGamesListParams) {
	games, _ := usecases.ListNonElectronicGames()
	response := make([]NonElectronicGame, len(games))

	for i, game := range games {
		response[i] = *NonElectronicGameFromModel(&game)
	}

	jsonRes, _ := json.Marshal(response)
	w.Write([]byte(jsonRes))
}

func (NonElectronicGamesHandler) NonElectronicGamesCreate(w http.ResponseWriter, r *http.Request) {
	var requestGame NonElectronicGame
	requestPayload, _ := io.ReadAll(r.Body)
	json.Unmarshal(requestPayload, &requestGame)

	modelGame, err := requestGame.ToModel()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdGame, _ := usecases.CreateNonElectronicGame(*modelGame)
	jsonRes, _ := json.Marshal(NonElectronicGameFromModel(&createdGame))
	w.Write([]byte(jsonRes))
}

func (NonElectronicGamesHandler) NonElectronicGamesGet(w http.ResponseWriter, r *http.Request, id string) {
	game, err := usecases.GetNonElectronicGameByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(responseData{
		data:       NonElectronicGameFromModel(&game),
		usecaseErr: nil,
		w:          w,
		statusCode: http.StatusOK,
	})
}

func (NonElectronicGamesHandler) NonElectronicGamesUpdate(w http.ResponseWriter, r *http.Request, id string) {
	requestGame, _ := unmarshalFromReq[NonElectronicGame](r)
	requestGame.Id = &id
	modelGame, err := requestGame.ToModel()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedGame, err := usecases.UpdateNonElectronicGame(*modelGame)

	respondWithJSON(responseData{
		data:       NonElectronicGameFromModel(&updatedGame),
		usecaseErr: err,
		w:          w,
		statusCode: http.StatusOK,
	})
}

func (NonElectronicGamesHandler) NonElectronicGamesDelete(w http.ResponseWriter, r *http.Request, id string) {
	err := usecases.DeleteNonElectronicGame(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
