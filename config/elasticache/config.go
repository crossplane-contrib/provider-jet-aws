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
	"github.com/crossplane/terrajet/pkg/config"
)

// Configure adds configurations for elasticache group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_elasticache_parameter_group", func(r *config.Resource) {
		r.Version = "v1alpha2"
		r.ExternalName = config.NameAsIdentifier
	})

	p.AddResourceConfigurator("aws_elasticache_cluster", func(r *config.Resource) {
		r.Version = "v1alpha2"
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["cluster_id"] = name
			},
			OmittedFields: []string{
				"cluster_id",
			},
			GetExternalNameFn: config.IDAsExternalName,
			GetIDFn:           config.ExternalNameAsID,
		}
		r.References = config.References{
			"parameter_group_name": config.Reference{
				Type: "ParameterGroup",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_elasticache_replication_group", func(r *config.Resource) {
		r.Version = "v1alpha2"
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["replication_group_id"] = name
			},
			OmittedFields: []string{
				"replication_group_id",
			},
			GetExternalNameFn: config.IDAsExternalName,
			GetIDFn:           config.ExternalNameAsID,
		}
	})

	p.AddResourceConfigurator("aws_elasticache_user", func(r *config.Resource) {
		r.Version = "v1alpha2"
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["user_id"] = name
			},
			OmittedFields: []string{
				"user_id",
			},
			GetExternalNameFn: config.IDAsExternalName,
			GetIDFn:           config.ExternalNameAsID,
		}
	})

	p.AddResourceConfigurator("aws_elasticache_user_group", func(r *config.Resource) {
		r.Version = "v1alpha2"
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["user_group_id"] = name
			},
			OmittedFields: []string{
				"user_group_id",
			},
			GetExternalNameFn: config.IDAsExternalName,
			GetIDFn:           config.ExternalNameAsID,
		}
		r.References = config.References{
			"user_ids": config.Reference{
				Type:              "User",
				RefFieldName:      "UserIdRefs",
				SelectorFieldName: "UserIdSelector",
			},
		}
	})
}
