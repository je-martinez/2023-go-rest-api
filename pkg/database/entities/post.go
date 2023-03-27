package entities

import (
	"time"
)

type Post struct {
	PostID        string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID        string
	Content       string
	Comment_Count int64
	Like_Count    int64
	CreatedBy     string
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedBy     string
	UpdatedAt     time.Time `gorm:"default:null"`
	Active        bool      `gorm:"default:true"`
	Comments      *[]Comment
	Files         *[]File
}
