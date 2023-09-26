package service

import (
	"fmt"
	"encoding/json"
	"context"
	"errors"
	"sync"

	"github.com/Eli15x/client-crud/src/model"
	"github.com/Eli15x/client-crud/src/storage"
)

var (
	instanceClient CommandClient
	onceClient     sync.Once
)

type CommandClient interface {
	CreateClient(ctx context.Context, client model.Client) error
	GetClient(ctx context.Context, id string) (string, error)
	GetClients(ctx context.Context) (string, error)
	EditClient(ctx context.Context, client model.Client) error 
	DeleteClient(ctx context.Context, id string) error
}

type client struct{}

func GetInstanceClient() CommandClient {
	onceClient.Do(func() {
		instanceClient = &client{}
	})
	return instanceClient
}

func (c *client) CreateClient(ctx context.Context, client model.Client) error {

	err := storage.GetInstance().Insert(client)
	if err != nil {
		return errors.New("Create New Profile: problem to insert into Postgres")
	}
	return nil
}

func (c *client) EditClient(ctx context.Context, client model.Client) error {
	
	err := storage.GetInstance().Update(client)
	if err != nil {
		return errors.New("Edit Client: problem to update into Postgres")
	}

	return nil
}

func (c *client) DeleteClient(ctx context.Context, id string) error {
	
	err := storage.GetInstance().Delete(id)
	if err != nil {
		return errors.New("Delete Client: problem to delete into Postgres")
	}

	return nil
}


func (c *client) GetClient(ctx context.Context, id string) (string, error) {

	client, err := storage.GetInstance().GetClient(id)
	if err != nil {
		return "",errors.New("Get Client: problem to get cpf into Postgres")
	}

	b, err := json.Marshal(client)
    if err != nil {
        fmt.Printf("Error: %s", err)
        return "" , err
    }

	return string(b) , nil
}

func (c *client) GetClients(ctx context.Context) (string, error) {
	
	clients, err := storage.GetInstance().GetClients()
	if err != nil {
		return "", errors.New("Get Client: problem to get cpf into Postgres")
	}

	b, err := json.Marshal(clients)
    if err != nil {
        fmt.Printf("Error: %s", err)
        return "", err
    }

	return string(b), nil
}

