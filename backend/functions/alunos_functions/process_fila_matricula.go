// functions/alunos_functions/process_fila_matricula.go
package alunos_functions

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"time"

	"cloud.google.com/go/firestore"
	database "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/read"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

// ResultadoProcessamento resume o que aconteceu ao processar a fila de uma turma.
type ResultadoProcessamento struct {
	TurmaId       string   `json:"turmaId"`
	VagasTotais   int      `json:"vagasTotais"`
	VagasOcupadas int      `json:"vagasOcupadas"`
	Efetivados    []string `json:"efetivados"`    // alunoIds que entraram
	ListaEspera   []string `json:"listaEspera"`   // alunoIds que ficaram fora
	ProcessadoEm  string   `json:"processadoEm"`
}

// ProcessarFilaMatriculaHandler recebe um turmaId e processa toda a fila pendente.
// Deve ser chamado ao fim do período de solicitações (ex: cron job ou manualmente pelo coordenador).
//
// Requisição esperada (POST /processar-fila):
//
//	{ "turmaId": "TURMA_001" }
func ProcessarFilaMatriculaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido.", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()

	// --- 1. Decodifica o body ---
	var req struct {
		TurmaId string `json:"turmaId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.TurmaId == "" {
		http.Error(w, "turmaId é obrigatório.", http.StatusBadRequest)
		return
	}

	resultado, err := processarFila(ctx, req.TurmaId)
	if err != nil {
		log.Printf("Erro ao processar fila da turma %s: %v", req.TurmaId, err)
		http.Error(w, "Erro interno ao processar a fila.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)
}

// processarFila contém toda a lógica de negócio — separada do handler para ser
// reutilizável por um cron job ou trigger do Firestore futuramente.
func processarFila(ctx context.Context, turmaId string) (*ResultadoProcessamento, error) {
	turmaRef := database.Client.Collection("turmas").Doc(turmaId)

	// --- 2. Busca a turma para saber quantas vagas restam ---
	turma, err := read.GetTurmaById(ctx, database.Client, turmaId)
	if err != nil {
		return nil, fmt.Errorf("turma não encontrada: %w", err)
	}

	vagasDisponiveis := turma.VagasTotais - turma.VagasOcupadas
	if vagasDisponiveis < 0 {
		vagasDisponiveis = 0
	}

	// --- 3. Busca todas as matrículas PENDENTES (status=false) da turma ---
	docs, err := database.Client.Collection("matriculas").
		Where("turmaId", "==", turmaId).
		Where("status", "==", false).
		Documents(ctx).
		GetAll()
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar matrículas pendentes: %w", err)
	}

	if len(docs) == 0 {
		return &ResultadoProcessamento{
			TurmaId:       turmaId,
			VagasTotais:   turma.VagasTotais,
			VagasOcupadas: turma.VagasOcupadas,
			Efetivados:    []string{},
			ListaEspera:   []string{},
			ProcessadoEm:  time.Now().Format(time.RFC3339),
		}, nil
	}

	// --- 4. Mapeia docs para structs e ordena por prioridadenota DESC ---
	// Em caso de empate na nota, desempata por DataSolicitacao ASC (quem pediu antes)
	type candidato struct {
		ref      *firestore.DocumentRef
		matricula models.Matricula
	}

	candidatos := make([]candidato, 0, len(docs))
	for _, doc := range docs {
		var m models.Matricula
		if err := doc.DataTo(&m); err != nil {
			log.Printf("Aviso: erro ao decodificar matrícula %s, pulando: %v", doc.Ref.ID, err)
			continue
		}
		candidatos = append(candidatos, candidato{ref: doc.Ref, matricula: m})
	}

	sort.Slice(candidatos, func(i, j int) bool {
		if candidatos[i].matricula.PrioridadeNota != candidatos[j].matricula.PrioridadeNota {
			// Maior nota → maior prioridade
			return candidatos[i].matricula.PrioridadeNota > candidatos[j].matricula.PrioridadeNota
		}
		// Empate: quem solicitou primeiro entra antes
		return candidatos[i].matricula.DataSolicitacao.Before(candidatos[j].matricula.DataSolicitacao)
	})

	// --- 5. Efetiva em transação: marca aprovados como status=true e atualiza vagasOcupadas ---
	efetivados := make([]string, 0)
	listaEspera := make([]string, 0)

	// Divide a fila em aprovados e reprovados antes da transação
	aprovados := candidatos
	reprovados := []candidato{}
	if vagasDisponiveis < len(candidatos) {
		aprovados = candidatos[:vagasDisponiveis]
		reprovados = candidatos[vagasDisponiveis:]
	}

	err = database.Client.RunTransaction(ctx, func(ctx context.Context, t *firestore.Transaction) error {
		// Re-lê a turma dentro da transação para evitar race condition
		turmaDoc, err := t.Get(turmaRef)
		if err != nil {
			return err
		}
		var turmaAtual models.Turma
		if err := turmaDoc.DataTo(&turmaAtual); err != nil {
			return err
		}

		vagasRestantes := turmaAtual.VagasTotais - turmaAtual.VagasOcupadas
		if vagasRestantes < 0 {
			vagasRestantes = 0
		}

		// Recalcula aprovados/reprovados com o valor mais recente de vagas
		totalAprovados := vagasRestantes
		if totalAprovados > len(candidatos) {
			totalAprovados = len(candidatos)
		}

		for i, c := range candidatos {
			if i < totalAprovados {
				// Aprova
				t.Update(c.ref, []firestore.Update{
					{Path: "status", Value: true},
					{Path: "dataSolicitacao", Value: time.Now()},
				})
				efetivados = append(efetivados, c.matricula.AlunoId)
			} else {
				// Mantém na lista de espera (status continua false)
				// Você pode opcionalmente deletar ou marcar como "listaEspera"
				listaEspera = append(listaEspera, c.matricula.AlunoId)
			}
		}

		// Atualiza vagas ocupadas na turma
		t.Update(turmaRef, []firestore.Update{
			{Path: "vagasOcupadas", Value: turmaAtual.VagasOcupadas + totalAprovados},
		})

		// Suprime variável não usada de reprovados (já tratada no loop acima)
		_ = reprovados

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("erro na transação de efetivação: %w", err)
	}

	return &ResultadoProcessamento{
		TurmaId:       turmaId,
		VagasTotais:   turma.VagasTotais,
		VagasOcupadas: turma.VagasOcupadas + len(efetivados),
		Efetivados:    efetivados,
		ListaEspera:   listaEspera,
		ProcessadoEm:  time.Now().Format(time.RFC3339),
	}, nil
}