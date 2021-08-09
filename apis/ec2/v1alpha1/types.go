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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// VPCParameters are the configurable fields of a VPC.
type VPCParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// VPCObservation are the observable fields of a VPC.
type VPCObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A VPCSpec defines the desired state of a VPC.
type VPCSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VPCParameters `json:"forProvider"`
}

// A VPCStatus represents the observed state of a VPC.
type VPCStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          VPCObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A VPC is an example API type.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.bindingPhase"
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".status.atProvider.state"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// Please replace `PROVIDER-NAME` with your actual provider name, like `aws`, `azure`, `gcp`, `alibaba`
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,PROVIDER-NAME}
type VPC struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VPCSpec   `json:"spec"`
	Status VPCStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCList contains a list of VPC
type VPCList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPC `json:"items"`
}
