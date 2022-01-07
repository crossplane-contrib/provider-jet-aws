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

type InstanceGroupEBSConfigObservation struct {
}

type InstanceGroupEBSConfigParameters struct {

	// +kubebuilder:validation:Optional
	Iops *int64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// +kubebuilder:validation:Required
	Size *int64 `json:"size" tf:"size,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	VolumesPerInstance *int64 `json:"volumesPerInstance,omitempty" tf:"volumes_per_instance,omitempty"`
}

type InstanceGroupObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	RunningInstanceCount *int64 `json:"runningInstanceCount,omitempty" tf:"running_instance_count,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type InstanceGroupParameters struct {

	// +kubebuilder:validation:Optional
	AutoscalingPolicy *string `json:"autoscalingPolicy,omitempty" tf:"autoscaling_policy,omitempty"`

	// +kubebuilder:validation:Optional
	BidPrice *string `json:"bidPrice,omitempty" tf:"bid_price,omitempty"`

	// +kubebuilder:validation:Required
	ClusterID *string `json:"clusterId" tf:"cluster_id,omitempty"`

	// +kubebuilder:validation:Optional
	ConfigurationsJSON *string `json:"configurationsJson,omitempty" tf:"configurations_json,omitempty"`

	// +kubebuilder:validation:Optional
	EBSConfig []InstanceGroupEBSConfigParameters `json:"ebsConfig,omitempty" tf:"ebs_config,omitempty"`

	// +kubebuilder:validation:Optional
	EBSOptimized *bool `json:"ebsOptimized,omitempty" tf:"ebs_optimized,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceCount *int64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// +kubebuilder:validation:Required
	InstanceType *string `json:"instanceType" tf:"instance_type,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// InstanceGroupSpec defines the desired state of InstanceGroup
type InstanceGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceGroupParameters `json:"forProvider"`
}

// InstanceGroupStatus defines the observed state of InstanceGroup.
type InstanceGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceGroup is the Schema for the InstanceGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type InstanceGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceGroupSpec   `json:"spec"`
	Status            InstanceGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceGroupList contains a list of InstanceGroups
type InstanceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceGroup `json:"items"`
}

// Repository type metadata.
var (
	InstanceGroup_Kind             = "InstanceGroup"
	InstanceGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstanceGroup_Kind}.String()
	InstanceGroup_KindAPIVersion   = InstanceGroup_Kind + "." + CRDGroupVersion.String()
	InstanceGroup_GroupVersionKind = CRDGroupVersion.WithKind(InstanceGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&InstanceGroup{}, &InstanceGroupList{})
}
