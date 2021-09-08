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

type AppmeshVirtualServiceObservation struct {
	ARN string `json:"arn" tf:"arn"`

	CreatedDate string `json:"createdDate" tf:"created_date"`

	LastUpdatedDate string `json:"lastUpdatedDate" tf:"last_updated_date"`

	ResourceOwner string `json:"resourceOwner" tf:"resource_owner"`
}

type AppmeshVirtualServiceParameters struct {
	MeshName string `json:"meshName" tf:"mesh_name"`

	MeshOwner *string `json:"meshOwner,omitempty" tf:"mesh_owner"`

	Name string `json:"name" tf:"name"`

	Spec []AppmeshVirtualServiceSpecParameters `json:"spec" tf:"spec"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type AppmeshVirtualServiceSpecObservation struct {
}

type AppmeshVirtualServiceSpecParameters struct {
	Provider []ProviderParameters `json:"provider,omitempty" tf:"provider"`
}

type ProviderObservation struct {
}

type ProviderParameters struct {
	VirtualNode []VirtualNodeParameters `json:"virtualNode,omitempty" tf:"virtual_node"`

	VirtualRouter []VirtualRouterParameters `json:"virtualRouter,omitempty" tf:"virtual_router"`
}

type VirtualNodeObservation struct {
}

type VirtualNodeParameters struct {
	VirtualNodeName string `json:"virtualNodeName" tf:"virtual_node_name"`
}

type VirtualRouterObservation struct {
}

type VirtualRouterParameters struct {
	VirtualRouterName string `json:"virtualRouterName" tf:"virtual_router_name"`
}

// AppmeshVirtualServiceSpec defines the desired state of AppmeshVirtualService
type AppmeshVirtualServiceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AppmeshVirtualServiceParameters `json:"forProvider"`
}

// AppmeshVirtualServiceStatus defines the observed state of AppmeshVirtualService.
type AppmeshVirtualServiceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AppmeshVirtualServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppmeshVirtualService is the Schema for the AppmeshVirtualServices API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type AppmeshVirtualService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppmeshVirtualServiceSpec   `json:"spec"`
	Status            AppmeshVirtualServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppmeshVirtualServiceList contains a list of AppmeshVirtualServices
type AppmeshVirtualServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppmeshVirtualService `json:"items"`
}

// Repository type metadata.
var (
	AppmeshVirtualServiceKind             = "AppmeshVirtualService"
	AppmeshVirtualServiceGroupKind        = schema.GroupKind{Group: Group, Kind: AppmeshVirtualServiceKind}.String()
	AppmeshVirtualServiceKindAPIVersion   = AppmeshVirtualServiceKind + "." + GroupVersion.String()
	AppmeshVirtualServiceGroupVersionKind = GroupVersion.WithKind(AppmeshVirtualServiceKind)
)

func init() {
	SchemeBuilder.Register(&AppmeshVirtualService{}, &AppmeshVirtualServiceList{})
}
