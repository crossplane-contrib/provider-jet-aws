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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type StandardsSubscriptionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type StandardsSubscriptionParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	StandardsArn *string `json:"standardsArn" tf:"standards_arn,omitempty"`
}

// StandardsSubscriptionSpec defines the desired state of StandardsSubscription
type StandardsSubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StandardsSubscriptionParameters `json:"forProvider"`
}

// StandardsSubscriptionStatus defines the observed state of StandardsSubscription.
type StandardsSubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StandardsSubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StandardsSubscription is the Schema for the StandardsSubscriptions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type StandardsSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StandardsSubscriptionSpec   `json:"spec"`
	Status            StandardsSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StandardsSubscriptionList contains a list of StandardsSubscriptions
type StandardsSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StandardsSubscription `json:"items"`
}

// Repository type metadata.
var (
	StandardsSubscription_Kind             = "StandardsSubscription"
	StandardsSubscription_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StandardsSubscription_Kind}.String()
	StandardsSubscription_KindAPIVersion   = StandardsSubscription_Kind + "." + CRDGroupVersion.String()
	StandardsSubscription_GroupVersionKind = CRDGroupVersion.WithKind(StandardsSubscription_Kind)
)

func init() {
	SchemeBuilder.Register(&StandardsSubscription{}, &StandardsSubscriptionList{})
}
