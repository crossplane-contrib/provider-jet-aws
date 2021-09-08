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

type WafXssMatchSetObservation struct {
	ARN string `json:"arn" tf:"arn"`
}

type WafXssMatchSetParameters struct {
	Name string `json:"name" tf:"name"`

	XSSMatchTuples []XSSMatchTuplesParameters `json:"xssMatchTuples,omitempty" tf:"xss_match_tuples"`
}

type XSSMatchTuplesFieldToMatchObservation struct {
}

type XSSMatchTuplesFieldToMatchParameters struct {
	Data *string `json:"data,omitempty" tf:"data"`

	Type string `json:"type" tf:"type"`
}

type XSSMatchTuplesObservation struct {
}

type XSSMatchTuplesParameters struct {
	FieldToMatch []XSSMatchTuplesFieldToMatchParameters `json:"fieldToMatch" tf:"field_to_match"`

	TextTransformation string `json:"textTransformation" tf:"text_transformation"`
}

// WafXssMatchSetSpec defines the desired state of WafXssMatchSet
type WafXssMatchSetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       WafXssMatchSetParameters `json:"forProvider"`
}

// WafXssMatchSetStatus defines the observed state of WafXssMatchSet.
type WafXssMatchSetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          WafXssMatchSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WafXssMatchSet is the Schema for the WafXssMatchSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type WafXssMatchSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafXssMatchSetSpec   `json:"spec"`
	Status            WafXssMatchSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafXssMatchSetList contains a list of WafXssMatchSets
type WafXssMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafXssMatchSet `json:"items"`
}

// Repository type metadata.
var (
	WafXssMatchSetKind             = "WafXssMatchSet"
	WafXssMatchSetGroupKind        = schema.GroupKind{Group: Group, Kind: WafXssMatchSetKind}.String()
	WafXssMatchSetKindAPIVersion   = WafXssMatchSetKind + "." + GroupVersion.String()
	WafXssMatchSetGroupVersionKind = GroupVersion.WithKind(WafXssMatchSetKind)
)

func init() {
	SchemeBuilder.Register(&WafXssMatchSet{}, &WafXssMatchSetList{})
}
