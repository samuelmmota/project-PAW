package entity

type ImageTest struct {
	ID        uint64 `gorm:"primary_key:auto_increment" json:"id"`
	ImageTest []byte `gorm:"type:longblob" json:"image"`
	MediaType string `gorm:"type:varchar(20)" json:"mediaType"`
}
