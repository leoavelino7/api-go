package entities

import (
	"github.com/google/uuid"
)

type Cep struct {
	ID string `json:"id"`
	Cep string `json:"cep"`
	Street string `json:"street"`
	Complement string `json:"complement"`
	Neighborhood string `json:"neighborhood"`
	City string `json:"city"`
	State string `json:"state"`
	Ibge string `json:"ibge"`
}

func NewCep() *Cep {
	cep := Cep{
		ID: uuid.New().String(),
	}

	return &cep
}