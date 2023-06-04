package controller

import (
	"fmt"
	"log"
	"net/http"
	"pawAPIbackend/dto"
	"pawAPIbackend/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMessages(c *gin.Context) {

	submissionID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	userID, err := strconv.ParseUint(c.GetString("user_id"), 10, 64)

	log.Default().Println("GetMessages => submissionID= ", submissionID, "&&", "userID= ", userID)

	messages, err := service.GetMessages(submissionID, userID)

	if err != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	//Necessario pare returnar um array vazio em vez de null!!
	if messages == nil {
		messages = []dto.MessageResponseDTO{}
	}

	c.JSON(200, gin.H{
		"message":  "selected messages",
		"messages": messages,
	})
}

func AddMessage(c *gin.Context) {
	var messageCreateDTO dto.MessageCreateDTO

	userIDToken, err := strconv.ParseUint(c.GetString("user_id"), 10, 64)

	if err := c.ShouldBindJSON(&messageCreateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Invalid request body: failed %s", err.Error()),
			"error":   err.Error(),
		})
		return
	}

	if messageCreateDTO.ClinicalID != userIDToken {
		c.JSON(401, gin.H{
			"message": "error - Unauthorized ClinicalID from token is different from ClinicalID from body",
			"error":   "Unauthorized ClinicalID from token is different from ClinicalID from body",
		})
		return
	}

	message, err := service.AddMessage(messageCreateDTO)

	if err != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message":  "selected messages",
		"messages": message,
	})

}
