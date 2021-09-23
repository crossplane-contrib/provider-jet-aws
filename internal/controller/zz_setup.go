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

	defaultnetworkacl "github.com/crossplane-contrib/provider-tf-aws/internal/controller/default/defaultnetworkacl"
	defaultroutetable "github.com/crossplane-contrib/provider-tf-aws/internal/controller/default/defaultroutetable"
	defaultsecuritygroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/default/defaultsecuritygroup"
	defaultsubnet "github.com/crossplane-contrib/provider-tf-aws/internal/controller/default/defaultsubnet"
	defaultvpc "github.com/crossplane-contrib/provider-tf-aws/internal/controller/default/defaultvpc"
	defaultvpcdhcpoptions "github.com/crossplane-contrib/provider-tf-aws/internal/controller/default/defaultvpcdhcpoptions"
	ec2availabilityzonegroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/ec2availabilityzonegroup"
	ec2capacityreservation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/ec2capacityreservation"
	ec2carriergateway "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/ec2carriergateway"
	ec2clientvpnauthorizationrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/ec2clientvpnauthorizationrule"
	ec2clientvpnendpoint "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/ec2clientvpnendpoint"
	ec2clientvpnnetworkassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/ec2clientvpnnetworkassociation"
	ec2clientvpnroute "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/ec2clientvpnroute"
	ec2fleet "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/ec2fleet"
	ec2localgatewayroute "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/ec2localgatewayroute"
	ec2localgatewayroutetablevpcassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/ec2localgatewayroutetablevpcassociation"
	ec2managedprefixlist "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/ec2managedprefixlist"
	ec2tag "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/ec2tag"
	ec2trafficmirrorfilter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/ec2trafficmirrorfilter"
	ec2trafficmirrorfilterrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/ec2trafficmirrorfilterrule"
	ec2trafficmirrorsession "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/ec2trafficmirrorsession"
	ec2trafficmirrortarget "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/ec2trafficmirrortarget"
	ec2transitgateway "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/ec2transitgateway"
	ec2transitgatewaypeeringattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/ec2transitgatewaypeeringattachment"
	ec2transitgatewaypeeringattachmentaccepter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/ec2transitgatewaypeeringattachmentaccepter"
	ec2transitgatewayprefixlistreference "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/ec2transitgatewayprefixlistreference"
	ec2transitgatewayroute "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/ec2transitgatewayroute"
	ec2transitgatewayroutetable "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/ec2transitgatewayroutetable"
	ec2transitgatewayroutetableassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/ec2transitgatewayroutetableassociation"
	ec2transitgatewayroutetablepropagation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/ec2transitgatewayroutetablepropagation"
	ec2transitgatewayvpcattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/ec2transitgatewayvpcattachment"
	ec2transitgatewayvpcattachmentaccepter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/ec2transitgatewayvpcattachmentaccepter"
	eip "github.com/crossplane-contrib/provider-tf-aws/internal/controller/eip/eip"
	eipassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/eip/eipassociation"
	eksaddon "github.com/crossplane-contrib/provider-tf-aws/internal/controller/eks/eksaddon"
	ekscluster "github.com/crossplane-contrib/provider-tf-aws/internal/controller/eks/ekscluster"
	eksfargateprofile "github.com/crossplane-contrib/provider-tf-aws/internal/controller/eks/eksfargateprofile"
	eksidentityproviderconfig "github.com/crossplane-contrib/provider-tf-aws/internal/controller/eks/eksidentityproviderconfig"
	eksnodegroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/eks/eksnodegroup"
	elb "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elb/elb"
	elbattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elb/elbattachment"
	iamaccesskey "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamaccesskey"
	iamaccountalias "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamaccountalias"
	iamaccountpasswordpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamaccountpasswordpolicy"
	iamgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamgroup"
	iamgroupmembership "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamgroupmembership"
	iamgrouppolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamgrouppolicy"
	iamgrouppolicyattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamgrouppolicyattachment"
	iaminstanceprofile "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iaminstanceprofile"
	iamopenidconnectprovider "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamopenidconnectprovider"
	iampolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iampolicy"
	iampolicyattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iampolicyattachment"
	iamrole "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamrole"
	iamrolepolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamrolepolicy"
	iamrolepolicyattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamrolepolicyattachment"
	iamsamlprovider "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamsamlprovider"
	iamservercertificate "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamservercertificate"
	iamservicelinkedrole "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamservicelinkedrole"
	iamuser "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamuser"
	iamusergroupmembership "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamusergroupmembership"
	iamuserloginprofile "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamuserloginprofile"
	iamuserpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamuserpolicy"
	iamuserpolicyattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamuserpolicyattachment"
	iamusersshkey "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamusersshkey"
	instance "github.com/crossplane-contrib/provider-tf-aws/internal/controller/instance/instance"
	lb "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lb/lb"
	lbcookiestickinesspolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lb/lbcookiestickinesspolicy"
	lblistener "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lb/lblistener"
	lblistenercertificate "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lb/lblistenercertificate"
	lblistenerrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lb/lblistenerrule"
	lbsslnegotiationpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lb/lbsslnegotiationpolicy"
	lbtargetgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lb/lbtargetgroup"
	lbtargetgroupattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lb/lbtargetgroupattachment"
	mainroutetableassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/main/mainroutetableassociation"
	rdscluster "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/rdscluster"
	rdsclusterendpoint "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/rdsclusterendpoint"
	rdsclusterinstance "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/rdsclusterinstance"
	rdsclusterparametergroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/rdsclusterparametergroup"
	rdsclusterroleassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/rdsclusterroleassociation"
	rdsglobalcluster "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/rdsglobalcluster"
	route "github.com/crossplane-contrib/provider-tf-aws/internal/controller/route/route"
	routetable "github.com/crossplane-contrib/provider-tf-aws/internal/controller/route/routetable"
	routetableassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/route/routetableassociation"
	route53delegationset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/route53/route53delegationset"
	route53healthcheck "github.com/crossplane-contrib/provider-tf-aws/internal/controller/route53/route53healthcheck"
	route53hostedzonednssec "github.com/crossplane-contrib/provider-tf-aws/internal/controller/route53/route53hostedzonednssec"
	route53keysigningkey "github.com/crossplane-contrib/provider-tf-aws/internal/controller/route53/route53keysigningkey"
	route53querylog "github.com/crossplane-contrib/provider-tf-aws/internal/controller/route53/route53querylog"
	route53record "github.com/crossplane-contrib/provider-tf-aws/internal/controller/route53/route53record"
	route53resolverdnssecconfig "github.com/crossplane-contrib/provider-tf-aws/internal/controller/route53/route53resolverdnssecconfig"
	route53resolverendpoint "github.com/crossplane-contrib/provider-tf-aws/internal/controller/route53/route53resolverendpoint"
	route53resolverfirewallconfig "github.com/crossplane-contrib/provider-tf-aws/internal/controller/route53/route53resolverfirewallconfig"
	route53resolverfirewalldomainlist "github.com/crossplane-contrib/provider-tf-aws/internal/controller/route53/route53resolverfirewalldomainlist"
	route53resolverfirewallrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/route53/route53resolverfirewallrule"
	route53resolverfirewallrulegroupassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/route53/route53resolverfirewallrulegroupassociation"
	route53resolverquerylogconfig "github.com/crossplane-contrib/provider-tf-aws/internal/controller/route53/route53resolverquerylogconfig"
	route53resolverquerylogconfigassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/route53/route53resolverquerylogconfigassociation"
	route53resolverrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/route53/route53resolverrule"
	route53resolverruleassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/route53/route53resolverruleassociation"
	route53vpcassociationauthorization "github.com/crossplane-contrib/provider-tf-aws/internal/controller/route53/route53vpcassociationauthorization"
	route53zone "github.com/crossplane-contrib/provider-tf-aws/internal/controller/route53/route53zone"
	route53zoneassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/route53/route53zoneassociation"
	s3accesspoint "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3/s3accesspoint"
	s3accountpublicaccessblock "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3/s3accountpublicaccessblock"
	s3bucket "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3/s3bucket"
	s3bucketanalyticsconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3/s3bucketanalyticsconfiguration"
	s3bucketinventory "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3/s3bucketinventory"
	s3bucketmetric "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3/s3bucketmetric"
	s3bucketnotification "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3/s3bucketnotification"
	s3bucketobject "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3/s3bucketobject"
	s3bucketownershipcontrols "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3/s3bucketownershipcontrols"
	s3bucketpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3/s3bucketpolicy"
	s3bucketpublicaccessblock "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3/s3bucketpublicaccessblock"
	s3objectcopy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3/s3objectcopy"
	subnet "github.com/crossplane-contrib/provider-tf-aws/internal/controller/subnet/subnet"
	tfaws "github.com/crossplane-contrib/provider-tf-aws/internal/controller/tfaws"
	vpc "github.com/crossplane-contrib/provider-tf-aws/internal/controller/vpc/vpc"
	vpcdhcpoptions "github.com/crossplane-contrib/provider-tf-aws/internal/controller/vpc/vpcdhcpoptions"
	vpcdhcpoptionsassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/vpc/vpcdhcpoptionsassociation"
	vpcendpoint "github.com/crossplane-contrib/provider-tf-aws/internal/controller/vpc/vpcendpoint"
	vpcendpointconnectionnotification "github.com/crossplane-contrib/provider-tf-aws/internal/controller/vpc/vpcendpointconnectionnotification"
	vpcendpointroutetableassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/vpc/vpcendpointroutetableassociation"
	vpcendpointservice "github.com/crossplane-contrib/provider-tf-aws/internal/controller/vpc/vpcendpointservice"
	vpcendpointserviceallowedprincipal "github.com/crossplane-contrib/provider-tf-aws/internal/controller/vpc/vpcendpointserviceallowedprincipal"
	vpcendpointsubnetassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/vpc/vpcendpointsubnetassociation"
	vpcipv4cidrblockassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/vpc/vpcipv4cidrblockassociation"
	vpcpeeringconnection "github.com/crossplane-contrib/provider-tf-aws/internal/controller/vpc/vpcpeeringconnection"
	vpcpeeringconnectionoptions "github.com/crossplane-contrib/provider-tf-aws/internal/controller/vpc/vpcpeeringconnectionoptions"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, l logging.Logger, wl workqueue.RateLimiter, concurrency int) error {
	for _, setup := range []func(ctrl.Manager, logging.Logger, workqueue.RateLimiter, int) error{
		defaultnetworkacl.Setup,
		defaultroutetable.Setup,
		defaultsecuritygroup.Setup,
		defaultsubnet.Setup,
		defaultvpc.Setup,
		defaultvpcdhcpoptions.Setup,
		ec2availabilityzonegroup.Setup,
		ec2capacityreservation.Setup,
		ec2carriergateway.Setup,
		ec2clientvpnauthorizationrule.Setup,
		ec2clientvpnendpoint.Setup,
		ec2clientvpnnetworkassociation.Setup,
		ec2clientvpnroute.Setup,
		ec2fleet.Setup,
		ec2localgatewayroute.Setup,
		ec2localgatewayroutetablevpcassociation.Setup,
		ec2managedprefixlist.Setup,
		ec2tag.Setup,
		ec2trafficmirrorfilter.Setup,
		ec2trafficmirrorfilterrule.Setup,
		ec2trafficmirrorsession.Setup,
		ec2trafficmirrortarget.Setup,
		ec2transitgateway.Setup,
		ec2transitgatewaypeeringattachment.Setup,
		ec2transitgatewaypeeringattachmentaccepter.Setup,
		ec2transitgatewayprefixlistreference.Setup,
		ec2transitgatewayroute.Setup,
		ec2transitgatewayroutetable.Setup,
		ec2transitgatewayroutetableassociation.Setup,
		ec2transitgatewayroutetablepropagation.Setup,
		ec2transitgatewayvpcattachment.Setup,
		ec2transitgatewayvpcattachmentaccepter.Setup,
		eip.Setup,
		eipassociation.Setup,
		eksaddon.Setup,
		ekscluster.Setup,
		eksfargateprofile.Setup,
		eksidentityproviderconfig.Setup,
		eksnodegroup.Setup,
		elb.Setup,
		elbattachment.Setup,
		iamaccesskey.Setup,
		iamaccountalias.Setup,
		iamaccountpasswordpolicy.Setup,
		iamgroup.Setup,
		iamgroupmembership.Setup,
		iamgrouppolicy.Setup,
		iamgrouppolicyattachment.Setup,
		iaminstanceprofile.Setup,
		iamopenidconnectprovider.Setup,
		iampolicy.Setup,
		iampolicyattachment.Setup,
		iamrole.Setup,
		iamrolepolicy.Setup,
		iamrolepolicyattachment.Setup,
		iamsamlprovider.Setup,
		iamservercertificate.Setup,
		iamservicelinkedrole.Setup,
		iamuser.Setup,
		iamusergroupmembership.Setup,
		iamuserloginprofile.Setup,
		iamuserpolicy.Setup,
		iamuserpolicyattachment.Setup,
		iamusersshkey.Setup,
		instance.Setup,
		lb.Setup,
		lbcookiestickinesspolicy.Setup,
		lblistener.Setup,
		lblistenercertificate.Setup,
		lblistenerrule.Setup,
		lbsslnegotiationpolicy.Setup,
		lbtargetgroup.Setup,
		lbtargetgroupattachment.Setup,
		mainroutetableassociation.Setup,
		rdscluster.Setup,
		rdsclusterendpoint.Setup,
		rdsclusterinstance.Setup,
		rdsclusterparametergroup.Setup,
		rdsclusterroleassociation.Setup,
		rdsglobalcluster.Setup,
		route.Setup,
		route53delegationset.Setup,
		route53healthcheck.Setup,
		route53hostedzonednssec.Setup,
		route53keysigningkey.Setup,
		route53querylog.Setup,
		route53record.Setup,
		route53resolverdnssecconfig.Setup,
		route53resolverendpoint.Setup,
		route53resolverfirewallconfig.Setup,
		route53resolverfirewalldomainlist.Setup,
		route53resolverfirewallrule.Setup,
		route53resolverfirewallrulegroupassociation.Setup,
		route53resolverquerylogconfig.Setup,
		route53resolverquerylogconfigassociation.Setup,
		route53resolverrule.Setup,
		route53resolverruleassociation.Setup,
		route53vpcassociationauthorization.Setup,
		route53zone.Setup,
		route53zoneassociation.Setup,
		routetable.Setup,
		routetableassociation.Setup,
		s3accesspoint.Setup,
		s3accountpublicaccessblock.Setup,
		s3bucket.Setup,
		s3bucketanalyticsconfiguration.Setup,
		s3bucketinventory.Setup,
		s3bucketmetric.Setup,
		s3bucketnotification.Setup,
		s3bucketobject.Setup,
		s3bucketownershipcontrols.Setup,
		s3bucketpolicy.Setup,
		s3bucketpublicaccessblock.Setup,
		s3objectcopy.Setup,
		subnet.Setup,
		tfaws.Setup,
		vpc.Setup,
		vpcdhcpoptions.Setup,
		vpcdhcpoptionsassociation.Setup,
		vpcendpoint.Setup,
		vpcendpointconnectionnotification.Setup,
		vpcendpointroutetableassociation.Setup,
		vpcendpointservice.Setup,
		vpcendpointserviceallowedprincipal.Setup,
		vpcendpointsubnetassociation.Setup,
		vpcipv4cidrblockassociation.Setup,
		vpcpeeringconnection.Setup,
		vpcpeeringconnectionoptions.Setup,
	} {
		if err := setup(mgr, l, wl, concurrency); err != nil {
			return err
		}
	}
	return nil
}
