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

func GetSubmission(submissionID uint64) (entity.Submission, error) {
	var submission entity.Submission
	config.Db.Preload("User").First(&submission, submissionID)
	if submission.ID != 0 {
		return submission, nil
	}

	return submission, errors.New("submission do not exists")
}

func UpdateSubmission(submission entity.Submission) (entity.Submission, error) {

	if _, err := GetSubmission(submission.ID); err == nil {
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
