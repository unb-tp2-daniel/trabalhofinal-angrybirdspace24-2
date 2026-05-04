package p

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"google.golang.org/api/iterator"

	database "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/BD"
)

type Turma struct {
	CodigoTurma    string `json:"codigoTurma" firestore:"codigoTurma"`
	MateriaId      string `json:"materiaId" firestore:"materiaId"`
	NomeMateria    string `json:"nomeMateria" firestore:"nomeMateria"`
	Semestre       string `json:"semestre" firestore:"semestre"`
	Capacidade     int    `json:"capacidade" firestore:"capacidade"`
	Ocupadas       int    `json:"ocupadas" firestore:"ocupadas"`
	VagasRestantes int    `json:"vagasRestantes" firestore:"vagasRestantes"`
	Status         string `json:"status" firestore:"status"`
}

func init() {
	// Inicia a conexão com o banco assim que a Cloud Function "acordar"
	database.InitDB()
	functions.HTTP("ListarTurmasDB", ListarTurmasDBHandler)
}

func ListarTurmasDBHandler(w http.ResponseWriter, r *http.Request) {
	var turmas []Turma

	// Busca todos os documentos dentro da coleção "turmas"
	iter := database.Client.Collection("turmas").Documents(database.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Erro ao iterar turmas do banco: %v", err)
			http.Error(w, "Erro interno ao ler turmas", http.StatusInternalServerError)
			return
		}

		var t Turma
		// Joga os dados do Firestore pra Struct
		if err := doc.DataTo(&t); err != nil {
			log.Printf("Erro ao converter doc para struct: %v", err)
			continue
		}
		turmas = append(turmas, t)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(turmas); err != nil {
		log.Printf("Erro ao retornar JSON: %v", err)
		http.Error(w, "Erro na resposta", http.StatusInternalServerError)
	}
}
