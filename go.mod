module github.com/crossplane-contrib/provider-tf-aws

go 1.16

replace github.com/crossplane-contrib/terrajet => github.com/muvaf/terrajet v0.0.0-20211027185749-c68d2062ce88

require (
	github.com/aws/aws-sdk-go-v2 v0.23.0
	github.com/crossplane-contrib/terrajet v0.1.1-0.20211020104517-c9e21387298d
	github.com/crossplane/crossplane-runtime v0.15.1-0.20211004150827-579c1833b513
	github.com/crossplane/crossplane-tools v0.0.0-20210916125540-071de511ae8e
	github.com/crossplane/provider-aws v0.19.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.7.0
	github.com/iancoleman/strcase v0.2.0
	github.com/pkg/errors v0.9.1
	github.com/terraform-providers/terraform-provider-aws v1.60.1-0.20210811232925-d6f99829ec3f
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	k8s.io/apiextensions-apiserver v0.21.3
	k8s.io/apimachinery v0.22.0
	k8s.io/client-go v0.22.0
	sigs.k8s.io/controller-runtime v0.9.6
	sigs.k8s.io/controller-tools v0.6.2
)

// We have to have this replace statement because github.com/terraform-providers/terraform-provider-aws
// uses this fork.
replace github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/gdavison/terraform-plugin-sdk/v2 v2.0.2-0.20210714181518-b5a3dc95a675
