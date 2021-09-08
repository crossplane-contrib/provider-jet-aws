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

type ColumnsObservation struct {
}

type ColumnsParameters struct {
	Comment *string `json:"comment,omitempty" tf:"comment"`

	Name string `json:"name" tf:"name"`

	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters"`

	Type *string `json:"type,omitempty" tf:"type"`
}

type GlueCatalogTableObservation struct {
	ARN string `json:"arn" tf:"arn"`
}

type GlueCatalogTableParameters struct {
	CatalogID *string `json:"catalogID,omitempty" tf:"catalog_id"`

	DatabaseName string `json:"databaseName" tf:"database_name"`

	Description *string `json:"description,omitempty" tf:"description"`

	Name string `json:"name" tf:"name"`

	Owner *string `json:"owner,omitempty" tf:"owner"`

	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters"`

	PartitionIndex []PartitionIndexParameters `json:"partitionIndex,omitempty" tf:"partition_index"`

	PartitionKeys []PartitionKeysParameters `json:"partitionKeys,omitempty" tf:"partition_keys"`

	Retention *int64 `json:"retention,omitempty" tf:"retention"`

	StorageDescriptor []StorageDescriptorParameters `json:"storageDescriptor,omitempty" tf:"storage_descriptor"`

	TableType *string `json:"tableType,omitempty" tf:"table_type"`

	TargetTable []TargetTableParameters `json:"targetTable,omitempty" tf:"target_table"`

	ViewExpandedText *string `json:"viewExpandedText,omitempty" tf:"view_expanded_text"`

	ViewOriginalText *string `json:"viewOriginalText,omitempty" tf:"view_original_text"`
}

type PartitionIndexObservation struct {
	IndexStatus string `json:"indexStatus" tf:"index_status"`
}

type PartitionIndexParameters struct {
	IndexName string `json:"indexName" tf:"index_name"`

	Keys []string `json:"keys" tf:"keys"`
}

type PartitionKeysObservation struct {
}

type PartitionKeysParameters struct {
	Comment *string `json:"comment,omitempty" tf:"comment"`

	Name string `json:"name" tf:"name"`

	Type *string `json:"type,omitempty" tf:"type"`
}

type SchemaIDObservation struct {
}

type SchemaIDParameters struct {
	RegistryName *string `json:"registryName,omitempty" tf:"registry_name"`

	SchemaARN *string `json:"schemaARN,omitempty" tf:"schema_arn"`

	SchemaName *string `json:"schemaName,omitempty" tf:"schema_name"`
}

type SchemaReferenceObservation struct {
}

type SchemaReferenceParameters struct {
	SchemaID []SchemaIDParameters `json:"schemaID,omitempty" tf:"schema_id"`

	SchemaVersionID *string `json:"schemaVersionID,omitempty" tf:"schema_version_id"`

	SchemaVersionNumber int64 `json:"schemaVersionNumber" tf:"schema_version_number"`
}

type SerDeInfoObservation struct {
}

type SerDeInfoParameters struct {
	Name *string `json:"name,omitempty" tf:"name"`

	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters"`

	SerializationLibrary *string `json:"serializationLibrary,omitempty" tf:"serialization_library"`
}

type SkewedInfoObservation struct {
}

type SkewedInfoParameters struct {
	SkewedColumnNames []string `json:"skewedColumnNames,omitempty" tf:"skewed_column_names"`

	SkewedColumnValueLocationMaps map[string]string `json:"skewedColumnValueLocationMaps,omitempty" tf:"skewed_column_value_location_maps"`

	SkewedColumnValues []string `json:"skewedColumnValues,omitempty" tf:"skewed_column_values"`
}

type SortColumnsObservation struct {
}

type SortColumnsParameters struct {
	Column string `json:"column" tf:"column"`

	SortOrder int64 `json:"sortOrder" tf:"sort_order"`
}

type StorageDescriptorObservation struct {
}

type StorageDescriptorParameters struct {
	BucketColumns []string `json:"bucketColumns,omitempty" tf:"bucket_columns"`

	Columns []ColumnsParameters `json:"columns,omitempty" tf:"columns"`

	Compressed *bool `json:"compressed,omitempty" tf:"compressed"`

	InputFormat *string `json:"inputFormat,omitempty" tf:"input_format"`

	Location *string `json:"location,omitempty" tf:"location"`

	NumberOfBuckets *int64 `json:"numberOfBuckets,omitempty" tf:"number_of_buckets"`

	OutputFormat *string `json:"outputFormat,omitempty" tf:"output_format"`

	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters"`

	SchemaReference []SchemaReferenceParameters `json:"schemaReference,omitempty" tf:"schema_reference"`

	SerDeInfo []SerDeInfoParameters `json:"serDeInfo,omitempty" tf:"ser_de_info"`

	SkewedInfo []SkewedInfoParameters `json:"skewedInfo,omitempty" tf:"skewed_info"`

	SortColumns []SortColumnsParameters `json:"sortColumns,omitempty" tf:"sort_columns"`

	StoredAsSubDirectories *bool `json:"storedAsSubDirectories,omitempty" tf:"stored_as_sub_directories"`
}

type TargetTableObservation struct {
}

type TargetTableParameters struct {
	CatalogID string `json:"catalogID" tf:"catalog_id"`

	DatabaseName string `json:"databaseName" tf:"database_name"`

	Name string `json:"name" tf:"name"`
}

// GlueCatalogTableSpec defines the desired state of GlueCatalogTable
type GlueCatalogTableSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       GlueCatalogTableParameters `json:"forProvider"`
}

// GlueCatalogTableStatus defines the observed state of GlueCatalogTable.
type GlueCatalogTableStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          GlueCatalogTableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GlueCatalogTable is the Schema for the GlueCatalogTables API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type GlueCatalogTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlueCatalogTableSpec   `json:"spec"`
	Status            GlueCatalogTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlueCatalogTableList contains a list of GlueCatalogTables
type GlueCatalogTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlueCatalogTable `json:"items"`
}

// Repository type metadata.
var (
	GlueCatalogTableKind             = "GlueCatalogTable"
	GlueCatalogTableGroupKind        = schema.GroupKind{Group: Group, Kind: GlueCatalogTableKind}.String()
	GlueCatalogTableKindAPIVersion   = GlueCatalogTableKind + "." + GroupVersion.String()
	GlueCatalogTableGroupVersionKind = GroupVersion.WithKind(GlueCatalogTableKind)
)

func init() {
	SchemeBuilder.Register(&GlueCatalogTable{}, &GlueCatalogTableList{})
}
