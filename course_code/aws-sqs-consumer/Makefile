
start:
	AWS_ACCESS_KEY_ID=test AWS_SECRET_ACCESS_KEY=test AWS_DEFAULT_REGION=us-east-1 \
	go run cmd/main.go --queue sqs-sample-queue --endpoint-url http://localhost:4566 --wait-time 2 --region us-east-1

build-image:
	docker image build . -t aws-sqs-consumer

push-image:
	aws ecr get-login-password --region us-east-1 --profile prep-exam | docker login --username AWS --password-stdin 890416256469.dkr.ecr.us-east-1.amazonaws.com
	docker tag aws-sqs-consumer:latest 890416256469.dkr.ecr.us-east-1.amazonaws.com/sqs-consumer:latest
	docker push 890416256469.dkr.ecr.us-east-1.amazonaws.com/sqs-consumer:latest

create-stack:
	aws cloudformation --region=us-east-1 --profile=prep-exam create-stack --template-body file://template.yml \
 		--stack-name sqs-consumer --capabilities CAPABILITY_NAMED_IAM \
 		--parameters ParameterKey=Image,ParameterValue=890416256469.dkr.ecr.us-east-1.amazonaws.com/sqs-consumer:latest \
								 ParameterKey=Subnet,ParameterValue=subnet-00e3368f4357c4184

delete-stack:
	aws cloudformation --region=us-east-1 --profile=exam delete-stack --stack-name sqs-consumer
	
