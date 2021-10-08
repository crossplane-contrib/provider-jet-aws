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

	addon "github.com/crossplane-contrib/provider-tf-aws/internal/controller/eks/addon"
	cluster "github.com/crossplane-contrib/provider-tf-aws/internal/controller/eks/cluster"
	fargateprofile "github.com/crossplane-contrib/provider-tf-aws/internal/controller/eks/fargateprofile"
	identityproviderconfig "github.com/crossplane-contrib/provider-tf-aws/internal/controller/eks/identityproviderconfig"
	nodegroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/eks/nodegroup"
	clusterelasticache "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elasticache/cluster"
	parametergroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elasticache/parametergroup"
	replicationgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elasticache/replicationgroup"
	accesskey "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/accesskey"
	grouppolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/grouppolicy"
	grouppolicyattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/grouppolicyattachment"
	iamgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamgroup"
	instanceprofile "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/instanceprofile"
	policy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/policy"
	policyattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/policyattachment"
	role "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/role"
	rolepolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/rolepolicy"
	rolepolicyattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/rolepolicyattachment"
	user "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/user"
	usergroupmembership "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/usergroupmembership"
	userpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/userpolicy"
	userpolicyattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/userpolicyattachment"
	clusterrds "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/cluster"
	bucket "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3/bucket"
	tfaws "github.com/crossplane-contrib/provider-tf-aws/internal/controller/tfaws"
	securitygroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/vpc/securitygroup"
	securitygrouprule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/vpc/securitygrouprule"
	subnet "github.com/crossplane-contrib/provider-tf-aws/internal/controller/vpc/subnet"
	vpc "github.com/crossplane-contrib/provider-tf-aws/internal/controller/vpc/vpc"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, l logging.Logger, wl workqueue.RateLimiter, ps terraform.SetupFn, ws *terraform.WorkspaceStore, concurrency int) error {
	for _, setup := range []func(ctrl.Manager, logging.Logger, workqueue.RateLimiter, terraform.SetupFn, *terraform.WorkspaceStore, int) error{
		accesskey.Setup,
		addon.Setup,
		bucket.Setup,
		cluster.Setup,
		clusterelasticache.Setup,
		clusterrds.Setup,
		fargateprofile.Setup,
		grouppolicy.Setup,
		grouppolicyattachment.Setup,
		iamgroup.Setup,
		identityproviderconfig.Setup,
		instanceprofile.Setup,
		parametergroup.Setup,
		nodegroup.Setup,
		policy.Setup,
		policyattachment.Setup,
		replicationgroup.Setup,
		role.Setup,
		rolepolicy.Setup,
		rolepolicyattachment.Setup,
		securitygroup.Setup,
		securitygrouprule.Setup,
		subnet.Setup,
		tfaws.Setup,
		user.Setup,
		usergroupmembership.Setup,
		userpolicy.Setup,
		userpolicyattachment.Setup,
		vpc.Setup,
	} {
		if err := setup(mgr, l, wl, ps, ws, concurrency); err != nil {
			return err
		}
	}
	return nil
}
