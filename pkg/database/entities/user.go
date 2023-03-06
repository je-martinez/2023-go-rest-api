package entities

import "time"

type User struct {
	UserID         string `gorm:"type:uuid;primary_key;default:uuid_generate_v4();"`
	Username       string `gorm:"unique"`
	Email          string `gorm:"unique"`
	Fullname       string
	PasswordHash   string
	SignInProvider string
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	Active         bool      `gorm:"default:true"`
	Profile        Profile
	Posts          []Post
	Comments       []Comment
}
