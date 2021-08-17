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
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/iancoleman/strcase"

	"github.com/crossplane-contrib/terrajet/pkg/pipeline"
	"github.com/pkg/errors"
	"github.com/terraform-providers/terraform-provider-aws/aws"
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
		words := strings.Split(name, "_")
		groupName := words[1]
		if len(groups[groupName]) == 0 {
			groups[groupName] = map[string]*schema.Resource{}
		}
		groups[groupName][name] = resource
	}

	//groups = map[string]map[string]*schema.Resource{
	//	"wafv2": {
	//		"aws_wafv2_web_acl": aws.Provider().ResourcesMap["aws_wafv2_web_acl"],
	//	},
	//}

	for group, resources := range groups {
		groupGen := pipeline.NewVersionGenerator(wd, group+".aws.tf.crossplane.io", "v1alpha1")
		if err := groupGen.Generate(); err != nil {
			panic(errors.Wrap(err, "cannot generate version files"))
		}
		for name, resource := range resources {
			kind := strcase.ToCamel(name)
			crdGen := pipeline.NewCRDGenerator(wd, group+".aws.tf.crossplane.io", "v1alpha1", kind, resource)
			fmt.Printf("doing %s\n", name)
			if err := crdGen.Generate(); err != nil {
				panic(errors.Wrap(err, "cannot generate crd"))
			}
		}
	}
}
