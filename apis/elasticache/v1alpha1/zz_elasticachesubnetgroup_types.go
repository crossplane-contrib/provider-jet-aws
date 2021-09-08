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

type ElasticacheSubnetGroupObservation struct {
	ARN string `json:"arn" tf:"arn"`
}

type ElasticacheSubnetGroupParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	Name string `json:"name" tf:"name"`

	SubnetIds []string `json:"subnetIds" tf:"subnet_ids"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// ElasticacheSubnetGroupSpec defines the desired state of ElasticacheSubnetGroup
type ElasticacheSubnetGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ElasticacheSubnetGroupParameters `json:"forProvider"`
}

// ElasticacheSubnetGroupStatus defines the observed state of ElasticacheSubnetGroup.
type ElasticacheSubnetGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ElasticacheSubnetGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticacheSubnetGroup is the Schema for the ElasticacheSubnetGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ElasticacheSubnetGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticacheSubnetGroupSpec   `json:"spec"`
	Status            ElasticacheSubnetGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticacheSubnetGroupList contains a list of ElasticacheSubnetGroups
type ElasticacheSubnetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticacheSubnetGroup `json:"items"`
}

// Repository type metadata.
var (
	ElasticacheSubnetGroupKind             = "ElasticacheSubnetGroup"
	ElasticacheSubnetGroupGroupKind        = schema.GroupKind{Group: Group, Kind: ElasticacheSubnetGroupKind}.String()
	ElasticacheSubnetGroupKindAPIVersion   = ElasticacheSubnetGroupKind + "." + GroupVersion.String()
	ElasticacheSubnetGroupGroupVersionKind = GroupVersion.WithKind(ElasticacheSubnetGroupKind)
)

func init() {
	SchemeBuilder.Register(&ElasticacheSubnetGroup{}, &ElasticacheSubnetGroupList{})
}
