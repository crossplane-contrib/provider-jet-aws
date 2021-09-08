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

type IamServerCertificateObservation struct {
	ARN string `json:"arn" tf:"arn"`

	Expiration string `json:"expiration" tf:"expiration"`

	UploadDate string `json:"uploadDate" tf:"upload_date"`
}

type IamServerCertificateParameters struct {
	CertificateBody string `json:"certificateBody" tf:"certificate_body"`

	CertificateChain *string `json:"certificateChain,omitempty" tf:"certificate_chain"`

	Name *string `json:"name,omitempty" tf:"name"`

	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`

	Path *string `json:"path,omitempty" tf:"path"`

	PrivateKey string `json:"privateKey" tf:"private_key"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// IamServerCertificateSpec defines the desired state of IamServerCertificate
type IamServerCertificateSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       IamServerCertificateParameters `json:"forProvider"`
}

// IamServerCertificateStatus defines the observed state of IamServerCertificate.
type IamServerCertificateStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          IamServerCertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IamServerCertificate is the Schema for the IamServerCertificates API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type IamServerCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamServerCertificateSpec   `json:"spec"`
	Status            IamServerCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamServerCertificateList contains a list of IamServerCertificates
type IamServerCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamServerCertificate `json:"items"`
}

// Repository type metadata.
var (
	IamServerCertificateKind             = "IamServerCertificate"
	IamServerCertificateGroupKind        = schema.GroupKind{Group: Group, Kind: IamServerCertificateKind}.String()
	IamServerCertificateKindAPIVersion   = IamServerCertificateKind + "." + GroupVersion.String()
	IamServerCertificateGroupVersionKind = GroupVersion.WithKind(IamServerCertificateKind)
)

func init() {
	SchemeBuilder.Register(&IamServerCertificate{}, &IamServerCertificateList{})
}
