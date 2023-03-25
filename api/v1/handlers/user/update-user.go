package user_handlers

import (
	"fmt"
	"main/pkg/DTOs"
	"main/pkg/database"
	"main/pkg/database/entities"
	"main/pkg/utils"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {

	var updateData DTOs.UpdateUserDTO
	err := c.BindJSON(&updateData)
	if err != nil {
		utils.GinApiResponse(c, 400, "Error binding JSON", nil, []string{err.Error()})
		return
	}

	err = validate.Struct(updateData)
	if err != nil {
		utils.GinApiResponse(c, 400, "Error with the provided JSON", nil, utils.ValidateStructErrors(err))
		return
	}

	query := entities.User{UserID: updateData.UserId}
	userFind, errUserFind := database.UserRepository.Find(query)

	if errUserFind != nil {
		if utils.EntityNotFound(errUserFind) {
			utils.GinApiResponse(c, 404, fmt.Sprintf("User not found with id: %s", updateData.UserId), nil, nil)
			return
		}
		utils.GinApiResponse(c, 500, errUserFind.Error(), nil, nil)
		return
	}

	utils.GinApiResponse(c, 200, "", userFind.ToDTO(), nil)
	return
}
