# AWS CLI, SDK, IAM Roles & Policies

## AWS Policy Generator
It is a visual tool for generating AWS policies like S3 Policies, SQS Policies and IAM Policies

## AWS Policy Simulator Simulator
It is a visual tool for verifying if a given set of policies are allowed to perform some action.

## AWS CLI Dry Run
--dry-run is CLI flag that can be passed to a lot of aws commands. Its goal is to validate that the action can be performed without doing it.

## AWS CLI STS Decode
When an authorization error is returned, the error is encoded making it difficult to debug the problem. Because of that, it's possible to use the STS Decode CLI to decode the message

## AWS EC2 Instance Metadata
When logged into an EC2 instance, it's possible to get metadata information about it using the internal endpoint http://169.254.169.254/

## AWS CLI Profiles
Profiles are way to use multiple accounts in the same PC

## AWS CLI with MFA
In order to use the CLI in an account that has MFA enabled, it's necessary to make a call to sts GetSessionToken

## Exponential Backoff
It's good practice to use some retry mechanism with a backoff stratety in order to increase the application resilience against intermittent failures

## AWS Credentials Provider & Chain
The order in which the CLI get the credentials is: --profile, ENV variables, ~/.aws/crentials
The order in which the SDK get the credentials is: java system properties, ENV variables, EC2 Instance role

## AWS Signature v4 Signing
Most requests to the AWS Web API are signed using the access key id. Hence, when neither the SDK nor the CLI are being used, it's necessary to implement signing.

