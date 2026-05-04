package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/api-matriculas/models"

	// Imports de debug
	"fmt"
	"io"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("MORRA"))
}

func TokenHandler(w http.ResponseWriter, r *http.Request) {

	// Inicializa o cliente para fazer a requisição POST
	client := &http.Client{}

	// Simulação de pegar os dados de login do front
	authtest := models.Auth{
		InstitutionalKey: "ChaveInstitucional123",
		Id:               "20260001",
		Password:         "senha123", //TESTE
	}

	// Transforma o modelo Auth em json
	jsonData, err := json.Marshal(authtest)

	// Criação da requisição http para o auth da api
	request, err := http.NewRequest(http.MethodPost, "http://localhost:8080/auth", bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Coloca os headers para informar que um json será enviado e tambem
	// coloca o valor da instituional key
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("institutional_key", authtest.InstitutionalKey)

	// Faz a requisição para a api
	resp, err := client.Do(request)

	/* Faz o if esperando que não dê em nada, porque se tiver que
	debugar o request do token inteiro...*/
	if err != nil {
		http.Error(w, "Sonhos foram destruidos...", http.StatusBadRequest)
	}

	// Fecha o body
	defer resp.Body.Close()

	// Para Debug

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	w.Write([]byte("Token = " + string(body)))
}
