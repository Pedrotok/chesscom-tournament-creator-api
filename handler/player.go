package handler

import (
	"chesscom-tournament-creator-api/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePlayer(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("error reading from request body on create player")
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "error reading from request body on create player",
		})
		return
	}

	var player model.Player
	err = json.Unmarshal(jsonData, &player)
	if err != nil {
		fmt.Println("error parsing request body on create player")
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  "error",
			"message": "error parsing request body on create player",
		})
		return
	}

	// service.CreatePlayer(player)

	c.JSON(http.StatusOK, gin.H{
		"status": "posted",
	})
}
