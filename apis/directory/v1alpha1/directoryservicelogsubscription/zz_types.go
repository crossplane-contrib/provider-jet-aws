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

// +kubebuilder:object:generate=true
// +groupName=directory.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/directory/v1alpha1"
)

type DirectoryServiceLogSubscriptionObservation struct {
}

type DirectoryServiceLogSubscriptionParameters struct {
	DirectoryId string `json:"directoryId" tf:"directory_id"`

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
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
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
	DirectoryServiceLogSubscriptionGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: DirectoryServiceLogSubscriptionKind}.String()
	DirectoryServiceLogSubscriptionKindAPIVersion   = DirectoryServiceLogSubscriptionKind + "." + v1alpha1.GroupVersion.String()
	DirectoryServiceLogSubscriptionGroupVersionKind = v1alpha1.GroupVersion.WithKind(DirectoryServiceLogSubscriptionKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&DirectoryServiceLogSubscription{}, &DirectoryServiceLogSubscriptionList{})
}