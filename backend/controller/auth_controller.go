package controller

import (
	"log"
	"pawAPIbackend/dto"
	"pawAPIbackend/service"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	loginDTO := dto.LoginDTO{}

	//Testar o mapeamento from dto
	err := c.ShouldBind(&loginDTO)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error binding from dto",
			"error":   err.Error(),
		})
		return
	}

	token, err := service.Login(loginDTO)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "error - invalid username or password",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "login success",
		"token":   token})

}

func EvaluateToken(c *gin.Context) {
	log.Default().Println(c.GetHeader("Authorization"))

	/*evaluateTokenDTO := dto.EvaluateTokenDTO{}

	//Testar o mapeamento from dto
	err := c.ShouldBind(&evaluateTokenDTO)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error binding from dto",
			"error":   err.Error(),
		})
		return
	}*/

	token := c.GetHeader("Authorization")
	_, err := service.ValidateToken(token)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "error - invalid token",
			"error":   err.Error(),
		})
		return
	}
	/*
		// Extract the user_id from the token
		userID, err := service.ExtractUserID(validToken)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "error",
				"error":   err.Error(),
			})
			return
		}

		// Create a new token
		regeneratedToken, err := service.CreateToken(userID)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "error regenerating token",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "login success",
			"token":   regeneratedToken})*/

	c.JSON(200, gin.H{
		"message": "valid token",
		"token":   token})
}
