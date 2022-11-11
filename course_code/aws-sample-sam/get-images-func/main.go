package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
)

var service *GetImagesService

func init() {
	c := aws.NewConfig()
	// Support Local Testing
	c.Endpoint = aws.String(os.Getenv("AWS_ENDPOINT"))

	sess := session.Must(session.NewSession(aws.NewConfig()))

	dynamodbClient := dynamodb.New(sess)
	tableName := os.Getenv("TABLE_NAME")
	repo := NewDynamoDBRepository(dynamodbClient, tableName)

	s3Client := s3.New(sess)
	bucketName := os.Getenv("BUCKET_NAME")
	storage := NewS3Storage(s3Client, bucketName)

	service = NewGetImagesService(repo, storage)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var req GetImagesRequest

	images, err := service.GetImages(ctx, req)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	response, err := json.Marshal(images)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(response),
		Headers: map[string]string{
        "Access-Control-Allow-Headers" : "Content-Type",
        "Access-Control-Allow-Origin": "*",
        "Access-Control-Allow-Methods": "GET",
    },
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.StartWithOptions(handler, lambda.WithContext(context.Background()))
}
