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

package iam

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
)

var SelfPgkPath = "github.com/crossplane-contrib/provider-tf-aws/config/iam"

func ExternalNameAsName(base map[string]interface{}, externalName string) {
	base["name"] = externalName
}

func PolicyARNExtractor() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			// todo(hasan): should we log this error?
			return ""
		}
		r, err := paved.GetString("status.atProvider.arn")
		if err != nil {
			// todo(hasan): should we log this error?
			return ""
		}
		return r
	}
}

func init() {
	config.Store.SetForResource("aws_iam_access_key", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"user": config.Reference{
				Type: "IamUser",
			},
		},
	})

	config.Store.SetForResource("aws_iam_group", config.Resource{
		Kind: "IAMGroup",
		ExternalName: config.ExternalName{
			ConfigureFunctionPath: "iam.ExternalNameAsName",
			OmittedFields: []string{
				"name",
			},
		},
	})

	config.Store.SetForResource("aws_iam_group_policy", config.Resource{
		ExternalName: config.ExternalName{
			ConfigureFunctionPath: "iam.ExternalNameAsName",
			OmittedFields: []string{
				"name",
				"name_prefix",
			},
		},
		References: config.References{
			"group": config.Reference{
				Type: "IamGroup",
			},
		},
	})

	config.Store.SetForResource("aws_iam_group_policy_attachment", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"group": config.Reference{
				Type: "IamGroup",
			},
			"policy_arn": config.Reference{
				Type:      "IamPolicy",
				Extractor: SelfPgkPath + ".PolicyARNExtractor()",
			},
		},
	})

	config.Store.SetForResource("aws_iam_instance_profile", config.Resource{
		ExternalName: config.ExternalName{
			ConfigureFunctionPath: "iam.ExternalNameAsName",
			OmittedFields: []string{
				"name",
				"name_prefix",
			},
		},
		References: config.References{
			"role": config.Reference{
				Type: "IamRole",
			},
		},
	})

	config.Store.SetForResource("aws_iam_policy", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
	})

	config.Store.SetForResource("aws_iam_policy_attachment", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"users": config.Reference{
				Type: "IamUser",
			},
			"roles": config.Reference{
				Type: "IamRole",
			},
			"groups": config.Reference{
				Type: "IamGroup",
			},
			"policy_arn": config.Reference{
				Type:      "IamPolicy",
				Extractor: SelfPgkPath + ".PolicyARNExtractor()",
			},
		},
	})

	config.Store.SetForResource("aws_iam_role", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: false,
			ConfigureFunctionPath:  "iam.ExternalNameAsName",
			OmittedFields: []string{
				"name",
				"name_prefix",
			},
		},
	})

	config.Store.SetForResource("aws_iam_role_policy", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: false,
			ConfigureFunctionPath:  "iam.ExternalNameAsName",
			OmittedFields: []string{
				"name",
				"name_prefix",
			},
		},
	})
	config.Store.SetForResource("aws_iam_role_policy_attachment", config.Resource{})

	config.Store.SetForResource("aws_iam_user", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: false,
			ConfigureFunctionPath:  "iam.ExternalNameAsName",
			OmittedFields: []string{
				"name",
			},
		},
	})

	config.Store.SetForResource("aws_iam_user_group_membership", config.Resource{})
	config.Store.SetForResource("aws_iam_user_policy", config.Resource{})
	config.Store.SetForResource("aws_iam_user_policy_attachment", config.Resource{})
}
