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

func GetUserClinicals(c *gin.Context) {

	patientID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	//userID, err := strconv.ParseUint(c.GetString("user_id"), 10, 64)
	log.Default().Println("patientID= ", patientID)

	/*if patientID != userID {
		c.JSON(401, gin.H{
			"message": "error - Unauthorized",
			"error":   "Unauthorized",
		})
		return
	}
	*/

	if err != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	clinicals := service.GetClinicals(patientID) //userID)

	c.JSON(200, gin.H{
		"message":   "selected clinicals",
		"clinicals": clinicals,
	})
}

func AddUserClinical(c *gin.Context) {
	var clinicalCreateDTO dto.ClinicalCreateDTO

	patientID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	userID, err := strconv.ParseUint(c.GetString("user_id"), 10, 64)
	log.Default().Println("patientID= ", patientID)

	if patientID != userID {
		c.JSON(401, gin.H{
			"message": "error - Unauthorized",
			"error":   "Unauthorized",
		})
		return
	}

	if err := c.ShouldBindJSON(&clinicalCreateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Invalid request body: failed %s", err.Error()),
			"error":   err.Error(),
		})
		return
	}

	clinicalResponseDTO, err := service.AddUserClinical(patientID, clinicalCreateDTO)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Clinical creation: failed %s", err.Error()),
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Clinical added successfully",
		"clinical": clinicalResponseDTO,
	})
}

func RemoveUserClinical(c *gin.Context) {}

/*
func RemoveUserClinical(c *gin.Context) {
	var clinical entity.Clinical
	err := c.ShouldBind(&clinical)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Insersion error",
			"error":   err.Error(),
		})
		return
	}

	err2 := service.RemoveClinical(clinical)

	if err2 != nil {
		c.JSON(400, gin.H{
			"message": "Insersion error",
			"error":   err2.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "remove clinical",
	})
}
*/

func GetClinicalSubmissions(c *gin.Context) {

	clinicalID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	//userID, err := strconv.ParseUint(c.GetString("user_id"), 10, 64)
	log.Default().Println("clinicalID= ", clinicalID)

	/*if patientID != userID {
		c.JSON(401, gin.H{
			"message": "error - Unauthorized",
			"error":   "Unauthorized",
		})
		return
	}
	*/

	if err != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	clinicals, err := service.GetClinicalSubmissions(clinicalID) //userID)

	if err != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message":   "selected clinicals",
		"clinicals": clinicals,
	})
}
