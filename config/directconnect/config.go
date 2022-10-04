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

package directconnect

import (
	"github.com/crossplane/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-jet-aws/config/common"
)

// Configure adds configurations for directconnect group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_dx_bgp_peer", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
		r.References["virtual_interface_id"] = config.Reference{
			Type: "TransitVirtualInterface", // TODO: or PublicVirtualInterface or PrivateVirtualInterface
		}
	})

	p.AddResourceConfigurator("aws_dx_connection", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
	})

	p.AddResourceConfigurator("aws_dx_gateway", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
	})

	p.AddResourceConfigurator("aws_dx_gateway_association", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
		r.References["dx_gateway_id"] = config.Reference{
			Type: "Gateway",
		}
		r.References["associated_gateway_id"] = config.Reference{
			Type: "TransitGateway", // how to "or VPN Gateway"?
		}
	})

	p.AddResourceConfigurator("aws_dx_private_virtual_interface", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
		r.References["connection_id"] = config.Reference{
			Type: "Connection",
		}
	})

	p.AddResourceConfigurator("aws_dx_public_virtual_interface", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.NameAsIdentifier
		r.References["connection_id"] = config.Reference{
			Type: "Connection",
		}
	})

	p.AddResourceConfigurator("aws_dx_transit_virtual_interface", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
		r.References["connection_id"] = config.Reference{
			Type: "Connection",
		}
		r.References["dx_gateway_id"] = config.Reference{
			Type: "Gateway",
		}
	})

	p.AddResourceConfigurator("aws_dx_lag", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
	})

	p.AddResourceConfigurator("aws_dx_connection_association", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
		r.References["connection_id"] = config.Reference{
			Type: "Connection",
		}
		r.References["lag_id"] = config.Reference{
			Type: "Lag",
		}
	})
}
