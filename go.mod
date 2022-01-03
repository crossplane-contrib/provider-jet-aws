module github.com/crossplane-contrib/provider-jet-aws

go 1.16

require (
	github.com/aws/aws-sdk-go-v2 v0.23.0
	github.com/crossplane/crossplane-runtime v0.15.1-0.20211004150827-579c1833b513
	github.com/crossplane/crossplane-tools v0.0.0-20210916125540-071de511ae8e
	github.com/crossplane/provider-aws v0.19.0
	github.com/crossplane/terrajet v0.2.1-0.20220103083602-2c54e8491aba
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.7.0
	github.com/pkg/errors v0.9.1
	github.com/terraform-providers/terraform-provider-aws v1.60.1-0.20210811232925-d6f99829ec3f
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	k8s.io/apimachinery v0.22.0
	k8s.io/client-go v0.22.0
	k8s.io/utils v0.0.0-20210722164352-7f3ee0f31471
	sigs.k8s.io/controller-runtime v0.9.6
	sigs.k8s.io/controller-tools v0.6.2
)

replace github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/gdavison/terraform-plugin-sdk/v2 v2.0.2-0.20210714181518-b5a3dc95a675
