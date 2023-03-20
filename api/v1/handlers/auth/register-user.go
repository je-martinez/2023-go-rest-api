package auth_handlers

import (
	"context"
	"fmt"
	"main/pkg/DTOs"
	"main/pkg/database"
	"main/pkg/utils"

	"github.com/gin-gonic/gin"
	v "github.com/go-playground/validator/v10"
)

var validate = v.New()

func RegisterUser(c *gin.Context) {

	var registerData DTOs.RegisterUserDTO
	err := c.BindJSON(&registerData)
	if err != nil {
		utils.GinApiResponse(c, 400, "Error binding JSON", nil, nil)
		return
	}

	err = validate.Struct(registerData)
	if err != nil {
		utils.GinApiResponse(c, 400, "Error with the provided JSON", nil, utils.ValidateStructErrors(err))
		return
	}

	ctx := context.Background()
	passwordHash, _ := utils.GenerateHash(registerData.Password)
	record, errInsert := database.UserRepository.Insert(ctx, registerData.ToModel(passwordHash))
	if errInsert != nil {
		fmt.Println(errInsert)
		utils.GinApiResponse(c, 400, "Error Creating User", nil, errInsert.Error())
		return
	}
	utils.GinApiResponse(c, 200, "New User Created!", record.ToEntity(), nil)
}
