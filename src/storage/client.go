//This example uses the ORM jet
package storage

import (
	"fmt"
	"database/sql"

	"github.com/Eli15x/client-crud/src/model"
	_ "github.com/lib/pq"

	"time"
	//"os"
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
	var err error
	connStr := "postgres://tuto:admingres@localhost/postgres?sslmode=disable"
	c.db, err = sql.Open("postgres", connStr)
	
	if err != nil {
		return err
	}
	
	if err = c.db.Ping(); err != nil {
		return err
	}
	// this will be printed in the terminal, confirming the connection to the database
	fmt.Println("The database is connected")
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
	date, _ := time.Parse("2006-01-02", client.DataNascimento)
	_, err :=  c.db.Query(
        "INSERT INTO client(cpf, nome, sobrenome,contato,endereco,datanascimento) VALUES($1, $2, $3, $4, $5, $6)",
        client.Cpf, client.Nome, client.Sobrenome, client.Contato, client.Endereco, date)
		
		if err == nil {
			return err
		}

	return nil
}

func (c *clientSql) Update(client model.Client) error {
	date, _ := time.Parse("2006-01-02", client.DataNascimento)
	_, err :=
		  c.db.Exec("UPDATE client SET nome=$1, sobrenome=$2, contato=$3, endereco=$4, datanascimento=$5 WHERE cpf=$6",
		  client.Nome, client.Sobrenome, client.Contato, client.Endereco, date, client.Cpf)
	
	return err
}


func (c *clientSql) GetClients() ([]model.Client, error) {
	rows, err := c.db.Query(
		  "SELECT nome, sobrenome,contato,endereco,datanascimento,cpf FROM client")
		  
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
	fmt.Println(interfaceSql)
	var client model.Client
	err := c.db.QueryRow("SELECT nome, sobrenome,contato,endereco,datanascimento,cpf FROM client WHERE cpf=$1",
	interfaceSql).Scan(&client.Nome, &client.Sobrenome, &client.Contato, &client.Endereco, &client.DataNascimento, &client.Cpf)

	if err != nil {
		return client, err
	}

	return client,nil
}




