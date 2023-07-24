package post_handlers

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"strconv"
	"strings"
	"time"

	"github.com/je-martinez/2023-go-rest-api/pkg/DTOs"
	"github.com/je-martinez/2023-go-rest-api/pkg/bucket_manager"
	"github.com/je-martinez/2023-go-rest-api/pkg/constants"
	"github.com/je-martinez/2023-go-rest-api/pkg/database/entities"
	auth_types "github.com/je-martinez/2023-go-rest-api/pkg/types/auth"
	router_types "github.com/je-martinez/2023-go-rest-api/pkg/types/router"
	"github.com/je-martinez/2023-go-rest-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

func CreatePost(props *router_types.RouterHandlerProps) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		tmpCurrentUser, errCurrentUser := c.Get(constants.CURRENT_USER_KEY_CTX)

		if !errCurrentUser {
			utils.GinApiResponse(c, 500, constants.ERR_CURRENT_USER, nil, nil)
			return
		}

		currentUser := tmpCurrentUser.(*auth_types.CurrentUser)

		var post DTOs.CreatePostDTO

		err := c.ShouldBind(&post)
		if err != nil {
			utils.GinApiResponse(c, 400, constants.ERR_BIND_MULTIPART, nil, []string{err.Error()})
			return
		}

		newPost := post.ToEntity(post.Content, currentUser.UserID)

		err = props.Database.PostRepository.Create(newPost)

		if err != nil {
			utils.GinApiResponse(c, 500, constants.CREATE_POST_ERR, nil, utils.ValidateStructErrors(err))
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		postFiles, err := handleUploadFiles(props.BucketManager, ctx, post.Files, currentUser.UserID, newPost.PostID, currentUser.UserID)

		if len(postFiles) > 0 {
			err := props.Database.FileRepository.CreateBatch(postFiles)
			if err != nil {
				utils.GinApiResponse(c, 500, constants.UPLOAD_POST_FILES_ERR, nil, []string{})
			}
			newPost.Files = append(newPost.Files, postFiles...)

		}

		if err != nil {
			utils.GinApiResponse(c, 500, constants.UPLOAD_POST_FILES_ERR, nil, []string{})
			return
		}

		utils.GinApiResponse(c, 200, "", newPost.ToDTO(), nil)
	})
}

func handleUploadFiles(bucketManager *bucket_manager.MinioApiInstance, ctx context.Context, files []multipart.FileHeader, bucketName string, post_id string, current_user string) ([]entities.File, error) {

	if files == nil {
		//Nothing to do
		return nil, nil
	}

	postFiles := []entities.File{}

	for _, file := range files {
		tmpFile, err := file.Open()
		if err != nil {
			return nil, errors.New(constants.READ_POST_FILE_ERR)
		}
		tmpName := strconv.FormatInt(time.Now().Unix(), 10) + "." + getExtension(file.Filename)
		location := fmt.Sprintf("posts/%s/%s", post_id, tmpName)
		upload, err := bucketManager.UploadFile(ctx, bucketName, location, tmpFile, file.Size)
		if err != nil {
			return nil, err
		}
		postFiles = append(postFiles, entities.File{}.FromMinioUpload(upload.Key, location, current_user, post_id, file))
	}

	return postFiles, nil
}

func getExtension(filename string) string {
	return strings.Split(filename, ".")[1]
}
