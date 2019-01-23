package main

import (
	"go_api/api"
	"net/http"
	"log"

)


func main() {
	router := api.NewRouter()
	log.Printf("arrancado")
	log.Fatal(http.ListenAndServe(":3000", router))
}
