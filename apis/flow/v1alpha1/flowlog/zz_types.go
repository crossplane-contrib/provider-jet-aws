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
// +groupName=flow.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/flow/v1alpha1"
)

type FlowLogObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type FlowLogParameters struct {
	EniId *string `json:"eniId,omitempty" tf:"eni_id"`

	IamRoleArn *string `json:"iamRoleArn,omitempty" tf:"iam_role_arn"`

	LogDestination *string `json:"logDestination,omitempty" tf:"log_destination"`

	LogDestinationType *string `json:"logDestinationType,omitempty" tf:"log_destination_type"`

	LogFormat *string `json:"logFormat,omitempty" tf:"log_format"`

	LogGroupName *string `json:"logGroupName,omitempty" tf:"log_group_name"`

	MaxAggregationInterval *int64 `json:"maxAggregationInterval,omitempty" tf:"max_aggregation_interval"`

	SubnetId *string `json:"subnetId,omitempty" tf:"subnet_id"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	TrafficType string `json:"trafficType" tf:"traffic_type"`

	VpcId *string `json:"vpcId,omitempty" tf:"vpc_id"`
}

// FlowLogSpec defines the desired state of FlowLog
type FlowLogSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       FlowLogParameters `json:"forProvider"`
}

// FlowLogStatus defines the observed state of FlowLog.
type FlowLogStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          FlowLogObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FlowLog is the Schema for the FlowLogs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type FlowLog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlowLogSpec   `json:"spec"`
	Status            FlowLogStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FlowLogList contains a list of FlowLogs
type FlowLogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlowLog `json:"items"`
}

// Repository type metadata.
var (
	FlowLogKind             = "FlowLog"
	FlowLogGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: FlowLogKind}.String()
	FlowLogKindAPIVersion   = FlowLogKind + "." + v1alpha1.GroupVersion.String()
	FlowLogGroupVersionKind = v1alpha1.GroupVersion.WithKind(FlowLogKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&FlowLog{}, &FlowLogList{})
}