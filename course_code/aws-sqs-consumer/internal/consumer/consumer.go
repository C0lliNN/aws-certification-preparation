package consumer

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var maxMessages = int64(10)

type Handler interface {
	Handle(ctx context.Context, message string) error
}

type Config struct {
	Region string
	Endpoint          string
	QueueName         string
	WorkerPool        int
	VisibilityTimeout int
	WaitTime          int64
	Handler           Handler

	queueUrl string
	sqs      *sqs.SQS
}

type Consumer struct {
	config Config
}

func NewConsumer(c Config) (*Consumer, error) {
	if c.QueueName == "" {
		return nil, fmt.Errorf("QueueName is required")
	}

	if c.Region == "" {
		return nil, fmt.Errorf("Region is required")
	}

	svc := sqs.New(newSession(c))
	c.sqs = svc

	output, err := c.sqs.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &c.QueueName,
	})
	if err != nil {
		return nil, err
	}

	c.queueUrl = *output.QueueUrl

	if c.WorkerPool == 0 {
		c.WorkerPool = 10
	}

	if c.VisibilityTimeout == 0 {
		c.VisibilityTimeout = 10
	}

	if c.WaitTime == 0 {
		c.WaitTime = 2
	}

	return &Consumer{c}, nil
}

func (c *Consumer) Consume(ctx context.Context) {
	jobs := make(chan *sqs.Message)

	for w := 1; w <= c.config.WorkerPool; w++ {
		go c.worker(ctx, w, jobs)
	}

	for {
		output, err := c.config.sqs.ReceiveMessageWithContext(ctx, &sqs.ReceiveMessageInput{
			QueueUrl:            &c.config.queueUrl,
			MaxNumberOfMessages: &maxMessages,
			WaitTimeSeconds:     &c.config.WaitTime,
		})
		if err != nil {
			log.Printf("%s , retrying in 10s", err)
			time.Sleep(10 * time.Second)
			continue
		}

		for _, m := range output.Messages {
			jobs <- m
		}
	}
}

func (c *Consumer) worker(ctx context.Context, id int, messages <-chan *sqs.Message) {
	for m := range messages {
		if err := c.run(ctx, m); err != nil {
			log.Println(err.Error())
		}
	}
}

func (c *Consumer) run(ctx context.Context, message *sqs.Message) error {
	if err := c.config.Handler.Handle(ctx, *message.Body); err != nil {
		return err
	}

	return c.delete(ctx, message.ReceiptHandle)
}

func (c *Consumer) delete(ctx context.Context, receiptHandle *string) error {
	_, err := c.config.sqs.DeleteMessageWithContext(ctx, &sqs.DeleteMessageInput{
		QueueUrl: &c.config.queueUrl, 
		ReceiptHandle: receiptHandle,
	})

	return err
}

func newSession(c Config) *session.Session {
	var endpoint *string
	if c.Endpoint != "" {
		endpoint = &c.Endpoint
	}

	return session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: &c.Region, Endpoint: endpoint},
	}))
}
