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
