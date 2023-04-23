package DTOs

import (
	"main/pkg/database/entities"
	"mime/multipart"
)

type CreatePostDTO struct {
	Content string                 `form:"content"`
	Files   []multipart.FileHeader `form:"files" binding:"omitempty"`
}

func (p CreatePostDTO) ToEntity(content string, userID string) *entities.Post {
	return &entities.Post{
		UserID:    userID,
		Content:   content,
		CreatedBy: userID,
	}
}
