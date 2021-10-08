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

package controller

import (
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/crossplane-runtime/pkg/logging"

	"github.com/crossplane-contrib/terrajet/pkg/terraform"

	iamaccesskey "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamaccesskey"
	iamgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamgroup"
	iamgrouppolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamgrouppolicy"
	iamgrouppolicyattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamgrouppolicyattachment"
	iaminstanceprofile "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iaminstanceprofile"
	iampolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iampolicy"
	iampolicyattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iampolicyattachment"
	iamrole "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamrole"
	iamrolepolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamrolepolicy"
	iamrolepolicyattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamrolepolicyattachment"
	iamuser "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamuser"
	iamusergroupmembership "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamusergroupmembership"
	iamuserpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamuserpolicy"
	iamuserpolicyattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamuserpolicyattachment"
	rdscluster "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/rdscluster"
	s3bucket "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3/s3bucket"
	tfaws "github.com/crossplane-contrib/provider-tf-aws/internal/controller/tfaws"
	vpc "github.com/crossplane-contrib/provider-tf-aws/internal/controller/vpc/vpc"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, l logging.Logger, wl workqueue.RateLimiter, ps terraform.SetupFn, ws *terraform.WorkspaceStore, concurrency int) error {
	for _, setup := range []func(ctrl.Manager, logging.Logger, workqueue.RateLimiter, terraform.SetupFn, *terraform.WorkspaceStore, int) error{
		iamaccesskey.Setup,
		iamgroup.Setup,
		iamgrouppolicy.Setup,
		iamgrouppolicyattachment.Setup,
		iaminstanceprofile.Setup,
		iampolicy.Setup,
		iampolicyattachment.Setup,
		iamrole.Setup,
		iamrolepolicy.Setup,
		iamrolepolicyattachment.Setup,
		iamuser.Setup,
		iamusergroupmembership.Setup,
		iamuserpolicy.Setup,
		iamuserpolicyattachment.Setup,
		rdscluster.Setup,
		s3bucket.Setup,
		tfaws.Setup,
		vpc.Setup,
	} {
		if err := setup(mgr, l, wl, ps, ws, concurrency); err != nil {
			return err
		}
	}
	return nil
}
