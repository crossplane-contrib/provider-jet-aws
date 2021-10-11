package ecs

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-tf-aws/config/common"
)

func init() {

	config.Store.SetForResource("aws_ecs_cluster", config.Resource{
		ExternalName: config.ExternalName{
			// Note(turkenh): Seems like we have a case that breaks our
			// assumption that "id" field always used to import resource and
			// should be set as external-id. Here we could import the resource
			// with "name" but id contains "arn".
			// https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecs_cluster
			ConfigureFunctionPath: common.PathExternalNameAsName,
			OmittedFields: []string{
				"name",
			},
		},
		References: config.References{
			"capacity_providers": config.Reference{
				Type: "CapacityProvider",
			},
			"execute_command_configuration[*].kms_key_id": config.Reference{
				Type:              "github.com/crossplane-contrib/provider-tf-aws/apis/kms/v1alpha1.Key",
				RefFieldName:      "KmsKeyRef",
				SelectorFieldName: "KmsKeySelector",
			},
			"log_configuration[*].s3_bucket_name": config.Reference{
				Type:              "github.com/crossplane-contrib/provider-tf-aws/apis/s3/v1alpha1.Bucket",
				RefFieldName:      "BucketRef",
				SelectorFieldName: "BucketSelector",
			},
		},
		UseAsync: true,
	})

	config.Store.SetForResource("aws_ecs_service", config.Resource{
		ExternalName: config.ExternalName{
			ConfigureFunctionPath: common.PathExternalNameAsName,
			OmittedFields: []string{
				"name",
			},
		},
		References: config.References{
			"cluster": config.Reference{
				Type:      "Cluster",
				Extractor: common.PathARNExtractor,
			},
			"iam_role": config.Reference{
				Type:      "github.com/crossplane-contrib/provider-tf-aws/apis/iam/v1alpha1.Role",
				Extractor: common.PathARNExtractor,
			},
			"network_configuration[*].Subnets": config.Reference{
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/vpc/v1alpha1.Subnet",
			},
			"network_configuration[*].security_groups": config.Reference{
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/vpc/v1alpha1.SecurityGroup",
			},
		},

		UseAsync: true,
	})

	config.Store.SetForResource("aws_ecs_capacity_provider", config.Resource{
		ExternalName: config.ExternalName{
			ConfigureFunctionPath: common.PathExternalNameAsName,
			OmittedFields: []string{
				"name",
			},
		},
		References: config.References{
			"auto_scaling_group_provider[*].auto_scaling_group_arn": config.Reference{
				Type:      "github.com/crossplane-contrib/provider-tf-aws/apis/autoscaling/v1alpha1.AutoscalingGroup",
				Extractor: common.PathARNExtractor,
			},
		},
	})

	config.Store.SetForResource("aws_ecs_tag", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			// Note(turkenh): This reference could correspond multiple types as
			// per documentation, any ecs resource: https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecs_tag#resource_arn
			// However, we could only reference to one type.
			"resource_arn": config.Reference{
				Type:      "Cluster",
				Extractor: common.PathARNExtractor,
			},
		},
	})

	config.Store.SetForResource("aws_ecs_task_definition", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"execution_role_arn": config.Reference{
				Type:      "github.com/crossplane-contrib/provider-tf-aws/apis/iam/v1alpha1.Role",
				Extractor: common.PathARNExtractor,
			},
		},
	})
}
