package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	log "github.com/sirupsen/logrus"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func GetJson(url string, target interface{}, traceID string) error {
		log.WithField("trace-id", traceID).Info("GET request to URL:")
		r, err := myClient.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}

func SendJsonResponse (data []byte, w http.ResponseWriter, traceID string) {
	log.WithField("trace-id", traceID).Info("Sending JSON response")
	w.WriteHeader(200)
	w.Write(data)
	return
}