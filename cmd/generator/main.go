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
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/iancoleman/strcase"

	"github.com/crossplane-contrib/terrajet/pkg/pipeline"
	"github.com/pkg/errors"
	"github.com/terraform-providers/terraform-provider-aws/aws"
)

const (
	ModulePath  = "github.com/crossplane-contrib/provider-tf-aws"
	GroupSuffix = ".aws.tf.crossplane.io"
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
			continue
		}
		// Both aws_config_configuration_recorder_status and aws_config_configuration_recorder
		// exist as resource which makes kubebuilder confuse them as same resource
		// and expect a storage version.
		if name == "aws_config_configuration_recorder_status" {
			continue
		}
		words := strings.Split(name, "_")
		groupName := words[1]
		if len(groups[groupName]) == 0 {
			groups[groupName] = map[string]*schema.Resource{}
		}
		groups[groupName][name] = resource
	}

	for group, resources := range groups {
		version := "v1alpha1"
		groupGen := pipeline.NewVersionGenerator(wd, strings.ToLower(group)+GroupSuffix, version)
		if err := groupGen.Generate(); err != nil {
			panic(errors.Wrap(err, "cannot generate version files"))
		}
		groupDir := filepath.Join(wd, "apis", group)
		crdGen := pipeline.NewCRDGenerator(groupDir, ModulePath, strings.ToLower(group)+GroupSuffix)
		for name, resource := range resources {
			// We don't want Aws prefix in all kinds.
			kind := strings.TrimPrefix(strcase.ToCamel(name), "Aws")
			if err := crdGen.Generate(version, kind, resource); err != nil {
				panic(errors.Wrap(err, "cannot generate crd"))
			}
		}
	}
}
