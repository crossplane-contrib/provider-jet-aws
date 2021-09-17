package v1alpha1

import "github.com/crossplane-contrib/terrajet/pkg/terraform/resource"

// DBClusterExternalNamer is config for external name mechanism of DBCluster.
var DBClusterExternalNamer = resource.ExternalNamer{
	SelfVarPath: "github.com/crossplane-contrib/provider-tf-aws/apis/rds/v1alpha1.DBClusterExternalNamer",
	Configure: func(base map[string]interface{}, name string) {
		base["cluster_identifier"] = name
	},
	OmittedFields: []string{
		"cluster_identifier",
		"cluster_identifier_prefix",
	},
	// Set to false explicitly. For this resource, metadata.name is suitable
	// to be the default for external name.
	DisableNameInitializer: true,
}
