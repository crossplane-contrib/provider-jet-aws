package vpc

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"
)

func VpcExternalNameConfigure(_ map[string]interface{}, _ string) {}

var vpcCustomConfig = config.Resource{
	Kind: "Vpc",
	ExternalName: config.ExternalName{
		ConfigureFunctionPath: "vpc.VpcExternalNameConfigure",
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
