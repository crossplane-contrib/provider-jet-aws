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
	"context"
	"path/filepath"
	"strings"

	"github.com/crossplane/terrajet/pkg/config"
	"github.com/pkg/errors"

	"github.com/crossplane-contrib/provider-jet-aws/config/common"
)

// Configure adds configurations for ecs group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_ecs_cluster", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]interface{}) (string, error) {
			// expected id format: arn:aws:ecs:us-west-2:123456789123:cluster/example-cluster
			w := strings.Split(tfstate["id"].(string), "/")
			if len(w) != 2 {
				return "", errors.New("terraform ID should be the ARN of the cluster")
			}
			return w[len(w)-1], nil
		}
		r.References = config.References{
			"capacity_providers": config.Reference{
				Type: "CapacityProvider",
			},
			"execute_command_configuration.kms_key_id": config.Reference{
				Type: "github.com/crossplane-contrib/provider-jet-aws/apis/kms/v1alpha2.Key",
			},
			"log_configuration.s3_bucket_name": config.Reference{
				Type: "github.com/crossplane-contrib/provider-jet-aws/apis/s3/v1alpha2.Bucket",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_ecs_service", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]interface{}) (string, error) {
			// expected id format: arn:aws:ecs:us-east-2:123456789123:service/sample-cluster/sample-service
			w := strings.Split(tfstate["id"].(string), "/")
			if len(w) != 3 {
				return "", errors.New("terraform ID should be the ARN of the service")
			}
			return w[len(w)-1], nil
		}
		r.ExternalName.GetIDFn = func(_ context.Context, externalName string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
			cl, ok := parameters["cluster"].(string)
			if !ok {
				return "", errors.New("cannot generate id without cluster paramater")
			}
			return filepath.Join(cl, externalName), nil
		}
		r.References = config.References{
			"cluster": config.Reference{
				Type:      "Cluster",
				Extractor: common.PathARNExtractor,
			},
			"iam_role": config.Reference{
				Type:      "github.com/crossplane-contrib/provider-jet-aws/apis/iam/v1alpha2.Role",
				Extractor: common.PathARNExtractor,
			},
			"network_configuration.subnets": config.Reference{
				Type:              "github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2.Subnet",
				RefFieldName:      "SubnetRefs",
				SelectorFieldName: "SubnetSelector",
			},
			"network_configuration.security_groups": config.Reference{
				Type:              "github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2.SecurityGroup",
				RefFieldName:      "SecurityGroupRefs",
				SelectorFieldName: "SecurityGroupSelector",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_ecs_capacity_provider", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.NameAsIdentifier
		r.References = config.References{
			"auto_scaling_group_provider.auto_scaling_group_arn": config.Reference{
				Type:      "github.com/crossplane-contrib/provider-jet-aws/apis/autoscaling/v1alpha2.AutoscalingGroup",
				Extractor: common.PathARNExtractor,
			},
		}
	})

	p.AddResourceConfigurator("aws_ecs_task_definition", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
		r.References = config.References{
			"execution_role_arn": config.Reference{
				Type:      "github.com/crossplane-contrib/provider-jet-aws/apis/iam/v1alpha2.Role",
				Extractor: common.PathARNExtractor,
			},
		}
	})
}
