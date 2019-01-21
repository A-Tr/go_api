package api

import (
	"io/ioutil"
	"fmt"
	"time"
	"net/http"
)

var netClient = &http.Client{
	Timeout: time.Second * 10,
}

func GetAllPokemons(W http.ResponseWriter, req *http.Request) {
	response, err := netClient.Get("https://pokeapi.co/api/v2/pokemon")
	if (err != nil) {
		fmt.Println("Error doing the request")
	}
	buf, err := ioutil.ReadAll(response.Body)
	if (err != nil) {
		fmt.Println("Error reading the body")
	}
	s := string(buf)
	fmt.Println(s)
}
