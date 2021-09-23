package v1alpha1

import (
	"github.com/crossplane-contrib/terrajet/pkg/terraform/resource"
	"github.com/crossplane-contrib/terrajet/pkg/types"

	s3v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/s3/v1alpha1"
)

func rdsClusterExternalNameConfigure(base map[string]interface{}, name string) {
	base["cluster_identifier"] = name
}

var rdsClusterCustomConfig = &resource.Configuration{
	ExternalName: resource.ExternalNameConfiguration{
		ConfigureFunction: "rdsClusterExternalNameConfigure",
		OmittedFields: []string{
			"cluster_identifier",
			"cluster_identifier_prefix",
		},
		// Set to false explicitly. For this resource, metadata.name is suitable
		// to be the default for external name.
		DisableNameInitializer: true,
	},
	Reference: map[string]resource.FieldReferenceConfiguration{
		"S3Import.BucketName": {
			ReferenceToType: types.PathForType(s3v1alpha1.S3Bucket{}),
		},
	},
}

var ConfigStoreBuilder resource.ConfigStoreBuilder

func init() {
	ConfigStoreBuilder.Register(func(c *resource.ConfigStore) error {
		c.SetConfigForResource("aws_rds_cluster", rdsClusterCustomConfig)
		return nil
	})
}
