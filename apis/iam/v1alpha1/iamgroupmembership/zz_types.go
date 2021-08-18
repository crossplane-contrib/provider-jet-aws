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
// +groupName=iam.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/iam/v1alpha1"
)

type IamGroupMembershipObservation struct {
}

type IamGroupMembershipParameters struct {
	Group string `json:"group" tf:"group"`

	Name string `json:"name" tf:"name"`

	Users []string `json:"users" tf:"users"`
}

// IamGroupMembershipSpec defines the desired state of IamGroupMembership
type IamGroupMembershipSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       IamGroupMembershipParameters `json:"forProvider"`
}

// IamGroupMembershipStatus defines the observed state of IamGroupMembership.
type IamGroupMembershipStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          IamGroupMembershipObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IamGroupMembership is the Schema for the IamGroupMemberships API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type IamGroupMembership struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamGroupMembershipSpec   `json:"spec"`
	Status            IamGroupMembershipStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamGroupMembershipList contains a list of IamGroupMemberships
type IamGroupMembershipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamGroupMembership `json:"items"`
}

// Repository type metadata.
var (
	IamGroupMembershipKind             = "IamGroupMembership"
	IamGroupMembershipGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: IamGroupMembershipKind}.String()
	IamGroupMembershipKindAPIVersion   = IamGroupMembershipKind + "." + v1alpha1.GroupVersion.String()
	IamGroupMembershipGroupVersionKind = v1alpha1.GroupVersion.WithKind(IamGroupMembershipKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&IamGroupMembership{}, &IamGroupMembershipList{})
}