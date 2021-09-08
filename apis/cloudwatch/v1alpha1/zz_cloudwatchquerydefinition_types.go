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

type CloudwatchQueryDefinitionObservation struct {
	QueryDefinitionID string `json:"queryDefinitionID" tf:"query_definition_id"`
}

type CloudwatchQueryDefinitionParameters struct {
	LogGroupNames []string `json:"logGroupNames,omitempty" tf:"log_group_names"`

	Name string `json:"name" tf:"name"`

	QueryString string `json:"queryString" tf:"query_string"`
}

// CloudwatchQueryDefinitionSpec defines the desired state of CloudwatchQueryDefinition
type CloudwatchQueryDefinitionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CloudwatchQueryDefinitionParameters `json:"forProvider"`
}

// CloudwatchQueryDefinitionStatus defines the observed state of CloudwatchQueryDefinition.
type CloudwatchQueryDefinitionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CloudwatchQueryDefinitionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchQueryDefinition is the Schema for the CloudwatchQueryDefinitions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type CloudwatchQueryDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudwatchQueryDefinitionSpec   `json:"spec"`
	Status            CloudwatchQueryDefinitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchQueryDefinitionList contains a list of CloudwatchQueryDefinitions
type CloudwatchQueryDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudwatchQueryDefinition `json:"items"`
}

// Repository type metadata.
var (
	CloudwatchQueryDefinitionKind             = "CloudwatchQueryDefinition"
	CloudwatchQueryDefinitionGroupKind        = schema.GroupKind{Group: Group, Kind: CloudwatchQueryDefinitionKind}.String()
	CloudwatchQueryDefinitionKindAPIVersion   = CloudwatchQueryDefinitionKind + "." + GroupVersion.String()
	CloudwatchQueryDefinitionGroupVersionKind = GroupVersion.WithKind(CloudwatchQueryDefinitionKind)
)

func init() {
	SchemeBuilder.Register(&CloudwatchQueryDefinition{}, &CloudwatchQueryDefinitionList{})
}
