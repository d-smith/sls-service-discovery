# sls-service-discovery

This project demonstrates how to use route 53 service discovery in a VPC. The example will show how lambda function can look up a service address at runtime using DNS

## Network Set up

### vpc

Private route 53 service namespaces can be set up in a VPC. For the purposes of this sample, create a VPC if needed using the included private stack.

```console
aws cloudformation create-stack --stack-name vpc --template-body file://vpc.yml
```

### private namespace and service creation

Use the `r53-service-notebook.ipynb` Jupyter notebook to step through the create of your private zone and service CNAME. Note for this example we're simulating the lookup of something outside our domain. This will be expanded later to lookup local names using different record types.

### record lookup

Once you've create the namespace and service record, you can look it up. Since it was created in a VPC, you have to be in the VPC for the dns lookup to work.

So, launch an ec2 instance in a public subnet in your VPC. SSH into the instance, and perform a lookup.

```console
[ec2-user@ip-172-31-0-222 ~]$ nslookup examplesvc.sector7.internal
Server:		172.31.0.2
Address:	172.31.0.2#53

Non-authoritative answer:
examplesvc.sector7.internal	canonical name = www.example.com.
Name:	www.example.com
Address: 93.184.216.34
```

## Serverless Example

TODO

## Follow On Work

* Demonstrate with internal services that register
* Show how health checks work, responses based on weighting, etc