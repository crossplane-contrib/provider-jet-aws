# Terrajet AWS Provider

`provider-tf-aws` is a [Crossplane](https://crossplane.io/) provider that is
built using [Terrajet](https://github.com/crossplane-contrib/terrajet) code
generation tools and exposes XRM-conformant managed resources for
[Amazon AWS](https://aws.amazon.com/).

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://github.com/crossplane-contrib/provider-tf-aws/releases):
```
kubectl crossplane install provider crossplane/provider-tf-aws:v0.2.1
```

You can see the API reference [here](https://doc.crds.dev/github.com/crossplane-contrib/provider-tf-aws).

## Contributing: Adding a New Resource

| Please note: the steps provided here is subject for changes since Terrajet is still alpha and does not guarantee a stabilized interface yet. |
| --- |

Adding a new resource to [Terrajet](https://github.com/crossplane-contrib/terrajet)
based providers is a very straightforward process which is basically figuring
out and providing the required [custom configuration](/docs/custom-configuration.md)
for the resource.

### Steps to Follow

1. **Include Resource:**

   Add the Terraform resource name to the list of resources
   (`includedResources`) in the [generator main.go](/cmd/generator/main.go).

   For example, to add [aws_iam_access_key](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_access_key)
   we add [this line](https://github.com/crossplane-contrib/provider-tf-aws/blob/d2c0024f18b5760e4d2222c405ad0501c63ee0b2/cmd/generator/main.go#L108).


2. **Create and Register Group Config File:** (if not exist)

    1. Create the custom config file for the group:

      ```console
      GROUP=<group-name>
      RESOURCE=<terraform-resource-name>
      mkdir config/$GROUP
      cat << EOF > config/$GROUP/config.go
      package $GROUP

      import (
        "github.com/crossplane-contrib/terrajet/pkg/config"
      )
      
      func init() {
        config.Store.SetForResource("$RESOURCE", config.Resource{})
      }
      EOF
      ```

   For `aws_iam_access_key` example:

      ```
      GROUP=iam
      RESOURCE=aws_iam_access_key
      ```

    2. Register this custom configuration package in custom/register.go. See
       [this](https://github.com/crossplane-contrib/provider-tf-aws/blob/main/config/register.go#L11)
       for iam as an example.

3. **Add custom configuration** in `config/$GROUP/config.go`:

    1. [Configure external name](/docs/custom-configuration.md#external-name).

    2. [Cross resource referencing](/docs/custom-configuration.md#cross-resource-referencing).

    3. [Configure Connection Secret Keys](/docs/custom-configuration.md#additional-sensitive-fields-and-custom-connection-details).

    4. [Late initialization configuration](/docs/custom-configuration.md#late-initialization-behavior).

5. Run code-generation:

    1. Run Terrajet code-generation pipeline:

      ```console
      go run cmd/generator/main.go
      ```

    1. Run code-generation for Controller Tools and Crossplane Tools:

      ```console
      make generate
      ```

6. Run against a Kubernetes cluster:

    ```console
    kubectl apply -f package/crds
    make run
    ```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/crossplane/provider-tf-aws/issues).

## Contact

Please use the following to reach members of the community:

* Slack: Join our [slack channel](https://slack.crossplane.io)
* Forums:
  [crossplane-dev](https://groups.google.com/forum/#!forum/crossplane-dev)
* Twitter: [@crossplane_io](https://twitter.com/crossplane_io)
* Email: [info@crossplane.io](mailto:info@crossplane.io)

## Governance and Owners

provider-tf-aws is run according to the same
[Governance](https://github.com/crossplane/crossplane/blob/master/GOVERNANCE.md)
and [Ownership](https://github.com/crossplane/crossplane/blob/master/OWNERS.md)
structure as the core Crossplane project.

## Code of Conduct

provider-tf-aws adheres to the same [Code of
Conduct](https://github.com/crossplane/crossplane/blob/master/CODE_OF_CONDUCT.md)
as the core Crossplane project.

## Licensing

provider-tf-aws is under the Apache 2.0 license.
