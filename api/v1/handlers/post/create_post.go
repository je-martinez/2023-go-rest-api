package post_handlers

import (
	"context"

	"github.com/je-martinez/2023-go-rest-api/pkg/DTOs"
	"github.com/je-martinez/2023-go-rest-api/pkg/constants"
	auth_types "github.com/je-martinez/2023-go-rest-api/pkg/types/auth"
	router_types "github.com/je-martinez/2023-go-rest-api/pkg/types/router"
	"github.com/je-martinez/2023-go-rest-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

func CreatePost(props *router_types.RouterHandlerProps) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		tmpCurrentUser, errCurrentUser := c.Get(constants.CURRENT_USER_KEY_CTX)

		if !errCurrentUser {
			msg := constants.ERR_CURRENT_USER
			utils.GinApiResponse(c, 500, &msg, nil, nil)
			return
		}

		currentUser := tmpCurrentUser.(*auth_types.CurrentUser)

		var post DTOs.CreatePostDTO

		err := c.ShouldBind(&post)
		if err != nil {
			msg := constants.ERR_BIND_MULTIPART
			utils.GinApiResponse(c, 400, &msg, nil, []string{err.Error()})
			return
		}

		newPost := post.ToEntity(post.Content, currentUser.UserID)

		err = props.Database.PostRepository.Create(newPost)

		if err != nil {
			msg := constants.CREATE_POST_ERR
			utils.GinApiResponse(c, 500, &msg, nil, utils.ValidateStructErrors(err))
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		postFiles, err := handleUploadFiles(props.BucketManager, ctx, post.Files, currentUser.UserID, newPost.PostID, currentUser.UserID)

		if len(postFiles) > 0 {
			err := props.Database.FileRepository.CreateBatch(postFiles)
			if err != nil {
				msg := constants.UPLOAD_POST_FILES_ERR
				utils.GinApiResponse(c, 500, &msg, nil, []string{})
			}
			newPost.Files = append(newPost.Files, postFiles...)
		}

		if err != nil {
			msg := constants.UPLOAD_POST_FILES_ERR
			utils.GinApiResponse(c, 500, &msg, nil, []string{})
			return
		}

		utils.GinApiResponse(c, 200, nil, newPost.ToDTO(), nil)
	})
}
