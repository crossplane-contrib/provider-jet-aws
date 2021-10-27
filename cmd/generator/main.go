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

package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
	"github.com/terraform-providers/terraform-provider-aws/aws"

	"github.com/crossplane-contrib/terrajet/pkg/config"
	"github.com/crossplane-contrib/terrajet/pkg/pipeline"
	"github.com/crossplane-contrib/terrajet/pkg/types/comments"

	// This is to get init functions registering custom configs to be called.
	_ "github.com/crossplane-contrib/provider-tf-aws/config"
)

// Constants to use in generated artifacts.
const (
	modulePath        = "github.com/crossplane-contrib/provider-tf-aws"
	groupSuffix       = ".aws.tf.crossplane.io"
	providerShortName = "tfaws"
)

// These resources cannot be generated because of their suffixes colliding with
// kubebuilder-accepted type suffixes.
var skipList = map[string]struct{}{
	"aws_config_configuration_recorder_status": {}, // collision
	"aws_vpc_peering_connection_accepter":      {},
	"aws_elasticache_user_group":               {},
	"aws_waf_rule_group":                       {},
	"aws_wafregional_rule_group":               {},
	"aws_shield_protection_group":              {},
	"aws_route53_resolver_firewall_rule_group": {},
	"aws_kinesis_analytics_application":        {},
	"aws_ssm_parameter":                        {}, // collision
	"aws_alb_target_group":                     {}, // collision
	"aws_appmesh_virtual_service":              {}, // collision
	"aws_mwaa_environment":                     {}, // sensitive field type
	"aws_mq_configuration":                     {}, // collision
	"aws_lex_intent":                           {}, // collision
	"aws_glue_schema":                          {}, // collision
	"aws_glue_connection":                      {}, // sensitive field type
	"aws_placement_group":                      {}, // collision
	"aws_quicksight_group":                     {}, // collision
	"aws_resourcegroups_group":                 {}, // collision
	"aws_xray_group":                           {}, // collision
	"aws_cloudformation_type":                  {}, // collision with special word type
	"aws_wafv2_web_acl":                        {}, // CRD with 1.9MB size
}

var includedResources = map[string]struct{}{

	// VPC
	"aws_vpc":                             {},
	"aws_security_group":                  {},
	"aws_security_group_rule":             {},
	"aws_subnet":                          {},
	"aws_network_interface":               {},
	"aws_route":                           {},
	"aws_route_table":                     {},
	"aws_vpc_endpoint":                    {},
	"aws_vpc_ipv4_cidr_block_association": {},
	"aws_vpc_peering_connection":          {},
	"aws_route_table_association":         {},

	// Elastic Load Balancing v2 (ALB/NLB)
	"aws_lb":                         {},
	"aws_lb_listener":                {},
	"aws_lb_target_group":            {},
	"aws_lb_target_group_attachment": {},

	// ECR
	"aws_ecr_repository":       {},
	"aws_ecrpublic_repository": {},

	// RDS
	"aws_rds_cluster":        {},
	"aws_db_instance":        {},
	"aws_db_parameter_group": {},

	// S3
	"aws_s3_bucket": {},

	// Elasticache
	"aws_elasticache_cluster":           {},
	"aws_elasticache_parameter_group":   {},
	"aws_elasticache_replication_group": {},

	// ECS
	"aws_ecs_cluster":           {},
	"aws_ecs_service":           {},
	"aws_ecs_capacity_provider": {},
	"aws_ecs_tag":               {},
	"aws_ecs_task_definition":   {},

	// Autoscaling
	"aws_autoscaling_group":      {},
	"aws_autoscaling_attachment": {},

	// IAM
	"aws_iam_access_key":              {},
	"aws_iam_group":                   {},
	"aws_iam_group_policy":            {},
	"aws_iam_group_policy_attachment": {},
	"aws_iam_instance_profile":        {},
	"aws_iam_policy":                  {},
	"aws_iam_policy_attachment":       {},
	"aws_iam_role":                    {},
	"aws_iam_role_policy":             {},
	"aws_iam_role_policy_attachment":  {},
	"aws_iam_user":                    {},
	"aws_iam_user_group_membership":   {},
	"aws_iam_user_policy":             {},
	"aws_iam_user_policy_attachment":  {},

	// EKS
	"aws_eks_addon":                    {},
	"aws_eks_cluster":                  {},
	"aws_eks_fargate_profile":          {},
	"aws_eks_node_group":               {},
	"aws_eks_identity_provider_config": {},

	// EC2
	"aws_instance":                                    {},
	"aws_eip":                                         {},
	"aws_launch_template":                             {},
	"aws_ec2_transit_gateway":                         {},
	"aws_ec2_transit_gateway_route":                   {},
	"aws_ec2_transit_gateway_route_table":             {},
	"aws_ec2_transit_gateway_route_table_association": {},
	"aws_ec2_transit_gateway_vpc_attachment":          {},
	"aws_ec2_transit_gateway_vpc_attachment_accepter": {},
	"aws_ec2_transit_gateway_route_table_propagation": {},

	// KMS
	"aws_kms_key": {},

	// EBS
	"aws_ebs_volume": {},
}

func main() { // nolint:gocyclo
	wd, err := os.Getwd()
	if err != nil {
		panic(errors.Wrap(err, "cannot get working directory"))
	}

	groups := map[string]map[string]*schema.Resource{}
	resources := aws.Provider().ResourcesMap
	keys := make([]string, len(resources))
	i := 0
	for k := range resources {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, name := range keys {
		if _, ok := skipList[name]; ok {
			continue
		}
		resource := resources[name]
		if len(resource.Schema) == 0 {
			// There are resources with no schema, like aws_securityhub_account
			// that we will address later.
			fmt.Printf("Skipping resource %s because it has no schema\n", name)
			continue
		}
		words := strings.Split(name, "_")
		groupName := words[1]
		switch {
		case strings.Contains(name, "aws_security_group") ||
			strings.Contains(name, "aws_subnet") ||
			strings.Contains(name, "aws_network") ||
			strings.Contains(name, "aws_eip") ||
			strings.Contains(name, "aws_launch_template") ||
			strings.Contains(name, "aws_route") ||
			strings.Contains(name, "aws_instance"):
			groupName = "ec2"
		case groupName == "vpc":
			groupName = "ec2"
		case name == "aws_lb":
			groupName = "lb"
		case strings.Contains(name, "aws_db_"):
			groupName = "rds"
		case strings.Contains(name, "aws_ecrpublic_repository"):
			groupName = "ecr"
		}
		if len(groups[groupName]) == 0 {
			groups[groupName] = map[string]*schema.Resource{}
		}
		groups[groupName][name] = resource
	}
	count := 0
	versionPkgList := []string{
		"github.com/crossplane-contrib/provider-tf-aws/apis/v1alpha1",
	}
	controllerPkgList := []string{
		"github.com/crossplane-contrib/provider-tf-aws/internal/controller/tfaws",
	}

	regionSchema, err := getRegionSchema()
	if err != nil {
		panic(errors.Wrap(err, "cannot get region schema"))
	}

	for group, resources := range groups {
		version := "v1alpha1"
		versionGen := pipeline.NewVersionGenerator(wd, modulePath, strings.ToLower(group)+groupSuffix, version)

		crdGen := pipeline.NewCRDGenerator(versionGen.Package(), versionGen.DirectoryPath(), strings.ToLower(group)+groupSuffix, providerShortName)
		tfGen := pipeline.NewTerraformedGenerator(versionGen.Package(), versionGen.DirectoryPath())
		ctrlGen := pipeline.NewControllerGenerator(wd, modulePath, strings.ToLower(group)+groupSuffix)

		keys := make([]string, len(resources))
		i := 0
		for k := range resources {
			keys[i] = k
			i++
		}
		sort.Strings(keys)

		for _, name := range keys {
			// We don't want Aws prefix in all kinds.
			kind := strcase.ToCamel(strings.TrimPrefix(strings.TrimPrefix(name, "aws_"), group))
			if kind == "" {
				kind = strcase.ToCamel(group)
			}
			resource := resources[name]
			if group != "iam" {
				resource.Schema["region"] = regionSchema
			}
			// tags_all is used only in tfstate to accumulate provider-wide
			// default tags in TF, which is not something we support. So, we don't
			// need it as a parameter while "tags" is already in place.
			if t, ok := resource.Schema["tags_all"]; ok {
				t.Computed = true
				t.Optional = false
			}
			rc := config.NewResource(version, kind, name)
			if err = rc.OverrideConfig(config.Store.GetForResource(name)); err != nil {
				panic(errors.Wrap(err, "cannot override config"))
			}

			if err = crdGen.Generate(rc, resource); err != nil {
				panic(errors.Wrap(err, "cannot generate crd"))
			}
			if err = tfGen.Generate(rc, resource); err != nil {
				panic(errors.Wrap(err, "cannot generate terraformed"))
			}
			ctrlPkgPath, err := ctrlGen.Generate(rc, versionGen.Package().Path())
			if err != nil {
				panic(errors.Wrap(err, "cannot generate controller"))
			}
			controllerPkgList = append(controllerPkgList, ctrlPkgPath)
			count++
		}

		if err := versionGen.Generate(); err != nil {
			panic(errors.Wrap(err, "cannot generate version files"))
		}
		versionPkgList = append(versionPkgList, versionGen.Package().Path())
	}

	if err := pipeline.NewRegisterGenerator(wd, modulePath).Generate(versionPkgList); err != nil {
		panic(errors.Wrap(err, "cannot generate register file"))
	}
	if err := pipeline.NewSetupGenerator(wd, modulePath).Generate(controllerPkgList); err != nil {
		panic(errors.Wrap(err, "cannot generate setup file"))
	}
	fmt.Println("running goimports apis")
	if o, err := exec.Command("bash", "-c", "goimports -w $(find apis -iname 'zz_*')").CombinedOutput(); err != nil {
		panic(errors.Wrapf(err, "cannot run goimports for apis folder: %s", string(o)))
	}
	fmt.Println("running goimports internal")
	if o, err := exec.Command("bash", "-c", "goimports -w $(find internal -iname 'zz_*')").CombinedOutput(); err != nil {
		panic(errors.Wrapf(err, "cannot run goimports for internal folder: %s", string(o)))
	}
	fmt.Printf("\nGenerated %d resources!\n", count)
}

func getRegionSchema() (*schema.Schema, error) {
	c := "Region is the region you'd like your resource to be created in.\n"
	comment, err := comments.New(c, comments.WithTFTag("-"))
	if err != nil {
		return nil, errors.Wrap(err, "cannot build comment")
	}
	return &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: comment.String(),
	}, nil
}
