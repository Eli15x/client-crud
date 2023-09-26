package model

type Client struct {
	Nome            string  `json:"nome,omitempty" bson:"-"`
	Sobrenome       string  `json:"sobrenome,omitempty" bson:"-"`
	Contato         string	`json:"contato,omitempty" bson:"-"`
	Endereco		string	`json:"endereco,omitempty" bson:"-"`
	DataNascimento 	string `json:"dataNascimento,omitempty" bson:"-"`
	Cpf				string  `json:"cpf,omitempty" bson:"-"`
}
