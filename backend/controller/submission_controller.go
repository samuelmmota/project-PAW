package controller

import (
	"pawAPIbackend/dto"
	"pawAPIbackend/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllSubmissions(c *gin.Context) {

	c.JSON(200, gin.H{
		"message":     "select submissions",
		"submissions": service.GetAllSubmissions(),
	})
}

func InsertSubmission(c *gin.Context) {
	var submissionCreateDTO dto.SubmissionCreateDTO
	err := c.ShouldBind(&submissionCreateDTO)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	//vai buscar ao token(conveertido em string nos parematros) to jwt token
	userID, _ := strconv.ParseUint(c.GetString("user_id"), 10, 64)

	submission := service.InsertSubmission(submissionCreateDTO, userID)
	c.JSON(200, gin.H{
		"message":    "insert submission",
		"submission": submission,
	})
}

func GetSubmission(c *gin.Context) {
	submissionID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	submissionResponseDTO, err := service.GetSubmission(submissionID)
	userID, _ := strconv.ParseUint(c.GetString("user_id"), 10, 64)

	if err != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	if !service.IsAllowedToEdit(userID, submissionID) {
		c.JSON(401, gin.H{
			"message": "you do not have the permission - you are not the owner of this submission",
		})
		return
	}

	c.JSON(200, gin.H{
		"message":    "select submission",
		"submission": submissionResponseDTO,
	})
}

func UpdateSubmission(c *gin.Context) {
	submissionID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	userID, _ := strconv.ParseUint(c.GetString("user_id"), 10, 64)

	var submission dto.SubmissionUpdateDTO

	err := c.ShouldBind(&submission)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	if !service.IsAllowedToEdit(userID, submissionID) {
		c.JSON(401, gin.H{
			"message": "you do not have the permission - you are not the owner of this submission",
		})
		return
	}

	submission.ID = submissionID
	submission.UserID = userID
	submissionResponse, err := service.UpdateSubmission(submission)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message":    "update submission",
		"submission": submissionResponse,
	})
}

func DeleteSubmission(c *gin.Context) {
	submissionID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	userID, _ := strconv.ParseUint(c.GetString("user_id"), 10, 64)

	if !service.IsAllowedToEdit(userID, submissionID) {
		c.JSON(401, gin.H{
			"message": "you do not have the permission - you are not the owner of this submission",
		})
		return
	}

	err := service.DeleteSubmission(submissionID)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "submission deleted",
	})
}
