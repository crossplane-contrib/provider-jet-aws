package v1alpha1

import "github.com/crossplane-contrib/terrajet/pkg/terraform/resource"

var (
	// VPCExternalNameConfig is config for external name mechanism of VPC.
	VPCExternalNameConfig = resource.ExternalNameConfiguration{
		SelfVarPath: "github.com/crossplane-contrib/provider-tf-aws/apis/vpc/v1alpha1.VPCExternalNameConfig",
		Configure:   resource.NopConfigureWithName,
		// Set to true explicitly since the value is calculated by AWS.
		DisableNameInitializer: true,
	}
)
