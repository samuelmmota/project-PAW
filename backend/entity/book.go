package entity

type Book struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"title" binding:"required"`
	Year        string `gorm:"type:varchar(4)" json:"year"`
	Description string `gorm:"type:text" json:"description"`
	BookCover   string `gorm:"type:varchar(1024); default:'https://dl.acm.org/specs/products/acm/releasedAssets/images/cover-default--book.svg'" json:"book_cover"`
	UserID      uint64 `gorm:"not null" json:"-"`
	User        User   `gorm:"foreignkey:UserID;constraint;onUpdate:CASCADE;onDelete:CASCADE" json:"user"`
}
