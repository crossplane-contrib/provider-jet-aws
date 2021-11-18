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

	tjconfig "github.com/crossplane-contrib/terrajet/pkg/config"
	"github.com/crossplane-contrib/terrajet/pkg/terraform"

	attachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/autoscaling/attachment"
	autoscalinggroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/autoscaling/autoscalinggroup"
	volume "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ebs/volume"
	elasticip "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/elasticip"
	instance "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/instance"
	ipv4cidrblockassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/ipv4cidrblockassociation"
	launchtemplate "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/launchtemplate"
	networkinterface "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/networkinterface"
	route "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/route"
	routetable "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/routetable"
	routetableassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/routetableassociation"
	securitygroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/securitygroup"
	securitygrouprule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/securitygrouprule"
	subnet "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/subnet"
	transitgateway "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/transitgateway"
	transitgatewayroute "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/transitgatewayroute"
	transitgatewayroutetable "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/transitgatewayroutetable"
	transitgatewayroutetableassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/transitgatewayroutetableassociation"
	transitgatewayroutetablepropagation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/transitgatewayroutetablepropagation"
	transitgatewayvpcattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/transitgatewayvpcattachment"
	transitgatewayvpcattachmentaccepter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/transitgatewayvpcattachmentaccepter"
	vpc "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/vpc"
	vpcendpoint "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/vpcendpoint"
	vpcpeeringconnection "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/vpcpeeringconnection"
	repository "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ecr/repository"
	repositoryecrpublic "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ecrpublic/repository"
	capacityprovider "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ecs/capacityprovider"
	cluster "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ecs/cluster"
	service "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ecs/service"
	taskdefinition "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ecs/taskdefinition"
	addon "github.com/crossplane-contrib/provider-tf-aws/internal/controller/eks/addon"
	clustereks "github.com/crossplane-contrib/provider-tf-aws/internal/controller/eks/cluster"
	fargateprofile "github.com/crossplane-contrib/provider-tf-aws/internal/controller/eks/fargateprofile"
	identityproviderconfig "github.com/crossplane-contrib/provider-tf-aws/internal/controller/eks/identityproviderconfig"
	nodegroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/eks/nodegroup"
	clusterelasticache "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elasticache/cluster"
	parametergroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elasticache/parametergroup"
	replicationgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elasticache/replicationgroup"
	loadbalancer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elasticloadbalancing/loadbalancer"
	loadbalancerlistener "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elasticloadbalancing/loadbalancerlistener"
	targetgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elasticloadbalancing/targetgroup"
	targetgroupattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elasticloadbalancing/targetgroupattachment"
	accesskey "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/accesskey"
	group "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/group"
	grouppolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/grouppolicy"
	grouppolicyattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/grouppolicyattachment"
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
	key "github.com/crossplane-contrib/provider-tf-aws/internal/controller/kms/key"
	providerconfig "github.com/crossplane-contrib/provider-tf-aws/internal/controller/providerconfig"
	dbcluster "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/dbcluster"
	dbinstance "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/dbinstance"
	dbparametergroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/dbparametergroup"
	bucket "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3/bucket"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, l logging.Logger, wl workqueue.RateLimiter, ps terraform.SetupFn, ws *terraform.WorkspaceStore, cfg *tjconfig.Provider, concurrency int) error {
	for _, setup := range []func(ctrl.Manager, logging.Logger, workqueue.RateLimiter, terraform.SetupFn, *terraform.WorkspaceStore, *tjconfig.Provider, int) error{
		attachment.Setup,
		autoscalinggroup.Setup,
		volume.Setup,
		elasticip.Setup,
		instance.Setup,
		ipv4cidrblockassociation.Setup,
		launchtemplate.Setup,
		networkinterface.Setup,
		route.Setup,
		routetable.Setup,
		routetableassociation.Setup,
		securitygroup.Setup,
		securitygrouprule.Setup,
		subnet.Setup,
		transitgateway.Setup,
		transitgatewayroute.Setup,
		transitgatewayroutetable.Setup,
		transitgatewayroutetableassociation.Setup,
		transitgatewayroutetablepropagation.Setup,
		transitgatewayvpcattachment.Setup,
		transitgatewayvpcattachmentaccepter.Setup,
		vpc.Setup,
		vpcendpoint.Setup,
		vpcpeeringconnection.Setup,
		repository.Setup,
		repositoryecrpublic.Setup,
		capacityprovider.Setup,
		cluster.Setup,
		service.Setup,
		taskdefinition.Setup,
		addon.Setup,
		clustereks.Setup,
		fargateprofile.Setup,
		identityproviderconfig.Setup,
		nodegroup.Setup,
		clusterelasticache.Setup,
		parametergroup.Setup,
		replicationgroup.Setup,
		loadbalancer.Setup,
		loadbalancerlistener.Setup,
		targetgroup.Setup,
		targetgroupattachment.Setup,
		accesskey.Setup,
		group.Setup,
		grouppolicy.Setup,
		grouppolicyattachment.Setup,
		instanceprofile.Setup,
		policy.Setup,
		policyattachment.Setup,
		role.Setup,
		rolepolicy.Setup,
		rolepolicyattachment.Setup,
		user.Setup,
		usergroupmembership.Setup,
		userpolicy.Setup,
		userpolicyattachment.Setup,
		key.Setup,
		providerconfig.Setup,
		dbcluster.Setup,
		dbinstance.Setup,
		dbparametergroup.Setup,
		bucket.Setup,
	} {
		if err := setup(mgr, l, wl, ps, ws, cfg, concurrency); err != nil {
			return err
		}
	}
	return nil
}
