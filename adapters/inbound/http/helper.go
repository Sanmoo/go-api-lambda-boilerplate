// create a generic function to extract the request body and unmarshal it into a struct
package http

import (
	"encoding/json"
	"io"
	"net/http"
)

func unmarshalFromReq[T any](r *http.Request) (*T, error) {
	var requestData T
	requestPayload, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(requestPayload, &requestData)
	if err != nil {
		return nil, err
	}

	return &requestData, nil
}

type responseData struct {
	data       interface{}
	usecaseErr error
	w          http.ResponseWriter
	statusCode int
}

func respondWithJSON(responseData responseData) error {
	if responseData.usecaseErr != nil {
		http.Error(responseData.w, responseData.usecaseErr.Error(), http.StatusBadRequest)
		return nil
	}

	jsonRes, err := json.Marshal(responseData.data)

	if err != nil {
		// This should not happen
		panic(err)
	}

	responseData.w.Header().Set("Content-Type", "application/json")
	responseData.w.WriteHeader(responseData.statusCode)
	_, err = responseData.w.Write(jsonRes)
	return err
}
