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

type ParameterObservation struct {
}

type ParameterParameters struct {
	Name string `json:"name" tf:"name"`

	Value string `json:"value" tf:"value"`
}

type RedshiftParameterGroupObservation struct {
	ARN string `json:"arn" tf:"arn"`
}

type RedshiftParameterGroupParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	Family string `json:"family" tf:"family"`

	Name string `json:"name" tf:"name"`

	Parameter []ParameterParameters `json:"parameter,omitempty" tf:"parameter"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// RedshiftParameterGroupSpec defines the desired state of RedshiftParameterGroup
type RedshiftParameterGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       RedshiftParameterGroupParameters `json:"forProvider"`
}

// RedshiftParameterGroupStatus defines the observed state of RedshiftParameterGroup.
type RedshiftParameterGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          RedshiftParameterGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RedshiftParameterGroup is the Schema for the RedshiftParameterGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type RedshiftParameterGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedshiftParameterGroupSpec   `json:"spec"`
	Status            RedshiftParameterGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedshiftParameterGroupList contains a list of RedshiftParameterGroups
type RedshiftParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedshiftParameterGroup `json:"items"`
}

// Repository type metadata.
var (
	RedshiftParameterGroupKind             = "RedshiftParameterGroup"
	RedshiftParameterGroupGroupKind        = schema.GroupKind{Group: Group, Kind: RedshiftParameterGroupKind}.String()
	RedshiftParameterGroupKindAPIVersion   = RedshiftParameterGroupKind + "." + GroupVersion.String()
	RedshiftParameterGroupGroupVersionKind = GroupVersion.WithKind(RedshiftParameterGroupKind)
)

func init() {
	SchemeBuilder.Register(&RedshiftParameterGroup{}, &RedshiftParameterGroupList{})
}
