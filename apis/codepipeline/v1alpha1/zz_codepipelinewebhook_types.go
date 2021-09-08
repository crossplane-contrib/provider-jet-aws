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

type AuthenticationConfigurationObservation struct {
}

type AuthenticationConfigurationParameters struct {
	AllowedIPRange *string `json:"allowedIPRange,omitempty" tf:"allowed_ip_range"`

	SecretToken *string `json:"secretToken,omitempty" tf:"secret_token"`
}

type CodepipelineWebhookObservation struct {
	URL string `json:"url" tf:"url"`
}

type CodepipelineWebhookParameters struct {
	Authentication string `json:"authentication" tf:"authentication"`

	AuthenticationConfiguration []AuthenticationConfigurationParameters `json:"authenticationConfiguration,omitempty" tf:"authentication_configuration"`

	Filter []FilterParameters `json:"filter" tf:"filter"`

	Name string `json:"name" tf:"name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	TargetAction string `json:"targetAction" tf:"target_action"`

	TargetPipeline string `json:"targetPipeline" tf:"target_pipeline"`
}

type FilterObservation struct {
}

type FilterParameters struct {
	JSONPath string `json:"jsonPath" tf:"json_path"`

	MatchEquals string `json:"matchEquals" tf:"match_equals"`
}

// CodepipelineWebhookSpec defines the desired state of CodepipelineWebhook
type CodepipelineWebhookSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CodepipelineWebhookParameters `json:"forProvider"`
}

// CodepipelineWebhookStatus defines the observed state of CodepipelineWebhook.
type CodepipelineWebhookStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CodepipelineWebhookObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CodepipelineWebhook is the Schema for the CodepipelineWebhooks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type CodepipelineWebhook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodepipelineWebhookSpec   `json:"spec"`
	Status            CodepipelineWebhookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodepipelineWebhookList contains a list of CodepipelineWebhooks
type CodepipelineWebhookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodepipelineWebhook `json:"items"`
}

// Repository type metadata.
var (
	CodepipelineWebhookKind             = "CodepipelineWebhook"
	CodepipelineWebhookGroupKind        = schema.GroupKind{Group: Group, Kind: CodepipelineWebhookKind}.String()
	CodepipelineWebhookKindAPIVersion   = CodepipelineWebhookKind + "." + GroupVersion.String()
	CodepipelineWebhookGroupVersionKind = GroupVersion.WithKind(CodepipelineWebhookKind)
)

func init() {
	SchemeBuilder.Register(&CodepipelineWebhook{}, &CodepipelineWebhookList{})
}
