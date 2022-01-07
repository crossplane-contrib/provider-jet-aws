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

type NetworkInterfaceAttachmentObservation struct {
	AttachmentID *string `json:"attachmentId,omitempty" tf:"attachment_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type NetworkInterfaceAttachmentParameters struct {

	// +kubebuilder:validation:Required
	DeviceIndex *int64 `json:"deviceIndex" tf:"device_index,omitempty"`

	// +kubebuilder:validation:Required
	InstanceID *string `json:"instanceId" tf:"instance_id,omitempty"`

	// +kubebuilder:validation:Required
	NetworkInterfaceID *string `json:"networkInterfaceId" tf:"network_interface_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// NetworkInterfaceAttachmentSpec defines the desired state of NetworkInterfaceAttachment
type NetworkInterfaceAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkInterfaceAttachmentParameters `json:"forProvider"`
}

// NetworkInterfaceAttachmentStatus defines the observed state of NetworkInterfaceAttachment.
type NetworkInterfaceAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkInterfaceAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInterfaceAttachment is the Schema for the NetworkInterfaceAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type NetworkInterfaceAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkInterfaceAttachmentSpec   `json:"spec"`
	Status            NetworkInterfaceAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInterfaceAttachmentList contains a list of NetworkInterfaceAttachments
type NetworkInterfaceAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkInterfaceAttachment `json:"items"`
}

// Repository type metadata.
var (
	NetworkInterfaceAttachment_Kind             = "NetworkInterfaceAttachment"
	NetworkInterfaceAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkInterfaceAttachment_Kind}.String()
	NetworkInterfaceAttachment_KindAPIVersion   = NetworkInterfaceAttachment_Kind + "." + CRDGroupVersion.String()
	NetworkInterfaceAttachment_GroupVersionKind = CRDGroupVersion.WithKind(NetworkInterfaceAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkInterfaceAttachment{}, &NetworkInterfaceAttachmentList{})
}
