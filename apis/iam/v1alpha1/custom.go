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
	"github.com/crossplane-contrib/terrajet/pkg/config"
)

func iamExternalNameConfigure(base map[string]interface{}, externalName string) {
	base["name"] = externalName
}

var iamRoleCustomConfig = config.Resource{
	ExternalName: config.ExternalName{
		ConfigureFunctionPath: "iamExternalNameConfigure",
		OmittedFields: []string{
			"name",
			"name_prefix",
		},
		// Set to false explicitly. For this resource, metadata.name is suitable
		// to be the default for external name.
		DisableNameInitializer: false,
	},
}

var iamUserCustomConfig = config.Resource{
	ExternalName: config.ExternalName{
		ConfigureFunctionPath: "iamExternalNameConfigure",
		OmittedFields: []string{
			"name",
		},
		// Set to false explicitly. For this resource, metadata.name is suitable
		// to be the default for external name.
		DisableNameInitializer: false,
	},
}

var iamAccessKeyCustomConfig = config.Resource{
	ExternalName: config.ExternalName{
		DisableNameInitializer: true,
	},
	References: config.References{
		"user": config.Reference{
			Type: "IamUser",
		},
	},
}

func init() {
	config.Store.SetForResource("aws_iam_user", iamUserCustomConfig)
	config.Store.SetForResource("aws_iam_role", iamRoleCustomConfig)
	config.Store.SetForResource("aws_iam_access_key", iamAccessKeyCustomConfig)
}
