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

type AppSourceObservation struct {
}

type AppSourceParameters struct {

	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Revision *string `json:"revision,omitempty" tf:"revision,omitempty"`

	// +kubebuilder:validation:Optional
	SSHKeySecretRef *v1.SecretKeySelector `json:"sshKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type ApplicationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ApplicationParameters struct {

	// +kubebuilder:validation:Optional
	AppSource []AppSourceParameters `json:"appSource,omitempty" tf:"app_source,omitempty"`

	// +kubebuilder:validation:Optional
	AutoBundleOnDeploy *string `json:"autoBundleOnDeploy,omitempty" tf:"auto_bundle_on_deploy,omitempty"`

	// +kubebuilder:validation:Optional
	AwsFlowRubySettings *string `json:"awsFlowRubySettings,omitempty" tf:"aws_flow_ruby_settings,omitempty"`

	// +kubebuilder:validation:Optional
	DataSourceArn *string `json:"dataSourceArn,omitempty" tf:"data_source_arn,omitempty"`

	// +kubebuilder:validation:Optional
	DataSourceDatabaseName *string `json:"dataSourceDatabaseName,omitempty" tf:"data_source_database_name,omitempty"`

	// +kubebuilder:validation:Optional
	DataSourceType *string `json:"dataSourceType,omitempty" tf:"data_source_type,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DocumentRoot *string `json:"documentRoot,omitempty" tf:"document_root,omitempty"`

	// +kubebuilder:validation:Optional
	Domains []*string `json:"domains,omitempty" tf:"domains,omitempty"`

	// +kubebuilder:validation:Optional
	EnableSSL *bool `json:"enableSsl,omitempty" tf:"enable_ssl,omitempty"`

	// +kubebuilder:validation:Optional
	Environment []EnvironmentParameters `json:"environment,omitempty" tf:"environment,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	RailsEnv *string `json:"railsEnv,omitempty" tf:"rails_env,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SSLConfiguration []SSLConfigurationParameters `json:"sslConfiguration,omitempty" tf:"ssl_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	ShortName *string `json:"shortName,omitempty" tf:"short_name,omitempty"`

	// +kubebuilder:validation:Required
	StackID *string `json:"stackId" tf:"stack_id,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type EnvironmentObservation struct {
}

type EnvironmentParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Secure *bool `json:"secure,omitempty" tf:"secure,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type SSLConfigurationObservation struct {
}

type SSLConfigurationParameters struct {

	// +kubebuilder:validation:Required
	Certificate *string `json:"certificate" tf:"certificate,omitempty"`

	// +kubebuilder:validation:Optional
	Chain *string `json:"chain,omitempty" tf:"chain,omitempty"`

	// +kubebuilder:validation:Required
	PrivateKeySecretRef v1.SecretKeySelector `json:"privateKeySecretRef" tf:"-"`
}

// ApplicationSpec defines the desired state of Application
type ApplicationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationParameters `json:"forProvider"`
}

// ApplicationStatus defines the observed state of Application.
type ApplicationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Application is the Schema for the Applications API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationSpec   `json:"spec"`
	Status            ApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationList contains a list of Applications
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

// Repository type metadata.
var (
	Application_Kind             = "Application"
	Application_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Application_Kind}.String()
	Application_KindAPIVersion   = Application_Kind + "." + CRDGroupVersion.String()
	Application_GroupVersionKind = CRDGroupVersion.WithKind(Application_Kind)
)

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{})
}
