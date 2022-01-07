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

type PrivateVirtualInterfaceObservation struct {
	AmazonSideAsn *string `json:"amazonSideAsn,omitempty" tf:"amazon_side_asn,omitempty"`

	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	AwsDevice *string `json:"awsDevice,omitempty" tf:"aws_device,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	JumboFrameCapable *bool `json:"jumboFrameCapable,omitempty" tf:"jumbo_frame_capable,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type PrivateVirtualInterfaceParameters struct {

	// +kubebuilder:validation:Required
	AddressFamily *string `json:"addressFamily" tf:"address_family,omitempty"`

	// +kubebuilder:validation:Optional
	AmazonAddress *string `json:"amazonAddress,omitempty" tf:"amazon_address,omitempty"`

	// +kubebuilder:validation:Required
	BGPAsn *int64 `json:"bgpAsn" tf:"bgp_asn,omitempty"`

	// +kubebuilder:validation:Optional
	BGPAuthKey *string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key,omitempty"`

	// +kubebuilder:validation:Required
	ConnectionID *string `json:"connectionId" tf:"connection_id,omitempty"`

	// +kubebuilder:validation:Optional
	CustomerAddress *string `json:"customerAddress,omitempty" tf:"customer_address,omitempty"`

	// +kubebuilder:validation:Optional
	DxGatewayID *string `json:"dxGatewayId,omitempty" tf:"dx_gateway_id,omitempty"`

	// +kubebuilder:validation:Optional
	Mtu *int64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	VPNGatewayID *string `json:"vpnGatewayId,omitempty" tf:"vpn_gateway_id,omitempty"`

	// +kubebuilder:validation:Required
	Vlan *int64 `json:"vlan" tf:"vlan,omitempty"`
}

// PrivateVirtualInterfaceSpec defines the desired state of PrivateVirtualInterface
type PrivateVirtualInterfaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivateVirtualInterfaceParameters `json:"forProvider"`
}

// PrivateVirtualInterfaceStatus defines the observed state of PrivateVirtualInterface.
type PrivateVirtualInterfaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivateVirtualInterfaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateVirtualInterface is the Schema for the PrivateVirtualInterfaces API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type PrivateVirtualInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateVirtualInterfaceSpec   `json:"spec"`
	Status            PrivateVirtualInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateVirtualInterfaceList contains a list of PrivateVirtualInterfaces
type PrivateVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateVirtualInterface `json:"items"`
}

// Repository type metadata.
var (
	PrivateVirtualInterface_Kind             = "PrivateVirtualInterface"
	PrivateVirtualInterface_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrivateVirtualInterface_Kind}.String()
	PrivateVirtualInterface_KindAPIVersion   = PrivateVirtualInterface_Kind + "." + CRDGroupVersion.String()
	PrivateVirtualInterface_GroupVersionKind = CRDGroupVersion.WithKind(PrivateVirtualInterface_Kind)
)

func init() {
	SchemeBuilder.Register(&PrivateVirtualInterface{}, &PrivateVirtualInterfaceList{})
}
