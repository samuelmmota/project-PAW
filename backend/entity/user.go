package entity

type User struct {
	ID                  uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Password            string `gorm:"not null" json:"password"`
	Email               string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	ExportToReasearcher bool   `gorm:"default:false" json:"exportToReasearcher"`
	IsClinical          bool   `gorm:"default:false" json:"isClinical"`
	Token               string `gorm:"-" json:"token"`
}
