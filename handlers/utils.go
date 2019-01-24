package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	log "github.com/sirupsen/logrus"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func GetJson(url string, target interface{}, traceId string) error {
		log.Info("GET request to URL:", url, "trace-id: ", traceId)
		r, err := myClient.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}

func SendJsonResponse (data []byte, w http.ResponseWriter, traceId string) {
	log.Info("Sending JSON response. trace-id: ", traceId)
	w.WriteHeader(200)
	w.Write(data)
	return
}