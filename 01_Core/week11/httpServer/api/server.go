package api

import (
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)
// Gorilla Mux is helpful for router for method-based routes, and inherits from the mux router
type Item struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
}

type Server struct {
	*mux.Router

	shoppingItems []Item	
}

func NewServer() *Server {
	s := &Server{
		Router:		mux.NewRouter(),
		shoppingItems: []Item{},
	}
	return s
}
// Typically would not store data in our server struct but it is fine for demo it would be a connection to a database

// Create our write handlers

// tutorial https://www.youtube.com/watch?v=5BIylxkudaE
//Pulls from Shopping cart 

