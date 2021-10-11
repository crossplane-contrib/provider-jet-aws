package lbv2

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"
)

func init() {

	config.Store.SetForResource("aws_lb", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: map[string]config.Reference{
			"security_groups": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/vpc/v1alpha1.SecurityGroup",
			},
			"subnets": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/vpc/v1alpha1.Subnet",
			},
			"access_logs[*].bucket": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/s3/v1alpha1.Bucket",
			},
			"subnet_mapping[*].subnet_id": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/vpc/v1alpha1.Subnet",
			},
		},
		UseAsync: true,
	})

	config.Store.SetForResource("aws_lb_listener", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: map[string]config.Reference{
			"load_balancer_arn": {
				Type: "Lb",
			},
			"default_action[*].target_group_arn": {
				Type: "LbTargetGroup",
			},
			"default_action[*].forward[*].target_group[*].arn": {
				Type: "LbTargetGroup",
			},
		},
	})

	config.Store.SetForResource("aws_lb_target_group", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: map[string]config.Reference{
			"vpc_id": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/vpc/v1alpha1.VPC",
			},
		},
	})

	config.Store.SetForResource("aws_lb_target_group_attachment", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: map[string]config.Reference{
			"target_group_arn": {
				Type: "LbTargetGroup",
			},
		},
	})
}
