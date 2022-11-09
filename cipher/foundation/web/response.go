package web

import (
	"encoding/json"
	"net/http"
	"service/foundation/logger"
)

type Body struct {
	Status int `json:"status"`
	Data   any `json:"data"`
}

type Error struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

// Respond converts a Go value to JSON and sends it to the client.
func Respond(w http.ResponseWriter, data any, statusCode int) {

	// If there is nothing to marshal then set status code and return.
	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
		return
	}

	// Convert the response value to JSON.
	jsonData, err := json.Marshal(data)
	if err != nil {
		logger.New("")
		return
	}

	// Set the content type and headers once we know marshaling has succeeded.
	w.Header().Set("Content-Type", "application/json")

	// Write the status code to the response.
	w.WriteHeader(statusCode)

	// Send the result back to the client.
	if _, err := w.Write(jsonData); err != nil {
		logger.New("")
		return
	}
}
