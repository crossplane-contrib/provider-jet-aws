/*
Copyright 2022 The Crossplane Authors.

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

package grafana

import (
	"github.com/crossplane/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-jet-aws/config/common"
)

// Configure adds configurations for grafana group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_grafana_workspace", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
	})

	p.AddResourceConfigurator("aws_grafana_workspace_saml_configuration", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
		r.References = config.References{
			"workspace_id": config.Reference{
				Type: "Workspace",
			},
		}
	})

	p.AddResourceConfigurator("aws_grafana_role_association", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
		r.References = config.References{
			"workspace_id": config.Reference{
				Type: "Workspace",
			},
		}
	})
}
