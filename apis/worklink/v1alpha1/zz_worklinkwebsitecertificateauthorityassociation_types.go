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

type WorklinkWebsiteCertificateAuthorityAssociationObservation struct {
	WebsiteCaID string `json:"websiteCaID" tf:"website_ca_id"`
}

type WorklinkWebsiteCertificateAuthorityAssociationParameters struct {
	Certificate string `json:"certificate" tf:"certificate"`

	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`

	FleetARN string `json:"fleetARN" tf:"fleet_arn"`
}

// WorklinkWebsiteCertificateAuthorityAssociationSpec defines the desired state of WorklinkWebsiteCertificateAuthorityAssociation
type WorklinkWebsiteCertificateAuthorityAssociationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       WorklinkWebsiteCertificateAuthorityAssociationParameters `json:"forProvider"`
}

// WorklinkWebsiteCertificateAuthorityAssociationStatus defines the observed state of WorklinkWebsiteCertificateAuthorityAssociation.
type WorklinkWebsiteCertificateAuthorityAssociationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          WorklinkWebsiteCertificateAuthorityAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WorklinkWebsiteCertificateAuthorityAssociation is the Schema for the WorklinkWebsiteCertificateAuthorityAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type WorklinkWebsiteCertificateAuthorityAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorklinkWebsiteCertificateAuthorityAssociationSpec   `json:"spec"`
	Status            WorklinkWebsiteCertificateAuthorityAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorklinkWebsiteCertificateAuthorityAssociationList contains a list of WorklinkWebsiteCertificateAuthorityAssociations
type WorklinkWebsiteCertificateAuthorityAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorklinkWebsiteCertificateAuthorityAssociation `json:"items"`
}

// Repository type metadata.
var (
	WorklinkWebsiteCertificateAuthorityAssociationKind             = "WorklinkWebsiteCertificateAuthorityAssociation"
	WorklinkWebsiteCertificateAuthorityAssociationGroupKind        = schema.GroupKind{Group: Group, Kind: WorklinkWebsiteCertificateAuthorityAssociationKind}.String()
	WorklinkWebsiteCertificateAuthorityAssociationKindAPIVersion   = WorklinkWebsiteCertificateAuthorityAssociationKind + "." + GroupVersion.String()
	WorklinkWebsiteCertificateAuthorityAssociationGroupVersionKind = GroupVersion.WithKind(WorklinkWebsiteCertificateAuthorityAssociationKind)
)

func init() {
	SchemeBuilder.Register(&WorklinkWebsiteCertificateAuthorityAssociation{}, &WorklinkWebsiteCertificateAuthorityAssociationList{})
}
