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
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var service *UploadImageService

func init() {
	c := aws.NewConfig()
	// Support Local Testing
	c.Endpoint = aws.String(os.Getenv("AWS_ENDPOINT"))

	sess := session.Must(session.NewSession(aws.NewConfig()))

	dynamodbClient := dynamodb.New(sess)
	tableName := os.Getenv("TABLE_NAME")
	repo := NewDynamoDBRepository(dynamodbClient, tableName)

	s3Uploader := s3manager.NewUploader(sess)
	bucketName := os.Getenv("BUCKET_NAME")
	storage := NewS3Storage(*s3Uploader, bucketName)

	service = NewUploadImageService(repo, storage)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var req SaveImageRequest
	json.Unmarshal([]byte(request.Body), &req)

	image, err := service.SaveImage(ctx, req)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	response, err := json.Marshal(image)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(response),
		Headers: map[string]string{
        "Access-Control-Allow-Headers" : "Content-Type",
        "Access-Control-Allow-Origin": "*",
        "Access-Control-Allow-Methods": "POST",
    },
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.StartWithOptions(handler, lambda.WithContext(context.Background()))
}
