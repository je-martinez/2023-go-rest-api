package DTOs

import (
	"main/pkg/database/entities"
	t "main/pkg/database/extensions"
)

type LoginDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterUserDTO struct {
	Username string               `json:"username" validate:"required"`
	Fullname string               `json:"fullname" validate:"required"`
	Email    string               `json:"email" validate:"email,required"`
	Provider t.SignInProviderType `json:"provider" validate:"required"`
	Password string               `json:"password" validate:"required"`
}

func (data RegisterUserDTO) ToEntity(passwordHash string) *entities.User {
	return &entities.User{
		Username:       data.Username,
		Fullname:       data.Fullname,
		Email:          data.Email,
		SignInProvider: data.Provider,
		PasswordHash:   passwordHash,
	}
}

type UpdateUserDTO struct {
	UserId   string `json:"user_id" validate:"required"`
	Fullname string `json:"fullname" validate:"required"`
	Email    string `json:"email" validate:"email,required"`
}
