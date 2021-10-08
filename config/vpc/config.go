package vpc

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"
)

func init() {
	config.Store.SetForResource("aws_vpc", config.Resource{
		Kind: "VPC",
		ExternalName: config.ExternalName{
			// Set to true explicitly since the value is calculated by AWS.
			DisableNameInitializer: true,
		},
	})
}
