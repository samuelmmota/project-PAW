package entity

type Clinical struct {
	ID         uint64 `gorm:"primary_key:auto_increment" json:"id"`
	ClinicalID uint64 `gorm:"not null" json:"-"`
	Clinical   User   `gorm:"foreignkey:ClinicalID;constraint;onUpdate:CASCADE;onDelete:CASCADE" json:"clinical"`
	PatientID  uint64 `gorm:"not null" json:"-"`
	Patient    User   `gorm:"foreignkey:PatientID;constraint;onUpdate:CASCADE;onDelete:CASCADE" json:"patient"`
}
