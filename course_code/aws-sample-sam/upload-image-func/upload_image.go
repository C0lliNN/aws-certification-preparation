package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"strings"
	"time"

	"github.com/google/uuid"
)

type ImageKey string

type Image struct {
	ID      string   `dynamodbav:"id"`
	Key     ImageKey `dynamodbav:"key"`
	Created int64    `dynamodbav:"created"`
}

type Repository interface {
	SaveImage(ctx context.Context, image *Image) error
}

type Storage interface {
	StoreImage(ctx context.Context, key ImageKey, image image.Image) error
}

type UploadImageService struct {
	Repository Repository
	Storage    Storage
}

func NewUploadImageService(repo Repository, storage Storage) *UploadImageService {
	return &UploadImageService{repo, storage}
}

type SaveImageRequest struct {
	// ImageContent base64 encoded
	ImageContent string
	ImageName    string
}

func (s *UploadImageService) SaveImage(ctx context.Context, req SaveImageRequest) (*Image, error) {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(req.ImageContent))

	img, format, err := image.Decode(reader)
	if err != nil {
		return nil, fmt.Errorf("cannot decode file: %v", err)
	}

	currentTime := time.Now().Unix()
	key := ImageKey(fmt.Sprintf("%s_%d.%s", req.ImageName, currentTime, format))

	if err = s.Storage.StoreImage(ctx, key, img); err != nil {
		return nil, err
	}

	image := &Image{
		ID:      uuid.New().String(),
		Key:     key,
		Created: currentTime,
	}

	if err = s.Repository.SaveImage(ctx, image); err != nil {
		return nil, err
	}

	return image, nil
}
