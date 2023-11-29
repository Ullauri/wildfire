package server

import (
	"wildfire/internal/clients"

	"github.com/gin-gonic/gin"
)

func AddIndexRoutes(g *gin.RouterGroup, namesClient clients.NamesClient, jokesClient clients.JokesClient) {
	// TODO: implement proper error handling via app errors; for now just return 500
	g.GET("/", func(ctx *gin.Context) {
		randomName, err := namesClient.GetRandomName(ctx)
		if err != nil {
			ctx.JSON(500, gin.H{"namesClientError": err.Error()})
			return
		} else if randomName == nil {
			ctx.JSON(500, gin.H{"namesClientError": "randomName was nil"})
			return
		}

		joke, err := jokesClient.GetJoke(ctx, &randomName.FirstName, &randomName.LastName)
		if err != nil {
			ctx.JSON(500, gin.H{"jokesClientError": err.Error()})
			return
		} else if joke == nil {
			ctx.JSON(500, gin.H{"jokesClientError": "joke was nil"})
			return
		}

		ctx.String(200, joke.Value.Joke)
	})

	g.GET("/health", func(ctx *gin.Context) {
		ctx.String(200, "OK")
	})
}
