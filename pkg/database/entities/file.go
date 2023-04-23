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

type FileDTO struct {
	FileID    string `json:"file_id"`
	PostId    string `json:"post_id"`
	CommentId string `json:"comment_id"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	MimeType  string `json:"mimetype"`
}

func (i File) ToDTO() *FileDTO {
	return &FileDTO{
		FileID:    i.FileID,
		PostId:    i.PostId,
		CommentId: i.CommentId,
		Name:      i.Name,
		MimeType:  i.MimeType,
	}
}
