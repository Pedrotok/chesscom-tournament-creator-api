package main

import (
	"github.com/gin-gonic/gin"

	"chesscom-tournament-creator-api/config"
	"chesscom-tournament-creator-api/handler"
)

func main() {
	config.ConnectDB()

	r := gin.Default()
	r.POST("/createPlayer", handler.CreatePlayer)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
