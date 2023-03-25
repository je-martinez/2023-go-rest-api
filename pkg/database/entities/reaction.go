package entities

import (
	dbe "main/pkg/database/extensions"
	"time"
)

type Reaction struct {
	ReactionID   string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ReactionType dbe.ReactionType
	PostId       string
	CommentId    string
	UserId       string
	CreatedBy    string
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedBy    string
	UpdatedAt    time.Time `gorm:"default:null"`
	Active       bool      `gorm:"default:true"`
}
