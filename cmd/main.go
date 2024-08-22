package main

import (
	"urlzipper/internal/configs"

	"github.com/gin-gonic/gin"
)

func main() {
	serverEngine := gin.Default()

	app := configs.NewApp()
	app.Setup(serverEngine)

	err := serverEngine.Run()
	if err != nil {
		panic(err)
	}
}
