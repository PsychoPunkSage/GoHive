package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		// Which means Bug is on our side
		log.Println("Responding with 5XX level error:", msg)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	RespondWithJson(w, code, errorResponse{
		Error: msg,
	})
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload) // GoResponse -> JSON
	if err != nil {
		w.WriteHeader(500)
		log.Println("Failed to marshal JSON payload: ", payload)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
