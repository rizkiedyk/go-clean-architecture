package main

import (
	"go-clean-architecture/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	// timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	gin.Run(env.ServerAddress)
}
