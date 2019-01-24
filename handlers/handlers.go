package handlers

import (
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"encoding/json"
	md "go_api/models"
	"time"
	"net/http"
)

var netClient = &http.Client{
	Timeout: time.Second * 10,
}

func GetAllPokemons(w http.ResponseWriter, req *http.Request) {
	traceID := uuid.Must(uuid.NewV4()).String()
	pokeRes := new(md.PokeAPIResponse)
	log.Info("Hello there", "trace-id: ", traceID)
	err := GetJson("https://pokeapi.co/api/v2/pokemon", pokeRes, traceID)
	if err != nil {
		log.Printf("Error doing the request", "trace-id: ", traceID)
	}
	data, err := json.Marshal(pokeRes)

	SendJsonResponse(data, w, traceID)
}
