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
// +groupName=db.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/db/v1alpha1"
)

type DbSubnetGroupObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type DbSubnetGroupParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	Name *string `json:"name,omitempty" tf:"name"`

	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`

	SubnetIds []string `json:"subnetIds" tf:"subnet_ids"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// DbSubnetGroupSpec defines the desired state of DbSubnetGroup
type DbSubnetGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DbSubnetGroupParameters `json:"forProvider"`
}

// DbSubnetGroupStatus defines the observed state of DbSubnetGroup.
type DbSubnetGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DbSubnetGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DbSubnetGroup is the Schema for the DbSubnetGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DbSubnetGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DbSubnetGroupSpec   `json:"spec"`
	Status            DbSubnetGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DbSubnetGroupList contains a list of DbSubnetGroups
type DbSubnetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DbSubnetGroup `json:"items"`
}

// Repository type metadata.
var (
	DbSubnetGroupKind             = "DbSubnetGroup"
	DbSubnetGroupGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: DbSubnetGroupKind}.String()
	DbSubnetGroupKindAPIVersion   = DbSubnetGroupKind + "." + v1alpha1.GroupVersion.String()
	DbSubnetGroupGroupVersionKind = v1alpha1.GroupVersion.WithKind(DbSubnetGroupKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&DbSubnetGroup{}, &DbSubnetGroupList{})
}