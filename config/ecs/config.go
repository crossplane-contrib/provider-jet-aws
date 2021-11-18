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

package ecs

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-tf-aws/config/common"
)

// Configure adds configurations for ecs group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_ecs_cluster", func(r *config.Resource) {
		// todo: use new ID operation overrides because ID is ARN.
		r.ExternalName = config.IdentifierFromProvider
		r.References = config.References{
			"capacity_providers": config.Reference{
				Type: "CapacityProvider",
			},
			"execute_command_configuration.kms_key_id": config.Reference{
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/kms/v1alpha1.Key",
			},
			"log_configuration.s3_bucket_name": config.Reference{
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/s3/v1alpha1.Bucket",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_ecs_service", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.References = config.References{
			"cluster": config.Reference{
				Type:      "Cluster",
				Extractor: common.PathARNExtractor,
			},
			"iam_role": config.Reference{
				Type:      "github.com/crossplane-contrib/provider-tf-aws/apis/iam/v1alpha1.Role",
				Extractor: common.PathARNExtractor,
			},
			"network_configuration.subnets": config.Reference{
				Type:              "github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1.Subnet",
				RefFieldName:      "SubnetRefs",
				SelectorFieldName: "SubnetSelector",
			},
			"network_configuration.security_groups": config.Reference{
				Type:              "github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1.SecurityGroup",
				RefFieldName:      "SecurityGroupRefs",
				SelectorFieldName: "SecurityGroupSelector",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_ecs_capacity_provider", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.References = config.References{
			"auto_scaling_group_provider.auto_scaling_group_arn": config.Reference{
				Type:      "github.com/crossplane-contrib/provider-tf-aws/apis/autoscaling/v1alpha1.AutoscalingGroup",
				Extractor: common.PathARNExtractor,
			},
		}
	})

	p.AddResourceConfigurator("aws_ecs_tag", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.References = config.References{
			// Note(turkenh): This reference could correspond multiple types as
			// per documentation, any ecs resource: https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecs_tag#resource_arn
			// However, we could only reference to one type.
			"resource_arn": config.Reference{
				Type:      "Cluster",
				Extractor: common.PathARNExtractor,
			},
		}
	})
	p.AddResourceConfigurator("aws_ecs_task_definition", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.References = config.References{
			"execution_role_arn": config.Reference{
				Type:      "github.com/crossplane-contrib/provider-tf-aws/apis/iam/v1alpha1.Role",
				Extractor: common.PathARNExtractor,
			},
		}
	})
}
