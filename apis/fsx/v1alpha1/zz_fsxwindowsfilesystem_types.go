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

type AuditLogConfigurationObservation struct {
}

type AuditLogConfigurationParameters struct {
	AuditLogDestination *string `json:"auditLogDestination,omitempty" tf:"audit_log_destination"`

	FileAccessAuditLogLevel *string `json:"fileAccessAuditLogLevel,omitempty" tf:"file_access_audit_log_level"`

	FileShareAccessAuditLogLevel *string `json:"fileShareAccessAuditLogLevel,omitempty" tf:"file_share_access_audit_log_level"`
}

type FsxWindowsFileSystemObservation struct {
	ARN string `json:"arn" tf:"arn"`

	DNSName string `json:"dnsName" tf:"dns_name"`

	NetworkInterfaceIds []string `json:"networkInterfaceIds" tf:"network_interface_ids"`

	OwnerID string `json:"ownerID" tf:"owner_id"`

	PreferredFileServerIP string `json:"preferredFileServerIP" tf:"preferred_file_server_ip"`

	RemoteAdministrationEndpoint string `json:"remoteAdministrationEndpoint" tf:"remote_administration_endpoint"`

	VPCID string `json:"vpcID" tf:"vpc_id"`
}

type FsxWindowsFileSystemParameters struct {
	ActiveDirectoryID *string `json:"activeDirectoryID,omitempty" tf:"active_directory_id"`

	Aliases []string `json:"aliases,omitempty" tf:"aliases"`

	AuditLogConfiguration []AuditLogConfigurationParameters `json:"auditLogConfiguration,omitempty" tf:"audit_log_configuration"`

	AutomaticBackupRetentionDays *int64 `json:"automaticBackupRetentionDays,omitempty" tf:"automatic_backup_retention_days"`

	CopyTagsToBackups *bool `json:"copyTagsToBackups,omitempty" tf:"copy_tags_to_backups"`

	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime,omitempty" tf:"daily_automatic_backup_start_time"`

	DeploymentType *string `json:"deploymentType,omitempty" tf:"deployment_type"`

	KmsKeyID *string `json:"kmsKeyID,omitempty" tf:"kms_key_id"`

	PreferredSubnetID *string `json:"preferredSubnetID,omitempty" tf:"preferred_subnet_id"`

	SecurityGroupIds []string `json:"securityGroupIds,omitempty" tf:"security_group_ids"`

	SelfManagedActiveDirectory []SelfManagedActiveDirectoryParameters `json:"selfManagedActiveDirectory,omitempty" tf:"self_managed_active_directory"`

	SkipFinalBackup *bool `json:"skipFinalBackup,omitempty" tf:"skip_final_backup"`

	StorageCapacity int64 `json:"storageCapacity" tf:"storage_capacity"`

	StorageType *string `json:"storageType,omitempty" tf:"storage_type"`

	SubnetIds []string `json:"subnetIds" tf:"subnet_ids"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	ThroughputCapacity int64 `json:"throughputCapacity" tf:"throughput_capacity"`

	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime,omitempty" tf:"weekly_maintenance_start_time"`
}

type SelfManagedActiveDirectoryObservation struct {
}

type SelfManagedActiveDirectoryParameters struct {
	DNSIps []string `json:"dnsIps" tf:"dns_ips"`

	DomainName string `json:"domainName" tf:"domain_name"`

	FileSystemAdministratorsGroup *string `json:"fileSystemAdministratorsGroup,omitempty" tf:"file_system_administrators_group"`

	OrganizationalUnitDistinguishedName *string `json:"organizationalUnitDistinguishedName,omitempty" tf:"organizational_unit_distinguished_name"`

	Password string `json:"password" tf:"password"`

	Username string `json:"username" tf:"username"`
}

// FsxWindowsFileSystemSpec defines the desired state of FsxWindowsFileSystem
type FsxWindowsFileSystemSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       FsxWindowsFileSystemParameters `json:"forProvider"`
}

// FsxWindowsFileSystemStatus defines the observed state of FsxWindowsFileSystem.
type FsxWindowsFileSystemStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          FsxWindowsFileSystemObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FsxWindowsFileSystem is the Schema for the FsxWindowsFileSystems API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type FsxWindowsFileSystem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FsxWindowsFileSystemSpec   `json:"spec"`
	Status            FsxWindowsFileSystemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FsxWindowsFileSystemList contains a list of FsxWindowsFileSystems
type FsxWindowsFileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FsxWindowsFileSystem `json:"items"`
}

// Repository type metadata.
var (
	FsxWindowsFileSystemKind             = "FsxWindowsFileSystem"
	FsxWindowsFileSystemGroupKind        = schema.GroupKind{Group: Group, Kind: FsxWindowsFileSystemKind}.String()
	FsxWindowsFileSystemKindAPIVersion   = FsxWindowsFileSystemKind + "." + GroupVersion.String()
	FsxWindowsFileSystemGroupVersionKind = GroupVersion.WithKind(FsxWindowsFileSystemKind)
)

func init() {
	SchemeBuilder.Register(&FsxWindowsFileSystem{}, &FsxWindowsFileSystemList{})
}
