module github.com/crossplane-contrib/provider-jet-aws

go 1.16

require (
	github.com/aws/aws-sdk-go-v2 v0.23.0
	github.com/crossplane/crossplane-runtime v0.15.1-0.20211004150827-579c1833b513
	github.com/crossplane/crossplane-tools v0.0.0-20210916125540-071de511ae8e
	github.com/crossplane/provider-aws v0.19.0
	github.com/crossplane/terrajet v0.3.1
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.10.1
	github.com/hashicorp/terraform-provider-aws v1.60.1-0.20220106231712-d14b20ba8b08
	github.com/pkg/errors v0.9.1
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	k8s.io/apimachinery v0.22.0
	k8s.io/client-go v0.22.0
	k8s.io/utils v0.0.0-20211208161948-7d6a63dca704
	sigs.k8s.io/controller-runtime v0.9.6
	sigs.k8s.io/controller-tools v0.6.2
)

replace github.com/hashicorp/terraform-provider-aws => ./.work/.tfaws
