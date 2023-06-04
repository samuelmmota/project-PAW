package repository

import (
	"log"
	"pawAPIbackend/config"
	"pawAPIbackend/entity"
)

func GetClinicals(userPatientID uint64) []entity.Clinical {
	log.Default().Println(userPatientID)
	var clinicals []entity.Clinical
	config.Db.Preload("Clinical").Preload("Patient").Where("patient_id = ?", userPatientID).Find(&clinicals)

	return clinicals
}

func GetPatients(userClinicalID uint64) []entity.Clinical {
	log.Default().Println(userClinicalID)
	var clinicals []entity.Clinical
	config.Db.Preload("Clinical").Preload("Patient").Where("clinical_id = ?", userClinicalID).Find(&clinicals)

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

func GetCLinical(userClinicalID uint64, userPatientID uint64) (entity.Clinical, error) {
	var clinical entity.Clinical
	//config.Db.Preload("Clinical").Preload("Patient").Where("clinical_id = ?", userClinicalID).Find(&clinicals)
	err := config.Db.Where("patient_id = ? AND clinical_id = ?", userPatientID, userClinicalID).First(&clinical).Error
	return clinical, err
}
