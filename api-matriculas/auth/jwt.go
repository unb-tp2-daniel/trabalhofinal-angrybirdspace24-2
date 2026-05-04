package auth

import (
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/api-matriculas/models"

	"github.com/golang-jwt/jwt/v5"
)

func SignToken(tokenData models.TokenData, secretKey []byte) (string, error) {
	claims := jwt.MapClaims{
		"institution_id": tokenData.InstitutionID,
		"id":             tokenData.ID,
		"exp":            tokenData.ExpiresAt.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
