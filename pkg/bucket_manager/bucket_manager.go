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

func New(ctx *context.Context, cfg config.AWS) (*MinioApiInstance, error) {
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
		ctx:    ctx,
		client: client,
	}, nil
}

type MinioApiInstance struct {
	ctx    *context.Context
	client *minio.Client
}

func (m *MinioApiInstance) ValidateIfBucketExist(bucketName string) (sucess bool) {
	exists, errBucketExists := m.client.BucketExists(*m.ctx, bucketName)
	if errBucketExists != nil {
		l.ApiLogger.Error(constants.BUCKET_CREATION_ERROR, errBucketExists.Error())
		return false
	}
	l.ApiLogger.Infof(constants.BUCKET_ALREADY_EXISTS, bucketName)
	return exists
}

func (m *MinioApiInstance) CreateBucket(bucketName string, location string) (sucess bool) {
	err := m.client.MakeBucket(*m.ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		return m.ValidateIfBucketExist(bucketName)
	} else {
		l.ApiLogger.Info(constants.BUCKET_CREATED)
		return true
	}
}

func (m *MinioApiInstance) DeleteBucket(bucketName string) (sucess bool, notFound bool) {
	if !m.ValidateIfBucketExist(bucketName) {
		return false, true
	}
	err := m.client.RemoveBucket(*m.ctx, bucketName)
	if err != nil {
		l.ApiLogger.Infof(constants.BUCKET_DELETE_ERROR, bucketName)
		return false, false
	}
	l.ApiLogger.Infof(constants.BUCKET_DELETED, bucketName)
	return true, false
}

func (m *MinioApiInstance) UploadFile(bucketName string, name string, file multipart.File, size int64) (*minio.UploadInfo, error) {
	if !m.ValidateIfBucketExist(bucketName) {
		return nil, errors.New(constants.BUCKET_DOESNT_EXISTS)
	}
	upload, err := m.client.PutObject(*m.ctx, bucketName, name, file, size, minio.PutObjectOptions{})
	if err != nil {
		l.ApiLogger.Error(constants.UPLOAD_POST_FILE_ERR, err.Error())
	}
	return &upload, err
}

func (m *MinioApiInstance) DeleteFile(bucketName string, keyObject string) error {
	if !m.ValidateIfBucketExist(bucketName) {
		return errors.New(constants.BUCKET_DOESNT_EXISTS)
	}
	err := m.client.RemoveObject(*m.ctx, bucketName, keyObject, minio.RemoveObjectOptions{})
	if err != nil {
		l.ApiLogger.Error(constants.DELETE_POST_FILE_ERR, err.Error())
	}
	return err
}

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

func UploadFile(ctx context.Context, bucketName string, name string, file multipart.File, size int64) (minio.UploadInfo, error) {
	if MinioClient == nil {
		l.ApiLogger.Errorf(constants.BUCKET_MANAGER_NOT_STARTED, bucketName)
		return minio.UploadInfo{}, errors.New(constants.BUCKET_MANAGER_NOT_STARTED)
	}
	if !validateIfBucketExist(ctx, bucketName) {
		return minio.UploadInfo{}, errors.New(constants.BUCKET_DOESNT_EXISTS)
	}
	upload, err := MinioClient.PutObject(ctx, bucketName, name, file, size, minio.PutObjectOptions{})

	if err != nil {
		l.ApiLogger.Error(constants.UPLOAD_POST_FILE_ERR, err.Error())
	}

	return upload, err
}

func DeleteFile(ctx context.Context, bucketName string, keyObject string) error {
	if MinioClient == nil {
		l.ApiLogger.Errorf(constants.BUCKET_MANAGER_NOT_STARTED, bucketName)
		return errors.New(constants.BUCKET_MANAGER_NOT_STARTED)
	}
	if !validateIfBucketExist(ctx, bucketName) {
		return errors.New(constants.BUCKET_DOESNT_EXISTS)
	}
	err := MinioClient.RemoveObject(ctx, bucketName, keyObject, minio.RemoveObjectOptions{})

	if err != nil {
		l.ApiLogger.Error(constants.DELETE_POST_FILE_ERR, err.Error())
	}

	return err
}

func validateIfBucketExist(ctx context.Context, bucketName string) bool {
	exists, errBucketExists := MinioClient.BucketExists(ctx, bucketName)
	if errBucketExists != nil {
		l.ApiLogger.Error(constants.BUCKET_CREATION_ERROR, errBucketExists.Error())
		return false
	}
	l.ApiLogger.Infof(constants.BUCKET_ALREADY_EXISTS, bucketName)
	return exists
}
