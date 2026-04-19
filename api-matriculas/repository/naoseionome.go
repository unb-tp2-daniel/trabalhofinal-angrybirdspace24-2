package repository

import "fmt"

func FindInstitutionByID(institutionalKey string) (string, error) {
	// Simulação de busca no banco de dados
	if institutionalKey != "ChaveInstitucional123" {
		return "", fmt.Errorf("instituição não encontrada")
	}
	return "Unb", nil
}
