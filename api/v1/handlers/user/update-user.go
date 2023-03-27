package user_handlers

import (
	"fmt"
	"main/pkg/DTOs"
	"main/pkg/constants"
	"main/pkg/database"
	"main/pkg/database/entities"
	"main/pkg/utils"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {

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

	query := entities.User{UserID: currentUser.UserID}
	userFind, errUserFind, notfound := database.UserRepository.Find(query)

	if errUserFind != nil {
		if notfound {
			utils.GinApiResponse(c, 404, fmt.Sprintf(constants.ERR_ENTITY_NOT_FOUND_ID, "User", currentUser.UserID), nil, nil)
			return
		}
		utils.GinApiResponse(c, 500, errUserFind.Error(), nil, nil)
		return
	}

	utils.GinApiResponse(c, 200, "", userFind.ToDTO(), nil)
	return
}
