aws cloudformation --region=us-east-1 --profile=exam create-stack --template-body file://template.yml \
 --stack-name golang-sample --capabilities CAPABILITY_NAMED_IAM \
 --parameters ParameterKey=UserData,ParameterValue=(base64 -w0 ec2_script)  ParameterKey=KeyName,ParameterValue=exam

aws cloudformation --region=us-east-1 --profile=exam list-stacks

aws cloudformation --region=us-east-1 --profile=exam delete-stack --stack-name golang-sample
