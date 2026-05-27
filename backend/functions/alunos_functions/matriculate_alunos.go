// functions/alunos_functions/matriculate_aluno.go
package alunos_functions

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	// Importando as nossas pastas isoladas
	"cloud.google.com/go/firestore"
	database "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/create"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/read"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

// MatriculateAlunoHandler lida exclusivamente com a requisição da internet
func MatriculateAlunoHandler(w http.ResponseWriter, r *http.Request) {
	// permite que qualquer origem chegue aqui (ALTERAR DEPOIS; CORS)
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// chatice do go
	if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido.", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context();

	var matricula models.Matricula
	if err := json.NewDecoder(r.Body).Decode(&matricula); err != nil {
		log.Printf("Erro ao decodificar JSON da matrícula: %v", err)
		http.Error(w, "Formato de dados inválido", http.StatusBadRequest)
		return
	}

	matriculaId := matricula.AlunoId + "_" + matricula.TurmaId
	matriculaRef := database.Client.Collection("matriculas").Doc(matriculaId)
	turmaRef := database.Client.Collection("turmas").Doc(matricula.TurmaId)

	// transação do firebase. Basicamente, impede que outro registro ocorra enquanto esse ta ocorrendo
	// tratando as race conditions

	err := database.Client.RunTransaction(ctx, func(ctx context.Context, t *firestore.Transaction) error {
		// verifica se ja esta matriculado
		matDoc, err := t.Get(matriculaRef)
		if err == nil && matDoc.Exists() {
			return fmt.Errorf("aluno_ja_matriculado")
		}

		turmaDoc, err := t.Get(turmaRef)
		if err != nil {
			return err
		}

		var turma models.Turma
		err = turmaDoc.DataTo(&turma)
		if err != nil {
			return err
		}

		// verifica se ainda há vagas
		if turma.VagasOcupadas >= turma.VagasTotais {
			return fmt.Errorf("vagas_esgotadas")
		}

		// senão, realiza a matricula
		t.Update(turmaRef, []firestore.Update{
			{Path: "vagasOcupadas", Value: turma.VagasOcupadas + 1},
		})

		matricula.Status = true
		matricula.DataSolicitacao = time.Now()

		t.Set(matriculaRef, matricula)

		return nil
	})

	if err != nil {
		if err.Error() == "aluno_ja_matriculado" {
			http.Error(w, "Você já possui uma matrícula ou solicitação ativa nesta turma.", http.StatusConflict)
			return
		}
		if err.Error() == "vagas_esgotadas" {
			http.Error(w, "Vagas esgotadas.", http.StatusConflict)
			return
		}

		log.Printf("Erro na transação de matrícula: %v", err)
		http.Error(w, "Erro interno ao processar a matrícula", http.StatusInternalServerError)
		return
	}

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("Matrícula extraordinária realizada com sucesso!"))
}

// lida com a matricula normal e rematricula do aluno
func NormalMatriculateAlunoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido.", http.StatusMethodNotAllowed)
		return
	}
	var matricula models.Matricula
	if err := json.NewDecoder(r.Body).Decode(&matricula); err != nil {
		log.Printf("Erro ao decodificar JSON da matrícula: %v", err)
		http.Error(w, "Formato de dados inválido", http.StatusBadRequest)
		return
	}

	// não há necessidade de verificar se há disponibilidade na turmaref nesse caso

	ctx := r.Context()

	// aplicação de regras de prioridade
	aluno, err := read.GetAlunoById(ctx, database.Client, matricula.AlunoId)
	if err != nil {
		log.Printf("Error ao resgatar aluno no banco de dados %v", err)
		http.Error(w, "Error ao resgatar aluno no banco de dados", http.StatusInternalServerError)
		return
	}

	turmaref, err := read.GetTurmaById(ctx, database.Client, matricula.TurmaId)
	if err != nil {
		log.Printf("Error ao resgatar turmaref no banco de dados %v", err)
		http.Error(w, "Error ao resgatar turmaref no banco de dados", http.StatusInternalServerError)
		return
	}

	materia, err := read.GetMateriaById(ctx, database.Client, turmaref.MateriaId)
	if err != nil {
		log.Printf("Error ao resgatar materia no banco de dados %v", err)
		http.Error(w, "Error ao resgatar materia no banco de dados", http.StatusInternalServerError)
		return
	}

	curso, err := read.GetCursoById(ctx, database.Client, aluno.CursoId)
	if err != nil {
		log.Printf("Error ao resgatar curso no banco de dados %v", err)
		http.Error(w, "Error ao resgatar curso no banco de dados", http.StatusInternalServerError)
		return
	}

	matricula.PrioridadeNota = ApplyRules(*aluno, *materia, *curso)

	err = create.CreateMatricula(ctx, database.Client, matricula)
	if err != nil {
		log.Printf("Erro ao salvar matrícula no banco: %v", err)
		http.Error(w, "Erro interno ao salvar a matrícula", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Matrícula cadastrada com sucesso no banco de dados!"))	
}