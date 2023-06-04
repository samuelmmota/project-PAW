package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io/ioutil"
	"log"
	"mime/multipart"
	"pawAPIbackend/dto"
	"pawAPIbackend/entity"
	"pawAPIbackend/repository"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/scrypt"

	"github.com/mashingan/smapping"
)

const (
	saltSize            = 16
	encryptionNonceSize = 12
)

func GetAllSubmissions(userID uint64) []dto.SubmissionResponseDTO {
	var submissionResponse []dto.SubmissionResponseDTO
	submissions := repository.GetAllUserSubmissions(userID)

	for _, user := range submissions {
		response := dto.SubmissionResponseDTO{}
		err := smapping.FillStruct(&response, smapping.MapFields(&user))
		if err != nil {
			log.Fatal("failed to map submission to response ", err)
			return submissionResponse
		}

		imageDecrypted, err := GetImage(response.Media, userID)
		if err != nil {
			log.Fatal("Failed to decrypt image", err)
			return submissionResponse
		}

		response.Media = imageDecrypted

		submissionResponse = append(submissionResponse, response)
	}

	return submissionResponse
}

func InsertSubmission(submissionCreateDTO dto.SubmissionCreateDTO, multipartFile *multipart.FileHeader, userID uint64) (dto.SubmissionResponseDTO, error) {
	submission := entity.Submission{}
	submissionResponse := dto.SubmissionResponseDTO{}

	err := smapping.FillStruct(&submission, smapping.MapFields(&submissionCreateDTO))
	if err != nil {
		log.Fatal("failed to map ", err)
		return submissionResponse, err
	}

	err = isValidFormat(multipartFile.Filename)
	if err != nil {
		log.Fatal("Wrong format type inserted!", err)
		return submissionResponse, err
	}

	newDate, err := ConverterDate(submissionCreateDTO.Date)
	if err != nil {
		log.Fatal("Wrong format type inserted!", err)
		return submissionResponse, err
	}

	//TEST code
	// Open the uploaded file

	file, err := multipartFile.Open()
	if err != nil {
		log.Fatal("Failed to open image file ", err)
		return submissionResponse, err
	}
	defer file.Close()

	// Read the file content
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("Failed to read image file", err)
		return submissionResponse, err
	}

	imageEncrypted, err := InsertImage(fileBytes, userID)
	if err != nil {
		log.Fatal("Failed to encrypt image", err)
		return submissionResponse, err
	}

	submission.UserID = userID

	submission.Media = imageEncrypted

	submission.Date = newDate

	submission = repository.InsertSubmission(submission)

	err = smapping.FillStruct(&submissionResponse, smapping.MapFields(&submission))
	if err != nil {
		log.Fatal("failed to map to response ", err)
		return submissionResponse, err
	}

	return submissionResponse, err
}

func GetSubmission(submissionID uint64, userId uint64) (dto.SubmissionResponseDTO, error) {
	submissionResponse := dto.SubmissionResponseDTO{}

	//TODO: Validation user before get submission

	if submission, err := repository.GetSubmission(submissionID); err == nil {

		err = smapping.FillStruct(&submissionResponse, smapping.MapFields(&submission))

		if err != nil {
			log.Fatal("failed to map to response ", err)
			return submissionResponse, err
		}

		return submissionResponse, nil
	}

	imageDecrypted, err := GetImage(submissionResponse.Media, userId)
	if err != nil {
		log.Fatal("Failed to decrypt image", err)
		return submissionResponse, err
	}

	submissionResponse.Media = imageDecrypted

	return submissionResponse, errors.New("submission do not exist")
}

func UpdateSubmission(submissionDTO dto.SubmissionUpdateDTO) (dto.SubmissionResponseDTO, error) {
	submission, err := repository.GetSubmission(submissionDTO.ID)
	if err != nil {
		return dto.SubmissionResponseDTO{}, errors.New("Submission does not exist")
	}

	if submissionDTO.Description != "" {
		submission.Description = submissionDTO.Description
	}

	if submissionDTO.BodyPart != "" {
		submission.BodyPart = submissionDTO.BodyPart
	}

	if submissionDTO.Date != "" {
		submission.Date = submissionDTO.Date
	}

	if submissionDTO.ShareWithClinicals != "" {
		isSharedWithClinicals, err := strconv.ParseBool(submissionDTO.ShareWithClinicals)
		if err != nil {
			return dto.SubmissionResponseDTO{}, errors.New("Invalid value for IsClinical")
		}
		submission.ShareWithClinicals = isSharedWithClinicals
	}

	submissionResponse := dto.SubmissionResponseDTO{}

	if submissionUpdated, err := repository.UpdateSubmission(submission); err == nil {

		err = smapping.FillStruct(&submissionResponse, smapping.MapFields(&submissionUpdated))

		if err != nil {
			log.Fatal("failed to map to response ", err)
			return submissionResponse, err
		}

		return submissionResponse, nil
	}

	return submissionResponse, errors.New("submission do not exists")
}

func DeleteSubmission(submissionID uint64) error {
	if err := repository.DeleteSubmission(submissionID); err == nil {
		return nil
	}
	return errors.New("submission do not exists")
}

func IsAllowedToEdit(userID uint64, submissionID uint64) bool {
	submission := repository.GetTheSubmissionUsingID(submissionID)
	return userID == submission.UserID
}

func encryptImage(image []byte, userKey string) ([]byte, error) {

	// Derive a key from the user's password
	salt := make([]byte, saltSize)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	key, err := scrypt.Key([]byte(userKey), salt, 16384, 8, 1, 32)
	if err != nil {
		return nil, err
	}

	// Generate a random nonce
	nonce := make([]byte, encryptionNonceSize)
	_, err = rand.Read(nonce)
	if err != nil {
		return nil, err
	}

	// Create a new AES-GCM cipher block using the derived key
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Create a new GCM mode cipher with the block and nonce
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Encrypt the plaintext using GCM and the nonce
	encryptedData := aesGCM.Seal(nil, nonce, image, nil)

	// Prepend the salt and nonce to the encrypted data
	encryptedData = append(nonce, encryptedData...)
	encryptedData = append(salt, encryptedData...)

	// Return the encrypted data
	return encryptedData, nil
}
func decryptImage(encryptedData []byte, userKey string) ([]byte, error) {
	if len(encryptedData) < saltSize+encryptionNonceSize {
		return nil, errors.New("invalid encrypted data")
	}

	salt := encryptedData[:saltSize]
	nonce := encryptedData[saltSize : saltSize+encryptionNonceSize]
	cipherText := encryptedData[saltSize+encryptionNonceSize:]

	// Derive the key from the user's key and the stored salt
	key, err := scrypt.Key([]byte(userKey), salt, 16384, 8, 1, 32)
	if err != nil {
		return nil, err
	}

	// Create a new AES-GCM cipher block using the derived key
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Create a new GCM mode cipher with the block and nonce
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Decrypt the cipherText using GCM and the nonce
	decryptedData, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}

	// Return the decrypted data
	return decryptedData, nil
}


func InsertImage(image []byte, userId uint64) ([]byte, error) {

	user, err := repository.GetUser(userId)
	if err != nil {
		log.Fatal("failed to get user ", err)
		return nil, err
	}
	
	key := user.Key

	//Encrypt the image data using the user's key
	encryptedImage, err := encryptImage(image, key)
	if err != nil {
		log.Fatal("failed to encrypt image: ", err)
		return nil, err
	}

	imageResult := encryptedImage

	return imageResult, nil
}

func GetImage(image []byte, userId uint64) ([]byte, error) {

	user, err := repository.GetUser(userId)
	if err != nil {
		log.Fatal("failed to get user ", err)
		return nil, err
	}

	key := user.Key

	//Decrypt the image data using the user's key
	decryptedImage, err := decryptImage(image, key)
	if err != nil {
		log.Fatal("failed to decrypt image: ", err)
		return nil, err
	}

	return decryptedImage, nil
}

func isValidFormat(mediaType string) error {
	imageFormats := []string{"jpg", "jpeg", "png", "gif"}
	videoFormats := []string{"mp4", "avi", "mov"}

	fileExt := getFileExtension(mediaType)

	for _, format := range imageFormats {
		if format == fileExt {
			return nil
		}
	}

	for _, format := range videoFormats {
		if format == fileExt {
			return nil
		}
	}

	return errors.New("Not Valid Format!")
}

func getFileExtension(mediaType string) string {
	parts := strings.Split(mediaType, ".")
	if len(parts) > 1 {
		return strings.ToLower(parts[len(parts)-1])
	}
	return ""
}

func ConverterDate(date string) (string, error) {

	newInputDate:= removeSubstring(date)
	// Define o layout de entrada
	layoutEntrada := "Mon Jan 2 2006 15:04:05"
	// Converte a string para um objeto time.Time
	t, err := time.Parse(layoutEntrada, newInputDate)
	if err != nil {
		return "", err
	}

	// Define o layout de saída
	layoutSaida := "2006-01-02T15:04:05.000Z"
	// Formata o objeto time.Time no formato desejado
	newDate := t.UTC().Format(layoutSaida)

	return newDate, nil
}

func removeSubstring(str string) string {
	index := strings.Index(str, " GMT")
	if index != -1 {
		newString := str[:index]
		return newString
	}
	return str
}

