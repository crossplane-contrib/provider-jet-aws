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

package networkfirewall

import (
	"github.com/crossplane/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-jet-aws/config/common"
)

// Configure adds configurations for networkfirewall group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_networkfirewall_firewall", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider

		r.References = config.References{
			"firewall_policy_arn": {
				Type: "FirewallPolicy",
			},
			"subnet_mapping.subnet_id": {
				Type:              "github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2.Subnet",
				RefFieldName:      "SubnetIdRef",
				SelectorFieldName: "SubnetIdSelector",
			},
			"vpc_id": {
				Type:              "github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2.VPC",
				RefFieldName:      "VpcIdRef",
				SelectorFieldName: "VpcIdSelector",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_networkfirewall_firewall_policy", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider

		r.References = config.References{
			"firewall_policy.stateless_rule_group_reference.resource_arn": {
				Type: "RuleGroup",
			},
			"firewall_policy.stateful_rule_group_reference.resource_arn": {
				Type: "RuleGroup",
			},
		}
	})

	p.AddResourceConfigurator("aws_networkfirewall_logging_configuration", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider

		r.References = config.References{
			"firewall_arn": {
				Type: "Firewall",
			},
		}

	})

	p.AddResourceConfigurator("aws_networkfirewall_rule_group", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider

		r.UseAsync = true
	})
}
