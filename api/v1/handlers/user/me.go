package user_handlers

import (
	"fmt"

	"github.com/je-martinez/2023-go-rest-api/pkg/constants"
	"github.com/je-martinez/2023-go-rest-api/pkg/database/entities"
	types "github.com/je-martinez/2023-go-rest-api/pkg/types/database"
	router_types "github.com/je-martinez/2023-go-rest-api/pkg/types/router"
	"github.com/je-martinez/2023-go-rest-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

func Me(props *router_types.RouterHandlerProps) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		currentUser, errCurrentUser := utils.ExtractUserFromToken(c)

		if errCurrentUser != nil || currentUser == nil {
			utils.GinApiResponse(c, 500, constants.ERR_CURRENT_USER, nil, utils.ValidateStructErrors(errCurrentUser))
			return
		}

		query := types.QueryOptions{
			Query:    entities.User{UserID: currentUser.UserID},
			Preloads: []string{"Profile"},
		}
		userFind, notfound, errUserFind := props.Database.UserRepository.Find(query)

		if errUserFind != nil {
			if notfound {
				utils.GinApiResponse(c, 404, fmt.Sprintf(constants.ERR_ENTITY_NOT_FOUND_ID, "User", currentUser.UserID), nil, nil)
				return
			}
			utils.GinApiResponse(c, 500, fmt.Sprintf(constants.ERR_FIND_ENTITY, "User"), nil, nil)
			return
		}

		utils.GinApiResponse(c, 200, "", userFind.ToDTO(), nil)
	})
}
