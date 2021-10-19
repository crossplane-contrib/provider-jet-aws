package ecr

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-tf-aws/config/common"
)

func init() {

	config.Store.SetForResource("aws_ecr_repository", config.Resource{
		ExternalName: config.NameAsIdentifier,
		References: map[string]config.Reference{
			"encryption_configuration[*].kms_key": {
				Type:      "github.com/crossplane-contrib/provider-tf-aws/apis/kms/v1alpha1.Key",
				Extractor: common.PathARNExtractor,
			},
		},

		// Note(turkenh): only delete has a timeout of 20 mins, we only need to
		// use async with destroy:
		// https://github.com/crossplane-contrib/terrajet/issues/102
		UseAsync: true,
	})

	config.Store.SetForResource("aws_ecrpublic_repository", config.Resource{
		ExternalName: config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["repository_name"] = name
			},
			OmittedFields: []string{
				"repository_name",
			},
		},

		// Note(turkenh): only delete has a timeout of 20 mins, we only need to
		// use async with destroy:
		// https://github.com/crossplane-contrib/terrajet/issues/102
		UseAsync: true,
	})
}
