package routes

import (
	"api-go-rest/controllers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.HomeHandler)
	r.HandleFunc("/api/characters", controllers.AllCharactersHandler).Methods("Get")
	r.HandleFunc("/api/character/{id}", controllers.OneCharacterHandler).Methods("Get")
	r.HandleFunc("/api/characters", controllers.CreateCharacter).Methods("Post")
	r.HandleFunc("/api/character/{id}", controllers.DeleteCharacter).Methods("Delete")
	r.HandleFunc("/api/character/{id}", controllers.EditCharacter).Methods("Put")
	http.Handle("/", r)

	fmt.Println("Starting API REST on port :8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
