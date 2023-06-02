package repository

import (
	"log"
	"pawAPIbackend/config"
	"pawAPIbackend/entity"
)

func GetAllUserClinicals(userID uint64) []entity.Clinical {
	log.Default().Println(userID)
	var clinicals []entity.Clinical
	config.Db.Preload("Clinical").Preload("Patient").Where("patient_id = ?", userID).Find(&clinicals)

	return clinicals
}

func DoesClinicalExist(clinical entity.Clinical) bool {
	config.Db.Where("patient_id = ? AND clinical_id = ?", clinical.PatientID, clinical.ClinicalID).First(&clinical)
	if clinical.ID != 0 {
		return true
	}

	return false
}

func CreateClinical(clinical entity.Clinical) error {
	err := config.Db.Save(&clinical).Error
	if err != nil {
		return err
	}

	return nil
}
