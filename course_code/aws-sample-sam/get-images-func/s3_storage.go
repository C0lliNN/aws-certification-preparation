package main

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Storage struct {
	client     *s3.S3
	bucketName string
}

func NewS3Storage(client *s3.S3, bucketName string) *S3Storage {
	return &S3Storage{client: client, bucketName: bucketName}
}

func (s *S3Storage) GetURLFromKey(ctx context.Context, key ImageKey) (string, error) {
	req, _ := s.client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(string(key)),
	})

	return req.Presign(60 * time.Hour)
}
