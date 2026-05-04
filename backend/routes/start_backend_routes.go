package routes

import (
	"net/http"

	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/handlers"
)

func StartRoutes() {
	http.HandleFunc("/login", handlers.TokenHandler)
}
