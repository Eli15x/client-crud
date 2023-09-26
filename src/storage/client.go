//This example uses the ORM jet
package storage

import (
	"database/sql"

	"github.com/Eli15x/client-crud/src/model"
	"github.com/go-sql-driver/mysql"

	"fmt"
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
	GetClient(interfaceSql interface{}) error 
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
   connStr := "postgresql://" +os.Getenv("POSTGRES_USER")+":"+ os.Getenv("POSTGRES_PASSWORD")+"@"+os.Getenv("POSTGRES_DB")+"/todos?sslmode=disable"
   // Connect to database
   db, err := sql.Open("postgres", connStr)
   if err != nil {
       log.Fatal(err)
   }
}

func (c *clientSql) Delete(interfaceSql interface{}) error {
	_, err := c.db.Exec("DELETE FROM client WHERE cpf=$1", interfaceSql)
	if err != nil {
		return err
	}

	return nil
}

func (c *clientSql) Insert(client model.Client) error {
	err := db.QueryRow(
        "INSERT INTO client(nome, sobrenome,contato,endereco,dataNacimento,cpf) VALUES($1, $2, $3, $4, $5, $6) RETURNING id",
        client.Nome, client.Sobrenome, client.Contato, client.Endereco, client.DataNascimento, client.Cpf)
		
		if err != nil {
			return err
		}
}

func (c *clientSql) Update(client model.Client) error {
	_, err :=
		  db.Exec("UPDATE client SET nome=$1, sobrenome=$2, contato=$3, endereco=$4, dataNascimento=$5 WHERE id=$6",
		  client.Nome, client.Sobrenome, client.Contato, client.Endereco, client.DataNascimento, client.Cpf)
	
	return err
}


func (c *clientSql) GetClients() ([]model.Client, error) {
	rows, err := db.Query(
		  "SELECT nome, sobrenome,contato,endereco,dataNacimento,cpf FROM client")
		  
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	client := []model.Client{}
	
	for rows.Next() {
		var c client
		if err := rows.Scan(&client.Nome, &client.Sobrenome, &client.Contato, &client.Endereco, &client.DataNascimento ,&client.Cpf); err != nil {
			return
			nil, err
		}
		client = append(client, c)
	}
	
	return products, nil
}

func (c *clientSql) GetClient(interfaceSql interface{}) error {
	var client model.Client
	_, err := db.QueryRow("SELECT nome, sobrenome,contato,endereco,dataNacimento,cpf FROM client WHERE id=$1",
	interfaceSql).Scan(&client.Nome, &client.Sobrenome, &client.Contato, &client.Endereco, &client.DataNascimento, &client.Cpf)

	return client
}




