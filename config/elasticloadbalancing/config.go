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
	"github.com/crossplane-contrib/terrajet/pkg/config"
)

// Configure adds configurations for elasticloadbalancing group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_lb", func(r *config.Resource) {
		r.Kind = "LoadBalancer"
		r.ExternalName = config.IdentifierFromProvider
		r.References = config.References{
			"security_groups": {
				Type:              "github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1.SecurityGroup",
				RefFieldName:      "SecurityGroupRefs",
				SelectorFieldName: "SecurityGroupSelector",
			},
			"subnets": {
				Type:              "github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1.Subnet",
				RefFieldName:      "SubnetRefs",
				SelectorFieldName: "SubnetSelector",
			},
			"access_logs.bucket": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/s3/v1alpha1.Bucket",
			},
			"subnet_mapping.subnet_id": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1.Subnet",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_lb_listener", func(r *config.Resource) {
		r.Kind = "LoadBalancerListener"
		r.ExternalName = config.IdentifierFromProvider
		r.References = config.References{
			"load_balancer_arn": {
				Type: "LoadBalancer",
			},
			"default_action.target_group_arn": {
				Type: "TargetGroup",
			},
			"default_action.forward.target_group.arn": {
				Type: "TargetGroup",
			},
		}
	})

	p.AddResourceConfigurator("aws_lb_target_group", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.References = config.References{
			"vpc_id": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1.VPC",
			},
		}
	})
	p.AddResourceConfigurator("aws_lb_target_group_attachment", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.References = config.References{
			"target_group_arn": {
				Type: "TargetGroup",
			},
		}
	})
}
