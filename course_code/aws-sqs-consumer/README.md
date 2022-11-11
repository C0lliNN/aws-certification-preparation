# AWS SQS Consumer
A sample Golang AWS consumer

## How to run locally
* Run `docker-compose up -d`
* Run `make start`

## Useful commands
* Build Image:
```bash
docker image build . -t aws-sqs-consumer
```

* Run Container:
```bash
docker container run --name aws-sqs-consumer-test2 --network host -e AWS_ACCESS_KEY_ID=test -e AWS_SECRET_ACCESS_KEY=test aws-sqs-consumer --queue sqs-sample-queue --endpoint-url http://localhost:4566 --wait-time 2 --region us-east-1
```

## Deployment Steps
* Build Image
* Push Image to some Image Repository
* Create the Cloudformation steps providing the image uri and the subnet