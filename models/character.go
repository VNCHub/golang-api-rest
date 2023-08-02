package models

type Character struct {
	Name    string `json:"name"`
	History string `json:"history"`
}

var Characters []Character
