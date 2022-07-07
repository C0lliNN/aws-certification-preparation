# EC2 Instance Storage Options
There are couple of options available regarding how to persist files for a given ec2 instance

## EBS - Elastic Block Store
* An EBS volume is a network drive you can attach to instances while they run
* It allows your instances to persist data, even after their termination
* They usually be mounted at one instance at time, but one instance mmight have one or more volumes attached
* They are **bounded to a specific availability zone**
* It is possible to think about them like "network USB stick"
* It uses network to comunicate with the EC2 instance. Hence, some latency is usually involved
* It has a provisied capacity (GB and IOPS)
* EBS volumes are deleted in the instance termination, unless specificed otherwise

## EBS Snapshots
* Snapshots can be understood as backups. 
* It is recommended detaching the volume before creating an snapshot
* Snapshots are useful when it is need to transfer a volume between Availability zones
* We have different services for reducing costs like EBS Snapshot Archive
* Recycle Bin for EBS Snapshots is safe mechanism for preventing delete incidents


## AMI - Amazon Machine Image
* It is a customization of an EC2 instance.
* It is possible to add specific software, configuration, operation system, monitoring... 
* It allows faster book since the software is pre-packaged
* They are build for a specific region, but they can be copied
* It is like a Docker Image

## EC2 Instance Store
* EC2 Instance store, differently from EBS, is an actual hardware disk being attached to the instance.
* It offers better performance.
* However, data is lost when the instance stops

## EBS Volume Types
* gp2/gp3 SSD general purpose
* io1/io2 SSD better performance
* stl1 HDD general purpose
* scl HDD lowest performance

## EFS - Elastic File System
* Managed NFS (network file system) that can be mounted on may EC2
* It works in multi-AZ
* Highly available, scable and expensive
* its main usecases are: content management, web serving, data sharing and wordpress
* It's compatible only with Linux based AMI
