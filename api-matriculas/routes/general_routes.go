package routes

import (
	"api-matriculas/handlers"
	"net/http"
)

func GeneralRoutes() {
	http.HandleFunc("/", handlers.TestHandler)
	http.HandleFunc("/ping", handlers.PingHandler)
	http.HandleFunc("/time", handlers.TimeHandler)
	http.HandleFunc("/status", handlers.StatusHandler)

	// curl -H "institutional_key: KEY" localhost:8080/Verify/Key
	http.HandleFunc("/Verify/Key", handlers.VerifyKeyHandler)

	//http.HandleFunc("/bdcapacity", handlers.TestHandler) To do
}
