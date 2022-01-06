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
	// SelfPackagePath is the golang path for this package.
	SelfPackagePath = "github.com/crossplane-contrib/provider-jet-aws/config/common"

	// PathARNExtractor is the golang path to ARNExtractor function
	// in this package.
	PathARNExtractor = SelfPackagePath + ".ARNExtractor()"

	// VersionV1Alpha2 is used as minimum version for all manually configured
	// resources.
	VersionV1Alpha2 = "v1alpha2"
)

// ARNExtractor extracts ARN of the resources from "status.atProvider.arn" which
// is quite common among all AWS resources.
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

// func GetExternalNameFromARN() config.GetExternalNameFn {
//	return func(tfstate map[string]interface{}) (string, error) {
//		arn, ok := tfstate["id"].(string)
//		if !ok || arn == "" {
//			return "", errors.New("cannot get id from tfstate")
//		}
//		// arn:partition:service:region:account-id:resource-id
//		// arn:partition:service:region:account-id:resource-type/resource-id
//		// arn:partition:service:region:account-id:resource-type:resource-id
//		w := strings.Split(arn, ":")
//		name := w[len(w)-1]
//		if strings.Contains(name, "/") {
//			ids := strings.Split(name, "/")
//			return ids[len(ids)-1], nil
//		}
//		return name, nil
//	}
// }
