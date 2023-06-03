package repository

import (
	"errors"
	"pawAPIbackend/config"
	"pawAPIbackend/entity"
)

func InsertSubmission(submission entity.Submission) entity.Submission {
	config.Db.Save(&submission)
	config.Db.Preload("User").Find(&submission)

	return submission
}

func GetAllSubmissions() []entity.Submission {
	var submissions []entity.Submission
	config.Db.Preload("User").Find(&submissions)

	return submissions
}

func GetAllUserSubmissions(userID uint64) []entity.Submission {
	var submissions []entity.Submission
	config.Db.Preload("User").Where("user_id = ?", userID).Find(&submissions)
	return submissions
}

func GetSubmission(submissionID uint64) (entity.Submission, error) {
	var submission entity.Submission
	config.Db.Preload("User").First(&submission, submissionID)
	if submission.ID != 0 {
		return submission, nil
	}

	return submission, errors.New("submission do not exists")
}

func UpdateSubmission(submission entity.Submission) (entity.Submission, error) {

	if oldSubmission, err := GetSubmission(submission.ID); err == nil {
		submission.Media = oldSubmission.Media
		submission.MediaType = oldSubmission.MediaType
		config.Db.Save(&submission)
		config.Db.Preload("User").Find(&submission)
		return submission, nil
	}
	return submission, errors.New("submission do not exists")
}

func DeleteSubmission(submissionID uint64) error {
	var submission entity.Submission
	config.Db.First(&submission, submissionID)
	if submission.ID != 0 {
		config.Db.Delete(&submission)
		return nil
	}
	return errors.New("submission do not exists")
}

func GetTheSubmissionUsingID(submissionID uint64) entity.Submission {
	var submission entity.Submission
	config.Db.Preload("User").First(&submission, submissionID)
	return submission
}

func InsertImage(image entity.ImageTest) entity.ImageTest {
	config.Db.Save(&image)
	return image
}

func GetImageByID(imageID string) (*entity.ImageTest, error) {
	var image entity.ImageTest
	if err := config.Db.First(&image, imageID).Error; err != nil {
		return nil, err
	}
	return &image, nil
}

func GetAllImages() ([]entity.ImageTest, error) {
	var images []entity.ImageTest
	if err := config.Db.Find(&images).Error; err != nil {
		return nil, err
	}
	return images, nil
}

func GetClinicSubmissions(userID uint64) ([]entity.Submission, error) {
	var submissions []entity.Submission
	err := config.Db.
		Preload("Clinical").
		Preload("Patient").
		Joins("JOIN clinicals ON clinicals.patient_id = submissions.user_id").
		Where("clinicals.clinical_id = ?", userID).
		Find(&submissions).Error
	if err != nil {
		return nil, err
	}

	return submissions, nil
}

func GetEmailFromID(userID uint64) (string, error) {
	var user entity.User
	err := config.Db.First(&user, userID).Error
	if err != nil {
		return "", err
	}

	return user.Email, nil
}
