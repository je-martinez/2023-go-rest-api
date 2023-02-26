package entities

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	ProfileId      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	User           User
	PhotosUploaded int64
	VideosUploaded int64
	AudiosUploaded int64
	LikesCount     int64
	LastPost       int64
	UpdatedBy      User
	UpdatedAt      time.Time
	Active         bool
}
