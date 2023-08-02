package routes

import (
	"api-go-rest/controllers"
	"log"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/characters", controllers.AllCharacters)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
