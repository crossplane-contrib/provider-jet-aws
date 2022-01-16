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

package neptune

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/crossplane/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-jet-aws/config/common"
)

// Configure adds configurations for neptune group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_neptune_cluster", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
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
	})

	p.AddResourceConfigurator("aws_neptune_cluster_endpoint", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["cluster_endpoint_identifier"] = name
			},

			OmittedFields: []string{
				"cluster_endpoint_identifier",
			},

			GetExternalNameFn: func(tfstate map[string]interface{}) (string, error) {
				id, ok := tfstate["id"].(string)
				if !ok || id == "" {
					return "", errors.New("cannot get id from tfstate")
				}

				// my_cluster:my_cluster_endpoint
				w := strings.Split(id, ":")
				if len(w) != 2 {
					return "", errors.New("format of id should be my_cluster:my_cluster_endpoint")
				}
				return w[len(w)-1], nil
			},

			GetIDFn: func(_ context.Context, externalName string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
				ci, ok := parameters["cluster_identifier"].(string)
				if !ok || ci == "" {
					return "", errors.New("cannot get cluster_identifier from parameters")
				}
				return fmt.Sprintf("%s:%s", ci, externalName), nil
			},
		}

		r.References = config.References{
			"cluster_identifier": config.Reference{
				Type: "Cluster",
			},
		}
	})

	p.AddResourceConfigurator("aws_neptune_cluster_instance", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
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

		r.UseAsync = true

		r.References = config.References{
			"cluster_identifier": config.Reference{
				Type: "Cluster",
			},
		}
	})

	p.AddResourceConfigurator("aws_neptune_cluster_parameter_group", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["name"] = name
			},

			OmittedFields: []string{
				"name",
				"name_prefix",
			},

			GetExternalNameFn: config.IDAsExternalName,
			GetIDFn:           config.ExternalNameAsID,
		}
	})

	p.AddResourceConfigurator("aws_neptune_cluster_snapshot", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["db_cluster_snapshot_identifier"] = name
			},

			OmittedFields: []string{
				"db_cluster_snapshot_identifier",
			},

			GetExternalNameFn: config.IDAsExternalName,
			GetIDFn:           config.ExternalNameAsID,
		}

		r.UseAsync = true

		r.References = config.References{
			"db_cluster_identifier": config.Reference{
				Type: "Cluster",
			},
		}
	})

	p.AddResourceConfigurator("aws_neptune_event_subscription", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["name"] = name
			},

			OmittedFields: []string{
				"name",
				"name_prefix",
			},

			GetExternalNameFn: config.IDAsExternalName,
			GetIDFn:           config.ExternalNameAsID,
		}
	})

	p.AddResourceConfigurator("aws_neptune_parameter_group", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["name"] = name
			},

			OmittedFields: []string{
				"name",
				"name_prefix",
			},

			GetExternalNameFn: config.IDAsExternalName,
			GetIDFn:           config.ExternalNameAsID,
		}
	})

	p.AddResourceConfigurator("aws_neptune_subnet_group", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["name"] = name
			},

			OmittedFields: []string{
				"name",
				"name_prefix",
			},

			GetExternalNameFn: config.IDAsExternalName,
			GetIDFn:           config.ExternalNameAsID,
		}

		r.References = config.References{
			"subnet_ids": {
				Type:              "github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2.Subnet",
				RefFieldName:      "SubnetIdRefs",
				SelectorFieldName: "SubnetIdSelector",
			},
		}
	})
}
