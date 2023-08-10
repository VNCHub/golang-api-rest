package routes

import (
	"api-go-rest/controllers"
	"api-go-rest/middleware"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.HomeHandler)
	r.HandleFunc("/api/characters", controllers.AllCharactersHandler).Methods("Get")
	r.HandleFunc("/api/characters/{id}", controllers.OneCharacterHandler).Methods("Get")
	r.HandleFunc("/api/characters", controllers.CreateCharacter).Methods("Post")
	r.HandleFunc("/api/characters/{id}", controllers.DeleteCharacter).Methods("Delete")
	r.HandleFunc("/api/characters/{id}", controllers.EditCharacter).Methods("Put")
	http.Handle("/", r)

	fmt.Println("Starting API REST on port :8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
