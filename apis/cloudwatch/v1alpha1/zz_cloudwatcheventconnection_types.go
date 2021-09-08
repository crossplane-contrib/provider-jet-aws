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

type APIKeyObservation struct {
}

type APIKeyParameters struct {
	Key string `json:"key" tf:"key"`

	Value string `json:"value" tf:"value"`
}

type AuthParametersObservation struct {
}

type AuthParametersParameters struct {
	APIKey []APIKeyParameters `json:"apiKey,omitempty" tf:"api_key"`

	Basic []BasicParameters `json:"basic,omitempty" tf:"basic"`

	InvocationHTTPParameters []InvocationHTTPParametersParameters `json:"invocationHTTPParameters,omitempty" tf:"invocation_http_parameters"`

	Oauth []OauthParameters `json:"oauth,omitempty" tf:"oauth"`
}

type BasicObservation struct {
}

type BasicParameters struct {
	Password string `json:"password" tf:"password"`

	Username string `json:"username" tf:"username"`
}

type BodyObservation struct {
}

type BodyParameters struct {
	IsValueSecret *bool `json:"isValueSecret,omitempty" tf:"is_value_secret"`

	Key *string `json:"key,omitempty" tf:"key"`

	Value *string `json:"value,omitempty" tf:"value"`
}

type ClientParametersObservation struct {
}

type ClientParametersParameters struct {
	ClientID string `json:"clientID" tf:"client_id"`

	ClientSecret string `json:"clientSecret" tf:"client_secret"`
}

type CloudwatchEventConnectionObservation struct {
	ARN string `json:"arn" tf:"arn"`

	SecretARN string `json:"secretARN" tf:"secret_arn"`
}

type CloudwatchEventConnectionParameters struct {
	AuthParameters []AuthParametersParameters `json:"authParameters" tf:"auth_parameters"`

	AuthorizationType string `json:"authorizationType" tf:"authorization_type"`

	Description *string `json:"description,omitempty" tf:"description"`

	Name string `json:"name" tf:"name"`
}

type HeaderObservation struct {
}

type HeaderParameters struct {
	IsValueSecret *bool `json:"isValueSecret,omitempty" tf:"is_value_secret"`

	Key *string `json:"key,omitempty" tf:"key"`

	Value *string `json:"value,omitempty" tf:"value"`
}

type InvocationHTTPParametersObservation struct {
}

type InvocationHTTPParametersParameters struct {
	Body []BodyParameters `json:"body,omitempty" tf:"body"`

	Header []HeaderParameters `json:"header,omitempty" tf:"header"`

	QueryString []QueryStringParameters `json:"queryString,omitempty" tf:"query_string"`
}

type OauthHTTPParametersBodyObservation struct {
}

type OauthHTTPParametersBodyParameters struct {
	IsValueSecret *bool `json:"isValueSecret,omitempty" tf:"is_value_secret"`

	Key *string `json:"key,omitempty" tf:"key"`

	Value *string `json:"value,omitempty" tf:"value"`
}

type OauthHTTPParametersHeaderObservation struct {
}

type OauthHTTPParametersHeaderParameters struct {
	IsValueSecret *bool `json:"isValueSecret,omitempty" tf:"is_value_secret"`

	Key *string `json:"key,omitempty" tf:"key"`

	Value *string `json:"value,omitempty" tf:"value"`
}

type OauthHTTPParametersObservation struct {
}

type OauthHTTPParametersParameters struct {
	Body []OauthHTTPParametersBodyParameters `json:"body,omitempty" tf:"body"`

	Header []OauthHTTPParametersHeaderParameters `json:"header,omitempty" tf:"header"`

	QueryString []OauthHTTPParametersQueryStringParameters `json:"queryString,omitempty" tf:"query_string"`
}

type OauthHTTPParametersQueryStringObservation struct {
}

type OauthHTTPParametersQueryStringParameters struct {
	IsValueSecret *bool `json:"isValueSecret,omitempty" tf:"is_value_secret"`

	Key *string `json:"key,omitempty" tf:"key"`

	Value *string `json:"value,omitempty" tf:"value"`
}

type OauthObservation struct {
}

type OauthParameters struct {
	AuthorizationEndpoint string `json:"authorizationEndpoint" tf:"authorization_endpoint"`

	ClientParameters []ClientParametersParameters `json:"clientParameters,omitempty" tf:"client_parameters"`

	HTTPMethod string `json:"httpMethod" tf:"http_method"`

	OauthHTTPParameters []OauthHTTPParametersParameters `json:"oauthHTTPParameters" tf:"oauth_http_parameters"`
}

type QueryStringObservation struct {
}

type QueryStringParameters struct {
	IsValueSecret *bool `json:"isValueSecret,omitempty" tf:"is_value_secret"`

	Key *string `json:"key,omitempty" tf:"key"`

	Value *string `json:"value,omitempty" tf:"value"`
}

// CloudwatchEventConnectionSpec defines the desired state of CloudwatchEventConnection
type CloudwatchEventConnectionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CloudwatchEventConnectionParameters `json:"forProvider"`
}

// CloudwatchEventConnectionStatus defines the observed state of CloudwatchEventConnection.
type CloudwatchEventConnectionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CloudwatchEventConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchEventConnection is the Schema for the CloudwatchEventConnections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type CloudwatchEventConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudwatchEventConnectionSpec   `json:"spec"`
	Status            CloudwatchEventConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchEventConnectionList contains a list of CloudwatchEventConnections
type CloudwatchEventConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudwatchEventConnection `json:"items"`
}

// Repository type metadata.
var (
	CloudwatchEventConnectionKind             = "CloudwatchEventConnection"
	CloudwatchEventConnectionGroupKind        = schema.GroupKind{Group: Group, Kind: CloudwatchEventConnectionKind}.String()
	CloudwatchEventConnectionKindAPIVersion   = CloudwatchEventConnectionKind + "." + GroupVersion.String()
	CloudwatchEventConnectionGroupVersionKind = GroupVersion.WithKind(CloudwatchEventConnectionKind)
)

func init() {
	SchemeBuilder.Register(&CloudwatchEventConnection{}, &CloudwatchEventConnectionList{})
}
