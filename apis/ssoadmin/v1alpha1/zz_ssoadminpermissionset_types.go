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

type SsoadminPermissionSetObservation struct {
	ARN string `json:"arn" tf:"arn"`

	CreatedDate string `json:"createdDate" tf:"created_date"`
}

type SsoadminPermissionSetParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	InstanceARN string `json:"instanceARN" tf:"instance_arn"`

	Name string `json:"name" tf:"name"`

	RelayState *string `json:"relayState,omitempty" tf:"relay_state"`

	SessionDuration *string `json:"sessionDuration,omitempty" tf:"session_duration"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// SsoadminPermissionSetSpec defines the desired state of SsoadminPermissionSet
type SsoadminPermissionSetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SsoadminPermissionSetParameters `json:"forProvider"`
}

// SsoadminPermissionSetStatus defines the observed state of SsoadminPermissionSet.
type SsoadminPermissionSetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SsoadminPermissionSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SsoadminPermissionSet is the Schema for the SsoadminPermissionSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type SsoadminPermissionSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SsoadminPermissionSetSpec   `json:"spec"`
	Status            SsoadminPermissionSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SsoadminPermissionSetList contains a list of SsoadminPermissionSets
type SsoadminPermissionSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SsoadminPermissionSet `json:"items"`
}

// Repository type metadata.
var (
	SsoadminPermissionSetKind             = "SsoadminPermissionSet"
	SsoadminPermissionSetGroupKind        = schema.GroupKind{Group: Group, Kind: SsoadminPermissionSetKind}.String()
	SsoadminPermissionSetKindAPIVersion   = SsoadminPermissionSetKind + "." + GroupVersion.String()
	SsoadminPermissionSetGroupVersionKind = GroupVersion.WithKind(SsoadminPermissionSetKind)
)

func init() {
	SchemeBuilder.Register(&SsoadminPermissionSet{}, &SsoadminPermissionSetList{})
}
