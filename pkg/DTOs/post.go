package DTOs

import (
	"mime/multipart"

	"github.com/je-martinez/2023-go-rest-api/pkg/database/entities"
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
