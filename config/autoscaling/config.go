/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package autoscaling

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-jet-aws/config/common"
)

// Configure adds configurations for autoscaling group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_autoscaling_group", func(r *config.Resource) {
		r.Kind = "AutoscalingGroup"
		r.ExternalName = config.NameAsIdentifier
		r.References["vpc_zone_identifier"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha1.Subnet",
		}
		r.UseAsync = true

		// Managed by Attachment resource.
		if s, ok := r.TerraformResource.Schema["load_balancers"]; ok {
			s.Optional = false
			s.Computed = true
		}
		if s, ok := r.TerraformResource.Schema["target_group_arns"]; ok {
			s.Optional = false
			s.Computed = true
		}
	})
	p.AddResourceConfigurator("aws_autoscaling_attachment", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.References["autoscaling_group_name"] = config.Reference{
			Type: "AutoscalingGroup",
		}
		r.References["alb_target_group_arn"] = config.Reference{
			Type:      "github.com/crossplane-contrib/provider-jet-aws/apis/elasticloadbalancing/v1alpha1.TargetGroup",
			Extractor: common.PathARNExtractor,
		}
	})
}
