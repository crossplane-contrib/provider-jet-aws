package autoscaling

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-tf-aws/config/common"
)

func Configure(p config.Provider) {
	p.AddResourceConfigurator("aws_autoscaling_group", func(r *config.Resource) {
		r.Kind = "AutoscalingGroup"
		r.ExternalName = config.NameAsIdentifier
		r.References["vpc_zone_identifier"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1.Subnet",
		}
		r.References["target_group_arns"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tf-aws/apis/lb/v1alpha1.LBTargetGroup",
		}

		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_autoscaling_attachment", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.References["autoscaling_group_name"] = config.Reference{
			Type: "AutoscalingGroup",
		}
		r.References["alb_target_group_arn"] = config.Reference{
			Type:      "github.com/crossplane-contrib/provider-tf-aws/apis/lb/v1alpha1.LBTargetGroup",
			Extractor: common.PathARNExtractor,
		}
	})
}
