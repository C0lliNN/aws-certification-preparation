# VPC - Virtual Private Clould
It is a private network to deploy your resources (regional resource)

## Subnets
It is partition inside your VPC (AZ resource)

* Public subnet is a subnet that is accessible from the internet
* Private subnet is a subnet that is not accessible from the internet

## Route Tables
They are used to define acess to the internet and between subnets

## Internet Gateway
It helps the VPC instances connect with the internet. Public subnets have a route to the internet gateway

## NAT Gateways & NAT instances
It allows the instances in the private subnet to access the internet while remaining private.

## NACL
* A firewall which controls traffic from and to subnet
* Can have ALLOW and DENY rules
* Are attached at the SUbnet levels
* Rules only include IP addresses

## Security Groups
* A firewall that controls traffic to and from an ENI / an EC2 instance
* Can have only ALLOW rules
* Rules include IP addresses and other security groups

## VPC Flow Logs
* Capture information about IP traffic going into your interfaces
* Helps to monitor & troubleshoot connectivity issues
* Captures network information from AWS managed interfaces too: ELB, ElastiCache, RDS, Aurora.

## VPC Peering
* Connect two VPC, privately using AWS' network
* Make them behave as if they were in the same network
* VPC Peering connection is not transitive

## VPC Endpoints
* Provide private access to AWS Services within VPC
* Site to Site VPN and Direct Connect cannot access it

## Site to Site VPN
* Connect an on-premises VPN to AWS
* Encrypted connection
* Goes over the public internet

## Direct Connect (DX)
* Establish a physical connection between on-premises and AWS
* The connection is private, secure and fast
* Goes over a private network



