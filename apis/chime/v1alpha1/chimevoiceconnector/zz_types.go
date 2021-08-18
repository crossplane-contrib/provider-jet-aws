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
// +groupName=chime.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/chime/v1alpha1"
)

type ChimeVoiceConnectorObservation struct {
	OutboundHostName string `json:"outboundHostName" tf:"outbound_host_name"`
}

type ChimeVoiceConnectorParameters struct {
	AwsRegion *string `json:"awsRegion,omitempty" tf:"aws_region"`

	Name string `json:"name" tf:"name"`

	RequireEncryption bool `json:"requireEncryption" tf:"require_encryption"`
}

// ChimeVoiceConnectorSpec defines the desired state of ChimeVoiceConnector
type ChimeVoiceConnectorSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ChimeVoiceConnectorParameters `json:"forProvider"`
}

// ChimeVoiceConnectorStatus defines the observed state of ChimeVoiceConnector.
type ChimeVoiceConnectorStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ChimeVoiceConnectorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ChimeVoiceConnector is the Schema for the ChimeVoiceConnectors API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ChimeVoiceConnector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ChimeVoiceConnectorSpec   `json:"spec"`
	Status            ChimeVoiceConnectorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ChimeVoiceConnectorList contains a list of ChimeVoiceConnectors
type ChimeVoiceConnectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ChimeVoiceConnector `json:"items"`
}

// Repository type metadata.
var (
	ChimeVoiceConnectorKind             = "ChimeVoiceConnector"
	ChimeVoiceConnectorGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: ChimeVoiceConnectorKind}.String()
	ChimeVoiceConnectorKindAPIVersion   = ChimeVoiceConnectorKind + "." + v1alpha1.GroupVersion.String()
	ChimeVoiceConnectorGroupVersionKind = v1alpha1.GroupVersion.WithKind(ChimeVoiceConnectorKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&ChimeVoiceConnector{}, &ChimeVoiceConnectorList{})
}