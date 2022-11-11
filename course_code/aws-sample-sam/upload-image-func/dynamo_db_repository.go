package main

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DynamoDBRepository struct {
	client    *dynamodb.DynamoDB
	tableName string
}

func NewDynamoDBRepository(client *dynamodb.DynamoDB, tableName string) *DynamoDBRepository {
	return &DynamoDBRepository{client: client, tableName: tableName}
}

func (r *DynamoDBRepository) SaveImage(ctx context.Context, image *Image) error {
	item, err := dynamodbattribute.MarshalMap(image)
	if err != nil {
		return nil
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(r.tableName),
	}

	_, err = r.client.PutItem(input)
	return err
}
