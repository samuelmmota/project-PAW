package repository

import (
	"errors"
	"pawAPIbackend/config"
	"pawAPIbackend/entity"
)

func InsertBook(book entity.Book) entity.Book {
	config.Db.Save(&book)
	config.Db.Preload("User").Find(&book)

	return book
}

func GetAllBooks() []entity.Book {
	var books []entity.Book
	config.Db.Preload("User").Find(&books)

	return books
}

func GetBook(bookID uint64) (entity.Book, error) {
	var book entity.Book
	config.Db.Preload("User").First(&book, bookID)
	if book.ID != 0 {
		return book, nil
	}

	return book, errors.New("book do not exists")
}

func UpdateBook(book entity.Book) (entity.Book, error) {
	if _, err := GetBook(book.ID); err == nil {
		config.Db.Save(&book)
		config.Db.Preload("User").Find(&book)
		return book, nil
	}
	return book, errors.New("book do not exists")
}

func DeleteBook(bookID uint64) error {
	var book entity.Book
	config.Db.First(&book, bookID)
	if book.ID != 0 {
		config.Db.Delete(&book)
		return nil
	}
	return errors.New("book do not exists")
}

func GetTheBookUsingID(bookID uint64) entity.Book {
	var book entity.Book
	config.Db.Preload("User").First(&book, bookID)
	return book
}
