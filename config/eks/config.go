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
	"github.com/crossplane-contrib/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-tf-aws/config/common"
)

const (
	// SelfPackagePath is the golang path for this package.
	SelfPackagePath = "github.com/crossplane-contrib/provider-tf-aws/config/eks"
)

// ClusterExternalNameConfigure configures cluster name.
func ClusterExternalNameConfigure(base map[string]interface{}, name string) {
	base["name"] = name
}

// NodeGroupExternalNameConfigure configures nodegroup name.
func NodeGroupExternalNameConfigure(base map[string]interface{}, name string) {
	base["name"] = name
}

// FargateProfileExternalNameConfigure configures fargate profile name.
func FargateProfileExternalNameConfigure(base map[string]interface{}, name string) {
	base["fargate_profile_name"] = name
}

func init() {
	config.Store.SetForResource("aws_eks_cluster", config.Resource{
		ExternalName: config.ExternalName{
			ConfigureFunctionPath: SelfPackagePath + ".ClusterExternalNameConfigure",
			OmittedFields: []string{
				"name",
			},
		},
		References: config.References{
			"role_arn": {
				Type:      "github.com/crossplane-contrib/provider-tf-aws/apis/iam/v1alpha1.Role",
				Extractor: common.PathARNExtractor,
			},
			"vpc_config[*].subnet_ids": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1.Subnet",
			},
			"vpc_config[*].security_group_ids": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1.SecurityGroup",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("aws_eks_node_group", config.Resource{
		ExternalName: config.ExternalName{
			ConfigureFunctionPath: SelfPackagePath + ".NodeGroupExternalNameConfigure",
			OmittedFields: []string{
				"node_group_name",
				"node_group_name_prefix",
			},
		},
		References: config.References{
			"cluster_name": {
				Type: "Cluster",
			},
			"node_role_arn": {
				Type:      "github.com/crossplane-contrib/provider-tf-aws/apis/iam/v1alpha1.Role",
				Extractor: common.PathARNExtractor,
			},
			"remote_access[*].source_security_group_ids": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1.SecurityGroup",
			},
			"subnet_ids": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1.Subnet",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("aws_eks_identity_provider_config", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"cluster_name": {
				Type: "Cluster",
			},
		},
	})
	config.Store.SetForResource("aws_eks_fargate_profile", config.Resource{
		ExternalName: config.ExternalName{
			OmittedFields: []string{
				"fargate_profile_name",
			},
		},
		References: config.References{
			"cluster_name": {
				Type: "Cluster",
			},
			"pod_execution_role_arn": {
				Type:      "github.com/crossplane-contrib/provider-tf-aws/apis/iam/v1alpha1.Role",
				Extractor: common.PathARNExtractor,
			},
			"subnet_ids": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1.Subnet",
			},
		},
	})
	config.Store.SetForResource("aws_eks_addon", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"cluster_name": {
				Type: "Cluster",
			},
			"service_account_role_arn": {
				Type:      "github.com/crossplane-contrib/provider-tf-aws/apis/iam/v1alpha1.Role",
				Extractor: common.PathARNExtractor,
			},
		},
	})
}
