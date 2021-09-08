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

type CodeHookObservation struct {
}

type CodeHookParameters struct {
	MessageVersion string `json:"messageVersion" tf:"message_version"`

	URI string `json:"uri" tf:"uri"`
}

type ConclusionStatementMessageObservation struct {
}

type ConclusionStatementMessageParameters struct {
	Content string `json:"content" tf:"content"`

	ContentType string `json:"contentType" tf:"content_type"`

	GroupNumber *int64 `json:"groupNumber,omitempty" tf:"group_number"`
}

type ConclusionStatementObservation struct {
}

type ConclusionStatementParameters struct {
	Message []ConclusionStatementMessageParameters `json:"message" tf:"message"`

	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card"`
}

type ConfirmationPromptMessageObservation struct {
}

type ConfirmationPromptMessageParameters struct {
	Content string `json:"content" tf:"content"`

	ContentType string `json:"contentType" tf:"content_type"`

	GroupNumber *int64 `json:"groupNumber,omitempty" tf:"group_number"`
}

type ConfirmationPromptObservation struct {
}

type ConfirmationPromptParameters struct {
	MaxAttempts int64 `json:"maxAttempts" tf:"max_attempts"`

	Message []ConfirmationPromptMessageParameters `json:"message" tf:"message"`

	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card"`
}

type DialogCodeHookObservation struct {
}

type DialogCodeHookParameters struct {
	MessageVersion string `json:"messageVersion" tf:"message_version"`

	URI string `json:"uri" tf:"uri"`
}

type FollowUpPromptObservation struct {
}

type FollowUpPromptParameters struct {
	Prompt []PromptParameters `json:"prompt" tf:"prompt"`

	RejectionStatement []RejectionStatementParameters `json:"rejectionStatement" tf:"rejection_statement"`
}

type FulfillmentActivityObservation struct {
}

type FulfillmentActivityParameters struct {
	CodeHook []CodeHookParameters `json:"codeHook,omitempty" tf:"code_hook"`

	Type string `json:"type" tf:"type"`
}

type LexIntentObservation struct {
	ARN string `json:"arn" tf:"arn"`

	Checksum string `json:"checksum" tf:"checksum"`

	CreatedDate string `json:"createdDate" tf:"created_date"`

	LastUpdatedDate string `json:"lastUpdatedDate" tf:"last_updated_date"`

	Version string `json:"version" tf:"version"`
}

type LexIntentParameters struct {
	ConclusionStatement []ConclusionStatementParameters `json:"conclusionStatement,omitempty" tf:"conclusion_statement"`

	ConfirmationPrompt []ConfirmationPromptParameters `json:"confirmationPrompt,omitempty" tf:"confirmation_prompt"`

	CreateVersion *bool `json:"createVersion,omitempty" tf:"create_version"`

	Description *string `json:"description,omitempty" tf:"description"`

	DialogCodeHook []DialogCodeHookParameters `json:"dialogCodeHook,omitempty" tf:"dialog_code_hook"`

	FollowUpPrompt []FollowUpPromptParameters `json:"followUpPrompt,omitempty" tf:"follow_up_prompt"`

	FulfillmentActivity []FulfillmentActivityParameters `json:"fulfillmentActivity" tf:"fulfillment_activity"`

	Name string `json:"name" tf:"name"`

	ParentIntentSignature *string `json:"parentIntentSignature,omitempty" tf:"parent_intent_signature"`

	RejectionStatement []LexIntentRejectionStatementParameters `json:"rejectionStatement,omitempty" tf:"rejection_statement"`

	SampleUtterances []string `json:"sampleUtterances,omitempty" tf:"sample_utterances"`

	Slot []SlotParameters `json:"slot,omitempty" tf:"slot"`
}

type LexIntentRejectionStatementMessageObservation struct {
}

type LexIntentRejectionStatementMessageParameters struct {
	Content string `json:"content" tf:"content"`

	ContentType string `json:"contentType" tf:"content_type"`

	GroupNumber *int64 `json:"groupNumber,omitempty" tf:"group_number"`
}

type LexIntentRejectionStatementObservation struct {
}

type LexIntentRejectionStatementParameters struct {
	Message []LexIntentRejectionStatementMessageParameters `json:"message" tf:"message"`

	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card"`
}

type PromptMessageObservation struct {
}

type PromptMessageParameters struct {
	Content string `json:"content" tf:"content"`

	ContentType string `json:"contentType" tf:"content_type"`

	GroupNumber *int64 `json:"groupNumber,omitempty" tf:"group_number"`
}

type PromptObservation struct {
}

type PromptParameters struct {
	MaxAttempts int64 `json:"maxAttempts" tf:"max_attempts"`

	Message []PromptMessageParameters `json:"message" tf:"message"`

	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card"`
}

type RejectionStatementMessageObservation struct {
}

type RejectionStatementMessageParameters struct {
	Content string `json:"content" tf:"content"`

	ContentType string `json:"contentType" tf:"content_type"`

	GroupNumber *int64 `json:"groupNumber,omitempty" tf:"group_number"`
}

type RejectionStatementObservation struct {
}

type RejectionStatementParameters struct {
	Message []RejectionStatementMessageParameters `json:"message" tf:"message"`

	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card"`
}

type SlotObservation struct {
}

type SlotParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	Name string `json:"name" tf:"name"`

	Priority *int64 `json:"priority,omitempty" tf:"priority"`

	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card"`

	SampleUtterances []string `json:"sampleUtterances,omitempty" tf:"sample_utterances"`

	SlotConstraint string `json:"slotConstraint" tf:"slot_constraint"`

	SlotType string `json:"slotType" tf:"slot_type"`

	SlotTypeVersion *string `json:"slotTypeVersion,omitempty" tf:"slot_type_version"`

	ValueElicitationPrompt []ValueElicitationPromptParameters `json:"valueElicitationPrompt,omitempty" tf:"value_elicitation_prompt"`
}

type ValueElicitationPromptMessageObservation struct {
}

type ValueElicitationPromptMessageParameters struct {
	Content string `json:"content" tf:"content"`

	ContentType string `json:"contentType" tf:"content_type"`

	GroupNumber *int64 `json:"groupNumber,omitempty" tf:"group_number"`
}

type ValueElicitationPromptObservation struct {
}

type ValueElicitationPromptParameters struct {
	MaxAttempts int64 `json:"maxAttempts" tf:"max_attempts"`

	Message []ValueElicitationPromptMessageParameters `json:"message" tf:"message"`

	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card"`
}

// LexIntentSpec defines the desired state of LexIntent
type LexIntentSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       LexIntentParameters `json:"forProvider"`
}

// LexIntentStatus defines the observed state of LexIntent.
type LexIntentStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          LexIntentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LexIntent is the Schema for the LexIntents API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type LexIntent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LexIntentSpec   `json:"spec"`
	Status            LexIntentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LexIntentList contains a list of LexIntents
type LexIntentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LexIntent `json:"items"`
}

// Repository type metadata.
var (
	LexIntentKind             = "LexIntent"
	LexIntentGroupKind        = schema.GroupKind{Group: Group, Kind: LexIntentKind}.String()
	LexIntentKindAPIVersion   = LexIntentKind + "." + GroupVersion.String()
	LexIntentGroupVersionKind = GroupVersion.WithKind(LexIntentKind)
)

func init() {
	SchemeBuilder.Register(&LexIntent{}, &LexIntentList{})
}
