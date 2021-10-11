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
	"github.com/crossplane-contrib/provider-tf-aws/config/common"
	"github.com/crossplane-contrib/terrajet/pkg/config"
)

// Package path constants.
const (
	SelfPkgPath = "github.com/crossplane-contrib/provider-tf-aws/config/s3"
)

// BucketExternalNameConfigure configures bucket name.
func BucketExternalNameConfigure(base map[string]interface{}, name string) {
	base["bucket"] = name
}

func init() {
	config.Store.SetForResource("aws_s3_bucket", config.Resource{
		ExternalName: config.ExternalName{
			ConfigureFunctionPath: SelfPkgPath + ".BucketExternalNameConfigure",
			OmittedFields: []string{
				"bucket",
				"bucket_prefix",
			},
		},
		References: map[string]config.Reference{
			"server_side_encryption_configuration[*].rule[*].apply_server_side_encryption_by_default[*].kms_master_key_id": {
				Type:      "github.com/crossplane-contrib/provider-tf-aws/apis/kms/v1alpha1.Key",
				Extractor: common.PathARNExtractor,
			},
		},
	})
}
