package routes

import (
	"net/http"

	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/api-matriculas/handlers"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/api-matriculas/middleware"
)

func AuthRoutes() {
	http.HandleFunc("/auth", middleware.RequireInstitutionKey(handlers.GenerateTokenHandler))
}
