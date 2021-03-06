# sls-service-discovery

This project demonstrates how to use Route 53 service discovery in a VPC. The example shows how a lambda function can look up a service address at runtime using a DNS CNAME record lookup.

For more details on Route 53 Auto Naming API for service name management and discovery, refer to the [annoucement](https://aws.amazon.com/about-aws/whats-new/2017/12/amazon-route-53-releases-auto-naming-api-name-service-management/) and the [api documentation](https://docs.aws.amazon.com/Route53/latest/APIReference/API_Operations_Amazon_Route_53_Auto_Naming.html)

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

The provided serverless example does a lookup of the destination address via the CNAME associated with examplesvc.sector7.internal, does a get on it, then relays the result back to the caller.

To install:

```console
make
sls deploy --region <region> --aws-profile <your profile>
```

This is a go project - if you don't have go installed  you can build using docker:

```console
docker run --rm -v "$PWD":/go/src/sls-service-discovery -w /go/src/sls-service-discovery golang:1.9 make
```

Once deployed, curl the url the sls provides using the x-api-key shown as well. You should see the output of www.example.com.

To tear down the example, do:

```console
sls remove --region <region> --aws-profile <your profile>
```

**Note**

Tear down is very slow, as the ENI that was allocated for use by lambda takes seemingly forever to delete.


## Follow On Work

* Demonstrate with internal services that register their name with the service when they boot.
* Show how health checks work to winnow the CNAME lookup response to only healthy service providers, responses based on weighting, etc