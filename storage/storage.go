package storage

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
)

var (
	MINIO_ENDPOINT   = "MINIO_ENDPOINT"
	MINIO_ACCESS_KEY = "MINIO_ACCESS_KEY"
	MINIO_SECRET_KEY = "MINIO_SECRET_KEY"
	MINIO_BUCKET     = "MINIO_BUCKET"
)

type Storage struct {
	MinioClient *minio.Client
}

type UploadPayload struct {
	Folder      string
	File        io.Reader
	Filename    string
	ContentType string
	Size        int64
}

func NewStorage(log *logrus.Entry) (*Storage, error) {
	endpoint := os.Getenv(MINIO_ENDPOINT)
	accessKeyID := os.Getenv(MINIO_ACCESS_KEY)
	secretAccessKey := os.Getenv(MINIO_SECRET_KEY)

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: true,
	})
	if err != nil {
		log.Fatalln("Failed to initialize storage minio, err: ", err)
		return nil, err
	}

	storage := &Storage{
		MinioClient: minioClient,
	}

	return storage, nil
}

func (s *Storage) UploadFile(ctx context.Context, payload UploadPayload) error {
	filePath := fmt.Sprintf("%s/%s", payload.Folder, payload.Filename)

	_, err := s.MinioClient.PutObject(
		ctx,
		os.Getenv(MINIO_BUCKET),
		filePath,
		payload.File,
		payload.Size,
		minio.PutObjectOptions{
			ContentType: payload.ContentType,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) DeleteFile(ctx context.Context, folder, filename string) error {
	filePath := fmt.Sprintf("%s/%s", folder, filename)

	err := s.MinioClient.RemoveObject(
		ctx,
		os.Getenv(MINIO_BUCKET),
		filePath,
		minio.RemoveObjectOptions{},
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) IsFileExists(ctx context.Context, folder, filename string) bool {
	filePath := fmt.Sprintf("%s/%s", folder, filename)

	_, err := s.MinioClient.StatObject(ctx, os.Getenv(MINIO_BUCKET), filePath, minio.StatObjectOptions{})
	if err != nil {
		respErr, ok := err.(minio.ErrorResponse)
		if ok && respErr.Code == "NoSuchKey" {
			return false
		}
		return false
	}
	return true
}

func (s *Storage) CopyFile(ctx context.Context, folder, filename, newFilename string) error {
	bucket := os.Getenv(MINIO_BUCKET)

	srcPath := fmt.Sprintf("%s/%s", folder, filename)
	destPath := fmt.Sprintf("%s/%s", folder, newFilename)

	src := minio.CopySrcOptions{
		Bucket: bucket,
		Object: srcPath,
	}

	dst := minio.CopyDestOptions{
		Bucket: bucket,
		Object: destPath,
	}

	_, err := s.MinioClient.CopyObject(ctx, dst, src)
	return err
}
