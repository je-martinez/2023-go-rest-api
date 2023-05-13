package bucket_manager

import (
	"context"
	"errors"
	"mime/multipart"

	"github.com/je-martinez/2023-go-rest-api/config"
	"github.com/je-martinez/2023-go-rest-api/pkg/constants"
	l "github.com/je-martinez/2023-go-rest-api/pkg/logger"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func New(cfg *config.AWS) (*MinioApiInstance, error) {
	client, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MinioAccessKey, cfg.MinioSecretKey, ""),
		Secure: cfg.UseSSL,
	})
	if err != nil {
		l.ApiLogger.Fatal(constants.STARTING_BUCKET_MANAGER, err)
		return nil, err
	}
	l.ApiLogger.Info(constants.BUCKET_MANAGER_STARTED)
	return &MinioApiInstance{
		client: client,
	}, nil
}

type MinioApiInstance struct {
	client *minio.Client
}

func (m *MinioApiInstance) ValidateIfBucketExist(ctx context.Context, bucketName string) (sucess bool) {
	exists, errBucketExists := m.client.BucketExists(ctx, bucketName)
	if errBucketExists != nil {
		l.ApiLogger.Error(constants.BUCKET_CREATION_ERROR, errBucketExists.Error())
		return false
	}
	l.ApiLogger.Infof(constants.BUCKET_ALREADY_EXISTS, bucketName)
	return exists
}

func (m *MinioApiInstance) CreateBucket(ctx context.Context, bucketName string, location string) (sucess bool) {
	err := m.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		return m.ValidateIfBucketExist(ctx, bucketName)
	} else {
		l.ApiLogger.Info(constants.BUCKET_CREATED)
		return true
	}
}

func (m *MinioApiInstance) DeleteBucket(ctx context.Context, bucketName string) (sucess bool, notFound bool) {
	if !m.ValidateIfBucketExist(ctx, bucketName) {
		return false, true
	}
	err := m.client.RemoveBucket(ctx, bucketName)
	if err != nil {
		l.ApiLogger.Infof(constants.BUCKET_DELETE_ERROR, bucketName)
		return false, false
	}
	l.ApiLogger.Infof(constants.BUCKET_DELETED, bucketName)
	return true, false
}

func (m *MinioApiInstance) UploadFile(ctx context.Context, bucketName string, name string, file multipart.File, size int64) (*minio.UploadInfo, error) {
	if !m.ValidateIfBucketExist(ctx, bucketName) {
		return nil, errors.New(constants.BUCKET_DOESNT_EXISTS)
	}
	upload, err := m.client.PutObject(ctx, bucketName, name, file, size, minio.PutObjectOptions{})
	if err != nil {
		l.ApiLogger.Error(constants.UPLOAD_POST_FILE_ERR, err.Error())
	}
	return &upload, err
}

func (m *MinioApiInstance) DeleteFile(ctx context.Context, bucketName string, keyObject string) error {
	if !m.ValidateIfBucketExist(ctx, bucketName) {
		return errors.New(constants.BUCKET_DOESNT_EXISTS)
	}
	err := m.client.RemoveObject(ctx, bucketName, keyObject, minio.RemoveObjectOptions{})
	if err != nil {
		l.ApiLogger.Error(constants.DELETE_POST_FILE_ERR, err.Error())
	}
	return err
}

var GlobalInstance *MinioApiInstance

func StartGlobalInstance(cfg *config.AWS) (*MinioApiInstance, error) {
	var err error
	GlobalInstance, err = New(cfg)
	return GlobalInstance, err
}
