package user_handlers

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

func UpdateUser(props *router_types.RouterHandlerProps) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		var updateData DTOs.UpdateUserDTO
		err := c.BindJSON(&updateData)
		if err != nil {
			utils.GinApiResponse(c, 400, constants.ERR_BIND_JSON, nil, []string{err.Error()})
			return
		}

		err = validate.Struct(updateData)
		if err != nil {
			utils.GinApiResponse(c, 400, constants.ERR_INVALID_JSON, nil, utils.ValidateStructErrors(err))
			return
		}

		currentUser, errCurrentUser := utils.ExtractUserFromToken(c)

		if errCurrentUser != nil || currentUser == nil {
			utils.GinApiResponse(c, 500, constants.ERR_CURRENT_USER, nil, utils.ValidateStructErrors(errCurrentUser))
			return
		}

		query := types.QueryOptions{Query: entities.User{UserID: currentUser.UserID}}
		userFind, errUserFind, notFound := props.Database.UserRepository.Find(query)

		if errUserFind != nil {
			if notFound {
				utils.GinApiResponse(c, 404, fmt.Sprintf(constants.ERR_ENTITY_NOT_FOUND_ID, "User", currentUser.UserID), nil, nil)
				return
			}
			utils.GinApiResponse(c, 500, fmt.Sprintf(constants.ERR_FIND_ENTITY, "User"), nil, nil)
			return
		}

		errMsg, newPasswordHash, hasErr := handlerPasswordChange(updateData.OldPassword, updateData.NewPassword, userFind.PasswordHash)

		if hasErr {
			utils.GinApiResponse(c, 400, errMsg, nil, nil)
			return
		}

		if newPasswordHash != "" {
			userFind.PasswordHash = newPasswordHash
		}
		userFind.Fullname = updateData.Fullname
		userFind.Email = updateData.Email

		_, errUpdate := props.Database.UserRepository.Update(userFind)

		if errUpdate != nil {
			utils.GinApiResponse(c, 500, fmt.Sprintf(constants.ERR_UPDATE_ENTITY, "User"), nil, []string{errUpdate.Error()})
			return
		}

		utils.GinApiResponse(c, 200, "", userFind.ToDTO(), nil)
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
