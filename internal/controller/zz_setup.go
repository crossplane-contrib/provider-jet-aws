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

	analyzer "github.com/crossplane-contrib/provider-jet-aws/internal/controller/accessanalyzer/analyzer"
	certificate "github.com/crossplane-contrib/provider-jet-aws/internal/controller/acm/certificate"
	certificatevalidation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/acm/certificatevalidation"
	certificateacmpca "github.com/crossplane-contrib/provider-jet-aws/internal/controller/acmpca/certificate"
	certificateauthority "github.com/crossplane-contrib/provider-jet-aws/internal/controller/acmpca/certificateauthority"
	certificateauthoritycertificate "github.com/crossplane-contrib/provider-jet-aws/internal/controller/acmpca/certificateauthoritycertificate"
	listener "github.com/crossplane-contrib/provider-jet-aws/internal/controller/alb/listener"
	listenercertificate "github.com/crossplane-contrib/provider-jet-aws/internal/controller/alb/listenercertificate"
	listenerrule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/alb/listenerrule"
	targetgroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/alb/targetgroup"
	copy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ami/copy"
	frominstance "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ami/frominstance"
	launchpermission "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ami/launchpermission"
	app "github.com/crossplane-contrib/provider-jet-aws/internal/controller/amplify/app"
	backendenvironment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/amplify/backendenvironment"
	branch "github.com/crossplane-contrib/provider-jet-aws/internal/controller/amplify/branch"
	domainassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/amplify/domainassociation"
	webhook "github.com/crossplane-contrib/provider-jet-aws/internal/controller/amplify/webhook"
	gatewayaccount "github.com/crossplane-contrib/provider-jet-aws/internal/controller/api/gatewayaccount"
	gatewayapikey "github.com/crossplane-contrib/provider-jet-aws/internal/controller/api/gatewayapikey"
	gatewayauthorizer "github.com/crossplane-contrib/provider-jet-aws/internal/controller/api/gatewayauthorizer"
	gatewaybasepathmapping "github.com/crossplane-contrib/provider-jet-aws/internal/controller/api/gatewaybasepathmapping"
	gatewayclientcertificate "github.com/crossplane-contrib/provider-jet-aws/internal/controller/api/gatewayclientcertificate"
	gatewaydeployment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/api/gatewaydeployment"
	gatewaydocumentationpart "github.com/crossplane-contrib/provider-jet-aws/internal/controller/api/gatewaydocumentationpart"
	gatewaydocumentationversion "github.com/crossplane-contrib/provider-jet-aws/internal/controller/api/gatewaydocumentationversion"
	gatewaydomainname "github.com/crossplane-contrib/provider-jet-aws/internal/controller/api/gatewaydomainname"
	gatewaygatewayresponse "github.com/crossplane-contrib/provider-jet-aws/internal/controller/api/gatewaygatewayresponse"
	gatewayintegration "github.com/crossplane-contrib/provider-jet-aws/internal/controller/api/gatewayintegration"
	gatewayintegrationresponse "github.com/crossplane-contrib/provider-jet-aws/internal/controller/api/gatewayintegrationresponse"
	gatewaymethod "github.com/crossplane-contrib/provider-jet-aws/internal/controller/api/gatewaymethod"
	gatewaymethodresponse "github.com/crossplane-contrib/provider-jet-aws/internal/controller/api/gatewaymethodresponse"
	gatewaymethodsettings "github.com/crossplane-contrib/provider-jet-aws/internal/controller/api/gatewaymethodsettings"
	gatewaymodel "github.com/crossplane-contrib/provider-jet-aws/internal/controller/api/gatewaymodel"
	gatewayrequestvalidator "github.com/crossplane-contrib/provider-jet-aws/internal/controller/api/gatewayrequestvalidator"
	gatewayresource "github.com/crossplane-contrib/provider-jet-aws/internal/controller/api/gatewayresource"
	gatewayrestapi "github.com/crossplane-contrib/provider-jet-aws/internal/controller/api/gatewayrestapi"
	gatewayrestapipolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/api/gatewayrestapipolicy"
	gatewaystage "github.com/crossplane-contrib/provider-jet-aws/internal/controller/api/gatewaystage"
	gatewayusageplan "github.com/crossplane-contrib/provider-jet-aws/internal/controller/api/gatewayusageplan"
	gatewayusageplankey "github.com/crossplane-contrib/provider-jet-aws/internal/controller/api/gatewayusageplankey"
	gatewayvpclink "github.com/crossplane-contrib/provider-jet-aws/internal/controller/api/gatewayvpclink"
	api "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigatewayv2/api"
	apimapping "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigatewayv2/apimapping"
	authorizer "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigatewayv2/authorizer"
	deployment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigatewayv2/deployment"
	domainname "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigatewayv2/domainname"
	integration "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigatewayv2/integration"
	integrationresponse "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigatewayv2/integrationresponse"
	model "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigatewayv2/model"
	route "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigatewayv2/route"
	routeresponse "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigatewayv2/routeresponse"
	stage "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigatewayv2/stage"
	vpclink "github.com/crossplane-contrib/provider-jet-aws/internal/controller/apigatewayv2/vpclink"
	cookiestickinesspolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/app/cookiestickinesspolicy"
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
	apikey "github.com/crossplane-contrib/provider-jet-aws/internal/controller/appsync/apikey"
	datasource "github.com/crossplane-contrib/provider-jet-aws/internal/controller/appsync/datasource"
	function "github.com/crossplane-contrib/provider-jet-aws/internal/controller/appsync/function"
	graphqlapi "github.com/crossplane-contrib/provider-jet-aws/internal/controller/appsync/graphqlapi"
	resolver "github.com/crossplane-contrib/provider-jet-aws/internal/controller/appsync/resolver"
	database "github.com/crossplane-contrib/provider-jet-aws/internal/controller/athena/database"
	namedquery "github.com/crossplane-contrib/provider-jet-aws/internal/controller/athena/namedquery"
	workgroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/athena/workgroup"
	attachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/autoscaling/attachment"
	autoscalinggroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/autoscaling/autoscalinggroup"
	lifecyclehook "github.com/crossplane-contrib/provider-jet-aws/internal/controller/autoscaling/lifecyclehook"
	notification "github.com/crossplane-contrib/provider-jet-aws/internal/controller/autoscaling/notification"
	policyautoscaling "github.com/crossplane-contrib/provider-jet-aws/internal/controller/autoscaling/policy"
	schedule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/autoscaling/schedule"
	scalingplan "github.com/crossplane-contrib/provider-jet-aws/internal/controller/autoscalingplans/scalingplan"
	ami "github.com/crossplane-contrib/provider-jet-aws/internal/controller/aws/ami"
	cloudtrail "github.com/crossplane-contrib/provider-jet-aws/internal/controller/aws/cloudtrail"
	codepipeline "github.com/crossplane-contrib/provider-jet-aws/internal/controller/aws/codepipeline"
	elb "github.com/crossplane-contrib/provider-jet-aws/internal/controller/aws/elb"
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
	environmentec2 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloud9/environmentec2"
	cloudformationtype "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudformation/cloudformationtype"
	stack "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudformation/stack"
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
	v2cluster "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudhsm/v2cluster"
	v2hsm "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudhsm/v2hsm"
	compositealarm "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatch/compositealarm"
	dashboard "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatch/dashboard"
	eventapidestination "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatch/eventapidestination"
	eventarchive "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatch/eventarchive"
	eventbus "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatch/eventbus"
	eventbuspolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatch/eventbuspolicy"
	eventconnection "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatch/eventconnection"
	eventpermission "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatch/eventpermission"
	eventrule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatch/eventrule"
	eventtarget "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatch/eventtarget"
	logdestination "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatch/logdestination"
	logdestinationpolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatch/logdestinationpolicy"
	loggroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatch/loggroup"
	logmetricfilter "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatch/logmetricfilter"
	logresourcepolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatch/logresourcepolicy"
	logstream "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatch/logstream"
	logsubscriptionfilter "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatch/logsubscriptionfilter"
	metricalarm "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatch/metricalarm"
	metricstream "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatch/metricstream"
	querydefinition "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cloudwatch/querydefinition"
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
	webhookcodepipeline "github.com/crossplane-contrib/provider-jet-aws/internal/controller/codepipeline/webhook"
	connectioncodestarconnections "github.com/crossplane-contrib/provider-jet-aws/internal/controller/codestarconnections/connection"
	host "github.com/crossplane-contrib/provider-jet-aws/internal/controller/codestarconnections/host"
	notificationrule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/codestarnotifications/notificationrule"
	identitypool "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cognito/identitypool"
	identitypoolrolesattachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cognito/identitypoolrolesattachment"
	identityprovider "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cognito/identityprovider"
	resourceserver "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cognito/resourceserver"
	usergroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cognito/usergroup"
	userpool "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cognito/userpool"
	userpoolclient "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cognito/userpoolclient"
	userpooldomain "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cognito/userpooldomain"
	userpooluicustomization "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cognito/userpooluicustomization"
	aggregateauthorization "github.com/crossplane-contrib/provider-jet-aws/internal/controller/config/aggregateauthorization"
	awsconfigurationrecorderstatus "github.com/crossplane-contrib/provider-jet-aws/internal/controller/config/awsconfigurationrecorderstatus"
	configrule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/config/configrule"
	configurationaggregator "github.com/crossplane-contrib/provider-jet-aws/internal/controller/config/configurationaggregator"
	configurationrecorder "github.com/crossplane-contrib/provider-jet-aws/internal/controller/config/configurationrecorder"
	conformancepack "github.com/crossplane-contrib/provider-jet-aws/internal/controller/config/conformancepack"
	deliverychannel "github.com/crossplane-contrib/provider-jet-aws/internal/controller/config/deliverychannel"
	organizationconformancepack "github.com/crossplane-contrib/provider-jet-aws/internal/controller/config/organizationconformancepack"
	organizationcustomrule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/config/organizationcustomrule"
	organizationmanagedrule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/config/organizationmanagedrule"
	remediationconfiguration "github.com/crossplane-contrib/provider-jet-aws/internal/controller/config/remediationconfiguration"
	reportdefinition "github.com/crossplane-contrib/provider-jet-aws/internal/controller/cur/reportdefinition"
	gateway "github.com/crossplane-contrib/provider-jet-aws/internal/controller/customer/gateway"
	pipeline "github.com/crossplane-contrib/provider-jet-aws/internal/controller/datapipeline/pipeline"
	agent "github.com/crossplane-contrib/provider-jet-aws/internal/controller/datasync/agent"
	locationefs "github.com/crossplane-contrib/provider-jet-aws/internal/controller/datasync/locationefs"
	locationfsxwindowsfilesystem "github.com/crossplane-contrib/provider-jet-aws/internal/controller/datasync/locationfsxwindowsfilesystem"
	locationnfs "github.com/crossplane-contrib/provider-jet-aws/internal/controller/datasync/locationnfs"
	locations3 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/datasync/locations3"
	locationsmb "github.com/crossplane-contrib/provider-jet-aws/internal/controller/datasync/locationsmb"
	task "github.com/crossplane-contrib/provider-jet-aws/internal/controller/datasync/task"
	cluster "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dax/cluster"
	parametergroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dax/parametergroup"
	subnetgroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dax/subnetgroup"
	networkacl "github.com/crossplane-contrib/provider-jet-aws/internal/controller/default/networkacl"
	routetable "github.com/crossplane-contrib/provider-jet-aws/internal/controller/default/routetable"
	securitygroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/default/securitygroup"
	subnet "github.com/crossplane-contrib/provider-jet-aws/internal/controller/default/subnet"
	vpc "github.com/crossplane-contrib/provider-jet-aws/internal/controller/default/vpc"
	vpcdhcpoptions "github.com/crossplane-contrib/provider-jet-aws/internal/controller/default/vpcdhcpoptions"
	projectdevicefarm "github.com/crossplane-contrib/provider-jet-aws/internal/controller/devicefarm/project"
	serviceconditionalforwarder "github.com/crossplane-contrib/provider-jet-aws/internal/controller/directory/serviceconditionalforwarder"
	servicedirectory "github.com/crossplane-contrib/provider-jet-aws/internal/controller/directory/servicedirectory"
	servicelogsubscription "github.com/crossplane-contrib/provider-jet-aws/internal/controller/directory/servicelogsubscription"
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
	bgppeer "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dx/bgppeer"
	connectiondx "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dx/connection"
	connectionassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dx/connectionassociation"
	gatewaydx "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dx/gateway"
	gatewayassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dx/gatewayassociation"
	gatewayassociationproposal "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dx/gatewayassociationproposal"
	hostedprivatevirtualinterface "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dx/hostedprivatevirtualinterface"
	hostedprivatevirtualinterfaceaccepter "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dx/hostedprivatevirtualinterfaceaccepter"
	hostedpublicvirtualinterface "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dx/hostedpublicvirtualinterface"
	hostedpublicvirtualinterfaceaccepter "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dx/hostedpublicvirtualinterfaceaccepter"
	hostedtransitvirtualinterface "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dx/hostedtransitvirtualinterface"
	hostedtransitvirtualinterfaceaccepter "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dx/hostedtransitvirtualinterfaceaccepter"
	lag "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dx/lag"
	privatevirtualinterface "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dx/privatevirtualinterface"
	publicvirtualinterface "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dx/publicvirtualinterface"
	transitvirtualinterface "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dx/transitvirtualinterface"
	globaltable "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dynamodb/globaltable"
	kinesisstreamingdestination "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dynamodb/kinesisstreamingdestination"
	table "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dynamodb/table"
	tableitem "github.com/crossplane-contrib/provider-jet-aws/internal/controller/dynamodb/tableitem"
	defaultkmskey "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ebs/defaultkmskey"
	encryptionbydefault "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ebs/encryptionbydefault"
	snapshot "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ebs/snapshot"
	snapshotcopy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ebs/snapshotcopy"
	snapshotimport "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ebs/snapshotimport"
	volume "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ebs/volume"
	acl "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/acl"
	aclrule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/aclrule"
	association "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/association"
	availabilityzonegroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/availabilityzonegroup"
	capacityreservation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/capacityreservation"
	carriergateway "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/carriergateway"
	clientvpnauthorizationrule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/clientvpnauthorizationrule"
	clientvpnendpoint "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/clientvpnendpoint"
	clientvpnnetworkassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/clientvpnnetworkassociation"
	clientvpnroute "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/clientvpnroute"
	dhcpoptions "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/dhcpoptions"
	dhcpoptionsassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/dhcpoptionsassociation"
	elasticip "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/elasticip"
	endpointconnectionnotification "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/endpointconnectionnotification"
	endpointroutetableassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/endpointroutetableassociation"
	endpointservice "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/endpointservice"
	endpointserviceallowedprincipal "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/endpointserviceallowedprincipal"
	endpointsubnetassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/endpointsubnetassociation"
	fleet "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/fleet"
	instance "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/instance"
	interfaceattachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/interfaceattachment"
	interfacesgattachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/interfacesgattachment"
	ipv4cidrblockassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/ipv4cidrblockassociation"
	launchtemplate "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/launchtemplate"
	localgatewayroute "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/localgatewayroute"
	localgatewayroutetablevpcassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/localgatewayroutetablevpcassociation"
	mainroutetableassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/mainroutetableassociation"
	managedprefixlist "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/managedprefixlist"
	networkinterface "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/networkinterface"
	peeringconnectionaccepter "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/peeringconnectionaccepter"
	peeringconnectionoptions "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/peeringconnectionoptions"
	routeec2 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/route"
	routetableec2 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/routetable"
	routetableassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/routetableassociation"
	securitygroupec2 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/securitygroup"
	securitygrouprule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/securitygrouprule"
	subnetec2 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/subnet"
	tag "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/tag"
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
	vpcec2 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/vpc"
	vpcendpoint "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/vpcendpoint"
	vpcpeeringconnection "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ec2/vpcpeeringconnection"
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
	onlyinternetgateway "github.com/crossplane-contrib/provider-jet-aws/internal/controller/egress/onlyinternetgateway"
	addon "github.com/crossplane-contrib/provider-jet-aws/internal/controller/eks/addon"
	clustereks "github.com/crossplane-contrib/provider-jet-aws/internal/controller/eks/cluster"
	fargateprofile "github.com/crossplane-contrib/provider-jet-aws/internal/controller/eks/fargateprofile"
	identityproviderconfig "github.com/crossplane-contrib/provider-jet-aws/internal/controller/eks/identityproviderconfig"
	nodegroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/eks/nodegroup"
	beanstalkapplication "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elastic/beanstalkapplication"
	beanstalkapplicationversion "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elastic/beanstalkapplicationversion"
	beanstalkconfigurationtemplate "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elastic/beanstalkconfigurationtemplate"
	beanstalkenvironment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elastic/beanstalkenvironment"
	clusterelasticache "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticache/cluster"
	globalreplicationgroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticache/globalreplicationgroup"
	parametergroupelasticache "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticache/parametergroup"
	replicationgroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticache/replicationgroup"
	securitygroupelasticache "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticache/securitygroup"
	subnetgroupelasticache "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticache/subnetgroup"
	user "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticache/user"
	usergroupelasticache "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticache/usergroup"
	cookiestickinesspolicyelasticloadbalancing "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticloadbalancing/cookiestickinesspolicy"
	listenercertificateelasticloadbalancing "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticloadbalancing/listenercertificate"
	listenerruleelasticloadbalancing "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticloadbalancing/listenerrule"
	loadbalancer "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticloadbalancing/loadbalancer"
	loadbalancerlistener "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticloadbalancing/loadbalancerlistener"
	sslnegotiationpolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticloadbalancing/sslnegotiationpolicy"
	targetgroupelasticloadbalancing "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticloadbalancing/targetgroup"
	targetgroupattachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticloadbalancing/targetgroupattachment"
	domainelasticsearch "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticsearch/domain"
	domainpolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticsearch/domainpolicy"
	domainsamloptions "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elasticsearch/domainsamloptions"
	pipelineelastictranscoder "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elastictranscoder/pipeline"
	preset "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elastictranscoder/preset"
	attachmentelb "github.com/crossplane-contrib/provider-jet-aws/internal/controller/elb/attachment"
	clusteremr "github.com/crossplane-contrib/provider-jet-aws/internal/controller/emr/cluster"
	instancefleet "github.com/crossplane-contrib/provider-jet-aws/internal/controller/emr/instancefleet"
	instancegroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/emr/instancegroup"
	managedscalingpolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/emr/managedscalingpolicy"
	securityconfiguration "github.com/crossplane-contrib/provider-jet-aws/internal/controller/emr/securityconfiguration"
	log "github.com/crossplane-contrib/provider-jet-aws/internal/controller/flow/log"
	adminaccount "github.com/crossplane-contrib/provider-jet-aws/internal/controller/fms/adminaccount"
	policyfms "github.com/crossplane-contrib/provider-jet-aws/internal/controller/fms/policy"
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
	listenerglobalaccelerator "github.com/crossplane-contrib/provider-jet-aws/internal/controller/globalaccelerator/listener"
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
	resourcepolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/glue/resourcepolicy"
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
	group "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iam/group"
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
	gatewayinternet "github.com/crossplane-contrib/provider-jet-aws/internal/controller/internet/gateway"
	certificateiot "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iot/certificate"
	policyiot "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iot/policy"
	policyattachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iot/policyattachment"
	rolealias "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iot/rolealias"
	thing "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iot/thing"
	thingprincipalattachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iot/thingprincipalattachment"
	thingtype "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iot/thingtype"
	topicrule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/iot/topicrule"
	pair "github.com/crossplane-contrib/provider-jet-aws/internal/controller/key/pair"
	analyticsapplication "github.com/crossplane-contrib/provider-jet-aws/internal/controller/kinesis/analyticsapplication"
	firehosedeliverystream "github.com/crossplane-contrib/provider-jet-aws/internal/controller/kinesis/firehosedeliverystream"
	stream "github.com/crossplane-contrib/provider-jet-aws/internal/controller/kinesis/stream"
	streamconsumer "github.com/crossplane-contrib/provider-jet-aws/internal/controller/kinesis/streamconsumer"
	videostream "github.com/crossplane-contrib/provider-jet-aws/internal/controller/kinesis/videostream"
	applicationkinesisanalyticsv2 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/kinesisanalyticsv2/application"
	applicationsnapshot "github.com/crossplane-contrib/provider-jet-aws/internal/controller/kinesisanalyticsv2/applicationsnapshot"
	aliaskms "github.com/crossplane-contrib/provider-jet-aws/internal/controller/kms/alias"
	ciphertext "github.com/crossplane-contrib/provider-jet-aws/internal/controller/kms/ciphertext"
	externalkey "github.com/crossplane-contrib/provider-jet-aws/internal/controller/kms/externalkey"
	grant "github.com/crossplane-contrib/provider-jet-aws/internal/controller/kms/grant"
	key "github.com/crossplane-contrib/provider-jet-aws/internal/controller/kms/key"
	datalakesettings "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lakeformation/datalakesettings"
	permissions "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lakeformation/permissions"
	resource "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lakeformation/resource"
	aliaslambda "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lambda/alias"
	codesigningconfig "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lambda/codesigningconfig"
	eventsourcemapping "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lambda/eventsourcemapping"
	functionlambda "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lambda/function"
	functioneventinvokeconfig "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lambda/functioneventinvokeconfig"
	layerversion "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lambda/layerversion"
	permission "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lambda/permission"
	provisionedconcurrencyconfig "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lambda/provisionedconcurrencyconfig"
	configuration "github.com/crossplane-contrib/provider-jet-aws/internal/controller/launch/configuration"
	bot "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lex/bot"
	botalias "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lex/botalias"
	intent "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lex/intent"
	slottype "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lex/slottype"
	associationlicensemanager "github.com/crossplane-contrib/provider-jet-aws/internal/controller/licensemanager/association"
	licenseconfiguration "github.com/crossplane-contrib/provider-jet-aws/internal/controller/licensemanager/licenseconfiguration"
	domainlightsail "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lightsail/domain"
	instancelightsail "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lightsail/instance"
	instancepublicports "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lightsail/instancepublicports"
	keypair "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lightsail/keypair"
	staticip "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lightsail/staticip"
	staticipattachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/lightsail/staticipattachment"
	balancerbackendserverpolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/load/balancerbackendserverpolicy"
	balancerlistenerpolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/load/balancerlistenerpolicy"
	balancerpolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/load/balancerpolicy"
	memberaccountassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/macie/memberaccountassociation"
	s3bucketassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/macie/s3bucketassociation"
	account "github.com/crossplane-contrib/provider-jet-aws/internal/controller/macie2/account"
	classificationjob "github.com/crossplane-contrib/provider-jet-aws/internal/controller/macie2/classificationjob"
	customdataidentifier "github.com/crossplane-contrib/provider-jet-aws/internal/controller/macie2/customdataidentifier"
	findingsfilter "github.com/crossplane-contrib/provider-jet-aws/internal/controller/macie2/findingsfilter"
	invitationaccepter "github.com/crossplane-contrib/provider-jet-aws/internal/controller/macie2/invitationaccepter"
	membermacie2 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/macie2/member"
	organizationadminaccountmacie2 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/macie2/organizationadminaccount"
	convertqueue "github.com/crossplane-contrib/provider-jet-aws/internal/controller/media/convertqueue"
	packagechannel "github.com/crossplane-contrib/provider-jet-aws/internal/controller/media/packagechannel"
	storecontainer "github.com/crossplane-contrib/provider-jet-aws/internal/controller/media/storecontainer"
	storecontainerpolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/media/storecontainerpolicy"
	broker "github.com/crossplane-contrib/provider-jet-aws/internal/controller/mq/broker"
	configurationmq "github.com/crossplane-contrib/provider-jet-aws/internal/controller/mq/configuration"
	clustermsk "github.com/crossplane-contrib/provider-jet-aws/internal/controller/msk/cluster"
	configurationmsk "github.com/crossplane-contrib/provider-jet-aws/internal/controller/msk/configuration"
	scramsecretassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/msk/scramsecretassociation"
	gatewaynat "github.com/crossplane-contrib/provider-jet-aws/internal/controller/nat/gateway"
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
	groupplacement "github.com/crossplane-contrib/provider-jet-aws/internal/controller/placement/group"
	workspace "github.com/crossplane-contrib/provider-jet-aws/internal/controller/prometheus/workspace"
	providerconfig "github.com/crossplane-contrib/provider-jet-aws/internal/controller/providerconfig"
	protocolpolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/proxy/protocolpolicy"
	ledger "github.com/crossplane-contrib/provider-jet-aws/internal/controller/qldb/ledger"
	groupquicksight "github.com/crossplane-contrib/provider-jet-aws/internal/controller/quicksight/group"
	userquicksight "github.com/crossplane-contrib/provider-jet-aws/internal/controller/quicksight/user"
	principalassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ram/principalassociation"
	resourceassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ram/resourceassociation"
	resourceshare "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ram/resourceshare"
	resourceshareaccepter "github.com/crossplane-contrib/provider-jet-aws/internal/controller/ram/resourceshareaccepter"
	clusterendpointrds "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/clusterendpoint"
	clusterinstancerds "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/clusterinstance"
	clusterparametergrouprds "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/clusterparametergroup"
	clusterroleassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/clusterroleassociation"
	clustersnapshotrds "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/clustersnapshot"
	dbcluster "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/dbcluster"
	dbinstance "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/dbinstance"
	dbparametergroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/dbparametergroup"
	eventsubscriptionrds "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/eventsubscription"
	globalcluster "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/globalcluster"
	instanceroleassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/instanceroleassociation"
	optiongroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/optiongroup"
	proxy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/proxy"
	proxydefaulttargetgroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/proxydefaulttargetgroup"
	proxyendpoint "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/proxyendpoint"
	proxytarget "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/proxytarget"
	securitygrouprds "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/securitygroup"
	snapshotrds "github.com/crossplane-contrib/provider-jet-aws/internal/controller/rds/snapshot"
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
	resolverdnssecconfig "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/resolverdnssecconfig"
	resolverendpoint "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/resolverendpoint"
	resolverfirewallconfig "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/resolverfirewallconfig"
	resolverfirewalldomainlist "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/resolverfirewalldomainlist"
	resolverfirewallrule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/resolverfirewallrule"
	resolverfirewallrulegroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/resolverfirewallrulegroup"
	resolverfirewallrulegroupassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/resolverfirewallrulegroupassociation"
	resolverquerylogconfig "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/resolverquerylogconfig"
	resolverquerylogconfigassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/resolverquerylogconfigassociation"
	resolverrule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/resolverrule"
	resolverruleassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/resolverruleassociation"
	vpcassociationauthorization "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/vpcassociationauthorization"
	zone "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/zone"
	zoneassociation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/route53/zoneassociation"
	accesspoints3 "github.com/crossplane-contrib/provider-jet-aws/internal/controller/s3/accesspoint"
	accountpublicaccessblock "github.com/crossplane-contrib/provider-jet-aws/internal/controller/s3/accountpublicaccessblock"
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
	buckets3control "github.com/crossplane-contrib/provider-jet-aws/internal/controller/s3control/bucket"
	bucketlifecycleconfiguration "github.com/crossplane-contrib/provider-jet-aws/internal/controller/s3control/bucketlifecycleconfiguration"
	bucketpolicys3control "github.com/crossplane-contrib/provider-jet-aws/internal/controller/s3control/bucketpolicy"
	endpoints3outposts "github.com/crossplane-contrib/provider-jet-aws/internal/controller/s3outposts/endpoint"
	appsagemaker "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sagemaker/app"
	appimageconfig "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sagemaker/appimageconfig"
	coderepository "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sagemaker/coderepository"
	domainsagemaker "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sagemaker/domain"
	endpointsagemaker "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sagemaker/endpoint"
	endpointconfiguration "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sagemaker/endpointconfiguration"
	featuregroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sagemaker/featuregroup"
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
	actiontarget "github.com/crossplane-contrib/provider-jet-aws/internal/controller/securityhub/actiontarget"
	insight "github.com/crossplane-contrib/provider-jet-aws/internal/controller/securityhub/insight"
	inviteacceptersecurityhub "github.com/crossplane-contrib/provider-jet-aws/internal/controller/securityhub/inviteaccepter"
	membersecurityhub "github.com/crossplane-contrib/provider-jet-aws/internal/controller/securityhub/member"
	organizationadminaccountsecurityhub "github.com/crossplane-contrib/provider-jet-aws/internal/controller/securityhub/organizationadminaccount"
	organizationconfigurationsecurityhub "github.com/crossplane-contrib/provider-jet-aws/internal/controller/securityhub/organizationconfiguration"
	productsubscription "github.com/crossplane-contrib/provider-jet-aws/internal/controller/securityhub/productsubscription"
	standardscontrol "github.com/crossplane-contrib/provider-jet-aws/internal/controller/securityhub/standardscontrol"
	standardssubscription "github.com/crossplane-contrib/provider-jet-aws/internal/controller/securityhub/standardssubscription"
	cloudformationstack "github.com/crossplane-contrib/provider-jet-aws/internal/controller/serverlessapplicationrepository/cloudformationstack"
	discoveryhttpnamespace "github.com/crossplane-contrib/provider-jet-aws/internal/controller/service/discoveryhttpnamespace"
	discoveryprivatednsnamespace "github.com/crossplane-contrib/provider-jet-aws/internal/controller/service/discoveryprivatednsnamespace"
	discoverypublicdnsnamespace "github.com/crossplane-contrib/provider-jet-aws/internal/controller/service/discoverypublicdnsnamespace"
	discoveryservice "github.com/crossplane-contrib/provider-jet-aws/internal/controller/service/discoveryservice"
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
	createvolumepermission "github.com/crossplane-contrib/provider-jet-aws/internal/controller/snapshot/createvolumepermission"
	platformapplication "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sns/platformapplication"
	smspreferences "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sns/smspreferences"
	topic "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sns/topic"
	topicpolicy "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sns/topicpolicy"
	topicsubscription "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sns/topicsubscription"
	datafeedsubscription "github.com/crossplane-contrib/provider-jet-aws/internal/controller/spot/datafeedsubscription"
	fleetrequest "github.com/crossplane-contrib/provider-jet-aws/internal/controller/spot/fleetrequest"
	instancerequest "github.com/crossplane-contrib/provider-jet-aws/internal/controller/spot/instancerequest"
	queue "github.com/crossplane-contrib/provider-jet-aws/internal/controller/sqs/queue"
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
	attachmentvolume "github.com/crossplane-contrib/provider-jet-aws/internal/controller/volume/attachment"
	connectionvpn "github.com/crossplane-contrib/provider-jet-aws/internal/controller/vpn/connection"
	connectionroute "github.com/crossplane-contrib/provider-jet-aws/internal/controller/vpn/connectionroute"
	gatewayvpn "github.com/crossplane-contrib/provider-jet-aws/internal/controller/vpn/gateway"
	gatewayattachment "github.com/crossplane-contrib/provider-jet-aws/internal/controller/vpn/gatewayattachment"
	gatewayroutepropagation "github.com/crossplane-contrib/provider-jet-aws/internal/controller/vpn/gatewayroutepropagation"
	bytematchset "github.com/crossplane-contrib/provider-jet-aws/internal/controller/waf/bytematchset"
	geomatchset "github.com/crossplane-contrib/provider-jet-aws/internal/controller/waf/geomatchset"
	ipsetwaf "github.com/crossplane-contrib/provider-jet-aws/internal/controller/waf/ipset"
	ratebasedrule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/waf/ratebasedrule"
	regexmatchset "github.com/crossplane-contrib/provider-jet-aws/internal/controller/waf/regexmatchset"
	regexpatternset "github.com/crossplane-contrib/provider-jet-aws/internal/controller/waf/regexpatternset"
	rule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/waf/rule"
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
	directory "github.com/crossplane-contrib/provider-jet-aws/internal/controller/workspaces/directory"
	ipgroup "github.com/crossplane-contrib/provider-jet-aws/internal/controller/workspaces/ipgroup"
	workspaceworkspaces "github.com/crossplane-contrib/provider-jet-aws/internal/controller/workspaces/workspace"
	encryptionconfig "github.com/crossplane-contrib/provider-jet-aws/internal/controller/xray/encryptionconfig"
	groupxray "github.com/crossplane-contrib/provider-jet-aws/internal/controller/xray/group"
	samplingrule "github.com/crossplane-contrib/provider-jet-aws/internal/controller/xray/samplingrule"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, l logging.Logger, wl workqueue.RateLimiter, ps terraform.SetupFn, ws *terraform.WorkspaceStore, cfg *tjconfig.Provider, concurrency int) error {
	for _, setup := range []func(ctrl.Manager, logging.Logger, workqueue.RateLimiter, terraform.SetupFn, *terraform.WorkspaceStore, *tjconfig.Provider, int) error{
		analyzer.Setup,
		certificate.Setup,
		certificatevalidation.Setup,
		certificateacmpca.Setup,
		certificateauthority.Setup,
		certificateauthoritycertificate.Setup,
		listener.Setup,
		listenercertificate.Setup,
		listenerrule.Setup,
		targetgroup.Setup,
		copy.Setup,
		frominstance.Setup,
		launchpermission.Setup,
		app.Setup,
		backendenvironment.Setup,
		branch.Setup,
		domainassociation.Setup,
		webhook.Setup,
		gatewayaccount.Setup,
		gatewayapikey.Setup,
		gatewayauthorizer.Setup,
		gatewaybasepathmapping.Setup,
		gatewayclientcertificate.Setup,
		gatewaydeployment.Setup,
		gatewaydocumentationpart.Setup,
		gatewaydocumentationversion.Setup,
		gatewaydomainname.Setup,
		gatewaygatewayresponse.Setup,
		gatewayintegration.Setup,
		gatewayintegrationresponse.Setup,
		gatewaymethod.Setup,
		gatewaymethodresponse.Setup,
		gatewaymethodsettings.Setup,
		gatewaymodel.Setup,
		gatewayrequestvalidator.Setup,
		gatewayresource.Setup,
		gatewayrestapi.Setup,
		gatewayrestapipolicy.Setup,
		gatewaystage.Setup,
		gatewayusageplan.Setup,
		gatewayusageplankey.Setup,
		gatewayvpclink.Setup,
		api.Setup,
		apimapping.Setup,
		authorizer.Setup,
		deployment.Setup,
		domainname.Setup,
		integration.Setup,
		integrationresponse.Setup,
		model.Setup,
		route.Setup,
		routeresponse.Setup,
		stage.Setup,
		vpclink.Setup,
		cookiestickinesspolicy.Setup,
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
		apikey.Setup,
		datasource.Setup,
		function.Setup,
		graphqlapi.Setup,
		resolver.Setup,
		database.Setup,
		namedquery.Setup,
		workgroup.Setup,
		attachment.Setup,
		autoscalinggroup.Setup,
		lifecyclehook.Setup,
		notification.Setup,
		policyautoscaling.Setup,
		schedule.Setup,
		scalingplan.Setup,
		ami.Setup,
		cloudtrail.Setup,
		codepipeline.Setup,
		elb.Setup,
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
		environmentec2.Setup,
		cloudformationtype.Setup,
		stack.Setup,
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
		v2cluster.Setup,
		v2hsm.Setup,
		compositealarm.Setup,
		dashboard.Setup,
		eventapidestination.Setup,
		eventarchive.Setup,
		eventbus.Setup,
		eventbuspolicy.Setup,
		eventconnection.Setup,
		eventpermission.Setup,
		eventrule.Setup,
		eventtarget.Setup,
		logdestination.Setup,
		logdestinationpolicy.Setup,
		loggroup.Setup,
		logmetricfilter.Setup,
		logresourcepolicy.Setup,
		logstream.Setup,
		logsubscriptionfilter.Setup,
		metricalarm.Setup,
		metricstream.Setup,
		querydefinition.Setup,
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
		webhookcodepipeline.Setup,
		connectioncodestarconnections.Setup,
		host.Setup,
		notificationrule.Setup,
		identitypool.Setup,
		identitypoolrolesattachment.Setup,
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
		gateway.Setup,
		pipeline.Setup,
		agent.Setup,
		locationefs.Setup,
		locationfsxwindowsfilesystem.Setup,
		locationnfs.Setup,
		locations3.Setup,
		locationsmb.Setup,
		task.Setup,
		cluster.Setup,
		parametergroup.Setup,
		subnetgroup.Setup,
		networkacl.Setup,
		routetable.Setup,
		securitygroup.Setup,
		subnet.Setup,
		vpc.Setup,
		vpcdhcpoptions.Setup,
		projectdevicefarm.Setup,
		serviceconditionalforwarder.Setup,
		servicedirectory.Setup,
		servicelogsubscription.Setup,
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
		bgppeer.Setup,
		connectiondx.Setup,
		connectionassociation.Setup,
		gatewaydx.Setup,
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
		globaltable.Setup,
		kinesisstreamingdestination.Setup,
		table.Setup,
		tableitem.Setup,
		defaultkmskey.Setup,
		encryptionbydefault.Setup,
		snapshot.Setup,
		snapshotcopy.Setup,
		snapshotimport.Setup,
		volume.Setup,
		acl.Setup,
		aclrule.Setup,
		association.Setup,
		availabilityzonegroup.Setup,
		capacityreservation.Setup,
		carriergateway.Setup,
		clientvpnauthorizationrule.Setup,
		clientvpnendpoint.Setup,
		clientvpnnetworkassociation.Setup,
		clientvpnroute.Setup,
		dhcpoptions.Setup,
		dhcpoptionsassociation.Setup,
		elasticip.Setup,
		endpointconnectionnotification.Setup,
		endpointroutetableassociation.Setup,
		endpointservice.Setup,
		endpointserviceallowedprincipal.Setup,
		endpointsubnetassociation.Setup,
		fleet.Setup,
		instance.Setup,
		interfaceattachment.Setup,
		interfacesgattachment.Setup,
		ipv4cidrblockassociation.Setup,
		launchtemplate.Setup,
		localgatewayroute.Setup,
		localgatewayroutetablevpcassociation.Setup,
		mainroutetableassociation.Setup,
		managedprefixlist.Setup,
		networkinterface.Setup,
		peeringconnectionaccepter.Setup,
		peeringconnectionoptions.Setup,
		routeec2.Setup,
		routetableec2.Setup,
		routetableassociation.Setup,
		securitygroupec2.Setup,
		securitygrouprule.Setup,
		subnetec2.Setup,
		tag.Setup,
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
		vpcec2.Setup,
		vpcendpoint.Setup,
		vpcpeeringconnection.Setup,
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
		onlyinternetgateway.Setup,
		addon.Setup,
		clustereks.Setup,
		fargateprofile.Setup,
		identityproviderconfig.Setup,
		nodegroup.Setup,
		beanstalkapplication.Setup,
		beanstalkapplicationversion.Setup,
		beanstalkconfigurationtemplate.Setup,
		beanstalkenvironment.Setup,
		clusterelasticache.Setup,
		globalreplicationgroup.Setup,
		parametergroupelasticache.Setup,
		replicationgroup.Setup,
		securitygroupelasticache.Setup,
		subnetgroupelasticache.Setup,
		user.Setup,
		usergroupelasticache.Setup,
		cookiestickinesspolicyelasticloadbalancing.Setup,
		listenercertificateelasticloadbalancing.Setup,
		listenerruleelasticloadbalancing.Setup,
		loadbalancer.Setup,
		loadbalancerlistener.Setup,
		sslnegotiationpolicy.Setup,
		targetgroupelasticloadbalancing.Setup,
		targetgroupattachment.Setup,
		domainelasticsearch.Setup,
		domainpolicy.Setup,
		domainsamloptions.Setup,
		pipelineelastictranscoder.Setup,
		preset.Setup,
		attachmentelb.Setup,
		clusteremr.Setup,
		instancefleet.Setup,
		instancegroup.Setup,
		managedscalingpolicy.Setup,
		securityconfiguration.Setup,
		log.Setup,
		adminaccount.Setup,
		policyfms.Setup,
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
		listenerglobalaccelerator.Setup,
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
		resourcepolicy.Setup,
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
		group.Setup,
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
		gatewayinternet.Setup,
		certificateiot.Setup,
		policyiot.Setup,
		policyattachment.Setup,
		rolealias.Setup,
		thing.Setup,
		thingprincipalattachment.Setup,
		thingtype.Setup,
		topicrule.Setup,
		pair.Setup,
		analyticsapplication.Setup,
		firehosedeliverystream.Setup,
		stream.Setup,
		streamconsumer.Setup,
		videostream.Setup,
		applicationkinesisanalyticsv2.Setup,
		applicationsnapshot.Setup,
		aliaskms.Setup,
		ciphertext.Setup,
		externalkey.Setup,
		grant.Setup,
		key.Setup,
		datalakesettings.Setup,
		permissions.Setup,
		resource.Setup,
		aliaslambda.Setup,
		codesigningconfig.Setup,
		eventsourcemapping.Setup,
		functionlambda.Setup,
		functioneventinvokeconfig.Setup,
		layerversion.Setup,
		permission.Setup,
		provisionedconcurrencyconfig.Setup,
		configuration.Setup,
		bot.Setup,
		botalias.Setup,
		intent.Setup,
		slottype.Setup,
		associationlicensemanager.Setup,
		licenseconfiguration.Setup,
		domainlightsail.Setup,
		instancelightsail.Setup,
		instancepublicports.Setup,
		keypair.Setup,
		staticip.Setup,
		staticipattachment.Setup,
		balancerbackendserverpolicy.Setup,
		balancerlistenerpolicy.Setup,
		balancerpolicy.Setup,
		memberaccountassociation.Setup,
		s3bucketassociation.Setup,
		account.Setup,
		classificationjob.Setup,
		customdataidentifier.Setup,
		findingsfilter.Setup,
		invitationaccepter.Setup,
		membermacie2.Setup,
		organizationadminaccountmacie2.Setup,
		convertqueue.Setup,
		packagechannel.Setup,
		storecontainer.Setup,
		storecontainerpolicy.Setup,
		broker.Setup,
		configurationmq.Setup,
		clustermsk.Setup,
		configurationmsk.Setup,
		scramsecretassociation.Setup,
		gatewaynat.Setup,
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
		groupplacement.Setup,
		workspace.Setup,
		providerconfig.Setup,
		protocolpolicy.Setup,
		ledger.Setup,
		groupquicksight.Setup,
		userquicksight.Setup,
		principalassociation.Setup,
		resourceassociation.Setup,
		resourceshare.Setup,
		resourceshareaccepter.Setup,
		clusterendpointrds.Setup,
		clusterinstancerds.Setup,
		clusterparametergrouprds.Setup,
		clusterroleassociation.Setup,
		clustersnapshotrds.Setup,
		dbcluster.Setup,
		dbinstance.Setup,
		dbparametergroup.Setup,
		eventsubscriptionrds.Setup,
		globalcluster.Setup,
		instanceroleassociation.Setup,
		optiongroup.Setup,
		proxy.Setup,
		proxydefaulttargetgroup.Setup,
		proxyendpoint.Setup,
		proxytarget.Setup,
		securitygrouprds.Setup,
		snapshotrds.Setup,
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
		resolverdnssecconfig.Setup,
		resolverendpoint.Setup,
		resolverfirewallconfig.Setup,
		resolverfirewalldomainlist.Setup,
		resolverfirewallrule.Setup,
		resolverfirewallrulegroup.Setup,
		resolverfirewallrulegroupassociation.Setup,
		resolverquerylogconfig.Setup,
		resolverquerylogconfigassociation.Setup,
		resolverrule.Setup,
		resolverruleassociation.Setup,
		vpcassociationauthorization.Setup,
		zone.Setup,
		zoneassociation.Setup,
		accesspoints3.Setup,
		accountpublicaccessblock.Setup,
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
		buckets3control.Setup,
		bucketlifecycleconfiguration.Setup,
		bucketpolicys3control.Setup,
		endpoints3outposts.Setup,
		appsagemaker.Setup,
		appimageconfig.Setup,
		coderepository.Setup,
		domainsagemaker.Setup,
		endpointsagemaker.Setup,
		endpointconfiguration.Setup,
		featuregroup.Setup,
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
		discoveryhttpnamespace.Setup,
		discoveryprivatednsnamespace.Setup,
		discoverypublicdnsnamespace.Setup,
		discoveryservice.Setup,
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
		createvolumepermission.Setup,
		platformapplication.Setup,
		smspreferences.Setup,
		topic.Setup,
		topicpolicy.Setup,
		topicsubscription.Setup,
		datafeedsubscription.Setup,
		fleetrequest.Setup,
		instancerequest.Setup,
		queue.Setup,
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
		attachmentvolume.Setup,
		connectionvpn.Setup,
		connectionroute.Setup,
		gatewayvpn.Setup,
		gatewayattachment.Setup,
		gatewayroutepropagation.Setup,
		bytematchset.Setup,
		geomatchset.Setup,
		ipsetwaf.Setup,
		ratebasedrule.Setup,
		regexmatchset.Setup,
		regexpatternset.Setup,
		rule.Setup,
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
		directory.Setup,
		ipgroup.Setup,
		workspaceworkspaces.Setup,
		encryptionconfig.Setup,
		groupxray.Setup,
		samplingrule.Setup,
	} {
		if err := setup(mgr, l, wl, ps, ws, cfg, concurrency); err != nil {
			return err
		}
	}
	return nil
}
