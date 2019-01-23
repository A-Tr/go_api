package handlers

import (
	"net/http"
	"time"
)

var netClient = &http.Client{
	Timeout: time.Second * 10,
}