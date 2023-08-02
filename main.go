package main

import (
	"api-go-rest/models"
	"api-go-rest/routes"
	"fmt"
)

func main() {
	models.Characters = []models.Character{
		{Name: "Harry Potter", History: "The boy who survived..."},
		{Name: "Ron Weasley", History: "The best friend on Earth..."},
	}

	fmt.Println("Starting REST Server...")
	routes.HandleRequest()
}
