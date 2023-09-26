## Ola seja bem vindo

para iniciar a aplicação rode os comandos

sudo docker pull rabbitmq
sudo docker pull postgres

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

## Pontuações de melhorias e finalizações

a ideia da pasta docker-entrypoint era ser um script automatico para não ser necessário inserir o create da tabela diretamente na mão, que começei a implementação.

tanto quanto a implementação do golang no docker.

iria ajustar alguns nomes colocando tudo em padrão para ingles
e alguns pontos de arquitetura, como também alguns tratamentos de erros.

gostaria também de fazer a parte referente ao postgres generica.

o rabbitmq ainda está em finalização

colocaria também testes unitarios
e mudaria o retorno referente ao formato do /get clients, que não constitui o padrão correto.
