# Route 53
Managed DNS WEB Service

## DNS Concepts
DNS stands for Domain Name System, and it's responsible for translating hostnames to IP addresses.

* It uses hierarchical naming structure
### Terminologies

* Domain Registrar: It is a service that manages the reservation of domain names. EX: Amazon Route 53, GoDaddy 
* DNS Records: DNS records (aka zone files) are instructions that live in authoritative DNS servers and provide information about a domain including what IP address is associated with that domain and how to handle requests for that domain.
* DNS Records Types: A, AAAA, CNAME, NS
* Zone file: contains DNS records
* Name Server (NS): resolves DNS queries (Authoritative or Non-Auhtoritative)
* Top Level Domain (TLD): .com
* Second Level Domain (SDL): amazon.com

## Amazon Route 53
* A highly available, scalable, fully managed and Authoritative DNS
* Route 53 is also a Domain Registrar
* Includes Healthchecks
* The only AWS service which provides 100% availability SLA
* The name route 53 is a reference to the default DNS port


### Records
* Records define how to route traffic for a domain
* Each record contains:
	* Domain/subdomain Name
	* Record Type - A or AAAA
	* Value - IP Address or other hostname
	* Routing Policy
	* TTL
* Route 53 Suports the following DNS record types: A, AAA, CNAME, NS
* Record A: Maps a hostname to IPv4
* Record AAAA: Maps a hostname to IPv6
* CNAME: maps a hostname to another hostname
* NS: Control how traffic is routed for a domain

### Hosted Zones
* A contaienr for record that define how to route traffic to a domain and its subdomains
* Public Hosted Zones: contains records that specify how to route traffic on the Internet
* Private Hosted Zones: contain record that specify how you route traffic within one or more VPC

### CNAME vs Alias
* AWS resources expose an AWS hostname
* CNAME only works for NON ROOT DOMAIN
* Alias points a hostname to an AWS resource. It works for ROOT Domain
* Alias is free of chare for traffic
* Alias has native health check
* Alias don't allow to configure the TTL
* Alias cannot point to an EC2 DNS name

### Routing Policies
* Define how Route 53 responds to DNS queries
* Don't get confused by the word "Routing"
* Route 53 Supports the following Routing Policies:
	* Simple
	* Weighted
	* Failover
	* Latency based
	* Geolocation
	* Multi-Value Answer
	* Geoproximity

### Route 53 - Heath Checks
* Only fro public resources
* Health Check => Automated DNS failover
* Health checks monitor an enpoint, other health checks or even CloudWatch Alerts




