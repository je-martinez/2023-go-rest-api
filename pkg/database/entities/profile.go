package entities

import "time"

type Profile struct {
	ProfileID      string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID         string
	PhotosUploaded int64
	VideosUploaded int64
	AudiosUploaded int64
	LikesCount     int64
	LastPost       int64
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	Active         bool      `gorm:"default:true"`
}
