package entity

type User struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Password string `gorm:"not null" json:"password"`
	Email    string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Token    string `gorm:"-" json:"token"`
	Key      string `gorm:"type:varchar(20)" json:"key"`
}
