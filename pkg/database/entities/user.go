package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserId         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Username       string
	Email          string
	Fullname       string
	PasswordHash   string
	SignInProvider string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Active         bool
}
