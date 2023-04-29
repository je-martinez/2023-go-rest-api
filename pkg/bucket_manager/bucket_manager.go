package bucket_manager

import (
	"context"
	"main/config"
	"main/pkg/constants"
	l "main/pkg/logger"

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
	err := MinioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := MinioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			l.ApiLogger.Infof(constants.BUCKET_ALREADY_EXISTS, bucketName)
			return false
		} else {
			l.ApiLogger.Error(constants.BUCKET_CREATION_ERROR, err.Error())
			return false
		}
	} else {
		l.ApiLogger.Info(constants.BUCKET_CREATED)
		return true
	}
}
