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

type AccessLogSettingsObservation struct {
}

type AccessLogSettingsParameters struct {
	DestinationARN string `json:"destinationARN" tf:"destination_arn"`

	Format string `json:"format" tf:"format"`
}

type Apigatewayv2StageObservation struct {
	ARN string `json:"arn" tf:"arn"`

	ExecutionARN string `json:"executionARN" tf:"execution_arn"`

	InvokeURL string `json:"invokeURL" tf:"invoke_url"`
}

type Apigatewayv2StageParameters struct {
	APIID string `json:"apiID" tf:"api_id"`

	AccessLogSettings []AccessLogSettingsParameters `json:"accessLogSettings,omitempty" tf:"access_log_settings"`

	AutoDeploy *bool `json:"autoDeploy,omitempty" tf:"auto_deploy"`

	ClientCertificateID *string `json:"clientCertificateID,omitempty" tf:"client_certificate_id"`

	DefaultRouteSettings []DefaultRouteSettingsParameters `json:"defaultRouteSettings,omitempty" tf:"default_route_settings"`

	DeploymentID *string `json:"deploymentID,omitempty" tf:"deployment_id"`

	Description *string `json:"description,omitempty" tf:"description"`

	Name string `json:"name" tf:"name"`

	RouteSettings []RouteSettingsParameters `json:"routeSettings,omitempty" tf:"route_settings"`

	StageVariables map[string]string `json:"stageVariables,omitempty" tf:"stage_variables"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type DefaultRouteSettingsObservation struct {
}

type DefaultRouteSettingsParameters struct {
	DataTraceEnabled *bool `json:"dataTraceEnabled,omitempty" tf:"data_trace_enabled"`

	DetailedMetricsEnabled *bool `json:"detailedMetricsEnabled,omitempty" tf:"detailed_metrics_enabled"`

	LoggingLevel *string `json:"loggingLevel,omitempty" tf:"logging_level"`

	ThrottlingBurstLimit *int64 `json:"throttlingBurstLimit,omitempty" tf:"throttling_burst_limit"`

	ThrottlingRateLimit *float64 `json:"throttlingRateLimit,omitempty" tf:"throttling_rate_limit"`
}

type RouteSettingsObservation struct {
}

type RouteSettingsParameters struct {
	DataTraceEnabled *bool `json:"dataTraceEnabled,omitempty" tf:"data_trace_enabled"`

	DetailedMetricsEnabled *bool `json:"detailedMetricsEnabled,omitempty" tf:"detailed_metrics_enabled"`

	LoggingLevel *string `json:"loggingLevel,omitempty" tf:"logging_level"`

	RouteKey string `json:"routeKey" tf:"route_key"`

	ThrottlingBurstLimit *int64 `json:"throttlingBurstLimit,omitempty" tf:"throttling_burst_limit"`

	ThrottlingRateLimit *float64 `json:"throttlingRateLimit,omitempty" tf:"throttling_rate_limit"`
}

// Apigatewayv2StageSpec defines the desired state of Apigatewayv2Stage
type Apigatewayv2StageSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Apigatewayv2StageParameters `json:"forProvider"`
}

// Apigatewayv2StageStatus defines the observed state of Apigatewayv2Stage.
type Apigatewayv2StageStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Apigatewayv2StageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Apigatewayv2Stage is the Schema for the Apigatewayv2Stages API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Apigatewayv2Stage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Apigatewayv2StageSpec   `json:"spec"`
	Status            Apigatewayv2StageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Apigatewayv2StageList contains a list of Apigatewayv2Stages
type Apigatewayv2StageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Apigatewayv2Stage `json:"items"`
}

// Repository type metadata.
var (
	Apigatewayv2StageKind             = "Apigatewayv2Stage"
	Apigatewayv2StageGroupKind        = schema.GroupKind{Group: Group, Kind: Apigatewayv2StageKind}.String()
	Apigatewayv2StageKindAPIVersion   = Apigatewayv2StageKind + "." + GroupVersion.String()
	Apigatewayv2StageGroupVersionKind = GroupVersion.WithKind(Apigatewayv2StageKind)
)

func init() {
	SchemeBuilder.Register(&Apigatewayv2Stage{}, &Apigatewayv2StageList{})
}
