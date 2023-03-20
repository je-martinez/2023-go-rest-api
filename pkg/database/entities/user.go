package entities

import (
	dbe "main/pkg/database/extensions"
	"time"
)

type User struct {
	UserID         string `gorm:"type:uuid;primary_key;default:uuid_generate_v4();"`
	Username       string `gorm:"unique"`
	Email          string `gorm:"unique"`
	Fullname       string
	PasswordHash   string
	SignInProvider dbe.SignInProviderType `gorm:"type:sign_in_provider_type"`
	CreatedAt      time.Time              `gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedAt      time.Time              `gorm:"default:null"`
	Active         bool                   `gorm:"default:true"`
	Profile        Profile
	Posts          []Post
	Comments       []Comment
}

type UserModel struct {
	UserID         string                 `json:"user_id"`
	Username       string                 `json:"username"`
	Email          string                 `json:"email"`
	Fullname       string                 `json:"fullname"`
	PasswordHash   string                 `json:"password"`
	SignInProvider dbe.SignInProviderType `json:"sign_in_provider"`
	Active         bool                   `json:"active"`
}

func (i User) ToModel() (input UserModel) {
	return UserModel{
		UserID:         i.UserID,
		Username:       i.Username,
		Email:          i.Email,
		Fullname:       i.Fullname,
		PasswordHash:   i.PasswordHash,
		SignInProvider: i.SignInProvider,
		Active:         i.Active,
	}
}

func (i User) FromModel(entity UserModel) interface{} {
	return User{
		UserID:         entity.UserID,
		Username:       entity.Username,
		Email:          entity.Email,
		Fullname:       entity.Fullname,
		PasswordHash:   entity.PasswordHash,
		SignInProvider: entity.SignInProvider,
		Active:         entity.Active,
	}
}
