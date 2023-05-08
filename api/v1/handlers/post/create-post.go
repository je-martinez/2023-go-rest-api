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
	"github.com/je-martinez/2023-go-rest-api/pkg/database"
	"github.com/je-martinez/2023-go-rest-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {

	currentUser, errCurrentUser := utils.ExtractUserFromToken(c)

	if errCurrentUser != nil || currentUser == nil {
		utils.GinApiResponse(c, 500, constants.ERR_CURRENT_USER, nil, utils.ValidateStructErrors(errCurrentUser))
		return
	}

	var post DTOs.CreatePostDTO

	err := c.ShouldBind(&post)
	if err != nil {
		utils.GinApiResponse(c, 400, constants.ERR_BIND_MULTIPART, nil, []string{err.Error()})
		return
	}

	newPost := post.ToEntity(post.Content, currentUser.UserID)

	err = database.PostRepository.Create(newPost)

	if err != nil {
		utils.GinApiResponse(c, 500, constants.CREATE_POST_ERR, nil, utils.ValidateStructErrors(errCurrentUser))
		return
	}

	ctx := context.Background()
	err = handleUploadFiles(ctx, post.Files, currentUser.UserID, newPost.PostID)

	if err != nil {
		utils.GinApiResponse(c, 500, constants.UPLOAD_POST_FILES_ERR, nil, []string{})
		return
	}

	utils.GinApiResponse(c, 200, "", newPost.ToDTO(), nil)
}

func handleUploadFiles(ctx context.Context, files []multipart.FileHeader, bucketName string, post_id string) error {

	if files == nil {
		//Nothing to do
		return nil
	}

	for _, file := range files {

		tmpFile, err := file.Open()

		if err != nil {
			return errors.New(constants.READ_POST_FILE_ERR)
		}
		tmpName := strconv.FormatInt(time.Now().Unix(), 10) + "." + getExtension(file.Filename)
		location := fmt.Sprintf("%s/%s", post_id, tmpName)
		bucket_manager.UploadFile(ctx, bucketName, location, tmpFile, file.Size)
	}

	return nil
}

func getExtension(filename string) string {
	return strings.Split(filename, ".")[1]
}
