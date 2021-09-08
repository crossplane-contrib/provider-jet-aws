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

type ConfigOrganizationCustomRuleObservation struct {
	ARN string `json:"arn" tf:"arn"`
}

type ConfigOrganizationCustomRuleParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	ExcludedAccounts []string `json:"excludedAccounts,omitempty" tf:"excluded_accounts"`

	InputParameters *string `json:"inputParameters,omitempty" tf:"input_parameters"`

	LambdaFunctionARN string `json:"lambdaFunctionARN" tf:"lambda_function_arn"`

	MaximumExecutionFrequency *string `json:"maximumExecutionFrequency,omitempty" tf:"maximum_execution_frequency"`

	Name string `json:"name" tf:"name"`

	ResourceIDScope *string `json:"resourceIDScope,omitempty" tf:"resource_id_scope"`

	ResourceTypesScope []string `json:"resourceTypesScope,omitempty" tf:"resource_types_scope"`

	TagKeyScope *string `json:"tagKeyScope,omitempty" tf:"tag_key_scope"`

	TagValueScope *string `json:"tagValueScope,omitempty" tf:"tag_value_scope"`

	TriggerTypes []string `json:"triggerTypes" tf:"trigger_types"`
}

// ConfigOrganizationCustomRuleSpec defines the desired state of ConfigOrganizationCustomRule
type ConfigOrganizationCustomRuleSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ConfigOrganizationCustomRuleParameters `json:"forProvider"`
}

// ConfigOrganizationCustomRuleStatus defines the observed state of ConfigOrganizationCustomRule.
type ConfigOrganizationCustomRuleStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ConfigOrganizationCustomRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigOrganizationCustomRule is the Schema for the ConfigOrganizationCustomRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ConfigOrganizationCustomRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigOrganizationCustomRuleSpec   `json:"spec"`
	Status            ConfigOrganizationCustomRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigOrganizationCustomRuleList contains a list of ConfigOrganizationCustomRules
type ConfigOrganizationCustomRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigOrganizationCustomRule `json:"items"`
}

// Repository type metadata.
var (
	ConfigOrganizationCustomRuleKind             = "ConfigOrganizationCustomRule"
	ConfigOrganizationCustomRuleGroupKind        = schema.GroupKind{Group: Group, Kind: ConfigOrganizationCustomRuleKind}.String()
	ConfigOrganizationCustomRuleKindAPIVersion   = ConfigOrganizationCustomRuleKind + "." + GroupVersion.String()
	ConfigOrganizationCustomRuleGroupVersionKind = GroupVersion.WithKind(ConfigOrganizationCustomRuleKind)
)

func init() {
	SchemeBuilder.Register(&ConfigOrganizationCustomRule{}, &ConfigOrganizationCustomRuleList{})
}
