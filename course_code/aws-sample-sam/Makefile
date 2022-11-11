.PHONY: build deploy

build:
	@sam build
	@cd web && npm run build

deploy:
	@sam deploy --stack-name sample-sam --region us-east-1 \
		--resolve-s3 --region us-east-1 --capabilities CAPABILITY_IAM --profile exam

	@aws s3 sync ./web/build s3://collin-45342d-web-bucket/ --profile exam

delete:
	@sam delete --stack-name sample-sam --profile exam