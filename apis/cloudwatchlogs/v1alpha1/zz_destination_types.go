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

type DestinationObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DestinationParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/iam/v1alpha2.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	TargetArn *string `json:"targetArn" tf:"target_arn,omitempty"`
}

// DestinationSpec defines the desired state of Destination
type DestinationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DestinationParameters `json:"forProvider"`
}

// DestinationStatus defines the observed state of Destination.
type DestinationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DestinationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Destination is the Schema for the Destinations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Destination struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DestinationSpec   `json:"spec"`
	Status            DestinationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DestinationList contains a list of Destinations
type DestinationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Destination `json:"items"`
}

// Repository type metadata.
var (
	Destination_Kind             = "Destination"
	Destination_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Destination_Kind}.String()
	Destination_KindAPIVersion   = Destination_Kind + "." + CRDGroupVersion.String()
	Destination_GroupVersionKind = CRDGroupVersion.WithKind(Destination_Kind)
)

func init() {
	SchemeBuilder.Register(&Destination{}, &DestinationList{})
}
