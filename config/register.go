package config

import (
	_ "github.com/crossplane-contrib/provider-tf-aws/config/autoscaling" //nolint:golint
	_ "github.com/crossplane-contrib/provider-tf-aws/config/ebs"         //nolint:golint
	_ "github.com/crossplane-contrib/provider-tf-aws/config/ec2"         //nolint:golint
	_ "github.com/crossplane-contrib/provider-tf-aws/config/ecs"         //nolint:golint
	_ "github.com/crossplane-contrib/provider-tf-aws/config/eks"         //nolint:golint
	_ "github.com/crossplane-contrib/provider-tf-aws/config/elasticache" //nolint:golint
	_ "github.com/crossplane-contrib/provider-tf-aws/config/iam"         //nolint:golint
	_ "github.com/crossplane-contrib/provider-tf-aws/config/kms"         //nolint:golint
	_ "github.com/crossplane-contrib/provider-tf-aws/config/rds"         //nolint:golint
	_ "github.com/crossplane-contrib/provider-tf-aws/config/s3"          //nolint:golint
	_ "github.com/crossplane-contrib/provider-tf-aws/config/vpc"         //nolint:golint
)
