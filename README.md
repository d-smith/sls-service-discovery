# sls-service-discovery

## vpc

```console
aws cloudformation create-stack --stack-name vpc --template-body file://vpc.yml
```

##private hosted zone

```console
aws cloudformation create-stack \
--stack-name hosted-zone \
--template-body file://r53-private-hosted-zone.yml \
--parameters ParameterKey=Name,ParameterValue=sector7.com \
ParameterKey=Comment,ParameterValue=None \
ParameterKey=VpcId,ParameterValue=vpc-573ece2f 
```