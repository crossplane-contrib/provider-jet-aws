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

import "github.com/crossplane-contrib/terrajet/pkg/config"

const (
	SelfPackagePath = "github.com/crossplane-contrib/provider-tf-aws/config/ec2"
)

func init() {
	config.Store.SetForResource("aws_instance", config.Resource{
		ExternalName: config.ExternalName{
			// Set to true explicitly since the value is calculated by AWS.
			DisableNameInitializer: true,
		},
		References: map[string]config.Reference{
			"subnet_id": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/vpc/v1alpha1.Subnet",
			},
			"vpc_security_group_ids": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/vpc/v1alpha1.SecurityGroup",
			},
			"security_groups": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/vpc/v1alpha1.SecurityGroup",
			},
			"root_block_device[*].kms_key_id": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/kms/v1alpha1.Key",
			},
			"network_interface[*].network_interface_id": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/kms/v1alpha1.Key",
			},
			"ebs_block_device[*].kms_key_id": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/kms/v1alpha1.Key",
			},
		},
	})
}
