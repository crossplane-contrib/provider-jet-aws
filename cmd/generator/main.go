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
	"os/exec"
	"sort"
	"strings"

	"github.com/crossplane-contrib/terrajet/pkg/comments"
	"github.com/crossplane-contrib/terrajet/pkg/config"
	"github.com/crossplane-contrib/terrajet/pkg/pipeline"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
	"github.com/terraform-providers/terraform-provider-aws/aws"

	// This is to get init functions registering custom configs to be called.
	// TODO(hasan): Revisit getting custom config with pipeline abstraction work
	_ "github.com/crossplane-contrib/provider-tf-aws/apis"
)

// Constants to use in generated artifacts.
const (
	modulePath        = "github.com/crossplane-contrib/provider-tf-aws"
	groupSuffix       = ".aws.tf.crossplane.io"
	providerShortName = "tfaws"
)

// These resources cannot be generated because of their suffixes colliding with
// kubebuilder-accepted type suffixes.
var skipList = map[string]struct{}{
	"aws_config_configuration_recorder_status": {},
	"aws_vpc_peering_connection_accepter":      {},
	"aws_elasticache_user_group":               {},
	"aws_waf_rule_group":                       {},
	"aws_wafregional_rule_group":               {},
	"aws_shield_protection_group":              {},
	"aws_route53_resolver_firewall_rule_group": {},
	"aws_kinesis_analytics_application":        {},
}

var techPreviewIncludedGroups = map[string]struct{}{
	"vpc":      {},
	"rds":      {},
	"eks":      {},
	"ec2":      {},
	"s3":       {},
	"iam":      {},
	"default":  {},
	"eip":      {},
	"elb":      {},
	"instance": {},
	"lb":       {},
	"main":     {},
	"route":    {},
	"route53":  {},
	"subnet":   {},
}

func main() { // nolint:gocyclo
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
		if _, ok := skipList[name]; ok {
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
	versionPkgList := []string{
		"github.com/crossplane-contrib/provider-tf-aws/apis/v1alpha1",
	}
	controllerPkgList := []string{
		"github.com/crossplane-contrib/provider-tf-aws/internal/controller/tfaws",
	}

	regionSchema, err := getRegionSchema()
	if err != nil {
		panic(errors.Wrap(err, "cannot get region schema"))
	}

	for group, resources := range groups {
		if _, ok := techPreviewIncludedGroups[group]; !ok {
			continue
		}
		version := "v1alpha1"
		versionGen := pipeline.NewVersionGenerator(wd, modulePath, strings.ToLower(group)+groupSuffix, version)

		crdGen := pipeline.NewCRDGenerator(versionGen.Package(), versionGen.DirectoryPath(), strings.ToLower(group)+groupSuffix, providerShortName)
		tfGen := pipeline.NewTerraformedGenerator(versionGen.Package(), versionGen.DirectoryPath())
		ctrlGen := pipeline.NewControllerGenerator(wd, modulePath, strings.ToLower(group)+groupSuffix)

		keys := make([]string, len(resources))
		i := 0
		for k := range resources {
			keys[i] = k
			i++
		}
		sort.Strings(keys)

		for _, name := range keys {
			// We don't want Aws prefix in all kinds.
			kind := strings.TrimPrefix(strcase.ToCamel(name), "Aws")
			resource := resources[name]
			resource.Schema["region"] = regionSchema
			rc := config.NewResource(version, kind, name)
			if err = rc.OverrideConfig(config.Store.GetForResource(name)); err != nil {
				panic(errors.Wrap(err, "cannot override config"))
			}

			if err = crdGen.Generate(rc, resource); err != nil {
				panic(errors.Wrap(err, "cannot generate crd"))
			}
			if err = tfGen.Generate(rc); err != nil {
				panic(errors.Wrap(err, "cannot generate terraformed"))
			}
			ctrlPkgPath, err := ctrlGen.Generate(rc, versionGen.Package().Path())
			if err != nil {
				panic(errors.Wrap(err, "cannot generate controller"))
			}
			controllerPkgList = append(controllerPkgList, ctrlPkgPath)
			count++
		}

		if err := versionGen.Generate(); err != nil {
			panic(errors.Wrap(err, "cannot generate version files"))
		}
		versionPkgList = append(versionPkgList, versionGen.Package().Path())
	}

	if err := pipeline.NewRegisterGenerator(wd, modulePath).Generate(versionPkgList); err != nil {
		panic(errors.Wrap(err, "cannot generate register file"))
	}
	if err := pipeline.NewSetupGenerator(wd, modulePath).Generate(controllerPkgList); err != nil {
		panic(errors.Wrap(err, "cannot generate setup file"))
	}
	if err := exec.Command("bash", "-c", "goimports -w $(find apis -iname 'zz_*')").Run(); err != nil {
		panic(errors.Wrap(err, "cannot run goimports for apis folder"))
	}
	if err := exec.Command("bash", "-c", "goimports -w $(find internal -iname 'zz_*')").Run(); err != nil {
		panic(errors.Wrap(err, "cannot run goimports for internal folder"))
	}
	fmt.Printf("\nGenerated %d resources!\n", count)
}

func getRegionSchema() (*schema.Schema, error) {
	c := "Region is the region you'd like your resource to be created in.\n"
	comment, err := comments.New(c, comments.WithTFTag("-"))
	if err != nil {
		return nil, errors.Wrap(err, "cannot build comment")
	}
	return &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: comment.String(),
	}, nil
}
