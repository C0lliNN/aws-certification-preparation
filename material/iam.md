# IAM - Indentity and Access Management
Service responsible for managing access to the AWS root account. Through this service, it is possible to create "sub-users" (called IAM user) with limited permission to AWS. It is best practice to user IAM users instead of the Root Account


## Access Management
* User Groups are used to bundle users together and assign specific permissions to them
* Policy is a document describing things that can or cannot be done. It can resource specific, service specific and so on. It used formatted using JSON and it contains a version and a at least one statement
* Identity Providers are a way to access AWS using SSO. This the standard for organizations.
* MAF - Multi-factor authentication makes an account safer requiring both the password and access to protected device to be able to access the account. People should encouraged to use it
* Role is an identity you assign some specific permissions with credentials that are valid for short durations. They are usually used by other AWS Services instead of people
