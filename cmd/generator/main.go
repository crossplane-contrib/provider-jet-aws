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
	"path/filepath"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
	"github.com/terraform-providers/terraform-provider-aws/aws"

	"github.com/crossplane-contrib/terrajet/pkg/pipeline"
)

// Constants to use in generated artifacts.
const (
	modulePath  = "github.com/crossplane-contrib/provider-tf-aws"
	groupSuffix = ".aws.tf.crossplane.io"
)

var (
	// We expect a function called "ProviderConfigBuilder" of type
	// "ProviderConfigFn" (https://github.com/crossplane-contrib/terrajet/blob/4246657031f181fdaf0de83e83ceaa8735307180/pkg/terraform/controller.go#L33)
	// available at this path
	providerConfigBuilderPath = filepath.Join(modulePath, "internal", "clients")
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(errors.Wrap(err, "cannot get working directory"))
	}
	groups := map[string]map[string]*schema.Resource{}
	for name, resource := range aws.Provider().ResourcesMap {
		if len(resource.Schema) == 0 {
			// There are resources with no schema, like aws_securityhub_account
			// that we will address later.
			fmt.Printf("Skipping resource %s because it has no schema\n", name)
			continue
		}
		// Both aws_config_configuration_recorder_status and aws_config_configuration_recorder
		// exist as resource which makes kubebuilder confuse them as same resource
		// and expect a storage version.
		if name == "aws_config_configuration_recorder_status" {
			fmt.Printf("Skipping resource %s because there is a resource whose name is prefix of this resource\n", name)
			continue
		}
		words := strings.Split(name, "_")
		groupName := words[1]
		if len(groups[groupName]) == 0 {
			groups[groupName] = map[string]*schema.Resource{}
		}
		groups[groupName][name] = resource
	}
	count := 0
	for group, resources := range groups {
		version := "v1alpha1"
		groupGen := pipeline.NewVersionGenerator(wd, strings.ToLower(group)+groupSuffix, version)
		if err := groupGen.Generate(); err != nil {
			panic(errors.Wrap(err, "cannot generate version files"))
		}
		apiGroupDir := filepath.Join(wd, "apis", group)
		ctrlGroupDir := filepath.Join(wd, "internal", "controller", group)

		crdGen := pipeline.NewCRDGenerator(apiGroupDir, modulePath, strings.ToLower(group)+groupSuffix)
		tfGen := pipeline.NewTerraformedGenerator(apiGroupDir, modulePath, strings.ToLower(group)+groupSuffix)
		ctrlGen := pipeline.NewControllerGenerator(ctrlGroupDir, modulePath, strings.ToLower(group)+groupSuffix, providerConfigBuilderPath)

		for name, resource := range resources {
			// We don't want Aws prefix in all kinds.
			kind := strings.TrimPrefix(strcase.ToCamel(name), "Aws")
			if err := crdGen.Generate(version, kind, resource); err != nil {
				panic(errors.Wrap(err, "cannot generate crd"))
			}
			if err := tfGen.Generate(version, kind, name, "id"); err != nil {
				panic(errors.Wrap(err, "cannot generate terraformed"))
			}
			if err := ctrlGen.Generate(version, kind); err != nil {
				panic(errors.Wrap(err, "cannot generate controller"))
			}
			count++
		}
	}
	fmt.Printf("\nGenerated %d resources!\n", count)
}
