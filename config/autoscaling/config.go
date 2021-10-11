package autoscaling

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-tf-aws/config/common"
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
		References: map[string]config.Reference{
			"vpc_zone_identifier": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/vpc/v1alpha1.Subnet",
			},
			"target_group_arns": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/lb/v1alpha1.LBTargetGroup",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("aws_autoscaling_attachment", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: map[string]config.Reference{
			"autoscaling_group_name": {
				Type: "AutoscalingGroup",
			},
			"alb_target_group_arn": {
				Type:      "github.com/crossplane-contrib/provider-tf-aws/apis/lb/v1alpha1.LBTargetGroup",
				Extractor: common.PathARNExtractor,
			},
		},
	})
}
