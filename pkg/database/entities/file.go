package entities

import (
	"time"

	"github.com/google/uuid"
)

type File struct {
	FileId    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Post      Post
	CreatedBy User
	CreatedAt time.Time
	UpdatedBy User
	UpdatedAt time.Time
	Active    bool
}
