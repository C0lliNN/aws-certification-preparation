.PHONY : deploy

create-bucket:
	aws s3api create-bucket --bucket rctm-sample-lambda1 --profile exam
	aws s3api put-bucket-versioning --bucket rctm-sample-lambda1 --versioning-configuration Status=Enabled --profile exam

upload-code:
	@npm i > /dev/null 2>&1
	@-zip -r app2.zip . > /dev/null 2>&1
	@aws s3api put-object --bucket rctm-sample-lambda1  --key code/app2.zip --body app2.zip --profile exam

create-stack:
	aws cloudformation --region=us-east-1 --profile=prep-exam create-stack --template-body file://template.yml \
 		--stack-name lambda-sample2 --capabilities CAPABILITY_NAMED_IAM \
 		--parameters ParameterKey=Bucket,ParameterValue=rctm-sample-lambda1  \
								 ParameterKey=S3Key,ParameterValue=code/app2.zip \
								 ParameterKey=S3ObjectVersion,ParameterValue=.6ESGqyjlKTvK79BlFBV2x3.QHgf7ned
update-stack:
	aws cloudformation --region=us-east-1 --profile=prep-exam update-stack --template-body file://template.yml \
 		--stack-name lambda-sample2 --capabilities CAPABILITY_NAMED_IAM \
 		--parameters ParameterKey=Bucket,ParameterValue=rctm-sample-lambda1  \
								 ParameterKey=S3Key,ParameterValue=code/app2.zip \
								 ParameterKey=S3ObjectVersion,ParameterValue=$(version)

delete-stack:
	aws cloudformation --region=us-east-1 --profile=exam delete-stack --stack-name lambda-sample1