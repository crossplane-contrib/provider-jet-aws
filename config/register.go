package config

import (
	_ "github.com/crossplane-contrib/provider-tf-aws/config/eks"         //nolint:golint
	_ "github.com/crossplane-contrib/provider-tf-aws/config/elasticache" //nolint:golint
	_ "github.com/crossplane-contrib/provider-tf-aws/config/iam"         //nolint:golint
	_ "github.com/crossplane-contrib/provider-tf-aws/config/rds"         //nolint:golint
	_ "github.com/crossplane-contrib/provider-tf-aws/config/s3"          //nolint:golint
	_ "github.com/crossplane-contrib/provider-tf-aws/config/vpc"         //nolint:golint
)
