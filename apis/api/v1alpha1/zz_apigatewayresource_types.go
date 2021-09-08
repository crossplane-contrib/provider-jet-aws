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

type ApiGatewayResourceObservation struct {
	Path string `json:"path" tf:"path"`
}

type ApiGatewayResourceParameters struct {
	ParentID string `json:"parentID" tf:"parent_id"`

	PathPart string `json:"pathPart" tf:"path_part"`

	RestAPIID string `json:"restAPIID" tf:"rest_api_id"`
}

// ApiGatewayResourceSpec defines the desired state of ApiGatewayResource
type ApiGatewayResourceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApiGatewayResourceParameters `json:"forProvider"`
}

// ApiGatewayResourceStatus defines the observed state of ApiGatewayResource.
type ApiGatewayResourceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApiGatewayResourceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayResource is the Schema for the ApiGatewayResources API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ApiGatewayResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayResourceSpec   `json:"spec"`
	Status            ApiGatewayResourceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayResourceList contains a list of ApiGatewayResources
type ApiGatewayResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayResource `json:"items"`
}

// Repository type metadata.
var (
	ApiGatewayResourceKind             = "ApiGatewayResource"
	ApiGatewayResourceGroupKind        = schema.GroupKind{Group: Group, Kind: ApiGatewayResourceKind}.String()
	ApiGatewayResourceKindAPIVersion   = ApiGatewayResourceKind + "." + GroupVersion.String()
	ApiGatewayResourceGroupVersionKind = GroupVersion.WithKind(ApiGatewayResourceKind)
)

func init() {
	SchemeBuilder.Register(&ApiGatewayResource{}, &ApiGatewayResourceList{})
}
