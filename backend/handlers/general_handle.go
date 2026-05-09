package handlers

import (
	"net/http"
	"time"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC1123)
	w.Write([]byte("Current time: " + currentTime))
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Status: Online"))
}
