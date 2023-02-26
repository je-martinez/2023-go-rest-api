package entities

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	CommentId uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Post      Post
	File      File
	CreatedBy User
	CreatedAt time.Time
	UpdatedBy User
	UpdatedAt time.Time
	Active    bool
}
