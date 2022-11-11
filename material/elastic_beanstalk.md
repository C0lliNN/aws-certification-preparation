# Elastic Beanstalk
It's a service that provides a developer centric view of deploying an application on AWS. It's like an abstraction that makes easy to deploy software. It keeps the developer away from the low level configuration details.

* It uses components like EC2, ASG, ELB, RDS, etc...
* It's a managed service
	* Handles capacity provisioning, load balancing, scaling, etc..
	* Just the application code is the responsibility of the developer
* Free to use, but it's necessary to pay for the underlying services

## Components
* Application: collection of Elastic Beanstalk components (environment, version, configurations)
* Application Version: an iteration of your application code
* Environment
	* Collection of AWS resources running an application version
	* Tiers: Web Server Environment Tier & Worker Environment Tier

## Deployment Modes
* All at once: deploys everything in one go. It's fastest but there is some downtime.
* Rolling: Update a few instances at time, and them move onto the next bucket. It reduces the capacity in the deployment but no downtime is present. It's the cheapest way to deploy without downtime.
* Rolling with additional batches. Like rolling, but spins up new instances to move the batch aiming to not decrease the capacity
* Immutable: spins up new instances in a new ASG, deploys version to these instances. It's useful when it's necessary to rollback
* Blue / Gree: Not a direct feature of Elastic Beanstalk. Zero downtime and easy to rollback if the new release is not working as expected
* Canary Testing: New application version is deployed to a temporary ASG with the same capacity. A small % of traffic is sent to the temporary ASG for a configurable amount of time

## CLI
* It's possible to use the Elastic Beanstalk CLI in automated deployment pipelines

## Deployment Process
* Console: upload zip file and then deploy
* CLI: create new app version using CLI and then deploy

## Lifecycle Policy
* Beanstalk can store at most 1000 application versions
* To pahse out old application versions, it's possible to use a lifecycle policy

## Extensions
* Configurations for the Elastic Beanstalk can be set through the UI using a file stored in .ebextensions/something.config

## Cloning
It's a feature that allows to create a new environment of the exact configuration of an existing one

## ELB
* It's important to remember that it's not possible to change the ELB type after the EB environment is created

## RDS
* It's possible to attach a RDS database to Elastic Beanstalk. This is good for dev environments, but not so great for production.

## Docker
It's possible to run docker containers in Beanstalk. Single Docker Container or Multi Docker Containers can be executed. OBS: ECS is only used for Multi Docker Container

