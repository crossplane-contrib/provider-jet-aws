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

type AndObservation struct {
}

type AndParameters struct {
	SimpleScopeTerm []SimpleScopeTermParameters `json:"simpleScopeTerm,omitempty" tf:"simple_scope_term"`

	TagScopeTerm []TagScopeTermParameters `json:"tagScopeTerm,omitempty" tf:"tag_scope_term"`
}

type AndSimpleScopeTermObservation struct {
}

type AndSimpleScopeTermParameters struct {
	Comparator *string `json:"comparator,omitempty" tf:"comparator"`

	Key *string `json:"key,omitempty" tf:"key"`

	Values []string `json:"values,omitempty" tf:"values"`
}

type AndTagScopeTermObservation struct {
}

type AndTagScopeTermParameters struct {
	Comparator *string `json:"comparator,omitempty" tf:"comparator"`

	Key *string `json:"key,omitempty" tf:"key"`

	TagValues []TagScopeTermTagValuesParameters `json:"tagValues,omitempty" tf:"tag_values"`

	Target *string `json:"target,omitempty" tf:"target"`
}

type BucketDefinitionsObservation struct {
}

type BucketDefinitionsParameters struct {
	AccountID string `json:"accountID" tf:"account_id"`

	Buckets []string `json:"buckets" tf:"buckets"`
}

type ExcludesObservation struct {
}

type ExcludesParameters struct {
	And []AndParameters `json:"and,omitempty" tf:"and"`
}

type IncludesAndObservation struct {
}

type IncludesAndParameters struct {
	SimpleScopeTerm []AndSimpleScopeTermParameters `json:"simpleScopeTerm,omitempty" tf:"simple_scope_term"`

	TagScopeTerm []AndTagScopeTermParameters `json:"tagScopeTerm,omitempty" tf:"tag_scope_term"`
}

type IncludesObservation struct {
}

type IncludesParameters struct {
	And []IncludesAndParameters `json:"and,omitempty" tf:"and"`
}

type Macie2ClassificationJobObservation struct {
	CreatedAt string `json:"createdAt" tf:"created_at"`

	JobARN string `json:"jobARN" tf:"job_arn"`

	JobID string `json:"jobID" tf:"job_id"`

	UserPausedDetails []UserPausedDetailsObservation `json:"userPausedDetails" tf:"user_paused_details"`
}

type Macie2ClassificationJobParameters struct {
	CustomDataIdentifierIds []string `json:"customDataIdentifierIds,omitempty" tf:"custom_data_identifier_ids"`

	Description *string `json:"description,omitempty" tf:"description"`

	InitialRun *bool `json:"initialRun,omitempty" tf:"initial_run"`

	JobStatus *string `json:"jobStatus,omitempty" tf:"job_status"`

	JobType string `json:"jobType" tf:"job_type"`

	Name *string `json:"name,omitempty" tf:"name"`

	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`

	S3JobDefinition []S3JobDefinitionParameters `json:"s3JobDefinition" tf:"s3_job_definition"`

	SamplingPercentage *int64 `json:"samplingPercentage,omitempty" tf:"sampling_percentage"`

	ScheduleFrequency []ScheduleFrequencyParameters `json:"scheduleFrequency,omitempty" tf:"schedule_frequency"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type S3JobDefinitionObservation struct {
}

type S3JobDefinitionParameters struct {
	BucketDefinitions []BucketDefinitionsParameters `json:"bucketDefinitions,omitempty" tf:"bucket_definitions"`

	Scoping []ScopingParameters `json:"scoping,omitempty" tf:"scoping"`
}

type ScheduleFrequencyObservation struct {
}

type ScheduleFrequencyParameters struct {
	DailySchedule *bool `json:"dailySchedule,omitempty" tf:"daily_schedule"`

	MonthlySchedule *int64 `json:"monthlySchedule,omitempty" tf:"monthly_schedule"`

	WeeklySchedule *string `json:"weeklySchedule,omitempty" tf:"weekly_schedule"`
}

type ScopingObservation struct {
}

type ScopingParameters struct {
	Excludes []ExcludesParameters `json:"excludes,omitempty" tf:"excludes"`

	Includes []IncludesParameters `json:"includes,omitempty" tf:"includes"`
}

type SimpleScopeTermObservation struct {
}

type SimpleScopeTermParameters struct {
	Comparator *string `json:"comparator,omitempty" tf:"comparator"`

	Key *string `json:"key,omitempty" tf:"key"`

	Values []string `json:"values,omitempty" tf:"values"`
}

type TagScopeTermObservation struct {
}

type TagScopeTermParameters struct {
	Comparator *string `json:"comparator,omitempty" tf:"comparator"`

	Key *string `json:"key,omitempty" tf:"key"`

	TagValues []TagValuesParameters `json:"tagValues,omitempty" tf:"tag_values"`

	Target *string `json:"target,omitempty" tf:"target"`
}

type TagScopeTermTagValuesObservation struct {
}

type TagScopeTermTagValuesParameters struct {
	Key *string `json:"key,omitempty" tf:"key"`

	Value *string `json:"value,omitempty" tf:"value"`
}

type TagValuesObservation struct {
}

type TagValuesParameters struct {
	Key *string `json:"key,omitempty" tf:"key"`

	Value *string `json:"value,omitempty" tf:"value"`
}

type UserPausedDetailsObservation struct {
	JobExpiresAt string `json:"jobExpiresAt" tf:"job_expires_at"`

	JobImminentExpirationHealthEventARN string `json:"jobImminentExpirationHealthEventARN" tf:"job_imminent_expiration_health_event_arn"`

	JobPausedAt string `json:"jobPausedAt" tf:"job_paused_at"`
}

type UserPausedDetailsParameters struct {
}

// Macie2ClassificationJobSpec defines the desired state of Macie2ClassificationJob
type Macie2ClassificationJobSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Macie2ClassificationJobParameters `json:"forProvider"`
}

// Macie2ClassificationJobStatus defines the observed state of Macie2ClassificationJob.
type Macie2ClassificationJobStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Macie2ClassificationJobObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Macie2ClassificationJob is the Schema for the Macie2ClassificationJobs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Macie2ClassificationJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Macie2ClassificationJobSpec   `json:"spec"`
	Status            Macie2ClassificationJobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Macie2ClassificationJobList contains a list of Macie2ClassificationJobs
type Macie2ClassificationJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Macie2ClassificationJob `json:"items"`
}

// Repository type metadata.
var (
	Macie2ClassificationJobKind             = "Macie2ClassificationJob"
	Macie2ClassificationJobGroupKind        = schema.GroupKind{Group: Group, Kind: Macie2ClassificationJobKind}.String()
	Macie2ClassificationJobKindAPIVersion   = Macie2ClassificationJobKind + "." + GroupVersion.String()
	Macie2ClassificationJobGroupVersionKind = GroupVersion.WithKind(Macie2ClassificationJobKind)
)

func init() {
	SchemeBuilder.Register(&Macie2ClassificationJob{}, &Macie2ClassificationJobList{})
}
