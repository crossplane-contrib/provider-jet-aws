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

type Apigatewayv2VpcLinkObservation struct {
	ARN string `json:"arn" tf:"arn"`
}

type Apigatewayv2VpcLinkParameters struct {
	Name string `json:"name" tf:"name"`

	SecurityGroupIds []string `json:"securityGroupIds" tf:"security_group_ids"`

	SubnetIds []string `json:"subnetIds" tf:"subnet_ids"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// Apigatewayv2VpcLinkSpec defines the desired state of Apigatewayv2VpcLink
type Apigatewayv2VpcLinkSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Apigatewayv2VpcLinkParameters `json:"forProvider"`
}

// Apigatewayv2VpcLinkStatus defines the observed state of Apigatewayv2VpcLink.
type Apigatewayv2VpcLinkStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Apigatewayv2VpcLinkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Apigatewayv2VpcLink is the Schema for the Apigatewayv2VpcLinks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Apigatewayv2VpcLink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Apigatewayv2VpcLinkSpec   `json:"spec"`
	Status            Apigatewayv2VpcLinkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Apigatewayv2VpcLinkList contains a list of Apigatewayv2VpcLinks
type Apigatewayv2VpcLinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Apigatewayv2VpcLink `json:"items"`
}

// Repository type metadata.
var (
	Apigatewayv2VpcLinkKind             = "Apigatewayv2VpcLink"
	Apigatewayv2VpcLinkGroupKind        = schema.GroupKind{Group: Group, Kind: Apigatewayv2VpcLinkKind}.String()
	Apigatewayv2VpcLinkKindAPIVersion   = Apigatewayv2VpcLinkKind + "." + GroupVersion.String()
	Apigatewayv2VpcLinkGroupVersionKind = GroupVersion.WithKind(Apigatewayv2VpcLinkKind)
)

func init() {
	SchemeBuilder.Register(&Apigatewayv2VpcLink{}, &Apigatewayv2VpcLinkList{})
}
