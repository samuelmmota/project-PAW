package entity

type Submission struct {
	ID                 uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Date               string `gorm:"type:varchar(255)" json:"date"`
	Description        string `gorm:"type:text" json:"description"`
	BodyPart           string `gorm:"type:varchar(255)" json:"body_part"`
	Media              []byte `gorm:"type:longblob" json:"media"`
	MediaType          string `gorm:"type:varchar(20)" json:"media_type"`
	UserID             uint64 `gorm:"not null" json:"-"`
	ShareWithClinicals bool   `gorm:"default:false" json:"shareWithClinicals"`
	User               User   `gorm:"foreignkey:UserID;constraint;onUpdate:CASCADE;onDelete:CASCADE" json:"user"`
}
