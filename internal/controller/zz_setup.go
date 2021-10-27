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
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane-contrib/terrajet/pkg/terraform"

	analyzer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/accessanalyzer/analyzer"
	certificate "github.com/crossplane-contrib/provider-tf-aws/internal/controller/acm/certificate"
	certificatevalidation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/acm/certificatevalidation"
	certificateacmpca "github.com/crossplane-contrib/provider-tf-aws/internal/controller/acmpca/certificate"
	certificateauthority "github.com/crossplane-contrib/provider-tf-aws/internal/controller/acmpca/certificateauthority"
	certificateauthoritycertificate "github.com/crossplane-contrib/provider-tf-aws/internal/controller/acmpca/certificateauthoritycertificate"
	alb "github.com/crossplane-contrib/provider-tf-aws/internal/controller/alb/alb"
	listener "github.com/crossplane-contrib/provider-tf-aws/internal/controller/alb/listener"
	listenercertificate "github.com/crossplane-contrib/provider-tf-aws/internal/controller/alb/listenercertificate"
	listenerrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/alb/listenerrule"
	targetgroupattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/alb/targetgroupattachment"
	ami "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ami/ami"
	copy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ami/copy"
	frominstance "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ami/frominstance"
	launchpermission "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ami/launchpermission"
	appamplify "github.com/crossplane-contrib/provider-tf-aws/internal/controller/amplify/app"
	backendenvironment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/amplify/backendenvironment"
	branch "github.com/crossplane-contrib/provider-tf-aws/internal/controller/amplify/branch"
	domainassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/amplify/domainassociation"
	webhookamplify "github.com/crossplane-contrib/provider-tf-aws/internal/controller/amplify/webhook"
	gatewayaccount "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/gatewayaccount"
	gatewayapikey "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/gatewayapikey"
	gatewayauthorizer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/gatewayauthorizer"
	gatewaybasepathmapping "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/gatewaybasepathmapping"
	gatewayclientcertificate "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/gatewayclientcertificate"
	gatewaydeployment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/gatewaydeployment"
	gatewaydocumentationpart "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/gatewaydocumentationpart"
	gatewaydocumentationversion "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/gatewaydocumentationversion"
	gatewaydomainname "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/gatewaydomainname"
	gatewaygatewayresponse "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/gatewaygatewayresponse"
	gatewayintegration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/gatewayintegration"
	gatewayintegrationresponse "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/gatewayintegrationresponse"
	gatewaymethod "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/gatewaymethod"
	gatewaymethodresponse "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/gatewaymethodresponse"
	gatewaymethodsettings "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/gatewaymethodsettings"
	gatewaymodel "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/gatewaymodel"
	gatewayrequestvalidator "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/gatewayrequestvalidator"
	gatewayresource "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/gatewayresource"
	gatewayrestapi "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/gatewayrestapi"
	gatewayrestapipolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/gatewayrestapipolicy"
	gatewaystage "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/gatewaystage"
	gatewayusageplan "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/gatewayusageplan"
	gatewayusageplankey "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/gatewayusageplankey"
	gatewayvpclink "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/gatewayvpclink"
	api "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apigatewayv2/api"
	apimapping "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apigatewayv2/apimapping"
	authorizer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apigatewayv2/authorizer"
	deployment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apigatewayv2/deployment"
	domainname "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apigatewayv2/domainname"
	integration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apigatewayv2/integration"
	integrationresponse "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apigatewayv2/integrationresponse"
	model "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apigatewayv2/model"
	routeapigatewayv2 "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apigatewayv2/route"
	routeresponse "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apigatewayv2/routeresponse"
	stage "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apigatewayv2/stage"
	vpclink "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apigatewayv2/vpclink"
	cookiestickinesspolicyapp "github.com/crossplane-contrib/provider-tf-aws/internal/controller/app/cookiestickinesspolicy"
	policyappautoscaling "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appautoscaling/policy"
	scheduledaction "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appautoscaling/scheduledaction"
	target "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appautoscaling/target"
	applicationappconfig "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appconfig/application"
	configurationprofile "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appconfig/configurationprofile"
	deploymentappconfig "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appconfig/deployment"
	deploymentstrategy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appconfig/deploymentstrategy"
	environment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appconfig/environment"
	hostedconfigurationversion "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appconfig/hostedconfigurationversion"
	gatewayroute "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appmesh/gatewayroute"
	mesh "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appmesh/mesh"
	route "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appmesh/route"
	virtualgateway "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appmesh/virtualgateway"
	virtualnode "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appmesh/virtualnode"
	virtualrouter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appmesh/virtualrouter"
	autoscalingconfigurationversion "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apprunner/autoscalingconfigurationversion"
	connectionapprunner "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apprunner/connection"
	customdomainassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apprunner/customdomainassociation"
	service "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apprunner/service"
	apikey "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appsync/apikey"
	datasource "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appsync/datasource"
	functionappsync "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appsync/function"
	graphqlapi "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appsync/graphqlapi"
	resolver "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appsync/resolver"
	databaseathena "github.com/crossplane-contrib/provider-tf-aws/internal/controller/athena/database"
	namedquery "github.com/crossplane-contrib/provider-tf-aws/internal/controller/athena/namedquery"
	workgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/athena/workgroup"
	attachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/autoscaling/attachment"
	autoscalinggroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/autoscaling/autoscalinggroup"
	lifecyclehook "github.com/crossplane-contrib/provider-tf-aws/internal/controller/autoscaling/lifecyclehook"
	notification "github.com/crossplane-contrib/provider-tf-aws/internal/controller/autoscaling/notification"
	policy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/autoscaling/policy"
	schedule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/autoscaling/schedule"
	scalingplan "github.com/crossplane-contrib/provider-tf-aws/internal/controller/autoscalingplans/scalingplan"
	globalsettings "github.com/crossplane-contrib/provider-tf-aws/internal/controller/backup/globalsettings"
	plan "github.com/crossplane-contrib/provider-tf-aws/internal/controller/backup/plan"
	regionsettings "github.com/crossplane-contrib/provider-tf-aws/internal/controller/backup/regionsettings"
	selection "github.com/crossplane-contrib/provider-tf-aws/internal/controller/backup/selection"
	vault "github.com/crossplane-contrib/provider-tf-aws/internal/controller/backup/vault"
	vaultnotifications "github.com/crossplane-contrib/provider-tf-aws/internal/controller/backup/vaultnotifications"
	vaultpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/backup/vaultpolicy"
	computeenvironment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/batch/computeenvironment"
	jobdefinition "github.com/crossplane-contrib/provider-tf-aws/internal/controller/batch/jobdefinition"
	jobqueue "github.com/crossplane-contrib/provider-tf-aws/internal/controller/batch/jobqueue"
	budget "github.com/crossplane-contrib/provider-tf-aws/internal/controller/budgets/budget"
	budgetaction "github.com/crossplane-contrib/provider-tf-aws/internal/controller/budgets/budgetaction"
	voiceconnector "github.com/crossplane-contrib/provider-tf-aws/internal/controller/chime/voiceconnector"
	environmentec2 "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloud9/environmentec2"
	stackcloudformation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudformation/stack"
	stackset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudformation/stackset"
	stacksetinstance "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudformation/stacksetinstance"
	cachepolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudfront/cachepolicy"
	distribution "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudfront/distribution"
	functioncloudfront "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudfront/function"
	keygroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudfront/keygroup"
	monitoringsubscription "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudfront/monitoringsubscription"
	originaccessidentity "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudfront/originaccessidentity"
	originrequestpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudfront/originrequestpolicy"
	publickey "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudfront/publickey"
	realtimelogconfig "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudfront/realtimelogconfig"
	v2cluster "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudhsm/v2cluster"
	v2hsm "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudhsm/v2hsm"
	cloudtrail "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudtrail/cloudtrail"
	compositealarm "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/compositealarm"
	dashboard "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/dashboard"
	eventapidestination "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/eventapidestination"
	eventarchive "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/eventarchive"
	eventbus "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/eventbus"
	eventbuspolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/eventbuspolicy"
	eventconnection "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/eventconnection"
	eventpermission "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/eventpermission"
	eventrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/eventrule"
	eventtarget "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/eventtarget"
	logdestination "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/logdestination"
	logdestinationpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/logdestinationpolicy"
	loggroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/loggroup"
	logmetricfilter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/logmetricfilter"
	logresourcepolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/logresourcepolicy"
	logstream "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/logstream"
	logsubscriptionfilter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/logsubscriptionfilter"
	metricalarm "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/metricalarm"
	metricstream "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/metricstream"
	querydefinition "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/querydefinition"
	domaincodeartifact "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codeartifact/domain"
	domainpermissionspolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codeartifact/domainpermissionspolicy"
	repository "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codeartifact/repository"
	repositorypermissionspolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codeartifact/repositorypermissionspolicy"
	project "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codebuild/project"
	reportgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codebuild/reportgroup"
	sourcecredential "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codebuild/sourcecredential"
	webhookcodebuild "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codebuild/webhook"
	repositorycodecommit "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codecommit/repository"
	triggercodecommit "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codecommit/trigger"
	app "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codedeploy/app"
	deploymentconfig "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codedeploy/deploymentconfig"
	deploymentgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codedeploy/deploymentgroup"
	codepipeline "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codepipeline/codepipeline"
	webhook "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codepipeline/webhook"
	connectioncodestarconnections "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codestarconnections/connection"
	host "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codestarconnections/host"
	notificationrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codestarnotifications/notificationrule"
	identitypool "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cognito/identitypool"
	identitypoolrolesattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cognito/identitypoolrolesattachment"
	identityprovider "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cognito/identityprovider"
	resourceserver "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cognito/resourceserver"
	usergroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cognito/usergroup"
	userpool "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cognito/userpool"
	userpoolclient "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cognito/userpoolclient"
	userpooldomain "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cognito/userpooldomain"
	userpooluicustomization "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cognito/userpooluicustomization"
	aggregateauthorization "github.com/crossplane-contrib/provider-tf-aws/internal/controller/config/aggregateauthorization"
	configrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/config/configrule"
	configurationaggregator "github.com/crossplane-contrib/provider-tf-aws/internal/controller/config/configurationaggregator"
	configurationrecorder "github.com/crossplane-contrib/provider-tf-aws/internal/controller/config/configurationrecorder"
	conformancepack "github.com/crossplane-contrib/provider-tf-aws/internal/controller/config/conformancepack"
	deliverychannel "github.com/crossplane-contrib/provider-tf-aws/internal/controller/config/deliverychannel"
	organizationconformancepack "github.com/crossplane-contrib/provider-tf-aws/internal/controller/config/organizationconformancepack"
	organizationcustomrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/config/organizationcustomrule"
	organizationmanagedrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/config/organizationmanagedrule"
	remediationconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/config/remediationconfiguration"
	reportdefinition "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cur/reportdefinition"
	gatewaycustomer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/customer/gateway"
	pipelinedatapipeline "github.com/crossplane-contrib/provider-tf-aws/internal/controller/datapipeline/pipeline"
	agent "github.com/crossplane-contrib/provider-tf-aws/internal/controller/datasync/agent"
	locationefs "github.com/crossplane-contrib/provider-tf-aws/internal/controller/datasync/locationefs"
	locationfsxwindowsfilesystem "github.com/crossplane-contrib/provider-tf-aws/internal/controller/datasync/locationfsxwindowsfilesystem"
	locationnfs "github.com/crossplane-contrib/provider-tf-aws/internal/controller/datasync/locationnfs"
	locations3 "github.com/crossplane-contrib/provider-tf-aws/internal/controller/datasync/locations3"
	locationsmb "github.com/crossplane-contrib/provider-tf-aws/internal/controller/datasync/locationsmb"
	task "github.com/crossplane-contrib/provider-tf-aws/internal/controller/datasync/task"
	clusterdax "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dax/cluster"
	parametergroupdax "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dax/parametergroup"
	subnetgroupdax "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dax/subnetgroup"
	networkacl "github.com/crossplane-contrib/provider-tf-aws/internal/controller/default/networkacl"
	routetable "github.com/crossplane-contrib/provider-tf-aws/internal/controller/default/routetable"
	securitygroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/default/securitygroup"
	subnet "github.com/crossplane-contrib/provider-tf-aws/internal/controller/default/subnet"
	vpc "github.com/crossplane-contrib/provider-tf-aws/internal/controller/default/vpc"
	vpcdhcpoptions "github.com/crossplane-contrib/provider-tf-aws/internal/controller/default/vpcdhcpoptions"
	projectdevicefarm "github.com/crossplane-contrib/provider-tf-aws/internal/controller/devicefarm/project"
	serviceconditionalforwarder "github.com/crossplane-contrib/provider-tf-aws/internal/controller/directory/serviceconditionalforwarder"
	servicedirectory "github.com/crossplane-contrib/provider-tf-aws/internal/controller/directory/servicedirectory"
	servicelogsubscription "github.com/crossplane-contrib/provider-tf-aws/internal/controller/directory/servicelogsubscription"
	lifecyclepolicydlm "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dlm/lifecyclepolicy"
	certificatedms "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dms/certificate"
	endpoint "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dms/endpoint"
	eventsubscriptiondms "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dms/eventsubscription"
	replicationinstance "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dms/replicationinstance"
	replicationsubnetgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dms/replicationsubnetgroup"
	replicationtask "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dms/replicationtask"
	clusterdocdb "github.com/crossplane-contrib/provider-tf-aws/internal/controller/docdb/cluster"
	clusterinstancedocdb "github.com/crossplane-contrib/provider-tf-aws/internal/controller/docdb/clusterinstance"
	clusterparametergroupdocdb "github.com/crossplane-contrib/provider-tf-aws/internal/controller/docdb/clusterparametergroup"
	clustersnapshotdocdb "github.com/crossplane-contrib/provider-tf-aws/internal/controller/docdb/clustersnapshot"
	subnetgroupdocdb "github.com/crossplane-contrib/provider-tf-aws/internal/controller/docdb/subnetgroup"
	bgppeer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/bgppeer"
	connectiondx "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/connection"
	connectionassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/connectionassociation"
	gatewaydx "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/gateway"
	gatewayassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/gatewayassociation"
	gatewayassociationproposal "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/gatewayassociationproposal"
	hostedprivatevirtualinterface "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/hostedprivatevirtualinterface"
	hostedprivatevirtualinterfaceaccepter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/hostedprivatevirtualinterfaceaccepter"
	hostedpublicvirtualinterface "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/hostedpublicvirtualinterface"
	hostedpublicvirtualinterfaceaccepter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/hostedpublicvirtualinterfaceaccepter"
	hostedtransitvirtualinterface "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/hostedtransitvirtualinterface"
	hostedtransitvirtualinterfaceaccepter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/hostedtransitvirtualinterfaceaccepter"
	lag "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/lag"
	privatevirtualinterface "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/privatevirtualinterface"
	publicvirtualinterface "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/publicvirtualinterface"
	transitvirtualinterface "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/transitvirtualinterface"
	globaltable "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dynamodb/globaltable"
	kinesisstreamingdestination "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dynamodb/kinesisstreamingdestination"
	tabledynamodb "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dynamodb/table"
	tableitem "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dynamodb/tableitem"
	defaultkmskey "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ebs/defaultkmskey"
	encryptionbydefault "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ebs/encryptionbydefault"
	snapshot "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ebs/snapshot"
	snapshotcopy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ebs/snapshotcopy"
	snapshotimport "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ebs/snapshotimport"
	volume "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ebs/volume"
	availabilityzonegroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/availabilityzonegroup"
	capacityreservation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/capacityreservation"
	carriergateway "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/carriergateway"
	clientvpnauthorizationrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/clientvpnauthorizationrule"
	clientvpnendpoint "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/clientvpnendpoint"
	clientvpnnetworkassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/clientvpnnetworkassociation"
	clientvpnroute "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/clientvpnroute"
	ec2launchtemplate "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/ec2launchtemplate"
	ec2networkinterface "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/ec2networkinterface"
	eipassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/eipassociation"
	elasticip "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/elasticip"
	fleetec2 "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/fleet"
	instanceec2 "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/instance"
	ipv4cidrblockassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/ipv4cidrblockassociation"
	localgatewayroute "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/localgatewayroute"
	localgatewayroutetablevpcassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/localgatewayroutetablevpcassociation"
	managedprefixlist "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/managedprefixlist"
	networkaclec2 "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/networkacl"
	networkaclrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/networkaclrule"
	networkfirewallfirewall "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/networkfirewallfirewall"
	networkfirewallfirewallpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/networkfirewallfirewallpolicy"
	networkfirewallloggingconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/networkfirewallloggingconfiguration"
	networkfirewallresourcepolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/networkfirewallresourcepolicy"
	networkfirewallrulegroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/networkfirewallrulegroup"
	networkinterfaceattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/networkinterfaceattachment"
	networkinterfacesgattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/networkinterfacesgattachment"
	routeec2 "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/route"
	route53delegationset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/route53delegationset"
	route53healthcheck "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/route53healthcheck"
	route53hostedzonednssec "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/route53hostedzonednssec"
	route53keysigningkey "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/route53keysigningkey"
	route53querylog "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/route53querylog"
	route53record "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/route53record"
	route53resolverdnssecconfig "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/route53resolverdnssecconfig"
	route53resolverendpoint "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/route53resolverendpoint"
	route53resolverfirewallconfig "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/route53resolverfirewallconfig"
	route53resolverfirewalldomainlist "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/route53resolverfirewalldomainlist"
	route53resolverfirewallrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/route53resolverfirewallrule"
	route53resolverfirewallrulegroupassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/route53resolverfirewallrulegroupassociation"
	route53resolverquerylogconfig "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/route53resolverquerylogconfig"
	route53resolverquerylogconfigassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/route53resolverquerylogconfigassociation"
	route53resolverrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/route53resolverrule"
	route53resolverruleassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/route53resolverruleassociation"
	route53vpcassociationauthorization "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/route53vpcassociationauthorization"
	route53zone "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/route53zone"
	route53zoneassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/route53zoneassociation"
	routetableec2 "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/routetable"
	routetableassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/routetableassociation"
	securitygroupec2 "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/securitygroup"
	securitygrouprule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/securitygrouprule"
	subnetec2 "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/subnet"
	tag "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/tag"
	trafficmirrorfilter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/trafficmirrorfilter"
	trafficmirrorfilterrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/trafficmirrorfilterrule"
	trafficmirrorsession "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/trafficmirrorsession"
	trafficmirrortarget "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/trafficmirrortarget"
	transitgateway "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/transitgateway"
	transitgatewaypeeringattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/transitgatewaypeeringattachment"
	transitgatewaypeeringattachmentaccepter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/transitgatewaypeeringattachmentaccepter"
	transitgatewayprefixlistreference "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/transitgatewayprefixlistreference"
	transitgatewayroute "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/transitgatewayroute"
	transitgatewayroutetable "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/transitgatewayroutetable"
	transitgatewayroutetableassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/transitgatewayroutetableassociation"
	transitgatewayroutetablepropagation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/transitgatewayroutetablepropagation"
	transitgatewayvpcattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/transitgatewayvpcattachment"
	transitgatewayvpcattachmentaccepter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/transitgatewayvpcattachmentaccepter"
	vpcec2 "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/vpc"
	vpcdhcpoptionsec2 "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/vpcdhcpoptions"
	vpcdhcpoptionsassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/vpcdhcpoptionsassociation"
	vpcendpoint "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/vpcendpoint"
	vpcendpointconnectionnotification "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/vpcendpointconnectionnotification"
	vpcendpointroutetableassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/vpcendpointroutetableassociation"
	vpcendpointservice "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/vpcendpointservice"
	vpcendpointserviceallowedprincipal "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/vpcendpointserviceallowedprincipal"
	vpcendpointsubnetassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/vpcendpointsubnetassociation"
	vpcpeeringconnection "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/vpcpeeringconnection"
	vpcpeeringconnectionoptions "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ec2/vpcpeeringconnectionoptions"
	lifecyclepolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ecr/lifecyclepolicy"
	publicrepository "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ecr/publicrepository"
	registrypolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ecr/registrypolicy"
	replicationconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ecr/replicationconfiguration"
	repositoryecr "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ecr/repository"
	repositorypolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ecr/repositorypolicy"
	capacityprovider "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ecs/capacityprovider"
	clusterecs "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ecs/cluster"
	serviceecs "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ecs/service"
	taskdefinition "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ecs/taskdefinition"
	accesspointefs "github.com/crossplane-contrib/provider-tf-aws/internal/controller/efs/accesspoint"
	backuppolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/efs/backuppolicy"
	filesystem "github.com/crossplane-contrib/provider-tf-aws/internal/controller/efs/filesystem"
	filesystempolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/efs/filesystempolicy"
	mounttarget "github.com/crossplane-contrib/provider-tf-aws/internal/controller/efs/mounttarget"
	onlyinternetgateway "github.com/crossplane-contrib/provider-tf-aws/internal/controller/egress/onlyinternetgateway"
	addon "github.com/crossplane-contrib/provider-tf-aws/internal/controller/eks/addon"
	clustereks "github.com/crossplane-contrib/provider-tf-aws/internal/controller/eks/cluster"
	fargateprofile "github.com/crossplane-contrib/provider-tf-aws/internal/controller/eks/fargateprofile"
	identityproviderconfig "github.com/crossplane-contrib/provider-tf-aws/internal/controller/eks/identityproviderconfig"
	nodegroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/eks/nodegroup"
	beanstalkapplication "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elastic/beanstalkapplication"
	beanstalkapplicationversion "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elastic/beanstalkapplicationversion"
	beanstalkconfigurationtemplate "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elastic/beanstalkconfigurationtemplate"
	beanstalkenvironment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elastic/beanstalkenvironment"
	clusterelasticache "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elasticache/cluster"
	globalreplicationgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elasticache/globalreplicationgroup"
	parametergroupelasticache "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elasticache/parametergroup"
	replicationgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elasticache/replicationgroup"
	securitygroupelasticache "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elasticache/securitygroup"
	subnetgroupelasticache "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elasticache/subnetgroup"
	userelasticache "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elasticache/user"
	domain "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elasticsearch/domain"
	domainpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elasticsearch/domainpolicy"
	domainsamloptions "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elasticsearch/domainsamloptions"
	pipeline "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elastictranscoder/pipeline"
	preset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elastictranscoder/preset"
	attachmentelb "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elb/attachment"
	elb "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elb/elb"
	clusteremr "github.com/crossplane-contrib/provider-tf-aws/internal/controller/emr/cluster"
	instancefleet "github.com/crossplane-contrib/provider-tf-aws/internal/controller/emr/instancefleet"
	instancegroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/emr/instancegroup"
	managedscalingpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/emr/managedscalingpolicy"
	securityconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/emr/securityconfiguration"
	log "github.com/crossplane-contrib/provider-tf-aws/internal/controller/flow/log"
	adminaccount "github.com/crossplane-contrib/provider-tf-aws/internal/controller/fms/adminaccount"
	policyfms "github.com/crossplane-contrib/provider-tf-aws/internal/controller/fms/policy"
	lustrefilesystem "github.com/crossplane-contrib/provider-tf-aws/internal/controller/fsx/lustrefilesystem"
	windowsfilesystem "github.com/crossplane-contrib/provider-tf-aws/internal/controller/fsx/windowsfilesystem"
	alias "github.com/crossplane-contrib/provider-tf-aws/internal/controller/gamelift/alias"
	build "github.com/crossplane-contrib/provider-tf-aws/internal/controller/gamelift/build"
	fleet "github.com/crossplane-contrib/provider-tf-aws/internal/controller/gamelift/fleet"
	gamesessionqueue "github.com/crossplane-contrib/provider-tf-aws/internal/controller/gamelift/gamesessionqueue"
	vaultglacier "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glacier/vault"
	vaultlock "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glacier/vaultlock"
	accelerator "github.com/crossplane-contrib/provider-tf-aws/internal/controller/globalaccelerator/accelerator"
	endpointgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/globalaccelerator/endpointgroup"
	listenerglobalaccelerator "github.com/crossplane-contrib/provider-tf-aws/internal/controller/globalaccelerator/listener"
	catalogdatabase "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/catalogdatabase"
	catalogtable "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/catalogtable"
	classifier "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/classifier"
	crawler "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/crawler"
	datacatalogencryptionsettings "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/datacatalogencryptionsettings"
	devendpoint "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/devendpoint"
	job "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/job"
	mltransform "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/mltransform"
	partition "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/partition"
	registryglue "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/registry"
	resourcepolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/resourcepolicy"
	securityconfigurationglue "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/securityconfiguration"
	trigger "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/trigger"
	userdefinedfunction "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/userdefinedfunction"
	workflow "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/workflow"
	detector "github.com/crossplane-contrib/provider-tf-aws/internal/controller/guardduty/detector"
	filter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/guardduty/filter"
	inviteaccepterguardduty "github.com/crossplane-contrib/provider-tf-aws/internal/controller/guardduty/inviteaccepter"
	ipset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/guardduty/ipset"
	memberguardduty "github.com/crossplane-contrib/provider-tf-aws/internal/controller/guardduty/member"
	organizationadminaccountguardduty "github.com/crossplane-contrib/provider-tf-aws/internal/controller/guardduty/organizationadminaccount"
	organizationconfigurationguardduty "github.com/crossplane-contrib/provider-tf-aws/internal/controller/guardduty/organizationconfiguration"
	publishingdestination "github.com/crossplane-contrib/provider-tf-aws/internal/controller/guardduty/publishingdestination"
	threatintelset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/guardduty/threatintelset"
	accesskey "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/accesskey"
	accountalias "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/accountalias"
	accountpasswordpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/accountpasswordpolicy"
	groupmembership "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/groupmembership"
	grouppolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/grouppolicy"
	grouppolicyattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/grouppolicyattachment"
	iamgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/iamgroup"
	instanceprofile "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/instanceprofile"
	openidconnectprovider "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/openidconnectprovider"
	policyiam "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/policy"
	policyattachmentiam "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/policyattachment"
	role "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/role"
	rolepolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/rolepolicy"
	rolepolicyattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/rolepolicyattachment"
	samlprovider "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/samlprovider"
	servercertificate "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/servercertificate"
	servicelinkedrole "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/servicelinkedrole"
	useriam "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/user"
	usergroupmembership "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/usergroupmembership"
	userloginprofile "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/userloginprofile"
	userpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/userpolicy"
	userpolicyattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/userpolicyattachment"
	usersshkey "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iam/usersshkey"
	component "github.com/crossplane-contrib/provider-tf-aws/internal/controller/imagebuilder/component"
	distributionconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/imagebuilder/distributionconfiguration"
	image "github.com/crossplane-contrib/provider-tf-aws/internal/controller/imagebuilder/image"
	imagepipeline "github.com/crossplane-contrib/provider-tf-aws/internal/controller/imagebuilder/imagepipeline"
	imagerecipe "github.com/crossplane-contrib/provider-tf-aws/internal/controller/imagebuilder/imagerecipe"
	infrastructureconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/imagebuilder/infrastructureconfiguration"
	assessmenttarget "github.com/crossplane-contrib/provider-tf-aws/internal/controller/inspector/assessmenttarget"
	assessmenttemplate "github.com/crossplane-contrib/provider-tf-aws/internal/controller/inspector/assessmenttemplate"
	resourcegroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/inspector/resourcegroup"
	gatewayinternet "github.com/crossplane-contrib/provider-tf-aws/internal/controller/internet/gateway"
	certificateiot "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iot/certificate"
	policyiot "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iot/policy"
	policyattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iot/policyattachment"
	rolealias "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iot/rolealias"
	thing "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iot/thing"
	thingprincipalattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iot/thingprincipalattachment"
	thingtype "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iot/thingtype"
	topicrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iot/topicrule"
	pair "github.com/crossplane-contrib/provider-tf-aws/internal/controller/key/pair"
	firehosedeliverystream "github.com/crossplane-contrib/provider-tf-aws/internal/controller/kinesis/firehosedeliverystream"
	stream "github.com/crossplane-contrib/provider-tf-aws/internal/controller/kinesis/stream"
	streamconsumer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/kinesis/streamconsumer"
	videostream "github.com/crossplane-contrib/provider-tf-aws/internal/controller/kinesis/videostream"
	applicationkinesisanalyticsv2 "github.com/crossplane-contrib/provider-tf-aws/internal/controller/kinesisanalyticsv2/application"
	applicationsnapshot "github.com/crossplane-contrib/provider-tf-aws/internal/controller/kinesisanalyticsv2/applicationsnapshot"
	aliaskms "github.com/crossplane-contrib/provider-tf-aws/internal/controller/kms/alias"
	ciphertext "github.com/crossplane-contrib/provider-tf-aws/internal/controller/kms/ciphertext"
	externalkey "github.com/crossplane-contrib/provider-tf-aws/internal/controller/kms/externalkey"
	grant "github.com/crossplane-contrib/provider-tf-aws/internal/controller/kms/grant"
	key "github.com/crossplane-contrib/provider-tf-aws/internal/controller/kms/key"
	datalakesettings "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lakeformation/datalakesettings"
	permissions "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lakeformation/permissions"
	resource "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lakeformation/resource"
	aliaslambda "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lambda/alias"
	codesigningconfig "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lambda/codesigningconfig"
	eventsourcemapping "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lambda/eventsourcemapping"
	function "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lambda/function"
	functioneventinvokeconfig "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lambda/functioneventinvokeconfig"
	layerversion "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lambda/layerversion"
	permission "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lambda/permission"
	provisionedconcurrencyconfig "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lambda/provisionedconcurrencyconfig"
	configuration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/launch/configuration"
	cookiestickinesspolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lb/cookiestickinesspolicy"
	lblistener "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lb/lblistener"
	lbtargetgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lb/lbtargetgroup"
	lbtargetgroupattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lb/lbtargetgroupattachment"
	listenercertificatelb "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lb/listenercertificate"
	listenerrulelb "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lb/listenerrule"
	loadbalancer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lb/loadbalancer"
	sslnegotiationpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lb/sslnegotiationpolicy"
	bot "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lex/bot"
	botalias "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lex/botalias"
	slottype "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lex/slottype"
	association "github.com/crossplane-contrib/provider-tf-aws/internal/controller/licensemanager/association"
	licenseconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/licensemanager/licenseconfiguration"
	domainlightsail "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lightsail/domain"
	instancelightsail "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lightsail/instance"
	instancepublicports "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lightsail/instancepublicports"
	keypair "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lightsail/keypair"
	staticip "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lightsail/staticip"
	staticipattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lightsail/staticipattachment"
	balancerbackendserverpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/load/balancerbackendserverpolicy"
	balancerlistenerpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/load/balancerlistenerpolicy"
	balancerpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/load/balancerpolicy"
	memberaccountassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/macie/memberaccountassociation"
	s3bucketassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/macie/s3bucketassociation"
	account "github.com/crossplane-contrib/provider-tf-aws/internal/controller/macie2/account"
	classificationjob "github.com/crossplane-contrib/provider-tf-aws/internal/controller/macie2/classificationjob"
	customdataidentifier "github.com/crossplane-contrib/provider-tf-aws/internal/controller/macie2/customdataidentifier"
	findingsfilter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/macie2/findingsfilter"
	invitationaccepter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/macie2/invitationaccepter"
	membermacie2 "github.com/crossplane-contrib/provider-tf-aws/internal/controller/macie2/member"
	organizationadminaccountmacie2 "github.com/crossplane-contrib/provider-tf-aws/internal/controller/macie2/organizationadminaccount"
	routetableassociationmain "github.com/crossplane-contrib/provider-tf-aws/internal/controller/main/routetableassociation"
	convertqueue "github.com/crossplane-contrib/provider-tf-aws/internal/controller/media/convertqueue"
	packagechannel "github.com/crossplane-contrib/provider-tf-aws/internal/controller/media/packagechannel"
	storecontainer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/media/storecontainer"
	storecontainerpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/media/storecontainerpolicy"
	broker "github.com/crossplane-contrib/provider-tf-aws/internal/controller/mq/broker"
	clustermsk "github.com/crossplane-contrib/provider-tf-aws/internal/controller/msk/cluster"
	configurationmsk "github.com/crossplane-contrib/provider-tf-aws/internal/controller/msk/configuration"
	scramsecretassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/msk/scramsecretassociation"
	gatewaynat "github.com/crossplane-contrib/provider-tf-aws/internal/controller/nat/gateway"
	clusterneptune "github.com/crossplane-contrib/provider-tf-aws/internal/controller/neptune/cluster"
	clusterendpointneptune "github.com/crossplane-contrib/provider-tf-aws/internal/controller/neptune/clusterendpoint"
	clusterinstanceneptune "github.com/crossplane-contrib/provider-tf-aws/internal/controller/neptune/clusterinstance"
	clusterparametergroupneptune "github.com/crossplane-contrib/provider-tf-aws/internal/controller/neptune/clusterparametergroup"
	clustersnapshot "github.com/crossplane-contrib/provider-tf-aws/internal/controller/neptune/clustersnapshot"
	eventsubscription "github.com/crossplane-contrib/provider-tf-aws/internal/controller/neptune/eventsubscription"
	parametergroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/neptune/parametergroup"
	subnetgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/neptune/subnetgroup"
	application "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/application"
	customlayer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/customlayer"
	ganglialayer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/ganglialayer"
	haproxylayer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/haproxylayer"
	instance "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/instance"
	javaapplayer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/javaapplayer"
	memcachedlayer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/memcachedlayer"
	mysqllayer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/mysqllayer"
	nodejsapplayer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/nodejsapplayer"
	permissionopsworks "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/permission"
	phpapplayer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/phpapplayer"
	railsapplayer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/railsapplayer"
	rdsdbinstance "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/rdsdbinstance"
	stack "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/stack"
	staticweblayer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/staticweblayer"
	userprofile "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/userprofile"
	accountorganizations "github.com/crossplane-contrib/provider-tf-aws/internal/controller/organizations/account"
	delegatedadministrator "github.com/crossplane-contrib/provider-tf-aws/internal/controller/organizations/delegatedadministrator"
	organization "github.com/crossplane-contrib/provider-tf-aws/internal/controller/organizations/organization"
	organizationalunit "github.com/crossplane-contrib/provider-tf-aws/internal/controller/organizations/organizationalunit"
	policyorganizations "github.com/crossplane-contrib/provider-tf-aws/internal/controller/organizations/policy"
	policyattachmentorganizations "github.com/crossplane-contrib/provider-tf-aws/internal/controller/organizations/policyattachment"
	admchannel "github.com/crossplane-contrib/provider-tf-aws/internal/controller/pinpoint/admchannel"
	apnschannel "github.com/crossplane-contrib/provider-tf-aws/internal/controller/pinpoint/apnschannel"
	apnssandboxchannel "github.com/crossplane-contrib/provider-tf-aws/internal/controller/pinpoint/apnssandboxchannel"
	apnsvoipchannel "github.com/crossplane-contrib/provider-tf-aws/internal/controller/pinpoint/apnsvoipchannel"
	apnsvoipsandboxchannel "github.com/crossplane-contrib/provider-tf-aws/internal/controller/pinpoint/apnsvoipsandboxchannel"
	apppinpoint "github.com/crossplane-contrib/provider-tf-aws/internal/controller/pinpoint/app"
	baiduchannel "github.com/crossplane-contrib/provider-tf-aws/internal/controller/pinpoint/baiduchannel"
	emailchannel "github.com/crossplane-contrib/provider-tf-aws/internal/controller/pinpoint/emailchannel"
	eventstream "github.com/crossplane-contrib/provider-tf-aws/internal/controller/pinpoint/eventstream"
	gcmchannel "github.com/crossplane-contrib/provider-tf-aws/internal/controller/pinpoint/gcmchannel"
	smschannel "github.com/crossplane-contrib/provider-tf-aws/internal/controller/pinpoint/smschannel"
	workspace "github.com/crossplane-contrib/provider-tf-aws/internal/controller/prometheus/workspace"
	protocolpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/proxy/protocolpolicy"
	ledger "github.com/crossplane-contrib/provider-tf-aws/internal/controller/qldb/ledger"
	user "github.com/crossplane-contrib/provider-tf-aws/internal/controller/quicksight/user"
	principalassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ram/principalassociation"
	resourceassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ram/resourceassociation"
	resourceshare "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ram/resourceshare"
	resourceshareaccepter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ram/resourceshareaccepter"
	cluster "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/cluster"
	clusterendpoint "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/clusterendpoint"
	clusterinstance "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/clusterinstance"
	clusterparametergroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/clusterparametergroup"
	clusterroleassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/clusterroleassociation"
	dbclustersnapshot "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/dbclustersnapshot"
	dbeventsubscription "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/dbeventsubscription"
	dbinstance "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/dbinstance"
	dbinstanceroleassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/dbinstanceroleassociation"
	dboptiongroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/dboptiongroup"
	dbparametergroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/dbparametergroup"
	dbproxy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/dbproxy"
	dbproxydefaulttargetgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/dbproxydefaulttargetgroup"
	dbproxyendpoint "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/dbproxyendpoint"
	dbproxytarget "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/dbproxytarget"
	dbsecuritygroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/dbsecuritygroup"
	dbsnapshot "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/dbsnapshot"
	dbsubnetgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/dbsubnetgroup"
	globalcluster "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/globalcluster"
	clusterredshift "github.com/crossplane-contrib/provider-tf-aws/internal/controller/redshift/cluster"
	eventsubscriptionredshift "github.com/crossplane-contrib/provider-tf-aws/internal/controller/redshift/eventsubscription"
	parametergroupredshift "github.com/crossplane-contrib/provider-tf-aws/internal/controller/redshift/parametergroup"
	securitygroupredshift "github.com/crossplane-contrib/provider-tf-aws/internal/controller/redshift/securitygroup"
	snapshotcopygrant "github.com/crossplane-contrib/provider-tf-aws/internal/controller/redshift/snapshotcopygrant"
	snapshotschedule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/redshift/snapshotschedule"
	snapshotscheduleassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/redshift/snapshotscheduleassociation"
	subnetgroupredshift "github.com/crossplane-contrib/provider-tf-aws/internal/controller/redshift/subnetgroup"
	accesspoint "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3/accesspoint"
	accountpublicaccessblock "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3/accountpublicaccessblock"
	bucket "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3/bucket"
	bucketanalyticsconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3/bucketanalyticsconfiguration"
	bucketinventory "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3/bucketinventory"
	bucketmetric "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3/bucketmetric"
	bucketnotification "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3/bucketnotification"
	bucketobject "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3/bucketobject"
	bucketownershipcontrols "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3/bucketownershipcontrols"
	bucketpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3/bucketpolicy"
	bucketpublicaccessblock "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3/bucketpublicaccessblock"
	objectcopy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3/objectcopy"
	buckets3control "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3control/bucket"
	bucketlifecycleconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3control/bucketlifecycleconfiguration"
	bucketpolicys3control "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3control/bucketpolicy"
	endpoints3outposts "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3outposts/endpoint"
	appsagemaker "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/app"
	appimageconfig "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/appimageconfig"
	coderepository "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/coderepository"
	domainsagemaker "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/domain"
	endpointsagemaker "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/endpoint"
	endpointconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/endpointconfiguration"
	featuregroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/featuregroup"
	imagesagemaker "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/image"
	imageversion "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/imageversion"
	modelsagemaker "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/model"
	modelpackagegroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/modelpackagegroup"
	notebookinstance "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/notebookinstance"
	notebookinstancelifecycleconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/notebookinstancelifecycleconfiguration"
	userprofilesagemaker "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/userprofile"
	workforce "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/workforce"
	workteam "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/workteam"
	discoverer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/schemas/discoverer"
	registry "github.com/crossplane-contrib/provider-tf-aws/internal/controller/schemas/registry"
	schema "github.com/crossplane-contrib/provider-tf-aws/internal/controller/schemas/schema"
	secret "github.com/crossplane-contrib/provider-tf-aws/internal/controller/secretsmanager/secret"
	secretpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/secretsmanager/secretpolicy"
	secretrotation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/secretsmanager/secretrotation"
	secretversion "github.com/crossplane-contrib/provider-tf-aws/internal/controller/secretsmanager/secretversion"
	actiontarget "github.com/crossplane-contrib/provider-tf-aws/internal/controller/securityhub/actiontarget"
	insight "github.com/crossplane-contrib/provider-tf-aws/internal/controller/securityhub/insight"
	inviteaccepter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/securityhub/inviteaccepter"
	member "github.com/crossplane-contrib/provider-tf-aws/internal/controller/securityhub/member"
	organizationadminaccount "github.com/crossplane-contrib/provider-tf-aws/internal/controller/securityhub/organizationadminaccount"
	organizationconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/securityhub/organizationconfiguration"
	productsubscription "github.com/crossplane-contrib/provider-tf-aws/internal/controller/securityhub/productsubscription"
	standardscontrol "github.com/crossplane-contrib/provider-tf-aws/internal/controller/securityhub/standardscontrol"
	standardssubscription "github.com/crossplane-contrib/provider-tf-aws/internal/controller/securityhub/standardssubscription"
	cloudformationstack "github.com/crossplane-contrib/provider-tf-aws/internal/controller/serverlessapplicationrepository/cloudformationstack"
	discoveryhttpnamespace "github.com/crossplane-contrib/provider-tf-aws/internal/controller/service/discoveryhttpnamespace"
	discoveryprivatednsnamespace "github.com/crossplane-contrib/provider-tf-aws/internal/controller/service/discoveryprivatednsnamespace"
	discoverypublicdnsnamespace "github.com/crossplane-contrib/provider-tf-aws/internal/controller/service/discoverypublicdnsnamespace"
	discoveryservice "github.com/crossplane-contrib/provider-tf-aws/internal/controller/service/discoveryservice"
	budgetresourceassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/servicecatalog/budgetresourceassociation"
	constraint "github.com/crossplane-contrib/provider-tf-aws/internal/controller/servicecatalog/constraint"
	organizationsaccess "github.com/crossplane-contrib/provider-tf-aws/internal/controller/servicecatalog/organizationsaccess"
	portfolio "github.com/crossplane-contrib/provider-tf-aws/internal/controller/servicecatalog/portfolio"
	portfolioshare "github.com/crossplane-contrib/provider-tf-aws/internal/controller/servicecatalog/portfolioshare"
	principalportfolioassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/servicecatalog/principalportfolioassociation"
	product "github.com/crossplane-contrib/provider-tf-aws/internal/controller/servicecatalog/product"
	productportfolioassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/servicecatalog/productportfolioassociation"
	provisionedproduct "github.com/crossplane-contrib/provider-tf-aws/internal/controller/servicecatalog/provisionedproduct"
	provisioningartifact "github.com/crossplane-contrib/provider-tf-aws/internal/controller/servicecatalog/provisioningartifact"
	serviceaction "github.com/crossplane-contrib/provider-tf-aws/internal/controller/servicecatalog/serviceaction"
	tagoption "github.com/crossplane-contrib/provider-tf-aws/internal/controller/servicecatalog/tagoption"
	tagoptionresourceassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/servicecatalog/tagoptionresourceassociation"
	servicequota "github.com/crossplane-contrib/provider-tf-aws/internal/controller/servicequotas/servicequota"
	activereceiptruleset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ses/activereceiptruleset"
	configurationset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ses/configurationset"
	domaindkim "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ses/domaindkim"
	domainidentity "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ses/domainidentity"
	domainidentityverification "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ses/domainidentityverification"
	domainmailfrom "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ses/domainmailfrom"
	emailidentity "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ses/emailidentity"
	eventdestination "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ses/eventdestination"
	identitynotificationtopic "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ses/identitynotificationtopic"
	identitypolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ses/identitypolicy"
	receiptfilter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ses/receiptfilter"
	receiptrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ses/receiptrule"
	receiptruleset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ses/receiptruleset"
	template "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ses/template"
	activity "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sfn/activity"
	statemachine "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sfn/statemachine"
	protection "github.com/crossplane-contrib/provider-tf-aws/internal/controller/shield/protection"
	signingjob "github.com/crossplane-contrib/provider-tf-aws/internal/controller/signer/signingjob"
	signingprofile "github.com/crossplane-contrib/provider-tf-aws/internal/controller/signer/signingprofile"
	signingprofilepermission "github.com/crossplane-contrib/provider-tf-aws/internal/controller/signer/signingprofilepermission"
	domainsimpledb "github.com/crossplane-contrib/provider-tf-aws/internal/controller/simpledb/domain"
	createvolumepermission "github.com/crossplane-contrib/provider-tf-aws/internal/controller/snapshot/createvolumepermission"
	platformapplication "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sns/platformapplication"
	smspreferences "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sns/smspreferences"
	topic "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sns/topic"
	topicpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sns/topicpolicy"
	topicsubscription "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sns/topicsubscription"
	datafeedsubscription "github.com/crossplane-contrib/provider-tf-aws/internal/controller/spot/datafeedsubscription"
	fleetrequest "github.com/crossplane-contrib/provider-tf-aws/internal/controller/spot/fleetrequest"
	instancerequest "github.com/crossplane-contrib/provider-tf-aws/internal/controller/spot/instancerequest"
	queue "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sqs/queue"
	queuepolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sqs/queuepolicy"
	activation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ssm/activation"
	associationssm "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ssm/association"
	document "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ssm/document"
	maintenancewindow "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ssm/maintenancewindow"
	maintenancewindowtarget "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ssm/maintenancewindowtarget"
	maintenancewindowtask "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ssm/maintenancewindowtask"
	patchbaseline "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ssm/patchbaseline"
	patchgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ssm/patchgroup"
	resourcedatasync "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ssm/resourcedatasync"
	accountassignment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ssoadmin/accountassignment"
	managedpolicyattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ssoadmin/managedpolicyattachment"
	permissionset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ssoadmin/permissionset"
	permissionsetinlinepolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ssoadmin/permissionsetinlinepolicy"
	cache "github.com/crossplane-contrib/provider-tf-aws/internal/controller/storagegateway/cache"
	cachediscsivolume "github.com/crossplane-contrib/provider-tf-aws/internal/controller/storagegateway/cachediscsivolume"
	filesystemassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/storagegateway/filesystemassociation"
	gatewaystoragegateway "github.com/crossplane-contrib/provider-tf-aws/internal/controller/storagegateway/gateway"
	nfsfileshare "github.com/crossplane-contrib/provider-tf-aws/internal/controller/storagegateway/nfsfileshare"
	smbfileshare "github.com/crossplane-contrib/provider-tf-aws/internal/controller/storagegateway/smbfileshare"
	storediscsivolume "github.com/crossplane-contrib/provider-tf-aws/internal/controller/storagegateway/storediscsivolume"
	tapepool "github.com/crossplane-contrib/provider-tf-aws/internal/controller/storagegateway/tapepool"
	uploadbuffer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/storagegateway/uploadbuffer"
	workingstorage "github.com/crossplane-contrib/provider-tf-aws/internal/controller/storagegateway/workingstorage"
	domainswf "github.com/crossplane-contrib/provider-tf-aws/internal/controller/swf/domain"
	canary "github.com/crossplane-contrib/provider-tf-aws/internal/controller/synthetics/canary"
	tfaws "github.com/crossplane-contrib/provider-tf-aws/internal/controller/tfaws"
	database "github.com/crossplane-contrib/provider-tf-aws/internal/controller/timestreamwrite/database"
	table "github.com/crossplane-contrib/provider-tf-aws/internal/controller/timestreamwrite/table"
	server "github.com/crossplane-contrib/provider-tf-aws/internal/controller/transfer/server"
	sshkey "github.com/crossplane-contrib/provider-tf-aws/internal/controller/transfer/sshkey"
	usertransfer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/transfer/user"
	attachmentvolume "github.com/crossplane-contrib/provider-tf-aws/internal/controller/volume/attachment"
	connection "github.com/crossplane-contrib/provider-tf-aws/internal/controller/vpn/connection"
	connectionroute "github.com/crossplane-contrib/provider-tf-aws/internal/controller/vpn/connectionroute"
	gateway "github.com/crossplane-contrib/provider-tf-aws/internal/controller/vpn/gateway"
	gatewayattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/vpn/gatewayattachment"
	gatewayroutepropagation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/vpn/gatewayroutepropagation"
	bytematchset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/waf/bytematchset"
	geomatchset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/waf/geomatchset"
	ipsetwaf "github.com/crossplane-contrib/provider-tf-aws/internal/controller/waf/ipset"
	ratebasedrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/waf/ratebasedrule"
	regexmatchset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/waf/regexmatchset"
	regexpatternset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/waf/regexpatternset"
	rule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/waf/rule"
	sizeconstraintset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/waf/sizeconstraintset"
	sqlinjectionmatchset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/waf/sqlinjectionmatchset"
	webacl "github.com/crossplane-contrib/provider-tf-aws/internal/controller/waf/webacl"
	xssmatchset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/waf/xssmatchset"
	bytematchsetwafregional "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafregional/bytematchset"
	geomatchsetwafregional "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafregional/geomatchset"
	ipsetwafregional "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafregional/ipset"
	ratebasedrulewafregional "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafregional/ratebasedrule"
	regexmatchsetwafregional "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafregional/regexmatchset"
	regexpatternsetwafregional "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafregional/regexpatternset"
	rulewafregional "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafregional/rule"
	sizeconstraintsetwafregional "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafregional/sizeconstraintset"
	sqlinjectionmatchsetwafregional "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafregional/sqlinjectionmatchset"
	webaclwafregional "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafregional/webacl"
	webaclassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafregional/webaclassociation"
	xssmatchsetwafregional "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafregional/xssmatchset"
	ipsetwafv2 "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafv2/ipset"
	regexpatternsetwafv2 "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafv2/regexpatternset"
	rulegroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafv2/rulegroup"
	webaclassociationwafv2 "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafv2/webaclassociation"
	webaclloggingconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafv2/webaclloggingconfiguration"
	fleetworklink "github.com/crossplane-contrib/provider-tf-aws/internal/controller/worklink/fleet"
	websitecertificateauthorityassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/worklink/websitecertificateauthorityassociation"
	directory "github.com/crossplane-contrib/provider-tf-aws/internal/controller/workspaces/directory"
	ipgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/workspaces/ipgroup"
	workspaceworkspaces "github.com/crossplane-contrib/provider-tf-aws/internal/controller/workspaces/workspace"
	encryptionconfig "github.com/crossplane-contrib/provider-tf-aws/internal/controller/xray/encryptionconfig"
	samplingrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/xray/samplingrule"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, l logging.Logger, wl workqueue.RateLimiter, ps terraform.SetupFn, ws *terraform.WorkspaceStore, concurrency int) error {
	for _, setup := range []func(ctrl.Manager, logging.Logger, workqueue.RateLimiter, terraform.SetupFn, *terraform.WorkspaceStore, int) error{
		accelerator.Setup,
		accesskey.Setup,
		accesspoint.Setup,
		accesspointefs.Setup,
		account.Setup,
		accountalias.Setup,
		accountassignment.Setup,
		accountorganizations.Setup,
		accountpasswordpolicy.Setup,
		accountpublicaccessblock.Setup,
		actiontarget.Setup,
		activation.Setup,
		activereceiptruleset.Setup,
		activity.Setup,
		addon.Setup,
		admchannel.Setup,
		adminaccount.Setup,
		agent.Setup,
		aggregateauthorization.Setup,
		alb.Setup,
		alias.Setup,
		aliaskms.Setup,
		aliaslambda.Setup,
		ami.Setup,
		analyzer.Setup,
		api.Setup,
		apikey.Setup,
		apimapping.Setup,
		apnschannel.Setup,
		apnssandboxchannel.Setup,
		apnsvoipchannel.Setup,
		apnsvoipsandboxchannel.Setup,
		app.Setup,
		appamplify.Setup,
		appimageconfig.Setup,
		application.Setup,
		applicationappconfig.Setup,
		applicationkinesisanalyticsv2.Setup,
		applicationsnapshot.Setup,
		apppinpoint.Setup,
		appsagemaker.Setup,
		assessmenttarget.Setup,
		assessmenttemplate.Setup,
		association.Setup,
		associationssm.Setup,
		attachment.Setup,
		attachmentelb.Setup,
		attachmentvolume.Setup,
		authorizer.Setup,
		autoscalingconfigurationversion.Setup,
		autoscalinggroup.Setup,
		availabilityzonegroup.Setup,
		backendenvironment.Setup,
		backuppolicy.Setup,
		baiduchannel.Setup,
		balancerbackendserverpolicy.Setup,
		balancerlistenerpolicy.Setup,
		balancerpolicy.Setup,
		beanstalkapplication.Setup,
		beanstalkapplicationversion.Setup,
		beanstalkconfigurationtemplate.Setup,
		beanstalkenvironment.Setup,
		bgppeer.Setup,
		bot.Setup,
		botalias.Setup,
		branch.Setup,
		broker.Setup,
		bucket.Setup,
		bucketanalyticsconfiguration.Setup,
		bucketinventory.Setup,
		bucketlifecycleconfiguration.Setup,
		bucketmetric.Setup,
		bucketnotification.Setup,
		bucketobject.Setup,
		bucketownershipcontrols.Setup,
		bucketpolicy.Setup,
		bucketpolicys3control.Setup,
		bucketpublicaccessblock.Setup,
		buckets3control.Setup,
		budget.Setup,
		budgetaction.Setup,
		budgetresourceassociation.Setup,
		build.Setup,
		bytematchset.Setup,
		bytematchsetwafregional.Setup,
		cache.Setup,
		cachediscsivolume.Setup,
		cachepolicy.Setup,
		canary.Setup,
		capacityprovider.Setup,
		capacityreservation.Setup,
		carriergateway.Setup,
		catalogdatabase.Setup,
		catalogtable.Setup,
		certificate.Setup,
		certificateacmpca.Setup,
		certificateauthority.Setup,
		certificateauthoritycertificate.Setup,
		certificatedms.Setup,
		certificateiot.Setup,
		certificatevalidation.Setup,
		ciphertext.Setup,
		classificationjob.Setup,
		classifier.Setup,
		clientvpnauthorizationrule.Setup,
		clientvpnendpoint.Setup,
		clientvpnnetworkassociation.Setup,
		clientvpnroute.Setup,
		cloudformationstack.Setup,
		cloudtrail.Setup,
		cluster.Setup,
		clusterdax.Setup,
		clusterdocdb.Setup,
		clusterecs.Setup,
		clustereks.Setup,
		clusterelasticache.Setup,
		clusteremr.Setup,
		clusterendpoint.Setup,
		clusterendpointneptune.Setup,
		clusterinstance.Setup,
		clusterinstancedocdb.Setup,
		clusterinstanceneptune.Setup,
		clustermsk.Setup,
		clusterneptune.Setup,
		clusterparametergroup.Setup,
		clusterparametergroupdocdb.Setup,
		clusterparametergroupneptune.Setup,
		clusterredshift.Setup,
		clusterroleassociation.Setup,
		clustersnapshot.Setup,
		clustersnapshotdocdb.Setup,
		codepipeline.Setup,
		coderepository.Setup,
		codesigningconfig.Setup,
		component.Setup,
		compositealarm.Setup,
		computeenvironment.Setup,
		configrule.Setup,
		configuration.Setup,
		configurationaggregator.Setup,
		configurationmsk.Setup,
		configurationprofile.Setup,
		configurationrecorder.Setup,
		configurationset.Setup,
		conformancepack.Setup,
		connection.Setup,
		connectionapprunner.Setup,
		connectionassociation.Setup,
		connectioncodestarconnections.Setup,
		connectiondx.Setup,
		connectionroute.Setup,
		constraint.Setup,
		convertqueue.Setup,
		cookiestickinesspolicy.Setup,
		cookiestickinesspolicyapp.Setup,
		copy.Setup,
		crawler.Setup,
		createvolumepermission.Setup,
		customdataidentifier.Setup,
		customdomainassociation.Setup,
		customlayer.Setup,
		dashboard.Setup,
		database.Setup,
		databaseathena.Setup,
		datacatalogencryptionsettings.Setup,
		datafeedsubscription.Setup,
		datalakesettings.Setup,
		datasource.Setup,
		dbclustersnapshot.Setup,
		dbeventsubscription.Setup,
		dbinstance.Setup,
		dbinstanceroleassociation.Setup,
		dboptiongroup.Setup,
		dbparametergroup.Setup,
		dbproxy.Setup,
		dbproxydefaulttargetgroup.Setup,
		dbproxyendpoint.Setup,
		dbproxytarget.Setup,
		dbsecuritygroup.Setup,
		dbsnapshot.Setup,
		dbsubnetgroup.Setup,
		defaultkmskey.Setup,
		delegatedadministrator.Setup,
		deliverychannel.Setup,
		deployment.Setup,
		deploymentappconfig.Setup,
		deploymentconfig.Setup,
		deploymentgroup.Setup,
		deploymentstrategy.Setup,
		detector.Setup,
		devendpoint.Setup,
		directory.Setup,
		discoverer.Setup,
		discoveryhttpnamespace.Setup,
		discoveryprivatednsnamespace.Setup,
		discoverypublicdnsnamespace.Setup,
		discoveryservice.Setup,
		distribution.Setup,
		distributionconfiguration.Setup,
		document.Setup,
		domain.Setup,
		domainassociation.Setup,
		domaincodeartifact.Setup,
		domaindkim.Setup,
		domainidentity.Setup,
		domainidentityverification.Setup,
		domainlightsail.Setup,
		domainmailfrom.Setup,
		domainname.Setup,
		domainpermissionspolicy.Setup,
		domainpolicy.Setup,
		domainsagemaker.Setup,
		domainsamloptions.Setup,
		domainsimpledb.Setup,
		domainswf.Setup,
		ec2launchtemplate.Setup,
		ec2networkinterface.Setup,
		eipassociation.Setup,
		elasticip.Setup,
		elb.Setup,
		emailchannel.Setup,
		emailidentity.Setup,
		encryptionbydefault.Setup,
		encryptionconfig.Setup,
		endpoint.Setup,
		endpointconfiguration.Setup,
		endpointgroup.Setup,
		endpoints3outposts.Setup,
		endpointsagemaker.Setup,
		environment.Setup,
		environmentec2.Setup,
		eventapidestination.Setup,
		eventarchive.Setup,
		eventbus.Setup,
		eventbuspolicy.Setup,
		eventconnection.Setup,
		eventdestination.Setup,
		eventpermission.Setup,
		eventrule.Setup,
		eventsourcemapping.Setup,
		eventstream.Setup,
		eventsubscription.Setup,
		eventsubscriptiondms.Setup,
		eventsubscriptionredshift.Setup,
		eventtarget.Setup,
		externalkey.Setup,
		fargateprofile.Setup,
		featuregroup.Setup,
		filesystem.Setup,
		filesystemassociation.Setup,
		filesystempolicy.Setup,
		filter.Setup,
		findingsfilter.Setup,
		firehosedeliverystream.Setup,
		fleet.Setup,
		fleetec2.Setup,
		fleetrequest.Setup,
		fleetworklink.Setup,
		frominstance.Setup,
		function.Setup,
		functionappsync.Setup,
		functioncloudfront.Setup,
		functioneventinvokeconfig.Setup,
		gamesessionqueue.Setup,
		ganglialayer.Setup,
		gateway.Setup,
		gatewayaccount.Setup,
		gatewayapikey.Setup,
		gatewayassociation.Setup,
		gatewayassociationproposal.Setup,
		gatewayattachment.Setup,
		gatewayauthorizer.Setup,
		gatewaybasepathmapping.Setup,
		gatewayclientcertificate.Setup,
		gatewaycustomer.Setup,
		gatewaydeployment.Setup,
		gatewaydocumentationpart.Setup,
		gatewaydocumentationversion.Setup,
		gatewaydomainname.Setup,
		gatewaydx.Setup,
		gatewaygatewayresponse.Setup,
		gatewayintegration.Setup,
		gatewayintegrationresponse.Setup,
		gatewayinternet.Setup,
		gatewaymethod.Setup,
		gatewaymethodresponse.Setup,
		gatewaymethodsettings.Setup,
		gatewaymodel.Setup,
		gatewaynat.Setup,
		gatewayrequestvalidator.Setup,
		gatewayresource.Setup,
		gatewayrestapi.Setup,
		gatewayrestapipolicy.Setup,
		gatewayroute.Setup,
		gatewayroutepropagation.Setup,
		gatewaystage.Setup,
		gatewaystoragegateway.Setup,
		gatewayusageplan.Setup,
		gatewayusageplankey.Setup,
		gatewayvpclink.Setup,
		gcmchannel.Setup,
		geomatchset.Setup,
		geomatchsetwafregional.Setup,
		globalcluster.Setup,
		globalreplicationgroup.Setup,
		globalsettings.Setup,
		globaltable.Setup,
		grant.Setup,
		graphqlapi.Setup,
		groupmembership.Setup,
		grouppolicy.Setup,
		grouppolicyattachment.Setup,
		haproxylayer.Setup,
		host.Setup,
		hostedconfigurationversion.Setup,
		hostedprivatevirtualinterface.Setup,
		hostedprivatevirtualinterfaceaccepter.Setup,
		hostedpublicvirtualinterface.Setup,
		hostedpublicvirtualinterfaceaccepter.Setup,
		hostedtransitvirtualinterface.Setup,
		hostedtransitvirtualinterfaceaccepter.Setup,
		iamgroup.Setup,
		identitynotificationtopic.Setup,
		identitypolicy.Setup,
		identitypool.Setup,
		identitypoolrolesattachment.Setup,
		identityprovider.Setup,
		identityproviderconfig.Setup,
		image.Setup,
		imagepipeline.Setup,
		imagerecipe.Setup,
		imagesagemaker.Setup,
		imageversion.Setup,
		infrastructureconfiguration.Setup,
		insight.Setup,
		instance.Setup,
		instanceec2.Setup,
		instancefleet.Setup,
		instancegroup.Setup,
		instancelightsail.Setup,
		instanceprofile.Setup,
		instancepublicports.Setup,
		instancerequest.Setup,
		integration.Setup,
		integrationresponse.Setup,
		invitationaccepter.Setup,
		inviteaccepter.Setup,
		inviteaccepterguardduty.Setup,
		ipgroup.Setup,
		ipset.Setup,
		ipsetwaf.Setup,
		ipsetwafregional.Setup,
		ipsetwafv2.Setup,
		ipv4cidrblockassociation.Setup,
		javaapplayer.Setup,
		job.Setup,
		jobdefinition.Setup,
		jobqueue.Setup,
		key.Setup,
		keygroup.Setup,
		keypair.Setup,
		kinesisstreamingdestination.Setup,
		lag.Setup,
		launchpermission.Setup,
		layerversion.Setup,
		lblistener.Setup,
		lbtargetgroup.Setup,
		lbtargetgroupattachment.Setup,
		ledger.Setup,
		licenseconfiguration.Setup,
		lifecyclehook.Setup,
		lifecyclepolicy.Setup,
		lifecyclepolicydlm.Setup,
		listener.Setup,
		listenercertificate.Setup,
		listenercertificatelb.Setup,
		listenerglobalaccelerator.Setup,
		listenerrule.Setup,
		listenerrulelb.Setup,
		loadbalancer.Setup,
		localgatewayroute.Setup,
		localgatewayroutetablevpcassociation.Setup,
		locationefs.Setup,
		locationfsxwindowsfilesystem.Setup,
		locationnfs.Setup,
		locations3.Setup,
		locationsmb.Setup,
		log.Setup,
		logdestination.Setup,
		logdestinationpolicy.Setup,
		loggroup.Setup,
		logmetricfilter.Setup,
		logresourcepolicy.Setup,
		logstream.Setup,
		logsubscriptionfilter.Setup,
		lustrefilesystem.Setup,
		maintenancewindow.Setup,
		maintenancewindowtarget.Setup,
		maintenancewindowtask.Setup,
		managedpolicyattachment.Setup,
		managedprefixlist.Setup,
		managedscalingpolicy.Setup,
		member.Setup,
		memberaccountassociation.Setup,
		memberguardduty.Setup,
		membermacie2.Setup,
		memcachedlayer.Setup,
		mesh.Setup,
		metricalarm.Setup,
		metricstream.Setup,
		mltransform.Setup,
		model.Setup,
		modelpackagegroup.Setup,
		modelsagemaker.Setup,
		monitoringsubscription.Setup,
		mounttarget.Setup,
		mysqllayer.Setup,
		namedquery.Setup,
		networkacl.Setup,
		networkaclec2.Setup,
		networkaclrule.Setup,
		networkfirewallfirewall.Setup,
		networkfirewallfirewallpolicy.Setup,
		networkfirewallloggingconfiguration.Setup,
		networkfirewallresourcepolicy.Setup,
		networkfirewallrulegroup.Setup,
		networkinterfaceattachment.Setup,
		networkinterfacesgattachment.Setup,
		nfsfileshare.Setup,
		nodegroup.Setup,
		nodejsapplayer.Setup,
		notebookinstance.Setup,
		notebookinstancelifecycleconfiguration.Setup,
		notification.Setup,
		notificationrule.Setup,
		objectcopy.Setup,
		onlyinternetgateway.Setup,
		openidconnectprovider.Setup,
		organization.Setup,
		organizationadminaccount.Setup,
		organizationadminaccountguardduty.Setup,
		organizationadminaccountmacie2.Setup,
		organizationalunit.Setup,
		organizationconfiguration.Setup,
		organizationconfigurationguardduty.Setup,
		organizationconformancepack.Setup,
		organizationcustomrule.Setup,
		organizationmanagedrule.Setup,
		organizationsaccess.Setup,
		originaccessidentity.Setup,
		originrequestpolicy.Setup,
		packagechannel.Setup,
		pair.Setup,
		parametergroup.Setup,
		parametergroupdax.Setup,
		parametergroupelasticache.Setup,
		parametergroupredshift.Setup,
		partition.Setup,
		patchbaseline.Setup,
		patchgroup.Setup,
		permission.Setup,
		permissionopsworks.Setup,
		permissions.Setup,
		permissionset.Setup,
		permissionsetinlinepolicy.Setup,
		phpapplayer.Setup,
		pipeline.Setup,
		pipelinedatapipeline.Setup,
		plan.Setup,
		platformapplication.Setup,
		policy.Setup,
		policyappautoscaling.Setup,
		policyattachment.Setup,
		policyattachmentiam.Setup,
		policyattachmentorganizations.Setup,
		policyfms.Setup,
		policyiam.Setup,
		policyiot.Setup,
		policyorganizations.Setup,
		portfolio.Setup,
		portfolioshare.Setup,
		preset.Setup,
		principalassociation.Setup,
		principalportfolioassociation.Setup,
		privatevirtualinterface.Setup,
		product.Setup,
		productportfolioassociation.Setup,
		productsubscription.Setup,
		project.Setup,
		projectdevicefarm.Setup,
		protection.Setup,
		protocolpolicy.Setup,
		provisionedconcurrencyconfig.Setup,
		provisionedproduct.Setup,
		provisioningartifact.Setup,
		publickey.Setup,
		publicrepository.Setup,
		publicvirtualinterface.Setup,
		publishingdestination.Setup,
		querydefinition.Setup,
		queue.Setup,
		queuepolicy.Setup,
		railsapplayer.Setup,
		ratebasedrule.Setup,
		ratebasedrulewafregional.Setup,
		rdsdbinstance.Setup,
		realtimelogconfig.Setup,
		receiptfilter.Setup,
		receiptrule.Setup,
		receiptruleset.Setup,
		regexmatchset.Setup,
		regexmatchsetwafregional.Setup,
		regexpatternset.Setup,
		regexpatternsetwafregional.Setup,
		regexpatternsetwafv2.Setup,
		regionsettings.Setup,
		registry.Setup,
		registryglue.Setup,
		registrypolicy.Setup,
		remediationconfiguration.Setup,
		replicationconfiguration.Setup,
		replicationgroup.Setup,
		replicationinstance.Setup,
		replicationsubnetgroup.Setup,
		replicationtask.Setup,
		reportdefinition.Setup,
		reportgroup.Setup,
		repository.Setup,
		repositorycodecommit.Setup,
		repositoryecr.Setup,
		repositorypermissionspolicy.Setup,
		repositorypolicy.Setup,
		resolver.Setup,
		resource.Setup,
		resourceassociation.Setup,
		resourcedatasync.Setup,
		resourcegroup.Setup,
		resourcepolicy.Setup,
		resourceserver.Setup,
		resourceshare.Setup,
		resourceshareaccepter.Setup,
		role.Setup,
		rolealias.Setup,
		rolepolicy.Setup,
		rolepolicyattachment.Setup,
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
		routeapigatewayv2.Setup,
		routeec2.Setup,
		routeresponse.Setup,
		routetable.Setup,
		routetableassociation.Setup,
		routetableassociationmain.Setup,
		routetableec2.Setup,
		rule.Setup,
		rulegroup.Setup,
		rulewafregional.Setup,
		s3bucketassociation.Setup,
		samlprovider.Setup,
		samplingrule.Setup,
		scalingplan.Setup,
		schedule.Setup,
		scheduledaction.Setup,
		schema.Setup,
		scramsecretassociation.Setup,
		secret.Setup,
		secretpolicy.Setup,
		secretrotation.Setup,
		secretversion.Setup,
		securityconfiguration.Setup,
		securityconfigurationglue.Setup,
		securitygroup.Setup,
		securitygroupec2.Setup,
		securitygroupelasticache.Setup,
		securitygroupredshift.Setup,
		securitygrouprule.Setup,
		selection.Setup,
		server.Setup,
		servercertificate.Setup,
		service.Setup,
		serviceaction.Setup,
		serviceconditionalforwarder.Setup,
		servicedirectory.Setup,
		serviceecs.Setup,
		servicelinkedrole.Setup,
		servicelogsubscription.Setup,
		servicequota.Setup,
		signingjob.Setup,
		signingprofile.Setup,
		signingprofilepermission.Setup,
		sizeconstraintset.Setup,
		sizeconstraintsetwafregional.Setup,
		slottype.Setup,
		smbfileshare.Setup,
		smschannel.Setup,
		smspreferences.Setup,
		snapshot.Setup,
		snapshotcopy.Setup,
		snapshotcopygrant.Setup,
		snapshotimport.Setup,
		snapshotschedule.Setup,
		snapshotscheduleassociation.Setup,
		sourcecredential.Setup,
		sqlinjectionmatchset.Setup,
		sqlinjectionmatchsetwafregional.Setup,
		sshkey.Setup,
		sslnegotiationpolicy.Setup,
		stack.Setup,
		stackcloudformation.Setup,
		stackset.Setup,
		stacksetinstance.Setup,
		stage.Setup,
		standardscontrol.Setup,
		standardssubscription.Setup,
		statemachine.Setup,
		staticip.Setup,
		staticipattachment.Setup,
		staticweblayer.Setup,
		storecontainer.Setup,
		storecontainerpolicy.Setup,
		storediscsivolume.Setup,
		stream.Setup,
		streamconsumer.Setup,
		subnet.Setup,
		subnetec2.Setup,
		subnetgroup.Setup,
		subnetgroupdax.Setup,
		subnetgroupdocdb.Setup,
		subnetgroupelasticache.Setup,
		subnetgroupredshift.Setup,
		table.Setup,
		tabledynamodb.Setup,
		tableitem.Setup,
		tag.Setup,
		tagoption.Setup,
		tagoptionresourceassociation.Setup,
		tapepool.Setup,
		target.Setup,
		targetgroupattachment.Setup,
		task.Setup,
		taskdefinition.Setup,
		template.Setup,
		tfaws.Setup,
		thing.Setup,
		thingprincipalattachment.Setup,
		thingtype.Setup,
		threatintelset.Setup,
		topic.Setup,
		topicpolicy.Setup,
		topicrule.Setup,
		topicsubscription.Setup,
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
		transitvirtualinterface.Setup,
		trigger.Setup,
		triggercodecommit.Setup,
		uploadbuffer.Setup,
		user.Setup,
		userdefinedfunction.Setup,
		userelasticache.Setup,
		usergroup.Setup,
		usergroupmembership.Setup,
		useriam.Setup,
		userloginprofile.Setup,
		userpolicy.Setup,
		userpolicyattachment.Setup,
		userpool.Setup,
		userpoolclient.Setup,
		userpooldomain.Setup,
		userpooluicustomization.Setup,
		userprofile.Setup,
		userprofilesagemaker.Setup,
		usersshkey.Setup,
		usertransfer.Setup,
		v2cluster.Setup,
		v2hsm.Setup,
		vault.Setup,
		vaultglacier.Setup,
		vaultlock.Setup,
		vaultnotifications.Setup,
		vaultpolicy.Setup,
		videostream.Setup,
		virtualgateway.Setup,
		virtualnode.Setup,
		virtualrouter.Setup,
		voiceconnector.Setup,
		volume.Setup,
		vpc.Setup,
		vpcdhcpoptions.Setup,
		vpcdhcpoptionsassociation.Setup,
		vpcdhcpoptionsec2.Setup,
		vpcec2.Setup,
		vpcendpoint.Setup,
		vpcendpointconnectionnotification.Setup,
		vpcendpointroutetableassociation.Setup,
		vpcendpointservice.Setup,
		vpcendpointserviceallowedprincipal.Setup,
		vpcendpointsubnetassociation.Setup,
		vpclink.Setup,
		vpcpeeringconnection.Setup,
		vpcpeeringconnectionoptions.Setup,
		webacl.Setup,
		webaclassociation.Setup,
		webaclassociationwafv2.Setup,
		webaclloggingconfiguration.Setup,
		webaclwafregional.Setup,
		webhook.Setup,
		webhookamplify.Setup,
		webhookcodebuild.Setup,
		websitecertificateauthorityassociation.Setup,
		windowsfilesystem.Setup,
		workflow.Setup,
		workforce.Setup,
		workgroup.Setup,
		workingstorage.Setup,
		workspace.Setup,
		workspaceworkspaces.Setup,
		workteam.Setup,
		xssmatchset.Setup,
		xssmatchsetwafregional.Setup,
	} {
		if err := setup(mgr, l, wl, ps, ws, concurrency); err != nil {
			return err
		}
	}
	return nil
}
