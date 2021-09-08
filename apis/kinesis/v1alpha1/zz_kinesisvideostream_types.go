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

type KinesisVideoStreamObservation struct {
	ARN string `json:"arn" tf:"arn"`

	CreationTime string `json:"creationTime" tf:"creation_time"`

	Version string `json:"version" tf:"version"`
}

type KinesisVideoStreamParameters struct {
	DataRetentionInHours *int64 `json:"dataRetentionInHours,omitempty" tf:"data_retention_in_hours"`

	DeviceName *string `json:"deviceName,omitempty" tf:"device_name"`

	KmsKeyID *string `json:"kmsKeyID,omitempty" tf:"kms_key_id"`

	MediaType *string `json:"mediaType,omitempty" tf:"media_type"`

	Name string `json:"name" tf:"name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// KinesisVideoStreamSpec defines the desired state of KinesisVideoStream
type KinesisVideoStreamSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       KinesisVideoStreamParameters `json:"forProvider"`
}

// KinesisVideoStreamStatus defines the observed state of KinesisVideoStream.
type KinesisVideoStreamStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          KinesisVideoStreamObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KinesisVideoStream is the Schema for the KinesisVideoStreams API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type KinesisVideoStream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KinesisVideoStreamSpec   `json:"spec"`
	Status            KinesisVideoStreamStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KinesisVideoStreamList contains a list of KinesisVideoStreams
type KinesisVideoStreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KinesisVideoStream `json:"items"`
}

// Repository type metadata.
var (
	KinesisVideoStreamKind             = "KinesisVideoStream"
	KinesisVideoStreamGroupKind        = schema.GroupKind{Group: Group, Kind: KinesisVideoStreamKind}.String()
	KinesisVideoStreamKindAPIVersion   = KinesisVideoStreamKind + "." + GroupVersion.String()
	KinesisVideoStreamGroupVersionKind = GroupVersion.WithKind(KinesisVideoStreamKind)
)

func init() {
	SchemeBuilder.Register(&KinesisVideoStream{}, &KinesisVideoStreamList{})
}
