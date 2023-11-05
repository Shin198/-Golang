package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Shin198/go-be-rss/internal/database"
)

// call pointer here to handler
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameter struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		responeWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.db.CreateUser(r.Context(), database.CreateUserParams{
		CreateAt: time.Now().UTC(),
		UpdateAt: time.Now().UTC(),
		Name:     params.Name,
	})
	if err != nil {
		responeWithError(w, 400, fmt.Sprintf("Can't create user: %v", err))
		return
	}
	responeWithJSON(w, 200, databaseUserToUser(User))
}
