package DTOs

import (
	"main/pkg/database/entities"
	t "main/pkg/database/extensions"
)

type LoginDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthResponseDTO struct {
	Username string               `json:"username"`
	Fullname string               `json:"fullname"`
	Email    string               `json:"email"`
	Provider t.SignInProviderType `json:"provider"`
	Token    string               `json:"token"`
}

type RegisterUserDTO struct {
	Username string               `json:"username" validate:"required"`
	Fullname string               `json:"fullname" validate:"required"`
	Email    string               `json:"email" validate:"email,required"`
	Provider t.SignInProviderType `json:"provider" validate:"required"`
	Password string               `json:"password" validate:"required"`
}

func (r RegisterUserDTO) ToEntity(passwordHash string) *entities.User {
	return &entities.User{
		Username:       r.Username,
		Fullname:       r.Fullname,
		Email:          r.Email,
		SignInProvider: r.Provider,
		PasswordHash:   passwordHash,
		Profile: entities.Profile{
			PhotosUploaded: 0,
			VideosUploaded: 0,
			AudiosUploaded: 0,
		},
	}
}

type UpdateUserDTO struct {
	Fullname    string `json:"fullname" validate:"required"`
	Email       string `json:"email" validate:"email,required"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
