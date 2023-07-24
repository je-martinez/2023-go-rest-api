package entities

import (
	"mime/multipart"
	"time"
)

type File struct {
	FileID    string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	PostId    *string
	CommentId *string
	Name      string
	Key       string
	MimeType  string
	CreatedBy string
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedBy string
	UpdatedAt time.Time `gorm:"default:null"`
	Active    bool      `gorm:"default:true"`
}

type FileDTO struct {
	FileID    string  `json:"file_id"`
	PostId    *string `json:"post_id"`
	CommentId *string `json:"comment_id"`
	Name      string  `json:"name"`
	Key       string  `json:"key"`
	MimeType  string  `json:"mimetype"`
}

func (i File) ToDTO() *FileDTO {
	return &FileDTO{
		FileID:    i.FileID,
		PostId:    i.PostId,
		CommentId: i.CommentId,
		Name:      i.Name,
		Key:       i.Key,
		MimeType:  i.MimeType,
	}
}

func (i File) FromMinioUpload(key string, path string, user_id string, post_id string, file multipart.FileHeader) File {
	return File{
		PostId:    &post_id,
		Name:      file.Filename,
		Key:       key,
		CreatedBy: user_id,
		MimeType:  file.Header.Get("Content-Type"),
	}
}
