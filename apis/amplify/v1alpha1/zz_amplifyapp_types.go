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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type AmplifyAppObservation struct {
	ARN string `json:"arn" tf:"arn"`

	DefaultDomain string `json:"defaultDomain" tf:"default_domain"`

	ProductionBranch []ProductionBranchObservation `json:"productionBranch" tf:"production_branch"`
}

type AmplifyAppParameters struct {
	AccessToken *string `json:"accessToken,omitempty" tf:"access_token"`

	AutoBranchCreationConfig []AutoBranchCreationConfigParameters `json:"autoBranchCreationConfig,omitempty" tf:"auto_branch_creation_config"`

	AutoBranchCreationPatterns []string `json:"autoBranchCreationPatterns,omitempty" tf:"auto_branch_creation_patterns"`

	BasicAuthCredentials *string `json:"basicAuthCredentials,omitempty" tf:"basic_auth_credentials"`

	BuildSpec *string `json:"buildSpec,omitempty" tf:"build_spec"`

	CustomRule []CustomRuleParameters `json:"customRule,omitempty" tf:"custom_rule"`

	Description *string `json:"description,omitempty" tf:"description"`

	EnableAutoBranchCreation *bool `json:"enableAutoBranchCreation,omitempty" tf:"enable_auto_branch_creation"`

	EnableBasicAuth *bool `json:"enableBasicAuth,omitempty" tf:"enable_basic_auth"`

	EnableBranchAutoBuild *bool `json:"enableBranchAutoBuild,omitempty" tf:"enable_branch_auto_build"`

	EnableBranchAutoDeletion *bool `json:"enableBranchAutoDeletion,omitempty" tf:"enable_branch_auto_deletion"`

	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty" tf:"environment_variables"`

	IamServiceRoleARN *string `json:"iamServiceRoleARN,omitempty" tf:"iam_service_role_arn"`

	Name string `json:"name" tf:"name"`

	OauthToken *string `json:"oauthToken,omitempty" tf:"oauth_token"`

	Platform *string `json:"platform,omitempty" tf:"platform"`

	Repository *string `json:"repository,omitempty" tf:"repository"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type AutoBranchCreationConfigObservation struct {
}

type AutoBranchCreationConfigParameters struct {
	BasicAuthCredentials *string `json:"basicAuthCredentials,omitempty" tf:"basic_auth_credentials"`

	BuildSpec *string `json:"buildSpec,omitempty" tf:"build_spec"`

	EnableAutoBuild *bool `json:"enableAutoBuild,omitempty" tf:"enable_auto_build"`

	EnableBasicAuth *bool `json:"enableBasicAuth,omitempty" tf:"enable_basic_auth"`

	EnablePerformanceMode *bool `json:"enablePerformanceMode,omitempty" tf:"enable_performance_mode"`

	EnablePullRequestPreview *bool `json:"enablePullRequestPreview,omitempty" tf:"enable_pull_request_preview"`

	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty" tf:"environment_variables"`

	Framework *string `json:"framework,omitempty" tf:"framework"`

	PullRequestEnvironmentName *string `json:"pullRequestEnvironmentName,omitempty" tf:"pull_request_environment_name"`

	Stage *string `json:"stage,omitempty" tf:"stage"`
}

type CustomRuleObservation struct {
}

type CustomRuleParameters struct {
	Condition *string `json:"condition,omitempty" tf:"condition"`

	Source string `json:"source" tf:"source"`

	Status *string `json:"status,omitempty" tf:"status"`

	Target string `json:"target" tf:"target"`
}

type ProductionBranchObservation struct {
	BranchName string `json:"branchName" tf:"branch_name"`

	LastDeployTime string `json:"lastDeployTime" tf:"last_deploy_time"`

	Status string `json:"status" tf:"status"`

	ThumbnailURL string `json:"thumbnailURL" tf:"thumbnail_url"`
}

type ProductionBranchParameters struct {
}

// AmplifyAppSpec defines the desired state of AmplifyApp
type AmplifyAppSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AmplifyAppParameters `json:"forProvider"`
}

// AmplifyAppStatus defines the observed state of AmplifyApp.
type AmplifyAppStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AmplifyAppObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AmplifyApp is the Schema for the AmplifyApps API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type AmplifyApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AmplifyAppSpec   `json:"spec"`
	Status            AmplifyAppStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AmplifyAppList contains a list of AmplifyApps
type AmplifyAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AmplifyApp `json:"items"`
}

// Repository type metadata.
var (
	AmplifyAppKind             = "AmplifyApp"
	AmplifyAppGroupKind        = schema.GroupKind{Group: Group, Kind: AmplifyAppKind}.String()
	AmplifyAppKindAPIVersion   = AmplifyAppKind + "." + GroupVersion.String()
	AmplifyAppGroupVersionKind = GroupVersion.WithKind(AmplifyAppKind)
)

func init() {
	SchemeBuilder.Register(&AmplifyApp{}, &AmplifyAppList{})
}
