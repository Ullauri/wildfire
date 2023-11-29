package main

import (
	"fmt"
	"wildfire/internal/clients"
	"wildfire/internal/common"
	"wildfire/internal/server"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func main() {
	cfg := common.Config()

	zerolog.SetGlobalLevel(zerolog.Level(cfg.LogLevel))

	router := gin.Default()

	httpClient := clients.NewHttpClient()
	namesClient := clients.NewNamesClient(httpClient)
	jokesClient := clients.NewJokesClient(httpClient)

	root := router.Group("")
	{
		server.AddIndexRoutes(root, namesClient, jokesClient)
	}

	router.Run(fmt.Sprintf(":%d", cfg.Port))
}
