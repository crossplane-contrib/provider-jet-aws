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

type AdditionalAuthenticationProviderObservation struct {
}

type AdditionalAuthenticationProviderParameters struct {
	AuthenticationType string `json:"authenticationType" tf:"authentication_type"`

	OpenidConnectConfig []OpenidConnectConfigParameters `json:"openidConnectConfig,omitempty" tf:"openid_connect_config"`

	UserPoolConfig []UserPoolConfigParameters `json:"userPoolConfig,omitempty" tf:"user_pool_config"`
}

type AppsyncGraphqlApiObservation struct {
	ARN string `json:"arn" tf:"arn"`

	Uris map[string]string `json:"uris" tf:"uris"`
}

type AppsyncGraphqlApiOpenidConnectConfigObservation struct {
}

type AppsyncGraphqlApiOpenidConnectConfigParameters struct {
	AuthTTL *int64 `json:"authTTL,omitempty" tf:"auth_ttl"`

	ClientID *string `json:"clientID,omitempty" tf:"client_id"`

	IatTTL *int64 `json:"iatTTL,omitempty" tf:"iat_ttl"`

	Issuer string `json:"issuer" tf:"issuer"`
}

type AppsyncGraphqlApiParameters struct {
	AdditionalAuthenticationProvider []AdditionalAuthenticationProviderParameters `json:"additionalAuthenticationProvider,omitempty" tf:"additional_authentication_provider"`

	AuthenticationType string `json:"authenticationType" tf:"authentication_type"`

	LogConfig []LogConfigParameters `json:"logConfig,omitempty" tf:"log_config"`

	Name string `json:"name" tf:"name"`

	OpenidConnectConfig []AppsyncGraphqlApiOpenidConnectConfigParameters `json:"openidConnectConfig,omitempty" tf:"openid_connect_config"`

	Schema *string `json:"schema,omitempty" tf:"schema"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	UserPoolConfig []AppsyncGraphqlApiUserPoolConfigParameters `json:"userPoolConfig,omitempty" tf:"user_pool_config"`

	XrayEnabled *bool `json:"xrayEnabled,omitempty" tf:"xray_enabled"`
}

type AppsyncGraphqlApiUserPoolConfigObservation struct {
}

type AppsyncGraphqlApiUserPoolConfigParameters struct {
	AppIDClientRegex *string `json:"appIDClientRegex,omitempty" tf:"app_id_client_regex"`

	AwsRegion *string `json:"awsRegion,omitempty" tf:"aws_region"`

	DefaultAction string `json:"defaultAction" tf:"default_action"`

	UserPoolID string `json:"userPoolID" tf:"user_pool_id"`
}

type LogConfigObservation struct {
}

type LogConfigParameters struct {
	CloudwatchLogsRoleARN string `json:"cloudwatchLogsRoleARN" tf:"cloudwatch_logs_role_arn"`

	ExcludeVerboseContent *bool `json:"excludeVerboseContent,omitempty" tf:"exclude_verbose_content"`

	FieldLogLevel string `json:"fieldLogLevel" tf:"field_log_level"`
}

type OpenidConnectConfigObservation struct {
}

type OpenidConnectConfigParameters struct {
	AuthTTL *int64 `json:"authTTL,omitempty" tf:"auth_ttl"`

	ClientID *string `json:"clientID,omitempty" tf:"client_id"`

	IatTTL *int64 `json:"iatTTL,omitempty" tf:"iat_ttl"`

	Issuer string `json:"issuer" tf:"issuer"`
}

type UserPoolConfigObservation struct {
}

type UserPoolConfigParameters struct {
	AppIDClientRegex *string `json:"appIDClientRegex,omitempty" tf:"app_id_client_regex"`

	AwsRegion *string `json:"awsRegion,omitempty" tf:"aws_region"`

	UserPoolID string `json:"userPoolID" tf:"user_pool_id"`
}

// AppsyncGraphqlApiSpec defines the desired state of AppsyncGraphqlApi
type AppsyncGraphqlApiSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AppsyncGraphqlApiParameters `json:"forProvider"`
}

// AppsyncGraphqlApiStatus defines the observed state of AppsyncGraphqlApi.
type AppsyncGraphqlApiStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AppsyncGraphqlApiObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppsyncGraphqlApi is the Schema for the AppsyncGraphqlApis API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type AppsyncGraphqlApi struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppsyncGraphqlApiSpec   `json:"spec"`
	Status            AppsyncGraphqlApiStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppsyncGraphqlApiList contains a list of AppsyncGraphqlApis
type AppsyncGraphqlApiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppsyncGraphqlApi `json:"items"`
}

// Repository type metadata.
var (
	AppsyncGraphqlApiKind             = "AppsyncGraphqlApi"
	AppsyncGraphqlApiGroupKind        = schema.GroupKind{Group: Group, Kind: AppsyncGraphqlApiKind}.String()
	AppsyncGraphqlApiKindAPIVersion   = AppsyncGraphqlApiKind + "." + GroupVersion.String()
	AppsyncGraphqlApiGroupVersionKind = GroupVersion.WithKind(AppsyncGraphqlApiKind)
)

func init() {
	SchemeBuilder.Register(&AppsyncGraphqlApi{}, &AppsyncGraphqlApiList{})
}
