package main

import (
	"pawAPIbackend/config"
	"pawAPIbackend/controller"
	"pawAPIbackend/entity"
	"pawAPIbackend/middleware"

	"github.com/gin-gonic/gin"
)

var Users []entity.User

// Set the desired buffer size for handling multipart requests
const MaxMultipartMemory = 128 << 20 // 64 MB

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	config.ConnectDB()

	defer config.CloseDb()

	router := gin.Default()

	// Set the maximum multipart memory size
	router.MaxMultipartMemory = MaxMultipartMemory

	router.Use(CORSMiddleware())

	v1 := router.Group("/paw/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", controller.Login)
			auth.POST("/refreshtoken", controller.EvaluateToken)
		}

		user := v1.Group("/user")
		{
			user.GET("/", controller.GetAllUsers) // sem AUTH / com DTO RESPONSE (ID, NAME, EMAIL)
			user.POST("/", controller.Register)
			user.GET("/:id", middleware.Authorized(), controller.Profile)          // com AUTH + OWNER / com DTO RESPONSE
			user.PUT("/:id", middleware.Authorized(), controller.UpdateProfile)    // com AUTH + OWNER
			user.DELETE("/:id", middleware.Authorized(), controller.DeleteAccount) // com AUTH + OWNER
		}

		clinical := user.Group("/clinical")
		{
			clinical.GET("/:id", middleware.Authorized(), controller.GetUserClinicals)
			clinical.GET("/clinicals/:id", middleware.Authorized(), controller.GetUserClinicals)
			clinical.GET("/patients/:id", middleware.Authorized(), controller.GetUserClinicals)

			clinical.GET("/submission/:id", middleware.Authorized(), controller.GetClinicalSubmissions)

			clinical.POST("/:id", middleware.Authorized(), controller.AddUserClinical)
			clinical.DELETE("/:id", middleware.Authorized(), controller.RemoveUserClinical)
		}

		//TEST
		v1.POST("/image", controller.InsertImage)
		v1.GET("/image/:id", controller.GetImage)
		v1.GET("/image", controller.GetAllImages)

		submission := v1.Group("/submission")
		{
			submission.GET("/", middleware.Authorized(), controller.GetAllSubmissions)
			submission.GET("/:id", middleware.Authorized(), controller.GetSubmission) // DTO -> AUTH + OWNER -> necessario para checks no frontend
			submission.POST("/", middleware.Authorized(), controller.InsertSubmission)
			submission.PUT("/:id", middleware.Authorized(), controller.UpdateSubmission) // DTO -> AUTH + OWNER
			submission.DELETE("/:id", middleware.Authorized(), controller.DeleteSubmission)
		}

	}
	router.Run(":3000")
}
