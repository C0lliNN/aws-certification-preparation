docker image build . -t aws-sample-reactjs --build-arg BASE_URL=http://3.83.246.68

docker tag aws-sample-reactjs:latest 890416256469.dkr.ecr.us-east-1.amazonaws.com/aws-sample-reactjs:latest

docker push 890416256469.dkr.ecr.us-east-1.amazonaws.com/aws-sample-reactjs:latest