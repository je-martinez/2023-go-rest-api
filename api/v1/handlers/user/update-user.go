package user_handlers

import (
	"context"
	"main/pkg/DTOs"
	"main/pkg/database"
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

	ctx := context.Background()
	_, err = database.UserRepository.FindByID(ctx, updateData.UserId)

	if err != nil {
		utils.GinApiResponse(c, 404, err.Error(), nil, nil)
		return
	}
	utils.GinApiResponse(c, 200, "Great", nil, nil)

}
