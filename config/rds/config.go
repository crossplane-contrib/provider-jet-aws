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

package rds

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-tf-aws/config/common"
)

const (
	// SelfPackagePath is the golang path for this package.
	SelfPackagePath = "github.com/crossplane-contrib/provider-tf-aws/config/rds"
)

// ClusterExternalNameConfigure configures name of cluster.
func ClusterExternalNameConfigure(base map[string]interface{}, name string) {
	base["cluster_identifier"] = name
}

// DBInstanceExternalNameConfigure configures name of db instance.
func DBInstanceExternalNameConfigure(base map[string]interface{}, name string) {
	base["identifier"] = name
}

func init() {
	config.Store.SetForResource("aws_rds_cluster", config.Resource{
		ExternalName: config.ExternalName{
			ConfigureFunctionPath: SelfPackagePath + ".ClusterExternalNameConfigure",
			OmittedFields: []string{
				"cluster_identifier",
				"cluster_identifier_prefix",
			},
		},
		References: config.References{
			"s3_import[*].bucket_name": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/s3/v1alpha1.Bucket",
			},
			"vpc_security_group_ids": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1.SecurityGroup",
			},
			"restore_to_point_in_time[*].source_cluster_identifier": {
				Type: "Cluster",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("aws_db_instance", config.Resource{
		Kind: "DBInstance",
		ExternalName: config.ExternalName{
			ConfigureFunctionPath: SelfPackagePath + ".DBInstanceExternalNameConfigure",
			OmittedFields: []string{
				"identifier",
				"identifier_prefix",
			},
		},
		References: config.References{
			"restore_to_point_in_time[*].source_db_instance_identifier": {
				Type: "DBInstance",
			},
			"s3_import[*].bucket_name": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/s3/v1alpha1.Bucket",
			},
			"kms_key_id": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/kms/v1alpha1.Key",
			},
			"performance_insights_kms_key_id": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/kms/v1alpha1.Key",
			},
			"restore_to_point_in_time[*].source_cluster_identifier": {
				Type: "Cluster",
			},
			"security_group_names": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1.SecurityGroup",
			},
			"vpc_security_group_ids": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1.SecurityGroup",
			},
			"parameter_group_name": {
				Type: "DBParameterGroup",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("aws_db_parameter_group", config.Resource{
		Kind: "DBParameterGroup",
		ExternalName: config.ExternalName{
			ConfigureFunctionPath: common.PathExternalNameAsName,
			OmittedFields: []string{
				"name",
				"name_prefix",
			},
		},
	})
}
