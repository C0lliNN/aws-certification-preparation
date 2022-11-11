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

func (r *DynamoDBRepository) FindAll(ctx context.Context) ([]Image, error) {

	output, err := r.client.ScanWithContext(ctx, &dynamodb.ScanInput{TableName: aws.String(r.tableName)})
	if err != nil {
		return nil, err
	}

	items := make([]Image, 0)

	for _, i := range output.Items {
		var image Image
		if err = dynamodbattribute.UnmarshalMap(i, &image); err != nil {
			return nil, err
		}

		items = append(items, image)
	}

	return items, nil
}
