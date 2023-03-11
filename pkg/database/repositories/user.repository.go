package repositories

import (
	db "main/pkg/database"
	e "main/pkg/database/entities"
)

type UserRepository struct{}

func (UserRepository) Create(input *UserInput) (user *e.User, err error) {

	Row := &e.User{
		Username:       input.Username,
		Email:          input.Email,
		Fullname:       input.Fullname,
		PasswordHash:   input.PasswordHash,
		SignInProvider: input.SignInProvider,
	}

	operation := db.Database.Create(&Row)

	if operation.Error != nil {
		return &e.User{}, operation.Error
	}

	return Row, nil
}
