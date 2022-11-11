package main

import (
	"context"
)

type ImageKey string

type Image struct {
	ID      string   `dynamodbav:"id"`
	Key     ImageKey `dynamodbav:"key"`
	Created int64    `dynamodbav:"created"`
	URL     string   `dynamodbav:"-"`
}

type Repository interface {
	FindAll(ctx context.Context) ([]Image, error)
}

type Storage interface {
	GetURLFromKey(ctx context.Context, key ImageKey) (string, error)
}

type GetImagesService struct {
	Repository Repository
	Storage    Storage
}

func NewGetImagesService(repo Repository, storage Storage) *GetImagesService {
	return &GetImagesService{repo, storage}
}

type GetImagesRequest struct{}

func (s *GetImagesService) GetImages(ctx context.Context, req GetImagesRequest) ([]Image, error) {
	images, err := s.Repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(images); i++ {
		url, err := s.Storage.GetURLFromKey(ctx, images[i].Key)
		if err != nil {
			return nil, err
		}
		images[i].URL = url
	}

	return images, nil
}
