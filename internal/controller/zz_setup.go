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
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

	alertmanagerdefinition "github.com/crossplane-contrib/provider-jet-aws/internal/controller/amp/alertmanagerdefinition"
	rulegroupnamespace "github.com/crossplane-contrib/provider-jet-aws/internal/controller/amp/rulegroupnamespace"
	workspace "github.com/crossplane-contrib/provider-jet-aws/internal/controller/amp/workspace"
	attachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/autoscaling/attachment"
	autoscalinggroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/autoscaling/autoscalinggroup"
	ebsvolume "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/ebsvolume"
	eip "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/eip"
	instance "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/instance"
	internetgateway "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/internetgateway"
	launchtemplate "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/launchtemplate"
	networkinterface "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/networkinterface"
	route "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/route"
	routetable "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/routetable"
	routetableassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/routetableassociation"
	securitygroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/securitygroup"
	securitygrouprule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/securitygrouprule"
	subnet "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/subnet"
	transitgateway "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/transitgateway"
	transitgatewayroute "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/transitgatewayroute"
	transitgatewayroutetable "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/transitgatewayroutetable"
	transitgatewayroutetableassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/transitgatewayroutetableassociation"
	transitgatewayroutetablepropagation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/transitgatewayroutetablepropagation"
	transitgatewayvpcattachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/transitgatewayvpcattachment"
	transitgatewayvpcattachmentaccepter "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/transitgatewayvpcattachmentaccepter"
	vpc "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/vpc"
	vpcendpoint "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/vpcendpoint"
	vpcipv4cidrblockassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/vpcipv4cidrblockassociation"
	vpcpeeringconnection "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/vpcpeeringconnection"
	vpcpeeringconnectionaccepter "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/vpcpeeringconnectionaccepter"
	repository "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ecr/repository"
	repositoryecrpublic "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ecrpublic/repository"
	capacityprovider "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ecs/capacityprovider"
	cluster "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ecs/cluster"
	service "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ecs/service"
	taskdefinition "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ecs/taskdefinition"
	addon "github.com/crossplane-contrib/provider-jet-aws/internal/controller/eks/addon"
	clustereks "github.com/crossplane-contrib/provider-jet-aws/internal/controller/eks/cluster"
	fargateprofile "github.com/crossplane-contrib/provider-jet-aws/internal/controller/eks/fargateprofile"
	identityproviderconfig "github.com/crossplane-contrib/provider-jet-aws/internal/controller/eks/identityproviderconfig"
	nodegroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/eks/nodegroup"
	clusterelasticache "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticache/cluster"
	parametergroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticache/parametergroup"
	replicationgroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticache/replicationgroup"
	subnetgroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticache/subnetgroup"
	user "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticache/user"
	usergroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticache/usergroup"
	lb "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elbv2/lb"
	lblistener "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elbv2/lblistener"
	lbtargetgroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elbv2/lbtargetgroup"
	lbtargetgroupattachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elbv2/lbtargetgroupattachment"
	accelerator "github.com/crossplane-contrib/provider-jet-aws/internal/controller/globalaccelerator/accelerator"
	endpointgroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/globalaccelerator/endpointgroup"
	listener "github.com/crossplane-contrib/provider-jet-aws/internal/controller/globalaccelerator/listener"
	roleassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/grafana/roleassociation"
	workspacegrafana "github.com/crossplane-contrib/provider-jet-aws/internal/controller/grafana/workspace"
	workspacesamlconfiguration "github.com/crossplane-contrib/provider-jet-aws/internal/controller/grafana/workspacesamlconfiguration"
	accesskey "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/accesskey"
	group "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/group"
	grouppolicyattachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/grouppolicyattachment"
	instanceprofile "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/instanceprofile"
	openidconnectprovider "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/openidconnectprovider"
	policy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/policy"
	role "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/role"
	rolepolicyattachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/rolepolicyattachment"
	useriam "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/user"
	usergroupmembership "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/usergroupmembership"
	userpolicyattachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/userpolicyattachment"
	key "github.com/crossplane-contrib/provider-jet-aws/internal/controller/kms/key"
	broker "github.com/crossplane-contrib/provider-jet-aws/internal/controller/mq/broker"
	configuration "github.com/crossplane-contrib/provider-jet-aws/internal/controller/mq/configuration"
	clusterneptune "github.com/crossplane-contrib/provider-jet-aws/internal/controller/neptune/cluster"
	clusterendpoint "github.com/crossplane-contrib/provider-jet-aws/internal/controller/neptune/clusterendpoint"
	clusterinstance "github.com/crossplane-contrib/provider-jet-aws/internal/controller/neptune/clusterinstance"
	clusterparametergroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/neptune/clusterparametergroup"
	clustersnapshot "github.com/crossplane-contrib/provider-jet-aws/internal/controller/neptune/clustersnapshot"
	eventsubscription "github.com/crossplane-contrib/provider-jet-aws/internal/controller/neptune/eventsubscription"
	parametergroupneptune "github.com/crossplane-contrib/provider-jet-aws/internal/controller/neptune/parametergroup"
	subnetgroupneptune "github.com/crossplane-contrib/provider-jet-aws/internal/controller/neptune/subnetgroup"
	providerconfig "github.com/crossplane-contrib/provider-jet-aws/internal/controller/providerconfig"
	clusterrds "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/cluster"
	instancerds "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/instance"
	parametergrouprds "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/parametergroup"
	subnetgrouprds "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/subnetgroup"
	delegationset "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/delegationset"
	healthcheck "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/healthcheck"
	hostedzonednssec "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/hostedzonednssec"
	keysigningkey "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/keysigningkey"
	querylog "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/querylog"
	record "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/record"
	vpcassociationauthorization "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/vpcassociationauthorization"
	zone "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/zone"
	zoneassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/zoneassociation"
	dnssecconfig "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53resolver/dnssecconfig"
	endpoint "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53resolver/endpoint"
	firewallconfig "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53resolver/firewallconfig"
	firewalldomainlist "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53resolver/firewalldomainlist"
	firewallrule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53resolver/firewallrule"
	firewallrulegroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53resolver/firewallrulegroup"
	firewallrulegroupassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53resolver/firewallrulegroupassociation"
	querylogconfig "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53resolver/querylogconfig"
	querylogconfigassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53resolver/querylogconfigassociation"
	rule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53resolver/rule"
	ruleassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53resolver/ruleassociation"
	bucket "github.com/crossplane-contrib/provider-jet-aws/internal/controller/s3/bucket"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alertmanagerdefinition.Setup,
		rulegroupnamespace.Setup,
		workspace.Setup,
		attachment.Setup,
		autoscalinggroup.Setup,
		ebsvolume.Setup,
		eip.Setup,
		instance.Setup,
		internetgateway.Setup,
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
		vpcipv4cidrblockassociation.Setup,
		vpcpeeringconnection.Setup,
		vpcpeeringconnectionaccepter.Setup,
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
		subnetgroup.Setup,
		user.Setup,
		usergroup.Setup,
		lb.Setup,
		lblistener.Setup,
		lbtargetgroup.Setup,
		lbtargetgroupattachment.Setup,
		accelerator.Setup,
		endpointgroup.Setup,
		listener.Setup,
		roleassociation.Setup,
		workspacegrafana.Setup,
		workspacesamlconfiguration.Setup,
		accesskey.Setup,
		group.Setup,
		grouppolicyattachment.Setup,
		instanceprofile.Setup,
		openidconnectprovider.Setup,
		policy.Setup,
		role.Setup,
		rolepolicyattachment.Setup,
		useriam.Setup,
		usergroupmembership.Setup,
		userpolicyattachment.Setup,
		key.Setup,
		broker.Setup,
		configuration.Setup,
		clusterneptune.Setup,
		clusterendpoint.Setup,
		clusterinstance.Setup,
		clusterparametergroup.Setup,
		clustersnapshot.Setup,
		eventsubscription.Setup,
		parametergroupneptune.Setup,
		subnetgroupneptune.Setup,
		providerconfig.Setup,
		clusterrds.Setup,
		instancerds.Setup,
		parametergrouprds.Setup,
		subnetgrouprds.Setup,
		delegationset.Setup,
		healthcheck.Setup,
		hostedzonednssec.Setup,
		keysigningkey.Setup,
		querylog.Setup,
		record.Setup,
		vpcassociationauthorization.Setup,
		zone.Setup,
		zoneassociation.Setup,
		dnssecconfig.Setup,
		endpoint.Setup,
		firewallconfig.Setup,
		firewalldomainlist.Setup,
		firewallrule.Setup,
		firewallrulegroup.Setup,
		firewallrulegroupassociation.Setup,
		querylogconfig.Setup,
		querylogconfigassociation.Setup,
		rule.Setup,
		ruleassociation.Setup,
		bucket.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
