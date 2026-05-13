package models

import "time"

type Usuario struct {
	Matricula string `json:"matricula" firestore:"matricula"`

	Senha string `json:"-" firestore:"senha"`

	NivelAcesso string    `json:"nivelAcesso" firestore:"nivelAcesso"`
	Created     time.Time `json:"created" firestore:"created"`
}
