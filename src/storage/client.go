//This example uses the ORM jet
package storage

import (
	"database/sql"

	"github.com/Eli15x/client-crud/src/model"
	_ "github.com/lib/pq"
	"log"

	"os"
	"sync"
)

var (
	instanceClient CommandClient
	onceClient     sync.Once
)

type CommandClient interface {
	Connect() error
	Delete(interfaceSql interface{}) error
	Insert(client model.Client) error
	Update(client model.Client) error 
	GetClients() ([]model.Client, error)
	GetClient(interfaceSql interface{}) (model.Client,error)
}

type clientSql struct {
	db *sql.DB
}

func GetInstance() CommandClient {
	onceClient.Do(func() {
		instanceClient = &clientSql{}
	})
	return instanceClient
}

func (c *clientSql) Connect() error {
   connStr := "postgresql://" + os.Getenv("POSTGRES_USER")+":"+ os.Getenv("POSTGRES_PASSWORD")+"@"+os.Getenv("POSTGRES_DB")+"/todos?sslmode=disable"
   // Connect to database
   db, err := sql.Open("postgres", connStr)
   if err != nil {
       return err
   }

   err = db.Ping()
   if err != nil {
		return err
   }

   return nil
}

func (c *clientSql) Delete(interfaceSql interface{}) error {
	_, err := c.db.Exec("DELETE FROM client WHERE cpf=$1", interfaceSql)
	if err != nil {
		return err
	}

	return nil
}

func (c *clientSql) Insert(client model.Client) error {
	err := c.db.QueryRow(
        "INSERT INTO client(nome, sobrenome,contato,endereco,dataNacimento,cpf) VALUES($1, $2, $3, $4, $5, $6) RETURNING id",
        client.Nome, client.Sobrenome, client.Contato, client.Endereco, client.DataNascimento, client.Cpf)
		
		if err != nil {
			log.Fatal(err)
		}

	return nil
}

func (c *clientSql) Update(client model.Client) error {
	_, err :=
		  c.db.Exec("UPDATE client SET nome=$1, sobrenome=$2, contato=$3, endereco=$4, dataNascimento=$5 WHERE id=$6",
		  client.Nome, client.Sobrenome, client.Contato, client.Endereco, client.DataNascimento, client.Cpf)
	
	return err
}


func (c *clientSql) GetClients() ([]model.Client, error) {
	rows, err := c.db.Query(
		  "SELECT nome, sobrenome,contato,endereco,dataNacimento,cpf FROM client")
		  
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	clients := []model.Client{}
	
	for rows.Next() {
		var c model.Client
		if err := rows.Scan(&c.Nome, &c.Sobrenome, &c.Contato, &c.Endereco, &c.DataNascimento ,&c.Cpf); err != nil {
			return nil, err
		}
		clients = append(clients, c)
	}
	
	return clients, nil
}

func (c *clientSql) GetClient(interfaceSql interface{}) (model.Client,error) {
	var client model.Client
	err := c.db.QueryRow("SELECT nome, sobrenome,contato,endereco,dataNacimento,cpf FROM client WHERE cpf=$1",
	interfaceSql).Scan(&client.Nome, &client.Sobrenome, &client.Contato, &client.Endereco, &client.DataNascimento, &client.Cpf)

	if err != nil {
		return client, err
	}

	return client,nil
}




