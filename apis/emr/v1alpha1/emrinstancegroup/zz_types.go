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
// +groupName=emr.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/emr/v1alpha1"
)

type EbsConfigObservation struct {
}

type EbsConfigParameters struct {
	Iops *int64 `json:"iops,omitempty" tf:"iops"`

	Size int64 `json:"size" tf:"size"`

	Type string `json:"type" tf:"type"`

	VolumesPerInstance *int64 `json:"volumesPerInstance,omitempty" tf:"volumes_per_instance"`
}

type EmrInstanceGroupObservation struct {
	RunningInstanceCount int64 `json:"runningInstanceCount" tf:"running_instance_count"`

	Status string `json:"status" tf:"status"`
}

type EmrInstanceGroupParameters struct {
	AutoscalingPolicy *string `json:"autoscalingPolicy,omitempty" tf:"autoscaling_policy"`

	BidPrice *string `json:"bidPrice,omitempty" tf:"bid_price"`

	ClusterId string `json:"clusterId" tf:"cluster_id"`

	ConfigurationsJson *string `json:"configurationsJson,omitempty" tf:"configurations_json"`

	EbsConfig []EbsConfigParameters `json:"ebsConfig,omitempty" tf:"ebs_config"`

	EbsOptimized *bool `json:"ebsOptimized,omitempty" tf:"ebs_optimized"`

	InstanceCount *int64 `json:"instanceCount,omitempty" tf:"instance_count"`

	InstanceType string `json:"instanceType" tf:"instance_type"`

	Name *string `json:"name,omitempty" tf:"name"`
}

// EmrInstanceGroupSpec defines the desired state of EmrInstanceGroup
type EmrInstanceGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EmrInstanceGroupParameters `json:"forProvider"`
}

// EmrInstanceGroupStatus defines the observed state of EmrInstanceGroup.
type EmrInstanceGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EmrInstanceGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EmrInstanceGroup is the Schema for the EmrInstanceGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EmrInstanceGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EmrInstanceGroupSpec   `json:"spec"`
	Status            EmrInstanceGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EmrInstanceGroupList contains a list of EmrInstanceGroups
type EmrInstanceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmrInstanceGroup `json:"items"`
}

// Repository type metadata.
var (
	EmrInstanceGroupKind             = "EmrInstanceGroup"
	EmrInstanceGroupGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: EmrInstanceGroupKind}.String()
	EmrInstanceGroupKindAPIVersion   = EmrInstanceGroupKind + "." + v1alpha1.GroupVersion.String()
	EmrInstanceGroupGroupVersionKind = v1alpha1.GroupVersion.WithKind(EmrInstanceGroupKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&EmrInstanceGroup{}, &EmrInstanceGroupList{})
}