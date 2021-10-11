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

const (
	// SelfPackagePath is the golang path for this package.
	SelfPackagePath = "github.com/crossplane-contrib/provider-tf-aws/config/rds"
)

// ClusterExternalNameConfigure configures name of cluster.
func ClusterExternalNameConfigure(base map[string]interface{}, name string) {
	base["cluster_identifier"] = name
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
				Type:              "github.com/crossplane-contrib/provider-tf-aws/apis/s3/v1alpha1.Bucket",
				RefFieldName:      "BucketRef",
				SelectorFieldName: "BucketSelector",
			},
			"vpc_security_group_ids": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/vpc/v1alpha1.SecurityGroup",
			},
			"restore_to_point_in_time[*].source_cluster_identifier": {
				Type: "Cluster",
			},
		},
		UseAsync: true,
	})
}
