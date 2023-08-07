package user_handlers

import (
	"fmt"

	"github.com/je-martinez/2023-go-rest-api/pkg/constants"
	"github.com/je-martinez/2023-go-rest-api/pkg/database/entities"
	auth_types "github.com/je-martinez/2023-go-rest-api/pkg/types/auth"
	types "github.com/je-martinez/2023-go-rest-api/pkg/types/database"
	router_types "github.com/je-martinez/2023-go-rest-api/pkg/types/router"
	"github.com/je-martinez/2023-go-rest-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

func Me(props *router_types.RouterHandlerProps) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		tmpCurrentUser, errCurrentUser := c.Get(constants.CURRENT_USER_KEY_CTX)

		if !errCurrentUser {
			msg := constants.ERR_CURRENT_USER
			utils.GinApiResponse(c, 500, &msg, nil, nil)
			return
		}

		currentUser := tmpCurrentUser.(*auth_types.CurrentUser)

		query := types.QueryOptions{
			Query:    entities.User{UserID: currentUser.UserID},
			Preloads: []string{"Profile"},
		}
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

		utils.GinApiResponse(c, 200, nil, userFind.ToDTO(), nil)
	})
}
