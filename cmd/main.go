package main

import (
	"log/slog"
	"urlzipper/internal/configs"

	"github.com/gin-gonic/gin"
)

func main() {
	serverEngine := gin.Default()

	app := configs.NewApp()
	app.Setup(serverEngine)

	err := serverEngine.Run()
	if err != nil {
		slog.Error("Application failed to start")
		panic(err)
	} else {
		slog.Info("Application started successfully")
	}
}
