# Build docker image
docker image build . -t aws-sample-nodejs --build-arg PORT=80

docker tag aws-sample-nodejs:latest 890416256469.dkr.ecr.us-east-1.amazonaws.com/aws-sample-nodejs:latest

docker push 890416256469.dkr.ecr.us-east-1.amazonaws.com/aws-sample-nodejs:latest