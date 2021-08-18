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

// +kubebuilder:object:generate=true
// +groupName=media.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/media/v1alpha1"
)

type MediaConvertQueueObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type MediaConvertQueueParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	Name string `json:"name" tf:"name"`

	PricingPlan *string `json:"pricingPlan,omitempty" tf:"pricing_plan"`

	ReservationPlanSettings []ReservationPlanSettingsParameters `json:"reservationPlanSettings,omitempty" tf:"reservation_plan_settings"`

	Status *string `json:"status,omitempty" tf:"status"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type ReservationPlanSettingsObservation struct {
}

type ReservationPlanSettingsParameters struct {
	Commitment string `json:"commitment" tf:"commitment"`

	RenewalType string `json:"renewalType" tf:"renewal_type"`

	ReservedSlots int64 `json:"reservedSlots" tf:"reserved_slots"`
}

// MediaConvertQueueSpec defines the desired state of MediaConvertQueue
type MediaConvertQueueSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       MediaConvertQueueParameters `json:"forProvider"`
}

// MediaConvertQueueStatus defines the observed state of MediaConvertQueue.
type MediaConvertQueueStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          MediaConvertQueueObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MediaConvertQueue is the Schema for the MediaConvertQueues API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type MediaConvertQueue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MediaConvertQueueSpec   `json:"spec"`
	Status            MediaConvertQueueStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MediaConvertQueueList contains a list of MediaConvertQueues
type MediaConvertQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MediaConvertQueue `json:"items"`
}

// Repository type metadata.
var (
	MediaConvertQueueKind             = "MediaConvertQueue"
	MediaConvertQueueGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: MediaConvertQueueKind}.String()
	MediaConvertQueueKindAPIVersion   = MediaConvertQueueKind + "." + v1alpha1.GroupVersion.String()
	MediaConvertQueueGroupVersionKind = v1alpha1.GroupVersion.WithKind(MediaConvertQueueKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&MediaConvertQueue{}, &MediaConvertQueueList{})
}