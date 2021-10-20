package lb

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"
)

func init() {

	config.Store.SetForResource("aws_lb", config.Resource{
		Kind:         "LoadBalancer",
		ExternalName: config.IdentifierFromProvider,
		References: map[string]config.Reference{
			"security_groups": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1.SecurityGroup",
			},
			"subnets": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1.Subnet",
			},
			"access_logs[*].bucket": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/s3/v1alpha1.Bucket",
			},
			"subnet_mapping[*].subnet_id": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1.Subnet",
			},
		},
		UseAsync: true,
	})

	config.Store.SetForResource("aws_lb_listener", config.Resource{
		Kind:         "LBListener",
		ExternalName: config.IdentifierFromProvider,
		References: map[string]config.Reference{
			"load_balancer_arn": {
				Type: "LoadBalancer",
			},
			"default_action[*].target_group_arn": {
				Type: "LBTargetGroup",
			},
			"default_action[*].forward[*].target_group[*].arn": {
				Type: "LBTargetGroup",
			},
		},
	})

	config.Store.SetForResource("aws_lb_target_group", config.Resource{
		// Note(turkenh): Without setting Kind here, we get kind as "TargetGroup"
		// which sounds better however, generator fails with the following since
		// it conflicts with lblisters "TargetGroupParameters":
		//
		// panic: cannot generate crd: cannot build types for TargetGroup:
		//  cannot build the types: cannot generate parameters type name of
		//  TargetGroup: could not generate a unique name for TargetGroupParameters
		Kind:         "LBTargetGroup",
		ExternalName: config.IdentifierFromProvider,
		References: map[string]config.Reference{
			"vpc_id": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1.VPC",
			},
		},
	})

	config.Store.SetForResource("aws_lb_target_group_attachment", config.Resource{
		Kind:         "LBTargetGroupAttachment",
		ExternalName: config.IdentifierFromProvider,
		References: map[string]config.Reference{
			"target_group_arn": {
				Type: "LBTargetGroup",
			},
		},
	})
}
