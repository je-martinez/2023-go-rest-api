package post_handlers

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/je-martinez/2023-go-rest-api/pkg/bucket_manager"
	"github.com/je-martinez/2023-go-rest-api/pkg/constants"
	"github.com/je-martinez/2023-go-rest-api/pkg/database/entities"
	auth_types "github.com/je-martinez/2023-go-rest-api/pkg/types/auth"
	types "github.com/je-martinez/2023-go-rest-api/pkg/types/database"
	router_types "github.com/je-martinez/2023-go-rest-api/pkg/types/router"
	"github.com/je-martinez/2023-go-rest-api/pkg/utils"
)

func DeletePost(props *router_types.RouterHandlerProps) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		tmpCurrentUser, errCurrentUser := c.Get(constants.CURRENT_USER_KEY_CTX)

		if !errCurrentUser {
			msg := constants.ERR_CURRENT_USER
			utils.GinApiResponse(c, 500, &msg, nil, nil)
			return
		}

		currentUser := tmpCurrentUser.(*auth_types.CurrentUser)

		post_id := c.Param("post_id")

		query := types.QueryOptions{
			Query: entities.Post{PostID: post_id},
		}

		_, err, notFound := props.Database.PostRepository.Find(query)

		log.Println("Hola")

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

		_, err = props.Database.PostRepository.Delete(&entities.Post{PostID: post_id})

		if err != nil {
			msg := constants.DELETE_POST_ERR
			utils.GinApiResponse(c, 500, &msg, nil, []string{err.Error()})
			return
		}

		go handleDeleteFiles(props.BucketManager, currentUser.UserID, post_id)

		utils.GinApiResponse(c, 200, nil, nil, nil)
	})
}

func handleDeleteFiles(bucketManager *bucket_manager.MinioApiInstance, bucketName string, postId string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	bucketManager.DeleteFolder(ctx, bucketName, fmt.Sprintf("posts/%s", postId))
}
