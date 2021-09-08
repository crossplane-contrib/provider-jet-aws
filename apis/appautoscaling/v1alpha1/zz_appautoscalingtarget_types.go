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

type AppautoscalingTargetObservation struct {
}

type AppautoscalingTargetParameters struct {
	MaxCapacity int64 `json:"maxCapacity" tf:"max_capacity"`

	MinCapacity int64 `json:"minCapacity" tf:"min_capacity"`

	ResourceID string `json:"resourceID" tf:"resource_id"`

	RoleARN *string `json:"roleARN,omitempty" tf:"role_arn"`

	ScalableDimension string `json:"scalableDimension" tf:"scalable_dimension"`

	ServiceNamespace string `json:"serviceNamespace" tf:"service_namespace"`
}

// AppautoscalingTargetSpec defines the desired state of AppautoscalingTarget
type AppautoscalingTargetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AppautoscalingTargetParameters `json:"forProvider"`
}

// AppautoscalingTargetStatus defines the observed state of AppautoscalingTarget.
type AppautoscalingTargetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AppautoscalingTargetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppautoscalingTarget is the Schema for the AppautoscalingTargets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type AppautoscalingTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppautoscalingTargetSpec   `json:"spec"`
	Status            AppautoscalingTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppautoscalingTargetList contains a list of AppautoscalingTargets
type AppautoscalingTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppautoscalingTarget `json:"items"`
}

// Repository type metadata.
var (
	AppautoscalingTargetKind             = "AppautoscalingTarget"
	AppautoscalingTargetGroupKind        = schema.GroupKind{Group: Group, Kind: AppautoscalingTargetKind}.String()
	AppautoscalingTargetKindAPIVersion   = AppautoscalingTargetKind + "." + GroupVersion.String()
	AppautoscalingTargetGroupVersionKind = GroupVersion.WithKind(AppautoscalingTargetKind)
)

func init() {
	SchemeBuilder.Register(&AppautoscalingTarget{}, &AppautoscalingTargetList{})
}
