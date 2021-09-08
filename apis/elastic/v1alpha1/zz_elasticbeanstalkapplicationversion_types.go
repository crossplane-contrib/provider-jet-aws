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

type ElasticBeanstalkApplicationVersionObservation struct {
	ARN string `json:"arn" tf:"arn"`
}

type ElasticBeanstalkApplicationVersionParameters struct {
	Application string `json:"application" tf:"application"`

	Bucket string `json:"bucket" tf:"bucket"`

	Description *string `json:"description,omitempty" tf:"description"`

	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete"`

	Key string `json:"key" tf:"key"`

	Name string `json:"name" tf:"name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// ElasticBeanstalkApplicationVersionSpec defines the desired state of ElasticBeanstalkApplicationVersion
type ElasticBeanstalkApplicationVersionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ElasticBeanstalkApplicationVersionParameters `json:"forProvider"`
}

// ElasticBeanstalkApplicationVersionStatus defines the observed state of ElasticBeanstalkApplicationVersion.
type ElasticBeanstalkApplicationVersionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ElasticBeanstalkApplicationVersionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticBeanstalkApplicationVersion is the Schema for the ElasticBeanstalkApplicationVersions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ElasticBeanstalkApplicationVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticBeanstalkApplicationVersionSpec   `json:"spec"`
	Status            ElasticBeanstalkApplicationVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticBeanstalkApplicationVersionList contains a list of ElasticBeanstalkApplicationVersions
type ElasticBeanstalkApplicationVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticBeanstalkApplicationVersion `json:"items"`
}

// Repository type metadata.
var (
	ElasticBeanstalkApplicationVersionKind             = "ElasticBeanstalkApplicationVersion"
	ElasticBeanstalkApplicationVersionGroupKind        = schema.GroupKind{Group: Group, Kind: ElasticBeanstalkApplicationVersionKind}.String()
	ElasticBeanstalkApplicationVersionKindAPIVersion   = ElasticBeanstalkApplicationVersionKind + "." + GroupVersion.String()
	ElasticBeanstalkApplicationVersionGroupVersionKind = GroupVersion.WithKind(ElasticBeanstalkApplicationVersionKind)
)

func init() {
	SchemeBuilder.Register(&ElasticBeanstalkApplicationVersion{}, &ElasticBeanstalkApplicationVersionList{})
}
