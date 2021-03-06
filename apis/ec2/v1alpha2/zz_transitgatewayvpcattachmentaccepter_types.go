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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TransitGatewayVPCAttachmentAccepterObservation struct {
	ApplianceModeSupport *string `json:"applianceModeSupport,omitempty" tf:"appliance_mode_support,omitempty"`

	DNSSupport *string `json:"dnsSupport,omitempty" tf:"dns_support,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IPv6Support *string `json:"ipv6Support,omitempty" tf:"ipv6_support,omitempty"`

	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id,omitempty"`

	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	VPCOwnerID *string `json:"vpcOwnerId,omitempty" tf:"vpc_owner_id,omitempty"`
}

type TransitGatewayVPCAttachmentAccepterParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +crossplane:generate:reference:type=TransitGatewayVPCAttachment
	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentID *string `json:"transitGatewayAttachmentId,omitempty" tf:"transit_gateway_attachment_id,omitempty"`

	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentIDRef *v1.Reference `json:"transitGatewayAttachmentIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentIDSelector *v1.Selector `json:"transitGatewayAttachmentIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TransitGatewayDefaultRouteTableAssociation *bool `json:"transitGatewayDefaultRouteTableAssociation,omitempty" tf:"transit_gateway_default_route_table_association,omitempty"`

	// +kubebuilder:validation:Optional
	TransitGatewayDefaultRouteTablePropagation *bool `json:"transitGatewayDefaultRouteTablePropagation,omitempty" tf:"transit_gateway_default_route_table_propagation,omitempty"`
}

// TransitGatewayVPCAttachmentAccepterSpec defines the desired state of TransitGatewayVPCAttachmentAccepter
type TransitGatewayVPCAttachmentAccepterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitGatewayVPCAttachmentAccepterParameters `json:"forProvider"`
}

// TransitGatewayVPCAttachmentAccepterStatus defines the observed state of TransitGatewayVPCAttachmentAccepter.
type TransitGatewayVPCAttachmentAccepterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitGatewayVPCAttachmentAccepterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayVPCAttachmentAccepter is the Schema for the TransitGatewayVPCAttachmentAccepters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type TransitGatewayVPCAttachmentAccepter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransitGatewayVPCAttachmentAccepterSpec   `json:"spec"`
	Status            TransitGatewayVPCAttachmentAccepterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayVPCAttachmentAccepterList contains a list of TransitGatewayVPCAttachmentAccepters
type TransitGatewayVPCAttachmentAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitGatewayVPCAttachmentAccepter `json:"items"`
}

// Repository type metadata.
var (
	TransitGatewayVPCAttachmentAccepter_Kind             = "TransitGatewayVPCAttachmentAccepter"
	TransitGatewayVPCAttachmentAccepter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitGatewayVPCAttachmentAccepter_Kind}.String()
	TransitGatewayVPCAttachmentAccepter_KindAPIVersion   = TransitGatewayVPCAttachmentAccepter_Kind + "." + CRDGroupVersion.String()
	TransitGatewayVPCAttachmentAccepter_GroupVersionKind = CRDGroupVersion.WithKind(TransitGatewayVPCAttachmentAccepter_Kind)
)

func init() {
	SchemeBuilder.Register(&TransitGatewayVPCAttachmentAccepter{}, &TransitGatewayVPCAttachmentAccepterList{})
}
