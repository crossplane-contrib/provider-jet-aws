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

package elasticache

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"
)

const (
	CommonPkgPath = "github.com/crossplane-contrib/provider-tf-aws/config/common"
	SelfPkgPath   = "github.com/crossplane-contrib/provider-tf-aws/config/elasticache"
)

func ClusterExternalNameConfigure(base map[string]interface{}, name string) {
	base["cluster_id"] = name
}

func ReplicationGroupExternalNameConfigure(base map[string]interface{}, name string) {
	base["replication_group_id"] = name
}

func init() {

	config.Store.SetForResource("aws_elasticache_parameter_group", config.Resource{
		ExternalName: config.ExternalName{
			ConfigureFunctionPath: CommonPkgPath + ".ExternalNameAsName",
			OmittedFields: []string{
				"name",
			},
		},
	})

	config.Store.SetForResource("aws_elasticache_cluster", config.Resource{
		ExternalName: config.ExternalName{
			ConfigureFunctionPath: SelfPkgPath + ".ClusterExternalNameConfigure",
			OmittedFields: []string{
				"cluster_id",
			},
		},
		References: config.References{
			"parameter_group_name": config.Reference{
				Type:              "ParameterGroup",
				RefFieldName:      "ParameterGroupRef",
				SelectorFieldName: "ParameterGroupSelector",
			},
		},
		UseAsync: true,
	})

	config.Store.SetForResource("aws_elasticache_replication_group", config.Resource{
		ExternalName: config.ExternalName{
			ConfigureFunctionPath: SelfPkgPath + ".ReplicationGroupExternalNameConfigure",
		},
	})
}
