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
	UpdatedAt time.Time `gorm:"default:null"`
	Active    bool      `gorm:"default:true"`
	File      *File
}

type CommentDTO struct {
	CommentID string
	PostID    string
	UserID    string
	Content   string
	File      *FileDTO
}

func (i Comment) ToDTO() *CommentDTO {
	var file *FileDTO
	if i.File != nil {
		file = i.File.ToDTO()
	}
	return &CommentDTO{
		CommentID: i.CommentID,
		PostID:    i.PostID,
		UserID:    i.UserID,
		Content:   i.Content,
		File:      file,
	}
}
