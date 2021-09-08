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

type DbEventSubscriptionObservation struct {
	ARN string `json:"arn" tf:"arn"`

	CustomerAwsID string `json:"customerAwsID" tf:"customer_aws_id"`
}

type DbEventSubscriptionParameters struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	EventCategories []string `json:"eventCategories,omitempty" tf:"event_categories"`

	Name *string `json:"name,omitempty" tf:"name"`

	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`

	SnsTopic string `json:"snsTopic" tf:"sns_topic"`

	SourceIds []string `json:"sourceIds,omitempty" tf:"source_ids"`

	SourceType *string `json:"sourceType,omitempty" tf:"source_type"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// DbEventSubscriptionSpec defines the desired state of DbEventSubscription
type DbEventSubscriptionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DbEventSubscriptionParameters `json:"forProvider"`
}

// DbEventSubscriptionStatus defines the observed state of DbEventSubscription.
type DbEventSubscriptionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DbEventSubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DbEventSubscription is the Schema for the DbEventSubscriptions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type DbEventSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DbEventSubscriptionSpec   `json:"spec"`
	Status            DbEventSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DbEventSubscriptionList contains a list of DbEventSubscriptions
type DbEventSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DbEventSubscription `json:"items"`
}

// Repository type metadata.
var (
	DbEventSubscriptionKind             = "DbEventSubscription"
	DbEventSubscriptionGroupKind        = schema.GroupKind{Group: Group, Kind: DbEventSubscriptionKind}.String()
	DbEventSubscriptionKindAPIVersion   = DbEventSubscriptionKind + "." + GroupVersion.String()
	DbEventSubscriptionGroupVersionKind = GroupVersion.WithKind(DbEventSubscriptionKind)
)

func init() {
	SchemeBuilder.Register(&DbEventSubscription{}, &DbEventSubscriptionList{})
}
