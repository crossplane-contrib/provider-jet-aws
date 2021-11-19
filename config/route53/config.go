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

package route53

import (
	"strings"

	"github.com/crossplane-contrib/terrajet/pkg/config"
)

// Configure route53 resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_route53_delegation_set", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
	})
	p.AddResourceConfigurator("aws_route53_health_check", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
	})
	p.AddResourceConfigurator("aws_route53_hosted_zone_dnssec", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.References = config.References{
			"hosted_zone_id": config.Reference{
				Type: "Zone",
			},
		}
	})
	p.AddResourceConfigurator("aws_route53_key_signing_key", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.References = config.References{
			"hosted_zone_id": config.Reference{
				Type: "Zone",
			},
			"key_management_service_arn": config.Reference{
				Type:      "github.com/crossplane-contrib/provider-jet-aws/apis/kms/v1alpha1.Key",
				Extractor: "github.com/crossplane-contrib/provider-jet-aws/apis/kms/v1alpha1.KMSKeyARN()",
			},
		}
	})
	p.AddResourceConfigurator("aws_route53_query_log", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.References = config.References{
			"hosted_zone_id": config.Reference{
				Type: "Zone",
			},
		}
	})
	p.AddResourceConfigurator("aws_route53_record", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.References = config.References{
			"zone_id": config.Reference{
				Type: "Zone",
			},
			"health_check_id": config.Reference{
				Type: "HealthCheck",
			},
		}
	})
	p.AddResourceConfigurator("aws_route53_vpc_association_authorization", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ExternalName.SetIdentifierArgumentFn = func(base map[string]interface{}, externalName string) {
			words := strings.Split(externalName, ":")
			if len(words) != 2 {
				return
			}
			base["zone_id"] = words[0]
			base["vpc_id"] = words[1]
		}
		r.References = config.References{
			"zone_id": config.Reference{
				Type: "Zone",
			},
			"vpc_id": config.Reference{
				Type: "github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha1.VPC",
			},
		}
	})
	p.AddResourceConfigurator("aws_route53_zone", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.References = config.References{
			"delegation_set_id": config.Reference{
				Type: "DelegationSet",
			},
			"vpc.vpc_id": config.Reference{
				Type: "github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha1.VPC",
			},
		}
	})
	p.AddResourceConfigurator("aws_route53_zone_association", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// Z123456ABCDEFG:vpc-12345678
		// Z123456ABCDEFG:vpc-12345678:us-east-2
		r.ExternalName.SetIdentifierArgumentFn = func(base map[string]interface{}, externalName string) {
			words := strings.Split(externalName, ":")
			if len(words) >= 2 {
				base["zone_id"] = words[0]
				base["vpc_id"] = words[1]
			}
			if len(words) == 3 {
				base["vpc_region"] = words[2]
			}
		}
		r.References = config.References{
			"zone_id": config.Reference{
				Type: "Zone",
			},
			"vpc_id": config.Reference{
				Type: "github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha1.VPC",
			},
		}
	})
}
