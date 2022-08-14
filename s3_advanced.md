# Advanced S3

## S3 MFA Delete
It's possible to require an MFA confirmation when deleting objects or disabling versioning. This can be useful for increasing security

## S3 Default Encryption
It's possible to enforce that all objects in a bucket are encrypted either using policies or enabling S3 default Encryption.

## S3 Access Logs
In order to record all actions performed on an s3 bucket by users, S3 Access Logs can be enabled. It's important to redirect the logs to a different bucket

## S3 Replication
When replication is enabled, new writes are replicated to another bucket in the same region or in another region

## S3 Pre-signed urls
Pre-signed urls give temporary read/write access to a location is S3

## Storage Classes
* Amazon S3 Standard - General Purpose
* Amazon S3 Standard - Infrequent Access (IA)
* Amazon S3 One Zone - Infrequent Access
* Amazon S3 Glacier Instant Retrieval
* Amazon S3 Glacier Flexible Retrieval
* Amazon S3 Glacier Deep Archive
* Amazon S3 Intelligent Tiering

## S3 Lifecycle Rules
Lifecycle rules are an automated mechanism of changing the storage classes or even deleting objects based on the last access.

## S3 Performance
* It's important to note that throughput is based on the prefix
* KMS might limit performance
* Multi-part update can faster the writes
* S3 Transfer Acceleration
* S3 Byte-Range Fetches
* S3 Select - Server Side Filtering

## Amazon Athena
* Serverless quey service to perform analytics agains S3 objects
* Uses standard SQL language to query the files

