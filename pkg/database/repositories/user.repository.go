package repositories

import (
	"errors"
	db "main/pkg/database"
	e "main/pkg/database/entities"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct{}

func (UserRepository) GetAll() (users *[]e.User, err error) {
	Rows := &[]e.User{}

	operation := db.Database.Find(Rows)

	if err := operation.Error; err != nil && !errors.Is(err, gorm.ErrEmptySlice) {
		return Rows, err
	}

	return Rows, nil
}

func (UserRepository) Get(ID string) (user *e.User, err error, not_found bool) {

	Row := &e.User{}

	operation := db.Database.First(Row, ID)

	if operation.Error != nil {
		return &e.User{}, operation.Error, errors.Is(operation.Error, gorm.ErrRecordNotFound)
	}

	return Row, nil, false
}

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

	db.Database.Save(&Row)

	return Row, nil
}

func (r UserRepository) Update(input *UserInput) (user *e.User, err error, not_found bool) {

	RowToUpdated := &e.User{}

	operationFind := db.Database.Where(&e.User{Username: input.Username}).First(RowToUpdated)

	if err := operationFind; err != nil {
		return &e.User{}, operationFind.Error, errors.Is(operationFind.Error, gorm.ErrRecordNotFound)
	}

	RowToUpdated.Fullname = input.Fullname
	RowToUpdated.Email = input.Email
	RowToUpdated.UpdatedAt = time.Now()

	return RowToUpdated, nil, false
}

func (UserRepository) Delete(ID string) (err error, not_found bool) {

	operation := db.Database.Where(&e.User{UserID: ID}).Delete(&e.User{})

	if operation != nil {
		return operation.Error, errors.Is(operation.Error, gorm.ErrRecordNotFound)
	}
	return nil, false
}
