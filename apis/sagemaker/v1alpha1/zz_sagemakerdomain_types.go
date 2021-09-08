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

type CustomImageObservation struct {
}

type CustomImageParameters struct {
	AppImageConfigName string `json:"appImageConfigName" tf:"app_image_config_name"`

	ImageName string `json:"imageName" tf:"image_name"`

	ImageVersionNumber *int64 `json:"imageVersionNumber,omitempty" tf:"image_version_number"`
}

type DefaultResourceSpecObservation struct {
}

type DefaultResourceSpecParameters struct {
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type"`

	SagemakerImageARN *string `json:"sagemakerImageARN,omitempty" tf:"sagemaker_image_arn"`
}

type DefaultUserSettingsObservation struct {
}

type DefaultUserSettingsParameters struct {
	ExecutionRole string `json:"executionRole" tf:"execution_role"`

	JupyterServerAppSettings []JupyterServerAppSettingsParameters `json:"jupyterServerAppSettings,omitempty" tf:"jupyter_server_app_settings"`

	KernelGatewayAppSettings []KernelGatewayAppSettingsParameters `json:"kernelGatewayAppSettings,omitempty" tf:"kernel_gateway_app_settings"`

	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups"`

	SharingSettings []SharingSettingsParameters `json:"sharingSettings,omitempty" tf:"sharing_settings"`

	TensorBoardAppSettings []TensorBoardAppSettingsParameters `json:"tensorBoardAppSettings,omitempty" tf:"tensor_board_app_settings"`
}

type JupyterServerAppSettingsObservation struct {
}

type JupyterServerAppSettingsParameters struct {
	DefaultResourceSpec []DefaultResourceSpecParameters `json:"defaultResourceSpec,omitempty" tf:"default_resource_spec"`
}

type KernelGatewayAppSettingsDefaultResourceSpecObservation struct {
}

type KernelGatewayAppSettingsDefaultResourceSpecParameters struct {
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type"`

	SagemakerImageARN *string `json:"sagemakerImageARN,omitempty" tf:"sagemaker_image_arn"`
}

type KernelGatewayAppSettingsObservation struct {
}

type KernelGatewayAppSettingsParameters struct {
	CustomImage []CustomImageParameters `json:"customImage,omitempty" tf:"custom_image"`

	DefaultResourceSpec []KernelGatewayAppSettingsDefaultResourceSpecParameters `json:"defaultResourceSpec,omitempty" tf:"default_resource_spec"`
}

type RetentionPolicyObservation struct {
}

type RetentionPolicyParameters struct {
	HomeEfsFileSystem *string `json:"homeEfsFileSystem,omitempty" tf:"home_efs_file_system"`
}

type SagemakerDomainObservation struct {
	ARN string `json:"arn" tf:"arn"`

	HomeEfsFileSystemID string `json:"homeEfsFileSystemID" tf:"home_efs_file_system_id"`

	SingleSignOnManagedApplicationInstanceID string `json:"singleSignOnManagedApplicationInstanceID" tf:"single_sign_on_managed_application_instance_id"`

	URL string `json:"url" tf:"url"`
}

type SagemakerDomainParameters struct {
	AppNetworkAccessType *string `json:"appNetworkAccessType,omitempty" tf:"app_network_access_type"`

	AuthMode string `json:"authMode" tf:"auth_mode"`

	DefaultUserSettings []DefaultUserSettingsParameters `json:"defaultUserSettings" tf:"default_user_settings"`

	DomainName string `json:"domainName" tf:"domain_name"`

	KmsKeyID *string `json:"kmsKeyID,omitempty" tf:"kms_key_id"`

	RetentionPolicy []RetentionPolicyParameters `json:"retentionPolicy,omitempty" tf:"retention_policy"`

	SubnetIds []string `json:"subnetIds" tf:"subnet_ids"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	VPCID string `json:"vpcID" tf:"vpc_id"`
}

type SharingSettingsObservation struct {
}

type SharingSettingsParameters struct {
	NotebookOutputOption *string `json:"notebookOutputOption,omitempty" tf:"notebook_output_option"`

	S3KmsKeyID *string `json:"s3KmsKeyID,omitempty" tf:"s3_kms_key_id"`

	S3OutputPath *string `json:"s3OutputPath,omitempty" tf:"s3_output_path"`
}

type TensorBoardAppSettingsDefaultResourceSpecObservation struct {
}

type TensorBoardAppSettingsDefaultResourceSpecParameters struct {
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type"`

	SagemakerImageARN *string `json:"sagemakerImageARN,omitempty" tf:"sagemaker_image_arn"`
}

type TensorBoardAppSettingsObservation struct {
}

type TensorBoardAppSettingsParameters struct {
	DefaultResourceSpec []TensorBoardAppSettingsDefaultResourceSpecParameters `json:"defaultResourceSpec,omitempty" tf:"default_resource_spec"`
}

// SagemakerDomainSpec defines the desired state of SagemakerDomain
type SagemakerDomainSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SagemakerDomainParameters `json:"forProvider"`
}

// SagemakerDomainStatus defines the observed state of SagemakerDomain.
type SagemakerDomainStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SagemakerDomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SagemakerDomain is the Schema for the SagemakerDomains API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type SagemakerDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SagemakerDomainSpec   `json:"spec"`
	Status            SagemakerDomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SagemakerDomainList contains a list of SagemakerDomains
type SagemakerDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SagemakerDomain `json:"items"`
}

// Repository type metadata.
var (
	SagemakerDomainKind             = "SagemakerDomain"
	SagemakerDomainGroupKind        = schema.GroupKind{Group: Group, Kind: SagemakerDomainKind}.String()
	SagemakerDomainKindAPIVersion   = SagemakerDomainKind + "." + GroupVersion.String()
	SagemakerDomainGroupVersionKind = GroupVersion.WithKind(SagemakerDomainKind)
)

func init() {
	SchemeBuilder.Register(&SagemakerDomain{}, &SagemakerDomainList{})
}
