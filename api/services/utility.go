package services

import (
	"context"
	"io"
	"mime/multipart"
	"net/url"
	"strings"
	"time"
	"travel/infrastructure"

	"cloud.google.com/go/storage"
)

// StorageBucketService -> handles the file upload/download function
type StorageBucketService struct {
	logger infrastructure.Logger
	client *storage.Client
	env    infrastructure.Env
}

// NewStorageBucketService -> initiate for the StorageBucketService struct
func NewStorageBucketService(
	logger infrastructure.Logger,
	client *storage.Client,
	env infrastructure.Env,
) StorageBucketService {
	return StorageBucketService{
		logger: logger,
		client: client,
		env:    env,
	}
}

// UploadFile -> uploads the file to cloud storage
func (s StorageBucketService) UploadFile(
	ctx context.Context,
	file multipart.File,
	fileName string,
	fileType string,
) (string, error) {
	var bucketName = s.env.StorageBucketName
	if bucketName == "" {
		s.logger.Zap.Fatal("Please check bucket name in env file")
	}
	_, err := s.client.Bucket(bucketName).Attrs(ctx)
	if err == storage.ErrBucketNotExist {
		s.logger.Zap.Fatal("provided bucket %v doesn't exists", bucketName)
	}
	if err != nil {
		s.logger.Zap.Fatal("Cloud bucker error: %v does not exists", bucketName)
	}
	wc := s.client.Bucket(bucketName).Object(fileName).NewWriter(ctx)
	if fileType != "application/pdf" {
		wc.ContentType = "application/octet-stream"
	}
	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}
	u, err := url.ParseRequestURI("/" + bucketName + "/" + wc.Attrs().Name)
	if err != nil {
		return "", err
	}
	path := u.EscapedPath()
	path = strings.Replace(path, "/"+bucketName, "", 1)
	path = strings.Replace(path, "/", "", 1)
	return path, nil
}

// RemoveObject -> removes the file from the storage bucket
func (s StorageBucketService) RemoveObject(object string) error {
	var bucket = s.env.StorageBucketName
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		s.logger.Zap.Fatal("storage.NewClient: %v", err)
	}
	defer client.Close()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	o := client.Bucket(bucket).Object(object)
	if err := o.Delete(ctx); err != nil {
		s.logger.Zap.Fatal("Object(%q).Delete: %v", object, err)
	}
	s.logger.Zap.Fatal("Blob %v deleted.\n", object)
	return nil
}
