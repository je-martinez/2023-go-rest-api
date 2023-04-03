package entities

import (
	dbe "main/pkg/database/extensions"
	"time"
)

type Reaction struct {
	ReactionID   string           `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ReactionType dbe.ReactionType `gorm:"type:reaction_type"`
	PostId       string
	CommentId    string
	UserId       string
	CreatedBy    string
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedBy    string
	UpdatedAt    time.Time `gorm:"default:null"`
	Active       bool      `gorm:"default:true"`
}

type ReactionDTO struct {
	ReactionID   string
	ReactionType dbe.ReactionType
	PostId       string
	CommentId    string
	UserId       string
}

func (i Reaction) ToDTO() *ReactionDTO {
	return &ReactionDTO{
		ReactionID:   i.ReactionID,
		ReactionType: i.ReactionType,
		PostId:       i.PostId,
		CommentId:    i.CommentId,
		UserId:       i.UserId,
	}
}
