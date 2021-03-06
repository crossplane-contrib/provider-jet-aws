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

type TransitGatewayRouteTablePropagationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`
}

type TransitGatewayRouteTablePropagationParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +crossplane:generate:reference:type=TransitGatewayVPCAttachment
	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentID *string `json:"transitGatewayAttachmentId,omitempty" tf:"transit_gateway_attachment_id,omitempty"`

	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentIDRef *v1.Reference `json:"transitGatewayAttachmentIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentIDSelector *v1.Selector `json:"transitGatewayAttachmentIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=TransitGatewayRouteTable
	// +kubebuilder:validation:Optional
	TransitGatewayRouteTableID *string `json:"transitGatewayRouteTableId,omitempty" tf:"transit_gateway_route_table_id,omitempty"`

	// +kubebuilder:validation:Optional
	TransitGatewayRouteTableIDRef *v1.Reference `json:"transitGatewayRouteTableIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TransitGatewayRouteTableIDSelector *v1.Selector `json:"transitGatewayRouteTableIdSelector,omitempty" tf:"-"`
}

// TransitGatewayRouteTablePropagationSpec defines the desired state of TransitGatewayRouteTablePropagation
type TransitGatewayRouteTablePropagationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitGatewayRouteTablePropagationParameters `json:"forProvider"`
}

// TransitGatewayRouteTablePropagationStatus defines the observed state of TransitGatewayRouteTablePropagation.
type TransitGatewayRouteTablePropagationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitGatewayRouteTablePropagationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayRouteTablePropagation is the Schema for the TransitGatewayRouteTablePropagations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type TransitGatewayRouteTablePropagation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransitGatewayRouteTablePropagationSpec   `json:"spec"`
	Status            TransitGatewayRouteTablePropagationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayRouteTablePropagationList contains a list of TransitGatewayRouteTablePropagations
type TransitGatewayRouteTablePropagationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitGatewayRouteTablePropagation `json:"items"`
}

// Repository type metadata.
var (
	TransitGatewayRouteTablePropagation_Kind             = "TransitGatewayRouteTablePropagation"
	TransitGatewayRouteTablePropagation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitGatewayRouteTablePropagation_Kind}.String()
	TransitGatewayRouteTablePropagation_KindAPIVersion   = TransitGatewayRouteTablePropagation_Kind + "." + CRDGroupVersion.String()
	TransitGatewayRouteTablePropagation_GroupVersionKind = CRDGroupVersion.WithKind(TransitGatewayRouteTablePropagation_Kind)
)

func init() {
	SchemeBuilder.Register(&TransitGatewayRouteTablePropagation{}, &TransitGatewayRouteTablePropagationList{})
}
