package bucket_manager

import (
	"context"

	"github.com/je-martinez/2023-go-rest-api/config"
	"github.com/je-martinez/2023-go-rest-api/pkg/constants"
	l "github.com/je-martinez/2023-go-rest-api/pkg/logger"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

func Start(cfg *config.Config) *minio.Client {
	// Initialize minio client object.
	minioClient, err := minio.New(cfg.AWS.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AWS.MinioAccessKey, cfg.AWS.MinioSecretKey, ""),
		Secure: cfg.AWS.UseSSL,
	})
	if err != nil {
		l.ApiLogger.Fatal(constants.STARTING_BUCKET_MANAGER, err)
		return nil
	}
	MinioClient = minioClient
	l.ApiLogger.Info(constants.BUCKET_MANAGER_STARTED)

	return MinioClient
}

func CreateBucket(ctx context.Context, bucketName string, location string) bool {

	if MinioClient == nil {
		l.ApiLogger.Errorf(constants.BUCKET_MANAGER_NOT_STARTED, bucketName)
		return false
	}

	err := MinioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		return validateIfBucketExist(ctx, bucketName)
	} else {
		l.ApiLogger.Info(constants.BUCKET_CREATED)
		return true
	}
}

func DeleteBucket(ctx context.Context, bucketName string) bool {
	if MinioClient == nil {
		l.ApiLogger.Errorf(constants.BUCKET_MANAGER_NOT_STARTED, bucketName)
		return false
	}
	if !validateIfBucketExist(ctx, bucketName) {
		return false
	}

	err := MinioClient.RemoveBucket(ctx, bucketName)
	if err != nil {
		l.ApiLogger.Infof(constants.BUCKET_DELETE_ERROR, bucketName)
		return false
	}
	l.ApiLogger.Infof(constants.BUCKET_DELETED, bucketName)
	return true
}

func validateIfBucketExist(ctx context.Context, bucketName string) bool {
	exists, errBucketExists := MinioClient.BucketExists(ctx, bucketName)
	if errBucketExists == nil && exists {
		l.ApiLogger.Infof(constants.BUCKET_ALREADY_EXISTS, bucketName)
		return true
	} else {
		l.ApiLogger.Error(constants.BUCKET_CREATION_ERROR, errBucketExists.Error())
		return false
	}
}
