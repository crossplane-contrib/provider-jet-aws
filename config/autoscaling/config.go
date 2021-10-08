package autoscaling

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-tf-aws/config/common"
)

const (
	CommonPkgPath = "github.com/crossplane-contrib/provider-tf-aws/config/common"
	SelfPkgPath   = "github.com/crossplane-contrib/provider-tf-aws/config/autoscaling"
)

func init() {
	config.Store.SetForResource("aws_autoscaling_group", config.Resource{
		Kind: "AutoscalingGroup",
		ExternalName: config.ExternalName{
			ConfigureFunctionPath: common.PathExternalNameAsName,
			OmittedFields: []string{
				"name",
				"name_prefix",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("aws_autoscaling_attachment", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: map[string]config.Reference{
			"autoscaling_group_name": config.Reference{
				Type:              "AutoscalingGroup",
				RefFieldName:      "AutoscalingGroupRef",
				SelectorFieldName: "AutoscalingGroupSelector",
			},
		},
	})
}
