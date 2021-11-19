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
)

// Configure adds configurations for rds group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_rds_cluster", func(r *config.Resource) {
		r.Kind = "DBCluster"
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["cluster_identifier"] = name
			},
			OmittedFields: []string{
				"cluster_identifier",
				"cluster_identifier_prefix",
			},
			GetExternalNameFn: config.IDAsExternalName,
			GetIDFn:           config.ExternalNameAsID,
		}
		r.References = config.References{
			"s3_import.bucket_name": {
				Type: "github.com/crossplane-contrib/provider-jet-aws/apis/s3/v1alpha1.Bucket",
			},
			"vpc_security_group_ids": {
				Type:              "github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha1.SecurityGroup",
				RefFieldName:      "VpcSecurityGroupIdRefs",
				SelectorFieldName: "VpcSecurityGroupIdSelector",
			},
			"restore_to_point_in_time.source_cluster_identifier": {
				Type: "Cluster",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_db_instance", func(r *config.Resource) {
		r.Kind = "DBInstance"
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["identifier"] = name
			},
			OmittedFields: []string{
				"identifier",
				"identifier_prefix",
			},
			GetExternalNameFn: config.IDAsExternalName,
			GetIDFn:           config.ExternalNameAsID,
		}
		r.References = config.References{
			"restore_to_point_in_time.source_db_instance_identifier": {
				Type: "DBInstance",
			},
			"s3_import.bucket_name": {
				Type: "github.com/crossplane-contrib/provider-jet-aws/apis/s3/v1alpha1.Bucket",
			},
			"kms_key_id": {
				Type: "github.com/crossplane-contrib/provider-jet-aws/apis/kms/v1alpha1.Key",
			},
			"performance_insights_kms_key_id": {
				Type: "github.com/crossplane-contrib/provider-jet-aws/apis/kms/v1alpha1.Key",
			},
			"restore_to_point_in_time.source_cluster_identifier": {
				Type: "Cluster",
			},
			"security_group_names": {
				Type:              "github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha1.SecurityGroup",
				RefFieldName:      "SecurityGroupNameRefs",
				SelectorFieldName: "SecurityGroupNameSelector",
			},
			"vpc_security_group_ids": {
				Type:              "github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha1.SecurityGroup",
				RefFieldName:      "VpcSecurityGroupIdRefs",
				SelectorFieldName: "VpcSecurityGroupIdSelector",
			},
			"parameter_group_name": {
				Type: "DBParameterGroup",
			},
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_db_parameter_group", func(r *config.Resource) {
		r.Kind = "DBParameterGroup"
		r.ExternalName = config.NameAsIdentifier
	})
}
