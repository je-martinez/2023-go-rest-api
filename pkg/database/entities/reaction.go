package entities

import (
	"time"

	dbe "github.com/je-martinez/2023-go-rest-api/pkg/database/extensions"
)

type Reaction struct {
	ReactionID   string           `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ReactionType dbe.ReactionType `gorm:"type:reaction_type"`
	PostID       *string
	CommentID    *string
	UserID       string
	CreatedBy    string
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedBy    *string
	UpdatedAt    time.Time `gorm:"default:null"`
	Active       bool      `gorm:"default:true"`
}

type ReactionDTO struct {
	ReactionID   string           `json:"reaction_id"`
	ReactionType dbe.ReactionType `json:"reaction_type"`
	PostID       *string          `json:"post_id"`
	CommentID    *string          `json:"comment_id"`
	UserID       string           `json:"user_id"`
}

func (i Reaction) ToDTO() *ReactionDTO {
	return &ReactionDTO{
		ReactionID:   i.ReactionID,
		ReactionType: i.ReactionType,
		PostID:       i.PostID,
		CommentID:    i.CommentID,
		UserID:       i.UserID,
	}
}
