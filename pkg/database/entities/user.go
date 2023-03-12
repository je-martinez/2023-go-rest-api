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
	UpdatedAt      time.Time `gorm:"default:null"`
	Active         bool      `gorm:"default:true"`
	Profile        Profile
	Posts          []Post
	Comments       []Comment
}

type UserInput struct {
	UserID         string
	Username       string
	Email          string
	Fullname       string
	PasswordHash   string
	SignInProvider string
	Active         bool
}

func (i User) ToEntity() (input UserInput) {
	return UserInput{
		UserID:         i.UserID,
		Username:       i.Username,
		Email:          i.Email,
		Fullname:       i.Fullname,
		PasswordHash:   i.PasswordHash,
		SignInProvider: i.SignInProvider,
		Active:         i.Active,
	}
}

func (i User) FromEntity(entity UserInput) interface{} {
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
