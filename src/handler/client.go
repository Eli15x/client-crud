package handler

import (
	"context"
	"encoding/json"

	"net/http"
	"github.com/Eli15x/client-crud/src/service"
	"github.com/Eli15x/client-crud/src/model"
	"github.com/gin-gonic/gin"
)



func CreateClient(c *gin.Context) {

	var client model.Client
	err := json.NewDecoder(c.Request.Body).Decode(&client)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	if client.Nome == "" {
		c.String(http.StatusBadRequest, "Create Client Error: nome not find")
		return
	}

	if client.Sobrenome == "" {
		c.String(400, "Create Client Error: sobrenome not find")
		return
	}

	if client.Cpf == "" {
		c.String(400, "Create Client Error: cpf not find")
		return
	}

	if client.DataNascimento == "" {
		c.String(400, "Create Client  Error: dataNacimento not find")
		return
	}

	if client.Contato == "" {
		c.String(400, "Create Client Error: contato not find")
		return
	}

	if client.Endereco == "" {
		c.String(400, "Create Client Error: endereco not find")
		return
	}

	err = service.GetInstanceClient().CreateClient(context.Background(), client)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(http.StatusOK, "" )
}

func EditClient(c *gin.Context) {

	var client model.Client
	err := json.NewDecoder(c.Request.Body).Decode(&client)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	if client.Nome == "" {
		c.String(http.StatusBadRequest, "Edit Client Error: nome not find")
		return
	}

	if client.Sobrenome == "" {
		c.String(400, "Edit Client Error: sobrenome not find")
		return
	}

	if client.Cpf == "" {
		c.String(400, "Edit Client Error: cpf not find")
		return
	}

	if client.DataNascimento == "" {
		c.String(400, "Edit Client  Error: dataNascimento not find")
		return
	}

	if client.Contato == "" {
		c.String(400, "Edit Client Error: contato not find")
		return
	}

	if client.Endereco == "" {
		c.String(400, "Edit Client Error: contato not find")
		return
	}

	err = service.GetInstanceClient().EditClient(context.Background(), client)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}

func DeleteClient(c *gin.Context) {

	cpf := c.Param("cpf")

	if cpf == ""{
		c.String(400, "Delete Client: none cpf informed")
		return 
	}

	err := service.GetInstanceClient().DeleteClient(context.Background(), cpf)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}

func GetClient(c *gin.Context) {

	cpf := c.Param("cpf")

	if cpf == ""{
		c.String(400, "Get Client: none cpf informed")
		return 
	}

	client, err := service.GetInstanceClient().GetClient(context.Background(), cpf)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(http.StatusOK, client)
}


func GetClients(c *gin.Context) {

	clients, err := service.GetInstanceClient().GetClients(context.Background())
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.JSON(http.StatusOK, clients) //return list
}