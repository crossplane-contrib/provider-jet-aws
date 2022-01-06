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

package elasticloadbalancing

import (
	"github.com/crossplane/terrajet/pkg/config"
)

// Configure adds configurations for elasticloadbalancing group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_lb", func(r *config.Resource) {
		r.Version = "v1alpha2"
		r.ExternalName = config.IdentifierFromProvider
		r.ExternalName.OmittedFields = append(r.ExternalName.OmittedFields, "name_prefix")
		r.References = config.References{
			"security_groups": {
				Type:              "github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2.SecurityGroup",
				RefFieldName:      "SecurityGroupRefs",
				SelectorFieldName: "SecurityGroupSelector",
			},
			"subnets": {
				Type:              "github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2.Subnet",
				RefFieldName:      "SubnetRefs",
				SelectorFieldName: "SubnetSelector",
			},
			"access_logs.bucket": {
				Type: "github.com/crossplane-contrib/provider-jet-aws/apis/s3/v1alpha2.Bucket",
			},
			"subnet_mapping.subnet_id": {
				Type: "github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2.Subnet",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_lb_listener", func(r *config.Resource) {
		r.Version = "v1alpha2"
		r.ExternalName = config.IdentifierFromProvider
		r.References = config.References{
			"load_balancer_arn": {
				Type: "LB",
			},
			"default_action.target_group_arn": {
				Type: "LBTargetGroup",
			},
			"default_action.forward.target_group.arn": {
				Type: "LBTargetGroup",
			},
		}
	})

	p.AddResourceConfigurator("aws_lb_target_group", func(r *config.Resource) {
		r.Version = "v1alpha2"
		r.ExternalName = config.IdentifierFromProvider
		r.ExternalName.OmittedFields = append(r.ExternalName.OmittedFields, "name_prefix")
		r.References = config.References{
			"vpc_id": {
				Type: "github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2.VPC",
			},
		}
		if s, ok := r.TerraformResource.Schema["name"]; ok {
			s.Optional = false
			s.ForceNew = true
			s.Computed = false
		}
	})
	p.AddResourceConfigurator("aws_lb_target_group_attachment", func(r *config.Resource) {
		r.Version = "v1alpha2"
		r.ExternalName = config.IdentifierFromProvider
		r.References = config.References{
			"target_group_arn": {
				Type: "LBTargetGroup",
			},
		}
	})
}
