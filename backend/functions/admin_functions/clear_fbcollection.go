package admin_functions

import (
	"encoding/json"
	"log"
	"net/http"

	database "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/delete"
)

const TokenAdminSecreto = "unb_angrybirds_secret_token_2026" // alterar dps talvez sepa

type ColecoesPayload struct {
	Colecoes []string `json:"colecoes"`
}

func ClearCollectionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Admin-Token")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido.", http.StatusMethodNotAllowed)
		return
	}

	adminToken := r.Header.Get("X-Admin-Token")
	if adminToken == "" || adminToken != TokenAdminSecreto {
		log.Printf("Tentativa de reset não autorizada vinda do IP: %s", r.RemoteAddr)
		http.Error(w, "Não autorizado. Token administrativo inválido.", http.StatusUnauthorized)
		return
	}

	ctx := r.Context()

	var payload ColecoesPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Payload inválido.", http.StatusBadRequest)
		return
	}

	if len(payload.Colecoes) == 0 {
		http.Error(w, "Nenhuma coleção foi especificada para limpeza.", http.StatusBadRequest)
		return
	}

	for _, colecao := range payload.Colecoes {
		err := delete.ClearCollection(ctx, database.Client, colecao)
		if err != nil {
			log.Printf("Erro ao limpar a coleção %s: %v", colecao, err)
			http.Error(w, "Erro interno ao limpar a coleção: "+colecao, http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("coleções limpas com sucesso"))
}

