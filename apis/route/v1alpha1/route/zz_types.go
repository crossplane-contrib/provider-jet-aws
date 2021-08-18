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
// +groupName=route.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/route/v1alpha1"
)

type RouteObservation struct {
	InstanceOwnerId string `json:"instanceOwnerId" tf:"instance_owner_id"`

	Origin string `json:"origin" tf:"origin"`

	State string `json:"state" tf:"state"`
}

type RouteParameters struct {
	CarrierGatewayId *string `json:"carrierGatewayId,omitempty" tf:"carrier_gateway_id"`

	DestinationCidrBlock *string `json:"destinationCidrBlock,omitempty" tf:"destination_cidr_block"`

	DestinationIpv6CidrBlock *string `json:"destinationIpv6CidrBlock,omitempty" tf:"destination_ipv6_cidr_block"`

	DestinationPrefixListId *string `json:"destinationPrefixListId,omitempty" tf:"destination_prefix_list_id"`

	EgressOnlyGatewayId *string `json:"egressOnlyGatewayId,omitempty" tf:"egress_only_gateway_id"`

	GatewayId *string `json:"gatewayId,omitempty" tf:"gateway_id"`

	InstanceId *string `json:"instanceId,omitempty" tf:"instance_id"`

	LocalGatewayId *string `json:"localGatewayId,omitempty" tf:"local_gateway_id"`

	NatGatewayId *string `json:"natGatewayId,omitempty" tf:"nat_gateway_id"`

	NetworkInterfaceId *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id"`

	RouteTableId string `json:"routeTableId" tf:"route_table_id"`

	TransitGatewayId *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id"`

	VpcEndpointId *string `json:"vpcEndpointId,omitempty" tf:"vpc_endpoint_id"`

	VpcPeeringConnectionId *string `json:"vpcPeeringConnectionId,omitempty" tf:"vpc_peering_connection_id"`
}

// RouteSpec defines the desired state of Route
type RouteSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       RouteParameters `json:"forProvider"`
}

// RouteStatus defines the observed state of Route.
type RouteStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          RouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Route is the Schema for the Routes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Route struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteSpec   `json:"spec"`
	Status            RouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteList contains a list of Routes
type RouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route `json:"items"`
}

// Repository type metadata.
var (
	RouteKind             = "Route"
	RouteGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: RouteKind}.String()
	RouteKindAPIVersion   = RouteKind + "." + v1alpha1.GroupVersion.String()
	RouteGroupVersionKind = v1alpha1.GroupVersion.WithKind(RouteKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&Route{}, &RouteList{})
}