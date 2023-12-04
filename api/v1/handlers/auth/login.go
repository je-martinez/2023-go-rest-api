package auth_handlers

import (
	"fmt"

	"github.com/je-martinez/2023-go-rest-api/pkg/DTOs"
	"github.com/je-martinez/2023-go-rest-api/pkg/constants"
	"github.com/je-martinez/2023-go-rest-api/pkg/database/entities"
	types "github.com/je-martinez/2023-go-rest-api/pkg/types/database"
	router_types "github.com/je-martinez/2023-go-rest-api/pkg/types/router"
	"github.com/je-martinez/2023-go-rest-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

func Login(props *router_types.RouterHandlerProps) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		var loginData DTOs.LoginDTO
		err := c.BindJSON(&loginData)
		if err != nil {
			msg := constants.ERR_BIND_JSON
			utils.GinApiResponse(c, 400, &msg, nil, []string{err.Error()})
			return
		}

		err = validate.Struct(loginData)
		if err != nil {
			msg := constants.ERR_INVALID_JSON
			utils.GinApiResponse(c, 400, &msg, nil, utils.ValidateStructErrors(err))
			return
		}

		query := types.QueryOptions{
			Query: entities.User{Username: loginData.Username},
		}

		foundUser, err, notFound := props.Database.UserRepository.Find(query)

		if err != nil {
			if notFound {
				msg := fmt.Sprintf(constants.ERR_ENTITY_NOT_FOUND, "User")
				utils.GinApiResponse(c, 404, &msg, nil, nil)
				return
			}
		}

		isPasswordValid := utils.CheckPasswordHash(loginData.Password, foundUser.PasswordHash)

		if !isPasswordValid {
			msg := constants.ERR_USERNAME_PASSWORD_INVALID
			utils.GinApiResponse(c, 401, &msg, nil, nil)
			return
		}

		token, err := utils.GenerateToken(*foundUser)

		if err != nil {
			msg := constants.ERR_GENERATE_TOKEN
			utils.GinApiResponse(c, 500, &msg, nil, nil)
			return
		}

		responseData := &DTOs.AuthResponseDTO{
			Username: foundUser.Username,
			Fullname: foundUser.Fullname,
			Email:    foundUser.Email,
			Provider: foundUser.SignInProvider,
			Token:    token,
		}

		utils.GinApiResponse(c, 200, nil, responseData, nil)
	})
}
