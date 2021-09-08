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

type DbSnapshotObservation struct {
	AllocatedStorage int64 `json:"allocatedStorage" tf:"allocated_storage"`

	AvailabilityZone string `json:"availabilityZone" tf:"availability_zone"`

	DBSnapshotARN string `json:"dbSnapshotARN" tf:"db_snapshot_arn"`

	Encrypted bool `json:"encrypted" tf:"encrypted"`

	Engine string `json:"engine" tf:"engine"`

	EngineVersion string `json:"engineVersion" tf:"engine_version"`

	Iops int64 `json:"iops" tf:"iops"`

	KmsKeyID string `json:"kmsKeyID" tf:"kms_key_id"`

	LicenseModel string `json:"licenseModel" tf:"license_model"`

	OptionGroupName string `json:"optionGroupName" tf:"option_group_name"`

	Port int64 `json:"port" tf:"port"`

	SnapshotType string `json:"snapshotType" tf:"snapshot_type"`

	SourceDBSnapshotIdentifier string `json:"sourceDBSnapshotIdentifier" tf:"source_db_snapshot_identifier"`

	SourceRegion string `json:"sourceRegion" tf:"source_region"`

	Status string `json:"status" tf:"status"`

	StorageType string `json:"storageType" tf:"storage_type"`

	VPCID string `json:"vpcID" tf:"vpc_id"`
}

type DbSnapshotParameters struct {
	DBInstanceIdentifier string `json:"dbInstanceIdentifier" tf:"db_instance_identifier"`

	DBSnapshotIdentifier string `json:"dbSnapshotIdentifier" tf:"db_snapshot_identifier"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// DbSnapshotSpec defines the desired state of DbSnapshot
type DbSnapshotSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DbSnapshotParameters `json:"forProvider"`
}

// DbSnapshotStatus defines the observed state of DbSnapshot.
type DbSnapshotStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DbSnapshotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DbSnapshot is the Schema for the DbSnapshots API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type DbSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DbSnapshotSpec   `json:"spec"`
	Status            DbSnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DbSnapshotList contains a list of DbSnapshots
type DbSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DbSnapshot `json:"items"`
}

// Repository type metadata.
var (
	DbSnapshotKind             = "DbSnapshot"
	DbSnapshotGroupKind        = schema.GroupKind{Group: Group, Kind: DbSnapshotKind}.String()
	DbSnapshotKindAPIVersion   = DbSnapshotKind + "." + GroupVersion.String()
	DbSnapshotGroupVersionKind = GroupVersion.WithKind(DbSnapshotKind)
)

func init() {
	SchemeBuilder.Register(&DbSnapshot{}, &DbSnapshotList{})
}
