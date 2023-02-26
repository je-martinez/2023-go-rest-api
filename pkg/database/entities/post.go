package entities

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	PostId        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Content       string
	Comment_Count int64
	LiKe_Count    int64
	CreatedBy     User
	CreatedAt     time.Time
	UpdatedBy     User
	UpdatedAt     time.Time
	Active        bool
}
