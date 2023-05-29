package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"golang.org/x/crypto/scrypt"
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

func InsertSubmission(submissionCreateDTO dto.SubmissionCreateDTO, multipartFile *multipart.FileHeader, userID uint64, image entity.ImageTest) (dto.SubmissionResponseDTO, error) {
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

	imageEncrypted, err := InsertImage(image, userID)
	if err != nil {
		log.Fatal("Failed to encrypt image", err)
		return submissionResponse, err
	}

	submission.UserID = userID

	submission.Media = imageEncrypted.ImageTest
	submission.MediaType = imageEncrypted.MediaType

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
	encryptedData = append(salt, encryptedData...)
	encryptedData = append(nonce, encryptedData...)

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

	// Derive the key from the user's password and the stored salt
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

func InsertImage(image entity.ImageTest, userId uint64) (entity.ImageTest, error) {

	user, err := repository.GetUser(userId)
	if err != nil {
		log.Fatal("failed to get user ", err)
		return entity.ImageTest{}, err
	}

	key := user.Key

	//Encrypt the image data using the user's key
	encryptedImage, err := encryptImage(image.ImageTest, key)
	if err != nil {
		log.Fatal("failed to encrypt image: ", err)
		return entity.ImageTest{}, err
	}

	image.ImageTest = encryptedImage

	return image, nil
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
