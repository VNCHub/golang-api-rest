package models

type Character struct {
	Id      int    `"id"`
	Name    string `json:"name"`
	History string `json:"history"`
}

var Characters []Character
