package main

import (
	"pawAPIbackend/config"
	"pawAPIbackend/controller"
	"pawAPIbackend/entity"
	"pawAPIbackend/middleware"

	"github.com/gin-gonic/gin"
)

var Users []entity.User

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
	router.Use(CORSMiddleware())

	v1 := router.Group("/paw/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", controller.Login)
		}

		user := v1.Group("/user")
		{
			user.GET("/", controller.GetAllUsers) // sem AUTH / com DTO RESPONSE (ID, NAME, EMAIL)
			user.POST("/", controller.Register)
			user.GET("/:id", middleware.Authorized(), controller.Profile)          // com AUTH + OWNER / com DTO RESPONSE
			user.PUT("/:id", middleware.Authorized(), controller.UpdateProfile)    // com AUTH + OWNER
			user.DELETE("/:id", middleware.Authorized(), controller.DeleteAccount) // com AUTH + OWNER
		}
		/*
			book := v1.Group("/book")
			{
				book.GET("/", controller.GetAllBooks)
				book.GET("/:id", middleware.Authorized(), controller.GetBook) // DTO -> AUTH + OWNER -> necessario para checks no frontend
				book.POST("/", middleware.Authorized(), controller.InsertBook)
				book.PUT("/:id", middleware.Authorized(), controller.UpdateBook) // DTO -> AUTH + OWNER
				/**
				Error:
				2023/04/12 00:48:55 /home/samuelmota/Desktop/Modulo 5/repository/book_repository.go:35 Error 1452: Cannot add or update a child row: a foreign key constraint fails (`books`.`books`, CONSTRAINT `fk_books_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`))

		*/
		//	book.DELETE("/:id", middleware.Authorized(), controller.DeleteBook)
		//	}

	}
	router.Run(":3000")
}
