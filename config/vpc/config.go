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

package vpc

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"
)

func init() {
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
			"subnet_ids[*]": {
				Type: "Subnet",
			},
			"security_group_ids[*]": {
				Type: "SecurityGroup",
			},
			"route_table_ids[*]": {
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
			"egress[*].security_groups[*]": {
				Type: "SecurityGroup",
			},
			"ingress[*].security_groups[*]": {
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
				Type: "NetworkInterface",
			},
			"instance_id": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1.Instance",
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
}
