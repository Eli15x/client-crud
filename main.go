package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/Eli15x/client-crud/src/handler"
	"github.com/Eli15x/client-crud/src/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	//Context
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err := godotenv.Load(".env")
    if err != nil {
        fmt.Errorf("Error loading .env file")
    }

	if err := storage.GetInstance().Connect(); err != nil {
		fmt.Errorf("[PostgreSQL] problem to Connect : %s \n", err, "")
	}

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}
	config.AllowMethods = []string{"*"}
	config.AllowCredentials = true

	router.Use(cors.New(config))

	router.POST("/client", handler.CreateClient)
	router.DELETE("/client", handler.DeleteClient)
	router.PUT("/client/:cpf", handler.EditClient)
	router.GET("/client/:cpf", handler.GetClient)
	router.GET("/client", handler.GetClients)


	router.Run(":1323")
}
