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

func Configure(p config.Provider) {
	p.AddResourceConfigurator("aws_instance", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.References["subnet_id"] = config.Reference{
			Type: "Subnet",
		}
		r.References["vpc_security_group_ids"] = config.Reference{
			Type: "SecurityGroup",
		}
		r.References["security_groups"] = config.Reference{
			Type: "SecurityGroup",
		}
		r.References["root_block_device.kms_key_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tf-aws/apis/kms/v1alpha1.Key",
		}
		r.References["network_interface.network_interface_id"] = config.Reference{
			Type: "EC2NetworkInterface",
		}
		r.References["ebs_block_device.kms_key_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tf-aws/apis/kms/v1alpha1.Key",
		}
		r.LateInitializer = config.LateInitializer{
			// NOTE(muvaf): These are ignored because they conflict with each other.
			// See the following for more details: https://github.com/crossplane-contrib/terrajet/issues/107
			IgnoredFields: []string{
				"subnet_id",
				"network_interface",
				"private_ip",
				"source_dest_check",
			},
		}
	})
	p.AddResourceConfigurator("aws_eip", func(r *config.Resource) {
		r.Kind = "ElasticIP"
		r.ExternalName = config.IdentifierFromProvider
		r.References["instance"] = config.Reference{
			Type: "Instance",
		}

		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_ec2_transit_gateway", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
	})

	p.AddResourceConfigurator("aws_ec2_transit_gateway_route", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.References["transit_gateway_attachment_id"] = config.Reference{
			Type: "TransitGatewayVpcAttachment",
		}
		r.References["transit_gateway_route_table_id"] = config.Reference{
			Type: "TransitGatewayRouteTable",
		}
	})

	p.AddResourceConfigurator("aws_ec2_transit_gateway_route_table", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.References["transit_gateway_id"] = config.Reference{
			Type: "TransitGateway",
		}
	})

	p.AddResourceConfigurator("aws_ec2_transit_gateway_route_table_association", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.References["transit_gateway_attachment_id"] = config.Reference{
			Type: "TransitGatewayVpcAttachment",
		}
		r.References["transit_gateway_route_table_id"] = config.Reference{
			Type: "TransitGatewayRouteTable",
		}
	})

	p.AddResourceConfigurator("aws_ec2_transit_gateway_vpc_attachment", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.References["subnet_ids"] = config.Reference{
			Type: "Subnet",
		}
		r.References["transit_gateway_id"] = config.Reference{
			Type: "TransitGateway",
		}
		r.References["vpc_id"] = config.Reference{
			Type: "VPC",
		}
	})

	p.AddResourceConfigurator("aws_ec2_transit_gateway_vpc_attachment_accepter", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.References["transit_gateway_attachment_id"] = config.Reference{
			Type: "TransitGatewayVpcAttachment",
		}
	})

	p.AddResourceConfigurator("aws_launch_template", func(r *config.Resource) {
		// Note(turkenh): Kind as "LaunchTemplate" fails with:
		// panic: cannot generate crd: cannot build types for LaunchTemplate:
		//  cannot build the types: cannot generate parameters type name of
		//  LaunchTemplate: could not generate a unique name for
		//  LaunchTemplateParameters
		r.Kind = "EC2LaunchTemplate"
		r.ExternalName = config.IdentifierFromProvider
		r.References["security_group_names"] = config.Reference{
			Type: "SecurityGroup",
		}
		r.References["vpc_security_group_ids"] = config.Reference{
			Type: "SecurityGroup",
		}
		r.References["block_device_mappings.ebs.kms_key_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tf-aws/apis/kms/v1alpha1.Key",
		}
		r.References["iam_instance_profile.arn"] = config.Reference{
			Type:      "github.com/crossplane-contrib/provider-tf-aws/apis/iam/v1alpha1.InstanceProfile",
			Extractor: common.PathARNExtractor,
		}
		r.References["iam_instance_profile.name"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tf-aws/apis/iam/v1alpha1.InstanceProfile",
		}
		r.References["network_interfaces.network_interface_id"] = config.Reference{
			Type: "EC2NetworkInterface",
		}
		r.References["network_interfaces.security_groups"] = config.Reference{
			Type: "SecurityGroup",
		}
		r.References["network_interfaces.subnet_id"] = config.Reference{
			Type: "Subnet",
		}
	})

	p.AddResourceConfigurator("aws_vpc", func(r *config.Resource) {
		r.Kind = "VPC"
		r.ExternalName = config.IdentifierFromProvider
	})

	p.AddResourceConfigurator("aws_vpc_endpoint", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.References["vpc_id"] = config.Reference{
			Type: "VPC",
		}
		r.References["subnet_ids"] = config.Reference{
			Type: "Subnet",
		}
		r.References["security_group_ids"] = config.Reference{
			Type: "SecurityGroup",
		}
		r.References["route_table_ids"] = config.Reference{
			Type: "RouteTable",
		}
	})

	p.AddResourceConfigurator("aws_subnet", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.References["vpc_id"] = config.Reference{
			Type: "VPC",
		}
		r.LateInitializer = config.LateInitializer{
			// NOTE(muvaf): Conflicts with AvailabilityZone. See the following
			// for more details: https://github.com/crossplane-contrib/terrajet/issues/107
			IgnoredFields: []string{
				"availability_zone_id",
			},
		}
	})

	p.AddResourceConfigurator("aws_network_interface", func(r *config.Resource) {
		r.Kind = "EC2NetworkInterface"
		r.ExternalName = config.IdentifierFromProvider
		r.References["vpc_id"] = config.Reference{
			Type: "VPC",
		}
		r.References["subnet_id"] = config.Reference{
			Type: "Subnet",
		}
		r.References["security_groups"] = config.Reference{
			Type: "SecurityGroup",
		}
		r.References["attachment.instance"] = config.Reference{
			Type: "Instance",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"interface_type",
			},
		}
	})

	p.AddResourceConfigurator("aws_security_group", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.References["vpc_id"] = config.Reference{
			Type: "VPC",
		}
		r.References["egress.security_groups"] = config.Reference{
			Type: "SecurityGroup",
		}
		r.References["ingress.security_groups"] = config.Reference{
			Type: "SecurityGroup",
		}
	})

	p.AddResourceConfigurator("aws_vpc_ipv4_cidr_block_association", func(r *config.Resource) {
		r.Kind = "IPv4CIDRBlockAssociation"
		r.ExternalName = config.IdentifierFromProvider
		r.References["vpc_id"] = config.Reference{
			Type: "VPC",
		}
	})

	p.AddResourceConfigurator("aws_vpc_peering_connection", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.References["vpc_id"] = config.Reference{
			Type: "VPC",
		}
		r.References["peer_vpc_id"] = config.Reference{
			Type: "VPC",
		}
	})

	p.AddResourceConfigurator("aws_route", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.References["route_table_id"] = config.Reference{
			Type: "RouteTable",
		}
		r.References["instance_id"] = config.Reference{
			Type: "Instance",
		}
		r.References["network_interface_id"] = config.Reference{
			Type: "EC2NetworkInterface",
		}
		r.References["transit_gateway_id"] = config.Reference{
			Type: "TransitGateway",
		}
		r.References["vpc_peering_connection_id"] = config.Reference{
			Type: "VpcPeeringConnection",
		}
		r.References["vpc_endpoint_id"] = config.Reference{
			Type: "VpcEndpoint",
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_route_table", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.References["vpc_id"] = config.Reference{
			Type: "VPC",
		}

		r.References["route.vpc_peering_connection_id"] = config.Reference{
			Type: "VpcPeeringConnection",
		}
		r.References["route.vpc_endpoint_id"] = config.Reference{
			Type: "VpcEndpoint",
		}
		r.References["route.network_interface_id"] = config.Reference{
			Type: "EC2NetworkInterface",
		}
		r.References["route.instance_id"] = config.Reference{
			Type: "Instance",
		}
	})

	p.AddResourceConfigurator("aws_route_table_association", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.References["subnet_id"] = config.Reference{
			Type: "Subnet",
		}
		r.References["route_table_id"] = config.Reference{
			Type: "RouteTable",
		}
	})

	p.AddResourceConfigurator("aws_ec2_transit_gateway_route_table_propagation", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.References["transit_gateway_attachment_id"] = config.Reference{
			Type: "TransitGatewayVpcAttachment",
		}
		r.References["transit_gateway_route_table_id"] = config.Reference{
			Type: "TransitGatewayRouteTable",
		}
	})
}
