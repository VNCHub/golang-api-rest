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
	fmt.Println("Starting HandleFunc 1 (HomeHandler)...")
	r.HandleFunc("/", controllers.HomeHandler)

	fmt.Println("Starting HandleFunc 2 (AllCharactersHandler)...")
	r.HandleFunc("/api/characters", controllers.AllCharactersHandler)

	fmt.Println("Starting HandleFunc 3 (OneCharacterHandler)...")
	r.HandleFunc("/api/character/{id}", controllers.OneCharacterHandler)

	fmt.Println("Starting REST...")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
