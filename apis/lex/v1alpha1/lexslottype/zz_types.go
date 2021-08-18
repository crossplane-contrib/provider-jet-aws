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
// +groupName=lex.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/lex/v1alpha1"
)

type EnumerationValueObservation struct {
}

type EnumerationValueParameters struct {
	Synonyms []string `json:"synonyms,omitempty" tf:"synonyms"`

	Value string `json:"value" tf:"value"`
}

type LexSlotTypeObservation struct {
	Checksum string `json:"checksum" tf:"checksum"`

	CreatedDate string `json:"createdDate" tf:"created_date"`

	LastUpdatedDate string `json:"lastUpdatedDate" tf:"last_updated_date"`

	Version string `json:"version" tf:"version"`
}

type LexSlotTypeParameters struct {
	CreateVersion *bool `json:"createVersion,omitempty" tf:"create_version"`

	Description *string `json:"description,omitempty" tf:"description"`

	EnumerationValue []EnumerationValueParameters `json:"enumerationValue" tf:"enumeration_value"`

	Name string `json:"name" tf:"name"`

	ValueSelectionStrategy *string `json:"valueSelectionStrategy,omitempty" tf:"value_selection_strategy"`
}

// LexSlotTypeSpec defines the desired state of LexSlotType
type LexSlotTypeSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       LexSlotTypeParameters `json:"forProvider"`
}

// LexSlotTypeStatus defines the observed state of LexSlotType.
type LexSlotTypeStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          LexSlotTypeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LexSlotType is the Schema for the LexSlotTypes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LexSlotType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LexSlotTypeSpec   `json:"spec"`
	Status            LexSlotTypeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LexSlotTypeList contains a list of LexSlotTypes
type LexSlotTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LexSlotType `json:"items"`
}

// Repository type metadata.
var (
	LexSlotTypeKind             = "LexSlotType"
	LexSlotTypeGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: LexSlotTypeKind}.String()
	LexSlotTypeKindAPIVersion   = LexSlotTypeKind + "." + v1alpha1.GroupVersion.String()
	LexSlotTypeGroupVersionKind = v1alpha1.GroupVersion.WithKind(LexSlotTypeKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&LexSlotType{}, &LexSlotTypeList{})
}