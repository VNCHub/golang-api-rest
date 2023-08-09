package main

import (
	"api-go-rest/database"
	"api-go-rest/models"
	"api-go-rest/routes"
	"fmt"
)

func main() {
	models.Characters = []models.Character{
		{Id: 1, Name: "Harry Potter", History: "The boy who survived..."},
		{Id: 2, Name: "Ron Weasley", History: "The best friend on Earth..."},
	}
	fmt.Println("Starting Database Connection...")
	database.DatabaseConnect()
	fmt.Println("Starting HandleRequest...")
	routes.HandleRequest()
}
