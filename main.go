package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Eli15x/client-crud/src/handlers"
	"github.com/Eli15x/client-crud/src/storage"
	"github.com/bugsnag/bugsnag-go/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	//Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err := godotenv.Load(".env")
    if err != nil {
        fmt.Errorf("Error loading .env file")
    }

	//defer cancel()
	if err := client.GetInstance().Connect(); err != nil {
		log.Infof("[PostgreSQL] problem to Connect : %s \n", err, "")
	}

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}
	config.AllowMethods = []string{"*"}
	config.AllowCredentials = true

	router.Use(cors.New(config))

	router.POST("/client", handlers.CreateClient)
	router.DELETE("/client", handlers.DeleteClient)
	router.UPDATE("/client/:cpf", handlers.UpdateClient)
	router.GET("/client/:cpf", handlers.GetClient)
	router.GET("/client", handlers.GetClients)


	router.Run(":1323")
}
