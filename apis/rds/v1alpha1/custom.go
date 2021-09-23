package v1alpha1

import "github.com/crossplane-contrib/terrajet/pkg/terraform/resource"

// DBClusterExternalNameConfig is config for external name mechanism of DBCluster.
var (
	DBClusterExternalNameConfig = resource.ExternalNameConfiguration{
		SelfVarPath: "github.com/crossplane-contrib/provider-tf-aws/apis/rds/v1alpha1.DBClusterExternalNameConfig",
		Configure: func(base map[string]interface{}, externalName string) {
			base["cluster_identifier"] = externalName
		},
		OmittedFields: []string{
			"cluster_identifier",
			"cluster_identifier_prefix",
		},
		// Set to false explicitly. For this resource, metadata.name is suitable
		// to be the default for external name.
		DisableNameInitializer: false,
	}
)
