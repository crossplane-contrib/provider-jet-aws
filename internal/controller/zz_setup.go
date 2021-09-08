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

	accessanalyzeranalyzer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/accessanalyzer/accessanalyzeranalyzer"
	acmcertificate "github.com/crossplane-contrib/provider-tf-aws/internal/controller/acm/acmcertificate"
	acmcertificatevalidation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/acm/acmcertificatevalidation"
	acmpcacertificate "github.com/crossplane-contrib/provider-tf-aws/internal/controller/acmpca/acmpcacertificate"
	acmpcacertificateauthority "github.com/crossplane-contrib/provider-tf-aws/internal/controller/acmpca/acmpcacertificateauthority"
	acmpcacertificateauthoritycertificate "github.com/crossplane-contrib/provider-tf-aws/internal/controller/acmpca/acmpcacertificateauthoritycertificate"
	alb "github.com/crossplane-contrib/provider-tf-aws/internal/controller/alb/alb"
	alblistener "github.com/crossplane-contrib/provider-tf-aws/internal/controller/alb/alblistener"
	alblistenercertificate "github.com/crossplane-contrib/provider-tf-aws/internal/controller/alb/alblistenercertificate"
	alblistenerrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/alb/alblistenerrule"
	albtargetgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/alb/albtargetgroup"
	albtargetgroupattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/alb/albtargetgroupattachment"
	ami "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ami/ami"
	amicopy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ami/amicopy"
	amifrominstance "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ami/amifrominstance"
	amilaunchpermission "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ami/amilaunchpermission"
	amplifyapp "github.com/crossplane-contrib/provider-tf-aws/internal/controller/amplify/amplifyapp"
	amplifybackendenvironment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/amplify/amplifybackendenvironment"
	amplifybranch "github.com/crossplane-contrib/provider-tf-aws/internal/controller/amplify/amplifybranch"
	amplifydomainassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/amplify/amplifydomainassociation"
	amplifywebhook "github.com/crossplane-contrib/provider-tf-aws/internal/controller/amplify/amplifywebhook"
	apigatewayaccount "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/apigatewayaccount"
	apigatewayapikey "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/apigatewayapikey"
	apigatewayauthorizer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/apigatewayauthorizer"
	apigatewaybasepathmapping "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/apigatewaybasepathmapping"
	apigatewayclientcertificate "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/apigatewayclientcertificate"
	apigatewaydeployment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/apigatewaydeployment"
	apigatewaydocumentationpart "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/apigatewaydocumentationpart"
	apigatewaydocumentationversion "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/apigatewaydocumentationversion"
	apigatewaydomainname "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/apigatewaydomainname"
	apigatewaygatewayresponse "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/apigatewaygatewayresponse"
	apigatewayintegration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/apigatewayintegration"
	apigatewayintegrationresponse "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/apigatewayintegrationresponse"
	apigatewaymethod "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/apigatewaymethod"
	apigatewaymethodresponse "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/apigatewaymethodresponse"
	apigatewaymethodsettings "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/apigatewaymethodsettings"
	apigatewaymodel "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/apigatewaymodel"
	apigatewayrequestvalidator "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/apigatewayrequestvalidator"
	apigatewayresource "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/apigatewayresource"
	apigatewayrestapi "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/apigatewayrestapi"
	apigatewayrestapipolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/apigatewayrestapipolicy"
	apigatewaystage "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/apigatewaystage"
	apigatewayusageplan "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/apigatewayusageplan"
	apigatewayusageplankey "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/apigatewayusageplankey"
	apigatewayvpclink "github.com/crossplane-contrib/provider-tf-aws/internal/controller/api/apigatewayvpclink"
	apigatewayv2api "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apigatewayv2/apigatewayv2api"
	apigatewayv2apimapping "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apigatewayv2/apigatewayv2apimapping"
	apigatewayv2authorizer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apigatewayv2/apigatewayv2authorizer"
	apigatewayv2deployment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apigatewayv2/apigatewayv2deployment"
	apigatewayv2domainname "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apigatewayv2/apigatewayv2domainname"
	apigatewayv2integration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apigatewayv2/apigatewayv2integration"
	apigatewayv2integrationresponse "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apigatewayv2/apigatewayv2integrationresponse"
	apigatewayv2model "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apigatewayv2/apigatewayv2model"
	apigatewayv2route "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apigatewayv2/apigatewayv2route"
	apigatewayv2routeresponse "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apigatewayv2/apigatewayv2routeresponse"
	apigatewayv2stage "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apigatewayv2/apigatewayv2stage"
	apigatewayv2vpclink "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apigatewayv2/apigatewayv2vpclink"
	appcookiestickinesspolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/app/appcookiestickinesspolicy"
	appautoscalingpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appautoscaling/appautoscalingpolicy"
	appautoscalingscheduledaction "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appautoscaling/appautoscalingscheduledaction"
	appautoscalingtarget "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appautoscaling/appautoscalingtarget"
	appconfigapplication "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appconfig/appconfigapplication"
	appconfigconfigurationprofile "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appconfig/appconfigconfigurationprofile"
	appconfigdeployment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appconfig/appconfigdeployment"
	appconfigdeploymentstrategy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appconfig/appconfigdeploymentstrategy"
	appconfigenvironment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appconfig/appconfigenvironment"
	appconfighostedconfigurationversion "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appconfig/appconfighostedconfigurationversion"
	appmeshgatewayroute "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appmesh/appmeshgatewayroute"
	appmeshmesh "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appmesh/appmeshmesh"
	appmeshroute "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appmesh/appmeshroute"
	appmeshvirtualgateway "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appmesh/appmeshvirtualgateway"
	appmeshvirtualnode "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appmesh/appmeshvirtualnode"
	appmeshvirtualrouter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appmesh/appmeshvirtualrouter"
	appmeshvirtualservice "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appmesh/appmeshvirtualservice"
	apprunnerautoscalingconfigurationversion "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apprunner/apprunnerautoscalingconfigurationversion"
	apprunnerconnection "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apprunner/apprunnerconnection"
	apprunnercustomdomainassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apprunner/apprunnercustomdomainassociation"
	apprunnerservice "github.com/crossplane-contrib/provider-tf-aws/internal/controller/apprunner/apprunnerservice"
	appsyncapikey "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appsync/appsyncapikey"
	appsyncdatasource "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appsync/appsyncdatasource"
	appsyncfunction "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appsync/appsyncfunction"
	appsyncgraphqlapi "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appsync/appsyncgraphqlapi"
	appsyncresolver "github.com/crossplane-contrib/provider-tf-aws/internal/controller/appsync/appsyncresolver"
	athenadatabase "github.com/crossplane-contrib/provider-tf-aws/internal/controller/athena/athenadatabase"
	athenanamedquery "github.com/crossplane-contrib/provider-tf-aws/internal/controller/athena/athenanamedquery"
	athenaworkgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/athena/athenaworkgroup"
	autoscalingattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/autoscaling/autoscalingattachment"
	autoscalinggroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/autoscaling/autoscalinggroup"
	autoscalinglifecyclehook "github.com/crossplane-contrib/provider-tf-aws/internal/controller/autoscaling/autoscalinglifecyclehook"
	autoscalingnotification "github.com/crossplane-contrib/provider-tf-aws/internal/controller/autoscaling/autoscalingnotification"
	autoscalingpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/autoscaling/autoscalingpolicy"
	autoscalingschedule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/autoscaling/autoscalingschedule"
	autoscalingplansscalingplan "github.com/crossplane-contrib/provider-tf-aws/internal/controller/autoscalingplans/autoscalingplansscalingplan"
	backupglobalsettings "github.com/crossplane-contrib/provider-tf-aws/internal/controller/backup/backupglobalsettings"
	backupplan "github.com/crossplane-contrib/provider-tf-aws/internal/controller/backup/backupplan"
	backupregionsettings "github.com/crossplane-contrib/provider-tf-aws/internal/controller/backup/backupregionsettings"
	backupselection "github.com/crossplane-contrib/provider-tf-aws/internal/controller/backup/backupselection"
	backupvault "github.com/crossplane-contrib/provider-tf-aws/internal/controller/backup/backupvault"
	backupvaultnotifications "github.com/crossplane-contrib/provider-tf-aws/internal/controller/backup/backupvaultnotifications"
	backupvaultpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/backup/backupvaultpolicy"
	batchcomputeenvironment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/batch/batchcomputeenvironment"
	batchjobdefinition "github.com/crossplane-contrib/provider-tf-aws/internal/controller/batch/batchjobdefinition"
	batchjobqueue "github.com/crossplane-contrib/provider-tf-aws/internal/controller/batch/batchjobqueue"
	budgetsbudget "github.com/crossplane-contrib/provider-tf-aws/internal/controller/budgets/budgetsbudget"
	budgetsbudgetaction "github.com/crossplane-contrib/provider-tf-aws/internal/controller/budgets/budgetsbudgetaction"
	chimevoiceconnector "github.com/crossplane-contrib/provider-tf-aws/internal/controller/chime/chimevoiceconnector"
	cloud9environmentec2 "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloud9/cloud9environmentec2"
	cloudformationstack "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudformation/cloudformationstack"
	cloudformationstackset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudformation/cloudformationstackset"
	cloudformationstacksetinstance "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudformation/cloudformationstacksetinstance"
	cloudformationtype "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudformation/cloudformationtype"
	cloudfrontcachepolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudfront/cloudfrontcachepolicy"
	cloudfrontdistribution "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudfront/cloudfrontdistribution"
	cloudfrontfunction "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudfront/cloudfrontfunction"
	cloudfrontkeygroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudfront/cloudfrontkeygroup"
	cloudfrontmonitoringsubscription "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudfront/cloudfrontmonitoringsubscription"
	cloudfrontoriginaccessidentity "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudfront/cloudfrontoriginaccessidentity"
	cloudfrontoriginrequestpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudfront/cloudfrontoriginrequestpolicy"
	cloudfrontpublickey "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudfront/cloudfrontpublickey"
	cloudfrontrealtimelogconfig "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudfront/cloudfrontrealtimelogconfig"
	cloudhsmv2cluster "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudhsm/cloudhsmv2cluster"
	cloudhsmv2hsm "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudhsm/cloudhsmv2hsm"
	cloudtrail "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudtrail/cloudtrail"
	cloudwatchcompositealarm "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/cloudwatchcompositealarm"
	cloudwatchdashboard "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/cloudwatchdashboard"
	cloudwatcheventapidestination "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/cloudwatcheventapidestination"
	cloudwatcheventarchive "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/cloudwatcheventarchive"
	cloudwatcheventbus "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/cloudwatcheventbus"
	cloudwatcheventbuspolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/cloudwatcheventbuspolicy"
	cloudwatcheventconnection "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/cloudwatcheventconnection"
	cloudwatcheventpermission "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/cloudwatcheventpermission"
	cloudwatcheventrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/cloudwatcheventrule"
	cloudwatcheventtarget "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/cloudwatcheventtarget"
	cloudwatchlogdestination "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/cloudwatchlogdestination"
	cloudwatchlogdestinationpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/cloudwatchlogdestinationpolicy"
	cloudwatchloggroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/cloudwatchloggroup"
	cloudwatchlogmetricfilter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/cloudwatchlogmetricfilter"
	cloudwatchlogresourcepolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/cloudwatchlogresourcepolicy"
	cloudwatchlogstream "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/cloudwatchlogstream"
	cloudwatchlogsubscriptionfilter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/cloudwatchlogsubscriptionfilter"
	cloudwatchmetricalarm "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/cloudwatchmetricalarm"
	cloudwatchmetricstream "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/cloudwatchmetricstream"
	cloudwatchquerydefinition "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cloudwatch/cloudwatchquerydefinition"
	codeartifactdomain "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codeartifact/codeartifactdomain"
	codeartifactdomainpermissionspolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codeartifact/codeartifactdomainpermissionspolicy"
	codeartifactrepository "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codeartifact/codeartifactrepository"
	codeartifactrepositorypermissionspolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codeartifact/codeartifactrepositorypermissionspolicy"
	codebuildproject "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codebuild/codebuildproject"
	codebuildreportgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codebuild/codebuildreportgroup"
	codebuildsourcecredential "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codebuild/codebuildsourcecredential"
	codebuildwebhook "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codebuild/codebuildwebhook"
	codecommitrepository "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codecommit/codecommitrepository"
	codecommittrigger "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codecommit/codecommittrigger"
	codedeployapp "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codedeploy/codedeployapp"
	codedeploydeploymentconfig "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codedeploy/codedeploydeploymentconfig"
	codedeploydeploymentgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codedeploy/codedeploydeploymentgroup"
	codepipeline "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codepipeline/codepipeline"
	codepipelinewebhook "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codepipeline/codepipelinewebhook"
	codestarconnectionsconnection "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codestarconnections/codestarconnectionsconnection"
	codestarconnectionshost "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codestarconnections/codestarconnectionshost"
	codestarnotificationsnotificationrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/codestarnotifications/codestarnotificationsnotificationrule"
	cognitoidentitypool "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cognito/cognitoidentitypool"
	cognitoidentitypoolrolesattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cognito/cognitoidentitypoolrolesattachment"
	cognitoidentityprovider "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cognito/cognitoidentityprovider"
	cognitoresourceserver "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cognito/cognitoresourceserver"
	cognitousergroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cognito/cognitousergroup"
	cognitouserpool "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cognito/cognitouserpool"
	cognitouserpoolclient "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cognito/cognitouserpoolclient"
	cognitouserpooldomain "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cognito/cognitouserpooldomain"
	cognitouserpooluicustomization "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cognito/cognitouserpooluicustomization"
	configaggregateauthorization "github.com/crossplane-contrib/provider-tf-aws/internal/controller/config/configaggregateauthorization"
	configconfigrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/config/configconfigrule"
	configconfigurationaggregator "github.com/crossplane-contrib/provider-tf-aws/internal/controller/config/configconfigurationaggregator"
	configconfigurationrecorder "github.com/crossplane-contrib/provider-tf-aws/internal/controller/config/configconfigurationrecorder"
	configconformancepack "github.com/crossplane-contrib/provider-tf-aws/internal/controller/config/configconformancepack"
	configdeliverychannel "github.com/crossplane-contrib/provider-tf-aws/internal/controller/config/configdeliverychannel"
	configorganizationconformancepack "github.com/crossplane-contrib/provider-tf-aws/internal/controller/config/configorganizationconformancepack"
	configorganizationcustomrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/config/configorganizationcustomrule"
	configorganizationmanagedrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/config/configorganizationmanagedrule"
	configremediationconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/config/configremediationconfiguration"
	curreportdefinition "github.com/crossplane-contrib/provider-tf-aws/internal/controller/cur/curreportdefinition"
	customergateway "github.com/crossplane-contrib/provider-tf-aws/internal/controller/customer/customergateway"
	datapipelinepipeline "github.com/crossplane-contrib/provider-tf-aws/internal/controller/datapipeline/datapipelinepipeline"
	datasyncagent "github.com/crossplane-contrib/provider-tf-aws/internal/controller/datasync/datasyncagent"
	datasynclocationefs "github.com/crossplane-contrib/provider-tf-aws/internal/controller/datasync/datasynclocationefs"
	datasynclocationfsxwindowsfilesystem "github.com/crossplane-contrib/provider-tf-aws/internal/controller/datasync/datasynclocationfsxwindowsfilesystem"
	datasynclocationnfs "github.com/crossplane-contrib/provider-tf-aws/internal/controller/datasync/datasynclocationnfs"
	datasynclocations3 "github.com/crossplane-contrib/provider-tf-aws/internal/controller/datasync/datasynclocations3"
	datasynclocationsmb "github.com/crossplane-contrib/provider-tf-aws/internal/controller/datasync/datasynclocationsmb"
	datasynctask "github.com/crossplane-contrib/provider-tf-aws/internal/controller/datasync/datasynctask"
	daxcluster "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dax/daxcluster"
	daxparametergroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dax/daxparametergroup"
	daxsubnetgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dax/daxsubnetgroup"
	dbclustersnapshot "github.com/crossplane-contrib/provider-tf-aws/internal/controller/db/dbclustersnapshot"
	dbeventsubscription "github.com/crossplane-contrib/provider-tf-aws/internal/controller/db/dbeventsubscription"
	dbinstance "github.com/crossplane-contrib/provider-tf-aws/internal/controller/db/dbinstance"
	dbinstanceroleassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/db/dbinstanceroleassociation"
	dboptiongroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/db/dboptiongroup"
	dbparametergroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/db/dbparametergroup"
	dbproxy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/db/dbproxy"
	dbproxydefaulttargetgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/db/dbproxydefaulttargetgroup"
	dbproxyendpoint "github.com/crossplane-contrib/provider-tf-aws/internal/controller/db/dbproxyendpoint"
	dbproxytarget "github.com/crossplane-contrib/provider-tf-aws/internal/controller/db/dbproxytarget"
	dbsecuritygroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/db/dbsecuritygroup"
	dbsnapshot "github.com/crossplane-contrib/provider-tf-aws/internal/controller/db/dbsnapshot"
	dbsubnetgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/db/dbsubnetgroup"
	defaultnetworkacl "github.com/crossplane-contrib/provider-tf-aws/internal/controller/default/defaultnetworkacl"
	defaultroutetable "github.com/crossplane-contrib/provider-tf-aws/internal/controller/default/defaultroutetable"
	defaultsecuritygroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/default/defaultsecuritygroup"
	defaultsubnet "github.com/crossplane-contrib/provider-tf-aws/internal/controller/default/defaultsubnet"
	defaultvpc "github.com/crossplane-contrib/provider-tf-aws/internal/controller/default/defaultvpc"
	defaultvpcdhcpoptions "github.com/crossplane-contrib/provider-tf-aws/internal/controller/default/defaultvpcdhcpoptions"
	devicefarmproject "github.com/crossplane-contrib/provider-tf-aws/internal/controller/devicefarm/devicefarmproject"
	directoryserviceconditionalforwarder "github.com/crossplane-contrib/provider-tf-aws/internal/controller/directory/directoryserviceconditionalforwarder"
	directoryservicedirectory "github.com/crossplane-contrib/provider-tf-aws/internal/controller/directory/directoryservicedirectory"
	directoryservicelogsubscription "github.com/crossplane-contrib/provider-tf-aws/internal/controller/directory/directoryservicelogsubscription"
	dlmlifecyclepolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dlm/dlmlifecyclepolicy"
	dmscertificate "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dms/dmscertificate"
	dmsendpoint "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dms/dmsendpoint"
	dmseventsubscription "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dms/dmseventsubscription"
	dmsreplicationinstance "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dms/dmsreplicationinstance"
	dmsreplicationsubnetgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dms/dmsreplicationsubnetgroup"
	dmsreplicationtask "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dms/dmsreplicationtask"
	docdbcluster "github.com/crossplane-contrib/provider-tf-aws/internal/controller/docdb/docdbcluster"
	docdbclusterinstance "github.com/crossplane-contrib/provider-tf-aws/internal/controller/docdb/docdbclusterinstance"
	docdbclusterparametergroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/docdb/docdbclusterparametergroup"
	docdbclustersnapshot "github.com/crossplane-contrib/provider-tf-aws/internal/controller/docdb/docdbclustersnapshot"
	docdbsubnetgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/docdb/docdbsubnetgroup"
	dxbgppeer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/dxbgppeer"
	dxconnection "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/dxconnection"
	dxconnectionassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/dxconnectionassociation"
	dxgateway "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/dxgateway"
	dxgatewayassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/dxgatewayassociation"
	dxgatewayassociationproposal "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/dxgatewayassociationproposal"
	dxhostedprivatevirtualinterface "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/dxhostedprivatevirtualinterface"
	dxhostedprivatevirtualinterfaceaccepter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/dxhostedprivatevirtualinterfaceaccepter"
	dxhostedpublicvirtualinterface "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/dxhostedpublicvirtualinterface"
	dxhostedpublicvirtualinterfaceaccepter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/dxhostedpublicvirtualinterfaceaccepter"
	dxhostedtransitvirtualinterface "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/dxhostedtransitvirtualinterface"
	dxhostedtransitvirtualinterfaceaccepter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/dxhostedtransitvirtualinterfaceaccepter"
	dxlag "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/dxlag"
	dxprivatevirtualinterface "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/dxprivatevirtualinterface"
	dxpublicvirtualinterface "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/dxpublicvirtualinterface"
	dxtransitvirtualinterface "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dx/dxtransitvirtualinterface"
	dynamodbglobaltable "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dynamodb/dynamodbglobaltable"
	dynamodbkinesisstreamingdestination "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dynamodb/dynamodbkinesisstreamingdestination"
	dynamodbtable "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dynamodb/dynamodbtable"
	dynamodbtableitem "github.com/crossplane-contrib/provider-tf-aws/internal/controller/dynamodb/dynamodbtableitem"
	ebsdefaultkmskey "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ebs/ebsdefaultkmskey"
	ebsencryptionbydefault "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ebs/ebsencryptionbydefault"
	ebssnapshot "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ebs/ebssnapshot"
	ebssnapshotcopy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ebs/ebssnapshotcopy"
	ebssnapshotimport "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ebs/ebssnapshotimport"
	ebsvolume "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ebs/ebsvolume"
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
	ecrlifecyclepolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ecr/ecrlifecyclepolicy"
	ecrregistrypolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ecr/ecrregistrypolicy"
	ecrreplicationconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ecr/ecrreplicationconfiguration"
	ecrrepository "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ecr/ecrrepository"
	ecrrepositorypolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ecr/ecrrepositorypolicy"
	ecrpublicrepository "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ecrpublic/ecrpublicrepository"
	ecscapacityprovider "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ecs/ecscapacityprovider"
	ecscluster "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ecs/ecscluster"
	ecsservice "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ecs/ecsservice"
	ecstaskdefinition "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ecs/ecstaskdefinition"
	efsaccesspoint "github.com/crossplane-contrib/provider-tf-aws/internal/controller/efs/efsaccesspoint"
	efsbackuppolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/efs/efsbackuppolicy"
	efsfilesystem "github.com/crossplane-contrib/provider-tf-aws/internal/controller/efs/efsfilesystem"
	efsfilesystempolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/efs/efsfilesystempolicy"
	efsmounttarget "github.com/crossplane-contrib/provider-tf-aws/internal/controller/efs/efsmounttarget"
	egressonlyinternetgateway "github.com/crossplane-contrib/provider-tf-aws/internal/controller/egress/egressonlyinternetgateway"
	eip "github.com/crossplane-contrib/provider-tf-aws/internal/controller/eip/eip"
	eipassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/eip/eipassociation"
	eksaddon "github.com/crossplane-contrib/provider-tf-aws/internal/controller/eks/eksaddon"
	ekscluster "github.com/crossplane-contrib/provider-tf-aws/internal/controller/eks/ekscluster"
	eksfargateprofile "github.com/crossplane-contrib/provider-tf-aws/internal/controller/eks/eksfargateprofile"
	eksidentityproviderconfig "github.com/crossplane-contrib/provider-tf-aws/internal/controller/eks/eksidentityproviderconfig"
	eksnodegroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/eks/eksnodegroup"
	elasticbeanstalkapplication "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elastic/elasticbeanstalkapplication"
	elasticbeanstalkapplicationversion "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elastic/elasticbeanstalkapplicationversion"
	elasticbeanstalkconfigurationtemplate "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elastic/elasticbeanstalkconfigurationtemplate"
	elasticbeanstalkenvironment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elastic/elasticbeanstalkenvironment"
	elasticachecluster "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elasticache/elasticachecluster"
	elasticacheglobalreplicationgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elasticache/elasticacheglobalreplicationgroup"
	elasticacheparametergroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elasticache/elasticacheparametergroup"
	elasticachereplicationgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elasticache/elasticachereplicationgroup"
	elasticachesecuritygroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elasticache/elasticachesecuritygroup"
	elasticachesubnetgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elasticache/elasticachesubnetgroup"
	elasticacheuser "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elasticache/elasticacheuser"
	elasticsearchdomain "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elasticsearch/elasticsearchdomain"
	elasticsearchdomainpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elasticsearch/elasticsearchdomainpolicy"
	elasticsearchdomainsamloptions "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elasticsearch/elasticsearchdomainsamloptions"
	elastictranscoderpipeline "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elastictranscoder/elastictranscoderpipeline"
	elastictranscoderpreset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elastictranscoder/elastictranscoderpreset"
	elb "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elb/elb"
	elbattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/elb/elbattachment"
	emrcluster "github.com/crossplane-contrib/provider-tf-aws/internal/controller/emr/emrcluster"
	emrinstancefleet "github.com/crossplane-contrib/provider-tf-aws/internal/controller/emr/emrinstancefleet"
	emrinstancegroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/emr/emrinstancegroup"
	emrmanagedscalingpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/emr/emrmanagedscalingpolicy"
	emrsecurityconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/emr/emrsecurityconfiguration"
	flowlog "github.com/crossplane-contrib/provider-tf-aws/internal/controller/flow/flowlog"
	fmsadminaccount "github.com/crossplane-contrib/provider-tf-aws/internal/controller/fms/fmsadminaccount"
	fmspolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/fms/fmspolicy"
	fsxlustrefilesystem "github.com/crossplane-contrib/provider-tf-aws/internal/controller/fsx/fsxlustrefilesystem"
	fsxwindowsfilesystem "github.com/crossplane-contrib/provider-tf-aws/internal/controller/fsx/fsxwindowsfilesystem"
	gameliftalias "github.com/crossplane-contrib/provider-tf-aws/internal/controller/gamelift/gameliftalias"
	gameliftbuild "github.com/crossplane-contrib/provider-tf-aws/internal/controller/gamelift/gameliftbuild"
	gameliftfleet "github.com/crossplane-contrib/provider-tf-aws/internal/controller/gamelift/gameliftfleet"
	gameliftgamesessionqueue "github.com/crossplane-contrib/provider-tf-aws/internal/controller/gamelift/gameliftgamesessionqueue"
	glaciervault "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glacier/glaciervault"
	glaciervaultlock "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glacier/glaciervaultlock"
	globalacceleratoraccelerator "github.com/crossplane-contrib/provider-tf-aws/internal/controller/globalaccelerator/globalacceleratoraccelerator"
	globalacceleratorendpointgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/globalaccelerator/globalacceleratorendpointgroup"
	globalacceleratorlistener "github.com/crossplane-contrib/provider-tf-aws/internal/controller/globalaccelerator/globalacceleratorlistener"
	gluecatalogdatabase "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/gluecatalogdatabase"
	gluecatalogtable "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/gluecatalogtable"
	glueclassifier "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/glueclassifier"
	glueconnection "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/glueconnection"
	gluecrawler "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/gluecrawler"
	gluedatacatalogencryptionsettings "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/gluedatacatalogencryptionsettings"
	gluedevendpoint "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/gluedevendpoint"
	gluejob "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/gluejob"
	gluemltransform "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/gluemltransform"
	gluepartition "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/gluepartition"
	glueregistry "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/glueregistry"
	glueresourcepolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/glueresourcepolicy"
	glueschema "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/glueschema"
	gluesecurityconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/gluesecurityconfiguration"
	gluetrigger "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/gluetrigger"
	glueuserdefinedfunction "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/glueuserdefinedfunction"
	glueworkflow "github.com/crossplane-contrib/provider-tf-aws/internal/controller/glue/glueworkflow"
	guarddutydetector "github.com/crossplane-contrib/provider-tf-aws/internal/controller/guardduty/guarddutydetector"
	guarddutyfilter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/guardduty/guarddutyfilter"
	guarddutyinviteaccepter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/guardduty/guarddutyinviteaccepter"
	guarddutyipset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/guardduty/guarddutyipset"
	guarddutymember "github.com/crossplane-contrib/provider-tf-aws/internal/controller/guardduty/guarddutymember"
	guarddutyorganizationadminaccount "github.com/crossplane-contrib/provider-tf-aws/internal/controller/guardduty/guarddutyorganizationadminaccount"
	guarddutyorganizationconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/guardduty/guarddutyorganizationconfiguration"
	guarddutypublishingdestination "github.com/crossplane-contrib/provider-tf-aws/internal/controller/guardduty/guarddutypublishingdestination"
	guarddutythreatintelset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/guardduty/guarddutythreatintelset"
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
	imagebuildercomponent "github.com/crossplane-contrib/provider-tf-aws/internal/controller/imagebuilder/imagebuildercomponent"
	imagebuilderdistributionconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/imagebuilder/imagebuilderdistributionconfiguration"
	imagebuilderimage "github.com/crossplane-contrib/provider-tf-aws/internal/controller/imagebuilder/imagebuilderimage"
	imagebuilderimagepipeline "github.com/crossplane-contrib/provider-tf-aws/internal/controller/imagebuilder/imagebuilderimagepipeline"
	imagebuilderimagerecipe "github.com/crossplane-contrib/provider-tf-aws/internal/controller/imagebuilder/imagebuilderimagerecipe"
	imagebuilderinfrastructureconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/imagebuilder/imagebuilderinfrastructureconfiguration"
	inspectorassessmenttarget "github.com/crossplane-contrib/provider-tf-aws/internal/controller/inspector/inspectorassessmenttarget"
	inspectorassessmenttemplate "github.com/crossplane-contrib/provider-tf-aws/internal/controller/inspector/inspectorassessmenttemplate"
	inspectorresourcegroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/inspector/inspectorresourcegroup"
	instance "github.com/crossplane-contrib/provider-tf-aws/internal/controller/instance/instance"
	internetgateway "github.com/crossplane-contrib/provider-tf-aws/internal/controller/internet/internetgateway"
	iotcertificate "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iot/iotcertificate"
	iotpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iot/iotpolicy"
	iotpolicyattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iot/iotpolicyattachment"
	iotrolealias "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iot/iotrolealias"
	iotthing "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iot/iotthing"
	iotthingprincipalattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iot/iotthingprincipalattachment"
	iotthingtype "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iot/iotthingtype"
	iottopicrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/iot/iottopicrule"
	keypair "github.com/crossplane-contrib/provider-tf-aws/internal/controller/key/keypair"
	kinesisfirehosedeliverystream "github.com/crossplane-contrib/provider-tf-aws/internal/controller/kinesis/kinesisfirehosedeliverystream"
	kinesisstream "github.com/crossplane-contrib/provider-tf-aws/internal/controller/kinesis/kinesisstream"
	kinesisstreamconsumer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/kinesis/kinesisstreamconsumer"
	kinesisvideostream "github.com/crossplane-contrib/provider-tf-aws/internal/controller/kinesis/kinesisvideostream"
	kinesisanalyticsv2application "github.com/crossplane-contrib/provider-tf-aws/internal/controller/kinesisanalyticsv2/kinesisanalyticsv2application"
	kinesisanalyticsv2applicationsnapshot "github.com/crossplane-contrib/provider-tf-aws/internal/controller/kinesisanalyticsv2/kinesisanalyticsv2applicationsnapshot"
	kmsalias "github.com/crossplane-contrib/provider-tf-aws/internal/controller/kms/kmsalias"
	kmsciphertext "github.com/crossplane-contrib/provider-tf-aws/internal/controller/kms/kmsciphertext"
	kmsexternalkey "github.com/crossplane-contrib/provider-tf-aws/internal/controller/kms/kmsexternalkey"
	kmsgrant "github.com/crossplane-contrib/provider-tf-aws/internal/controller/kms/kmsgrant"
	kmskey "github.com/crossplane-contrib/provider-tf-aws/internal/controller/kms/kmskey"
	lakeformationdatalakesettings "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lakeformation/lakeformationdatalakesettings"
	lakeformationpermissions "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lakeformation/lakeformationpermissions"
	lakeformationresource "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lakeformation/lakeformationresource"
	lambdaalias "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lambda/lambdaalias"
	lambdacodesigningconfig "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lambda/lambdacodesigningconfig"
	lambdaeventsourcemapping "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lambda/lambdaeventsourcemapping"
	lambdafunction "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lambda/lambdafunction"
	lambdafunctioneventinvokeconfig "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lambda/lambdafunctioneventinvokeconfig"
	lambdalayerversion "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lambda/lambdalayerversion"
	lambdapermission "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lambda/lambdapermission"
	lambdaprovisionedconcurrencyconfig "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lambda/lambdaprovisionedconcurrencyconfig"
	launchconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/launch/launchconfiguration"
	launchtemplate "github.com/crossplane-contrib/provider-tf-aws/internal/controller/launch/launchtemplate"
	lb "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lb/lb"
	lbcookiestickinesspolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lb/lbcookiestickinesspolicy"
	lblistener "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lb/lblistener"
	lblistenercertificate "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lb/lblistenercertificate"
	lblistenerrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lb/lblistenerrule"
	lbsslnegotiationpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lb/lbsslnegotiationpolicy"
	lbtargetgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lb/lbtargetgroup"
	lbtargetgroupattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lb/lbtargetgroupattachment"
	lexbot "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lex/lexbot"
	lexbotalias "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lex/lexbotalias"
	lexintent "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lex/lexintent"
	lexslottype "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lex/lexslottype"
	licensemanagerassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/licensemanager/licensemanagerassociation"
	licensemanagerlicenseconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/licensemanager/licensemanagerlicenseconfiguration"
	lightsaildomain "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lightsail/lightsaildomain"
	lightsailinstance "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lightsail/lightsailinstance"
	lightsailinstancepublicports "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lightsail/lightsailinstancepublicports"
	lightsailkeypair "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lightsail/lightsailkeypair"
	lightsailstaticip "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lightsail/lightsailstaticip"
	lightsailstaticipattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/lightsail/lightsailstaticipattachment"
	loadbalancerbackendserverpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/load/loadbalancerbackendserverpolicy"
	loadbalancerlistenerpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/load/loadbalancerlistenerpolicy"
	loadbalancerpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/load/loadbalancerpolicy"
	maciememberaccountassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/macie/maciememberaccountassociation"
	macies3bucketassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/macie/macies3bucketassociation"
	macie2account "github.com/crossplane-contrib/provider-tf-aws/internal/controller/macie2/macie2account"
	macie2classificationjob "github.com/crossplane-contrib/provider-tf-aws/internal/controller/macie2/macie2classificationjob"
	macie2customdataidentifier "github.com/crossplane-contrib/provider-tf-aws/internal/controller/macie2/macie2customdataidentifier"
	macie2findingsfilter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/macie2/macie2findingsfilter"
	macie2invitationaccepter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/macie2/macie2invitationaccepter"
	macie2member "github.com/crossplane-contrib/provider-tf-aws/internal/controller/macie2/macie2member"
	macie2organizationadminaccount "github.com/crossplane-contrib/provider-tf-aws/internal/controller/macie2/macie2organizationadminaccount"
	mainroutetableassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/main/mainroutetableassociation"
	mediaconvertqueue "github.com/crossplane-contrib/provider-tf-aws/internal/controller/media/mediaconvertqueue"
	mediapackagechannel "github.com/crossplane-contrib/provider-tf-aws/internal/controller/media/mediapackagechannel"
	mediastorecontainer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/media/mediastorecontainer"
	mediastorecontainerpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/media/mediastorecontainerpolicy"
	mqbroker "github.com/crossplane-contrib/provider-tf-aws/internal/controller/mq/mqbroker"
	mqconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/mq/mqconfiguration"
	mskcluster "github.com/crossplane-contrib/provider-tf-aws/internal/controller/msk/mskcluster"
	mskconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/msk/mskconfiguration"
	mskscramsecretassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/msk/mskscramsecretassociation"
	mwaaenvironment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/mwaa/mwaaenvironment"
	natgateway "github.com/crossplane-contrib/provider-tf-aws/internal/controller/nat/natgateway"
	neptunecluster "github.com/crossplane-contrib/provider-tf-aws/internal/controller/neptune/neptunecluster"
	neptuneclusterendpoint "github.com/crossplane-contrib/provider-tf-aws/internal/controller/neptune/neptuneclusterendpoint"
	neptuneclusterinstance "github.com/crossplane-contrib/provider-tf-aws/internal/controller/neptune/neptuneclusterinstance"
	neptuneclusterparametergroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/neptune/neptuneclusterparametergroup"
	neptuneclustersnapshot "github.com/crossplane-contrib/provider-tf-aws/internal/controller/neptune/neptuneclustersnapshot"
	neptuneeventsubscription "github.com/crossplane-contrib/provider-tf-aws/internal/controller/neptune/neptuneeventsubscription"
	neptuneparametergroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/neptune/neptuneparametergroup"
	neptunesubnetgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/neptune/neptunesubnetgroup"
	networkacl "github.com/crossplane-contrib/provider-tf-aws/internal/controller/network/networkacl"
	networkaclrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/network/networkaclrule"
	networkinterface "github.com/crossplane-contrib/provider-tf-aws/internal/controller/network/networkinterface"
	networkinterfaceattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/network/networkinterfaceattachment"
	networkinterfacesgattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/network/networkinterfacesgattachment"
	networkfirewallfirewall "github.com/crossplane-contrib/provider-tf-aws/internal/controller/networkfirewall/networkfirewallfirewall"
	networkfirewallfirewallpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/networkfirewall/networkfirewallfirewallpolicy"
	networkfirewallloggingconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/networkfirewall/networkfirewallloggingconfiguration"
	networkfirewallresourcepolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/networkfirewall/networkfirewallresourcepolicy"
	networkfirewallrulegroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/networkfirewall/networkfirewallrulegroup"
	opsworksapplication "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/opsworksapplication"
	opsworkscustomlayer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/opsworkscustomlayer"
	opsworksganglialayer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/opsworksganglialayer"
	opsworkshaproxylayer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/opsworkshaproxylayer"
	opsworksinstance "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/opsworksinstance"
	opsworksjavaapplayer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/opsworksjavaapplayer"
	opsworksmemcachedlayer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/opsworksmemcachedlayer"
	opsworksmysqllayer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/opsworksmysqllayer"
	opsworksnodejsapplayer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/opsworksnodejsapplayer"
	opsworkspermission "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/opsworkspermission"
	opsworksphpapplayer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/opsworksphpapplayer"
	opsworksrailsapplayer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/opsworksrailsapplayer"
	opsworksrdsdbinstance "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/opsworksrdsdbinstance"
	opsworksstack "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/opsworksstack"
	opsworksstaticweblayer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/opsworksstaticweblayer"
	opsworksuserprofile "github.com/crossplane-contrib/provider-tf-aws/internal/controller/opsworks/opsworksuserprofile"
	organizationsaccount "github.com/crossplane-contrib/provider-tf-aws/internal/controller/organizations/organizationsaccount"
	organizationsdelegatedadministrator "github.com/crossplane-contrib/provider-tf-aws/internal/controller/organizations/organizationsdelegatedadministrator"
	organizationsorganization "github.com/crossplane-contrib/provider-tf-aws/internal/controller/organizations/organizationsorganization"
	organizationsorganizationalunit "github.com/crossplane-contrib/provider-tf-aws/internal/controller/organizations/organizationsorganizationalunit"
	organizationspolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/organizations/organizationspolicy"
	organizationspolicyattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/organizations/organizationspolicyattachment"
	pinpointadmchannel "github.com/crossplane-contrib/provider-tf-aws/internal/controller/pinpoint/pinpointadmchannel"
	pinpointapnschannel "github.com/crossplane-contrib/provider-tf-aws/internal/controller/pinpoint/pinpointapnschannel"
	pinpointapnssandboxchannel "github.com/crossplane-contrib/provider-tf-aws/internal/controller/pinpoint/pinpointapnssandboxchannel"
	pinpointapnsvoipchannel "github.com/crossplane-contrib/provider-tf-aws/internal/controller/pinpoint/pinpointapnsvoipchannel"
	pinpointapnsvoipsandboxchannel "github.com/crossplane-contrib/provider-tf-aws/internal/controller/pinpoint/pinpointapnsvoipsandboxchannel"
	pinpointapp "github.com/crossplane-contrib/provider-tf-aws/internal/controller/pinpoint/pinpointapp"
	pinpointbaiduchannel "github.com/crossplane-contrib/provider-tf-aws/internal/controller/pinpoint/pinpointbaiduchannel"
	pinpointemailchannel "github.com/crossplane-contrib/provider-tf-aws/internal/controller/pinpoint/pinpointemailchannel"
	pinpointeventstream "github.com/crossplane-contrib/provider-tf-aws/internal/controller/pinpoint/pinpointeventstream"
	pinpointgcmchannel "github.com/crossplane-contrib/provider-tf-aws/internal/controller/pinpoint/pinpointgcmchannel"
	pinpointsmschannel "github.com/crossplane-contrib/provider-tf-aws/internal/controller/pinpoint/pinpointsmschannel"
	placementgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/placement/placementgroup"
	prometheusworkspace "github.com/crossplane-contrib/provider-tf-aws/internal/controller/prometheus/prometheusworkspace"
	proxyprotocolpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/proxy/proxyprotocolpolicy"
	qldbledger "github.com/crossplane-contrib/provider-tf-aws/internal/controller/qldb/qldbledger"
	quicksightgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/quicksight/quicksightgroup"
	quicksightuser "github.com/crossplane-contrib/provider-tf-aws/internal/controller/quicksight/quicksightuser"
	ramprincipalassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ram/ramprincipalassociation"
	ramresourceassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ram/ramresourceassociation"
	ramresourceshare "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ram/ramresourceshare"
	ramresourceshareaccepter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ram/ramresourceshareaccepter"
	rdscluster "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/rdscluster"
	rdsclusterendpoint "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/rdsclusterendpoint"
	rdsclusterinstance "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/rdsclusterinstance"
	rdsclusterparametergroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/rdsclusterparametergroup"
	rdsclusterroleassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/rdsclusterroleassociation"
	rdsglobalcluster "github.com/crossplane-contrib/provider-tf-aws/internal/controller/rds/rdsglobalcluster"
	redshiftcluster "github.com/crossplane-contrib/provider-tf-aws/internal/controller/redshift/redshiftcluster"
	redshifteventsubscription "github.com/crossplane-contrib/provider-tf-aws/internal/controller/redshift/redshifteventsubscription"
	redshiftparametergroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/redshift/redshiftparametergroup"
	redshiftsecuritygroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/redshift/redshiftsecuritygroup"
	redshiftsnapshotcopygrant "github.com/crossplane-contrib/provider-tf-aws/internal/controller/redshift/redshiftsnapshotcopygrant"
	redshiftsnapshotschedule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/redshift/redshiftsnapshotschedule"
	redshiftsnapshotscheduleassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/redshift/redshiftsnapshotscheduleassociation"
	redshiftsubnetgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/redshift/redshiftsubnetgroup"
	resourcegroupsgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/resourcegroups/resourcegroupsgroup"
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
	s3controlbucket "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3control/s3controlbucket"
	s3controlbucketlifecycleconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3control/s3controlbucketlifecycleconfiguration"
	s3controlbucketpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3control/s3controlbucketpolicy"
	s3outpostsendpoint "github.com/crossplane-contrib/provider-tf-aws/internal/controller/s3outposts/s3outpostsendpoint"
	sagemakerapp "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/sagemakerapp"
	sagemakerappimageconfig "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/sagemakerappimageconfig"
	sagemakercoderepository "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/sagemakercoderepository"
	sagemakerdomain "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/sagemakerdomain"
	sagemakerendpoint "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/sagemakerendpoint"
	sagemakerendpointconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/sagemakerendpointconfiguration"
	sagemakerfeaturegroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/sagemakerfeaturegroup"
	sagemakerimage "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/sagemakerimage"
	sagemakerimageversion "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/sagemakerimageversion"
	sagemakermodel "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/sagemakermodel"
	sagemakermodelpackagegroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/sagemakermodelpackagegroup"
	sagemakernotebookinstance "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/sagemakernotebookinstance"
	sagemakernotebookinstancelifecycleconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/sagemakernotebookinstancelifecycleconfiguration"
	sagemakeruserprofile "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/sagemakeruserprofile"
	sagemakerworkforce "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/sagemakerworkforce"
	sagemakerworkteam "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sagemaker/sagemakerworkteam"
	schemasdiscoverer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/schemas/schemasdiscoverer"
	schemasregistry "github.com/crossplane-contrib/provider-tf-aws/internal/controller/schemas/schemasregistry"
	schemasschema "github.com/crossplane-contrib/provider-tf-aws/internal/controller/schemas/schemasschema"
	secretsmanagersecret "github.com/crossplane-contrib/provider-tf-aws/internal/controller/secretsmanager/secretsmanagersecret"
	secretsmanagersecretpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/secretsmanager/secretsmanagersecretpolicy"
	secretsmanagersecretrotation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/secretsmanager/secretsmanagersecretrotation"
	secretsmanagersecretversion "github.com/crossplane-contrib/provider-tf-aws/internal/controller/secretsmanager/secretsmanagersecretversion"
	securitygroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/security/securitygroup"
	securitygrouprule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/security/securitygrouprule"
	securityhubactiontarget "github.com/crossplane-contrib/provider-tf-aws/internal/controller/securityhub/securityhubactiontarget"
	securityhubinsight "github.com/crossplane-contrib/provider-tf-aws/internal/controller/securityhub/securityhubinsight"
	securityhubinviteaccepter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/securityhub/securityhubinviteaccepter"
	securityhubmember "github.com/crossplane-contrib/provider-tf-aws/internal/controller/securityhub/securityhubmember"
	securityhuborganizationadminaccount "github.com/crossplane-contrib/provider-tf-aws/internal/controller/securityhub/securityhuborganizationadminaccount"
	securityhuborganizationconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/securityhub/securityhuborganizationconfiguration"
	securityhubproductsubscription "github.com/crossplane-contrib/provider-tf-aws/internal/controller/securityhub/securityhubproductsubscription"
	securityhubstandardscontrol "github.com/crossplane-contrib/provider-tf-aws/internal/controller/securityhub/securityhubstandardscontrol"
	securityhubstandardssubscription "github.com/crossplane-contrib/provider-tf-aws/internal/controller/securityhub/securityhubstandardssubscription"
	serverlessapplicationrepositorycloudformationstack "github.com/crossplane-contrib/provider-tf-aws/internal/controller/serverlessapplicationrepository/serverlessapplicationrepositorycloudformationstack"
	servicediscoveryhttpnamespace "github.com/crossplane-contrib/provider-tf-aws/internal/controller/service/servicediscoveryhttpnamespace"
	servicediscoveryprivatednsnamespace "github.com/crossplane-contrib/provider-tf-aws/internal/controller/service/servicediscoveryprivatednsnamespace"
	servicediscoverypublicdnsnamespace "github.com/crossplane-contrib/provider-tf-aws/internal/controller/service/servicediscoverypublicdnsnamespace"
	servicediscoveryservice "github.com/crossplane-contrib/provider-tf-aws/internal/controller/service/servicediscoveryservice"
	servicecatalogbudgetresourceassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/servicecatalog/servicecatalogbudgetresourceassociation"
	servicecatalogconstraint "github.com/crossplane-contrib/provider-tf-aws/internal/controller/servicecatalog/servicecatalogconstraint"
	servicecatalogorganizationsaccess "github.com/crossplane-contrib/provider-tf-aws/internal/controller/servicecatalog/servicecatalogorganizationsaccess"
	servicecatalogportfolio "github.com/crossplane-contrib/provider-tf-aws/internal/controller/servicecatalog/servicecatalogportfolio"
	servicecatalogportfolioshare "github.com/crossplane-contrib/provider-tf-aws/internal/controller/servicecatalog/servicecatalogportfolioshare"
	servicecatalogprincipalportfolioassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/servicecatalog/servicecatalogprincipalportfolioassociation"
	servicecatalogproduct "github.com/crossplane-contrib/provider-tf-aws/internal/controller/servicecatalog/servicecatalogproduct"
	servicecatalogproductportfolioassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/servicecatalog/servicecatalogproductportfolioassociation"
	servicecatalogprovisionedproduct "github.com/crossplane-contrib/provider-tf-aws/internal/controller/servicecatalog/servicecatalogprovisionedproduct"
	servicecatalogprovisioningartifact "github.com/crossplane-contrib/provider-tf-aws/internal/controller/servicecatalog/servicecatalogprovisioningartifact"
	servicecatalogserviceaction "github.com/crossplane-contrib/provider-tf-aws/internal/controller/servicecatalog/servicecatalogserviceaction"
	servicecatalogtagoption "github.com/crossplane-contrib/provider-tf-aws/internal/controller/servicecatalog/servicecatalogtagoption"
	servicecatalogtagoptionresourceassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/servicecatalog/servicecatalogtagoptionresourceassociation"
	servicequotasservicequota "github.com/crossplane-contrib/provider-tf-aws/internal/controller/servicequotas/servicequotasservicequota"
	sesactivereceiptruleset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ses/sesactivereceiptruleset"
	sesconfigurationset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ses/sesconfigurationset"
	sesdomaindkim "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ses/sesdomaindkim"
	sesdomainidentity "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ses/sesdomainidentity"
	sesdomainidentityverification "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ses/sesdomainidentityverification"
	sesdomainmailfrom "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ses/sesdomainmailfrom"
	sesemailidentity "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ses/sesemailidentity"
	seseventdestination "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ses/seseventdestination"
	sesidentitynotificationtopic "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ses/sesidentitynotificationtopic"
	sesidentitypolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ses/sesidentitypolicy"
	sesreceiptfilter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ses/sesreceiptfilter"
	sesreceiptrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ses/sesreceiptrule"
	sesreceiptruleset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ses/sesreceiptruleset"
	sestemplate "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ses/sestemplate"
	sfnactivity "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sfn/sfnactivity"
	sfnstatemachine "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sfn/sfnstatemachine"
	shieldprotection "github.com/crossplane-contrib/provider-tf-aws/internal/controller/shield/shieldprotection"
	signersigningjob "github.com/crossplane-contrib/provider-tf-aws/internal/controller/signer/signersigningjob"
	signersigningprofile "github.com/crossplane-contrib/provider-tf-aws/internal/controller/signer/signersigningprofile"
	signersigningprofilepermission "github.com/crossplane-contrib/provider-tf-aws/internal/controller/signer/signersigningprofilepermission"
	simpledbdomain "github.com/crossplane-contrib/provider-tf-aws/internal/controller/simpledb/simpledbdomain"
	snapshotcreatevolumepermission "github.com/crossplane-contrib/provider-tf-aws/internal/controller/snapshot/snapshotcreatevolumepermission"
	snsplatformapplication "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sns/snsplatformapplication"
	snssmspreferences "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sns/snssmspreferences"
	snstopic "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sns/snstopic"
	snstopicpolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sns/snstopicpolicy"
	snstopicsubscription "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sns/snstopicsubscription"
	spotdatafeedsubscription "github.com/crossplane-contrib/provider-tf-aws/internal/controller/spot/spotdatafeedsubscription"
	spotfleetrequest "github.com/crossplane-contrib/provider-tf-aws/internal/controller/spot/spotfleetrequest"
	spotinstancerequest "github.com/crossplane-contrib/provider-tf-aws/internal/controller/spot/spotinstancerequest"
	sqsqueue "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sqs/sqsqueue"
	sqsqueuepolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/sqs/sqsqueuepolicy"
	ssmactivation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ssm/ssmactivation"
	ssmassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ssm/ssmassociation"
	ssmdocument "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ssm/ssmdocument"
	ssmmaintenancewindow "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ssm/ssmmaintenancewindow"
	ssmmaintenancewindowtarget "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ssm/ssmmaintenancewindowtarget"
	ssmmaintenancewindowtask "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ssm/ssmmaintenancewindowtask"
	ssmparameter "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ssm/ssmparameter"
	ssmpatchbaseline "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ssm/ssmpatchbaseline"
	ssmpatchgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ssm/ssmpatchgroup"
	ssmresourcedatasync "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ssm/ssmresourcedatasync"
	ssoadminaccountassignment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ssoadmin/ssoadminaccountassignment"
	ssoadminmanagedpolicyattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ssoadmin/ssoadminmanagedpolicyattachment"
	ssoadminpermissionset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ssoadmin/ssoadminpermissionset"
	ssoadminpermissionsetinlinepolicy "github.com/crossplane-contrib/provider-tf-aws/internal/controller/ssoadmin/ssoadminpermissionsetinlinepolicy"
	storagegatewaycache "github.com/crossplane-contrib/provider-tf-aws/internal/controller/storagegateway/storagegatewaycache"
	storagegatewaycachediscsivolume "github.com/crossplane-contrib/provider-tf-aws/internal/controller/storagegateway/storagegatewaycachediscsivolume"
	storagegatewayfilesystemassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/storagegateway/storagegatewayfilesystemassociation"
	storagegatewaygateway "github.com/crossplane-contrib/provider-tf-aws/internal/controller/storagegateway/storagegatewaygateway"
	storagegatewaynfsfileshare "github.com/crossplane-contrib/provider-tf-aws/internal/controller/storagegateway/storagegatewaynfsfileshare"
	storagegatewaysmbfileshare "github.com/crossplane-contrib/provider-tf-aws/internal/controller/storagegateway/storagegatewaysmbfileshare"
	storagegatewaystorediscsivolume "github.com/crossplane-contrib/provider-tf-aws/internal/controller/storagegateway/storagegatewaystorediscsivolume"
	storagegatewaytapepool "github.com/crossplane-contrib/provider-tf-aws/internal/controller/storagegateway/storagegatewaytapepool"
	storagegatewayuploadbuffer "github.com/crossplane-contrib/provider-tf-aws/internal/controller/storagegateway/storagegatewayuploadbuffer"
	storagegatewayworkingstorage "github.com/crossplane-contrib/provider-tf-aws/internal/controller/storagegateway/storagegatewayworkingstorage"
	subnet "github.com/crossplane-contrib/provider-tf-aws/internal/controller/subnet/subnet"
	swfdomain "github.com/crossplane-contrib/provider-tf-aws/internal/controller/swf/swfdomain"
	syntheticscanary "github.com/crossplane-contrib/provider-tf-aws/internal/controller/synthetics/syntheticscanary"
	tfaws "github.com/crossplane-contrib/provider-tf-aws/internal/controller/tfaws"
	timestreamwritedatabase "github.com/crossplane-contrib/provider-tf-aws/internal/controller/timestreamwrite/timestreamwritedatabase"
	timestreamwritetable "github.com/crossplane-contrib/provider-tf-aws/internal/controller/timestreamwrite/timestreamwritetable"
	transferserver "github.com/crossplane-contrib/provider-tf-aws/internal/controller/transfer/transferserver"
	transfersshkey "github.com/crossplane-contrib/provider-tf-aws/internal/controller/transfer/transfersshkey"
	transferuser "github.com/crossplane-contrib/provider-tf-aws/internal/controller/transfer/transferuser"
	volumeattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/volume/volumeattachment"
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
	vpnconnection "github.com/crossplane-contrib/provider-tf-aws/internal/controller/vpn/vpnconnection"
	vpnconnectionroute "github.com/crossplane-contrib/provider-tf-aws/internal/controller/vpn/vpnconnectionroute"
	vpngateway "github.com/crossplane-contrib/provider-tf-aws/internal/controller/vpn/vpngateway"
	vpngatewayattachment "github.com/crossplane-contrib/provider-tf-aws/internal/controller/vpn/vpngatewayattachment"
	vpngatewayroutepropagation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/vpn/vpngatewayroutepropagation"
	wafbytematchset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/waf/wafbytematchset"
	wafgeomatchset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/waf/wafgeomatchset"
	wafipset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/waf/wafipset"
	wafratebasedrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/waf/wafratebasedrule"
	wafregexmatchset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/waf/wafregexmatchset"
	wafregexpatternset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/waf/wafregexpatternset"
	wafrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/waf/wafrule"
	wafsizeconstraintset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/waf/wafsizeconstraintset"
	wafsqlinjectionmatchset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/waf/wafsqlinjectionmatchset"
	wafwebacl "github.com/crossplane-contrib/provider-tf-aws/internal/controller/waf/wafwebacl"
	wafxssmatchset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/waf/wafxssmatchset"
	wafregionalbytematchset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafregional/wafregionalbytematchset"
	wafregionalgeomatchset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafregional/wafregionalgeomatchset"
	wafregionalipset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafregional/wafregionalipset"
	wafregionalratebasedrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafregional/wafregionalratebasedrule"
	wafregionalregexmatchset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafregional/wafregionalregexmatchset"
	wafregionalregexpatternset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafregional/wafregionalregexpatternset"
	wafregionalrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafregional/wafregionalrule"
	wafregionalsizeconstraintset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafregional/wafregionalsizeconstraintset"
	wafregionalsqlinjectionmatchset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafregional/wafregionalsqlinjectionmatchset"
	wafregionalwebacl "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafregional/wafregionalwebacl"
	wafregionalwebaclassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafregional/wafregionalwebaclassociation"
	wafregionalxssmatchset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafregional/wafregionalxssmatchset"
	wafv2ipset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafv2/wafv2ipset"
	wafv2regexpatternset "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafv2/wafv2regexpatternset"
	wafv2rulegroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafv2/wafv2rulegroup"
	wafv2webacl "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafv2/wafv2webacl"
	wafv2webaclassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafv2/wafv2webaclassociation"
	wafv2webaclloggingconfiguration "github.com/crossplane-contrib/provider-tf-aws/internal/controller/wafv2/wafv2webaclloggingconfiguration"
	worklinkfleet "github.com/crossplane-contrib/provider-tf-aws/internal/controller/worklink/worklinkfleet"
	worklinkwebsitecertificateauthorityassociation "github.com/crossplane-contrib/provider-tf-aws/internal/controller/worklink/worklinkwebsitecertificateauthorityassociation"
	workspacesdirectory "github.com/crossplane-contrib/provider-tf-aws/internal/controller/workspaces/workspacesdirectory"
	workspacesipgroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/workspaces/workspacesipgroup"
	workspacesworkspace "github.com/crossplane-contrib/provider-tf-aws/internal/controller/workspaces/workspacesworkspace"
	xrayencryptionconfig "github.com/crossplane-contrib/provider-tf-aws/internal/controller/xray/xrayencryptionconfig"
	xraygroup "github.com/crossplane-contrib/provider-tf-aws/internal/controller/xray/xraygroup"
	xraysamplingrule "github.com/crossplane-contrib/provider-tf-aws/internal/controller/xray/xraysamplingrule"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, l logging.Logger, wl workqueue.RateLimiter) error {
	for _, setup := range []func(ctrl.Manager, logging.Logger, workqueue.RateLimiter) error{
		accessanalyzeranalyzer.Setup,
		acmcertificate.Setup,
		acmcertificatevalidation.Setup,
		acmpcacertificate.Setup,
		acmpcacertificateauthority.Setup,
		acmpcacertificateauthoritycertificate.Setup,
		alb.Setup,
		alblistener.Setup,
		alblistenercertificate.Setup,
		alblistenerrule.Setup,
		albtargetgroup.Setup,
		albtargetgroupattachment.Setup,
		ami.Setup,
		amicopy.Setup,
		amifrominstance.Setup,
		amilaunchpermission.Setup,
		amplifyapp.Setup,
		amplifybackendenvironment.Setup,
		amplifybranch.Setup,
		amplifydomainassociation.Setup,
		amplifywebhook.Setup,
		apigatewayaccount.Setup,
		apigatewayapikey.Setup,
		apigatewayauthorizer.Setup,
		apigatewaybasepathmapping.Setup,
		apigatewayclientcertificate.Setup,
		apigatewaydeployment.Setup,
		apigatewaydocumentationpart.Setup,
		apigatewaydocumentationversion.Setup,
		apigatewaydomainname.Setup,
		apigatewaygatewayresponse.Setup,
		apigatewayintegration.Setup,
		apigatewayintegrationresponse.Setup,
		apigatewaymethod.Setup,
		apigatewaymethodresponse.Setup,
		apigatewaymethodsettings.Setup,
		apigatewaymodel.Setup,
		apigatewayrequestvalidator.Setup,
		apigatewayresource.Setup,
		apigatewayrestapi.Setup,
		apigatewayrestapipolicy.Setup,
		apigatewaystage.Setup,
		apigatewayusageplan.Setup,
		apigatewayusageplankey.Setup,
		apigatewayv2api.Setup,
		apigatewayv2apimapping.Setup,
		apigatewayv2authorizer.Setup,
		apigatewayv2deployment.Setup,
		apigatewayv2domainname.Setup,
		apigatewayv2integration.Setup,
		apigatewayv2integrationresponse.Setup,
		apigatewayv2model.Setup,
		apigatewayv2route.Setup,
		apigatewayv2routeresponse.Setup,
		apigatewayv2stage.Setup,
		apigatewayv2vpclink.Setup,
		apigatewayvpclink.Setup,
		appautoscalingpolicy.Setup,
		appautoscalingscheduledaction.Setup,
		appautoscalingtarget.Setup,
		appconfigapplication.Setup,
		appconfigconfigurationprofile.Setup,
		appconfigdeployment.Setup,
		appconfigdeploymentstrategy.Setup,
		appconfigenvironment.Setup,
		appconfighostedconfigurationversion.Setup,
		appcookiestickinesspolicy.Setup,
		appmeshgatewayroute.Setup,
		appmeshmesh.Setup,
		appmeshroute.Setup,
		appmeshvirtualgateway.Setup,
		appmeshvirtualnode.Setup,
		appmeshvirtualrouter.Setup,
		appmeshvirtualservice.Setup,
		apprunnerautoscalingconfigurationversion.Setup,
		apprunnerconnection.Setup,
		apprunnercustomdomainassociation.Setup,
		apprunnerservice.Setup,
		appsyncapikey.Setup,
		appsyncdatasource.Setup,
		appsyncfunction.Setup,
		appsyncgraphqlapi.Setup,
		appsyncresolver.Setup,
		athenadatabase.Setup,
		athenanamedquery.Setup,
		athenaworkgroup.Setup,
		autoscalingattachment.Setup,
		autoscalinggroup.Setup,
		autoscalinglifecyclehook.Setup,
		autoscalingnotification.Setup,
		autoscalingplansscalingplan.Setup,
		autoscalingpolicy.Setup,
		autoscalingschedule.Setup,
		backupglobalsettings.Setup,
		backupplan.Setup,
		backupregionsettings.Setup,
		backupselection.Setup,
		backupvault.Setup,
		backupvaultnotifications.Setup,
		backupvaultpolicy.Setup,
		batchcomputeenvironment.Setup,
		batchjobdefinition.Setup,
		batchjobqueue.Setup,
		budgetsbudget.Setup,
		budgetsbudgetaction.Setup,
		chimevoiceconnector.Setup,
		cloud9environmentec2.Setup,
		cloudformationstack.Setup,
		cloudformationstackset.Setup,
		cloudformationstacksetinstance.Setup,
		cloudformationtype.Setup,
		cloudfrontcachepolicy.Setup,
		cloudfrontdistribution.Setup,
		cloudfrontfunction.Setup,
		cloudfrontkeygroup.Setup,
		cloudfrontmonitoringsubscription.Setup,
		cloudfrontoriginaccessidentity.Setup,
		cloudfrontoriginrequestpolicy.Setup,
		cloudfrontpublickey.Setup,
		cloudfrontrealtimelogconfig.Setup,
		cloudhsmv2cluster.Setup,
		cloudhsmv2hsm.Setup,
		cloudtrail.Setup,
		cloudwatchcompositealarm.Setup,
		cloudwatchdashboard.Setup,
		cloudwatcheventapidestination.Setup,
		cloudwatcheventarchive.Setup,
		cloudwatcheventbus.Setup,
		cloudwatcheventbuspolicy.Setup,
		cloudwatcheventconnection.Setup,
		cloudwatcheventpermission.Setup,
		cloudwatcheventrule.Setup,
		cloudwatcheventtarget.Setup,
		cloudwatchlogdestination.Setup,
		cloudwatchlogdestinationpolicy.Setup,
		cloudwatchloggroup.Setup,
		cloudwatchlogmetricfilter.Setup,
		cloudwatchlogresourcepolicy.Setup,
		cloudwatchlogstream.Setup,
		cloudwatchlogsubscriptionfilter.Setup,
		cloudwatchmetricalarm.Setup,
		cloudwatchmetricstream.Setup,
		cloudwatchquerydefinition.Setup,
		codeartifactdomain.Setup,
		codeartifactdomainpermissionspolicy.Setup,
		codeartifactrepository.Setup,
		codeartifactrepositorypermissionspolicy.Setup,
		codebuildproject.Setup,
		codebuildreportgroup.Setup,
		codebuildsourcecredential.Setup,
		codebuildwebhook.Setup,
		codecommitrepository.Setup,
		codecommittrigger.Setup,
		codedeployapp.Setup,
		codedeploydeploymentconfig.Setup,
		codedeploydeploymentgroup.Setup,
		codepipeline.Setup,
		codepipelinewebhook.Setup,
		codestarconnectionsconnection.Setup,
		codestarconnectionshost.Setup,
		codestarnotificationsnotificationrule.Setup,
		cognitoidentitypool.Setup,
		cognitoidentitypoolrolesattachment.Setup,
		cognitoidentityprovider.Setup,
		cognitoresourceserver.Setup,
		cognitousergroup.Setup,
		cognitouserpool.Setup,
		cognitouserpoolclient.Setup,
		cognitouserpooldomain.Setup,
		cognitouserpooluicustomization.Setup,
		configaggregateauthorization.Setup,
		configconfigrule.Setup,
		configconfigurationaggregator.Setup,
		configconfigurationrecorder.Setup,
		configconformancepack.Setup,
		configdeliverychannel.Setup,
		configorganizationconformancepack.Setup,
		configorganizationcustomrule.Setup,
		configorganizationmanagedrule.Setup,
		configremediationconfiguration.Setup,
		curreportdefinition.Setup,
		customergateway.Setup,
		datapipelinepipeline.Setup,
		datasyncagent.Setup,
		datasynclocationefs.Setup,
		datasynclocationfsxwindowsfilesystem.Setup,
		datasynclocationnfs.Setup,
		datasynclocations3.Setup,
		datasynclocationsmb.Setup,
		datasynctask.Setup,
		daxcluster.Setup,
		daxparametergroup.Setup,
		daxsubnetgroup.Setup,
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
		defaultnetworkacl.Setup,
		defaultroutetable.Setup,
		defaultsecuritygroup.Setup,
		defaultsubnet.Setup,
		defaultvpc.Setup,
		defaultvpcdhcpoptions.Setup,
		devicefarmproject.Setup,
		directoryserviceconditionalforwarder.Setup,
		directoryservicedirectory.Setup,
		directoryservicelogsubscription.Setup,
		dlmlifecyclepolicy.Setup,
		dmscertificate.Setup,
		dmsendpoint.Setup,
		dmseventsubscription.Setup,
		dmsreplicationinstance.Setup,
		dmsreplicationsubnetgroup.Setup,
		dmsreplicationtask.Setup,
		docdbcluster.Setup,
		docdbclusterinstance.Setup,
		docdbclusterparametergroup.Setup,
		docdbclustersnapshot.Setup,
		docdbsubnetgroup.Setup,
		dxbgppeer.Setup,
		dxconnection.Setup,
		dxconnectionassociation.Setup,
		dxgateway.Setup,
		dxgatewayassociation.Setup,
		dxgatewayassociationproposal.Setup,
		dxhostedprivatevirtualinterface.Setup,
		dxhostedprivatevirtualinterfaceaccepter.Setup,
		dxhostedpublicvirtualinterface.Setup,
		dxhostedpublicvirtualinterfaceaccepter.Setup,
		dxhostedtransitvirtualinterface.Setup,
		dxhostedtransitvirtualinterfaceaccepter.Setup,
		dxlag.Setup,
		dxprivatevirtualinterface.Setup,
		dxpublicvirtualinterface.Setup,
		dxtransitvirtualinterface.Setup,
		dynamodbglobaltable.Setup,
		dynamodbkinesisstreamingdestination.Setup,
		dynamodbtable.Setup,
		dynamodbtableitem.Setup,
		ebsdefaultkmskey.Setup,
		ebsencryptionbydefault.Setup,
		ebssnapshot.Setup,
		ebssnapshotcopy.Setup,
		ebssnapshotimport.Setup,
		ebsvolume.Setup,
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
		ecrlifecyclepolicy.Setup,
		ecrpublicrepository.Setup,
		ecrregistrypolicy.Setup,
		ecrreplicationconfiguration.Setup,
		ecrrepository.Setup,
		ecrrepositorypolicy.Setup,
		ecscapacityprovider.Setup,
		ecscluster.Setup,
		ecsservice.Setup,
		ecstaskdefinition.Setup,
		efsaccesspoint.Setup,
		efsbackuppolicy.Setup,
		efsfilesystem.Setup,
		efsfilesystempolicy.Setup,
		efsmounttarget.Setup,
		egressonlyinternetgateway.Setup,
		eip.Setup,
		eipassociation.Setup,
		eksaddon.Setup,
		ekscluster.Setup,
		eksfargateprofile.Setup,
		eksidentityproviderconfig.Setup,
		eksnodegroup.Setup,
		elasticachecluster.Setup,
		elasticacheglobalreplicationgroup.Setup,
		elasticacheparametergroup.Setup,
		elasticachereplicationgroup.Setup,
		elasticachesecuritygroup.Setup,
		elasticachesubnetgroup.Setup,
		elasticacheuser.Setup,
		elasticbeanstalkapplication.Setup,
		elasticbeanstalkapplicationversion.Setup,
		elasticbeanstalkconfigurationtemplate.Setup,
		elasticbeanstalkenvironment.Setup,
		elasticsearchdomain.Setup,
		elasticsearchdomainpolicy.Setup,
		elasticsearchdomainsamloptions.Setup,
		elastictranscoderpipeline.Setup,
		elastictranscoderpreset.Setup,
		elb.Setup,
		elbattachment.Setup,
		emrcluster.Setup,
		emrinstancefleet.Setup,
		emrinstancegroup.Setup,
		emrmanagedscalingpolicy.Setup,
		emrsecurityconfiguration.Setup,
		flowlog.Setup,
		fmsadminaccount.Setup,
		fmspolicy.Setup,
		fsxlustrefilesystem.Setup,
		fsxwindowsfilesystem.Setup,
		gameliftalias.Setup,
		gameliftbuild.Setup,
		gameliftfleet.Setup,
		gameliftgamesessionqueue.Setup,
		glaciervault.Setup,
		glaciervaultlock.Setup,
		globalacceleratoraccelerator.Setup,
		globalacceleratorendpointgroup.Setup,
		globalacceleratorlistener.Setup,
		gluecatalogdatabase.Setup,
		gluecatalogtable.Setup,
		glueclassifier.Setup,
		glueconnection.Setup,
		gluecrawler.Setup,
		gluedatacatalogencryptionsettings.Setup,
		gluedevendpoint.Setup,
		gluejob.Setup,
		gluemltransform.Setup,
		gluepartition.Setup,
		glueregistry.Setup,
		glueresourcepolicy.Setup,
		glueschema.Setup,
		gluesecurityconfiguration.Setup,
		gluetrigger.Setup,
		glueuserdefinedfunction.Setup,
		glueworkflow.Setup,
		guarddutydetector.Setup,
		guarddutyfilter.Setup,
		guarddutyinviteaccepter.Setup,
		guarddutyipset.Setup,
		guarddutymember.Setup,
		guarddutyorganizationadminaccount.Setup,
		guarddutyorganizationconfiguration.Setup,
		guarddutypublishingdestination.Setup,
		guarddutythreatintelset.Setup,
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
		imagebuildercomponent.Setup,
		imagebuilderdistributionconfiguration.Setup,
		imagebuilderimage.Setup,
		imagebuilderimagepipeline.Setup,
		imagebuilderimagerecipe.Setup,
		imagebuilderinfrastructureconfiguration.Setup,
		inspectorassessmenttarget.Setup,
		inspectorassessmenttemplate.Setup,
		inspectorresourcegroup.Setup,
		instance.Setup,
		internetgateway.Setup,
		iotcertificate.Setup,
		iotpolicy.Setup,
		iotpolicyattachment.Setup,
		iotrolealias.Setup,
		iotthing.Setup,
		iotthingprincipalattachment.Setup,
		iotthingtype.Setup,
		iottopicrule.Setup,
		keypair.Setup,
		kinesisanalyticsv2application.Setup,
		kinesisanalyticsv2applicationsnapshot.Setup,
		kinesisfirehosedeliverystream.Setup,
		kinesisstream.Setup,
		kinesisstreamconsumer.Setup,
		kinesisvideostream.Setup,
		kmsalias.Setup,
		kmsciphertext.Setup,
		kmsexternalkey.Setup,
		kmsgrant.Setup,
		kmskey.Setup,
		lakeformationdatalakesettings.Setup,
		lakeformationpermissions.Setup,
		lakeformationresource.Setup,
		lambdaalias.Setup,
		lambdacodesigningconfig.Setup,
		lambdaeventsourcemapping.Setup,
		lambdafunction.Setup,
		lambdafunctioneventinvokeconfig.Setup,
		lambdalayerversion.Setup,
		lambdapermission.Setup,
		lambdaprovisionedconcurrencyconfig.Setup,
		launchconfiguration.Setup,
		launchtemplate.Setup,
		lb.Setup,
		lbcookiestickinesspolicy.Setup,
		lblistener.Setup,
		lblistenercertificate.Setup,
		lblistenerrule.Setup,
		lbsslnegotiationpolicy.Setup,
		lbtargetgroup.Setup,
		lbtargetgroupattachment.Setup,
		lexbot.Setup,
		lexbotalias.Setup,
		lexintent.Setup,
		lexslottype.Setup,
		licensemanagerassociation.Setup,
		licensemanagerlicenseconfiguration.Setup,
		lightsaildomain.Setup,
		lightsailinstance.Setup,
		lightsailinstancepublicports.Setup,
		lightsailkeypair.Setup,
		lightsailstaticip.Setup,
		lightsailstaticipattachment.Setup,
		loadbalancerbackendserverpolicy.Setup,
		loadbalancerlistenerpolicy.Setup,
		loadbalancerpolicy.Setup,
		macie2account.Setup,
		macie2classificationjob.Setup,
		macie2customdataidentifier.Setup,
		macie2findingsfilter.Setup,
		macie2invitationaccepter.Setup,
		macie2member.Setup,
		macie2organizationadminaccount.Setup,
		maciememberaccountassociation.Setup,
		macies3bucketassociation.Setup,
		mainroutetableassociation.Setup,
		mediaconvertqueue.Setup,
		mediapackagechannel.Setup,
		mediastorecontainer.Setup,
		mediastorecontainerpolicy.Setup,
		mqbroker.Setup,
		mqconfiguration.Setup,
		mskcluster.Setup,
		mskconfiguration.Setup,
		mskscramsecretassociation.Setup,
		mwaaenvironment.Setup,
		natgateway.Setup,
		neptunecluster.Setup,
		neptuneclusterendpoint.Setup,
		neptuneclusterinstance.Setup,
		neptuneclusterparametergroup.Setup,
		neptuneclustersnapshot.Setup,
		neptuneeventsubscription.Setup,
		neptuneparametergroup.Setup,
		neptunesubnetgroup.Setup,
		networkacl.Setup,
		networkaclrule.Setup,
		networkfirewallfirewall.Setup,
		networkfirewallfirewallpolicy.Setup,
		networkfirewallloggingconfiguration.Setup,
		networkfirewallresourcepolicy.Setup,
		networkfirewallrulegroup.Setup,
		networkinterface.Setup,
		networkinterfaceattachment.Setup,
		networkinterfacesgattachment.Setup,
		opsworksapplication.Setup,
		opsworkscustomlayer.Setup,
		opsworksganglialayer.Setup,
		opsworkshaproxylayer.Setup,
		opsworksinstance.Setup,
		opsworksjavaapplayer.Setup,
		opsworksmemcachedlayer.Setup,
		opsworksmysqllayer.Setup,
		opsworksnodejsapplayer.Setup,
		opsworkspermission.Setup,
		opsworksphpapplayer.Setup,
		opsworksrailsapplayer.Setup,
		opsworksrdsdbinstance.Setup,
		opsworksstack.Setup,
		opsworksstaticweblayer.Setup,
		opsworksuserprofile.Setup,
		organizationsaccount.Setup,
		organizationsdelegatedadministrator.Setup,
		organizationsorganization.Setup,
		organizationsorganizationalunit.Setup,
		organizationspolicy.Setup,
		organizationspolicyattachment.Setup,
		pinpointadmchannel.Setup,
		pinpointapnschannel.Setup,
		pinpointapnssandboxchannel.Setup,
		pinpointapnsvoipchannel.Setup,
		pinpointapnsvoipsandboxchannel.Setup,
		pinpointapp.Setup,
		pinpointbaiduchannel.Setup,
		pinpointemailchannel.Setup,
		pinpointeventstream.Setup,
		pinpointgcmchannel.Setup,
		pinpointsmschannel.Setup,
		placementgroup.Setup,
		prometheusworkspace.Setup,
		proxyprotocolpolicy.Setup,
		qldbledger.Setup,
		quicksightgroup.Setup,
		quicksightuser.Setup,
		ramprincipalassociation.Setup,
		ramresourceassociation.Setup,
		ramresourceshare.Setup,
		ramresourceshareaccepter.Setup,
		rdscluster.Setup,
		rdsclusterendpoint.Setup,
		rdsclusterinstance.Setup,
		rdsclusterparametergroup.Setup,
		rdsclusterroleassociation.Setup,
		rdsglobalcluster.Setup,
		redshiftcluster.Setup,
		redshifteventsubscription.Setup,
		redshiftparametergroup.Setup,
		redshiftsecuritygroup.Setup,
		redshiftsnapshotcopygrant.Setup,
		redshiftsnapshotschedule.Setup,
		redshiftsnapshotscheduleassociation.Setup,
		redshiftsubnetgroup.Setup,
		resourcegroupsgroup.Setup,
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
		s3controlbucket.Setup,
		s3controlbucketlifecycleconfiguration.Setup,
		s3controlbucketpolicy.Setup,
		s3objectcopy.Setup,
		s3outpostsendpoint.Setup,
		sagemakerapp.Setup,
		sagemakerappimageconfig.Setup,
		sagemakercoderepository.Setup,
		sagemakerdomain.Setup,
		sagemakerendpoint.Setup,
		sagemakerendpointconfiguration.Setup,
		sagemakerfeaturegroup.Setup,
		sagemakerimage.Setup,
		sagemakerimageversion.Setup,
		sagemakermodel.Setup,
		sagemakermodelpackagegroup.Setup,
		sagemakernotebookinstance.Setup,
		sagemakernotebookinstancelifecycleconfiguration.Setup,
		sagemakeruserprofile.Setup,
		sagemakerworkforce.Setup,
		sagemakerworkteam.Setup,
		schemasdiscoverer.Setup,
		schemasregistry.Setup,
		schemasschema.Setup,
		secretsmanagersecret.Setup,
		secretsmanagersecretpolicy.Setup,
		secretsmanagersecretrotation.Setup,
		secretsmanagersecretversion.Setup,
		securitygroup.Setup,
		securitygrouprule.Setup,
		securityhubactiontarget.Setup,
		securityhubinsight.Setup,
		securityhubinviteaccepter.Setup,
		securityhubmember.Setup,
		securityhuborganizationadminaccount.Setup,
		securityhuborganizationconfiguration.Setup,
		securityhubproductsubscription.Setup,
		securityhubstandardscontrol.Setup,
		securityhubstandardssubscription.Setup,
		serverlessapplicationrepositorycloudformationstack.Setup,
		servicecatalogbudgetresourceassociation.Setup,
		servicecatalogconstraint.Setup,
		servicecatalogorganizationsaccess.Setup,
		servicecatalogportfolio.Setup,
		servicecatalogportfolioshare.Setup,
		servicecatalogprincipalportfolioassociation.Setup,
		servicecatalogproduct.Setup,
		servicecatalogproductportfolioassociation.Setup,
		servicecatalogprovisionedproduct.Setup,
		servicecatalogprovisioningartifact.Setup,
		servicecatalogserviceaction.Setup,
		servicecatalogtagoption.Setup,
		servicecatalogtagoptionresourceassociation.Setup,
		servicediscoveryhttpnamespace.Setup,
		servicediscoveryprivatednsnamespace.Setup,
		servicediscoverypublicdnsnamespace.Setup,
		servicediscoveryservice.Setup,
		servicequotasservicequota.Setup,
		sesactivereceiptruleset.Setup,
		sesconfigurationset.Setup,
		sesdomaindkim.Setup,
		sesdomainidentity.Setup,
		sesdomainidentityverification.Setup,
		sesdomainmailfrom.Setup,
		sesemailidentity.Setup,
		seseventdestination.Setup,
		sesidentitynotificationtopic.Setup,
		sesidentitypolicy.Setup,
		sesreceiptfilter.Setup,
		sesreceiptrule.Setup,
		sesreceiptruleset.Setup,
		sestemplate.Setup,
		sfnactivity.Setup,
		sfnstatemachine.Setup,
		shieldprotection.Setup,
		signersigningjob.Setup,
		signersigningprofile.Setup,
		signersigningprofilepermission.Setup,
		simpledbdomain.Setup,
		snapshotcreatevolumepermission.Setup,
		snsplatformapplication.Setup,
		snssmspreferences.Setup,
		snstopic.Setup,
		snstopicpolicy.Setup,
		snstopicsubscription.Setup,
		spotdatafeedsubscription.Setup,
		spotfleetrequest.Setup,
		spotinstancerequest.Setup,
		sqsqueue.Setup,
		sqsqueuepolicy.Setup,
		ssmactivation.Setup,
		ssmassociation.Setup,
		ssmdocument.Setup,
		ssmmaintenancewindow.Setup,
		ssmmaintenancewindowtarget.Setup,
		ssmmaintenancewindowtask.Setup,
		ssmparameter.Setup,
		ssmpatchbaseline.Setup,
		ssmpatchgroup.Setup,
		ssmresourcedatasync.Setup,
		ssoadminaccountassignment.Setup,
		ssoadminmanagedpolicyattachment.Setup,
		ssoadminpermissionset.Setup,
		ssoadminpermissionsetinlinepolicy.Setup,
		storagegatewaycache.Setup,
		storagegatewaycachediscsivolume.Setup,
		storagegatewayfilesystemassociation.Setup,
		storagegatewaygateway.Setup,
		storagegatewaynfsfileshare.Setup,
		storagegatewaysmbfileshare.Setup,
		storagegatewaystorediscsivolume.Setup,
		storagegatewaytapepool.Setup,
		storagegatewayuploadbuffer.Setup,
		storagegatewayworkingstorage.Setup,
		subnet.Setup,
		swfdomain.Setup,
		syntheticscanary.Setup,
		tfaws.Setup,
		timestreamwritedatabase.Setup,
		timestreamwritetable.Setup,
		transferserver.Setup,
		transfersshkey.Setup,
		transferuser.Setup,
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
		vpcpeeringconnectionoptions.Setup,
		vpnconnection.Setup,
		vpnconnectionroute.Setup,
		vpngateway.Setup,
		vpngatewayattachment.Setup,
		vpngatewayroutepropagation.Setup,
		wafbytematchset.Setup,
		wafgeomatchset.Setup,
		wafipset.Setup,
		wafratebasedrule.Setup,
		wafregexmatchset.Setup,
		wafregexpatternset.Setup,
		wafregionalbytematchset.Setup,
		wafregionalgeomatchset.Setup,
		wafregionalipset.Setup,
		wafregionalratebasedrule.Setup,
		wafregionalregexmatchset.Setup,
		wafregionalregexpatternset.Setup,
		wafregionalrule.Setup,
		wafregionalsizeconstraintset.Setup,
		wafregionalsqlinjectionmatchset.Setup,
		wafregionalwebacl.Setup,
		wafregionalwebaclassociation.Setup,
		wafregionalxssmatchset.Setup,
		wafrule.Setup,
		wafsizeconstraintset.Setup,
		wafsqlinjectionmatchset.Setup,
		wafv2ipset.Setup,
		wafv2regexpatternset.Setup,
		wafv2rulegroup.Setup,
		wafv2webacl.Setup,
		wafv2webaclassociation.Setup,
		wafv2webaclloggingconfiguration.Setup,
		wafwebacl.Setup,
		wafxssmatchset.Setup,
		worklinkfleet.Setup,
		worklinkwebsitecertificateauthorityassociation.Setup,
		workspacesdirectory.Setup,
		workspacesipgroup.Setup,
		workspacesworkspace.Setup,
		xrayencryptionconfig.Setup,
		xraygroup.Setup,
		xraysamplingrule.Setup,
	} {
		if err := setup(mgr, l, wl); err != nil {
			return err
		}
	}
	return nil
}
