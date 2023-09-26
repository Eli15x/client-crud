package model

import (
	"time"
)

type Contatolient struct {
	Nome            string  `json:"nome,omitempty" bson:"-"`
	Sobrenome       string  `json:"sobrenome,omitempty" bson:"-"`
	Contato         string	`json:"contato,omitempty" bson:"-"`
	Endereco		string	`json:"endereco,omitempty" bson:"-"`
	DataNascimento time.Time  `json:"dataNacimento,omitempty" bson:"-"`
	CPF				string  `json:"cpf,omitempty" bson:"-"`
}
