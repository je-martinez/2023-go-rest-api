package DTOs

import (
	"main/pkg/database/entities"
	t "main/pkg/database/extensions"
)

type RegisterUserDTO struct {
	Username string               `json:"username"`
	Fullname string               `json:"fullname"`
	Email    string               `json:"email"`
	Provider t.SignInProviderType `json:"provider"`
	Password string               `json:"password"`
}

func (data RegisterUserDTO) ToModel(passwordHash string) *entities.UserModel {
	return &entities.UserModel{
		Username:       data.Username,
		Fullname:       data.Fullname,
		Email:          data.Email,
		SignInProvider: data.Provider,
		PasswordHash:   passwordHash,
	}
}
