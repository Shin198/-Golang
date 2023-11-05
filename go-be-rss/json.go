package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responeWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responde with 5XX error: ", msg)
	}
	//convert to JSON responde
	type errResponse struct {
		Error string `json:"error"`
	}

	responeWithJSON(w, code, errResponse{
		Error: msg,
	})
}
func responeWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
