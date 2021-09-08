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

type DirectoryServiceLogSubscriptionObservation struct {
}

type DirectoryServiceLogSubscriptionParameters struct {
	DirectoryID string `json:"directoryID" tf:"directory_id"`

	LogGroupName string `json:"logGroupName" tf:"log_group_name"`
}

// DirectoryServiceLogSubscriptionSpec defines the desired state of DirectoryServiceLogSubscription
type DirectoryServiceLogSubscriptionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DirectoryServiceLogSubscriptionParameters `json:"forProvider"`
}

// DirectoryServiceLogSubscriptionStatus defines the observed state of DirectoryServiceLogSubscription.
type DirectoryServiceLogSubscriptionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DirectoryServiceLogSubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DirectoryServiceLogSubscription is the Schema for the DirectoryServiceLogSubscriptions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type DirectoryServiceLogSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DirectoryServiceLogSubscriptionSpec   `json:"spec"`
	Status            DirectoryServiceLogSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DirectoryServiceLogSubscriptionList contains a list of DirectoryServiceLogSubscriptions
type DirectoryServiceLogSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DirectoryServiceLogSubscription `json:"items"`
}

// Repository type metadata.
var (
	DirectoryServiceLogSubscriptionKind             = "DirectoryServiceLogSubscription"
	DirectoryServiceLogSubscriptionGroupKind        = schema.GroupKind{Group: Group, Kind: DirectoryServiceLogSubscriptionKind}.String()
	DirectoryServiceLogSubscriptionKindAPIVersion   = DirectoryServiceLogSubscriptionKind + "." + GroupVersion.String()
	DirectoryServiceLogSubscriptionGroupVersionKind = GroupVersion.WithKind(DirectoryServiceLogSubscriptionKind)
)

func init() {
	SchemeBuilder.Register(&DirectoryServiceLogSubscription{}, &DirectoryServiceLogSubscriptionList{})
}
