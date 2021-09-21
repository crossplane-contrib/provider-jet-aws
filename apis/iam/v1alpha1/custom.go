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

package v1alpha1

import "github.com/crossplane-contrib/terrajet/pkg/terraform/resource"

var (
	IAMRoleExternalNamer = resource.ExternalNameConfiguration{
		SelfVarPath: "github.com/crossplane-contrib/provider-tf-aws/apis/iam/v1alpha1.IAMRoleExternalNamer",
		Configure: func(base map[string]interface{}, name string) {
			base["name"] = name
		},
		OmittedFields: []string{
			"name",
			"name_prefix",
		},
		// Set to false explicitly. For this resource, metadata.name is suitable
		// to be the default for external name.
		DisableNameInitializer: false,
	}
	IAMUserExternalNamer = resource.ExternalNameConfiguration{
		SelfVarPath: "github.com/crossplane-contrib/provider-tf-aws/apis/iam/v1alpha1.IAMUserExternalNamer",
		Configure: func(base map[string]interface{}, name string) {
			base["name"] = name
		},
		OmittedFields: []string{
			"name",
		},
		// Set to false explicitly. For this resource, metadata.name is suitable
		// to be the default for external name.
		DisableNameInitializer: false,
	}
)
