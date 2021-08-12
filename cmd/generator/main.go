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
	"strings"

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
	groupGen := pipeline.NewVersionGenerator(wd, "ec2.aws.tf.crossplane.io", "v1alpha1")
	if err := groupGen.Generate(); err != nil {
		panic(errors.Wrap(err, "cannot generate version files"))
	}
	list := []string{
		"aws_vpc",
		"aws_instance",
		"aws_db_instance",
	}
	for _, r := range list {
		kind := strcase.ToCamel(strings.TrimPrefix(r, "aws_"))
		crdGen := pipeline.NewCRDGenerator(wd, "ec2.aws.tf.crossplane.io", "v1alpha1", kind, aws.Provider().ResourcesMap[r])
		if err := crdGen.Generate(); err != nil {
			panic(errors.Wrap(err, "cannot generate crd"))
		}
	}
}
