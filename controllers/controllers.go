package controllers

import (
	"api-go-rest/database"
	"api-go-rest/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllCharactersHandler(w http.ResponseWriter, r *http.Request) {
	var p []models.Character
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func OneCharacterHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Character
	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}

func CreateCharacter(w http.ResponseWriter, r *http.Request) {
	var char models.Character
	json.NewDecoder(r.Body).Decode(&char) //recebe dados e decodifica no char
	database.DB.Create(&char)
	json.NewEncoder(w).Encode(char) //encodifica char e exibe no ReponseWriter

}
