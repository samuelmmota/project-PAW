package controller

import (
	"fmt"
	"pawAPIbackend/dto"
	"pawAPIbackend/entity"
	"pawAPIbackend/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// need upper case in begining to export the function on the package
func GetAllUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "select all Users",
		"users":   service.GetAllUsers(),
	})
}

func Register(c *gin.Context) {
	var user entity.User

	err := c.ShouldBind(&user)

	fmt.Println(user)
	fmt.Printf("Received user: %+v\n", user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	if user.Email == "" || user.Password == "" {
		c.JSON(400, gin.H{
			"message": "error",
			"error":   "email or password is empty",
		})
		return
	}

	if len(user.Password) < 6 {
		c.JSON(400, gin.H{
			"message": "error",
			"error":   "password must be at least 6 characters",
		})
		return
	}

	if user, _ := service.CheckEmail(user.Email); user.ID != 0 {
		c.JSON(400, gin.H{
			"message": "error",
			"error":   "User already exists",
		})
		return
	}

	user, err = service.Register(user)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Registered User",
		"user":    user,
	})
}

func Profile(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	userID_token, err := strconv.ParseUint(c.GetString("user_id"), 10, 64)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !service.IsAllowedUser(userID, userID_token) {
		c.JSON(401, gin.H{
			"message": "you do not have the permissions to see the user \n- you are not the owner",
		})
		return
	}

	userProfileResponseDTO, err := service.Profile(userID)

	if err != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message":      "selected user profile" + c.Param("id"),
		"user":         userProfileResponseDTO,
		"userID":       userID,
		"userID_Token": userID_token,
	})
}

func UpdateProfile(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	userID_token, err := strconv.ParseUint(c.GetString("user_id"), 10, 64)

	var user dto.UserUpdateDTO
	err = c.ShouldBind(&user)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !service.IsAllowedUser(userID, userID_token) {
		c.JSON(401, gin.H{
			"message": "you do not have the permissions to edit the user \n- you are not the owner",
		})
		return
	}

	user.ID = userID
	userResponse, err3 := service.UpdateProfile(user)

	if err3 != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   err3.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Updated user",
		"user":    userResponse,
	})
}

func DeleteAccount(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	userID_token, err := strconv.ParseUint(c.GetString("user_id"), 10, 64)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !service.IsAllowedUser(userID, userID_token) {
		c.JSON(401, gin.H{
			"message": "you do not have the permissions to edit the user \n- you are not the owner",
		})
		return
	}

	err = service.DeleteAccount(userID)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "user deleted",
	})
}
