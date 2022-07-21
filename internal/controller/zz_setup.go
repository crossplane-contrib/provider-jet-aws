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

	analyzer "github.com/crossplane-contrib/provider-jet-aws/internal/controller/accessanalyzer/analyzer"
	certificate "github.com/crossplane-contrib/provider-jet-aws/internal/controller/acm/certificate"
	certificatevalidation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/acm/certificatevalidation"
	certificateacmpca "github.com/crossplane-contrib/provider-jet-aws/internal/controller/acmpca/certificate"
	certificateauthority "github.com/crossplane-contrib/provider-jet-aws/internal/controller/acmpca/certificateauthority"
	certificateauthoritycertificate "github.com/crossplane-contrib/provider-jet-aws/internal/controller/acmpca/certificateauthoritycertificate"
	workspace "github.com/crossplane-contrib/provider-jet-aws/internal/controller/amp/workspace"
	app "github.com/crossplane-contrib/provider-jet-aws/internal/controller/amplify/app"
	backendenvironment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/amplify/backendenvironment"
	branch "github.com/crossplane-contrib/provider-jet-aws/internal/controller/amplify/branch"
	domainassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/amplify/domainassociation"
	webhook "github.com/crossplane-contrib/provider-jet-aws/internal/controller/amplify/webhook"
	account "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigateway/account"
	apikey "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigateway/apikey"
	authorizer "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigateway/authorizer"
	basepathmapping "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigateway/basepathmapping"
	clientcertificate "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigateway/clientcertificate"
	deployment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigateway/deployment"
	documentationpart "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigateway/documentationpart"
	documentationversion "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigateway/documentationversion"
	domainname "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigateway/domainname"
	gatewayresponse "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigateway/gatewayresponse"
	integration "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigateway/integration"
	integrationresponse "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigateway/integrationresponse"
	method "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigateway/method"
	methodresponse "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigateway/methodresponse"
	methodsettings "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigateway/methodsettings"
	model "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigateway/model"
	requestvalidator "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigateway/requestvalidator"
	resource "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigateway/resource"
	restapi "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigateway/restapi"
	restapipolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigateway/restapipolicy"
	stage "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigateway/stage"
	usageplan "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigateway/usageplan"
	usageplankey "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigateway/usageplankey"
	vpclink "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigateway/vpclink"
	api "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigatewayv2/api"
	apimapping "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigatewayv2/apimapping"
	authorizerapigatewayv2 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigatewayv2/authorizer"
	deploymentapigatewayv2 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigatewayv2/deployment"
	domainnameapigatewayv2 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigatewayv2/domainname"
	integrationapigatewayv2 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigatewayv2/integration"
	integrationresponseapigatewayv2 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigatewayv2/integrationresponse"
	modelapigatewayv2 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigatewayv2/model"
	route "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigatewayv2/route"
	routeresponse "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigatewayv2/routeresponse"
	stageapigatewayv2 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigatewayv2/stage"
	vpclinkapigatewayv2 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigatewayv2/vpclink"
	policy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/appautoscaling/policy"
	scheduledaction "github.com/crossplane-contrib/provider-jet-aws/internal/controller/appautoscaling/scheduledaction"
	target "github.com/crossplane-contrib/provider-jet-aws/internal/controller/appautoscaling/target"
	application "github.com/crossplane-contrib/provider-jet-aws/internal/controller/appconfig/application"
	configurationprofile "github.com/crossplane-contrib/provider-jet-aws/internal/controller/appconfig/configurationprofile"
	deploymentappconfig "github.com/crossplane-contrib/provider-jet-aws/internal/controller/appconfig/deployment"
	deploymentstrategy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/appconfig/deploymentstrategy"
	environment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/appconfig/environment"
	hostedconfigurationversion "github.com/crossplane-contrib/provider-jet-aws/internal/controller/appconfig/hostedconfigurationversion"
	gatewayroute "github.com/crossplane-contrib/provider-jet-aws/internal/controller/appmesh/gatewayroute"
	mesh "github.com/crossplane-contrib/provider-jet-aws/internal/controller/appmesh/mesh"
	routeappmesh "github.com/crossplane-contrib/provider-jet-aws/internal/controller/appmesh/route"
	virtualgateway "github.com/crossplane-contrib/provider-jet-aws/internal/controller/appmesh/virtualgateway"
	virtualnode "github.com/crossplane-contrib/provider-jet-aws/internal/controller/appmesh/virtualnode"
	virtualrouter "github.com/crossplane-contrib/provider-jet-aws/internal/controller/appmesh/virtualrouter"
	virtualservice "github.com/crossplane-contrib/provider-jet-aws/internal/controller/appmesh/virtualservice"
	autoscalingconfigurationversion "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apprunner/autoscalingconfigurationversion"
	connection "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apprunner/connection"
	customdomainassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apprunner/customdomainassociation"
	service "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apprunner/service"
	stack "github.com/crossplane-contrib/provider-jet-aws/internal/controller/appstream/stack"
	apikeyappsync "github.com/crossplane-contrib/provider-jet-aws/internal/controller/appsync/apikey"
	datasource "github.com/crossplane-contrib/provider-jet-aws/internal/controller/appsync/datasource"
	function "github.com/crossplane-contrib/provider-jet-aws/internal/controller/appsync/function"
	graphqlapi "github.com/crossplane-contrib/provider-jet-aws/internal/controller/appsync/graphqlapi"
	resolver "github.com/crossplane-contrib/provider-jet-aws/internal/controller/appsync/resolver"
	database "github.com/crossplane-contrib/provider-jet-aws/internal/controller/athena/database"
	namedquery "github.com/crossplane-contrib/provider-jet-aws/internal/controller/athena/namedquery"
	workgroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/athena/workgroup"
	attachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/autoscaling/attachment"
	autoscalinggroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/autoscaling/autoscalinggroup"
	grouptag "github.com/crossplane-contrib/provider-jet-aws/internal/controller/autoscaling/grouptag"
	launchconfiguration "github.com/crossplane-contrib/provider-jet-aws/internal/controller/autoscaling/launchconfiguration"
	lifecyclehook "github.com/crossplane-contrib/provider-jet-aws/internal/controller/autoscaling/lifecyclehook"
	notification "github.com/crossplane-contrib/provider-jet-aws/internal/controller/autoscaling/notification"
	policyautoscaling "github.com/crossplane-contrib/provider-jet-aws/internal/controller/autoscaling/policy"
	schedule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/autoscaling/schedule"
	scalingplan "github.com/crossplane-contrib/provider-jet-aws/internal/controller/autoscalingplans/scalingplan"
	globalsettings "github.com/crossplane-contrib/provider-jet-aws/internal/controller/backup/globalsettings"
	plan "github.com/crossplane-contrib/provider-jet-aws/internal/controller/backup/plan"
	regionsettings "github.com/crossplane-contrib/provider-jet-aws/internal/controller/backup/regionsettings"
	selection "github.com/crossplane-contrib/provider-jet-aws/internal/controller/backup/selection"
	vault "github.com/crossplane-contrib/provider-jet-aws/internal/controller/backup/vault"
	vaultnotifications "github.com/crossplane-contrib/provider-jet-aws/internal/controller/backup/vaultnotifications"
	vaultpolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/backup/vaultpolicy"
	computeenvironment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/batch/computeenvironment"
	jobdefinition "github.com/crossplane-contrib/provider-jet-aws/internal/controller/batch/jobdefinition"
	jobqueue "github.com/crossplane-contrib/provider-jet-aws/internal/controller/batch/jobqueue"
	budget "github.com/crossplane-contrib/provider-jet-aws/internal/controller/budgets/budget"
	budgetaction "github.com/crossplane-contrib/provider-jet-aws/internal/controller/budgets/budgetaction"
	voiceconnector "github.com/crossplane-contrib/provider-jet-aws/internal/controller/chime/voiceconnector"
	voiceconnectorgroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/chime/voiceconnectorgroup"
	environmentec2 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloud9/environmentec2"
	cloudformationtype "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudformation/cloudformationtype"
	stackcloudformation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudformation/stack"
	stackset "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudformation/stackset"
	stacksetinstance "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudformation/stacksetinstance"
	cachepolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudfront/cachepolicy"
	distribution "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudfront/distribution"
	functioncloudfront "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudfront/function"
	keygroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudfront/keygroup"
	monitoringsubscription "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudfront/monitoringsubscription"
	originaccessidentity "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudfront/originaccessidentity"
	originrequestpolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudfront/originrequestpolicy"
	publickey "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudfront/publickey"
	realtimelogconfig "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudfront/realtimelogconfig"
	cluster "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudhsmv2/cluster"
	hsm "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudhsmv2/hsm"
	trail "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudtrail/trail"
	compositealarm "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatch/compositealarm"
	dashboard "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatch/dashboard"
	metricalarm "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatch/metricalarm"
	metricstream "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatch/metricstream"
	definition "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatchlogs/definition"
	destination "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatchlogs/destination"
	destinationpolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatchlogs/destinationpolicy"
	group "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatchlogs/group"
	metricfilter "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatchlogs/metricfilter"
	resourcepolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatchlogs/resourcepolicy"
	stream "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatchlogs/stream"
	subscriptionfilter "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatchlogs/subscriptionfilter"
	domain "github.com/crossplane-contrib/provider-jet-aws/internal/controller/codeartifact/domain"
	domainpermissionspolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/codeartifact/domainpermissionspolicy"
	repository "github.com/crossplane-contrib/provider-jet-aws/internal/controller/codeartifact/repository"
	repositorypermissionspolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/codeartifact/repositorypermissionspolicy"
	project "github.com/crossplane-contrib/provider-jet-aws/internal/controller/codebuild/project"
	reportgroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/codebuild/reportgroup"
	sourcecredential "github.com/crossplane-contrib/provider-jet-aws/internal/controller/codebuild/sourcecredential"
	webhookcodebuild "github.com/crossplane-contrib/provider-jet-aws/internal/controller/codebuild/webhook"
	repositorycodecommit "github.com/crossplane-contrib/provider-jet-aws/internal/controller/codecommit/repository"
	trigger "github.com/crossplane-contrib/provider-jet-aws/internal/controller/codecommit/trigger"
	appcodedeploy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/codedeploy/app"
	deploymentconfig "github.com/crossplane-contrib/provider-jet-aws/internal/controller/codedeploy/deploymentconfig"
	deploymentgroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/codedeploy/deploymentgroup"
	codepipeline "github.com/crossplane-contrib/provider-jet-aws/internal/controller/codepipeline/codepipeline"
	webhookcodepipeline "github.com/crossplane-contrib/provider-jet-aws/internal/controller/codepipeline/webhook"
	connectioncodestarconnections "github.com/crossplane-contrib/provider-jet-aws/internal/controller/codestarconnections/connection"
	host "github.com/crossplane-contrib/provider-jet-aws/internal/controller/codestarconnections/host"
	notificationrule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/codestarnotifications/notificationrule"
	pool "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cognitoidentity/pool"
	poolrolesattachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cognitoidentity/poolrolesattachment"
	identityprovider "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cognitoidp/identityprovider"
	resourceserver "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cognitoidp/resourceserver"
	usergroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cognitoidp/usergroup"
	userpool "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cognitoidp/userpool"
	userpoolclient "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cognitoidp/userpoolclient"
	userpooldomain "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cognitoidp/userpooldomain"
	userpooluicustomization "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cognitoidp/userpooluicustomization"
	aggregateauthorization "github.com/crossplane-contrib/provider-jet-aws/internal/controller/configservice/aggregateauthorization"
	awsconfigurationrecorderstatus "github.com/crossplane-contrib/provider-jet-aws/internal/controller/configservice/awsconfigurationrecorderstatus"
	configrule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/configservice/configrule"
	configurationaggregator "github.com/crossplane-contrib/provider-jet-aws/internal/controller/configservice/configurationaggregator"
	configurationrecorder "github.com/crossplane-contrib/provider-jet-aws/internal/controller/configservice/configurationrecorder"
	conformancepack "github.com/crossplane-contrib/provider-jet-aws/internal/controller/configservice/conformancepack"
	deliverychannel "github.com/crossplane-contrib/provider-jet-aws/internal/controller/configservice/deliverychannel"
	organizationconformancepack "github.com/crossplane-contrib/provider-jet-aws/internal/controller/configservice/organizationconformancepack"
	organizationcustomrule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/configservice/organizationcustomrule"
	organizationmanagedrule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/configservice/organizationmanagedrule"
	remediationconfiguration "github.com/crossplane-contrib/provider-jet-aws/internal/controller/configservice/remediationconfiguration"
	reportdefinition "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cur/reportdefinition"
	pipeline "github.com/crossplane-contrib/provider-jet-aws/internal/controller/datapipeline/pipeline"
	agent "github.com/crossplane-contrib/provider-jet-aws/internal/controller/datasync/agent"
	locationefs "github.com/crossplane-contrib/provider-jet-aws/internal/controller/datasync/locationefs"
	locationfsxwindowsfilesystem "github.com/crossplane-contrib/provider-jet-aws/internal/controller/datasync/locationfsxwindowsfilesystem"
	locationnfs "github.com/crossplane-contrib/provider-jet-aws/internal/controller/datasync/locationnfs"
	locations3 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/datasync/locations3"
	locationsmb "github.com/crossplane-contrib/provider-jet-aws/internal/controller/datasync/locationsmb"
	task "github.com/crossplane-contrib/provider-jet-aws/internal/controller/datasync/task"
	clusterdax "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dax/cluster"
	parametergroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dax/parametergroup"
	subnetgroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dax/subnetgroup"
	projectdevicefarm "github.com/crossplane-contrib/provider-jet-aws/internal/controller/devicefarm/project"
	bgppeer "github.com/crossplane-contrib/provider-jet-aws/internal/controller/directconnect/bgppeer"
	connectiondirectconnect "github.com/crossplane-contrib/provider-jet-aws/internal/controller/directconnect/connection"
	connectionassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/directconnect/connectionassociation"
	gateway "github.com/crossplane-contrib/provider-jet-aws/internal/controller/directconnect/gateway"
	gatewayassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/directconnect/gatewayassociation"
	gatewayassociationproposal "github.com/crossplane-contrib/provider-jet-aws/internal/controller/directconnect/gatewayassociationproposal"
	hostedprivatevirtualinterface "github.com/crossplane-contrib/provider-jet-aws/internal/controller/directconnect/hostedprivatevirtualinterface"
	hostedprivatevirtualinterfaceaccepter "github.com/crossplane-contrib/provider-jet-aws/internal/controller/directconnect/hostedprivatevirtualinterfaceaccepter"
	hostedpublicvirtualinterface "github.com/crossplane-contrib/provider-jet-aws/internal/controller/directconnect/hostedpublicvirtualinterface"
	hostedpublicvirtualinterfaceaccepter "github.com/crossplane-contrib/provider-jet-aws/internal/controller/directconnect/hostedpublicvirtualinterfaceaccepter"
	hostedtransitvirtualinterface "github.com/crossplane-contrib/provider-jet-aws/internal/controller/directconnect/hostedtransitvirtualinterface"
	hostedtransitvirtualinterfaceaccepter "github.com/crossplane-contrib/provider-jet-aws/internal/controller/directconnect/hostedtransitvirtualinterfaceaccepter"
	lag "github.com/crossplane-contrib/provider-jet-aws/internal/controller/directconnect/lag"
	privatevirtualinterface "github.com/crossplane-contrib/provider-jet-aws/internal/controller/directconnect/privatevirtualinterface"
	publicvirtualinterface "github.com/crossplane-contrib/provider-jet-aws/internal/controller/directconnect/publicvirtualinterface"
	transitvirtualinterface "github.com/crossplane-contrib/provider-jet-aws/internal/controller/directconnect/transitvirtualinterface"
	lifecyclepolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dlm/lifecyclepolicy"
	certificatedms "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dms/certificate"
	endpoint "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dms/endpoint"
	eventsubscription "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dms/eventsubscription"
	replicationinstance "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dms/replicationinstance"
	replicationsubnetgroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dms/replicationsubnetgroup"
	replicationtask "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dms/replicationtask"
	clusterdocdb "github.com/crossplane-contrib/provider-jet-aws/internal/controller/docdb/cluster"
	clusterinstance "github.com/crossplane-contrib/provider-jet-aws/internal/controller/docdb/clusterinstance"
	clusterparametergroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/docdb/clusterparametergroup"
	clustersnapshot "github.com/crossplane-contrib/provider-jet-aws/internal/controller/docdb/clustersnapshot"
	subnetgroupdocdb "github.com/crossplane-contrib/provider-jet-aws/internal/controller/docdb/subnetgroup"
	conditionalforwarder "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ds/conditionalforwarder"
	directory "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ds/directory"
	logsubscription "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ds/logsubscription"
	globaltable "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dynamodb/globaltable"
	kinesisstreamingdestination "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dynamodb/kinesisstreamingdestination"
	table "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dynamodb/table"
	tableitem "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dynamodb/tableitem"
	tag "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dynamodb/tag"
	ami "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/ami"
	amicopy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/amicopy"
	amifrominstance "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/amifrominstance"
	amilaunchpermission "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/amilaunchpermission"
	availabilityzonegroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/availabilityzonegroup"
	capacityreservation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/capacityreservation"
	carriergateway "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/carriergateway"
	clientvpnauthorizationrule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/clientvpnauthorizationrule"
	clientvpnendpoint "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/clientvpnendpoint"
	clientvpnnetworkassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/clientvpnnetworkassociation"
	clientvpnroute "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/clientvpnroute"
	customergateway "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/customergateway"
	defaultnetworkacl "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/defaultnetworkacl"
	defaultroutetable "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/defaultroutetable"
	defaultsecuritygroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/defaultsecuritygroup"
	defaultsubnet "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/defaultsubnet"
	defaultvpc "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/defaultvpc"
	defaultvpcdhcpoptions "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/defaultvpcdhcpoptions"
	ebsdefaultkmskey "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/ebsdefaultkmskey"
	ebsencryptionbydefault "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/ebsencryptionbydefault"
	ebssnapshot "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/ebssnapshot"
	ebssnapshotcopy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/ebssnapshotcopy"
	ebssnapshotimport "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/ebssnapshotimport"
	ebsvolume "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/ebsvolume"
	egressonlyinternetgateway "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/egressonlyinternetgateway"
	eip "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/eip"
	eipassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/eipassociation"
	fleet "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/fleet"
	flowlog "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/flowlog"
	instance "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/instance"
	internetgateway "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/internetgateway"
	keypair "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/keypair"
	launchtemplate "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/launchtemplate"
	localgatewayroute "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/localgatewayroute"
	localgatewayroutetablevpcassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/localgatewayroutetablevpcassociation"
	mainroutetableassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/mainroutetableassociation"
	managedprefixlist "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/managedprefixlist"
	natgateway "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/natgateway"
	networkacl "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/networkacl"
	networkaclrule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/networkaclrule"
	networkinterface "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/networkinterface"
	networkinterfaceattachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/networkinterfaceattachment"
	networkinterfacesgattachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/networkinterfacesgattachment"
	placementgroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/placementgroup"
	routeec2 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/route"
	routetable "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/routetable"
	routetableassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/routetableassociation"
	securitygroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/securitygroup"
	securitygrouprule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/securitygrouprule"
	snapshotcreatevolumepermission "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/snapshotcreatevolumepermission"
	spotdatafeedsubscription "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/spotdatafeedsubscription"
	spotfleetrequest "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/spotfleetrequest"
	spotinstancerequest "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/spotinstancerequest"
	subnet "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/subnet"
	tagec2 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/tag"
	trafficmirrorfilter "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/trafficmirrorfilter"
	trafficmirrorfilterrule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/trafficmirrorfilterrule"
	trafficmirrorsession "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/trafficmirrorsession"
	trafficmirrortarget "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/trafficmirrortarget"
	transitgateway "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/transitgateway"
	transitgatewaypeeringattachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/transitgatewaypeeringattachment"
	transitgatewaypeeringattachmentaccepter "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/transitgatewaypeeringattachmentaccepter"
	transitgatewayprefixlistreference "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/transitgatewayprefixlistreference"
	transitgatewayroute "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/transitgatewayroute"
	transitgatewayroutetable "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/transitgatewayroutetable"
	transitgatewayroutetableassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/transitgatewayroutetableassociation"
	transitgatewayroutetablepropagation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/transitgatewayroutetablepropagation"
	transitgatewayvpcattachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/transitgatewayvpcattachment"
	transitgatewayvpcattachmentaccepter "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/transitgatewayvpcattachmentaccepter"
	volumeattachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/volumeattachment"
	vpc "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/vpc"
	vpcdhcpoptions "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/vpcdhcpoptions"
	vpcdhcpoptionsassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/vpcdhcpoptionsassociation"
	vpcendpoint "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/vpcendpoint"
	vpcendpointconnectionnotification "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/vpcendpointconnectionnotification"
	vpcendpointroutetableassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/vpcendpointroutetableassociation"
	vpcendpointservice "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/vpcendpointservice"
	vpcendpointserviceallowedprincipal "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/vpcendpointserviceallowedprincipal"
	vpcendpointsubnetassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/vpcendpointsubnetassociation"
	vpcipv4cidrblockassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/vpcipv4cidrblockassociation"
	vpcpeeringconnection "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/vpcpeeringconnection"
	vpcpeeringconnectionaccepter "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/vpcpeeringconnectionaccepter"
	vpcpeeringconnectionoptions "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/vpcpeeringconnectionoptions"
	vpnconnection "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/vpnconnection"
	vpnconnectionroute "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/vpnconnectionroute"
	vpngateway "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/vpngateway"
	vpngatewayattachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/vpngatewayattachment"
	vpngatewayroutepropagation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/vpngatewayroutepropagation"
	lifecyclepolicyecr "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ecr/lifecyclepolicy"
	registrypolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ecr/registrypolicy"
	replicationconfiguration "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ecr/replicationconfiguration"
	repositoryecr "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ecr/repository"
	repositorypolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ecr/repositorypolicy"
	repositoryecrpublic "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ecrpublic/repository"
	capacityprovider "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ecs/capacityprovider"
	clusterecs "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ecs/cluster"
	serviceecs "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ecs/service"
	taskdefinition "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ecs/taskdefinition"
	accesspoint "github.com/crossplane-contrib/provider-jet-aws/internal/controller/efs/accesspoint"
	backuppolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/efs/backuppolicy"
	filesystem "github.com/crossplane-contrib/provider-jet-aws/internal/controller/efs/filesystem"
	filesystempolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/efs/filesystempolicy"
	mounttarget "github.com/crossplane-contrib/provider-jet-aws/internal/controller/efs/mounttarget"
	addon "github.com/crossplane-contrib/provider-jet-aws/internal/controller/eks/addon"
	clustereks "github.com/crossplane-contrib/provider-jet-aws/internal/controller/eks/cluster"
	fargateprofile "github.com/crossplane-contrib/provider-jet-aws/internal/controller/eks/fargateprofile"
	identityproviderconfig "github.com/crossplane-contrib/provider-jet-aws/internal/controller/eks/identityproviderconfig"
	nodegroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/eks/nodegroup"
	clusterelasticache "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticache/cluster"
	globalreplicationgroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticache/globalreplicationgroup"
	parametergroupelasticache "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticache/parametergroup"
	replicationgroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticache/replicationgroup"
	securitygroupelasticache "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticache/securitygroup"
	subnetgroupelasticache "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticache/subnetgroup"
	user "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticache/user"
	usergroupelasticache "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticache/usergroup"
	applicationelasticbeanstalk "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticbeanstalk/application"
	applicationversion "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticbeanstalk/applicationversion"
	configurationtemplate "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticbeanstalk/configurationtemplate"
	environmentelasticbeanstalk "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticbeanstalk/environment"
	domainelasticsearch "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticsearch/domain"
	domainpolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticsearch/domainpolicy"
	domainsamloptions "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticsearch/domainsamloptions"
	pipelineelastictranscoder "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elastictranscoder/pipeline"
	preset "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elastictranscoder/preset"
	appcookiestickinesspolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elb/appcookiestickinesspolicy"
	attachmentelb "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elb/attachment"
	backendserverpolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elb/backendserverpolicy"
	elb "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elb/elb"
	lbcookiestickinesspolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elb/lbcookiestickinesspolicy"
	lbsslnegotiationpolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elb/lbsslnegotiationpolicy"
	listenerpolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elb/listenerpolicy"
	policyelb "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elb/policy"
	proxyprotocolpolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elb/proxyprotocolpolicy"
	alblistener "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elbv2/alblistener"
	alblistenercertificate "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elbv2/alblistenercertificate"
	alblistenerrule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elbv2/alblistenerrule"
	albtargetgroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elbv2/albtargetgroup"
	lb "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elbv2/lb"
	lblistener "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elbv2/lblistener"
	lblistenercertificate "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elbv2/lblistenercertificate"
	lblistenerrule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elbv2/lblistenerrule"
	lbtargetgroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elbv2/lbtargetgroup"
	lbtargetgroupattachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elbv2/lbtargetgroupattachment"
	clusteremr "github.com/crossplane-contrib/provider-jet-aws/internal/controller/emr/cluster"
	instancefleet "github.com/crossplane-contrib/provider-jet-aws/internal/controller/emr/instancefleet"
	instancegroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/emr/instancegroup"
	managedscalingpolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/emr/managedscalingpolicy"
	securityconfiguration "github.com/crossplane-contrib/provider-jet-aws/internal/controller/emr/securityconfiguration"
	apidestination "github.com/crossplane-contrib/provider-jet-aws/internal/controller/events/apidestination"
	archive "github.com/crossplane-contrib/provider-jet-aws/internal/controller/events/archive"
	bus "github.com/crossplane-contrib/provider-jet-aws/internal/controller/events/bus"
	buspolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/events/buspolicy"
	connectionevents "github.com/crossplane-contrib/provider-jet-aws/internal/controller/events/connection"
	permission "github.com/crossplane-contrib/provider-jet-aws/internal/controller/events/permission"
	rule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/events/rule"
	targetevents "github.com/crossplane-contrib/provider-jet-aws/internal/controller/events/target"
	deliverystream "github.com/crossplane-contrib/provider-jet-aws/internal/controller/firehose/deliverystream"
	adminaccount "github.com/crossplane-contrib/provider-jet-aws/internal/controller/fms/adminaccount"
	policyfms "github.com/crossplane-contrib/provider-jet-aws/internal/controller/fms/policy"
	backup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/fsx/backup"
	lustrefilesystem "github.com/crossplane-contrib/provider-jet-aws/internal/controller/fsx/lustrefilesystem"
	windowsfilesystem "github.com/crossplane-contrib/provider-jet-aws/internal/controller/fsx/windowsfilesystem"
	alias "github.com/crossplane-contrib/provider-jet-aws/internal/controller/gamelift/alias"
	build "github.com/crossplane-contrib/provider-jet-aws/internal/controller/gamelift/build"
	fleetgamelift "github.com/crossplane-contrib/provider-jet-aws/internal/controller/gamelift/fleet"
	gamesessionqueue "github.com/crossplane-contrib/provider-jet-aws/internal/controller/gamelift/gamesessionqueue"
	vaultglacier "github.com/crossplane-contrib/provider-jet-aws/internal/controller/glacier/vault"
	vaultlock "github.com/crossplane-contrib/provider-jet-aws/internal/controller/glacier/vaultlock"
	accelerator "github.com/crossplane-contrib/provider-jet-aws/internal/controller/globalaccelerator/accelerator"
	endpointgroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/globalaccelerator/endpointgroup"
	listener "github.com/crossplane-contrib/provider-jet-aws/internal/controller/globalaccelerator/listener"
	catalogdatabase "github.com/crossplane-contrib/provider-jet-aws/internal/controller/glue/catalogdatabase"
	catalogtable "github.com/crossplane-contrib/provider-jet-aws/internal/controller/glue/catalogtable"
	classifier "github.com/crossplane-contrib/provider-jet-aws/internal/controller/glue/classifier"
	crawler "github.com/crossplane-contrib/provider-jet-aws/internal/controller/glue/crawler"
	datacatalogencryptionsettings "github.com/crossplane-contrib/provider-jet-aws/internal/controller/glue/datacatalogencryptionsettings"
	devendpoint "github.com/crossplane-contrib/provider-jet-aws/internal/controller/glue/devendpoint"
	job "github.com/crossplane-contrib/provider-jet-aws/internal/controller/glue/job"
	mltransform "github.com/crossplane-contrib/provider-jet-aws/internal/controller/glue/mltransform"
	partition "github.com/crossplane-contrib/provider-jet-aws/internal/controller/glue/partition"
	registry "github.com/crossplane-contrib/provider-jet-aws/internal/controller/glue/registry"
	resourcepolicyglue "github.com/crossplane-contrib/provider-jet-aws/internal/controller/glue/resourcepolicy"
	schema "github.com/crossplane-contrib/provider-jet-aws/internal/controller/glue/schema"
	securityconfigurationglue "github.com/crossplane-contrib/provider-jet-aws/internal/controller/glue/securityconfiguration"
	triggerglue "github.com/crossplane-contrib/provider-jet-aws/internal/controller/glue/trigger"
	userdefinedfunction "github.com/crossplane-contrib/provider-jet-aws/internal/controller/glue/userdefinedfunction"
	workflow "github.com/crossplane-contrib/provider-jet-aws/internal/controller/glue/workflow"
	detector "github.com/crossplane-contrib/provider-jet-aws/internal/controller/guardduty/detector"
	filter "github.com/crossplane-contrib/provider-jet-aws/internal/controller/guardduty/filter"
	inviteaccepter "github.com/crossplane-contrib/provider-jet-aws/internal/controller/guardduty/inviteaccepter"
	ipset "github.com/crossplane-contrib/provider-jet-aws/internal/controller/guardduty/ipset"
	member "github.com/crossplane-contrib/provider-jet-aws/internal/controller/guardduty/member"
	organizationadminaccount "github.com/crossplane-contrib/provider-jet-aws/internal/controller/guardduty/organizationadminaccount"
	organizationconfiguration "github.com/crossplane-contrib/provider-jet-aws/internal/controller/guardduty/organizationconfiguration"
	publishingdestination "github.com/crossplane-contrib/provider-jet-aws/internal/controller/guardduty/publishingdestination"
	threatintelset "github.com/crossplane-contrib/provider-jet-aws/internal/controller/guardduty/threatintelset"
	accesskey "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/accesskey"
	accountalias "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/accountalias"
	accountpasswordpolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/accountpasswordpolicy"
	groupiam "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/group"
	groupmembership "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/groupmembership"
	grouppolicyattachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/grouppolicyattachment"
	instanceprofile "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/instanceprofile"
	openidconnectprovider "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/openidconnectprovider"
	policyiam "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/policy"
	role "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/role"
	rolepolicyattachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/rolepolicyattachment"
	samlprovider "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/samlprovider"
	servercertificate "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/servercertificate"
	servicelinkedrole "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/servicelinkedrole"
	useriam "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/user"
	usergroupmembership "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/usergroupmembership"
	userloginprofile "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/userloginprofile"
	userpolicyattachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/userpolicyattachment"
	usersshkey "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/usersshkey"
	component "github.com/crossplane-contrib/provider-jet-aws/internal/controller/imagebuilder/component"
	distributionconfiguration "github.com/crossplane-contrib/provider-jet-aws/internal/controller/imagebuilder/distributionconfiguration"
	image "github.com/crossplane-contrib/provider-jet-aws/internal/controller/imagebuilder/image"
	imagepipeline "github.com/crossplane-contrib/provider-jet-aws/internal/controller/imagebuilder/imagepipeline"
	imagerecipe "github.com/crossplane-contrib/provider-jet-aws/internal/controller/imagebuilder/imagerecipe"
	infrastructureconfiguration "github.com/crossplane-contrib/provider-jet-aws/internal/controller/imagebuilder/infrastructureconfiguration"
	assessmenttarget "github.com/crossplane-contrib/provider-jet-aws/internal/controller/inspector/assessmenttarget"
	assessmenttemplate "github.com/crossplane-contrib/provider-jet-aws/internal/controller/inspector/assessmenttemplate"
	resourcegroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/inspector/resourcegroup"
	certificateiot "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iot/certificate"
	policyiot "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iot/policy"
	policyattachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iot/policyattachment"
	rolealias "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iot/rolealias"
	thing "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iot/thing"
	thingprincipalattachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iot/thingprincipalattachment"
	thingtype "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iot/thingtype"
	topicrule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iot/topicrule"
	clusterkafka "github.com/crossplane-contrib/provider-jet-aws/internal/controller/kafka/cluster"
	configuration "github.com/crossplane-contrib/provider-jet-aws/internal/controller/kafka/configuration"
	scramsecretassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/kafka/scramsecretassociation"
	streamkinesis "github.com/crossplane-contrib/provider-jet-aws/internal/controller/kinesis/stream"
	streamconsumer "github.com/crossplane-contrib/provider-jet-aws/internal/controller/kinesis/streamconsumer"
	applicationkinesisanalytics "github.com/crossplane-contrib/provider-jet-aws/internal/controller/kinesisanalytics/application"
	applicationkinesisanalyticsv2 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/kinesisanalyticsv2/application"
	applicationsnapshot "github.com/crossplane-contrib/provider-jet-aws/internal/controller/kinesisanalyticsv2/applicationsnapshot"
	streamkinesisvideo "github.com/crossplane-contrib/provider-jet-aws/internal/controller/kinesisvideo/stream"
	aliaskms "github.com/crossplane-contrib/provider-jet-aws/internal/controller/kms/alias"
	ciphertext "github.com/crossplane-contrib/provider-jet-aws/internal/controller/kms/ciphertext"
	externalkey "github.com/crossplane-contrib/provider-jet-aws/internal/controller/kms/externalkey"
	grant "github.com/crossplane-contrib/provider-jet-aws/internal/controller/kms/grant"
	key "github.com/crossplane-contrib/provider-jet-aws/internal/controller/kms/key"
	datalakesettings "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lakeformation/datalakesettings"
	permissions "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lakeformation/permissions"
	resourcelakeformation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lakeformation/resource"
	aliaslambda "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lambda/alias"
	codesigningconfig "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lambda/codesigningconfig"
	eventsourcemapping "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lambda/eventsourcemapping"
	functionlambda "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lambda/function"
	functioneventinvokeconfig "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lambda/functioneventinvokeconfig"
	layerversion "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lambda/layerversion"
	permissionlambda "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lambda/permission"
	provisionedconcurrencyconfig "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lambda/provisionedconcurrencyconfig"
	bot "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lexmodels/bot"
	botalias "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lexmodels/botalias"
	intent "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lexmodels/intent"
	slottype "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lexmodels/slottype"
	association "github.com/crossplane-contrib/provider-jet-aws/internal/controller/licensemanager/association"
	licenseconfiguration "github.com/crossplane-contrib/provider-jet-aws/internal/controller/licensemanager/licenseconfiguration"
	domainlightsail "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lightsail/domain"
	instancelightsail "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lightsail/instance"
	instancepublicports "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lightsail/instancepublicports"
	keypairlightsail "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lightsail/keypair"
	staticip "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lightsail/staticip"
	staticipattachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lightsail/staticipattachment"
	memberaccountassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/macie/memberaccountassociation"
	s3bucketassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/macie/s3bucketassociation"
	accountmacie2 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/macie2/account"
	classificationjob "github.com/crossplane-contrib/provider-jet-aws/internal/controller/macie2/classificationjob"
	customdataidentifier "github.com/crossplane-contrib/provider-jet-aws/internal/controller/macie2/customdataidentifier"
	findingsfilter "github.com/crossplane-contrib/provider-jet-aws/internal/controller/macie2/findingsfilter"
	invitationaccepter "github.com/crossplane-contrib/provider-jet-aws/internal/controller/macie2/invitationaccepter"
	membermacie2 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/macie2/member"
	organizationadminaccountmacie2 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/macie2/organizationadminaccount"
	queue "github.com/crossplane-contrib/provider-jet-aws/internal/controller/mediaconvert/queue"
	channel "github.com/crossplane-contrib/provider-jet-aws/internal/controller/mediapackage/channel"
	container "github.com/crossplane-contrib/provider-jet-aws/internal/controller/mediastore/container"
	containerpolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/mediastore/containerpolicy"
	broker "github.com/crossplane-contrib/provider-jet-aws/internal/controller/mq/broker"
	configurationmq "github.com/crossplane-contrib/provider-jet-aws/internal/controller/mq/configuration"
	clusterneptune "github.com/crossplane-contrib/provider-jet-aws/internal/controller/neptune/cluster"
	clusterendpoint "github.com/crossplane-contrib/provider-jet-aws/internal/controller/neptune/clusterendpoint"
	clusterinstanceneptune "github.com/crossplane-contrib/provider-jet-aws/internal/controller/neptune/clusterinstance"
	clusterparametergroupneptune "github.com/crossplane-contrib/provider-jet-aws/internal/controller/neptune/clusterparametergroup"
	clustersnapshotneptune "github.com/crossplane-contrib/provider-jet-aws/internal/controller/neptune/clustersnapshot"
	eventsubscriptionneptune "github.com/crossplane-contrib/provider-jet-aws/internal/controller/neptune/eventsubscription"
	parametergroupneptune "github.com/crossplane-contrib/provider-jet-aws/internal/controller/neptune/parametergroup"
	subnetgroupneptune "github.com/crossplane-contrib/provider-jet-aws/internal/controller/neptune/subnetgroup"
	firewall "github.com/crossplane-contrib/provider-jet-aws/internal/controller/networkfirewall/firewall"
	firewallpolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/networkfirewall/firewallpolicy"
	loggingconfiguration "github.com/crossplane-contrib/provider-jet-aws/internal/controller/networkfirewall/loggingconfiguration"
	resourcepolicynetworkfirewall "github.com/crossplane-contrib/provider-jet-aws/internal/controller/networkfirewall/resourcepolicy"
	rulegroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/networkfirewall/rulegroup"
	applicationopsworks "github.com/crossplane-contrib/provider-jet-aws/internal/controller/opsworks/application"
	customlayer "github.com/crossplane-contrib/provider-jet-aws/internal/controller/opsworks/customlayer"
	ganglialayer "github.com/crossplane-contrib/provider-jet-aws/internal/controller/opsworks/ganglialayer"
	haproxylayer "github.com/crossplane-contrib/provider-jet-aws/internal/controller/opsworks/haproxylayer"
	instanceopsworks "github.com/crossplane-contrib/provider-jet-aws/internal/controller/opsworks/instance"
	javaapplayer "github.com/crossplane-contrib/provider-jet-aws/internal/controller/opsworks/javaapplayer"
	memcachedlayer "github.com/crossplane-contrib/provider-jet-aws/internal/controller/opsworks/memcachedlayer"
	mysqllayer "github.com/crossplane-contrib/provider-jet-aws/internal/controller/opsworks/mysqllayer"
	nodejsapplayer "github.com/crossplane-contrib/provider-jet-aws/internal/controller/opsworks/nodejsapplayer"
	permissionopsworks "github.com/crossplane-contrib/provider-jet-aws/internal/controller/opsworks/permission"
	phpapplayer "github.com/crossplane-contrib/provider-jet-aws/internal/controller/opsworks/phpapplayer"
	railsapplayer "github.com/crossplane-contrib/provider-jet-aws/internal/controller/opsworks/railsapplayer"
	rdsdbinstance "github.com/crossplane-contrib/provider-jet-aws/internal/controller/opsworks/rdsdbinstance"
	stackopsworks "github.com/crossplane-contrib/provider-jet-aws/internal/controller/opsworks/stack"
	staticweblayer "github.com/crossplane-contrib/provider-jet-aws/internal/controller/opsworks/staticweblayer"
	userprofile "github.com/crossplane-contrib/provider-jet-aws/internal/controller/opsworks/userprofile"
	accountorganizations "github.com/crossplane-contrib/provider-jet-aws/internal/controller/organizations/account"
	delegatedadministrator "github.com/crossplane-contrib/provider-jet-aws/internal/controller/organizations/delegatedadministrator"
	organization "github.com/crossplane-contrib/provider-jet-aws/internal/controller/organizations/organization"
	organizationalunit "github.com/crossplane-contrib/provider-jet-aws/internal/controller/organizations/organizationalunit"
	policyorganizations "github.com/crossplane-contrib/provider-jet-aws/internal/controller/organizations/policy"
	policyattachmentorganizations "github.com/crossplane-contrib/provider-jet-aws/internal/controller/organizations/policyattachment"
	admchannel "github.com/crossplane-contrib/provider-jet-aws/internal/controller/pinpoint/admchannel"
	apnschannel "github.com/crossplane-contrib/provider-jet-aws/internal/controller/pinpoint/apnschannel"
	apnssandboxchannel "github.com/crossplane-contrib/provider-jet-aws/internal/controller/pinpoint/apnssandboxchannel"
	apnsvoipchannel "github.com/crossplane-contrib/provider-jet-aws/internal/controller/pinpoint/apnsvoipchannel"
	apnsvoipsandboxchannel "github.com/crossplane-contrib/provider-jet-aws/internal/controller/pinpoint/apnsvoipsandboxchannel"
	apppinpoint "github.com/crossplane-contrib/provider-jet-aws/internal/controller/pinpoint/app"
	baiduchannel "github.com/crossplane-contrib/provider-jet-aws/internal/controller/pinpoint/baiduchannel"
	emailchannel "github.com/crossplane-contrib/provider-jet-aws/internal/controller/pinpoint/emailchannel"
	eventstream "github.com/crossplane-contrib/provider-jet-aws/internal/controller/pinpoint/eventstream"
	gcmchannel "github.com/crossplane-contrib/provider-jet-aws/internal/controller/pinpoint/gcmchannel"
	smschannel "github.com/crossplane-contrib/provider-jet-aws/internal/controller/pinpoint/smschannel"
	providerconfig "github.com/crossplane-contrib/provider-jet-aws/internal/controller/providerconfig"
	ledger "github.com/crossplane-contrib/provider-jet-aws/internal/controller/qldb/ledger"
	groupquicksight "github.com/crossplane-contrib/provider-jet-aws/internal/controller/quicksight/group"
	userquicksight "github.com/crossplane-contrib/provider-jet-aws/internal/controller/quicksight/user"
	principalassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ram/principalassociation"
	resourceassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ram/resourceassociation"
	resourceshare "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ram/resourceshare"
	resourceshareaccepter "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ram/resourceshareaccepter"
	clusterrds "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/cluster"
	clusterendpointrds "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/clusterendpoint"
	clusterinstancerds "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/clusterinstance"
	clusterparametergrouprds "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/clusterparametergroup"
	clusterroleassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/clusterroleassociation"
	clustersnapshotrds "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/clustersnapshot"
	eventsubscriptionrds "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/eventsubscription"
	globalcluster "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/globalcluster"
	instancerds "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/instance"
	instanceroleassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/instanceroleassociation"
	optiongroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/optiongroup"
	parametergrouprds "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/parametergroup"
	proxy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/proxy"
	proxydefaulttargetgroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/proxydefaulttargetgroup"
	proxyendpoint "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/proxyendpoint"
	proxytarget "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/proxytarget"
	securitygrouprds "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/securitygroup"
	snapshot "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/snapshot"
	subnetgrouprds "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/subnetgroup"
	clusterredshift "github.com/crossplane-contrib/provider-jet-aws/internal/controller/redshift/cluster"
	eventsubscriptionredshift "github.com/crossplane-contrib/provider-jet-aws/internal/controller/redshift/eventsubscription"
	parametergroupredshift "github.com/crossplane-contrib/provider-jet-aws/internal/controller/redshift/parametergroup"
	securitygroupredshift "github.com/crossplane-contrib/provider-jet-aws/internal/controller/redshift/securitygroup"
	snapshotcopygrant "github.com/crossplane-contrib/provider-jet-aws/internal/controller/redshift/snapshotcopygrant"
	snapshotschedule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/redshift/snapshotschedule"
	snapshotscheduleassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/redshift/snapshotscheduleassociation"
	subnetgroupredshift "github.com/crossplane-contrib/provider-jet-aws/internal/controller/redshift/subnetgroup"
	groupresourcegroups "github.com/crossplane-contrib/provider-jet-aws/internal/controller/resourcegroups/group"
	delegationset "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/delegationset"
	healthcheck "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/healthcheck"
	hostedzonednssec "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/hostedzonednssec"
	keysigningkey "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/keysigningkey"
	querylog "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/querylog"
	record "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/record"
	vpcassociationauthorization "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/vpcassociationauthorization"
	zone "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/zone"
	zoneassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/zoneassociation"
	clusterroute53recoverycontrolconfig "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53recoverycontrolconfig/cluster"
	controlpanel "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53recoverycontrolconfig/controlpanel"
	routingcontrol "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53recoverycontrolconfig/routingcontrol"
	safetyrule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53recoverycontrolconfig/safetyrule"
	cell "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53recoveryreadiness/cell"
	readinesscheck "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53recoveryreadiness/readinesscheck"
	recoverygroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53recoveryreadiness/recoverygroup"
	resourceset "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53recoveryreadiness/resourceset"
	dnssecconfig "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53resolver/dnssecconfig"
	endpointroute53resolver "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53resolver/endpoint"
	firewallconfig "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53resolver/firewallconfig"
	firewalldomainlist "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53resolver/firewalldomainlist"
	firewallrule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53resolver/firewallrule"
	firewallrulegroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53resolver/firewallrulegroup"
	firewallrulegroupassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53resolver/firewallrulegroupassociation"
	querylogconfig "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53resolver/querylogconfig"
	querylogconfigassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53resolver/querylogconfigassociation"
	ruleroute53resolver "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53resolver/rule"
	ruleassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53resolver/ruleassociation"
	bucket "github.com/crossplane-contrib/provider-jet-aws/internal/controller/s3/bucket"
	bucketanalyticsconfiguration "github.com/crossplane-contrib/provider-jet-aws/internal/controller/s3/bucketanalyticsconfiguration"
	bucketinventory "github.com/crossplane-contrib/provider-jet-aws/internal/controller/s3/bucketinventory"
	bucketmetric "github.com/crossplane-contrib/provider-jet-aws/internal/controller/s3/bucketmetric"
	bucketnotification "github.com/crossplane-contrib/provider-jet-aws/internal/controller/s3/bucketnotification"
	bucketobject "github.com/crossplane-contrib/provider-jet-aws/internal/controller/s3/bucketobject"
	bucketownershipcontrols "github.com/crossplane-contrib/provider-jet-aws/internal/controller/s3/bucketownershipcontrols"
	bucketpolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/s3/bucketpolicy"
	bucketpublicaccessblock "github.com/crossplane-contrib/provider-jet-aws/internal/controller/s3/bucketpublicaccessblock"
	objectcopy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/s3/objectcopy"
	accesspoints3control "github.com/crossplane-contrib/provider-jet-aws/internal/controller/s3control/accesspoint"
	accountpublicaccessblock "github.com/crossplane-contrib/provider-jet-aws/internal/controller/s3control/accountpublicaccessblock"
	buckets3control "github.com/crossplane-contrib/provider-jet-aws/internal/controller/s3control/bucket"
	bucketlifecycleconfiguration "github.com/crossplane-contrib/provider-jet-aws/internal/controller/s3control/bucketlifecycleconfiguration"
	bucketpolicys3control "github.com/crossplane-contrib/provider-jet-aws/internal/controller/s3control/bucketpolicy"
	endpoints3outposts "github.com/crossplane-contrib/provider-jet-aws/internal/controller/s3outposts/endpoint"
	appsagemaker "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sagemaker/app"
	appimageconfig "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sagemaker/appimageconfig"
	coderepository "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sagemaker/coderepository"
	devicefleet "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sagemaker/devicefleet"
	domainsagemaker "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sagemaker/domain"
	endpointsagemaker "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sagemaker/endpoint"
	endpointconfiguration "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sagemaker/endpointconfiguration"
	featuregroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sagemaker/featuregroup"
	humantaskui "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sagemaker/humantaskui"
	imagesagemaker "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sagemaker/image"
	imageversion "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sagemaker/imageversion"
	modelsagemaker "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sagemaker/model"
	modelpackagegroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sagemaker/modelpackagegroup"
	notebookinstance "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sagemaker/notebookinstance"
	notebookinstancelifecycleconfiguration "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sagemaker/notebookinstancelifecycleconfiguration"
	userprofilesagemaker "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sagemaker/userprofile"
	workforce "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sagemaker/workforce"
	workteam "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sagemaker/workteam"
	discoverer "github.com/crossplane-contrib/provider-jet-aws/internal/controller/schemas/discoverer"
	registryschemas "github.com/crossplane-contrib/provider-jet-aws/internal/controller/schemas/registry"
	schemaschemas "github.com/crossplane-contrib/provider-jet-aws/internal/controller/schemas/schema"
	secret "github.com/crossplane-contrib/provider-jet-aws/internal/controller/secretsmanager/secret"
	secretpolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/secretsmanager/secretpolicy"
	secretrotation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/secretsmanager/secretrotation"
	secretversion "github.com/crossplane-contrib/provider-jet-aws/internal/controller/secretsmanager/secretversion"
	accountsecurityhub "github.com/crossplane-contrib/provider-jet-aws/internal/controller/securityhub/account"
	actiontarget "github.com/crossplane-contrib/provider-jet-aws/internal/controller/securityhub/actiontarget"
	insight "github.com/crossplane-contrib/provider-jet-aws/internal/controller/securityhub/insight"
	inviteacceptersecurityhub "github.com/crossplane-contrib/provider-jet-aws/internal/controller/securityhub/inviteaccepter"
	membersecurityhub "github.com/crossplane-contrib/provider-jet-aws/internal/controller/securityhub/member"
	organizationadminaccountsecurityhub "github.com/crossplane-contrib/provider-jet-aws/internal/controller/securityhub/organizationadminaccount"
	organizationconfigurationsecurityhub "github.com/crossplane-contrib/provider-jet-aws/internal/controller/securityhub/organizationconfiguration"
	productsubscription "github.com/crossplane-contrib/provider-jet-aws/internal/controller/securityhub/productsubscription"
	standardscontrol "github.com/crossplane-contrib/provider-jet-aws/internal/controller/securityhub/standardscontrol"
	standardssubscription "github.com/crossplane-contrib/provider-jet-aws/internal/controller/securityhub/standardssubscription"
	cloudformationstack "github.com/crossplane-contrib/provider-jet-aws/internal/controller/serverlessrepo/cloudformationstack"
	budgetresourceassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/servicecatalog/budgetresourceassociation"
	constraint "github.com/crossplane-contrib/provider-jet-aws/internal/controller/servicecatalog/constraint"
	organizationsaccess "github.com/crossplane-contrib/provider-jet-aws/internal/controller/servicecatalog/organizationsaccess"
	portfolio "github.com/crossplane-contrib/provider-jet-aws/internal/controller/servicecatalog/portfolio"
	portfolioshare "github.com/crossplane-contrib/provider-jet-aws/internal/controller/servicecatalog/portfolioshare"
	principalportfolioassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/servicecatalog/principalportfolioassociation"
	product "github.com/crossplane-contrib/provider-jet-aws/internal/controller/servicecatalog/product"
	productportfolioassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/servicecatalog/productportfolioassociation"
	provisionedproduct "github.com/crossplane-contrib/provider-jet-aws/internal/controller/servicecatalog/provisionedproduct"
	provisioningartifact "github.com/crossplane-contrib/provider-jet-aws/internal/controller/servicecatalog/provisioningartifact"
	serviceaction "github.com/crossplane-contrib/provider-jet-aws/internal/controller/servicecatalog/serviceaction"
	tagoption "github.com/crossplane-contrib/provider-jet-aws/internal/controller/servicecatalog/tagoption"
	tagoptionresourceassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/servicecatalog/tagoptionresourceassociation"
	httpnamespace "github.com/crossplane-contrib/provider-jet-aws/internal/controller/servicediscovery/httpnamespace"
	privatednsnamespace "github.com/crossplane-contrib/provider-jet-aws/internal/controller/servicediscovery/privatednsnamespace"
	publicdnsnamespace "github.com/crossplane-contrib/provider-jet-aws/internal/controller/servicediscovery/publicdnsnamespace"
	serviceservicediscovery "github.com/crossplane-contrib/provider-jet-aws/internal/controller/servicediscovery/service"
	servicequota "github.com/crossplane-contrib/provider-jet-aws/internal/controller/servicequotas/servicequota"
	activereceiptruleset "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ses/activereceiptruleset"
	configurationset "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ses/configurationset"
	domaindkim "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ses/domaindkim"
	domainidentity "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ses/domainidentity"
	domainidentityverification "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ses/domainidentityverification"
	domainmailfrom "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ses/domainmailfrom"
	emailidentity "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ses/emailidentity"
	eventdestination "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ses/eventdestination"
	identitynotificationtopic "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ses/identitynotificationtopic"
	identitypolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ses/identitypolicy"
	receiptfilter "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ses/receiptfilter"
	receiptrule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ses/receiptrule"
	receiptruleset "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ses/receiptruleset"
	template "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ses/template"
	activity "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sfn/activity"
	statemachine "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sfn/statemachine"
	protection "github.com/crossplane-contrib/provider-jet-aws/internal/controller/shield/protection"
	protectiongroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/shield/protectiongroup"
	signingjob "github.com/crossplane-contrib/provider-jet-aws/internal/controller/signer/signingjob"
	signingprofile "github.com/crossplane-contrib/provider-jet-aws/internal/controller/signer/signingprofile"
	signingprofilepermission "github.com/crossplane-contrib/provider-jet-aws/internal/controller/signer/signingprofilepermission"
	domainsimpledb "github.com/crossplane-contrib/provider-jet-aws/internal/controller/simpledb/domain"
	platformapplication "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sns/platformapplication"
	smspreferences "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sns/smspreferences"
	topic "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sns/topic"
	topicpolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sns/topicpolicy"
	topicsubscription "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sns/topicsubscription"
	queuesqs "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sqs/queue"
	queuepolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sqs/queuepolicy"
	activation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ssm/activation"
	associationssm "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ssm/association"
	document "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ssm/document"
	maintenancewindow "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ssm/maintenancewindow"
	maintenancewindowtarget "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ssm/maintenancewindowtarget"
	maintenancewindowtask "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ssm/maintenancewindowtask"
	parameter "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ssm/parameter"
	patchbaseline "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ssm/patchbaseline"
	patchgroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ssm/patchgroup"
	resourcedatasync "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ssm/resourcedatasync"
	accountassignment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ssoadmin/accountassignment"
	managedpolicyattachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ssoadmin/managedpolicyattachment"
	permissionset "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ssoadmin/permissionset"
	permissionsetinlinepolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ssoadmin/permissionsetinlinepolicy"
	cache "github.com/crossplane-contrib/provider-jet-aws/internal/controller/storagegateway/cache"
	cachediscsivolume "github.com/crossplane-contrib/provider-jet-aws/internal/controller/storagegateway/cachediscsivolume"
	filesystemassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/storagegateway/filesystemassociation"
	gatewaystoragegateway "github.com/crossplane-contrib/provider-jet-aws/internal/controller/storagegateway/gateway"
	nfsfileshare "github.com/crossplane-contrib/provider-jet-aws/internal/controller/storagegateway/nfsfileshare"
	smbfileshare "github.com/crossplane-contrib/provider-jet-aws/internal/controller/storagegateway/smbfileshare"
	storediscsivolume "github.com/crossplane-contrib/provider-jet-aws/internal/controller/storagegateway/storediscsivolume"
	tapepool "github.com/crossplane-contrib/provider-jet-aws/internal/controller/storagegateway/tapepool"
	uploadbuffer "github.com/crossplane-contrib/provider-jet-aws/internal/controller/storagegateway/uploadbuffer"
	workingstorage "github.com/crossplane-contrib/provider-jet-aws/internal/controller/storagegateway/workingstorage"
	domainswf "github.com/crossplane-contrib/provider-jet-aws/internal/controller/swf/domain"
	canary "github.com/crossplane-contrib/provider-jet-aws/internal/controller/synthetics/canary"
	databasetimestreamwrite "github.com/crossplane-contrib/provider-jet-aws/internal/controller/timestreamwrite/database"
	tabletimestreamwrite "github.com/crossplane-contrib/provider-jet-aws/internal/controller/timestreamwrite/table"
	server "github.com/crossplane-contrib/provider-jet-aws/internal/controller/transfer/server"
	sshkey "github.com/crossplane-contrib/provider-jet-aws/internal/controller/transfer/sshkey"
	usertransfer "github.com/crossplane-contrib/provider-jet-aws/internal/controller/transfer/user"
	bytematchset "github.com/crossplane-contrib/provider-jet-aws/internal/controller/waf/bytematchset"
	geomatchset "github.com/crossplane-contrib/provider-jet-aws/internal/controller/waf/geomatchset"
	ipsetwaf "github.com/crossplane-contrib/provider-jet-aws/internal/controller/waf/ipset"
	ratebasedrule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/waf/ratebasedrule"
	regexmatchset "github.com/crossplane-contrib/provider-jet-aws/internal/controller/waf/regexmatchset"
	regexpatternset "github.com/crossplane-contrib/provider-jet-aws/internal/controller/waf/regexpatternset"
	rulewaf "github.com/crossplane-contrib/provider-jet-aws/internal/controller/waf/rule"
	sizeconstraintset "github.com/crossplane-contrib/provider-jet-aws/internal/controller/waf/sizeconstraintset"
	sqlinjectionmatchset "github.com/crossplane-contrib/provider-jet-aws/internal/controller/waf/sqlinjectionmatchset"
	webacl "github.com/crossplane-contrib/provider-jet-aws/internal/controller/waf/webacl"
	xssmatchset "github.com/crossplane-contrib/provider-jet-aws/internal/controller/waf/xssmatchset"
	bytematchsetwafregional "github.com/crossplane-contrib/provider-jet-aws/internal/controller/wafregional/bytematchset"
	geomatchsetwafregional "github.com/crossplane-contrib/provider-jet-aws/internal/controller/wafregional/geomatchset"
	ipsetwafregional "github.com/crossplane-contrib/provider-jet-aws/internal/controller/wafregional/ipset"
	ratebasedrulewafregional "github.com/crossplane-contrib/provider-jet-aws/internal/controller/wafregional/ratebasedrule"
	regexmatchsetwafregional "github.com/crossplane-contrib/provider-jet-aws/internal/controller/wafregional/regexmatchset"
	regexpatternsetwafregional "github.com/crossplane-contrib/provider-jet-aws/internal/controller/wafregional/regexpatternset"
	rulewafregional "github.com/crossplane-contrib/provider-jet-aws/internal/controller/wafregional/rule"
	sizeconstraintsetwafregional "github.com/crossplane-contrib/provider-jet-aws/internal/controller/wafregional/sizeconstraintset"
	sqlinjectionmatchsetwafregional "github.com/crossplane-contrib/provider-jet-aws/internal/controller/wafregional/sqlinjectionmatchset"
	webaclwafregional "github.com/crossplane-contrib/provider-jet-aws/internal/controller/wafregional/webacl"
	webaclassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/wafregional/webaclassociation"
	xssmatchsetwafregional "github.com/crossplane-contrib/provider-jet-aws/internal/controller/wafregional/xssmatchset"
	ipsetwafv2 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/wafv2/ipset"
	regexpatternsetwafv2 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/wafv2/regexpatternset"
	rulegroupwafv2 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/wafv2/rulegroup"
	webaclwafv2 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/wafv2/webacl"
	webaclassociationwafv2 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/wafv2/webaclassociation"
	webaclloggingconfiguration "github.com/crossplane-contrib/provider-jet-aws/internal/controller/wafv2/webaclloggingconfiguration"
	fleetworklink "github.com/crossplane-contrib/provider-jet-aws/internal/controller/worklink/fleet"
	websitecertificateauthorityassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/worklink/websitecertificateauthorityassociation"
	directoryworkspaces "github.com/crossplane-contrib/provider-jet-aws/internal/controller/workspaces/directory"
	ipgroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/workspaces/ipgroup"
	workspaceworkspaces "github.com/crossplane-contrib/provider-jet-aws/internal/controller/workspaces/workspace"
	encryptionconfig "github.com/crossplane-contrib/provider-jet-aws/internal/controller/xray/encryptionconfig"
	groupxray "github.com/crossplane-contrib/provider-jet-aws/internal/controller/xray/group"
	samplingrule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/xray/samplingrule"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		analyzer.Setup,
		certificate.Setup,
		certificatevalidation.Setup,
		certificateacmpca.Setup,
		certificateauthority.Setup,
		certificateauthoritycertificate.Setup,
		workspace.Setup,
		app.Setup,
		backendenvironment.Setup,
		branch.Setup,
		domainassociation.Setup,
		webhook.Setup,
		account.Setup,
		apikey.Setup,
		authorizer.Setup,
		basepathmapping.Setup,
		clientcertificate.Setup,
		deployment.Setup,
		documentationpart.Setup,
		documentationversion.Setup,
		domainname.Setup,
		gatewayresponse.Setup,
		integration.Setup,
		integrationresponse.Setup,
		method.Setup,
		methodresponse.Setup,
		methodsettings.Setup,
		model.Setup,
		requestvalidator.Setup,
		resource.Setup,
		restapi.Setup,
		restapipolicy.Setup,
		stage.Setup,
		usageplan.Setup,
		usageplankey.Setup,
		vpclink.Setup,
		api.Setup,
		apimapping.Setup,
		authorizerapigatewayv2.Setup,
		deploymentapigatewayv2.Setup,
		domainnameapigatewayv2.Setup,
		integrationapigatewayv2.Setup,
		integrationresponseapigatewayv2.Setup,
		modelapigatewayv2.Setup,
		route.Setup,
		routeresponse.Setup,
		stageapigatewayv2.Setup,
		vpclinkapigatewayv2.Setup,
		policy.Setup,
		scheduledaction.Setup,
		target.Setup,
		application.Setup,
		configurationprofile.Setup,
		deploymentappconfig.Setup,
		deploymentstrategy.Setup,
		environment.Setup,
		hostedconfigurationversion.Setup,
		gatewayroute.Setup,
		mesh.Setup,
		routeappmesh.Setup,
		virtualgateway.Setup,
		virtualnode.Setup,
		virtualrouter.Setup,
		virtualservice.Setup,
		autoscalingconfigurationversion.Setup,
		connection.Setup,
		customdomainassociation.Setup,
		service.Setup,
		stack.Setup,
		apikeyappsync.Setup,
		datasource.Setup,
		function.Setup,
		graphqlapi.Setup,
		resolver.Setup,
		database.Setup,
		namedquery.Setup,
		workgroup.Setup,
		attachment.Setup,
		autoscalinggroup.Setup,
		grouptag.Setup,
		launchconfiguration.Setup,
		lifecyclehook.Setup,
		notification.Setup,
		policyautoscaling.Setup,
		schedule.Setup,
		scalingplan.Setup,
		globalsettings.Setup,
		plan.Setup,
		regionsettings.Setup,
		selection.Setup,
		vault.Setup,
		vaultnotifications.Setup,
		vaultpolicy.Setup,
		computeenvironment.Setup,
		jobdefinition.Setup,
		jobqueue.Setup,
		budget.Setup,
		budgetaction.Setup,
		voiceconnector.Setup,
		voiceconnectorgroup.Setup,
		environmentec2.Setup,
		cloudformationtype.Setup,
		stackcloudformation.Setup,
		stackset.Setup,
		stacksetinstance.Setup,
		cachepolicy.Setup,
		distribution.Setup,
		functioncloudfront.Setup,
		keygroup.Setup,
		monitoringsubscription.Setup,
		originaccessidentity.Setup,
		originrequestpolicy.Setup,
		publickey.Setup,
		realtimelogconfig.Setup,
		cluster.Setup,
		hsm.Setup,
		trail.Setup,
		compositealarm.Setup,
		dashboard.Setup,
		metricalarm.Setup,
		metricstream.Setup,
		definition.Setup,
		destination.Setup,
		destinationpolicy.Setup,
		group.Setup,
		metricfilter.Setup,
		resourcepolicy.Setup,
		stream.Setup,
		subscriptionfilter.Setup,
		domain.Setup,
		domainpermissionspolicy.Setup,
		repository.Setup,
		repositorypermissionspolicy.Setup,
		project.Setup,
		reportgroup.Setup,
		sourcecredential.Setup,
		webhookcodebuild.Setup,
		repositorycodecommit.Setup,
		trigger.Setup,
		appcodedeploy.Setup,
		deploymentconfig.Setup,
		deploymentgroup.Setup,
		codepipeline.Setup,
		webhookcodepipeline.Setup,
		connectioncodestarconnections.Setup,
		host.Setup,
		notificationrule.Setup,
		pool.Setup,
		poolrolesattachment.Setup,
		identityprovider.Setup,
		resourceserver.Setup,
		usergroup.Setup,
		userpool.Setup,
		userpoolclient.Setup,
		userpooldomain.Setup,
		userpooluicustomization.Setup,
		aggregateauthorization.Setup,
		awsconfigurationrecorderstatus.Setup,
		configrule.Setup,
		configurationaggregator.Setup,
		configurationrecorder.Setup,
		conformancepack.Setup,
		deliverychannel.Setup,
		organizationconformancepack.Setup,
		organizationcustomrule.Setup,
		organizationmanagedrule.Setup,
		remediationconfiguration.Setup,
		reportdefinition.Setup,
		pipeline.Setup,
		agent.Setup,
		locationefs.Setup,
		locationfsxwindowsfilesystem.Setup,
		locationnfs.Setup,
		locations3.Setup,
		locationsmb.Setup,
		task.Setup,
		clusterdax.Setup,
		parametergroup.Setup,
		subnetgroup.Setup,
		projectdevicefarm.Setup,
		bgppeer.Setup,
		connectiondirectconnect.Setup,
		connectionassociation.Setup,
		gateway.Setup,
		gatewayassociation.Setup,
		gatewayassociationproposal.Setup,
		hostedprivatevirtualinterface.Setup,
		hostedprivatevirtualinterfaceaccepter.Setup,
		hostedpublicvirtualinterface.Setup,
		hostedpublicvirtualinterfaceaccepter.Setup,
		hostedtransitvirtualinterface.Setup,
		hostedtransitvirtualinterfaceaccepter.Setup,
		lag.Setup,
		privatevirtualinterface.Setup,
		publicvirtualinterface.Setup,
		transitvirtualinterface.Setup,
		lifecyclepolicy.Setup,
		certificatedms.Setup,
		endpoint.Setup,
		eventsubscription.Setup,
		replicationinstance.Setup,
		replicationsubnetgroup.Setup,
		replicationtask.Setup,
		clusterdocdb.Setup,
		clusterinstance.Setup,
		clusterparametergroup.Setup,
		clustersnapshot.Setup,
		subnetgroupdocdb.Setup,
		conditionalforwarder.Setup,
		directory.Setup,
		logsubscription.Setup,
		globaltable.Setup,
		kinesisstreamingdestination.Setup,
		table.Setup,
		tableitem.Setup,
		tag.Setup,
		ami.Setup,
		amicopy.Setup,
		amifrominstance.Setup,
		amilaunchpermission.Setup,
		availabilityzonegroup.Setup,
		capacityreservation.Setup,
		carriergateway.Setup,
		clientvpnauthorizationrule.Setup,
		clientvpnendpoint.Setup,
		clientvpnnetworkassociation.Setup,
		clientvpnroute.Setup,
		customergateway.Setup,
		defaultnetworkacl.Setup,
		defaultroutetable.Setup,
		defaultsecuritygroup.Setup,
		defaultsubnet.Setup,
		defaultvpc.Setup,
		defaultvpcdhcpoptions.Setup,
		ebsdefaultkmskey.Setup,
		ebsencryptionbydefault.Setup,
		ebssnapshot.Setup,
		ebssnapshotcopy.Setup,
		ebssnapshotimport.Setup,
		ebsvolume.Setup,
		egressonlyinternetgateway.Setup,
		eip.Setup,
		eipassociation.Setup,
		fleet.Setup,
		flowlog.Setup,
		instance.Setup,
		internetgateway.Setup,
		keypair.Setup,
		launchtemplate.Setup,
		localgatewayroute.Setup,
		localgatewayroutetablevpcassociation.Setup,
		mainroutetableassociation.Setup,
		managedprefixlist.Setup,
		natgateway.Setup,
		networkacl.Setup,
		networkaclrule.Setup,
		networkinterface.Setup,
		networkinterfaceattachment.Setup,
		networkinterfacesgattachment.Setup,
		placementgroup.Setup,
		routeec2.Setup,
		routetable.Setup,
		routetableassociation.Setup,
		securitygroup.Setup,
		securitygrouprule.Setup,
		snapshotcreatevolumepermission.Setup,
		spotdatafeedsubscription.Setup,
		spotfleetrequest.Setup,
		spotinstancerequest.Setup,
		subnet.Setup,
		tagec2.Setup,
		trafficmirrorfilter.Setup,
		trafficmirrorfilterrule.Setup,
		trafficmirrorsession.Setup,
		trafficmirrortarget.Setup,
		transitgateway.Setup,
		transitgatewaypeeringattachment.Setup,
		transitgatewaypeeringattachmentaccepter.Setup,
		transitgatewayprefixlistreference.Setup,
		transitgatewayroute.Setup,
		transitgatewayroutetable.Setup,
		transitgatewayroutetableassociation.Setup,
		transitgatewayroutetablepropagation.Setup,
		transitgatewayvpcattachment.Setup,
		transitgatewayvpcattachmentaccepter.Setup,
		volumeattachment.Setup,
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
		vpcpeeringconnectionaccepter.Setup,
		vpcpeeringconnectionoptions.Setup,
		vpnconnection.Setup,
		vpnconnectionroute.Setup,
		vpngateway.Setup,
		vpngatewayattachment.Setup,
		vpngatewayroutepropagation.Setup,
		lifecyclepolicyecr.Setup,
		registrypolicy.Setup,
		replicationconfiguration.Setup,
		repositoryecr.Setup,
		repositorypolicy.Setup,
		repositoryecrpublic.Setup,
		capacityprovider.Setup,
		clusterecs.Setup,
		serviceecs.Setup,
		taskdefinition.Setup,
		accesspoint.Setup,
		backuppolicy.Setup,
		filesystem.Setup,
		filesystempolicy.Setup,
		mounttarget.Setup,
		addon.Setup,
		clustereks.Setup,
		fargateprofile.Setup,
		identityproviderconfig.Setup,
		nodegroup.Setup,
		clusterelasticache.Setup,
		globalreplicationgroup.Setup,
		parametergroupelasticache.Setup,
		replicationgroup.Setup,
		securitygroupelasticache.Setup,
		subnetgroupelasticache.Setup,
		user.Setup,
		usergroupelasticache.Setup,
		applicationelasticbeanstalk.Setup,
		applicationversion.Setup,
		configurationtemplate.Setup,
		environmentelasticbeanstalk.Setup,
		domainelasticsearch.Setup,
		domainpolicy.Setup,
		domainsamloptions.Setup,
		pipelineelastictranscoder.Setup,
		preset.Setup,
		appcookiestickinesspolicy.Setup,
		attachmentelb.Setup,
		backendserverpolicy.Setup,
		elb.Setup,
		lbcookiestickinesspolicy.Setup,
		lbsslnegotiationpolicy.Setup,
		listenerpolicy.Setup,
		policyelb.Setup,
		proxyprotocolpolicy.Setup,
		alblistener.Setup,
		alblistenercertificate.Setup,
		alblistenerrule.Setup,
		albtargetgroup.Setup,
		lb.Setup,
		lblistener.Setup,
		lblistenercertificate.Setup,
		lblistenerrule.Setup,
		lbtargetgroup.Setup,
		lbtargetgroupattachment.Setup,
		clusteremr.Setup,
		instancefleet.Setup,
		instancegroup.Setup,
		managedscalingpolicy.Setup,
		securityconfiguration.Setup,
		apidestination.Setup,
		archive.Setup,
		bus.Setup,
		buspolicy.Setup,
		connectionevents.Setup,
		permission.Setup,
		rule.Setup,
		targetevents.Setup,
		deliverystream.Setup,
		adminaccount.Setup,
		policyfms.Setup,
		backup.Setup,
		lustrefilesystem.Setup,
		windowsfilesystem.Setup,
		alias.Setup,
		build.Setup,
		fleetgamelift.Setup,
		gamesessionqueue.Setup,
		vaultglacier.Setup,
		vaultlock.Setup,
		accelerator.Setup,
		endpointgroup.Setup,
		listener.Setup,
		catalogdatabase.Setup,
		catalogtable.Setup,
		classifier.Setup,
		crawler.Setup,
		datacatalogencryptionsettings.Setup,
		devendpoint.Setup,
		job.Setup,
		mltransform.Setup,
		partition.Setup,
		registry.Setup,
		resourcepolicyglue.Setup,
		schema.Setup,
		securityconfigurationglue.Setup,
		triggerglue.Setup,
		userdefinedfunction.Setup,
		workflow.Setup,
		detector.Setup,
		filter.Setup,
		inviteaccepter.Setup,
		ipset.Setup,
		member.Setup,
		organizationadminaccount.Setup,
		organizationconfiguration.Setup,
		publishingdestination.Setup,
		threatintelset.Setup,
		accesskey.Setup,
		accountalias.Setup,
		accountpasswordpolicy.Setup,
		groupiam.Setup,
		groupmembership.Setup,
		grouppolicyattachment.Setup,
		instanceprofile.Setup,
		openidconnectprovider.Setup,
		policyiam.Setup,
		role.Setup,
		rolepolicyattachment.Setup,
		samlprovider.Setup,
		servercertificate.Setup,
		servicelinkedrole.Setup,
		useriam.Setup,
		usergroupmembership.Setup,
		userloginprofile.Setup,
		userpolicyattachment.Setup,
		usersshkey.Setup,
		component.Setup,
		distributionconfiguration.Setup,
		image.Setup,
		imagepipeline.Setup,
		imagerecipe.Setup,
		infrastructureconfiguration.Setup,
		assessmenttarget.Setup,
		assessmenttemplate.Setup,
		resourcegroup.Setup,
		certificateiot.Setup,
		policyiot.Setup,
		policyattachment.Setup,
		rolealias.Setup,
		thing.Setup,
		thingprincipalattachment.Setup,
		thingtype.Setup,
		topicrule.Setup,
		clusterkafka.Setup,
		configuration.Setup,
		scramsecretassociation.Setup,
		streamkinesis.Setup,
		streamconsumer.Setup,
		applicationkinesisanalytics.Setup,
		applicationkinesisanalyticsv2.Setup,
		applicationsnapshot.Setup,
		streamkinesisvideo.Setup,
		aliaskms.Setup,
		ciphertext.Setup,
		externalkey.Setup,
		grant.Setup,
		key.Setup,
		datalakesettings.Setup,
		permissions.Setup,
		resourcelakeformation.Setup,
		aliaslambda.Setup,
		codesigningconfig.Setup,
		eventsourcemapping.Setup,
		functionlambda.Setup,
		functioneventinvokeconfig.Setup,
		layerversion.Setup,
		permissionlambda.Setup,
		provisionedconcurrencyconfig.Setup,
		bot.Setup,
		botalias.Setup,
		intent.Setup,
		slottype.Setup,
		association.Setup,
		licenseconfiguration.Setup,
		domainlightsail.Setup,
		instancelightsail.Setup,
		instancepublicports.Setup,
		keypairlightsail.Setup,
		staticip.Setup,
		staticipattachment.Setup,
		memberaccountassociation.Setup,
		s3bucketassociation.Setup,
		accountmacie2.Setup,
		classificationjob.Setup,
		customdataidentifier.Setup,
		findingsfilter.Setup,
		invitationaccepter.Setup,
		membermacie2.Setup,
		organizationadminaccountmacie2.Setup,
		queue.Setup,
		channel.Setup,
		container.Setup,
		containerpolicy.Setup,
		broker.Setup,
		configurationmq.Setup,
		clusterneptune.Setup,
		clusterendpoint.Setup,
		clusterinstanceneptune.Setup,
		clusterparametergroupneptune.Setup,
		clustersnapshotneptune.Setup,
		eventsubscriptionneptune.Setup,
		parametergroupneptune.Setup,
		subnetgroupneptune.Setup,
		firewall.Setup,
		firewallpolicy.Setup,
		loggingconfiguration.Setup,
		resourcepolicynetworkfirewall.Setup,
		rulegroup.Setup,
		applicationopsworks.Setup,
		customlayer.Setup,
		ganglialayer.Setup,
		haproxylayer.Setup,
		instanceopsworks.Setup,
		javaapplayer.Setup,
		memcachedlayer.Setup,
		mysqllayer.Setup,
		nodejsapplayer.Setup,
		permissionopsworks.Setup,
		phpapplayer.Setup,
		railsapplayer.Setup,
		rdsdbinstance.Setup,
		stackopsworks.Setup,
		staticweblayer.Setup,
		userprofile.Setup,
		accountorganizations.Setup,
		delegatedadministrator.Setup,
		organization.Setup,
		organizationalunit.Setup,
		policyorganizations.Setup,
		policyattachmentorganizations.Setup,
		admchannel.Setup,
		apnschannel.Setup,
		apnssandboxchannel.Setup,
		apnsvoipchannel.Setup,
		apnsvoipsandboxchannel.Setup,
		apppinpoint.Setup,
		baiduchannel.Setup,
		emailchannel.Setup,
		eventstream.Setup,
		gcmchannel.Setup,
		smschannel.Setup,
		providerconfig.Setup,
		ledger.Setup,
		groupquicksight.Setup,
		userquicksight.Setup,
		principalassociation.Setup,
		resourceassociation.Setup,
		resourceshare.Setup,
		resourceshareaccepter.Setup,
		clusterrds.Setup,
		clusterendpointrds.Setup,
		clusterinstancerds.Setup,
		clusterparametergrouprds.Setup,
		clusterroleassociation.Setup,
		clustersnapshotrds.Setup,
		eventsubscriptionrds.Setup,
		globalcluster.Setup,
		instancerds.Setup,
		instanceroleassociation.Setup,
		optiongroup.Setup,
		parametergrouprds.Setup,
		proxy.Setup,
		proxydefaulttargetgroup.Setup,
		proxyendpoint.Setup,
		proxytarget.Setup,
		securitygrouprds.Setup,
		snapshot.Setup,
		subnetgrouprds.Setup,
		clusterredshift.Setup,
		eventsubscriptionredshift.Setup,
		parametergroupredshift.Setup,
		securitygroupredshift.Setup,
		snapshotcopygrant.Setup,
		snapshotschedule.Setup,
		snapshotscheduleassociation.Setup,
		subnetgroupredshift.Setup,
		groupresourcegroups.Setup,
		delegationset.Setup,
		healthcheck.Setup,
		hostedzonednssec.Setup,
		keysigningkey.Setup,
		querylog.Setup,
		record.Setup,
		vpcassociationauthorization.Setup,
		zone.Setup,
		zoneassociation.Setup,
		clusterroute53recoverycontrolconfig.Setup,
		controlpanel.Setup,
		routingcontrol.Setup,
		safetyrule.Setup,
		cell.Setup,
		readinesscheck.Setup,
		recoverygroup.Setup,
		resourceset.Setup,
		dnssecconfig.Setup,
		endpointroute53resolver.Setup,
		firewallconfig.Setup,
		firewalldomainlist.Setup,
		firewallrule.Setup,
		firewallrulegroup.Setup,
		firewallrulegroupassociation.Setup,
		querylogconfig.Setup,
		querylogconfigassociation.Setup,
		ruleroute53resolver.Setup,
		ruleassociation.Setup,
		bucket.Setup,
		bucketanalyticsconfiguration.Setup,
		bucketinventory.Setup,
		bucketmetric.Setup,
		bucketnotification.Setup,
		bucketobject.Setup,
		bucketownershipcontrols.Setup,
		bucketpolicy.Setup,
		bucketpublicaccessblock.Setup,
		objectcopy.Setup,
		accesspoints3control.Setup,
		accountpublicaccessblock.Setup,
		buckets3control.Setup,
		bucketlifecycleconfiguration.Setup,
		bucketpolicys3control.Setup,
		endpoints3outposts.Setup,
		appsagemaker.Setup,
		appimageconfig.Setup,
		coderepository.Setup,
		devicefleet.Setup,
		domainsagemaker.Setup,
		endpointsagemaker.Setup,
		endpointconfiguration.Setup,
		featuregroup.Setup,
		humantaskui.Setup,
		imagesagemaker.Setup,
		imageversion.Setup,
		modelsagemaker.Setup,
		modelpackagegroup.Setup,
		notebookinstance.Setup,
		notebookinstancelifecycleconfiguration.Setup,
		userprofilesagemaker.Setup,
		workforce.Setup,
		workteam.Setup,
		discoverer.Setup,
		registryschemas.Setup,
		schemaschemas.Setup,
		secret.Setup,
		secretpolicy.Setup,
		secretrotation.Setup,
		secretversion.Setup,
		accountsecurityhub.Setup,
		actiontarget.Setup,
		insight.Setup,
		inviteacceptersecurityhub.Setup,
		membersecurityhub.Setup,
		organizationadminaccountsecurityhub.Setup,
		organizationconfigurationsecurityhub.Setup,
		productsubscription.Setup,
		standardscontrol.Setup,
		standardssubscription.Setup,
		cloudformationstack.Setup,
		budgetresourceassociation.Setup,
		constraint.Setup,
		organizationsaccess.Setup,
		portfolio.Setup,
		portfolioshare.Setup,
		principalportfolioassociation.Setup,
		product.Setup,
		productportfolioassociation.Setup,
		provisionedproduct.Setup,
		provisioningartifact.Setup,
		serviceaction.Setup,
		tagoption.Setup,
		tagoptionresourceassociation.Setup,
		httpnamespace.Setup,
		privatednsnamespace.Setup,
		publicdnsnamespace.Setup,
		serviceservicediscovery.Setup,
		servicequota.Setup,
		activereceiptruleset.Setup,
		configurationset.Setup,
		domaindkim.Setup,
		domainidentity.Setup,
		domainidentityverification.Setup,
		domainmailfrom.Setup,
		emailidentity.Setup,
		eventdestination.Setup,
		identitynotificationtopic.Setup,
		identitypolicy.Setup,
		receiptfilter.Setup,
		receiptrule.Setup,
		receiptruleset.Setup,
		template.Setup,
		activity.Setup,
		statemachine.Setup,
		protection.Setup,
		protectiongroup.Setup,
		signingjob.Setup,
		signingprofile.Setup,
		signingprofilepermission.Setup,
		domainsimpledb.Setup,
		platformapplication.Setup,
		smspreferences.Setup,
		topic.Setup,
		topicpolicy.Setup,
		topicsubscription.Setup,
		queuesqs.Setup,
		queuepolicy.Setup,
		activation.Setup,
		associationssm.Setup,
		document.Setup,
		maintenancewindow.Setup,
		maintenancewindowtarget.Setup,
		maintenancewindowtask.Setup,
		parameter.Setup,
		patchbaseline.Setup,
		patchgroup.Setup,
		resourcedatasync.Setup,
		accountassignment.Setup,
		managedpolicyattachment.Setup,
		permissionset.Setup,
		permissionsetinlinepolicy.Setup,
		cache.Setup,
		cachediscsivolume.Setup,
		filesystemassociation.Setup,
		gatewaystoragegateway.Setup,
		nfsfileshare.Setup,
		smbfileshare.Setup,
		storediscsivolume.Setup,
		tapepool.Setup,
		uploadbuffer.Setup,
		workingstorage.Setup,
		domainswf.Setup,
		canary.Setup,
		databasetimestreamwrite.Setup,
		tabletimestreamwrite.Setup,
		server.Setup,
		sshkey.Setup,
		usertransfer.Setup,
		bytematchset.Setup,
		geomatchset.Setup,
		ipsetwaf.Setup,
		ratebasedrule.Setup,
		regexmatchset.Setup,
		regexpatternset.Setup,
		rulewaf.Setup,
		sizeconstraintset.Setup,
		sqlinjectionmatchset.Setup,
		webacl.Setup,
		xssmatchset.Setup,
		bytematchsetwafregional.Setup,
		geomatchsetwafregional.Setup,
		ipsetwafregional.Setup,
		ratebasedrulewafregional.Setup,
		regexmatchsetwafregional.Setup,
		regexpatternsetwafregional.Setup,
		rulewafregional.Setup,
		sizeconstraintsetwafregional.Setup,
		sqlinjectionmatchsetwafregional.Setup,
		webaclwafregional.Setup,
		webaclassociation.Setup,
		xssmatchsetwafregional.Setup,
		ipsetwafv2.Setup,
		regexpatternsetwafv2.Setup,
		rulegroupwafv2.Setup,
		webaclwafv2.Setup,
		webaclassociationwafv2.Setup,
		webaclloggingconfiguration.Setup,
		fleetworklink.Setup,
		websitecertificateauthorityassociation.Setup,
		directoryworkspaces.Setup,
		ipgroup.Setup,
		workspaceworkspaces.Setup,
		encryptionconfig.Setup,
		groupxray.Setup,
		samplingrule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
