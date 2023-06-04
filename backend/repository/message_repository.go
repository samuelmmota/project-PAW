package repository

import (
	"pawAPIbackend/config"
	"pawAPIbackend/entity"
)

func GetMessages(submissionID uint64) []entity.Message {
	var messages []entity.Message
	config.Db.Preload("Message").Preload("Submission").Where("submission_id = ?", submissionID).Find(&messages)

	return messages
}

func AddMessage(message entity.Message) error {
	err := config.Db.Save(&message).Error
	if err != nil {
		return err
	}

	return nil
}
