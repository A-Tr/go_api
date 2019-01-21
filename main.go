package main

import (
	"go_api/api"
	"github.com/gorilla/mux"
	"net/http"

)
type Pokemon struct {
	ID   int
	Name string
}

type PokemonService interface {
	Pokemon(id int) (*Pokemon, error)
	Pokemons() ([]*Pokemon, error)
}


func main() {
	server := mux.NewRouter()
	server.HandleFunc("/", api.GetAllPokemons)
	http.ListenAndServe(":8080", server)
}
