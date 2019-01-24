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
	log.WithField("trace-id", traceID).Info("HOli")
	err := GetJson("https://pokeapi.co/api/v2/pokemon", pokeRes, traceID)
	if err != nil {
		log.WithField("trace-id", traceID).Error("Error doing the request")
	}
	data, err := json.Marshal(pokeRes)

	SendJsonResponse(data, w, traceID)
}
