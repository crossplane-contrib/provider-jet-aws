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

type GuarddutyOrganizationAdminAccountObservation struct {
}

type GuarddutyOrganizationAdminAccountParameters struct {
	AdminAccountID string `json:"adminAccountID" tf:"admin_account_id"`
}

// GuarddutyOrganizationAdminAccountSpec defines the desired state of GuarddutyOrganizationAdminAccount
type GuarddutyOrganizationAdminAccountSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       GuarddutyOrganizationAdminAccountParameters `json:"forProvider"`
}

// GuarddutyOrganizationAdminAccountStatus defines the observed state of GuarddutyOrganizationAdminAccount.
type GuarddutyOrganizationAdminAccountStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          GuarddutyOrganizationAdminAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GuarddutyOrganizationAdminAccount is the Schema for the GuarddutyOrganizationAdminAccounts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type GuarddutyOrganizationAdminAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GuarddutyOrganizationAdminAccountSpec   `json:"spec"`
	Status            GuarddutyOrganizationAdminAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GuarddutyOrganizationAdminAccountList contains a list of GuarddutyOrganizationAdminAccounts
type GuarddutyOrganizationAdminAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GuarddutyOrganizationAdminAccount `json:"items"`
}

// Repository type metadata.
var (
	GuarddutyOrganizationAdminAccountKind             = "GuarddutyOrganizationAdminAccount"
	GuarddutyOrganizationAdminAccountGroupKind        = schema.GroupKind{Group: Group, Kind: GuarddutyOrganizationAdminAccountKind}.String()
	GuarddutyOrganizationAdminAccountKindAPIVersion   = GuarddutyOrganizationAdminAccountKind + "." + GroupVersion.String()
	GuarddutyOrganizationAdminAccountGroupVersionKind = GroupVersion.WithKind(GuarddutyOrganizationAdminAccountKind)
)

func init() {
	SchemeBuilder.Register(&GuarddutyOrganizationAdminAccount{}, &GuarddutyOrganizationAdminAccountList{})
}
