package auth_handlers

import (
	"fmt"

	"github.com/je-martinez/2023-go-rest-api/pkg/DTOs"
	"github.com/je-martinez/2023-go-rest-api/pkg/constants"
	"github.com/je-martinez/2023-go-rest-api/pkg/database"
	"github.com/je-martinez/2023-go-rest-api/pkg/database/entities"
	types "github.com/je-martinez/2023-go-rest-api/pkg/types/database"
	"github.com/je-martinez/2023-go-rest-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var loginData DTOs.LoginDTO
	err := c.BindJSON(&loginData)
	if err != nil {
		utils.GinApiResponse(c, 400, constants.ERR_BIND_JSON, nil, []string{err.Error()})
		return
	}

	err = validate.Struct(loginData)
	if err != nil {
		utils.GinApiResponse(c, 400, constants.ERR_INVALID_JSON, nil, utils.ValidateStructErrors(err))
		return
	}

	query := types.QueryOptions{
		Query: entities.User{Username: loginData.Username},
	}

	foundUser, notFound, err := database.GlobalInstance.UserRepository.Find(query)

	if err != nil {
		if notFound {
			utils.GinApiResponse(c, 404, fmt.Sprintf(constants.ERR_ENTITY_NOT_FOUND, "User"), nil, nil)
			return
		}
	}

	isPasswordValid := utils.CheckPasswordHash(loginData.Password, foundUser.PasswordHash)

	if !isPasswordValid {
		utils.GinApiResponse(c, 401, constants.ERR_USERNAME_PASSWORD_INVALID, nil, nil)
		return
	}

	token, err := utils.GenerateToken(*foundUser)

	if err != nil {
		utils.GinApiResponse(c, 500, constants.ERR_GENERATE_TOKEN, nil, nil)
		return
	}

	responseData := &DTOs.AuthResponseDTO{
		Username: foundUser.Username,
		Fullname: foundUser.Fullname,
		Email:    foundUser.Email,
		Provider: foundUser.SignInProvider,
		Token:    token,
	}

	utils.GinApiResponse(c, 200, "", responseData, nil)
}
