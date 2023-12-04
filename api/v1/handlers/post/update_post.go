package post_handlers

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/je-martinez/2023-go-rest-api/pkg/DTOs"
	"github.com/je-martinez/2023-go-rest-api/pkg/constants"
	"github.com/je-martinez/2023-go-rest-api/pkg/database/entities"
	auth_types "github.com/je-martinez/2023-go-rest-api/pkg/types/auth"
	types "github.com/je-martinez/2023-go-rest-api/pkg/types/database"
	router_types "github.com/je-martinez/2023-go-rest-api/pkg/types/router"
	"github.com/je-martinez/2023-go-rest-api/pkg/utils"
)

func UpdatePost(props *router_types.RouterHandlerProps) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		tmpCurrentUser, errCurrentUser := c.Get(constants.CURRENT_USER_KEY_CTX)

		if !errCurrentUser {
			msg := constants.ERR_CURRENT_USER
			utils.GinApiResponse(c, 500, &msg, nil, nil)
			return
		}

		currentUser := tmpCurrentUser.(*auth_types.CurrentUser)

		var postToEdit DTOs.EditPostDTO
		post_id := c.Param("post_id")
		errBind := c.ShouldBind(&postToEdit)
		if errBind != nil {
			msg := constants.ERR_BIND_MULTIPART
			utils.GinApiResponse(c, 400, &msg, nil, []string{errBind.Error()})
			return
		}

		query := types.QueryOptions{
			Query:    entities.Post{PostID: post_id},
			Preloads: []string{"Comments", "Files", "Reactions"},
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
		post.UpdatedBy = &currentUser.UserID
		_, err = props.Database.PostRepository.Update(post)
		if err != nil {
			msg := constants.UPDATE_POST_ERR
			utils.GinApiResponse(c, 500, &msg, nil, []string{err.Error()})
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		//Create new added files
		postFiles, err := handleUploadFiles(props.BucketManager, ctx, postToEdit.NewFilesToAdd, currentUser.UserID, post.PostID, currentUser.UserID)

		if err != nil {
			msg := constants.UPLOAD_POST_FILES_ERR
			utils.GinApiResponse(c, 500, &msg, nil, []string{})
			return
		}

		if len(postFiles) > 0 {
			err := props.Database.FileRepository.CreateBatch(postFiles)
			if err != nil {
				msg := constants.UPLOAD_POST_FILES_ERR
				utils.GinApiResponse(c, 500, &msg, nil, []string{})
			}
			post.Files = append(post.Files, postFiles...)
		}

		utils.GinApiResponse(c, 200, nil, post.ToDTO(), nil)
	})
}
