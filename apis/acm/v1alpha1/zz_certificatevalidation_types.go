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

type CertificateValidationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CertificateValidationParameters struct {

	// +kubebuilder:validation:Required
	CertificateArn *string `json:"certificateArn" tf:"certificate_arn,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	ValidationRecordFqdns []*string `json:"validationRecordFqdns,omitempty" tf:"validation_record_fqdns,omitempty"`
}

// CertificateValidationSpec defines the desired state of CertificateValidation
type CertificateValidationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CertificateValidationParameters `json:"forProvider"`
}

// CertificateValidationStatus defines the observed state of CertificateValidation.
type CertificateValidationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CertificateValidationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateValidation is the Schema for the CertificateValidations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type CertificateValidation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CertificateValidationSpec   `json:"spec"`
	Status            CertificateValidationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateValidationList contains a list of CertificateValidations
type CertificateValidationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertificateValidation `json:"items"`
}

// Repository type metadata.
var (
	CertificateValidation_Kind             = "CertificateValidation"
	CertificateValidation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CertificateValidation_Kind}.String()
	CertificateValidation_KindAPIVersion   = CertificateValidation_Kind + "." + CRDGroupVersion.String()
	CertificateValidation_GroupVersionKind = CRDGroupVersion.WithKind(CertificateValidation_Kind)
)

func init() {
	SchemeBuilder.Register(&CertificateValidation{}, &CertificateValidationList{})
}
