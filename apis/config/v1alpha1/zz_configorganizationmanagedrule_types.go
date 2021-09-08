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

type ConfigOrganizationManagedRuleObservation struct {
	ARN string `json:"arn" tf:"arn"`
}

type ConfigOrganizationManagedRuleParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	ExcludedAccounts []string `json:"excludedAccounts,omitempty" tf:"excluded_accounts"`

	InputParameters *string `json:"inputParameters,omitempty" tf:"input_parameters"`

	MaximumExecutionFrequency *string `json:"maximumExecutionFrequency,omitempty" tf:"maximum_execution_frequency"`

	Name string `json:"name" tf:"name"`

	ResourceIDScope *string `json:"resourceIDScope,omitempty" tf:"resource_id_scope"`

	ResourceTypesScope []string `json:"resourceTypesScope,omitempty" tf:"resource_types_scope"`

	RuleIdentifier string `json:"ruleIdentifier" tf:"rule_identifier"`

	TagKeyScope *string `json:"tagKeyScope,omitempty" tf:"tag_key_scope"`

	TagValueScope *string `json:"tagValueScope,omitempty" tf:"tag_value_scope"`
}

// ConfigOrganizationManagedRuleSpec defines the desired state of ConfigOrganizationManagedRule
type ConfigOrganizationManagedRuleSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ConfigOrganizationManagedRuleParameters `json:"forProvider"`
}

// ConfigOrganizationManagedRuleStatus defines the observed state of ConfigOrganizationManagedRule.
type ConfigOrganizationManagedRuleStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ConfigOrganizationManagedRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigOrganizationManagedRule is the Schema for the ConfigOrganizationManagedRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ConfigOrganizationManagedRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigOrganizationManagedRuleSpec   `json:"spec"`
	Status            ConfigOrganizationManagedRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigOrganizationManagedRuleList contains a list of ConfigOrganizationManagedRules
type ConfigOrganizationManagedRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigOrganizationManagedRule `json:"items"`
}

// Repository type metadata.
var (
	ConfigOrganizationManagedRuleKind             = "ConfigOrganizationManagedRule"
	ConfigOrganizationManagedRuleGroupKind        = schema.GroupKind{Group: Group, Kind: ConfigOrganizationManagedRuleKind}.String()
	ConfigOrganizationManagedRuleKindAPIVersion   = ConfigOrganizationManagedRuleKind + "." + GroupVersion.String()
	ConfigOrganizationManagedRuleGroupVersionKind = GroupVersion.WithKind(ConfigOrganizationManagedRuleKind)
)

func init() {
	SchemeBuilder.Register(&ConfigOrganizationManagedRule{}, &ConfigOrganizationManagedRuleList{})
}
