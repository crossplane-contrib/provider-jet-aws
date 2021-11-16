package elasticloadbalancing

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_lb", func(r *config.Resource) {
		r.Kind = "LoadBalancer"
		r.Group = "elasticloadbalancing"
		r.ExternalName = config.IdentifierFromProvider
		r.References = config.References{
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
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_lb_listener", func(r *config.Resource) {
		r.Kind = "LoadBalancerListener"
		r.Group = "elasticloadbalancing"
		r.ExternalName = config.IdentifierFromProvider
		r.References = config.References{
			"load_balancer_arn": {
				Type: "LoadBalancer",
			},
			"default_action[*].target_group_arn": {
				Type: "LBTargetGroup",
			},
			"default_action[*].forward[*].target_group[*].arn": {
				Type: "LBTargetGroup",
			},
		}
	})

	p.AddResourceConfigurator("aws_lb_target_group", func(r *config.Resource) {
		r.Group = "elasticloadbalancing"
		r.ExternalName = config.IdentifierFromProvider
		r.References = config.References{
			"vpc_id": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1.VPC",
			},
		}
	})
	p.AddResourceConfigurator("aws_lb_target_group_attachment", func(r *config.Resource) {
		r.Group = "elasticloadbalancing"
		r.ExternalName = config.IdentifierFromProvider
		r.References = config.References{
			"target_group_arn": {
				Type: "LBTargetGroup",
			},
		}
	})
}
