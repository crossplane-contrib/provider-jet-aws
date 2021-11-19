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

package config

import (
	"regexp"

	"github.com/crossplane-contrib/terrajet/pkg/types/comments"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	tjconfig "github.com/crossplane-contrib/terrajet/pkg/config"
)

// GroupMap contains all overrides we'd like to make to the default group search.
var GroupMap = map[string]string{
	"^aws_security_group.*":    "ec2",
	"^aws_subnet.*":            "ec2",
	"^aws_network_.*":          "ec2",
	"^aws_eip_.*":              "ec2",
	"^aws_eip$":                "ec2",
	"^aws_launch_template_.*":  "ec2",
	"^aws_launch_template$":    "ec2",
	"^aws_route_.*":            "ec2",
	"^aws_route$":              "ec2",
	"^aws_main_route_table_.*": "ec2",
	"^aws_instance_.*":         "ec2",
	"^aws_instance$":           "ec2",
	"^aws_vpc_.*":              "ec2",
	"^aws_vpc$":                "ec2",
	"^aws_lb_.*":               "elasticloadbalancing",
	"^aws_lb$":                 "elasticloadbalancing",
	"^aws_db_.*":               "rds",
	"^aws_db$":                 "rds",
}

// GroupOverrides overrides the group of the resource if it matches expressions
// in given group map.
func GroupOverrides() tjconfig.ResourceOption {
	return func(r *tjconfig.Resource) {
		if group, ok := matches(r.Name, GroupMap); ok {
			r.ShortGroup = group
		}
	}
}

// RegionAddition adds region to the spec of all resources except iam group which
// does not have a region notion.
func RegionAddition() tjconfig.ResourceOption {
	return func(r *tjconfig.Resource) {
		if r.ShortGroup == "iam" {
			return
		}
		c := "Region is the region you'd like your resource to be created in.\n"
		comment, err := comments.New(c, comments.WithTFTag("-"))
		if err != nil {
			panic(errors.Wrap(err, "cannot build comment for region"))
		}
		r.TerraformResource.Schema["region"] = &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: comment.String(),
		}
	}
}

// TagsAllRemoval removes the tags_all field that is used only in tfstate to
// accumulate provider-wide default tags in TF, which is not something we support.
// So, we don't need it as a parameter while "tags" is already in place.
func TagsAllRemoval() tjconfig.ResourceOption {
	return func(r *tjconfig.Resource) {
		if t, ok := r.TerraformResource.Schema["tags_all"]; ok {
			t.Computed = true
			t.Optional = false
		}
	}
}

// IdentifierAssignedByAWS will work for all AWS types because even if the ID
// is assigned by user, we'll see it in the TF State ID.
// The resource-specific configurations should override this whenever possible.
func IdentifierAssignedByAWS() tjconfig.ResourceOption {
	return func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
	}
}

// NamePrefixRemoval makes sure we remove name_prefix from all since it is mostly
// for Terraform functionality.
func NamePrefixRemoval() tjconfig.ResourceOption {
	return func(r *tjconfig.Resource) {
		for _, f := range r.ExternalName.OmittedFields {
			if f == "name_prefix" {
				return
			}
		}
		r.ExternalName.OmittedFields = append(r.ExternalName.OmittedFields, "name_prefix")
	}
}

func matches(name string, regexes map[string]string) (string, bool) {
	for k, v := range regexes {
		ok, err := regexp.MatchString(k, name)
		if err != nil {
			panic(errors.Wrap(err, "cannot run regular expression match"))
		}
		if ok {
			return v, true
		}
	}
	return "", false
}
