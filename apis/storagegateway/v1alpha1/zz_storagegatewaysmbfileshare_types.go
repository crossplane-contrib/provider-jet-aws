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

type StoragegatewaySmbFileShareCacheAttributesObservation struct {
}

type StoragegatewaySmbFileShareCacheAttributesParameters struct {
	CacheStaleTimeoutInSeconds *int64 `json:"cacheStaleTimeoutInSeconds,omitempty" tf:"cache_stale_timeout_in_seconds"`
}

type StoragegatewaySmbFileShareObservation struct {
	ARN string `json:"arn" tf:"arn"`

	FileshareID string `json:"fileshareID" tf:"fileshare_id"`

	Path string `json:"path" tf:"path"`
}

type StoragegatewaySmbFileShareParameters struct {
	AccessBasedEnumeration *bool `json:"accessBasedEnumeration,omitempty" tf:"access_based_enumeration"`

	AdminUserList []string `json:"adminUserList,omitempty" tf:"admin_user_list"`

	AuditDestinationARN *string `json:"auditDestinationARN,omitempty" tf:"audit_destination_arn"`

	Authentication *string `json:"authentication,omitempty" tf:"authentication"`

	BucketRegion *string `json:"bucketRegion,omitempty" tf:"bucket_region"`

	CacheAttributes []StoragegatewaySmbFileShareCacheAttributesParameters `json:"cacheAttributes,omitempty" tf:"cache_attributes"`

	CaseSensitivity *string `json:"caseSensitivity,omitempty" tf:"case_sensitivity"`

	DefaultStorageClass *string `json:"defaultStorageClass,omitempty" tf:"default_storage_class"`

	FileShareName *string `json:"fileShareName,omitempty" tf:"file_share_name"`

	GatewayARN string `json:"gatewayARN" tf:"gateway_arn"`

	GuessMimeTypeEnabled *bool `json:"guessMimeTypeEnabled,omitempty" tf:"guess_mime_type_enabled"`

	InvalidUserList []string `json:"invalidUserList,omitempty" tf:"invalid_user_list"`

	KmsEncrypted *bool `json:"kmsEncrypted,omitempty" tf:"kms_encrypted"`

	KmsKeyARN *string `json:"kmsKeyARN,omitempty" tf:"kms_key_arn"`

	LocationARN string `json:"locationARN" tf:"location_arn"`

	NotificationPolicy *string `json:"notificationPolicy,omitempty" tf:"notification_policy"`

	ObjectACL *string `json:"objectACL,omitempty" tf:"object_acl"`

	OplocksEnabled *bool `json:"oplocksEnabled,omitempty" tf:"oplocks_enabled"`

	ReadOnly *bool `json:"readOnly,omitempty" tf:"read_only"`

	RequesterPays *bool `json:"requesterPays,omitempty" tf:"requester_pays"`

	RoleARN string `json:"roleARN" tf:"role_arn"`

	SmbACLEnabled *bool `json:"smbACLEnabled,omitempty" tf:"smb_acl_enabled"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	VPCEndpointDNSName *string `json:"vpcEndpointDNSName,omitempty" tf:"vpc_endpoint_dns_name"`

	ValidUserList []string `json:"validUserList,omitempty" tf:"valid_user_list"`
}

// StoragegatewaySmbFileShareSpec defines the desired state of StoragegatewaySmbFileShare
type StoragegatewaySmbFileShareSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       StoragegatewaySmbFileShareParameters `json:"forProvider"`
}

// StoragegatewaySmbFileShareStatus defines the observed state of StoragegatewaySmbFileShare.
type StoragegatewaySmbFileShareStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          StoragegatewaySmbFileShareObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StoragegatewaySmbFileShare is the Schema for the StoragegatewaySmbFileShares API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type StoragegatewaySmbFileShare struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StoragegatewaySmbFileShareSpec   `json:"spec"`
	Status            StoragegatewaySmbFileShareStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StoragegatewaySmbFileShareList contains a list of StoragegatewaySmbFileShares
type StoragegatewaySmbFileShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StoragegatewaySmbFileShare `json:"items"`
}

// Repository type metadata.
var (
	StoragegatewaySmbFileShareKind             = "StoragegatewaySmbFileShare"
	StoragegatewaySmbFileShareGroupKind        = schema.GroupKind{Group: Group, Kind: StoragegatewaySmbFileShareKind}.String()
	StoragegatewaySmbFileShareKindAPIVersion   = StoragegatewaySmbFileShareKind + "." + GroupVersion.String()
	StoragegatewaySmbFileShareGroupVersionKind = GroupVersion.WithKind(StoragegatewaySmbFileShareKind)
)

func init() {
	SchemeBuilder.Register(&StoragegatewaySmbFileShare{}, &StoragegatewaySmbFileShareList{})
}
