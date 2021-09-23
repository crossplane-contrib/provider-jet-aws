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

import (
	"github.com/crossplane-contrib/terrajet/pkg/terraform/resource"
	"github.com/crossplane-contrib/terrajet/pkg/types"
)

func iamExternalNameConfigure(base map[string]interface{}, externalName string) {
	base["name"] = externalName
}

var iamRoleCustomConfig = &resource.Configuration{
	ExternalName: resource.ExternalNameConfiguration{
		ConfigureFunction: "iamExternalNameConfigure",
		OmittedFields: []string{
			"name",
			"name_prefix",
		},
		// Set to false explicitly. For this resource, metadata.name is suitable
		// to be the default for external name.
		DisableNameInitializer: false,
	},
	Reference: map[string]resource.FieldReferenceConfiguration{
		"ManagedPolicyArns": {
			ReferenceToType: types.PathForType(IamPolicy{}),
			// TODO(hasan): fix below
			ReferenceExtractor: "github.com/crossplane/provider-aws/apis/ec2/v1beta1.SubnetARN()",
		},
	},
}

var iamUserCustomConfig = &resource.Configuration{
	ExternalName: resource.ExternalNameConfiguration{
		ConfigureFunction: "iamExternalNameConfigure",
		OmittedFields: []string{
			"name",
		},
		// Set to false explicitly. For this resource, metadata.name is suitable
		// to be the default for external name.
		DisableNameInitializer: false,
	},
}

var ConfigStoreBuilder resource.ConfigStoreBuilder

func init() {
	ConfigStoreBuilder.Register(func(c *resource.ConfigStore) error {
		c.SetConfigForResource("aws_iam_user", iamUserCustomConfig)
		c.SetConfigForResource("aws_iam_role", iamRoleCustomConfig)
		return nil
	})
}
