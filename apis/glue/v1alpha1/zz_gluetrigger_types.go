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

type ActionsNotificationPropertyObservation struct {
}

type ActionsNotificationPropertyParameters struct {
	NotifyDelayAfter *int64 `json:"notifyDelayAfter,omitempty" tf:"notify_delay_after"`
}

type ActionsObservation struct {
}

type ActionsParameters struct {
	Arguments map[string]string `json:"arguments,omitempty" tf:"arguments"`

	CrawlerName *string `json:"crawlerName,omitempty" tf:"crawler_name"`

	JobName *string `json:"jobName,omitempty" tf:"job_name"`

	NotificationProperty []ActionsNotificationPropertyParameters `json:"notificationProperty,omitempty" tf:"notification_property"`

	SecurityConfiguration *string `json:"securityConfiguration,omitempty" tf:"security_configuration"`

	Timeout *int64 `json:"timeout,omitempty" tf:"timeout"`
}

type ConditionsObservation struct {
}

type ConditionsParameters struct {
	CrawlState *string `json:"crawlState,omitempty" tf:"crawl_state"`

	CrawlerName *string `json:"crawlerName,omitempty" tf:"crawler_name"`

	JobName *string `json:"jobName,omitempty" tf:"job_name"`

	LogicalOperator *string `json:"logicalOperator,omitempty" tf:"logical_operator"`

	State *string `json:"state,omitempty" tf:"state"`
}

type GlueTriggerObservation struct {
	ARN string `json:"arn" tf:"arn"`

	State string `json:"state" tf:"state"`
}

type GlueTriggerParameters struct {
	Actions []ActionsParameters `json:"actions" tf:"actions"`

	Description *string `json:"description,omitempty" tf:"description"`

	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	Name string `json:"name" tf:"name"`

	Predicate []PredicateParameters `json:"predicate,omitempty" tf:"predicate"`

	Schedule *string `json:"schedule,omitempty" tf:"schedule"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	Type string `json:"type" tf:"type"`

	WorkflowName *string `json:"workflowName,omitempty" tf:"workflow_name"`
}

type PredicateObservation struct {
}

type PredicateParameters struct {
	Conditions []ConditionsParameters `json:"conditions" tf:"conditions"`

	Logical *string `json:"logical,omitempty" tf:"logical"`
}

// GlueTriggerSpec defines the desired state of GlueTrigger
type GlueTriggerSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       GlueTriggerParameters `json:"forProvider"`
}

// GlueTriggerStatus defines the observed state of GlueTrigger.
type GlueTriggerStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          GlueTriggerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GlueTrigger is the Schema for the GlueTriggers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type GlueTrigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlueTriggerSpec   `json:"spec"`
	Status            GlueTriggerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlueTriggerList contains a list of GlueTriggers
type GlueTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlueTrigger `json:"items"`
}

// Repository type metadata.
var (
	GlueTriggerKind             = "GlueTrigger"
	GlueTriggerGroupKind        = schema.GroupKind{Group: Group, Kind: GlueTriggerKind}.String()
	GlueTriggerKindAPIVersion   = GlueTriggerKind + "." + GroupVersion.String()
	GlueTriggerGroupVersionKind = GroupVersion.WithKind(GlueTriggerKind)
)

func init() {
	SchemeBuilder.Register(&GlueTrigger{}, &GlueTriggerList{})
}
