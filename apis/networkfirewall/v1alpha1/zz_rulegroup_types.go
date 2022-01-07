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

type ActionDefinitionPublishMetricActionObservation struct {
}

type ActionDefinitionPublishMetricActionParameters struct {

	// +kubebuilder:validation:Required
	Dimension []PublishMetricActionDimensionParameters `json:"dimension" tf:"dimension,omitempty"`
}

type CustomActionActionDefinitionObservation struct {
}

type CustomActionActionDefinitionParameters struct {

	// +kubebuilder:validation:Required
	PublishMetricAction []ActionDefinitionPublishMetricActionParameters `json:"publishMetricAction" tf:"publish_metric_action,omitempty"`
}

type CustomActionObservation struct {
}

type CustomActionParameters struct {

	// +kubebuilder:validation:Required
	ActionDefinition []CustomActionActionDefinitionParameters `json:"actionDefinition" tf:"action_definition,omitempty"`

	// +kubebuilder:validation:Required
	ActionName *string `json:"actionName" tf:"action_name,omitempty"`
}

type DestinationObservation struct {
}

type DestinationParameters struct {

	// +kubebuilder:validation:Required
	AddressDefinition *string `json:"addressDefinition" tf:"address_definition,omitempty"`
}

type DestinationPortObservation struct {
}

type DestinationPortParameters struct {

	// +kubebuilder:validation:Required
	FromPort *int64 `json:"fromPort" tf:"from_port,omitempty"`

	// +kubebuilder:validation:Optional
	ToPort *int64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type HeaderObservation struct {
}

type HeaderParameters struct {

	// +kubebuilder:validation:Required
	Destination *string `json:"destination" tf:"destination,omitempty"`

	// +kubebuilder:validation:Required
	DestinationPort *string `json:"destinationPort" tf:"destination_port,omitempty"`

	// +kubebuilder:validation:Required
	Direction *string `json:"direction" tf:"direction,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Required
	Source *string `json:"source" tf:"source,omitempty"`

	// +kubebuilder:validation:Required
	SourcePort *string `json:"sourcePort" tf:"source_port,omitempty"`
}

type IPSetObservation struct {
}

type IPSetParameters struct {

	// +kubebuilder:validation:Required
	Definition []*string `json:"definition" tf:"definition,omitempty"`
}

type IPSetsObservation struct {
}

type IPSetsParameters struct {

	// +kubebuilder:validation:Required
	IPSet []IPSetParameters `json:"ipSet" tf:"ip_set,omitempty"`

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`
}

type MatchAttributesObservation struct {
}

type MatchAttributesParameters struct {

	// +kubebuilder:validation:Optional
	Destination []DestinationParameters `json:"destination,omitempty" tf:"destination,omitempty"`

	// +kubebuilder:validation:Optional
	DestinationPort []DestinationPortParameters `json:"destinationPort,omitempty" tf:"destination_port,omitempty"`

	// +kubebuilder:validation:Optional
	Protocols []*int64 `json:"protocols,omitempty" tf:"protocols,omitempty"`

	// +kubebuilder:validation:Optional
	Source []SourceParameters `json:"source,omitempty" tf:"source,omitempty"`

	// +kubebuilder:validation:Optional
	SourcePort []SourcePortParameters `json:"sourcePort,omitempty" tf:"source_port,omitempty"`

	// +kubebuilder:validation:Optional
	TCPFlag []TCPFlagParameters `json:"tcpFlag,omitempty" tf:"tcp_flag,omitempty"`
}

type PortSetObservation struct {
}

type PortSetParameters struct {

	// +kubebuilder:validation:Required
	Definition []*string `json:"definition" tf:"definition,omitempty"`
}

type PortSetsObservation struct {
}

type PortSetsParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	PortSet []PortSetParameters `json:"portSet" tf:"port_set,omitempty"`
}

type PublishMetricActionDimensionObservation struct {
}

type PublishMetricActionDimensionParameters struct {

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type RuleDefinitionObservation struct {
}

type RuleDefinitionParameters struct {

	// +kubebuilder:validation:Required
	Actions []*string `json:"actions" tf:"actions,omitempty"`

	// +kubebuilder:validation:Required
	MatchAttributes []MatchAttributesParameters `json:"matchAttributes" tf:"match_attributes,omitempty"`
}

type RuleGroupObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	UpdateToken *string `json:"updateToken,omitempty" tf:"update_token,omitempty"`
}

type RuleGroupParameters struct {

	// +kubebuilder:validation:Required
	Capacity *int64 `json:"capacity" tf:"capacity,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RuleGroup []RuleGroupRuleGroupParameters `json:"ruleGroup,omitempty" tf:"rule_group,omitempty"`

	// +kubebuilder:validation:Optional
	Rules *string `json:"rules,omitempty" tf:"rules,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type RuleGroupRuleGroupObservation struct {
}

type RuleGroupRuleGroupParameters struct {

	// +kubebuilder:validation:Optional
	RuleVariables []RuleVariablesParameters `json:"ruleVariables,omitempty" tf:"rule_variables,omitempty"`

	// +kubebuilder:validation:Required
	RulesSource []RulesSourceParameters `json:"rulesSource" tf:"rules_source,omitempty"`
}

type RuleOptionObservation struct {
}

type RuleOptionParameters struct {

	// +kubebuilder:validation:Required
	Keyword *string `json:"keyword" tf:"keyword,omitempty"`

	// +kubebuilder:validation:Optional
	Settings []*string `json:"settings,omitempty" tf:"settings,omitempty"`
}

type RuleVariablesObservation struct {
}

type RuleVariablesParameters struct {

	// +kubebuilder:validation:Optional
	IPSets []IPSetsParameters `json:"ipSets,omitempty" tf:"ip_sets,omitempty"`

	// +kubebuilder:validation:Optional
	PortSets []PortSetsParameters `json:"portSets,omitempty" tf:"port_sets,omitempty"`
}

type RulesSourceListObservation struct {
}

type RulesSourceListParameters struct {

	// +kubebuilder:validation:Required
	GeneratedRulesType *string `json:"generatedRulesType" tf:"generated_rules_type,omitempty"`

	// +kubebuilder:validation:Required
	TargetTypes []*string `json:"targetTypes" tf:"target_types,omitempty"`

	// +kubebuilder:validation:Required
	Targets []*string `json:"targets" tf:"targets,omitempty"`
}

type RulesSourceObservation struct {
}

type RulesSourceParameters struct {

	// +kubebuilder:validation:Optional
	RulesSourceList []RulesSourceListParameters `json:"rulesSourceList,omitempty" tf:"rules_source_list,omitempty"`

	// +kubebuilder:validation:Optional
	RulesString *string `json:"rulesString,omitempty" tf:"rules_string,omitempty"`

	// +kubebuilder:validation:Optional
	StatefulRule []StatefulRuleParameters `json:"statefulRule,omitempty" tf:"stateful_rule,omitempty"`

	// +kubebuilder:validation:Optional
	StatelessRulesAndCustomActions []StatelessRulesAndCustomActionsParameters `json:"statelessRulesAndCustomActions,omitempty" tf:"stateless_rules_and_custom_actions,omitempty"`
}

type SourceObservation struct {
}

type SourceParameters struct {

	// +kubebuilder:validation:Required
	AddressDefinition *string `json:"addressDefinition" tf:"address_definition,omitempty"`
}

type SourcePortObservation struct {
}

type SourcePortParameters struct {

	// +kubebuilder:validation:Required
	FromPort *int64 `json:"fromPort" tf:"from_port,omitempty"`

	// +kubebuilder:validation:Optional
	ToPort *int64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type StatefulRuleObservation struct {
}

type StatefulRuleParameters struct {

	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// +kubebuilder:validation:Required
	Header []HeaderParameters `json:"header" tf:"header,omitempty"`

	// +kubebuilder:validation:Required
	RuleOption []RuleOptionParameters `json:"ruleOption" tf:"rule_option,omitempty"`
}

type StatelessRuleObservation struct {
}

type StatelessRuleParameters struct {

	// +kubebuilder:validation:Required
	Priority *int64 `json:"priority" tf:"priority,omitempty"`

	// +kubebuilder:validation:Required
	RuleDefinition []RuleDefinitionParameters `json:"ruleDefinition" tf:"rule_definition,omitempty"`
}

type StatelessRulesAndCustomActionsObservation struct {
}

type StatelessRulesAndCustomActionsParameters struct {

	// +kubebuilder:validation:Optional
	CustomAction []CustomActionParameters `json:"customAction,omitempty" tf:"custom_action,omitempty"`

	// +kubebuilder:validation:Required
	StatelessRule []StatelessRuleParameters `json:"statelessRule" tf:"stateless_rule,omitempty"`
}

type TCPFlagObservation struct {
}

type TCPFlagParameters struct {

	// +kubebuilder:validation:Required
	Flags []*string `json:"flags" tf:"flags,omitempty"`

	// +kubebuilder:validation:Optional
	Masks []*string `json:"masks,omitempty" tf:"masks,omitempty"`
}

// RuleGroupSpec defines the desired state of RuleGroup
type RuleGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RuleGroupParameters `json:"forProvider"`
}

// RuleGroupStatus defines the observed state of RuleGroup.
type RuleGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RuleGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RuleGroup is the Schema for the RuleGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type RuleGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RuleGroupSpec   `json:"spec"`
	Status            RuleGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RuleGroupList contains a list of RuleGroups
type RuleGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RuleGroup `json:"items"`
}

// Repository type metadata.
var (
	RuleGroup_Kind             = "RuleGroup"
	RuleGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RuleGroup_Kind}.String()
	RuleGroup_KindAPIVersion   = RuleGroup_Kind + "." + CRDGroupVersion.String()
	RuleGroup_GroupVersionKind = CRDGroupVersion.WithKind(RuleGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&RuleGroup{}, &RuleGroupList{})
}
