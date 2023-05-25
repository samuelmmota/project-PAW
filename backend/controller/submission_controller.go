package controller

import (
	"io/ioutil"
	"net/http"
	"pawAPIbackend/dto"
	"pawAPIbackend/entity"
	"pawAPIbackend/repository"
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
	// Read the Media file from the request body

	multipartFile, err := c.FormFile("media")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Media upload failed",
			"error":   err.Error(),
		})
		return
	}

	err = c.ShouldBind(&submissionCreateDTO)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Insersion error",
			"error":   err.Error(),
		})
		return
	}
	//vai buscar ao token(conveertido em string nos parematros) to jwt token
	userID, _ := strconv.ParseUint(c.GetString("user_id"), 10, 64)

	submission, err := service.InsertSubmission(submissionCreateDTO, multipartFile, userID)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Insersion error",
			"error":   err.Error(),
		})
		return
	}

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

func InsertImage(c *gin.Context) {
	// Read the image file from the request body
	imageFile, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Image upload failed",
			"error":   err.Error(),
		})
		return
	}

	// Open the uploaded file
	file, err := imageFile.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to open image file",
			"error":   err.Error(),
		})
		return
	}
	defer file.Close()

	// Read the file content
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to read image file",
			"error":   err.Error(),
		})
		return
	}

	// Create an instance of the ImageTest entity
	image := entity.ImageTest{
		ImageTest: fileBytes,
	}

	// Insert the image into the database
	insertedImage := repository.InsertImage(image)

	c.JSON(http.StatusOK, gin.H{
		"message": "Image inserted successfully",
		"image":   insertedImage,
	})
}

func GetImage(c *gin.Context) {
	// Get the image ID from the request parameters
	imageID := c.Param("id")

	// Retrieve the image from the database by ID
	image, err := repository.GetImageByID(imageID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Image not found",
			"error":   err.Error(),
		})
		return
	}

	// Set the appropriate response headers
	c.Header("Content-Type", "image/jpeg")
	c.Header("Content-Disposition", "attachment; filename=image.jpg")

	// Write the image bytes as the response body
	c.Writer.Write(image.ImageTest)
}

func GetAllImages(c *gin.Context) {
	// Retrieve all images from the database
	images, err := repository.GetAllImages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve images",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, images)
}
