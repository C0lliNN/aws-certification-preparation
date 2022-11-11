# ELB - Elastic Load Balancer
Load Balancers are servers that forward traffic to multiple downstream services (like EC2, and ECS)

## Scalability
Scalability means that an application can handle greater loads by adapting. There two kinds of scalability: vertical and horizontal

## High Availability
It usually goes hand in hand with horizontal scaling. Its goal is to survive a data center loss.

## Why Load Balancers
1. Spread load across multiple downstream instances
2. Expose a single point of access (DNS) to your application
3. Handle failures of downstream instances
4. Provide SSL Termination for websites

## Why ELB
* An Elastic Load Balancer is **managed load balancer**
	* AWS garantees that it will be working
	* AWS takes cares of the upgrades, maintenance, high availability
* It costs less to setup your own load balancer but it will be a lot more effort on your end
* It is integrated with many AWS offerings
	* EC2, EC2 Auto Scaling Groups, Amazon ECS
	* AWS Certificate Manager, CloudWatch
	* Route 53, AWS WAF, AWS Global Accelerator

## Types of Load Balancers
1. Classic Load Balancer (HTTP, HTTPS, TCP, SSL). It is deprecated
2. Application Load Balancer - HTTP, HTTPS. It provides a lot of cool features
3. Network Load Balancer - It acts in the transport layer. It has great performance
4. Gateway Load Balancer - It operates in Network layer

## Target Groups
It is possible to think about target groups as the a set of servers of the same application. They are useful when you wanna use the same load balancer for multiple applications.

## SNI - Server Name Indication
* It solves the problem of loading multiple SSL certificates onto one web server
* It is a newer protocol, and requires the client to indicate the hostname of the target server in the initial SSL handshake
* It only works for ALB & NLB

## Connection Draining
* Type to complete in-flight requests while the instance is de-registering or unhealthy
* Stops sending new requests to the EC2 instance which is de-registering
