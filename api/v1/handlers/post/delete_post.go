package post_handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/je-martinez/2023-go-rest-api/pkg/constants"
	"github.com/je-martinez/2023-go-rest-api/pkg/database/entities"
	router_types "github.com/je-martinez/2023-go-rest-api/pkg/types/router"
	"github.com/je-martinez/2023-go-rest-api/pkg/utils"
)

func DeletePost(props *router_types.RouterHandlerProps) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {

		post_id := c.Param("post_id")

		_, err, notFound := props.Database.PostRepository.FindByStringID(post_id, "Files")

		if err != nil {
			if notFound {
				utils.GinApiResponse(c, 400, fmt.Sprintf(constants.POST_NOT_FOUND, post_id), nil, []string{err.Error()})
				return
			}
			utils.GinApiResponse(c, 400, fmt.Sprintf(constants.FETCH_POST_FAILED, post_id), nil, []string{err.Error()})
			return
		}

		_, err = props.Database.PostRepository.Delete(&entities.Post{PostID: post_id})

		if err != nil {
			utils.GinApiResponse(c, 500, constants.DELETE_POST_ERR, nil, []string{err.Error()})
			return
		}

		utils.GinApiResponse(c, 200, "", nil, []string{err.Error()})
	})
}
