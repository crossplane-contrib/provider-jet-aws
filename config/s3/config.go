package s3

import "github.com/crossplane-contrib/terrajet/pkg/config"

var SelfPkgPath = "github.com/crossplane-contrib/provider-tf-aws/config/s3"

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
	})
}
