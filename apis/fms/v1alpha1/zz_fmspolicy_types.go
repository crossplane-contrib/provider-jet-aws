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

type ExcludeMapObservation struct {
}

type ExcludeMapParameters struct {
	Account []string `json:"account,omitempty" tf:"account"`

	Orgunit []string `json:"orgunit,omitempty" tf:"orgunit"`
}

type FmsPolicyObservation struct {
	ARN string `json:"arn" tf:"arn"`

	PolicyUpdateToken string `json:"policyUpdateToken" tf:"policy_update_token"`
}

type FmsPolicyParameters struct {
	DeleteAllPolicyResources *bool `json:"deleteAllPolicyResources,omitempty" tf:"delete_all_policy_resources"`

	ExcludeMap []ExcludeMapParameters `json:"excludeMap,omitempty" tf:"exclude_map"`

	ExcludeResourceTags bool `json:"excludeResourceTags" tf:"exclude_resource_tags"`

	IncludeMap []IncludeMapParameters `json:"includeMap,omitempty" tf:"include_map"`

	Name string `json:"name" tf:"name"`

	RemediationEnabled *bool `json:"remediationEnabled,omitempty" tf:"remediation_enabled"`

	ResourceTags map[string]string `json:"resourceTags,omitempty" tf:"resource_tags"`

	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type"`

	ResourceTypeList []string `json:"resourceTypeList,omitempty" tf:"resource_type_list"`

	SecurityServicePolicyData []SecurityServicePolicyDataParameters `json:"securityServicePolicyData" tf:"security_service_policy_data"`
}

type IncludeMapObservation struct {
}

type IncludeMapParameters struct {
	Account []string `json:"account,omitempty" tf:"account"`

	Orgunit []string `json:"orgunit,omitempty" tf:"orgunit"`
}

type SecurityServicePolicyDataObservation struct {
}

type SecurityServicePolicyDataParameters struct {
	ManagedServiceData *string `json:"managedServiceData,omitempty" tf:"managed_service_data"`

	Type string `json:"type" tf:"type"`
}

// FmsPolicySpec defines the desired state of FmsPolicy
type FmsPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       FmsPolicyParameters `json:"forProvider"`
}

// FmsPolicyStatus defines the observed state of FmsPolicy.
type FmsPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          FmsPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FmsPolicy is the Schema for the FmsPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type FmsPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FmsPolicySpec   `json:"spec"`
	Status            FmsPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FmsPolicyList contains a list of FmsPolicys
type FmsPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FmsPolicy `json:"items"`
}

// Repository type metadata.
var (
	FmsPolicyKind             = "FmsPolicy"
	FmsPolicyGroupKind        = schema.GroupKind{Group: Group, Kind: FmsPolicyKind}.String()
	FmsPolicyKindAPIVersion   = FmsPolicyKind + "." + GroupVersion.String()
	FmsPolicyGroupVersionKind = GroupVersion.WithKind(FmsPolicyKind)
)

func init() {
	SchemeBuilder.Register(&FmsPolicy{}, &FmsPolicyList{})
}
