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

type ConfigurationObservation struct {
}

type ConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	BytesScannedCutoffPerQuery *int64 `json:"bytesScannedCutoffPerQuery,omitempty" tf:"bytes_scanned_cutoff_per_query,omitempty"`

	// +kubebuilder:validation:Optional
	EnforceWorkgroupConfiguration *bool `json:"enforceWorkgroupConfiguration,omitempty" tf:"enforce_workgroup_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	PublishCloudwatchMetricsEnabled *bool `json:"publishCloudwatchMetricsEnabled,omitempty" tf:"publish_cloudwatch_metrics_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	RequesterPaysEnabled *bool `json:"requesterPaysEnabled,omitempty" tf:"requester_pays_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	ResultConfiguration []ResultConfigurationParameters `json:"resultConfiguration,omitempty" tf:"result_configuration,omitempty"`
}

type ResultConfigurationEncryptionConfigurationObservation struct {
}

type ResultConfigurationEncryptionConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	EncryptionOption *string `json:"encryptionOption,omitempty" tf:"encryption_option,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`
}

type ResultConfigurationObservation struct {
}

type ResultConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	EncryptionConfiguration []ResultConfigurationEncryptionConfigurationParameters `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	OutputLocation *string `json:"outputLocation,omitempty" tf:"output_location,omitempty"`
}

type WorkgroupObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type WorkgroupParameters struct {

	// +kubebuilder:validation:Optional
	Configuration []ConfigurationParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// WorkgroupSpec defines the desired state of Workgroup
type WorkgroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkgroupParameters `json:"forProvider"`
}

// WorkgroupStatus defines the observed state of Workgroup.
type WorkgroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkgroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Workgroup is the Schema for the Workgroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Workgroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkgroupSpec   `json:"spec"`
	Status            WorkgroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkgroupList contains a list of Workgroups
type WorkgroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workgroup `json:"items"`
}

// Repository type metadata.
var (
	Workgroup_Kind             = "Workgroup"
	Workgroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Workgroup_Kind}.String()
	Workgroup_KindAPIVersion   = Workgroup_Kind + "." + CRDGroupVersion.String()
	Workgroup_GroupVersionKind = CRDGroupVersion.WithKind(Workgroup_Kind)
)

func init() {
	SchemeBuilder.Register(&Workgroup{}, &WorkgroupList{})
}
