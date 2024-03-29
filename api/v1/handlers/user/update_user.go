package user_handlers

import (
	"fmt"

	"github.com/je-martinez/2023-go-rest-api/pkg/DTOs"
	"github.com/je-martinez/2023-go-rest-api/pkg/constants"
	"github.com/je-martinez/2023-go-rest-api/pkg/database/entities"
	auth_types "github.com/je-martinez/2023-go-rest-api/pkg/types/auth"
	types "github.com/je-martinez/2023-go-rest-api/pkg/types/database"
	router_types "github.com/je-martinez/2023-go-rest-api/pkg/types/router"
	"github.com/je-martinez/2023-go-rest-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

func UpdateUser(props *router_types.RouterHandlerProps) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		var updateData DTOs.UpdateUserDTO
		err := c.BindJSON(&updateData)
		if err != nil {
			msg := constants.ERR_BIND_JSON
			utils.GinApiResponse(c, 400, &msg, nil, []string{err.Error()})
			return
		}

		err = validate.Struct(updateData)
		if err != nil {
			msg := constants.ERR_INVALID_JSON
			utils.GinApiResponse(c, 400, &msg, nil, utils.ValidateStructErrors(err))
			return
		}

		tmpCurrentUser, errCurrentUser := c.Get(constants.CURRENT_USER_KEY_CTX)

		if !errCurrentUser {
			msg := constants.ERR_CURRENT_USER
			utils.GinApiResponse(c, 500, &msg, nil, nil)
			return
		}

		currentUser := tmpCurrentUser.(*auth_types.CurrentUser)

		query := types.QueryOptions{Query: entities.User{UserID: currentUser.UserID}}
		userFind, errUserFind, notFound := props.Database.UserRepository.Find(query)

		if errUserFind != nil {
			if notFound {
				msg := fmt.Sprintf(constants.ERR_ENTITY_NOT_FOUND_ID, "User", currentUser.UserID)
				utils.GinApiResponse(c, 404, &msg, nil, nil)
				return
			}
			msg := fmt.Sprintf(constants.ERR_FIND_ENTITY, "User")
			utils.GinApiResponse(c, 500, &msg, nil, nil)
			return
		}

		errMsg, newPasswordHash, hasErr := handlerPasswordChange(updateData.OldPassword, updateData.NewPassword, userFind.PasswordHash)

		if hasErr {
			utils.GinApiResponse(c, 400, &errMsg, nil, nil)
			return
		}

		if newPasswordHash != "" {
			userFind.PasswordHash = newPasswordHash
		}
		userFind.Fullname = updateData.Fullname
		userFind.Email = updateData.Email

		_, errUpdate := props.Database.UserRepository.Update(userFind)

		if errUpdate != nil {
			msg := fmt.Sprintf(constants.ERR_UPDATE_ENTITY, "User")
			utils.GinApiResponse(c, 500, &msg, nil, []string{errUpdate.Error()})
			return
		}

		utils.GinApiResponse(c, 200, nil, userFind.ToDTO(), nil)
	})
}

func handlerPasswordChange(oldPassword string, newPassword string, hash string) (string, string, bool) {
	if oldPassword == "" || newPassword == "" {
		//Ignoring password change
		return "", "", false
	}

	isOldPasswordValid := utils.CheckPasswordHash(oldPassword, hash)

	if !isOldPasswordValid {
		//Error old password is incorrect
		return constants.ERR_OLD_PASSWORD_MISMATCH, "", true
	}

	newPasswordHash, errHash := utils.GenerateHash(newPassword)

	if errHash != nil {
		//Error generating hash
		return fmt.Sprintf(constants.ERR_GENERATE_HASH, "password"), "", true
	}

	return "", newPasswordHash, false
}
