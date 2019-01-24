package main

import (
	log"github.com/sirupsen/logrus"
	"go_api/api"
	"net/http"

)


func main() {
	router := api.NewRouter()
	log.SetFormatter(&log.JSONFormatter{})
	
	log.Fatal(http.ListenAndServe(":3000", router))
}