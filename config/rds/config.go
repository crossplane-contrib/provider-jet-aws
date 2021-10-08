package rds

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"
)

func RdsClusterExternalNameConfigure(base map[string]interface{}, name string) {
	base["cluster_identifier"] = name
}

var rdsClusterCustomConfig = config.Resource{
	ExternalName: config.ExternalName{
		ConfigureFunctionPath: "rds.RdsClusterExternalNameConfigure",
		OmittedFields: []string{
			"cluster_identifier",
			"cluster_identifier_prefix",
		},
	},
	References: config.References{
		"s3_import[*].bucket_name": {
			Type: "github.com/crossplane-contrib/provider-tf-aws/apis/s3/v1alpha1.S3Bucket",
		},
	},
	UseAsync: true,
}

func init() {
	config.Store.SetForResource("aws_rds_cluster", rdsClusterCustomConfig)
}
