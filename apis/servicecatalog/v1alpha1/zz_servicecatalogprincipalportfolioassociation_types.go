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

type ServicecatalogPrincipalPortfolioAssociationObservation struct {
}

type ServicecatalogPrincipalPortfolioAssociationParameters struct {
	AcceptLanguage *string `json:"acceptLanguage,omitempty" tf:"accept_language"`

	PortfolioID string `json:"portfolioID" tf:"portfolio_id"`

	PrincipalARN string `json:"principalARN" tf:"principal_arn"`

	PrincipalType *string `json:"principalType,omitempty" tf:"principal_type"`
}

// ServicecatalogPrincipalPortfolioAssociationSpec defines the desired state of ServicecatalogPrincipalPortfolioAssociation
type ServicecatalogPrincipalPortfolioAssociationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ServicecatalogPrincipalPortfolioAssociationParameters `json:"forProvider"`
}

// ServicecatalogPrincipalPortfolioAssociationStatus defines the observed state of ServicecatalogPrincipalPortfolioAssociation.
type ServicecatalogPrincipalPortfolioAssociationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ServicecatalogPrincipalPortfolioAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServicecatalogPrincipalPortfolioAssociation is the Schema for the ServicecatalogPrincipalPortfolioAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ServicecatalogPrincipalPortfolioAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicecatalogPrincipalPortfolioAssociationSpec   `json:"spec"`
	Status            ServicecatalogPrincipalPortfolioAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServicecatalogPrincipalPortfolioAssociationList contains a list of ServicecatalogPrincipalPortfolioAssociations
type ServicecatalogPrincipalPortfolioAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServicecatalogPrincipalPortfolioAssociation `json:"items"`
}

// Repository type metadata.
var (
	ServicecatalogPrincipalPortfolioAssociationKind             = "ServicecatalogPrincipalPortfolioAssociation"
	ServicecatalogPrincipalPortfolioAssociationGroupKind        = schema.GroupKind{Group: Group, Kind: ServicecatalogPrincipalPortfolioAssociationKind}.String()
	ServicecatalogPrincipalPortfolioAssociationKindAPIVersion   = ServicecatalogPrincipalPortfolioAssociationKind + "." + GroupVersion.String()
	ServicecatalogPrincipalPortfolioAssociationGroupVersionKind = GroupVersion.WithKind(ServicecatalogPrincipalPortfolioAssociationKind)
)

func init() {
	SchemeBuilder.Register(&ServicecatalogPrincipalPortfolioAssociation{}, &ServicecatalogPrincipalPortfolioAssociationList{})
}
