package routes

import (
	"fmt"
	"log"
	"net/http"

	database "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/BD"
)

func DatabaseRoutes() {

	http.HandleFunc("/teste-banco", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		_, err := database.Client.Collection("testes").Doc("ping").Set(ctx, map[string]interface{}{
			"mensagem": "A API conectou",
			"status":   "Sucesso",
		})

		if err != nil {

			log.Printf("ERRO FIRESTORE: %v", err)

			// Mostra o erro detalhado na tela do navegador
			mensagemErro := fmt.Sprintf("Erro ao gravar: %v", err)
			http.Error(w, mensagemErro, http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Olhe o console do Firebase"))
	})
}
