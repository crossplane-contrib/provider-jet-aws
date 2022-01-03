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

package s3

import (
	"github.com/crossplane/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-jet-aws/config/common"
)

// Configure adds configurations for s3 group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_rds_cluster", func(r *config.Resource) {
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["bucket"] = name
			},
			OmittedFields: []string{
				"bucket",
				"bucket_prefix",
			},
			GetExternalNameFn: config.IDAsExternalName,
			GetIDFn:           config.ExternalNameAsID,
		}
		r.References = config.References{
			"server_side_encryption_configuration.rule.apply_server_side_encryption_by_default.kms_master_key_id": {
				Type:      "github.com/crossplane-contrib/provider-jet-aws/apis/kms/v1alpha1.Key",
				Extractor: common.PathARNExtractor,
			},
		}
	})
}
