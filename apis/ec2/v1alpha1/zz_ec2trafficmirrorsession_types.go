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

type Ec2TrafficMirrorSessionObservation struct {
	ARN string `json:"arn" tf:"arn"`

	OwnerID string `json:"ownerID" tf:"owner_id"`
}

type Ec2TrafficMirrorSessionParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	NetworkInterfaceID string `json:"networkInterfaceID" tf:"network_interface_id"`

	PacketLength *int64 `json:"packetLength,omitempty" tf:"packet_length"`

	SessionNumber int64 `json:"sessionNumber" tf:"session_number"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	TrafficMirrorFilterID string `json:"trafficMirrorFilterID" tf:"traffic_mirror_filter_id"`

	TrafficMirrorTargetID string `json:"trafficMirrorTargetID" tf:"traffic_mirror_target_id"`

	VirtualNetworkID *int64 `json:"virtualNetworkID,omitempty" tf:"virtual_network_id"`
}

// Ec2TrafficMirrorSessionSpec defines the desired state of Ec2TrafficMirrorSession
type Ec2TrafficMirrorSessionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Ec2TrafficMirrorSessionParameters `json:"forProvider"`
}

// Ec2TrafficMirrorSessionStatus defines the observed state of Ec2TrafficMirrorSession.
type Ec2TrafficMirrorSessionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Ec2TrafficMirrorSessionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TrafficMirrorSession is the Schema for the Ec2TrafficMirrorSessions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Ec2TrafficMirrorSession struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2TrafficMirrorSessionSpec   `json:"spec"`
	Status            Ec2TrafficMirrorSessionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TrafficMirrorSessionList contains a list of Ec2TrafficMirrorSessions
type Ec2TrafficMirrorSessionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2TrafficMirrorSession `json:"items"`
}

// Repository type metadata.
var (
	Ec2TrafficMirrorSessionKind             = "Ec2TrafficMirrorSession"
	Ec2TrafficMirrorSessionGroupKind        = schema.GroupKind{Group: Group, Kind: Ec2TrafficMirrorSessionKind}.String()
	Ec2TrafficMirrorSessionKindAPIVersion   = Ec2TrafficMirrorSessionKind + "." + GroupVersion.String()
	Ec2TrafficMirrorSessionGroupVersionKind = GroupVersion.WithKind(Ec2TrafficMirrorSessionKind)
)

func init() {
	SchemeBuilder.Register(&Ec2TrafficMirrorSession{}, &Ec2TrafficMirrorSessionList{})
}
