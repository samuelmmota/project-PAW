package controller

import (
	"pawAPIbackend/dto"
	"pawAPIbackend/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "select books",
		"books":   service.GetAllBooks(),
	})
}

func InsertBook(c *gin.Context) {
	var book dto.BookCreatedDTO
	err := c.ShouldBind(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	//vai buscar ao token(conveertido em string nos parematros) to jwt token
	userID, _ := strconv.ParseUint(c.GetString("user_id"), 10, 64)

	b := service.InsertBook(book, userID)
	c.JSON(200, gin.H{
		"message": "insert book",
		"book":    b,
	})
}

func GetBook(c *gin.Context) {
	bookID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	bookResponseDTO, err := service.GetBook(bookID)
	userID, _ := strconv.ParseUint(c.GetString("user_id"), 10, 64)

	if err != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	if !service.IsAllowedToEdit(userID, bookID) {
		c.JSON(401, gin.H{
			"message": "you do not have the permission - you are not the owner of this book",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "select book",
		"book":    bookResponseDTO,
	})
}

func UpdateBook(c *gin.Context) {
	bookID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	userID, _ := strconv.ParseUint(c.GetString("user_id"), 10, 64)

	var book dto.BookUpdateDTO

	err := c.ShouldBind(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	if !service.IsAllowedToEdit(userID, bookID) {
		c.JSON(401, gin.H{
			"message": "you do not have the permission - you are not the owner of this book",
		})
		return
	}

	book.ID = bookID
	bookResponse, err := service.UpdateBook(book)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "update book",
		"book":    bookResponse,
	})
}

func DeleteBook(c *gin.Context) {
	bookID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	userID, _ := strconv.ParseUint(c.GetString("user_id"), 10, 64)

	if !service.IsAllowedToEdit(userID, bookID) {
		c.JSON(401, gin.H{
			"message": "you do not have the permission - you are not the owner of this book",
		})
		return
	}

	err := service.DeleteBook(bookID)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "book deleted",
	})
}
