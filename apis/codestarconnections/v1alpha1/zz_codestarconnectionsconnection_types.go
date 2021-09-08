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

type CodestarconnectionsConnectionObservation struct {
	ARN string `json:"arn" tf:"arn"`

	ConnectionStatus string `json:"connectionStatus" tf:"connection_status"`
}

type CodestarconnectionsConnectionParameters struct {
	HostARN *string `json:"hostARN,omitempty" tf:"host_arn"`

	Name string `json:"name" tf:"name"`

	ProviderType *string `json:"providerType,omitempty" tf:"provider_type"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// CodestarconnectionsConnectionSpec defines the desired state of CodestarconnectionsConnection
type CodestarconnectionsConnectionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CodestarconnectionsConnectionParameters `json:"forProvider"`
}

// CodestarconnectionsConnectionStatus defines the observed state of CodestarconnectionsConnection.
type CodestarconnectionsConnectionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CodestarconnectionsConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CodestarconnectionsConnection is the Schema for the CodestarconnectionsConnections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type CodestarconnectionsConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodestarconnectionsConnectionSpec   `json:"spec"`
	Status            CodestarconnectionsConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodestarconnectionsConnectionList contains a list of CodestarconnectionsConnections
type CodestarconnectionsConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodestarconnectionsConnection `json:"items"`
}

// Repository type metadata.
var (
	CodestarconnectionsConnectionKind             = "CodestarconnectionsConnection"
	CodestarconnectionsConnectionGroupKind        = schema.GroupKind{Group: Group, Kind: CodestarconnectionsConnectionKind}.String()
	CodestarconnectionsConnectionKindAPIVersion   = CodestarconnectionsConnectionKind + "." + GroupVersion.String()
	CodestarconnectionsConnectionGroupVersionKind = GroupVersion.WithKind(CodestarconnectionsConnectionKind)
)

func init() {
	SchemeBuilder.Register(&CodestarconnectionsConnection{}, &CodestarconnectionsConnectionList{})
}
