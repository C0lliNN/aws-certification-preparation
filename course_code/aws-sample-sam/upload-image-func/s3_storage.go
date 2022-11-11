package main

import (
	"bytes"
	"context"
	"image"
	"image/png"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Storage struct {
	uploader   s3manager.Uploader
	bucketName string
}

func NewS3Storage(uploader s3manager.Uploader, bucketName string) *S3Storage {
	return &S3Storage{uploader: uploader, bucketName: bucketName}
}

func (s *S3Storage) StoreImage(ctx context.Context, key ImageKey, img image.Image) error {
	buff := bytes.NewBuffer(nil)
	if err := png.Encode(buff, img); err != nil {
		return err
	}

	params := &s3manager.UploadInput{
		Bucket: &s.bucketName,
		Key:    aws.String(string(key)),
		Body:   buff,
	}

	_, err := s.uploader.Upload(params)
	return err
}
