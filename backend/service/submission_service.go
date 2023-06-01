package service

import (
	"errors"
	"io/ioutil"
	"log"
	"mime/multipart"
	"pawAPIbackend/dto"
	"pawAPIbackend/entity"
	"pawAPIbackend/repository"

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

	submission.Media = fileBytes

	submission.UserID = userID
	submission = repository.InsertSubmission(submission)

	err = smapping.FillStruct(&submissionResponse, smapping.MapFields(&submission))
	if err != nil {
		log.Fatal("failed to map to response ", err)
		return submissionResponse, err
	}

	return submissionResponse, err
}

func GetSubmission(submissionID uint64) (dto.SubmissionResponseDTO, error) {
	submissionResponse := dto.SubmissionResponseDTO{}

	if submission, err := repository.GetSubmission(submissionID); err == nil {

		err = smapping.FillStruct(&submissionResponse, smapping.MapFields(&submission))

		if err != nil {
			log.Fatal("failed to map to response ", err)
			return submissionResponse, err
		}

		return submissionResponse, nil
	}
	return submissionResponse, errors.New("submission do not exist")
}

func UpdateSubmission(submissionDTO dto.SubmissionUpdateDTO) (dto.SubmissionResponseDTO, error) {
	submission := entity.Submission{}
	submissionResponse := dto.SubmissionResponseDTO{}

	err := smapping.FillStruct(&submission, smapping.MapFields(&submissionDTO))
	if err != nil {
		log.Fatal("failed to map ", err)
		return submissionResponse, nil
	}

	if submission, err = repository.UpdateSubmission(submission); err == nil {

		err = smapping.FillStruct(&submissionResponse, smapping.MapFields(&submission))

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

/*
func encryptImage(dto *dto.SubmissionCreateDTO, userPassword string) ([]byte, error) {
	plainText := dto.Image

	// Derive a key from the user's password
	salt := make([]byte, saltSize)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	key, err := scrypt.Key([]byte(userPassword), salt, 16384, 8, 1, 32)
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
	encryptedData := aesGCM.Seal(nil, nonce, plainText, nil)

	// Prepend the salt and nonce to the encrypted data
	encryptedData = append(salt, encryptedData...)
	encryptedData = append(nonce, encryptedData...)

	// Return the encrypted data
	return encryptedData, nil
}
func decryptImage(encryptedData []byte, userPassword string) ([]byte, error) {
	if len(encryptedData) < saltSize+encryptionNonceSize {
		return nil, errors.New("invalid encrypted data")
	}

	salt := encryptedData[:saltSize]
	nonce := encryptedData[saltSize : saltSize+encryptionNonceSize]
	cipherText := encryptedData[saltSize+encryptionNonceSize:]

	// Derive the key from the user's password and the stored salt
	key, err := scrypt.Key([]byte(userPassword), salt, 16384, 8, 1, 32)
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
}*/

func InsertImage(image entity.ImageTest) entity.ImageTest {
	// Get the user's password-> podemos vir buscar diretamente a hashed password porque como estamos no service
	//já tems a certeza que o user é o owner da submission!
	//userPassword, err := repository.GetUser(userID)

	// Encrypt the image data using the user's password
	/*encryptedImage, err := encryptImage(&submissionCreateDTO, userPassword.Password)
	if err != nil {
		log.Fatal("failed to encrypt image: ", err)
		return submissionResponse
	}

	submission.Image = encryptedImage*/

	imageReceived := repository.InsertImage(image)

	/*submissionResponse.Image, err = decryptImage(submissionResponse.Image, userPassword.Password)
	if err != nil {
		log.Fatal("failed to decrypt image data and map to response ", err)
		return submissionResponse
	}*/

	return imageReceived

}
