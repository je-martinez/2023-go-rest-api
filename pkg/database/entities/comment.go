package entities

import (
	"time"
)

type Comment struct {
	CommentID string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	PostID    string
	UserID    string
	Content   string
	CreatedBy string
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedBy string
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	Active    bool      `gorm:"default:true"`
	File      File
}
