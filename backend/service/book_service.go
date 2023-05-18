package service

import (
	"errors"
	"log"
	"pawAPIbackend/dto"
	"pawAPIbackend/entity"
	"pawAPIbackend/repository"

	"github.com/mashingan/smapping"
)

func GetAllBooks() []dto.BookResponseDTO {
	var booksResponse []dto.BookResponseDTO
	books := repository.GetAllBooks()

	for _, user := range books {
		response := dto.BookResponseDTO{}
		err := smapping.FillStruct(&response, smapping.MapFields(&user))
		if err != nil {
			log.Fatal("failed to map to response ", err)
			return booksResponse
		}
		booksResponse = append(booksResponse, response)
	}

	return booksResponse
}

func InsertBook(bookDTO dto.BookCreatedDTO, userID uint64) dto.BookResponseDTO {
	book := entity.Book{}
	bookResponse := dto.BookResponseDTO{}

	//mapear entidades diferentes, neste caso BookCreatedDTO to Book
	//codpia diretamente de uma estrutura par adetnroe de outra
	///preenche com todos os campos que são semelhantes
	err := smapping.FillStruct(&book, smapping.MapFields(&bookDTO))
	if err != nil {
		log.Fatal("failed to map ", err)
		return bookResponse
	}

	book.UserID = userID
	book = repository.InsertBook(book)

	err = smapping.FillStruct(&bookResponse, smapping.MapFields(&book))
	if err != nil {
		log.Fatal("failed to map to response ", err)
		return bookResponse
	}

	return bookResponse
}

func GetBook(bookID uint64) (dto.BookResponseDTO, error) {
	bookResponse := dto.BookResponseDTO{}

	if book, err := repository.GetBook(bookID); err == nil {

		err = smapping.FillStruct(&bookResponse, smapping.MapFields(&book))

		if err != nil {
			log.Fatal("failed to map to response ", err)
			return bookResponse, err
		}

		return bookResponse, nil
	}
	return bookResponse, errors.New("book do not exist")
}

func UpdateBook(bookDTO dto.BookUpdateDTO) (dto.BookResponseDTO, error) {
	book := entity.Book{}
	bookResponse := dto.BookResponseDTO{}

	err := smapping.FillStruct(&book, smapping.MapFields(&bookDTO))
	if err != nil {
		log.Fatal("failed to map ", err)
		return bookResponse, nil
	}

	if book, err = repository.UpdateBook(book); err == nil {

		err = smapping.FillStruct(&bookResponse, smapping.MapFields(&book))

		if err != nil {
			log.Fatal("failed to map to response ", err)
			return bookResponse, err
		}

		return bookResponse, nil
	}

	return bookResponse, errors.New("book do not exists")
}

func DeleteBook(bookID uint64) error {
	if err := repository.DeleteBook(bookID); err == nil {
		return nil
	}
	return errors.New("book do not exists")
}

func IsAllowedToEdit(userID uint64, bookID uint64) bool {
	b := repository.GetTheBookUsingID(bookID)
	return userID == b.UserID
}
