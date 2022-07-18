# RDS - Relational Database Service
A managed relational DB service in AWS

## Available SQL Engines
* Postgres
* MySQL
* MariaDB
* Oracle
* Microsoft SQL Server

## Advantages over deploying a DB on EC2
* Automated provisioning
* OS Patching
* Backups
* Scaling capability
* Maintenance windows

## Disadvantage over deploying a DB on EC2
* It's not possible SSH

## RDS Backups
* Automatically enabled
* Daily full backup of the database
* Ability to restore to any point in time
* 7 days retention

## DB Snapshots
* Manually triggered by the user
* Retention of backup for as long as necessary

## RDS - Storage Auto Scaling
* It's possible to Maximum Storage Threshold
* Avoid manually scaling your database storage
* Useful for applications with unpredictable workloads

## Read Replicas for read scalability
* Up to 5 Replicas
* Withing AZ, Cross AZ or Cross Region
* Replication is ASYNC
* Applications must update the connection string to leverage read replicas
* Useful for reporting
* For RDS Read replicas withing the same region, the network cost is free

## RDS Multi AZ (Disaster Recovery)
* SYNC Replication
* One DNS name
* Increase availability
* Not used for scaling
* Note: The Read Replicas can be setup as Multi AZ for Disaster Recovery
* Migrating from Single-AZ to Multi-AZ is a zero downtime operation

## RDS Encryption

### At rest encryption
* Possible to encrypt using AWS KMS - AES-256
* Encryption must be defined at lunch time
* If the master is not encrypted, the read replicas cannot be encrypted

### In-flight encryption
* SSL certificates to encrypt data to RDS in flight
* It's possible to enforce SSL

## RDS Security
* RDS databases are usually deployed withing a private subnet, not in a public one
* RDS security works be leveraging security groups
* IAM Policies can be used as well
* IAM database authentication works with MySQL and PostgreSQL


## Amazon Aurora
* Aurora is proprietaru technology from AWS
* Prostgres and MySQL are both supported as Aurora DB
* Aurora is "AWS cloud optimized" and claims 5x performance improvement
* Aurora storage automatically grows in increments of 10GB, up to 128TB
* Aurora can have 15 replicas while MySQL has 5, and the replication process is faster
* Aurora costs more than RDS
* 6 Copies of the data across 3 AZ
	* 4 copies out of 6 needed for writes
	* 3 copies out of 6 needed for reads








