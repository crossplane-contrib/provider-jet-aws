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

type Ec2ClientVpnNetworkAssociationObservation struct {
	AssociationID string `json:"associationID" tf:"association_id"`

	Status string `json:"status" tf:"status"`

	VPCID string `json:"vpcID" tf:"vpc_id"`
}

type Ec2ClientVpnNetworkAssociationParameters struct {
	ClientVpnEndpointID string `json:"clientVpnEndpointID" tf:"client_vpn_endpoint_id"`

	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups"`

	SubnetID string `json:"subnetID" tf:"subnet_id"`
}

// Ec2ClientVpnNetworkAssociationSpec defines the desired state of Ec2ClientVpnNetworkAssociation
type Ec2ClientVpnNetworkAssociationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Ec2ClientVpnNetworkAssociationParameters `json:"forProvider"`
}

// Ec2ClientVpnNetworkAssociationStatus defines the observed state of Ec2ClientVpnNetworkAssociation.
type Ec2ClientVpnNetworkAssociationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Ec2ClientVpnNetworkAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2ClientVpnNetworkAssociation is the Schema for the Ec2ClientVpnNetworkAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Ec2ClientVpnNetworkAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2ClientVpnNetworkAssociationSpec   `json:"spec"`
	Status            Ec2ClientVpnNetworkAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2ClientVpnNetworkAssociationList contains a list of Ec2ClientVpnNetworkAssociations
type Ec2ClientVpnNetworkAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2ClientVpnNetworkAssociation `json:"items"`
}

// Repository type metadata.
var (
	Ec2ClientVpnNetworkAssociationKind             = "Ec2ClientVpnNetworkAssociation"
	Ec2ClientVpnNetworkAssociationGroupKind        = schema.GroupKind{Group: Group, Kind: Ec2ClientVpnNetworkAssociationKind}.String()
	Ec2ClientVpnNetworkAssociationKindAPIVersion   = Ec2ClientVpnNetworkAssociationKind + "." + GroupVersion.String()
	Ec2ClientVpnNetworkAssociationGroupVersionKind = GroupVersion.WithKind(Ec2ClientVpnNetworkAssociationKind)
)

func init() {
	SchemeBuilder.Register(&Ec2ClientVpnNetworkAssociation{}, &Ec2ClientVpnNetworkAssociationList{})
}
