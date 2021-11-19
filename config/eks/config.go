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

package eks

import (
	"context"
	"fmt"
	"strings"

	"github.com/crossplane-contrib/terrajet/pkg/config"
	"github.com/pkg/errors"

	"github.com/crossplane-contrib/provider-jet-aws/config/common"
)

// Configure adds configurations for eks group.
func Configure(p *config.Provider) { // nolint:gocyclo
	p.AddResourceConfigurator("aws_eks_cluster", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.References = config.References{
			"role_arn": {
				Type:      "github.com/crossplane-contrib/provider-jet-aws/apis/iam/v1alpha1.Role",
				Extractor: common.PathARNExtractor,
			},
			"vpc_config.subnet_ids": {
				Type:              "github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha1.Subnet",
				RefFieldName:      "SubnetIdRefs",
				SelectorFieldName: "SubnetIdSelector",
			},
			"vpc_config.security_group_ids": {
				Type:              "github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha1.SecurityGroup",
				RefFieldName:      "SecurityGroupIdRefs",
				SelectorFieldName: "SecurityGroupIdSelector",
			},
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_eks_node_group", func(r *config.Resource) {
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["node_group_name"] = name
			},
			OmittedFields: []string{
				"node_group_name",
				"node_group_name_prefix",
			},
			GetExternalNameFn: config.IDAsExternalName,
			GetIDFn:           config.ExternalNameAsID,
		}
		r.References = config.References{
			"cluster_name": {
				Type: "Cluster",
			},
			"node_role_arn": {
				Type:      "github.com/crossplane-contrib/provider-jet-aws/apis/iam/v1alpha1.Role",
				Extractor: common.PathARNExtractor,
			},
			"remote_access.source_security_group_ids": {
				Type:              "github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha1.SecurityGroup",
				RefFieldName:      "SourceSecurityGroupIdRefs",
				SelectorFieldName: "SourceSecurityGroupIdSelector",
			},
			"subnet_ids": {
				Type:              "github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha1.Subnet",
				RefFieldName:      "SubnetIdRefs",
				SelectorFieldName: "SubnetIdSelector",
			},
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_eks_identity_provider_config", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.References = config.References{
			"cluster_name": {
				Type: "Cluster",
			},
		}
	})

	p.AddResourceConfigurator("aws_eks_fargate_profile", func(r *config.Resource) {
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["fargate_profile_name"] = name
			},
			OmittedFields: []string{
				"fargate_profile_name",
			},
			GetIDFn: func(_ context.Context, externalName string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
				cl, ok := parameters["cluster_name"].(string)
				if !ok || cl == "" {
					return "", errors.New("cannot get cluster_name from parameters")
				}
				return fmt.Sprintf("%s:%s", cl, externalName), nil
			},
			GetExternalNameFn: func(tfstate map[string]interface{}) (string, error) {
				id, ok := tfstate["id"].(string)
				if !ok || id == "" {
					return "", errors.New("cannot get id from tfstate")
				}
				// my_cluster:my_fargate_profile
				w := strings.Split(id, ":")
				if len(w) != 2 {
					return "", errors.New("format of id should be my_cluster:my_fargate_profile")
				}
				return w[len(w)-1], nil
			},
		}

		r.References = config.References{
			"cluster_name": {
				Type: "Cluster",
			},
			"pod_execution_role_arn": {
				Type:      "github.com/crossplane-contrib/provider-jet-aws/apis/iam/v1alpha1.Role",
				Extractor: common.PathARNExtractor,
			},
			"subnet_ids": {
				Type:              "github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha1.Subnet",
				RefFieldName:      "SubnetIdRefs",
				SelectorFieldName: "SubnetIdSelector",
			},
		}
	})
	p.AddResourceConfigurator("aws_eks_addon", func(r *config.Resource) {
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, externalName string) {
				base["addon_name"] = externalName
			},
			OmittedFields: []string{
				"addon_name",
			},
			GetIDFn: func(_ context.Context, externalName string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
				cl, ok := parameters["cluster_name"].(string)
				if !ok || cl == "" {
					return "", errors.New("cannot get cluster_name from parameters")
				}
				return fmt.Sprintf("%s:%s", cl, externalName), nil
			},
			GetExternalNameFn: func(tfstate map[string]interface{}) (string, error) {
				id, ok := tfstate["id"].(string)
				if !ok || id == "" {
					return "", errors.New("cannot get id from tfstate")
				}
				// my_cluster:my_fargate_profile
				w := strings.Split(id, ":")
				if len(w) != 2 {
					return "", errors.New("format of id should be my_cluster:my_fargate_profile")
				}
				return w[len(w)-1], nil
			},
		}
		r.References = config.References{
			"cluster_name": {
				Type: "Cluster",
			},
			"service_account_role_arn": {
				Type:      "github.com/crossplane-contrib/provider-jet-aws/apis/iam/v1alpha1.Role",
				Extractor: common.PathARNExtractor,
			},
		}
	})
}
