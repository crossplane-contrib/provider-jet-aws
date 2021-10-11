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

type AccepterObservation struct {
}

type AccepterParameters struct {

	// +kubebuilder:validation:Optional
	AllowClassicLinkToRemoteVpc *bool `json:"allowClassicLinkToRemoteVpc,omitempty" tf:"allow_classic_link_to_remote_vpc,omitempty"`

	// +kubebuilder:validation:Optional
	AllowRemoteVpcDNSResolution *bool `json:"allowRemoteVpcDnsResolution,omitempty" tf:"allow_remote_vpc_dns_resolution,omitempty"`

	// +kubebuilder:validation:Optional
	AllowVpcToRemoteClassicLink *bool `json:"allowVpcToRemoteClassicLink,omitempty" tf:"allow_vpc_to_remote_classic_link,omitempty"`
}

type PeeringConnectionObservation struct {
	AcceptStatus *string `json:"acceptStatus,omitempty" tf:"accept_status,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type PeeringConnectionParameters struct {

	// +kubebuilder:validation:Optional
	Accepter []AccepterParameters `json:"accepter,omitempty" tf:"accepter,omitempty"`

	// +kubebuilder:validation:Optional
	AutoAccept *bool `json:"autoAccept,omitempty" tf:"auto_accept,omitempty"`

	// +kubebuilder:validation:Optional
	PeerOwnerID *string `json:"peerOwnerId,omitempty" tf:"peer_owner_id,omitempty"`

	// +kubebuilder:validation:Optional
	PeerRegion *string `json:"peerRegion,omitempty" tf:"peer_region,omitempty"`

	// +crossplane:generate:reference:type=VPC
	// +kubebuilder:validation:Optional
	PeerVpcID *string `json:"peerVpcId,omitempty" tf:"peer_vpc_id,omitempty"`

	// +kubebuilder:validation:Optional
	PeerVpcIDRef *v1.Reference `json:"peerVpcIDRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PeerVpcIDSelector *v1.Selector `json:"peerVpcIDSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-,omitempty"`

	// +kubebuilder:validation:Optional
	Requester []RequesterParameters `json:"requester,omitempty" tf:"requester,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +crossplane:generate:reference:type=VPC
	// +kubebuilder:validation:Optional
	VpcID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// +kubebuilder:validation:Optional
	VpcIDRef *v1.Reference `json:"vpcIDRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VpcIDSelector *v1.Selector `json:"vpcIDSelector,omitempty" tf:"-"`
}

type RequesterObservation struct {
}

type RequesterParameters struct {

	// +kubebuilder:validation:Optional
	AllowClassicLinkToRemoteVpc *bool `json:"allowClassicLinkToRemoteVpc,omitempty" tf:"allow_classic_link_to_remote_vpc,omitempty"`

	// +kubebuilder:validation:Optional
	AllowRemoteVpcDNSResolution *bool `json:"allowRemoteVpcDnsResolution,omitempty" tf:"allow_remote_vpc_dns_resolution,omitempty"`

	// +kubebuilder:validation:Optional
	AllowVpcToRemoteClassicLink *bool `json:"allowVpcToRemoteClassicLink,omitempty" tf:"allow_vpc_to_remote_classic_link,omitempty"`
}

// PeeringConnectionSpec defines the desired state of PeeringConnection
type PeeringConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PeeringConnectionParameters `json:"forProvider"`
}

// PeeringConnectionStatus defines the observed state of PeeringConnection.
type PeeringConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PeeringConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PeeringConnection is the Schema for the PeeringConnections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type PeeringConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PeeringConnectionSpec   `json:"spec"`
	Status            PeeringConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PeeringConnectionList contains a list of PeeringConnections
type PeeringConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PeeringConnection `json:"items"`
}

// Repository type metadata.
var (
	PeeringConnectionKind             = "PeeringConnection"
	PeeringConnectionGroupKind        = schema.GroupKind{Group: Group, Kind: PeeringConnectionKind}.String()
	PeeringConnectionKindAPIVersion   = PeeringConnectionKind + "." + GroupVersion.String()
	PeeringConnectionGroupVersionKind = GroupVersion.WithKind(PeeringConnectionKind)
)

func init() {
	SchemeBuilder.Register(&PeeringConnection{}, &PeeringConnectionList{})
}