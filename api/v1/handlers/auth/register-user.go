package auth_handlers

import (
	"context"
	"fmt"
	"main/pkg/DTOs"
	"main/pkg/database"
	"main/pkg/utils"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {

	var registerData DTOs.RegisterUserDTO
	err := c.BindJSON(&registerData)
	if err != nil {
		utils.GinApiResponse(c, 400, "Error binding JSON", err, nil)
		return
	}
	ctx := context.Background()
	errInsert := database.UserRepository.Insert(ctx, registerData.ToModel("Holis"))
	if errInsert != nil {
		fmt.Println(errInsert)
		utils.GinApiResponse(c, 400, "Error Creating User", nil, errInsert.Error())
		return
	}
	utils.GinApiResponse(c, 200, "New User Created!", nil, nil)
}
