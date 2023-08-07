package reactions_handlers

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

func RemovePostReaction(props *router_types.RouterHandlerProps) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		tmpCurrentUser, errCurrentUser := c.Get(constants.CURRENT_USER_KEY_CTX)

		if !errCurrentUser {
			utils.GinApiResponse(c, 500, constants.ERR_CURRENT_USER, nil, nil)
			return
		}

		currentUser := tmpCurrentUser.(*auth_types.CurrentUser)

		post_id := c.Param("post_id")
		reaction_type := c.Param("reaction_type")

		isValidReaction, Reaction := utils.IsSupportedReaction(reaction_type)
		if !isValidReaction {
			utils.GinApiResponse(c, 404, fmt.Sprintf(constants.UNSUPPORTED_REACTION, reaction_type), nil, nil)
			return
		}

		query := types.QueryOptions{
			Query:    entities.Post{PostID: post_id},
			Preloads: []string{},
		}

		_, err, postNotFound := props.Database.PostRepository.Find(query)

		if err != nil {
			if postNotFound {
				utils.GinApiResponse(c, 404, fmt.Sprintf(constants.ERR_ENTITY_NOT_FOUND_ID, "Post", post_id), nil, nil)
				return
			}
			utils.GinApiResponse(c, 500, fmt.Sprintf(constants.ERR_FIND_ENTITY, "Post"), nil, nil)
			return
		}

		query = types.QueryOptions{
			Query:    entities.Reaction{PostID: &post_id, ReactionType: *Reaction, UserID: currentUser.UserID},
			Preloads: []string{},
		}

		reaction, err, notFound := props.Database.ReactionRepository.Find(query)
		if err != nil {
			if notFound {
				utils.GinApiResponse(c, 404, fmt.Sprintf(constants.ERR_ENTITY_NOT_FOUND, "Reaction"), nil, nil)
				return
			}
			utils.GinApiResponse(c, 500, fmt.Sprintf(constants.ERR_FIND_ENTITY, "Reaction"), nil, nil)
			return
		}

		_, err = props.Database.ReactionRepository.Delete(&entities.Reaction{ReactionID: reaction.ReactionID})

		if err != nil {
			utils.GinApiResponse(c, 400, fmt.Sprintf(constants.ERR_DELETE_ENTITY, "Reaction"), nil, nil)
			return
		}

		utils.GinApiResponse(c, 200, "", nil, nil)
	})
}
