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

type AvailabilityZoneGroupObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AvailabilityZoneGroupParameters struct {

	// +kubebuilder:validation:Required
	GroupName *string `json:"groupName" tf:"group_name,omitempty"`

	// +kubebuilder:validation:Required
	OptInStatus *string `json:"optInStatus" tf:"opt_in_status,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// AvailabilityZoneGroupSpec defines the desired state of AvailabilityZoneGroup
type AvailabilityZoneGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AvailabilityZoneGroupParameters `json:"forProvider"`
}

// AvailabilityZoneGroupStatus defines the observed state of AvailabilityZoneGroup.
type AvailabilityZoneGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AvailabilityZoneGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AvailabilityZoneGroup is the Schema for the AvailabilityZoneGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type AvailabilityZoneGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AvailabilityZoneGroupSpec   `json:"spec"`
	Status            AvailabilityZoneGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AvailabilityZoneGroupList contains a list of AvailabilityZoneGroups
type AvailabilityZoneGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AvailabilityZoneGroup `json:"items"`
}

// Repository type metadata.
var (
	AvailabilityZoneGroup_Kind             = "AvailabilityZoneGroup"
	AvailabilityZoneGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AvailabilityZoneGroup_Kind}.String()
	AvailabilityZoneGroup_KindAPIVersion   = AvailabilityZoneGroup_Kind + "." + CRDGroupVersion.String()
	AvailabilityZoneGroup_GroupVersionKind = CRDGroupVersion.WithKind(AvailabilityZoneGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&AvailabilityZoneGroup{}, &AvailabilityZoneGroupList{})
}
