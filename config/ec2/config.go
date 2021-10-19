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

package ec2

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-tf-aws/config/common"
)

func init() {
	config.Store.SetForResource("aws_instance", config.Resource{
		ExternalName: config.ExternalName{
			// Set to true explicitly since the value is calculated by AWS.
			DisableNameInitializer: true,
		},
		References: map[string]config.Reference{
			"subnet_id": {
				Type: "Subnet",
			},
			"vpc_security_group_ids": {
				Type: "SecurityGroup",
			},
			"security_groups": {
				Type: "SecurityGroup",
			},
			"root_block_device[*].kms_key_id": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/kms/v1alpha1.Key",
			},
			"network_interface[*].network_interface_id": {
				Type: "EC2NetworkInterface",
			},
			"ebs_block_device[*].kms_key_id": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/kms/v1alpha1.Key",
			},
		},
		LateInitializer: config.LateInitializer{
			// NOTE(muvaf): These are ignored because they conflict with each other.
			// See the following for more details: https://github.com/crossplane-contrib/terrajet/issues/107
			IgnoredFields: []string{
				"SubnetID",
				"NetworkInterface",
				"PrivateIP",
				"SourceDestCheck",
			},
		},
	})
	config.Store.SetForResource("aws_eip", config.Resource{
		Kind: "ElasticIP",
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: map[string]config.Reference{
			"instance": {
				Type: "Instance",
			},
		},

		UseAsync: true,
	})

	config.Store.SetForResource("aws_ec2_transit_gateway", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
	})

	config.Store.SetForResource("aws_ec2_transit_gateway_route", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: map[string]config.Reference{
			"transit_gateway_attachment_id": {
				Type: "TransitGatewayVpcAttachment",
			},
			"transit_gateway_route_table_id": {
				Type: "TransitGatewayRouteTable",
			},
		},
	})

	config.Store.SetForResource("aws_ec2_transit_gateway_route_table", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: map[string]config.Reference{
			"transit_gateway_id": {
				Type: "TransitGateway",
			},
		},
	})
	config.Store.SetForResource("aws_ec2_transit_gateway_route_table_association", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: map[string]config.Reference{
			"transit_gateway_attachment_id": {
				Type: "TransitGatewayVpcAttachment",
			},
			"transit_gateway_route_table_id": {
				Type: "TransitGatewayRouteTable",
			},
		},
	})
	config.Store.SetForResource("aws_ec2_transit_gateway_vpc_attachment", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: map[string]config.Reference{
			"subnet_ids": {
				Type: "Subnet",
			},
			"transit_gateway_id": {
				Type: "TransitGateway",
			},
			"vpc_id": {
				Type: "VPC",
			},
		},
	})
	config.Store.SetForResource("aws_ec2_transit_gateway_vpc_attachment_accepter", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: map[string]config.Reference{
			"transit_gateway_attachment_id": {
				Type: "TransitGatewayVpcAttachment",
			},
		},
	})

	config.Store.SetForResource("aws_launch_template", config.Resource{
		// Note(turkenh): Kind as "LaunchTemplate" fails with:
		// panic: cannot generate crd: cannot build types for LaunchTemplate:
		//  cannot build the types: cannot generate parameters type name of
		//  LaunchTemplate: could not generate a unique name for
		//  LaunchTemplateParameters
		Kind: "EC2LaunchTemplate",

		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},

		References: map[string]config.Reference{
			"security_group_names": {
				Type: "SecurityGroup",
			},
			"vpc_security_group_ids": {
				Type: "SecurityGroup",
			},
			"block_device_mappings[*].ebs[*].kms_key_id": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/kms/v1alpha1.Key",
			},
			"iam_instance_profile[*].arn": {
				Type:      "github.com/crossplane-contrib/provider-tf-aws/apis/iam/v1alpha1.InstanceProfile",
				Extractor: common.PathARNExtractor,
			},
			"iam_instance_profile[*].name": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/iam/v1alpha1.InstanceProfile",
			},
			"network_interfaces[*].network_interface_id": {
				Type: "EC2NetworkInterface",
			},
			"network_interfaces[*].security_groups": {
				Type: "SecurityGroup",
			},
			"network_interfaces[*].subnet_id": {
				Type: "Subnet",
			},
		},
	})
	config.Store.SetForResource("aws_vpc", config.Resource{
		Kind: "VPC",
		ExternalName: config.ExternalName{
			// Set to true explicitly since the value is calculated by AWS.
			DisableNameInitializer: true,
		},
	})
	config.Store.SetForResource("aws_vpc_endpoint", config.Resource{
		ExternalName: config.ExternalName{
			// Set to true explicitly since the value is calculated by AWS.
			DisableNameInitializer: true,
		},
		References: map[string]config.Reference{
			"vpc_id": {
				Type: "VPC",
			},
			"subnet_ids": {
				Type: "Subnet",
			},
			"security_group_ids": {
				Type: "SecurityGroup",
			},
			"route_table_ids": {
				Type: "RouteTable",
			},
		},
	})
	config.Store.SetForResource("aws_subnet", config.Resource{
		ExternalName: config.ExternalName{
			// Set to true explicitly since the value is calculated by AWS.
			DisableNameInitializer: true,
		},
		References: map[string]config.Reference{
			"vpc_id": {
				Type: "VPC",
			},
		},
		LateInitializer: config.LateInitializer{
			// NOTE(muvaf): Conflicts with AvailabilityZone. See the following
			// for more details: https://github.com/crossplane-contrib/terrajet/issues/107
			IgnoredFields: []string{
				"AvailabilityZoneID",
			},
		},
	})
	config.Store.SetForResource("aws_network_interface", config.Resource{
		Kind: "EC2NetworkInterface",
		ExternalName: config.ExternalName{
			// Set to true explicitly since the value is calculated by AWS.
			DisableNameInitializer: true,
		},
		References: map[string]config.Reference{
			"vpc_id": {
				Type: "VPC",
			},
			"subnet_id": {
				Type: "Subnet",
			},
			"security_groups": {
				Type: "SecurityGroup",
			},
			"attachment[*].instance": {
				Type: "Instance",
			},
		},
		LateInitializer: config.LateInitializer{
			IgnoredFields: []string{
				"InterfaceType",
			},
		},
	})
	config.Store.SetForResource("aws_security_group", config.Resource{
		ExternalName: config.ExternalName{
			// Set to true explicitly since the value is calculated by AWS.
			DisableNameInitializer: true,
		},
		References: map[string]config.Reference{
			"vpc_id": {
				Type: "VPC",
			},
			"egress[*].security_groups": {
				Type: "SecurityGroup",
			},
			"ingress[*].security_groups": {
				Type: "SecurityGroup",
			},
		},
	})
	config.Store.SetForResource("aws_vpc_ipv4_cidr_block_association", config.Resource{
		Kind: "IPv4CIDRBlockAssociation",
		ExternalName: config.ExternalName{
			// Set to true explicitly since the value is calculated by AWS.
			DisableNameInitializer: true,
		},
		References: map[string]config.Reference{
			"vpc_id": {
				Type: "VPC",
			},
		},
	})
	config.Store.SetForResource("aws_vpc_peering_connection", config.Resource{
		ExternalName: config.ExternalName{
			// Set to true explicitly since the value is calculated by AWS.
			DisableNameInitializer: true,
		},
		References: map[string]config.Reference{
			"vpc_id": {
				Type: "VPC",
			},
			"peer_vpc_id": {
				Type: "VPC",
			},
		},
	})
	config.Store.SetForResource("aws_route_table", config.Resource{
		ExternalName: config.ExternalName{
			// Set to true explicitly since the value is calculated by AWS.
			DisableNameInitializer: true,
		},
		References: map[string]config.Reference{
			"vpc_peering_connection_id": {
				Type: "PeeringConnection",
			},
			"vpc_endpoint_id": {
				Type: "Endpoint",
			},
			"network_interface_id": {
				Type: "EC2NetworkInterface",
			},
			"instance_id": {
				Type: "Instance",
			},
		},
	})
	config.Store.SetForResource("aws_route_table_association", config.Resource{
		ExternalName: config.ExternalName{
			// Set to true explicitly since the value is calculated by AWS.
			DisableNameInitializer: true,
		},
		References: map[string]config.Reference{
			"subnet_id": {
				Type: "Subnet",
			},
			"route_table_id": {
				Type: "RouteTable",
			},
		},
	})
	config.Store.SetForResource("aws_ec2_transit_gateway_route_table_propagation", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: map[string]config.Reference{
			"transit_gateway_attachment_id": {
				Type: "TransitGatewayVpcAttachment",
			},
			"transit_gateway_route_table_id": {
				Type: "TransitGatewayRouteTable",
			},
		},
	})
}
