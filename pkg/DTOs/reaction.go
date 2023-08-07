package DTOs

import (
	"github.com/je-martinez/2023-go-rest-api/pkg/database/entities"
	db_extensions "github.com/je-martinez/2023-go-rest-api/pkg/database/extensions"
)

type CreatePostReactionDTO struct {
	ReactionType db_extensions.ReactionType
	PostID       string
	UserID       string
}

func (r *CreatePostReactionDTO) ToEntity() *entities.Reaction {
	return &entities.Reaction{
		PostID:       &r.PostID,
		ReactionType: r.ReactionType,
		UserID:       r.UserID,
		CreatedBy:    r.UserID,
	}
}
