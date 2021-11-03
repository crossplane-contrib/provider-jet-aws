package config

import (
	"strings"

	tjconfig "github.com/crossplane-contrib/terrajet/pkg/config"
	"github.com/crossplane-contrib/terrajet/pkg/types/comments"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	tf "github.com/terraform-providers/terraform-provider-aws/aws"

	"github.com/crossplane-contrib/provider-tf-aws/config/autoscaling"
	"github.com/crossplane-contrib/provider-tf-aws/config/ec2"
)

const resourcePrefix = "aws"

var regionSchema = getRegionSchema()

var includedResources = []string{

	/*	// VPC
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

		// S3
		"aws_s3_bucket$",

		// Elasticache
		"aws_elasticache_cluster$",
		"aws_elasticache_parameter_group$",
		"aws_elasticache_replication_group$",

		// ECS
		"aws_ecs_cluster$",
		"aws_ecs_service$",
		"aws_ecs_capacity_provider$",
		"aws_ecs_tag$",
		"aws_ecs_task_definition$",
	*/

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
	/*
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
	*/
}

// These resources cannot be generated because of their suffixes colliding with
// kubebuilder-accepted type suffixes.
var skipList = []string{
	"aws_config_configuration_recorder_status$",
	"aws_vpc_peering_connection_accepter$",
	"aws_elasticache_user_group$",
	"aws_waf_rule_group$",
	"aws_wafregional_rule_group$",
	"aws_shield_protection_group$",
	"aws_route53_resolver_firewall_rule_group$",
	"aws_kinesis_analytics_application$",
}

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	pc := tjconfig.NewProvider(
		tf.Provider().ResourcesMap, resourcePrefix, "github.com/crossplane-contrib/provider-tf-aws",
		tjconfig.WithIncludeList(includedResources),
		tjconfig.WithSkipList(skipList),
	)

	for name, resource := range pc.Resources {
		switch {
		case strings.HasPrefix(name, "aws_security_group") ||
			strings.HasPrefix(name, "aws_subnet") ||
			strings.HasPrefix(name, "aws_network") ||
			strings.HasPrefix(name, "aws_eip") ||
			strings.HasPrefix(name, "aws_launch_template") ||
			strings.HasPrefix(name, "aws_route") ||
			strings.HasPrefix(name, "aws_instance") ||
			strings.HasPrefix(name, "aws_vpc"):
			pc.AddResourceConfigurator(name, resourceGroupConfigurator("ec2"))
		case resource.Group == "vpc":
			pc.AddResourceConfigurator(name, resourceGroupConfigurator("ec2"))
		case name == "aws_lb":
			pc.AddResourceConfigurator(name, resourceGroupConfigurator("lb"))
		case strings.Contains(name, "aws_db_"):
			pc.AddResourceConfigurator(name, resourceGroupConfigurator("rds"))
		case strings.Contains(name, "aws_ecrpublic_repository"):
			pc.AddResourceConfigurator(name, resourceGroupConfigurator("ecr"))
		}

		if resource.Group != "iam" {
			pc.AddResourceConfigurator(name, func(r *tjconfig.Resource) {
				r.TerraformResource.Schema["region"] = regionSchema
			})
		}

		// tags_all is used only in tfstate to accumulate provider-wide
		// default tags in TF, which is not something we support. So, we don't
		// need it as a parameter while "tags" is already in place.
		pc.AddResourceConfigurator(name, func(r *tjconfig.Resource) {
			if t, ok := r.TerraformResource.Schema["tags_all"]; ok {
				t.Computed = true
				t.Optional = false
			}
		})
	}

	for _, configure := range []func(provider *tjconfig.Provider){
		// add custom config functions
		autoscaling.Configure,
		ec2.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

func getRegionSchema() *schema.Schema {
	c := "Region is the region you'd like your resource to be created in.\n"

	comment, err := comments.New(c, comments.WithTFTag("-"))
	if err != nil {
		panic(errors.Wrap(err, "cannot build comment for region"))
	}
	return &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: comment.String(),
	}
}

func resourceGroupConfigurator(group string) tjconfig.ResourceConfiguratorFn {
	return func(r *tjconfig.Resource) {
		r.Group = group
	}
}
