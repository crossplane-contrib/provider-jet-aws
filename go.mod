module github.com/crossplane-contrib/provider-tf-aws

go 1.16

replace github.com/crossplane-contrib/terrajet => github.com/muvaf/terrajet v0.0.0-20210818141613-e4fa1ecf79bc

require (
	github.com/crossplane-contrib/terrajet v0.0.0-20210809201716-65ef979f8f10
	github.com/crossplane/crossplane-runtime v0.14.1-0.20210805220729-047d9387efbb
	github.com/crossplane/crossplane-tools v0.0.0-20210320162312-1baca298c527
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.7.0
	github.com/iancoleman/strcase v0.2.0
	github.com/pkg/errors v0.9.1
	github.com/terraform-providers/terraform-provider-aws v1.60.1-0.20210811232925-d6f99829ec3f
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	k8s.io/apimachinery v0.22.0
	k8s.io/client-go v0.22.0
	sigs.k8s.io/controller-runtime v0.9.2
	sigs.k8s.io/controller-tools v0.4.0
)

replace github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/gdavison/terraform-plugin-sdk/v2 v2.0.2-0.20210714181518-b5a3dc95a675
