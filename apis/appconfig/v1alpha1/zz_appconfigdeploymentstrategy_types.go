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

type AppconfigDeploymentStrategyObservation struct {
	ARN string `json:"arn" tf:"arn"`
}

type AppconfigDeploymentStrategyParameters struct {
	DeploymentDurationInMinutes int64 `json:"deploymentDurationInMinutes" tf:"deployment_duration_in_minutes"`

	Description *string `json:"description,omitempty" tf:"description"`

	FinalBakeTimeInMinutes *int64 `json:"finalBakeTimeInMinutes,omitempty" tf:"final_bake_time_in_minutes"`

	GrowthFactor float64 `json:"growthFactor" tf:"growth_factor"`

	GrowthType *string `json:"growthType,omitempty" tf:"growth_type"`

	Name string `json:"name" tf:"name"`

	ReplicateTo string `json:"replicateTo" tf:"replicate_to"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// AppconfigDeploymentStrategySpec defines the desired state of AppconfigDeploymentStrategy
type AppconfigDeploymentStrategySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AppconfigDeploymentStrategyParameters `json:"forProvider"`
}

// AppconfigDeploymentStrategyStatus defines the observed state of AppconfigDeploymentStrategy.
type AppconfigDeploymentStrategyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AppconfigDeploymentStrategyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppconfigDeploymentStrategy is the Schema for the AppconfigDeploymentStrategys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type AppconfigDeploymentStrategy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppconfigDeploymentStrategySpec   `json:"spec"`
	Status            AppconfigDeploymentStrategyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppconfigDeploymentStrategyList contains a list of AppconfigDeploymentStrategys
type AppconfigDeploymentStrategyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppconfigDeploymentStrategy `json:"items"`
}

// Repository type metadata.
var (
	AppconfigDeploymentStrategyKind             = "AppconfigDeploymentStrategy"
	AppconfigDeploymentStrategyGroupKind        = schema.GroupKind{Group: Group, Kind: AppconfigDeploymentStrategyKind}.String()
	AppconfigDeploymentStrategyKindAPIVersion   = AppconfigDeploymentStrategyKind + "." + GroupVersion.String()
	AppconfigDeploymentStrategyGroupVersionKind = GroupVersion.WithKind(AppconfigDeploymentStrategyKind)
)

func init() {
	SchemeBuilder.Register(&AppconfigDeploymentStrategy{}, &AppconfigDeploymentStrategyList{})
}
