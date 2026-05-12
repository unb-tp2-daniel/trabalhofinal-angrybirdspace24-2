package models

import (
	"time"
)

type Auth struct {
	InstitutionalKey string `json:"institutional_key"`
	Id               string `json:"id"`
}
type TokenData struct {
	ID            string
	InstitutionID string
	ExpiresAt     time.Time
}

type SignedToken struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}
