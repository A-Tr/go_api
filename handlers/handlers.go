package handlers

import (
	"encoding/json"
	md "go_api/models"
	"log"
	"net/http"
)

func WrappedGetAllPokemons() {

}

func GetAllPokemons(w http.ResponseWriter, req *http.Request) {
	response, err := netClient.Get("https://pokeapi.co/api/v2/pokemon")
	var pokeRes md.PokeAPIResponse

	if err != nil {
		log.Printf("Error doing the request")
	}
	err = json.NewDecoder(response.Body).Decode(&pokeRes)
	if err != nil {
		log.Printf("Error reading the body")
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pokeRes)
}
