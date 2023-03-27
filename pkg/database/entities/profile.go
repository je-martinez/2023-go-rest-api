package entities

import (
	"database/sql"
	"time"
)

type Profile struct {
	ProfileID      string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID         string
	PhotosUploaded int64
	VideosUploaded int64
	AudiosUploaded int64
	LastPostID     sql.NullString
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedAt      time.Time `gorm:"default:null"`
	Active         bool      `gorm:"default:true"`
	LastPost       Post      `gorm:"foreignKey:LastPostID"`
}

type ProfileDTO struct {
	ProfileID      string `json:"profile_id"`
	UserID         string `json:"user_id"`
	PhotosUploaded int64  `json:"photos_uploaded"`
	VideosUploaded int64  `json:"videos_uploaded"`
	AudiosUploaded int64  `json:"audios_uploaded"`
}

func (i Profile) ToDTO() (input ProfileDTO) {
	return ProfileDTO{
		ProfileID:      i.ProfileID,
		UserID:         i.UserID,
		PhotosUploaded: i.PhotosUploaded,
		VideosUploaded: i.VideosUploaded,
		AudiosUploaded: i.AudiosUploaded,
	}
}

func (i Profile) FromDTO(model ProfileDTO) Profile {
	return Profile{
		ProfileID:      i.ProfileID,
		UserID:         i.UserID,
		PhotosUploaded: i.PhotosUploaded,
		VideosUploaded: i.VideosUploaded,
		AudiosUploaded: i.AudiosUploaded,
	}
}
