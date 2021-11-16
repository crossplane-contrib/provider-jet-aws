package ecrpublic

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_ecrpublic_repository", func(r *config.Resource) {
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["repository_name"] = name
			},
			OmittedFields: []string{
				"repository_name",
			},
		}
	})
}
