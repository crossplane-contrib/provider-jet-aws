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

type GroupObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type GroupParameters struct {

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/kms/v1alpha2.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RetentionInDays *int64 `json:"retentionInDays,omitempty" tf:"retention_in_days,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// GroupSpec defines the desired state of Group
type GroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupParameters `json:"forProvider"`
}

// GroupStatus defines the observed state of Group.
type GroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Group is the Schema for the Groups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Group struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GroupSpec   `json:"spec"`
	Status            GroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupList contains a list of Groups
type GroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Group `json:"items"`
}

// Repository type metadata.
var (
	Group_Kind             = "Group"
	Group_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Group_Kind}.String()
	Group_KindAPIVersion   = Group_Kind + "." + CRDGroupVersion.String()
	Group_GroupVersionKind = CRDGroupVersion.WithKind(Group_Kind)
)

func init() {
	SchemeBuilder.Register(&Group{}, &GroupList{})
}
