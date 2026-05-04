package routes

import (
	"net/http"

	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/api-matriculas/handlers"
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
