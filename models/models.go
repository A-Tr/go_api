package models

type PokeAPIResponse struct {
	Count   int             `json:"count"`
	Next    string          `json:"next"`
	Results []SinglePokemon `json:"results"`
}

type SinglePokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}