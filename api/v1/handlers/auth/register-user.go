package auth_handlers

import (
	"context"
	"main/pkg/DTOs"
	"main/pkg/database"
	"main/pkg/utils"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {

	var registerData DTOs.RegisterUserDTO
	err := c.BindJSON(&registerData)
	if err != nil {
		utils.GinApiResponse(c, 400, "Error binding JSON", nil, []string{err.Error()})
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
		utils.GinApiResponse(c, 400, "Error Creating User", nil, []string{errInsert.Error()})
		return
	}
	utils.GinApiResponse(c, 200, "User Created!", record.ToModel(), nil)
}
