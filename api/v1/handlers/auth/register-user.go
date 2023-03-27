package auth_handlers

import (
	"fmt"
	"main/pkg/DTOs"
	"main/pkg/constants"
	"main/pkg/database"
	"main/pkg/utils"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {

	var registerData DTOs.RegisterUserDTO
	err := c.BindJSON(&registerData)
	if err != nil {
		utils.GinApiResponse(c, 400, constants.ERR_BIND_JSON, nil, []string{err.Error()})
		return
	}

	err = validate.Struct(registerData)
	if err != nil {
		utils.GinApiResponse(c, 400, constants.ERR_INVALID_JSON, nil, utils.ValidateStructErrors(err))
		return
	}

	passwordHash, _ := utils.GenerateHash(registerData.Password)
	newRecord := registerData.ToEntity(passwordHash)
	errInsert := database.UserRepository.Create(newRecord)
	if errInsert != nil {
		utils.GinApiResponse(c, 400, fmt.Sprintf(constants.ERR_CREATE_ENTITY, "User"), nil, []string{errInsert.Error()})
		return
	}

	token, err := utils.GenerateToken(*newRecord)

	if err != nil {
		utils.GinApiResponse(c, 500, constants.ERR_GENERATE_TOKEN, nil, nil)
		return
	}

	responseData := &DTOs.AuthResponseDTO{
		Username: newRecord.Username,
		Fullname: newRecord.Fullname,
		Email:    newRecord.Email,
		Provider: newRecord.SignInProvider,
		Token:    token,
	}

	utils.GinApiResponse(c, 200, "", responseData, nil)
	return
}
