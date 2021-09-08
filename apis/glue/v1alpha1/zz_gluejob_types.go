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

type CommandObservation struct {
}

type CommandParameters struct {
	Name *string `json:"name,omitempty" tf:"name"`

	PythonVersion *string `json:"pythonVersion,omitempty" tf:"python_version"`

	ScriptLocation string `json:"scriptLocation" tf:"script_location"`
}

type ExecutionPropertyObservation struct {
}

type ExecutionPropertyParameters struct {
	MaxConcurrentRuns *int64 `json:"maxConcurrentRuns,omitempty" tf:"max_concurrent_runs"`
}

type GlueJobObservation struct {
	ARN string `json:"arn" tf:"arn"`
}

type GlueJobParameters struct {
	Command []CommandParameters `json:"command" tf:"command"`

	Connections []string `json:"connections,omitempty" tf:"connections"`

	DefaultArguments map[string]string `json:"defaultArguments,omitempty" tf:"default_arguments"`

	Description *string `json:"description,omitempty" tf:"description"`

	ExecutionProperty []ExecutionPropertyParameters `json:"executionProperty,omitempty" tf:"execution_property"`

	GlueVersion *string `json:"glueVersion,omitempty" tf:"glue_version"`

	MaxCapacity *float64 `json:"maxCapacity,omitempty" tf:"max_capacity"`

	MaxRetries *int64 `json:"maxRetries,omitempty" tf:"max_retries"`

	Name string `json:"name" tf:"name"`

	NonOverridableArguments map[string]string `json:"nonOverridableArguments,omitempty" tf:"non_overridable_arguments"`

	NotificationProperty []NotificationPropertyParameters `json:"notificationProperty,omitempty" tf:"notification_property"`

	NumberOfWorkers *int64 `json:"numberOfWorkers,omitempty" tf:"number_of_workers"`

	RoleARN string `json:"roleARN" tf:"role_arn"`

	SecurityConfiguration *string `json:"securityConfiguration,omitempty" tf:"security_configuration"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	Timeout *int64 `json:"timeout,omitempty" tf:"timeout"`

	WorkerType *string `json:"workerType,omitempty" tf:"worker_type"`
}

type NotificationPropertyObservation struct {
}

type NotificationPropertyParameters struct {
	NotifyDelayAfter *int64 `json:"notifyDelayAfter,omitempty" tf:"notify_delay_after"`
}

// GlueJobSpec defines the desired state of GlueJob
type GlueJobSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       GlueJobParameters `json:"forProvider"`
}

// GlueJobStatus defines the observed state of GlueJob.
type GlueJobStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          GlueJobObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GlueJob is the Schema for the GlueJobs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type GlueJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlueJobSpec   `json:"spec"`
	Status            GlueJobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlueJobList contains a list of GlueJobs
type GlueJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlueJob `json:"items"`
}

// Repository type metadata.
var (
	GlueJobKind             = "GlueJob"
	GlueJobGroupKind        = schema.GroupKind{Group: Group, Kind: GlueJobKind}.String()
	GlueJobKindAPIVersion   = GlueJobKind + "." + GroupVersion.String()
	GlueJobGroupVersionKind = GroupVersion.WithKind(GlueJobKind)
)

func init() {
	SchemeBuilder.Register(&GlueJob{}, &GlueJobList{})
}
