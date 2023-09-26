package model

import (
	"time"
)

type Client struct {
	Nome            string  `json:"nome,omitempty" bson:"-"`
	Sobrenome       string  `json:"sobrenome,omitempty" bson:"-"`
	Contato         string	`json:"contato,omitempty" bson:"-"`
	Endereco		string	`json:"endereco,omitempty" bson:"-"`
	DataNascimento time.Time  `json:"dataNacimento,omitempty" bson:"-"`
	Cpf				string  `json:"cpf,omitempty" bson:"-"`
}
