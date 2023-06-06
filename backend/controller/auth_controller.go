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

	token := c.GetHeader("Authorization")
	_, err := service.ValidateToken(token)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "error - invalid token",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "valid token",
		"token":   token})
}
