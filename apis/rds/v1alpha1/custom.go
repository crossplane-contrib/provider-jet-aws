package v1alpha1

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"
	"github.com/crossplane-contrib/terrajet/pkg/types"

	s3v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/s3/v1alpha1"
)

func rdsClusterExternalNameConfigure(base map[string]interface{}, name string) {
	base["cluster_identifier"] = name
}

var rdsClusterCustomConfig = config.Resource{
	ExternalName: config.ExternalName{
		ConfigureFunctionPath: "rdsClusterExternalNameConfigure",
		OmittedFields: []string{
			"cluster_identifier",
			"cluster_identifier_prefix",
		},
		// Set to false explicitly. For this resource, metadata.name is suitable
		// to be the default for external name.
		DisableNameInitializer: true,
	},
	References: config.References{
		"S3Import.BucketName": {
			Type: types.TypePath(s3v1alpha1.S3Bucket{}),
		},
	},
}

func init() {
	config.Store.SetForResource("aws_rds_cluster", rdsClusterCustomConfig)
}
