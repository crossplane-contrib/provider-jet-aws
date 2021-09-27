package v1alpha1

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"
)

func vpcExternalNameConfigure(_ map[string]interface{}, _ string) {}

var vpcCustomConfig = config.Resource{
	ExternalName: config.ExternalName{
		ConfigureFunctionPath: "vpcExternalNameConfigure",
		OmittedFields: []string{
			"cluster_identifier",
			"cluster_identifier_prefix",
		},
		// Set to true explicitly since the value is calculated by AWS.
		DisableNameInitializer: true,
	},
}

func init() {
	config.Store.SetForResource("aws_vpc", vpcCustomConfig)
}
