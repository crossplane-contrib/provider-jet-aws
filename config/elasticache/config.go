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

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_elasticache_parameter_group", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
	})

	p.AddResourceConfigurator("aws_elasticache_cluster", func(r *config.Resource) {
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["cluster_id"] = name
			},
			OmittedFields: []string{
				"cluster_id",
			},
		}
		r.References = config.References{
			"parameter_group_name": config.Reference{
				Type: "ParameterGroup",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_elasticache_replication_group", func(r *config.Resource) {
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["replication_group_id"] = name
			},
			OmittedFields: []string{
				"replication_group_id",
			},
		}
	})
}
