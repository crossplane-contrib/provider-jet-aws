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

type IamUserLoginProfileObservation struct {
	EncryptedPassword string `json:"encryptedPassword" tf:"encrypted_password"`

	KeyFingerprint string `json:"keyFingerprint" tf:"key_fingerprint"`
}

type IamUserLoginProfileParameters struct {
	PasswordLength *int64 `json:"passwordLength,omitempty" tf:"password_length"`

	PasswordResetRequired *bool `json:"passwordResetRequired,omitempty" tf:"password_reset_required"`

	PgpKey string `json:"pgpKey" tf:"pgp_key"`

	User string `json:"user" tf:"user"`
}

// IamUserLoginProfileSpec defines the desired state of IamUserLoginProfile
type IamUserLoginProfileSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       IamUserLoginProfileParameters `json:"forProvider"`
}

// IamUserLoginProfileStatus defines the observed state of IamUserLoginProfile.
type IamUserLoginProfileStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          IamUserLoginProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IamUserLoginProfile is the Schema for the IamUserLoginProfiles API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type IamUserLoginProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamUserLoginProfileSpec   `json:"spec"`
	Status            IamUserLoginProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamUserLoginProfileList contains a list of IamUserLoginProfiles
type IamUserLoginProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamUserLoginProfile `json:"items"`
}

// Repository type metadata.
var (
	IamUserLoginProfileKind             = "IamUserLoginProfile"
	IamUserLoginProfileGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: IamUserLoginProfileKind}.String()
	IamUserLoginProfileKindAPIVersion   = IamUserLoginProfileKind + "." + v1alpha1.GroupVersion.String()
	IamUserLoginProfileGroupVersionKind = v1alpha1.GroupVersion.WithKind(IamUserLoginProfileKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&IamUserLoginProfile{}, &IamUserLoginProfileList{})
}