package main

import (
	"c0llinn/aws-sqs-consumer/internal/consumer"
	"c0llinn/aws-sqs-consumer/internal/handler"
	"context"
	"flag"
	"log"
)

var queue = flag.String("queue", "", "The name of the queue")
var endpoint = flag.String("endpoint-url", "", "The AWS endpoint")
var region = flag.String("region", "", "The AWS region")
var waitTime = flag.Int64("wait-time", 2, "How long, in seconds, to wait for long polling")

func main() {
	flag.Parse()

	handler := handler.NewHandler(handler.Config{})

	c, err := consumer.NewConsumer(consumer.Config{
		Region: *region,
		QueueName: *queue,
		Endpoint:   *endpoint,
		WorkerPool: 30,
		WaitTime:   *waitTime,
		Handler:    handler,
	})

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	c.Consume(ctx)
}