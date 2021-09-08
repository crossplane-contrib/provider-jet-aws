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

type SsmMaintenanceWindowTargetObservation struct {
}

type SsmMaintenanceWindowTargetParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	Name *string `json:"name,omitempty" tf:"name"`

	OwnerInformation *string `json:"ownerInformation,omitempty" tf:"owner_information"`

	ResourceType string `json:"resourceType" tf:"resource_type"`

	Targets []SsmMaintenanceWindowTargetTargetsParameters `json:"targets" tf:"targets"`

	WindowID string `json:"windowID" tf:"window_id"`
}

type SsmMaintenanceWindowTargetTargetsObservation struct {
}

type SsmMaintenanceWindowTargetTargetsParameters struct {
	Key string `json:"key" tf:"key"`

	Values []string `json:"values" tf:"values"`
}

// SsmMaintenanceWindowTargetSpec defines the desired state of SsmMaintenanceWindowTarget
type SsmMaintenanceWindowTargetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SsmMaintenanceWindowTargetParameters `json:"forProvider"`
}

// SsmMaintenanceWindowTargetStatus defines the observed state of SsmMaintenanceWindowTarget.
type SsmMaintenanceWindowTargetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SsmMaintenanceWindowTargetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SsmMaintenanceWindowTarget is the Schema for the SsmMaintenanceWindowTargets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type SsmMaintenanceWindowTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SsmMaintenanceWindowTargetSpec   `json:"spec"`
	Status            SsmMaintenanceWindowTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SsmMaintenanceWindowTargetList contains a list of SsmMaintenanceWindowTargets
type SsmMaintenanceWindowTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SsmMaintenanceWindowTarget `json:"items"`
}

// Repository type metadata.
var (
	SsmMaintenanceWindowTargetKind             = "SsmMaintenanceWindowTarget"
	SsmMaintenanceWindowTargetGroupKind        = schema.GroupKind{Group: Group, Kind: SsmMaintenanceWindowTargetKind}.String()
	SsmMaintenanceWindowTargetKindAPIVersion   = SsmMaintenanceWindowTargetKind + "." + GroupVersion.String()
	SsmMaintenanceWindowTargetGroupVersionKind = GroupVersion.WithKind(SsmMaintenanceWindowTargetKind)
)

func init() {
	SchemeBuilder.Register(&SsmMaintenanceWindowTarget{}, &SsmMaintenanceWindowTargetList{})
}
