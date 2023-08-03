package main

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/google/uuid"
)

func(apiCfg *apiConfig) handlerReadiness(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params:= parameters{}
	err:=decoder.Decode(&params)
	if err!=nil{
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user,err:=apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
	ID:uuid.New(),
	CreatedAt:time.Now().UTC(),
	UpdatedAt:time.Now().UTC(),
	Name:params.Name,
	})

	if err!=nil{
		respondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {

	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))

}