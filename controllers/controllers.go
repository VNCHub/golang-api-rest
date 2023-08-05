package controllers

import (
	"api-go-rest/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllCharactersHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Characters)
}

func OneCharacterHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, char := range models.Characters {
		id, _ := strconv.Atoi(id)
		if char.Id == id {
			json.NewEncoder(w).Encode(char)
		}
	}
}
