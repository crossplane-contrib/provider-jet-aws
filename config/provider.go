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

package config

import (
	// Note(ezgidemirel): we are importing this to embed provider schema document
	_ "embed"

	"github.com/crossplane-contrib/provider-jet-aws/config/mq"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	tjconfig "github.com/crossplane/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-jet-aws/config/autoscaling"
	"github.com/crossplane-contrib/provider-jet-aws/config/ebs"
	"github.com/crossplane-contrib/provider-jet-aws/config/ec2"
	"github.com/crossplane-contrib/provider-jet-aws/config/ecr"
	"github.com/crossplane-contrib/provider-jet-aws/config/ecrpublic"
	"github.com/crossplane-contrib/provider-jet-aws/config/ecs"
	"github.com/crossplane-contrib/provider-jet-aws/config/eks"
	"github.com/crossplane-contrib/provider-jet-aws/config/elasticache"
	"github.com/crossplane-contrib/provider-jet-aws/config/elasticloadbalancing"
	"github.com/crossplane-contrib/provider-jet-aws/config/globalaccelerator"
	"github.com/crossplane-contrib/provider-jet-aws/config/iam"
	"github.com/crossplane-contrib/provider-jet-aws/config/kms"
	"github.com/crossplane-contrib/provider-jet-aws/config/neptune"
	"github.com/crossplane-contrib/provider-jet-aws/config/rds"
	"github.com/crossplane-contrib/provider-jet-aws/config/route53"
	"github.com/crossplane-contrib/provider-jet-aws/config/s3"
)

//go:embed schema.json
var providerSchema string

// IncludedResources lists all resource patterns included in small set release.
var IncludedResources = []string{
	// Elastic Load Balancing v2 (ALB/NLB)
	"aws_lb$",
	"aws_lb_listener$",
	"aws_lb_target_group$",
	"aws_lb_target_group_attachment$",

	// ECR
	"aws_ecr_repository$",
	"aws_ecrpublic_repository$",

	// RDS
	"aws_rds_cluster$",
	"aws_db_instance$",
	"aws_db_parameter_group$",
	"aws_db_subnet_group$",

	// S3
	"aws_s3_bucket$",

	// Elasticache
	"aws_elasticache_cluster$",
	"aws_elasticache_subnet_group$",
	"aws_elasticache_parameter_group$",
	"aws_elasticache_replication_group$",
	"aws_elasticache_user$",
	"aws_elasticache_user_group$",

	// ECS
	"aws_ecs_cluster$",
	"aws_ecs_service$",
	"aws_ecs_capacity_provider$",
	"aws_ecs_tag$",
	"aws_ecs_task_definition$",

	// Autoscaling
	"aws_autoscaling_group$",
	"aws_autoscaling_attachment$",

	// EC2
	"aws_instance$",
	"aws_eip$",
	"aws_launch_template$",
	"aws_ec2_transit_gateway$",
	"aws_ec2_transit_gateway_route$",
	"aws_ec2_transit_gateway_route_table$",
	"aws_ec2_transit_gateway_route_table_association$",
	"aws_ec2_transit_gateway_vpc_attachment$",
	"aws_ec2_transit_gateway_vpc_attachment_accepter$",
	"aws_ec2_transit_gateway_route_table_propagation$",
	"aws_vpc$",
	"aws_security_group$",
	"aws_security_group_rule$",
	"aws_subnet$",
	"aws_network_interface$",
	"aws_route$",
	"aws_route_table$",
	"aws_vpc_endpoint$",
	"aws_vpc_ipv4_cidr_block_association$",
	"aws_vpc_peering_connection$",
	"aws_route_table_association$",
	"aws_internet_gateway$",
	"aws_nat_gateway$",

	// IAM
	"aws_iam_access_key$",
	"aws_iam_group$",
	"aws_iam_group_policy$",
	"aws_iam_group_policy_attachment$",
	"aws_iam_instance_profile$",
	"aws_iam_policy$",
	"aws_iam_policy_attachment$",
	"aws_iam_role$",
	"aws_iam_role_policy$",
	"aws_iam_role_policy_attachment$",
	"aws_iam_user$",
	"aws_iam_user_group_membership$",
	"aws_iam_user_policy$",
	"aws_iam_user_policy_attachment$",
	"aws_iam_openid_connect_provider$",

	// EKS
	"aws_eks_addon$",
	"aws_eks_cluster$",
	"aws_eks_fargate_profile$",
	"aws_eks_node_group$",
	"aws_eks_identity_provider_config$",

	// KMS
	"aws_kms_key$",

	// EBS
	"aws_ebs_volume$",

	// Route53
	"aws_route53_.*",

	// Neptune
	"aws_neptune_cluster$",
	"aws_neptune_cluster_endpoint$",
	"aws_neptune_cluster_instance$",
	"aws_neptune_cluster_parameter_group$",
	"aws_neptune_cluster_snapshot$",
	"aws_neptune_event_subscription",
	"aws_neptune_parameter_group$",
	"aws_neptune_subnet_group$",

	// MQ
	"aws_mq_broker$",
	"aws_mq_configuration$",

	// Global Accelerator
	"aws_globalaccelerator_accelerator",
	"aws_globalaccelerator_endpoint_group",
	"aws_globalaccelerator_listener",
}

var skipList = []string{
	"aws_waf_rule_group$",              // Too big CRD schema
	"aws_wafregional_rule_group$",      // Too big CRD schema
	"aws_glue_connection$",             // See https://github.com/crossplane-contrib/terrajet/issues/100
	"aws_mwaa_environment$",            // See https://github.com/crossplane-contrib/terrajet/issues/100
	"aws_ecs_tag$",                     // tags are already managed by ecs resources.
	"aws_alb$",                         // identical with aws_lb
	"aws_alb_target_group_attachment$", // identical with aws_lb_target_group_attachment
	"aws_iam_policy_attachment$",       // identical with aws_iam_*_policy_attachment resources.
	"aws_iam_group_policy$",            // identical with aws_iam_*_policy_attachment resources.
	"aws_iam_role_policy$",             // identical with aws_iam_*_policy_attachment resources.
	"aws_iam_user_policy$",             // identical with aws_iam_*_policy_attachment resources.
}

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	pc := tjconfig.NewProviderWithSchema([]byte(providerSchema), "aws", "github.com/crossplane-contrib/provider-jet-aws",
		tjconfig.WithShortName("awsjet"),
		tjconfig.WithRootGroup("aws.jet.crossplane.io"),
		tjconfig.WithIncludeList(IncludedResources),
		tjconfig.WithSkipList(skipList),
		tjconfig.WithDefaultResourceFn(DefaultResource(
			GroupKindOverrides(),
			KindOverrides(),
			RegionAddition(),
			TagsAllRemoval(),
			IdentifierAssignedByAWS(),
			NamePrefixRemoval(),
			KnownReferencers(),
			AddExternalTagsField(),
		)),
	)

	for _, configure := range []func(provider *tjconfig.Provider){
		autoscaling.Configure,
		ebs.Configure,
		ec2.Configure,
		ecr.Configure,
		ecrpublic.Configure,
		ecs.Configure,
		eks.Configure,
		elasticache.Configure,
		elasticloadbalancing.Configure,
		globalaccelerator.Configure,
		iam.Configure,
		kms.Configure,
		rds.Configure,
		s3.Configure,
		route53.Configure,
		neptune.Configure,
		mq.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// DefaultResource returns a DefaultResoruceFn that makes sure the original
// DefaultResource call is made with given options here.
func DefaultResource(opts ...tjconfig.ResourceOption) tjconfig.DefaultResourceFn {
	return func(name string, terraformResource *schema.Resource, orgOpts ...tjconfig.ResourceOption) *tjconfig.Resource {
		return tjconfig.DefaultResource(name, terraformResource, append(orgOpts, opts...)...)
	}
}
