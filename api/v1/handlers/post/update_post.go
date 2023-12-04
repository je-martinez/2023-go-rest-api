package post_handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/je-martinez/2023-go-rest-api/pkg/DTOs"
	"github.com/je-martinez/2023-go-rest-api/pkg/constants"
	"github.com/je-martinez/2023-go-rest-api/pkg/database/entities"
	types "github.com/je-martinez/2023-go-rest-api/pkg/types/database"
	router_types "github.com/je-martinez/2023-go-rest-api/pkg/types/router"
	"github.com/je-martinez/2023-go-rest-api/pkg/utils"
)

func UpdatePost(props *router_types.RouterHandlerProps) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		_, errCurrentUser := c.Get(constants.CURRENT_USER_KEY_CTX)

		if !errCurrentUser {
			msg := constants.ERR_CURRENT_USER
			utils.GinApiResponse(c, 500, &msg, nil, nil)
			return
		}

		var postToEdit DTOs.EditPostDTO
		post_id := c.Param("post_id")
		errBind := c.ShouldBind(&postToEdit)
		if errBind != nil {
			msg := constants.ERR_BIND_MULTIPART
			utils.GinApiResponse(c, 400, &msg, nil, []string{errBind.Error()})
			return
		}

		query := types.QueryOptions{
			Query: entities.Post{PostID: post_id},
		}

		post, err, notFound := props.Database.PostRepository.Find(query)

		if err != nil {
			if notFound {
				msg := fmt.Sprintf(constants.POST_NOT_FOUND, post_id)
				utils.GinApiResponse(c, 400, &msg, nil, []string{err.Error()})
				return
			}
			msg := fmt.Sprintf(constants.FETCH_POST_FAILED, post_id)
			utils.GinApiResponse(c, 400, &msg, nil, []string{err.Error()})
			return
		}

		post.Content = postToEdit.Content
		_, err = props.Database.PostRepository.Update(post)

		if err != nil {
			msg := constants.UPDATE_POST_ERR
			utils.GinApiResponse(c, 500, &msg, nil, []string{err.Error()})
			return
		}

		utils.GinApiResponse(c, 200, nil, post.ToDTO(), nil)
	})
}
