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

type CloudwatchEventApiDestinationObservation struct {
	ARN string `json:"arn" tf:"arn"`
}

type CloudwatchEventApiDestinationParameters struct {
	ConnectionARN string `json:"connectionARN" tf:"connection_arn"`

	Description *string `json:"description,omitempty" tf:"description"`

	HTTPMethod string `json:"httpMethod" tf:"http_method"`

	InvocationEndpoint string `json:"invocationEndpoint" tf:"invocation_endpoint"`

	InvocationRateLimitPerSecond *int64 `json:"invocationRateLimitPerSecond,omitempty" tf:"invocation_rate_limit_per_second"`

	Name string `json:"name" tf:"name"`
}

// CloudwatchEventApiDestinationSpec defines the desired state of CloudwatchEventApiDestination
type CloudwatchEventApiDestinationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CloudwatchEventApiDestinationParameters `json:"forProvider"`
}

// CloudwatchEventApiDestinationStatus defines the observed state of CloudwatchEventApiDestination.
type CloudwatchEventApiDestinationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CloudwatchEventApiDestinationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchEventApiDestination is the Schema for the CloudwatchEventApiDestinations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type CloudwatchEventApiDestination struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudwatchEventApiDestinationSpec   `json:"spec"`
	Status            CloudwatchEventApiDestinationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchEventApiDestinationList contains a list of CloudwatchEventApiDestinations
type CloudwatchEventApiDestinationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudwatchEventApiDestination `json:"items"`
}

// Repository type metadata.
var (
	CloudwatchEventApiDestinationKind             = "CloudwatchEventApiDestination"
	CloudwatchEventApiDestinationGroupKind        = schema.GroupKind{Group: Group, Kind: CloudwatchEventApiDestinationKind}.String()
	CloudwatchEventApiDestinationKindAPIVersion   = CloudwatchEventApiDestinationKind + "." + GroupVersion.String()
	CloudwatchEventApiDestinationGroupVersionKind = GroupVersion.WithKind(CloudwatchEventApiDestinationKind)
)

func init() {
	SchemeBuilder.Register(&CloudwatchEventApiDestination{}, &CloudwatchEventApiDestinationList{})
}
