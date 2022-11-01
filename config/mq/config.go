/*
Copyright 2022 The Crossplane Authors.

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

package mq

import (
	"github.com/crossplane/terrajet/pkg/config"

	"github.com/dkb-bank/provider-jet-aws/config/common"
)

// Configure adds configurations for rds group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_mq_broker", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		// Due to a terrajet limitation, we cannot use "metedata.name" field as the name of the resource
		// Therefore, "spec.forProvider.brokerName" field is not omitted
		// Details can be found in https://github.com/crossplane/terrajet/issues/280
		r.ExternalName = config.IdentifierFromProvider
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_mq_configuration", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		// Due to a terrajet limitation, we cannot use "metedata.name" field as the name of the resource
		// Therefore, "spec.forProvider.name" field is not omitted
		// Details can be found in https://github.com/crossplane/terrajet/issues/280
		r.ExternalName = config.IdentifierFromProvider
	})

}
