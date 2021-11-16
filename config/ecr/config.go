package ecr

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-tf-aws/config/common"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_ecr_repository", func(r *config.Resource) {
		r.References = map[string]config.Reference{
			"encryption_configuration[*].kms_key": {
				Type:      "github.com/crossplane-contrib/provider-tf-aws/apis/kms/v1alpha1.Key",
				Extractor: common.PathARNExtractor,
			},
		}
	})
}
