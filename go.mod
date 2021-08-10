module github.com/crossplane/provider-tf-aws

go 1.16

require (
	github.com/crossplane-contrib/terrajet v0.0.0-00010101000000-000000000000
	github.com/crossplane/crossplane-runtime v0.14.1-0.20210805220729-047d9387efbb
	github.com/crossplane/crossplane-tools v0.0.0-20210320162312-1baca298c527
	github.com/json-iterator/go v1.1.11
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	k8s.io/apimachinery v0.21.2
	k8s.io/client-go v0.21.2
	sigs.k8s.io/controller-runtime v0.9.2
	sigs.k8s.io/controller-tools v0.3.0
)

replace github.com/crossplane-contrib/terrajet => github.com/turkenh/terrajet v0.0.0-20210810130304-17c1f04a6b45
