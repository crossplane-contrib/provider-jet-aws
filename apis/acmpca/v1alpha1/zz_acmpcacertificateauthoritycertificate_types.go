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

type AcmpcaCertificateAuthorityCertificateObservation struct {
}

type AcmpcaCertificateAuthorityCertificateParameters struct {
	Certificate string `json:"certificate" tf:"certificate"`

	CertificateAuthorityARN string `json:"certificateAuthorityARN" tf:"certificate_authority_arn"`

	CertificateChain *string `json:"certificateChain,omitempty" tf:"certificate_chain"`
}

// AcmpcaCertificateAuthorityCertificateSpec defines the desired state of AcmpcaCertificateAuthorityCertificate
type AcmpcaCertificateAuthorityCertificateSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AcmpcaCertificateAuthorityCertificateParameters `json:"forProvider"`
}

// AcmpcaCertificateAuthorityCertificateStatus defines the observed state of AcmpcaCertificateAuthorityCertificate.
type AcmpcaCertificateAuthorityCertificateStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AcmpcaCertificateAuthorityCertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AcmpcaCertificateAuthorityCertificate is the Schema for the AcmpcaCertificateAuthorityCertificates API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type AcmpcaCertificateAuthorityCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AcmpcaCertificateAuthorityCertificateSpec   `json:"spec"`
	Status            AcmpcaCertificateAuthorityCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AcmpcaCertificateAuthorityCertificateList contains a list of AcmpcaCertificateAuthorityCertificates
type AcmpcaCertificateAuthorityCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AcmpcaCertificateAuthorityCertificate `json:"items"`
}

// Repository type metadata.
var (
	AcmpcaCertificateAuthorityCertificateKind             = "AcmpcaCertificateAuthorityCertificate"
	AcmpcaCertificateAuthorityCertificateGroupKind        = schema.GroupKind{Group: Group, Kind: AcmpcaCertificateAuthorityCertificateKind}.String()
	AcmpcaCertificateAuthorityCertificateKindAPIVersion   = AcmpcaCertificateAuthorityCertificateKind + "." + GroupVersion.String()
	AcmpcaCertificateAuthorityCertificateGroupVersionKind = GroupVersion.WithKind(AcmpcaCertificateAuthorityCertificateKind)
)

func init() {
	SchemeBuilder.Register(&AcmpcaCertificateAuthorityCertificate{}, &AcmpcaCertificateAuthorityCertificateList{})
}
