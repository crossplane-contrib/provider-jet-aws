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

type CloudwatchMetricStreamObservation struct {
	ARN string `json:"arn" tf:"arn"`

	CreationDate string `json:"creationDate" tf:"creation_date"`

	LastUpdateDate string `json:"lastUpdateDate" tf:"last_update_date"`

	State string `json:"state" tf:"state"`
}

type CloudwatchMetricStreamParameters struct {
	ExcludeFilter []ExcludeFilterParameters `json:"excludeFilter,omitempty" tf:"exclude_filter"`

	FirehoseARN string `json:"firehoseARN" tf:"firehose_arn"`

	IncludeFilter []IncludeFilterParameters `json:"includeFilter,omitempty" tf:"include_filter"`

	Name *string `json:"name,omitempty" tf:"name"`

	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`

	OutputFormat string `json:"outputFormat" tf:"output_format"`

	RoleARN string `json:"roleARN" tf:"role_arn"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type ExcludeFilterObservation struct {
}

type ExcludeFilterParameters struct {
	Namespace string `json:"namespace" tf:"namespace"`
}

type IncludeFilterObservation struct {
}

type IncludeFilterParameters struct {
	Namespace string `json:"namespace" tf:"namespace"`
}

// CloudwatchMetricStreamSpec defines the desired state of CloudwatchMetricStream
type CloudwatchMetricStreamSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CloudwatchMetricStreamParameters `json:"forProvider"`
}

// CloudwatchMetricStreamStatus defines the observed state of CloudwatchMetricStream.
type CloudwatchMetricStreamStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CloudwatchMetricStreamObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchMetricStream is the Schema for the CloudwatchMetricStreams API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type CloudwatchMetricStream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudwatchMetricStreamSpec   `json:"spec"`
	Status            CloudwatchMetricStreamStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchMetricStreamList contains a list of CloudwatchMetricStreams
type CloudwatchMetricStreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudwatchMetricStream `json:"items"`
}

// Repository type metadata.
var (
	CloudwatchMetricStreamKind             = "CloudwatchMetricStream"
	CloudwatchMetricStreamGroupKind        = schema.GroupKind{Group: Group, Kind: CloudwatchMetricStreamKind}.String()
	CloudwatchMetricStreamKindAPIVersion   = CloudwatchMetricStreamKind + "." + GroupVersion.String()
	CloudwatchMetricStreamGroupVersionKind = GroupVersion.WithKind(CloudwatchMetricStreamKind)
)

func init() {
	SchemeBuilder.Register(&CloudwatchMetricStream{}, &CloudwatchMetricStreamList{})
}
