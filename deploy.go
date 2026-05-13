package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Recebe o nome da função (name) e monta a lista de argumentos.
func deployFunction(name string) error {
	args := []string{
		"functions",
		"deploy",
		name,
		"--gen2",
		"--runtime=go126",
		"--region=southamerica-east1",
		"--source=.",
		"--entry-point=" + name,
		"--trigger-http",
		"--allow-unauthenticated",
		"--set-build-env-vars=GOFLAGS=-buildvcs=false",
		"--project=matriculas242",
	}

	cmd := exec.Command("gcloud", args...)

	// Redireciona a saída do gcloud para o seu terminal do VS Code,
	// assim você vê a barra de progresso do Google Cloud.
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Executa o comando
	return cmd.Run()
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Erro: Faltou informar o nome da função.")
		fmt.Println("Como usar: go run deploy.go NomeDaSuaFuncao")
		fmt.Println("Exemplo: go run deploy.go CriarTurma")
		return
	}

	// Pega exatamente a palavra que você digitou após o nome do script
	nomeDaFuncao := os.Args[1]

	fmt.Printf("Iniciando o deploy: %s...\n", nomeDaFuncao)

	// Chama a função que faz o trabalho pesado
	err := deployFunction(nomeDaFuncao)

	if err != nil {
		fmt.Printf("Erro ao realizar o deploy %s: %v\n", nomeDaFuncao, err)
		return
	}

	fmt.Printf("Deploy da função %s concluído com sucesso!\n", nomeDaFuncao)
}
