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
	"regexp"

	"github.com/crossplane/terrajet/pkg/types/comments"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	tjconfig "github.com/crossplane/terrajet/pkg/config"
)

// GroupMap contains all overrides we'd like to make to the default group search.
// It's written with data from TF Provider AWS repo service grouping in here:
// https://github.com/hashicorp/terraform-provider-aws/tree/main/internal/service
//
// At the end, all of them are based on grouping of the AWS Go SDK.
var GroupMap = map[string]string{
	"^aws_api_gateway.*":                        "apigateway",
	"^aws_launch_configuration":                 "autoscaling",
	"^aws_cloudhsm_v2_.*":                       "cloudhsmv2",
	"^aws_cloudtrail":                           "cloudtrail",
	"^aws_cur_.*":                               "costandusagereportservice",
	"^aws_dms_.*":                               "databasemigrationservice",
	"^aws_dx_.*":                                "directconnect",
	"^aws_directory_service.*":                  "directoryservice",
	"^aws_ami.*":                                "ec2",
	"^aws_customer_gateway$":                    "ec2",
	"^aws_default_network_acl$":                 "ec2",
	"^aws_default_route_table$":                 "ec2",
	"^aws_default_security_group$":              "ec2",
	"^aws_default_subnet$":                      "ec2",
	"^aws_default_vpc$":                         "ec2",
	"^aws_default_vpc_dhcp_options$":            "ec2",
	"^aws_egress_only_internet_gateway$":        "ec2",
	"^aws_flow_log$":                            "ec2",
	"^aws_internet_gateway$":                    "ec2",
	"^aws_main_route_table_association$":        "ec2",
	"^aws_nat_gateway$":                         "ec2",
	"^aws_network_acl$":                         "ec2",
	"^aws_network_acl_rule$":                    "ec2",
	"^aws_network_interface$":                   "ec2",
	"^aws_network_interface_attachment$":        "ec2",
	"^aws_network_interface_sg_attachment$":     "ec2",
	"^aws_vpn.*":                                "ec2",
	"^aws_ebs_default_kms_key$":                 "ec2",
	"^aws_ebs_encryption_by_default$":           "ec2",
	"^aws_eip.*":                                "ec2",
	"^aws_instance":                             "ec2",
	"^aws_security_group.*":                     "ec2",
	"^aws_subnet.*":                             "ec2",
	"^aws_network_.*":                           "ec2",
	"^aws_key_pair.*":                           "ec2",
	"^aws_launch_template.*":                    "ec2",
	"^aws_main_route_table_.*":                  "ec2",
	"^aws_placement_group.*":                    "ec2",
	"^aws_route_.*":                             "ec2",
	"^aws_route$":                               "ec2",
	"^aws_snapshot_create_volume_permission$":   "ec2",
	"^aws_spot.*":                               "ec2",
	"^aws_volume_attachment$":                   "ec2",
	"^aws_elastic_beanstalk.*":                  "elasticbeanstalk",
	"^aws_app_cookie_stickiness_policy$":        "elb",
	"^aws_elb$":                                 "elb",
	"^aws_elb_attachment$":                      "elb",
	"^aws_lb_cookie_stickiness_policy$":         "elb",
	"^aws_lb_ssl_negotiation_policy$":           "elb",
	"^aws_load_balancer_backend_server_policy$": "elb",
	"^aws_load_balancer_listener_policy$":       "elb",
	"^aws_load_balancer_policy$":                "elb",
	"^aws_proxy_protocol_policy$":               "elb",
	"^aws_lb$":                                  "elbv2",
	"^aws_lb_listener$":                         "elbv2",
	"^aws_lb_listener_certificate$":             "elbv2",
	"^aws_lb_listener_rule$":                    "elbv2",
	"^aws_lb_target_group$":                     "elbv2",
	"^aws_lb_target_group_attachment$":          "elbv2",
	"^aws_schemas_discoverer$":                  "eventbridge",
	"^aws_schemas_registry$":                    "eventbridge",
	"^aws_schemas_schema$":                      "eventbridge",
	"^aws_media_convert.*":                      "mediaconvert",
	"^aws_media_package.*":                      "mediapackage",
	"^aws_media_store.*":                        "mediastore",
	"aws_db_*":                                  "rds",
	"^aws_route53_resolver_.*":                  "route53resolver",
	"^aws_service_discovery_.*.*":               "servicediscovery",

	"^aws_vpc.*": "ec2",
	"^aws_db_.*": "rds",
	"^aws_db$":   "rds",
}

// KindMap contains kind string overrides.
var KindMap = map[string]string{
	"aws_ami":                                         "AMI",
	"aws_ami_copy":                                    "AMICopy",
	"aws_ami_from_instance":                           "AMIFromInstance",
	"aws_ami_launch_permission":                       "AMILaunchPermission",
	"aws_autoscaling_group":                           "AutoscalingGroup",
	"aws_cloudformation_type":                         "CloudFormationType",
	"aws_config_configuration_recorder_status":        "AWSConfigurationRecorderStatus",
	"aws_cloudtrail":                                  "Trail",
	"aws_db_instance":                                 "Instance",
	"aws_db_parameter_group":                          "ParameterGroup",
	"aws_ebs_default_kms_key":                         "EBSDefaultKMSKey",
	"aws_ebs_encryption_by_default":                   "EBSEncryptionByDefault",
	"aws_ec2_transit_gateway_vpc_attachment":          "TransitGatewayVPCAttachment",
	"aws_ec2_transit_gateway_vpc_attachment_accepter": "TransitGatewayVPCAttachmentAccepter",
	"aws_eip":                                      "ElasticIP",
	"aws_eip_association":                          "ElasticIPAssociation",
	"aws_elastic_beanstalk_application":            "Application",
	"aws_elastic_beanstalk_application_version":    "ApplicationVersion",
	"aws_elastic_beanstalk_configuration_template": "ConfigurationTemplate",
	"aws_elastic_beanstalk_environment":            "Environment",
	"aws_elb":                                      "ElasticLoadBalancer",
	"aws_elb_attachment":                           "ElasticLoadBalancerAttachment",
	"aws_key_pair":                                 "KeyPair",
	"aws_launch_configuration":                     "LaunchConfiguration",
	"aws_launch_template":                          "LaunchTemplate",
	"aws_lb":                                       "LoadBalancer",
	"aws_lb_listener":                              "LoadBalancerListener",
	"aws_load_balancer_backend_server_policy":      "ElasticLoadBalancerBackendServerPolicy",
	"aws_load_balancer_listener_policy":            "ElasticLoadBalancerListenerPolicy",
	"aws_load_balancer_policy":                     "ElasticLoadBalancerPolicy",
	"aws_network_interface":                        "NetworkInterface",
	"aws_placement_group":                          "PlacementGroup",
	"aws_proxy_protocol_policy$":                   "ProxyProtocolPolicy",
	"aws_route_table":                              "RouteTable",
	"aws_route_table_association":                  "RouteTableAssociation",
	"aws_main_route_table_association":             "MainRouteTableAssociation",
	"aws_rds_cluster":                              "DBCluster",
	"aws_schemas_discoverer":                       "SchemasDiscoverer",
	"aws_security_group":                           "SecurityGroup",
	"aws_security_group_rule":                      "SecurityGroupRule",
	"aws_snapshot_create_volume_permission":        "SnapshotCreateVolumePermission",
	"aws_volume_attachment":                        "VolumeAttachment",
	"aws_vpc":                                      "VPC",
	"aws_vpc_endpoint":                             "VPCEndpoint",
	"aws_vpc_ipv4_cidr_block_association":          "IPv4CIDRBlockAssociation",
	"aws_vpc_peering_connection":                   "VPCPeeringConnection",
}

// GroupOverrides overrides the group of the resource if it matches expressions
// in given group map.
func GroupOverrides() tjconfig.ResourceOption {
	return func(r *tjconfig.Resource) {
		if group, ok := matches(r.Name, GroupMap); ok {
			r.ShortGroup = group
		}
	}
}

// KindOverrides overrides the kind of the resources given in KindMap.
func KindOverrides() tjconfig.ResourceOption {
	return func(r *tjconfig.Resource) {
		if k, ok := KindMap[r.Name]; ok {
			r.Kind = k
		}
	}
}

// RegionAddition adds region to the spec of all resources except iam group which
// does not have a region notion.
func RegionAddition() tjconfig.ResourceOption {
	return func(r *tjconfig.Resource) {
		if r.ShortGroup == "iam" {
			return
		}
		c := "Region is the region you'd like your resource to be created in.\n"
		comment, err := comments.New(c, comments.WithTFTag("-"))
		if err != nil {
			panic(errors.Wrap(err, "cannot build comment for region"))
		}
		r.TerraformResource.Schema["region"] = &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: comment.String(),
		}
	}
}

// TagsAllRemoval removes the tags_all field that is used only in tfstate to
// accumulate provider-wide default tags in TF, which is not something we support.
// So, we don't need it as a parameter while "tags" is already in place.
func TagsAllRemoval() tjconfig.ResourceOption {
	return func(r *tjconfig.Resource) {
		if t, ok := r.TerraformResource.Schema["tags_all"]; ok {
			t.Computed = true
			t.Optional = false
		}
	}
}

// IdentifierAssignedByAWS will work for all AWS types because even if the ID
// is assigned by user, we'll see it in the TF State ID.
// The resource-specific configurations should override this whenever possible.
func IdentifierAssignedByAWS() tjconfig.ResourceOption {
	return func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
	}
}

// NamePrefixRemoval makes sure we remove name_prefix from all since it is mostly
// for Terraform functionality.
func NamePrefixRemoval() tjconfig.ResourceOption {
	return func(r *tjconfig.Resource) {
		for _, f := range r.ExternalName.OmittedFields {
			if f == "name_prefix" {
				return
			}
		}
		r.ExternalName.OmittedFields = append(r.ExternalName.OmittedFields, "name_prefix")
	}
}

func matches(name string, regexes map[string]string) (string, bool) {
	for k, v := range regexes {
		ok, err := regexp.MatchString(k, name)
		if err != nil {
			panic(errors.Wrap(err, "cannot run regular expression match"))
		}
		if ok {
			return v, true
		}
	}
	return "", false
}
