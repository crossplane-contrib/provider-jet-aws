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

package common

import (
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
)

const (
	SelfPackagePath = "github.com/crossplane-contrib/provider-tf-aws/config/common"
)

var (
	PathExternalNameAsName = SelfPackagePath + ".ExternalNameAsName"
	PathARNExtractor       = SelfPackagePath + ".ARNExtractor()"
)

// ExternalNameAsName is used by the resources whose schema includes a name field.
func ExternalNameAsName(base map[string]interface{}, externalName string) {
	base["name"] = externalName
}

func ARNExtractor() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			// todo(hasan): should we log this error?
			return ""
		}
		r, err := paved.GetString("status.atProvider.arn")
		if err != nil {
			// todo(hasan): should we log this error?
			return ""
		}
		return r
	}
}
