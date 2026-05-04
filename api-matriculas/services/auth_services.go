package services

import (
	"crypto/rand"
	"time"

	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/api-matriculas/auth"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/api-matriculas/models"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/api-matriculas/repository"
)

var secretKey = []byte("minha-chave-secreta-super-forte")

func GenerateAccessToken(req models.Auth) (models.SignedToken, error) {
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return models.SignedToken{}, err
	}

	institution, err := repository.FindInstitutionByID(req.InstitutionalKey)

	expiresAt := time.Now().Add(30 * time.Minute)

	tokenData := models.TokenData{
		InstitutionID: institution,
		ID:            req.Id,
		ExpiresAt:     expiresAt,
	}

	token, err := auth.SignToken(tokenData, secretKey)
	if err != nil {
		return models.SignedToken{}, err
	}

	return models.SignedToken{
		Token:     token,
		ExpiresAt: expiresAt,
	}, nil
}
