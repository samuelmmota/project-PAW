package entity

type Message struct {
	ID             uint64     `gorm:"primary_key:auto_increment" json:"id"`
	Date           string     `gorm:"type:varchar(255)" json:"date"`
	MessageContent string     `gorm:"type:text" json:"message_content"`
	SubmissionID   uint64     `gorm:"not null" json:"-"`
	Submission     Submission `gorm:"foreignkey:SubmissionID;constraint;onUpdate:CASCADE;onDelete:CASCADE" json:"submission"`
	UserID         uint64     `gorm:"not null" json:"-"`
	User           User       `gorm:"foreignkey:UserID;constraint;onUpdate:CASCADE;onDelete:CASCADE" json:"user"`
}
