## How can assumeRoleARN be used with provider-jet-aws ?

provider-jet-aws will be configured to connect to AWS Account A via `InjectedIdentity` or `Secret` , request security credentials, and then `assumeRoleARN` to assume a role in AWS Account B to manage the resources within AWS Account B.

The first thing that needs to be done is to create an IAM role within AWS Account B that provider-jet-aws will `assumeRoleARN` into.

- From within the AWS console of AWS Account B, navigate to IAM > Roles > Create role > Another AWS account.

  -  Enter the Account ID of Account A (the account provider-jet-aws will call `assumeRoleARN` from).

Next, the provider-jet-aws must be configured to use `assumeRoleARN`. The code snippet below shows how to configure provider-jet-aws to connect to AWS Account A and assumeRoleARN into a role within AWS Account B.

```bash
cat > provider-config.yaml <<EOF
apiVersion: aws.jet.crossplane.io/v1alpha1
kind: ProviderConfig
metadata:
  name: account-b
spec:
  assumeRoleARN: "arn:aws:iam::999999999999:role/account-b"
  credentials:
    source: InjectedIdentity
EOF
```