# ECS - Elastic Container Service
AWS own container Platform

## ECS Launch Type
* Launch Type is backend/infrastructure that an ECS cluster run on
* The possible types are: Fargate and EC2
* In fargate, it's not necessary to provision the infrastructure. It's all serverless

## ECS IAM Roles
* There's two types roles in the ECS Cluster: the EC2 Instance Role (for EC2 Launch Type Only) and the Task Role
* The EC2 role should be able to access ECS, ECR, ClouldWatch and maybe Secrets Manager
* The ECS task itself should be attached to the roles the application uses
* It's considered a best practice to use one role per task definition

## ECS LoadBalancer Integration
* App Load Balancer: supported and recommend for most usecases
* Network Load Balancer: recommended only for high throughput usecases
* Classic Load Balancer: supported but not recommended

## ECS Data Volumes
* It's possible to mound EFS file systems onto ECS tasks
* It works for both EC2 and Fargate launch types
* Fargate + EFS = Serverless

## ECS Service Auto Scaling
* Automatically increase/decrease the desired number of ECS tasks
* Amazon ECS auto scaling uses AWS Application Auto Scaling
* Target Tracking: scale based on target value for a specific CloudWatch metric
* Step Scaling: scale based on a specified ClouldWatch Alarm
* Scheduled Scaling: scale based on a specified date/time
* ECS Service Auto Scaling != EC2 Auto Scaling
* Fargate Auto Scaling is much easier to setup (because Serverless)

## ECS Rolling Updates
* When updating the task definition from v1 to v2, it's possible to control how many tasks can be started and stopped, and in which order.
* It's a way to deploy application without downtime

## ECS Task Definitions
* Task definitios are metadata in JSON format to tell ECS how to run a docker container
* It contains information such as:
	* Image Name
	* Port Binding for Container and Host
	* Memory and CPU required
	* Environment variables
	* Networking information
	* IAM Role
* It's possible to define up to 10 containers in a task definition. (side car)

## Amazon ECS - Data Volues
* Share data between multiple containers in the same task definition
* Works for both EC2 and Fargate tasks

## Amazon ECR - Elastic Container Registry
* It's used for storing and managing docker images on AWS
* Private and public repository
* Fully integrated with ECS, backed by Amazon S3

## Amazon EKS - Elastic Kubernetes Service
* It is a way to launch managed Kubernetes clusters on AWS
* Kubernetes is an open-source system for automatic deployment, scaling and management of containerized (usually Docker) application
* It's an alternative to ECS, similar goal but different API


