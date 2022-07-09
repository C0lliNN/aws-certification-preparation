# ASG - Auto Scaling Group
The goal of ASG is to sacle out to match an increased load, and scale in to match a decreased load.

ASG are free (you only pay for the underlying EC2 instances)

## Important Terminology

* Minimum capacity: the minimum amount of instances regardless of the load
* Desired capacity: the desired amount of instances necessary to handle the load
* Maximum capacity: the maximum amount of instances regardless of the load

## ASG Attributes
* Launch Template: defines the configuration of the new instances (including AMI, EC2 data, EBS volumes and so on)
* Min Size / Max Size / Initial Capacity
* Scaling Policies

## ASG and ClouldWatch
* It is possible to scale an ASG based on CloulWatch Alarms
* Based on the alarm
	* It's possible to create scale-out policies
	* It's possible to create scale-in policies

## ASG Dynamic Scaling Policies
These are ways of modifying the number of instancers without manual intervention
* Target Tracking Scaling
	* Most simple and easy to set-up
	* I want the average ASG CPU to stay at around 40%
* Step Scaling
	* When a CloudWatch alarm is triggered (example CPU > 70%), then add 2 units
	* When a CloudWatch alarm is triggered (example CPU < 30%), then remove 1
* Scheduled Actions
	* Antecipate a scaling based on known usage patters

