package rds

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"
)

var SelfPkgPath = "github.com/crossplane-contrib/provider-tf-aws/config/rds"

func ClusterExternalNameConfigure(base map[string]interface{}, name string) {
	base["cluster_identifier"] = name
}

func init() {
	config.Store.SetForResource("aws_rds_cluster", config.Resource{
		ExternalName: config.ExternalName{
			ConfigureFunctionPath: SelfPkgPath + ".ClusterExternalNameConfigure",
			OmittedFields: []string{
				"cluster_identifier",
				"cluster_identifier_prefix",
			},
		},
		References: config.References{
			"s3_import[*].bucket_name": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/s3/v1alpha1.Bucket",
			},
		},
		UseAsync: true,
	})
}
