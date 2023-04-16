package main

import (
	"ufp.edu.pt/project-paw/config"
	"ufp.edu.pt/project-paw/controller"
	"ufp.edu.pt/project-paw/entity"
	"ufp.edu.pt/project-paw/middleware"

	"github.com/gin-gonic/gin"
)

var Users []entity.User

func main() {
	config.ConnectDB()

	defer config.CloseDb()

	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", controller.Login)
		}

		book := v1.Group("/book")
		{
			book.GET("/", controller.GetAllBooks)
			book.GET("/:id", controller.GetBook)
			book.POST("/", middleware.Authorized(), controller.InsertBook)
			book.PUT("/:id", middleware.Authorized(), controller.UpdateBook) // DTO -> AUTH + OWNER
			/**
			Error:
			2023/04/12 00:48:55 /home/samuelmota/Desktop/Modulo 5/repository/book_repository.go:35 Error 1452: Cannot add or update a child row: a foreign key constraint fails (`books`.`books`, CONSTRAINT `fk_books_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`))

			*/
			book.DELETE("/:id", middleware.Authorized(), controller.DeleteBook)
		}

		user := v1.Group("/user")
		{
			user.GET("/", controller.GetAllUsers) // SEM AUTH -> DTO RESPONSE (ID, NAME, EMAIL)
			user.POST("/", controller.Register)
			user.GET("/:id", middleware.Authorized(), controller.Profile)          // AUTH - OWNER -> DTO RESPONSE (ID, NAME, EMAIL, PROFILE PICTURE)
			user.PUT("/:id", middleware.Authorized(), controller.UpdateProfile)    // AUTH - OWNER
			user.DELETE("/:id", middleware.Authorized(), controller.DeleteAccount) // AUTH - OWNER
		}

	}
	router.Run(":3000")
}
