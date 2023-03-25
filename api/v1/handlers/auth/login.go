package auth_handlers

import (
	"main/pkg/DTOs"
	"main/pkg/database"
	"main/pkg/database/entities"
	"main/pkg/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var loginData DTOs.LoginDTO
	err := c.BindJSON(&loginData)
	if err != nil {
		utils.GinApiResponse(c, 400, "Error binding JSON", nil, []string{err.Error()})
		return
	}

	err = validate.Struct(loginData)
	if err != nil {
		utils.GinApiResponse(c, 400, "Error with the provided JSON", nil, utils.ValidateStructErrors(err))
		return
	}

	foundUser, err, notFound := database.UserRepository.Find(entities.User{Username: loginData.Username})

	if err != nil {
		if notFound {
			utils.GinApiResponse(c, 404, "User not found in our database", nil, nil)
			return
		}
	}

	isPasswordValid := utils.CheckPasswordHash(loginData.Password, foundUser.PasswordHash)

	if !isPasswordValid {
		utils.GinApiResponse(c, 401, "Username or password are invalid", nil, nil)
		return
	}

	token, err := utils.GenerateToken(*foundUser)

	if err != nil {
		utils.GinApiResponse(c, 500, "Error trying to generate a new access token", nil, nil)
		return
	}

	responseData := &DTOs.LoginResponse{
		Token: token,
	}

	utils.GinApiResponse(c, 200, "", responseData, nil)
	return
}
