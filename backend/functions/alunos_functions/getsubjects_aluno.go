package alunos_functions

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"fmt"

	database "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

type InfoMatriculaHome struct {
	NomeMateria string    `json:"nome"`
	Local       string    `json:"local"`
	Horario     string    `json:"horario"`
	TurmaId     string    `json:"turma"`
	CodigoData  time.Time `json:"dataSolicitacao"`
}

func GetSubjectsAlunoHandler(w http.ResponseWriter, r *http.Request) {
	// permite que qualquer origem chegue aqui (ALTERAR DEPOIS; CORS)
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// chatice do go
	if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido.", http.StatusMethodNotAllowed)
		return
	}

	alunoId := r.URL.Query().Get("id")
	if alunoId == "" {
		http.Error(w, "O parâmetro 'id' do aluno é obrigatório.", http.StatusBadRequest)
		return
	}

	ctx := r.Context();

	var listaMatriculas []models.Matricula

	iterator := database.Client.Collection("matriculas").Where("alunoId", "==", alunoId).Documents(ctx)
	docs, err := iterator.GetAll()

	if err != nil {
		log.Printf("erro ao buscar matrículas do aluno %s: %v", alunoId, err)
		http.Error(w, "erro ao buscar matrículas do aluno.", http.StatusBadRequest)
		return
	}
	
	for _, doc := range docs {
		var matricula models.Matricula
		if err := doc.DataTo(&matricula); err != nil {
			continue 
		}

		listaMatriculas = append(listaMatriculas, matricula)
	}

	var gradeCompleta []InfoMatriculaHome

	for _, m := range listaMatriculas {
		turmaDoc, err := database.Client.Collection("turmas").Doc(m.TurmaId).Get(ctx)
		if err != nil {continue}
		
		materiaDoc, err := database.Client.Collection("materias").Doc(m.MateriaId).Get(ctx)
		if err != nil {continue}

		var nomeMateria, local, horario string
		
		if nome, err := materiaDoc.DataAt("nomeMateria"); err == nil {
			nomeMateria = fmt.Sprintf("%v", nome)	
		}
		if loc, err := turmaDoc.DataAt("local"); err == nil {
			local = fmt.Sprintf("%v", loc)
		}
		if hor, err := turmaDoc.DataAt("horario"); err == nil {
			horario = fmt.Sprintf("%v", hor)
		}

		gradeCompleta = append(gradeCompleta, InfoMatriculaHome{
			NomeMateria: nomeMateria,
			Local:       local,
			Horario:     horario,
			TurmaId:     m.TurmaId,
			CodigoData:  m.DataSolicitacao,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(gradeCompleta)
}
