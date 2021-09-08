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

type APIStagesObservation struct {
}

type APIStagesParameters struct {
	APIID string `json:"apiID" tf:"api_id"`

	Stage string `json:"stage" tf:"stage"`
}

type ApiGatewayUsagePlanObservation struct {
	ARN string `json:"arn" tf:"arn"`
}

type ApiGatewayUsagePlanParameters struct {
	APIStages []APIStagesParameters `json:"apiStages,omitempty" tf:"api_stages"`

	Description *string `json:"description,omitempty" tf:"description"`

	Name string `json:"name" tf:"name"`

	ProductCode *string `json:"productCode,omitempty" tf:"product_code"`

	QuotaSettings []QuotaSettingsParameters `json:"quotaSettings,omitempty" tf:"quota_settings"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	ThrottleSettings []ApiGatewayUsagePlanThrottleSettingsParameters `json:"throttleSettings,omitempty" tf:"throttle_settings"`
}

type ApiGatewayUsagePlanThrottleSettingsObservation struct {
}

type ApiGatewayUsagePlanThrottleSettingsParameters struct {
	BurstLimit *int64 `json:"burstLimit,omitempty" tf:"burst_limit"`

	RateLimit *float64 `json:"rateLimit,omitempty" tf:"rate_limit"`
}

type QuotaSettingsObservation struct {
}

type QuotaSettingsParameters struct {
	Limit int64 `json:"limit" tf:"limit"`

	Offset *int64 `json:"offset,omitempty" tf:"offset"`

	Period string `json:"period" tf:"period"`
}

// ApiGatewayUsagePlanSpec defines the desired state of ApiGatewayUsagePlan
type ApiGatewayUsagePlanSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApiGatewayUsagePlanParameters `json:"forProvider"`
}

// ApiGatewayUsagePlanStatus defines the observed state of ApiGatewayUsagePlan.
type ApiGatewayUsagePlanStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApiGatewayUsagePlanObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayUsagePlan is the Schema for the ApiGatewayUsagePlans API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ApiGatewayUsagePlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayUsagePlanSpec   `json:"spec"`
	Status            ApiGatewayUsagePlanStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayUsagePlanList contains a list of ApiGatewayUsagePlans
type ApiGatewayUsagePlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayUsagePlan `json:"items"`
}

// Repository type metadata.
var (
	ApiGatewayUsagePlanKind             = "ApiGatewayUsagePlan"
	ApiGatewayUsagePlanGroupKind        = schema.GroupKind{Group: Group, Kind: ApiGatewayUsagePlanKind}.String()
	ApiGatewayUsagePlanKindAPIVersion   = ApiGatewayUsagePlanKind + "." + GroupVersion.String()
	ApiGatewayUsagePlanGroupVersionKind = GroupVersion.WithKind(ApiGatewayUsagePlanKind)
)

func init() {
	SchemeBuilder.Register(&ApiGatewayUsagePlan{}, &ApiGatewayUsagePlanList{})
}
