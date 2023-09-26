## Ola seja bem vindo

para iniciar a aplicação rode os comandos

docker-compose up 

para criar a tabela :

CREATE TABLE client (
    cpf VARCHAR PRIMARY KEY,
    nome VARCHAR NOT NULL,
    sobrenome VARCHAR NOT NULL,
    contato VARCHAR NOT NULL,
    endereco VARCHAR NOT NULL,
    datanascimento DATE NOT NULL,
)


### formato para inserção

POST /client
body
{
    	"nome":    "elisa",
		"sobrenome": "castro",
        "contato": "19982847002",
        "endereco": "rua ze",
        "dataNascimento": "2002-14-02",
        "cpf": "00000000"
}

PUT /client
body
{
    	"nome":    "elisa",
		"sobrenome": "souza",
        "contato": "19982847002",
        "endereco": "rua ze",
        "dataNascimento": "2002-14-02",
        "cpf": "00000000"
}


DELETE /client/{cpf}
body
{

}

pegar cliente especifico
GET /client/{cpf}
body
{

}

pegar todos
GET /client/{cpf}
body
{

}