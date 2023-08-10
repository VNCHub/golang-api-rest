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
	var char []models.Character
	database.DB.Find(&char)
	json.NewEncoder(w).Encode(char)
}

func OneCharacterHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var char models.Character
	database.DB.First(&char, id)
	json.NewEncoder(w).Encode(char)
}

func CreateCharacter(w http.ResponseWriter, r *http.Request) {
	var char models.Character
	json.NewDecoder(r.Body).Decode(&char) //recebe dados e decodifica no char
	database.DB.Create(&char)
	json.NewEncoder(w).Encode(char) //encodifica char e exibe no ReponseWriter

}

func DeleteCharacter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var char models.Character
	database.DB.Delete(&char, id) //deleta uma personalidade
	json.NewEncoder(w).Encode(char)
}

func EditCharacter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var char models.Character
	database.DB.First(&char, id)
	json.NewDecoder(r.Body).Decode(&char)
	database.DB.Save(&char)
	json.NewEncoder(w).Encode(char)
}
