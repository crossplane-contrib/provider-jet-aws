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

type DbProxyTargetObservation struct {
	Endpoint string `json:"endpoint" tf:"endpoint"`

	Port int64 `json:"port" tf:"port"`

	RdsResourceID string `json:"rdsResourceID" tf:"rds_resource_id"`

	TargetARN string `json:"targetARN" tf:"target_arn"`

	TrackedClusterID string `json:"trackedClusterID" tf:"tracked_cluster_id"`

	Type string `json:"type" tf:"type"`
}

type DbProxyTargetParameters struct {
	DBClusterIdentifier *string `json:"dbClusterIdentifier,omitempty" tf:"db_cluster_identifier"`

	DBInstanceIdentifier *string `json:"dbInstanceIdentifier,omitempty" tf:"db_instance_identifier"`

	DBProxyName string `json:"dbProxyName" tf:"db_proxy_name"`

	TargetGroupName string `json:"targetGroupName" tf:"target_group_name"`
}

// DbProxyTargetSpec defines the desired state of DbProxyTarget
type DbProxyTargetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DbProxyTargetParameters `json:"forProvider"`
}

// DbProxyTargetStatus defines the observed state of DbProxyTarget.
type DbProxyTargetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DbProxyTargetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DbProxyTarget is the Schema for the DbProxyTargets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type DbProxyTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DbProxyTargetSpec   `json:"spec"`
	Status            DbProxyTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DbProxyTargetList contains a list of DbProxyTargets
type DbProxyTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DbProxyTarget `json:"items"`
}

// Repository type metadata.
var (
	DbProxyTargetKind             = "DbProxyTarget"
	DbProxyTargetGroupKind        = schema.GroupKind{Group: Group, Kind: DbProxyTargetKind}.String()
	DbProxyTargetKindAPIVersion   = DbProxyTargetKind + "." + GroupVersion.String()
	DbProxyTargetGroupVersionKind = GroupVersion.WithKind(DbProxyTargetKind)
)

func init() {
	SchemeBuilder.Register(&DbProxyTarget{}, &DbProxyTargetList{})
}
