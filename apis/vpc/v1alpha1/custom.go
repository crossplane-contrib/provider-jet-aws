package v1alpha1

import (
	"github.com/crossplane-contrib/terrajet/pkg/configuration"
)

func vpcExternalNameConfigure(_ map[string]interface{}, _ string) {}

var vpcCustomConfig = configuration.Resource{
	ExternalName: configuration.ExternalName{
		ConfigureFunction: "vpcExternalNameConfigure",
		OmittedFields: []string{
			"cluster_identifier",
			"cluster_identifier_prefix",
		},
		// Set to true explicitly since the value is calculated by AWS.
		DisableNameInitializer: true,
	},
}

var ConfigStoreBuilder configuration.StoreBuilder

func init() {
	ConfigStoreBuilder.Register(func(s *configuration.Store) error {
		s.SetForResource("aws_vpc", vpcCustomConfig)
		return nil
	})
}
