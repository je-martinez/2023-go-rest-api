package auth_handlers

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/je-martinez/2023-go-rest-api/pkg/DTOs"
	"github.com/je-martinez/2023-go-rest-api/pkg/constants"
	router_types "github.com/je-martinez/2023-go-rest-api/pkg/types/router"
	"github.com/je-martinez/2023-go-rest-api/pkg/utils"
)

func RegisterUser(props *router_types.RouterHandlerProps) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {

		var registerData DTOs.RegisterUserDTO
		err := c.BindJSON(&registerData)
		if err != nil {
			msg := constants.ERR_BIND_JSON
			utils.GinApiResponse(c, 400, &msg, nil, []string{err.Error()})
			return
		}

		err = validate.Struct(registerData)
		if err != nil {
			msg := constants.ERR_INVALID_JSON
			utils.GinApiResponse(c, 400, &msg, nil, utils.ValidateStructErrors(err))
			return
		}

		passwordHash, _ := utils.GenerateHash(registerData.Password)
		newRecord := registerData.ToEntity(passwordHash)
		errInsert := props.Database.UserRepository.Create(newRecord)
		if errInsert != nil {
			msg := fmt.Sprintf(constants.ERR_CREATE_ENTITY, "User")
			utils.GinApiResponse(c, 400, &msg, nil, []string{errInsert.Error()})
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		bucketCreated := props.BucketManager.CreateBucket(ctx, newRecord.UserID, constants.US_EAST_NORTH_VIRGINIA)

		if !bucketCreated {
			msg := fmt.Sprintf(constants.BUCKET_CREATION_USER_ERROR, newRecord.UserID)
			utils.GinApiResponse(c, 500, &msg, nil, []string{errInsert.Error()})
			return
		}

		token, err := utils.GenerateToken(*newRecord)

		if err != nil {
			msg := constants.ERR_GENERATE_TOKEN
			utils.GinApiResponse(c, 500, &msg, nil, nil)
			return
		}

		responseData := &DTOs.AuthResponseDTO{
			Username: newRecord.Username,
			Fullname: newRecord.Fullname,
			Email:    newRecord.Email,
			Provider: newRecord.SignInProvider,
			Token:    token,
		}

		utils.GinApiResponse(c, 200, nil, responseData, nil)
	})
}
