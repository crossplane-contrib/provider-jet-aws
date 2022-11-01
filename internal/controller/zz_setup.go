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

	alertmanagerdefinition "github.com/dkb-bank/provider-jet-aws/internal/controller/amp/alertmanagerdefinition"
	rulegroupnamespace "github.com/dkb-bank/provider-jet-aws/internal/controller/amp/rulegroupnamespace"
	workspace "github.com/dkb-bank/provider-jet-aws/internal/controller/amp/workspace"
	attachment "github.com/dkb-bank/provider-jet-aws/internal/controller/autoscaling/attachment"
	autoscalinggroup "github.com/dkb-bank/provider-jet-aws/internal/controller/autoscaling/autoscalinggroup"
	group "github.com/dkb-bank/provider-jet-aws/internal/controller/cloudwatchlogs/group"
	ebsvolume "github.com/dkb-bank/provider-jet-aws/internal/controller/ec2/ebsvolume"
	eip "github.com/dkb-bank/provider-jet-aws/internal/controller/ec2/eip"
	instance "github.com/dkb-bank/provider-jet-aws/internal/controller/ec2/instance"
	internetgateway "github.com/dkb-bank/provider-jet-aws/internal/controller/ec2/internetgateway"
	launchtemplate "github.com/dkb-bank/provider-jet-aws/internal/controller/ec2/launchtemplate"
	networkinterface "github.com/dkb-bank/provider-jet-aws/internal/controller/ec2/networkinterface"
	route "github.com/dkb-bank/provider-jet-aws/internal/controller/ec2/route"
	routetable "github.com/dkb-bank/provider-jet-aws/internal/controller/ec2/routetable"
	routetableassociation "github.com/dkb-bank/provider-jet-aws/internal/controller/ec2/routetableassociation"
	securitygroup "github.com/dkb-bank/provider-jet-aws/internal/controller/ec2/securitygroup"
	securitygrouprule "github.com/dkb-bank/provider-jet-aws/internal/controller/ec2/securitygrouprule"
	subnet "github.com/dkb-bank/provider-jet-aws/internal/controller/ec2/subnet"
	tag "github.com/dkb-bank/provider-jet-aws/internal/controller/ec2/tag"
	transitgateway "github.com/dkb-bank/provider-jet-aws/internal/controller/ec2/transitgateway"
	transitgatewayroute "github.com/dkb-bank/provider-jet-aws/internal/controller/ec2/transitgatewayroute"
	transitgatewayroutetable "github.com/dkb-bank/provider-jet-aws/internal/controller/ec2/transitgatewayroutetable"
	transitgatewayroutetableassociation "github.com/dkb-bank/provider-jet-aws/internal/controller/ec2/transitgatewayroutetableassociation"
	transitgatewayroutetablepropagation "github.com/dkb-bank/provider-jet-aws/internal/controller/ec2/transitgatewayroutetablepropagation"
	transitgatewayvpcattachment "github.com/dkb-bank/provider-jet-aws/internal/controller/ec2/transitgatewayvpcattachment"
	transitgatewayvpcattachmentaccepter "github.com/dkb-bank/provider-jet-aws/internal/controller/ec2/transitgatewayvpcattachmentaccepter"
	vpc "github.com/dkb-bank/provider-jet-aws/internal/controller/ec2/vpc"
	vpcendpoint "github.com/dkb-bank/provider-jet-aws/internal/controller/ec2/vpcendpoint"
	vpcipv4cidrblockassociation "github.com/dkb-bank/provider-jet-aws/internal/controller/ec2/vpcipv4cidrblockassociation"
	vpcpeeringconnection "github.com/dkb-bank/provider-jet-aws/internal/controller/ec2/vpcpeeringconnection"
	repository "github.com/dkb-bank/provider-jet-aws/internal/controller/ecr/repository"
	repositoryecrpublic "github.com/dkb-bank/provider-jet-aws/internal/controller/ecrpublic/repository"
	capacityprovider "github.com/dkb-bank/provider-jet-aws/internal/controller/ecs/capacityprovider"
	cluster "github.com/dkb-bank/provider-jet-aws/internal/controller/ecs/cluster"
	service "github.com/dkb-bank/provider-jet-aws/internal/controller/ecs/service"
	taskdefinition "github.com/dkb-bank/provider-jet-aws/internal/controller/ecs/taskdefinition"
	addon "github.com/dkb-bank/provider-jet-aws/internal/controller/eks/addon"
	clustereks "github.com/dkb-bank/provider-jet-aws/internal/controller/eks/cluster"
	fargateprofile "github.com/dkb-bank/provider-jet-aws/internal/controller/eks/fargateprofile"
	identityproviderconfig "github.com/dkb-bank/provider-jet-aws/internal/controller/eks/identityproviderconfig"
	nodegroup "github.com/dkb-bank/provider-jet-aws/internal/controller/eks/nodegroup"
	clusterelasticache "github.com/dkb-bank/provider-jet-aws/internal/controller/elasticache/cluster"
	parametergroup "github.com/dkb-bank/provider-jet-aws/internal/controller/elasticache/parametergroup"
	replicationgroup "github.com/dkb-bank/provider-jet-aws/internal/controller/elasticache/replicationgroup"
	subnetgroup "github.com/dkb-bank/provider-jet-aws/internal/controller/elasticache/subnetgroup"
	user "github.com/dkb-bank/provider-jet-aws/internal/controller/elasticache/user"
	usergroup "github.com/dkb-bank/provider-jet-aws/internal/controller/elasticache/usergroup"
	lb "github.com/dkb-bank/provider-jet-aws/internal/controller/elbv2/lb"
	lblistener "github.com/dkb-bank/provider-jet-aws/internal/controller/elbv2/lblistener"
	lblistenercertificate "github.com/dkb-bank/provider-jet-aws/internal/controller/elbv2/lblistenercertificate"
	lblistenerrule "github.com/dkb-bank/provider-jet-aws/internal/controller/elbv2/lblistenerrule"
	lbtargetgroup "github.com/dkb-bank/provider-jet-aws/internal/controller/elbv2/lbtargetgroup"
	lbtargetgroupattachment "github.com/dkb-bank/provider-jet-aws/internal/controller/elbv2/lbtargetgroupattachment"
	accelerator "github.com/dkb-bank/provider-jet-aws/internal/controller/globalaccelerator/accelerator"
	endpointgroup "github.com/dkb-bank/provider-jet-aws/internal/controller/globalaccelerator/endpointgroup"
	listener "github.com/dkb-bank/provider-jet-aws/internal/controller/globalaccelerator/listener"
	roleassociation "github.com/dkb-bank/provider-jet-aws/internal/controller/grafana/roleassociation"
	workspacegrafana "github.com/dkb-bank/provider-jet-aws/internal/controller/grafana/workspace"
	workspacesamlconfiguration "github.com/dkb-bank/provider-jet-aws/internal/controller/grafana/workspacesamlconfiguration"
	accesskey "github.com/dkb-bank/provider-jet-aws/internal/controller/iam/accesskey"
	groupiam "github.com/dkb-bank/provider-jet-aws/internal/controller/iam/group"
	grouppolicyattachment "github.com/dkb-bank/provider-jet-aws/internal/controller/iam/grouppolicyattachment"
	instanceprofile "github.com/dkb-bank/provider-jet-aws/internal/controller/iam/instanceprofile"
	openidconnectprovider "github.com/dkb-bank/provider-jet-aws/internal/controller/iam/openidconnectprovider"
	policy "github.com/dkb-bank/provider-jet-aws/internal/controller/iam/policy"
	role "github.com/dkb-bank/provider-jet-aws/internal/controller/iam/role"
	rolepolicyattachment "github.com/dkb-bank/provider-jet-aws/internal/controller/iam/rolepolicyattachment"
	useriam "github.com/dkb-bank/provider-jet-aws/internal/controller/iam/user"
	usergroupmembership "github.com/dkb-bank/provider-jet-aws/internal/controller/iam/usergroupmembership"
	userpolicyattachment "github.com/dkb-bank/provider-jet-aws/internal/controller/iam/userpolicyattachment"
	key "github.com/dkb-bank/provider-jet-aws/internal/controller/kms/key"
	broker "github.com/dkb-bank/provider-jet-aws/internal/controller/mq/broker"
	configuration "github.com/dkb-bank/provider-jet-aws/internal/controller/mq/configuration"
	clusterneptune "github.com/dkb-bank/provider-jet-aws/internal/controller/neptune/cluster"
	clusterendpoint "github.com/dkb-bank/provider-jet-aws/internal/controller/neptune/clusterendpoint"
	clusterinstance "github.com/dkb-bank/provider-jet-aws/internal/controller/neptune/clusterinstance"
	clusterparametergroup "github.com/dkb-bank/provider-jet-aws/internal/controller/neptune/clusterparametergroup"
	clustersnapshot "github.com/dkb-bank/provider-jet-aws/internal/controller/neptune/clustersnapshot"
	eventsubscription "github.com/dkb-bank/provider-jet-aws/internal/controller/neptune/eventsubscription"
	parametergroupneptune "github.com/dkb-bank/provider-jet-aws/internal/controller/neptune/parametergroup"
	subnetgroupneptune "github.com/dkb-bank/provider-jet-aws/internal/controller/neptune/subnetgroup"
	firewall "github.com/dkb-bank/provider-jet-aws/internal/controller/networkfirewall/firewall"
	firewallpolicy "github.com/dkb-bank/provider-jet-aws/internal/controller/networkfirewall/firewallpolicy"
	loggingconfiguration "github.com/dkb-bank/provider-jet-aws/internal/controller/networkfirewall/loggingconfiguration"
	rulegroup "github.com/dkb-bank/provider-jet-aws/internal/controller/networkfirewall/rulegroup"
	providerconfig "github.com/dkb-bank/provider-jet-aws/internal/controller/providerconfig"
	clusterrds "github.com/dkb-bank/provider-jet-aws/internal/controller/rds/cluster"
	instancerds "github.com/dkb-bank/provider-jet-aws/internal/controller/rds/instance"
	parametergrouprds "github.com/dkb-bank/provider-jet-aws/internal/controller/rds/parametergroup"
	subnetgrouprds "github.com/dkb-bank/provider-jet-aws/internal/controller/rds/subnetgroup"
	delegationset "github.com/dkb-bank/provider-jet-aws/internal/controller/route53/delegationset"
	healthcheck "github.com/dkb-bank/provider-jet-aws/internal/controller/route53/healthcheck"
	hostedzonednssec "github.com/dkb-bank/provider-jet-aws/internal/controller/route53/hostedzonednssec"
	keysigningkey "github.com/dkb-bank/provider-jet-aws/internal/controller/route53/keysigningkey"
	querylog "github.com/dkb-bank/provider-jet-aws/internal/controller/route53/querylog"
	record "github.com/dkb-bank/provider-jet-aws/internal/controller/route53/record"
	vpcassociationauthorization "github.com/dkb-bank/provider-jet-aws/internal/controller/route53/vpcassociationauthorization"
	zone "github.com/dkb-bank/provider-jet-aws/internal/controller/route53/zone"
	zoneassociation "github.com/dkb-bank/provider-jet-aws/internal/controller/route53/zoneassociation"
	dnssecconfig "github.com/dkb-bank/provider-jet-aws/internal/controller/route53resolver/dnssecconfig"
	endpoint "github.com/dkb-bank/provider-jet-aws/internal/controller/route53resolver/endpoint"
	firewallconfig "github.com/dkb-bank/provider-jet-aws/internal/controller/route53resolver/firewallconfig"
	firewalldomainlist "github.com/dkb-bank/provider-jet-aws/internal/controller/route53resolver/firewalldomainlist"
	firewallrule "github.com/dkb-bank/provider-jet-aws/internal/controller/route53resolver/firewallrule"
	firewallrulegroup "github.com/dkb-bank/provider-jet-aws/internal/controller/route53resolver/firewallrulegroup"
	firewallrulegroupassociation "github.com/dkb-bank/provider-jet-aws/internal/controller/route53resolver/firewallrulegroupassociation"
	querylogconfig "github.com/dkb-bank/provider-jet-aws/internal/controller/route53resolver/querylogconfig"
	querylogconfigassociation "github.com/dkb-bank/provider-jet-aws/internal/controller/route53resolver/querylogconfigassociation"
	rule "github.com/dkb-bank/provider-jet-aws/internal/controller/route53resolver/rule"
	ruleassociation "github.com/dkb-bank/provider-jet-aws/internal/controller/route53resolver/ruleassociation"
	bucket "github.com/dkb-bank/provider-jet-aws/internal/controller/s3/bucket"
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
		group.Setup,
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
		tag.Setup,
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
		lblistenercertificate.Setup,
		lblistenerrule.Setup,
		lbtargetgroup.Setup,
		lbtargetgroupattachment.Setup,
		accelerator.Setup,
		endpointgroup.Setup,
		listener.Setup,
		roleassociation.Setup,
		workspacegrafana.Setup,
		workspacesamlconfiguration.Setup,
		accesskey.Setup,
		groupiam.Setup,
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
		firewall.Setup,
		firewallpolicy.Setup,
		loggingconfiguration.Setup,
		rulegroup.Setup,
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
