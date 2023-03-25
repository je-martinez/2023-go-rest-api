package entities

import (
	"time"
)

type File struct {
	FileID    string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	PostId    string
	CommentId string
	Name      string
	Path      string
	MimeType  string
	CreatedBy string
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedBy string
	UpdatedAt time.Time `gorm:"default:null"`
	Active    bool      `gorm:"default:true"`
}
