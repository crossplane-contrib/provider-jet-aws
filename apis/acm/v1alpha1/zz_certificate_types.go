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

type CertificateObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	DomainValidationOptions []DomainValidationOptionsObservation `json:"domainValidationOptions,omitempty" tf:"domain_validation_options,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	ValidationEmails []*string `json:"validationEmails,omitempty" tf:"validation_emails,omitempty"`
}

type CertificateParameters struct {

	// +kubebuilder:validation:Optional
	CertificateAuthorityArn *string `json:"certificateAuthorityArn,omitempty" tf:"certificate_authority_arn,omitempty"`

	// +kubebuilder:validation:Optional
	CertificateBody *string `json:"certificateBody,omitempty" tf:"certificate_body,omitempty"`

	// +kubebuilder:validation:Optional
	CertificateChain *string `json:"certificateChain,omitempty" tf:"certificate_chain,omitempty"`

	// +kubebuilder:validation:Optional
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// +kubebuilder:validation:Optional
	Options []OptionsParameters `json:"options,omitempty" tf:"options,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateKeySecretRef *v1.SecretKeySelector `json:"privateKeySecretRef,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SubjectAlternativeNames []*string `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	ValidationMethod *string `json:"validationMethod,omitempty" tf:"validation_method,omitempty"`
}

type DomainValidationOptionsObservation struct {
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	ResourceRecordName *string `json:"resourceRecordName,omitempty" tf:"resource_record_name,omitempty"`

	ResourceRecordType *string `json:"resourceRecordType,omitempty" tf:"resource_record_type,omitempty"`

	ResourceRecordValue *string `json:"resourceRecordValue,omitempty" tf:"resource_record_value,omitempty"`
}

type DomainValidationOptionsParameters struct {
}

type OptionsObservation struct {
}

type OptionsParameters struct {

	// +kubebuilder:validation:Optional
	CertificateTransparencyLoggingPreference *string `json:"certificateTransparencyLoggingPreference,omitempty" tf:"certificate_transparency_logging_preference,omitempty"`
}

// CertificateSpec defines the desired state of Certificate
type CertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CertificateParameters `json:"forProvider"`
}

// CertificateStatus defines the observed state of Certificate.
type CertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Certificate is the Schema for the Certificates API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Certificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CertificateSpec   `json:"spec"`
	Status            CertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateList contains a list of Certificates
type CertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Certificate `json:"items"`
}

// Repository type metadata.
var (
	Certificate_Kind             = "Certificate"
	Certificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Certificate_Kind}.String()
	Certificate_KindAPIVersion   = Certificate_Kind + "." + CRDGroupVersion.String()
	Certificate_GroupVersionKind = CRDGroupVersion.WithKind(Certificate_Kind)
)

func init() {
	SchemeBuilder.Register(&Certificate{}, &CertificateList{})
}
