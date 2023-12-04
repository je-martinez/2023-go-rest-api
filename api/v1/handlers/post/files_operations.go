package post_handlers

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"strconv"
	"strings"
	"time"

	"github.com/je-martinez/2023-go-rest-api/pkg/bucket_manager"
	"github.com/je-martinez/2023-go-rest-api/pkg/constants"
	"github.com/je-martinez/2023-go-rest-api/pkg/database/entities"
)

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

func handleDeleteFiles(bucketManager *bucket_manager.MinioApiInstance, bucketName string, postId string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	bucketManager.DeleteFolder(ctx, bucketName, fmt.Sprintf("posts/%s", postId))
}

func handleDeleteFilesByKeys(bucketManager *bucket_manager.MinioApiInstance, bucketName string, postId string, keys []string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	bucketManager.DeleteFiles(ctx, bucketName, postId, keys)
}
