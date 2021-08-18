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

type DbSecurityGroupObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type DbSecurityGroupParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	Ingress []IngressParameters `json:"ingress" tf:"ingress"`

	Name string `json:"name" tf:"name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type IngressObservation struct {
}

type IngressParameters struct {
	Cidr *string `json:"cidr,omitempty" tf:"cidr"`

	SecurityGroupId *string `json:"securityGroupId,omitempty" tf:"security_group_id"`

	SecurityGroupName *string `json:"securityGroupName,omitempty" tf:"security_group_name"`

	SecurityGroupOwnerId *string `json:"securityGroupOwnerId,omitempty" tf:"security_group_owner_id"`
}

// DbSecurityGroupSpec defines the desired state of DbSecurityGroup
type DbSecurityGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DbSecurityGroupParameters `json:"forProvider"`
}

// DbSecurityGroupStatus defines the observed state of DbSecurityGroup.
type DbSecurityGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DbSecurityGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DbSecurityGroup is the Schema for the DbSecurityGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DbSecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DbSecurityGroupSpec   `json:"spec"`
	Status            DbSecurityGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DbSecurityGroupList contains a list of DbSecurityGroups
type DbSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DbSecurityGroup `json:"items"`
}

// Repository type metadata.
var (
	DbSecurityGroupKind             = "DbSecurityGroup"
	DbSecurityGroupGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: DbSecurityGroupKind}.String()
	DbSecurityGroupKindAPIVersion   = DbSecurityGroupKind + "." + v1alpha1.GroupVersion.String()
	DbSecurityGroupGroupVersionKind = v1alpha1.GroupVersion.WithKind(DbSecurityGroupKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&DbSecurityGroup{}, &DbSecurityGroupList{})
}