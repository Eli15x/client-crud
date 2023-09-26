package service

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/Eli15x/client-crud/src/model"
	"github.com/Eli15x/client-crud/src/storage"
	"github.com/fatih/structs"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	instanceClient CommandClient
	onceClient     sync.Once
)

type CommandClient interface {
	CreateClient(ctx context.Context, client model.Client) error
	GetClient(ctx context.Context, id string) ([]bson.M, error)
	GetClients(ctx context.Context, id string) ([]bson.M, error)
	EditClient(ctx context.Context, id string, client model.Client) ([]bson.M, error)
	DeleteClient(ctx context.Context, id string) ([]bson.M, error)
}

type client struct{}

func GetInstanceClient() CommandClient {
	onceClient.Do(func() {
		instanceClient = &client{}
	})
	return instanceClient
}

func (c *client) CreateClient(ctx context.Context, client model.Client) error {

	ClientInsert := structs.Map(client)

	_, err := storage.GetInstance().Insert(ctx, "profile", profileInsert)
	if err != nil {
		return errors.New("Create New Profile: problem to insert into Postgrees")
	}
	return nil
}

func (c *client) EditClient(ctx context.Context, id string, client model.Client) ([]bson.M, error) {
	var profile models.Profile

	userId := map[string]interface{}{"UserId": id}


	return result, nil
}

func (c *client) DeleteClient(ctx context.Context, id string) ([]bson.M, error) {
	var profile models.Profile

	userId := map[string]interface{}{"UserId": id}


	return result, nil
}


func (c *client) GetClient(ctx context.Context, id string) ([]bson.M, error) {
	var profile models.Profile

	userId := map[string]interface{}{"UserId": id}


	return result, nil
}

func (c *client) GetClients(ctx context.Context, id string) ([]bson.M, error) {
	var profile models.Profile

	userId := map[string]interface{}{"UserId": id}


	return result, nil
}
