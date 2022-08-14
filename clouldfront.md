# CloudFront
CDN service in AWS

## Features
* 215 Point of Presence Globally
* DDos protection
* Can expose external HTTPS and can talk to internal HTTPS backend
* Can be used in front of a S3 Website, EC2, or ALB

## Origin Access Identity (OAI)
It's a feature used to restrict access to an S3 bucket to only the ClouldFront Distribution

## CloudFront Geo Restriction
It's possible to deny/allow access based on the Country

## CloudFront vs S3 Cross Region Replication
ClouldFront for static content that must be **available** everywhere, whereas S3 CRR is great for **dynamic** content accross a few regions.

## CloudFront Caching
* Cache based on Headers, Session Cookies and Query String Parameters
* The cache lives at each CloudFront Edge Location
* You can invalidate part of the cache using the CreateInvalidation API

## Signed URLs / Cookies URL
* If the requirement is to distribute paid shared content to premium users over the world, Signed URL or Cookie can be used.
* Signed URL = access to individual files
* Signed Cookies = access to multiple files

## CloudFront Signed URL vs S3 Pre-Signed URL
* S3 URLs have a limited lifetime
* S3 uses the IAM key of the signing 
* CF URls are more flexible and allow filtering by IP, path, date and expiration
* CF can leverage caching features

## CloudFront - Multiple Origin
* It's possible to route to differnt kind of origing based on the content type or path

## Clould - Origin Groups
* To increase high-availability and do failover
* Origin Group: one primary and secondary origin
* If the primary origin fails, the second is used.

