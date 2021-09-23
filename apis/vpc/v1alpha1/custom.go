package v1alpha1

import (
	"github.com/crossplane-contrib/terrajet/pkg/terraform/resource"
)

func vpcExternalNameConfigure(_ map[string]interface{}, _ string) {}

var vpcCustomConfig = &resource.Configuration{
	ExternalName: resource.ExternalNameConfiguration{
		ConfigureFunction: "vpcExternalNameConfigure",
		OmittedFields: []string{
			"cluster_identifier",
			"cluster_identifier_prefix",
		},
		// Set to true explicitly since the value is calculated by AWS.
		DisableNameInitializer: true,
	},
}

var ConfigStoreBuilder resource.ConfigStoreBuilder

func init() {
	ConfigStoreBuilder.Register(func(c *resource.ConfigStore) error {
		c.SetConfigForResource("aws_vpc", vpcCustomConfig)
		return nil
	})
}
