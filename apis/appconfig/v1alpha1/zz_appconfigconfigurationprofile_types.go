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

type AppconfigConfigurationProfileObservation struct {
	ARN string `json:"arn" tf:"arn"`

	ConfigurationProfileID string `json:"configurationProfileID" tf:"configuration_profile_id"`
}

type AppconfigConfigurationProfileParameters struct {
	ApplicationID string `json:"applicationID" tf:"application_id"`

	Description *string `json:"description,omitempty" tf:"description"`

	LocationURI string `json:"locationURI" tf:"location_uri"`

	Name string `json:"name" tf:"name"`

	RetrievalRoleARN *string `json:"retrievalRoleARN,omitempty" tf:"retrieval_role_arn"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	Validator []ValidatorParameters `json:"validator,omitempty" tf:"validator"`
}

type ValidatorObservation struct {
}

type ValidatorParameters struct {
	Content *string `json:"content,omitempty" tf:"content"`

	Type string `json:"type" tf:"type"`
}

// AppconfigConfigurationProfileSpec defines the desired state of AppconfigConfigurationProfile
type AppconfigConfigurationProfileSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AppconfigConfigurationProfileParameters `json:"forProvider"`
}

// AppconfigConfigurationProfileStatus defines the observed state of AppconfigConfigurationProfile.
type AppconfigConfigurationProfileStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AppconfigConfigurationProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppconfigConfigurationProfile is the Schema for the AppconfigConfigurationProfiles API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type AppconfigConfigurationProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppconfigConfigurationProfileSpec   `json:"spec"`
	Status            AppconfigConfigurationProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppconfigConfigurationProfileList contains a list of AppconfigConfigurationProfiles
type AppconfigConfigurationProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppconfigConfigurationProfile `json:"items"`
}

// Repository type metadata.
var (
	AppconfigConfigurationProfileKind             = "AppconfigConfigurationProfile"
	AppconfigConfigurationProfileGroupKind        = schema.GroupKind{Group: Group, Kind: AppconfigConfigurationProfileKind}.String()
	AppconfigConfigurationProfileKindAPIVersion   = AppconfigConfigurationProfileKind + "." + GroupVersion.String()
	AppconfigConfigurationProfileGroupVersionKind = GroupVersion.WithKind(AppconfigConfigurationProfileKind)
)

func init() {
	SchemeBuilder.Register(&AppconfigConfigurationProfile{}, &AppconfigConfigurationProfileList{})
}
