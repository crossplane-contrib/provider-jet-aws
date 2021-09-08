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

type EcrLifecyclePolicyObservation struct {
	RegistryID string `json:"registryID" tf:"registry_id"`
}

type EcrLifecyclePolicyParameters struct {
	Policy string `json:"policy" tf:"policy"`

	Repository string `json:"repository" tf:"repository"`
}

// EcrLifecyclePolicySpec defines the desired state of EcrLifecyclePolicy
type EcrLifecyclePolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EcrLifecyclePolicyParameters `json:"forProvider"`
}

// EcrLifecyclePolicyStatus defines the observed state of EcrLifecyclePolicy.
type EcrLifecyclePolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EcrLifecyclePolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EcrLifecyclePolicy is the Schema for the EcrLifecyclePolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type EcrLifecyclePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EcrLifecyclePolicySpec   `json:"spec"`
	Status            EcrLifecyclePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EcrLifecyclePolicyList contains a list of EcrLifecyclePolicys
type EcrLifecyclePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EcrLifecyclePolicy `json:"items"`
}

// Repository type metadata.
var (
	EcrLifecyclePolicyKind             = "EcrLifecyclePolicy"
	EcrLifecyclePolicyGroupKind        = schema.GroupKind{Group: Group, Kind: EcrLifecyclePolicyKind}.String()
	EcrLifecyclePolicyKindAPIVersion   = EcrLifecyclePolicyKind + "." + GroupVersion.String()
	EcrLifecyclePolicyGroupVersionKind = GroupVersion.WithKind(EcrLifecyclePolicyKind)
)

func init() {
	SchemeBuilder.Register(&EcrLifecyclePolicy{}, &EcrLifecyclePolicyList{})
}
