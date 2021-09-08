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

type BlockDeviceMappingsObservation struct {
}

type BlockDeviceMappingsParameters struct {
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name"`

	Ebs []EbsParameters `json:"ebs,omitempty" tf:"ebs"`

	NoDevice *string `json:"noDevice,omitempty" tf:"no_device"`

	VirtualName *string `json:"virtualName,omitempty" tf:"virtual_name"`
}

type CPUOptionsObservation struct {
}

type CPUOptionsParameters struct {
	CoreCount *int64 `json:"coreCount,omitempty" tf:"core_count"`

	ThreadsPerCore *int64 `json:"threadsPerCore,omitempty" tf:"threads_per_core"`
}

type CapacityReservationSpecificationObservation struct {
}

type CapacityReservationSpecificationParameters struct {
	CapacityReservationPreference *string `json:"capacityReservationPreference,omitempty" tf:"capacity_reservation_preference"`

	CapacityReservationTarget []CapacityReservationTargetParameters `json:"capacityReservationTarget,omitempty" tf:"capacity_reservation_target"`
}

type CapacityReservationTargetObservation struct {
}

type CapacityReservationTargetParameters struct {
	CapacityReservationID *string `json:"capacityReservationID,omitempty" tf:"capacity_reservation_id"`
}

type CreditSpecificationObservation struct {
}

type CreditSpecificationParameters struct {
	CPUCredits *string `json:"cpuCredits,omitempty" tf:"cpu_credits"`
}

type EbsObservation struct {
}

type EbsParameters struct {
	DeleteOnTermination *string `json:"deleteOnTermination,omitempty" tf:"delete_on_termination"`

	Encrypted *string `json:"encrypted,omitempty" tf:"encrypted"`

	Iops *int64 `json:"iops,omitempty" tf:"iops"`

	KmsKeyID *string `json:"kmsKeyID,omitempty" tf:"kms_key_id"`

	SnapshotID *string `json:"snapshotID,omitempty" tf:"snapshot_id"`

	Throughput *int64 `json:"throughput,omitempty" tf:"throughput"`

	VolumeSize *int64 `json:"volumeSize,omitempty" tf:"volume_size"`

	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type"`
}

type ElasticGpuSpecificationsObservation struct {
}

type ElasticGpuSpecificationsParameters struct {
	Type string `json:"type" tf:"type"`
}

type ElasticInferenceAcceleratorObservation struct {
}

type ElasticInferenceAcceleratorParameters struct {
	Type string `json:"type" tf:"type"`
}

type EnclaveOptionsObservation struct {
}

type EnclaveOptionsParameters struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
}

type HibernationOptionsObservation struct {
}

type HibernationOptionsParameters struct {
	Configured bool `json:"configured" tf:"configured"`
}

type IamInstanceProfileObservation struct {
}

type IamInstanceProfileParameters struct {
	ARN *string `json:"arn,omitempty" tf:"arn"`

	Name *string `json:"name,omitempty" tf:"name"`
}

type InstanceMarketOptionsObservation struct {
}

type InstanceMarketOptionsParameters struct {
	MarketType *string `json:"marketType,omitempty" tf:"market_type"`

	SpotOptions []SpotOptionsParameters `json:"spotOptions,omitempty" tf:"spot_options"`
}

type LaunchTemplateMetadataOptionsObservation struct {
}

type LaunchTemplateMetadataOptionsParameters struct {
	HTTPEndpoint *string `json:"httpEndpoint,omitempty" tf:"http_endpoint"`

	HTTPPutResponseHopLimit *int64 `json:"httpPutResponseHopLimit,omitempty" tf:"http_put_response_hop_limit"`

	HTTPTokens *string `json:"httpTokens,omitempty" tf:"http_tokens"`
}

type LaunchTemplateObservation struct {
	ARN string `json:"arn" tf:"arn"`

	LatestVersion int64 `json:"latestVersion" tf:"latest_version"`
}

type LaunchTemplateParameters struct {
	BlockDeviceMappings []BlockDeviceMappingsParameters `json:"blockDeviceMappings,omitempty" tf:"block_device_mappings"`

	CPUOptions []CPUOptionsParameters `json:"cpuOptions,omitempty" tf:"cpu_options"`

	CapacityReservationSpecification []CapacityReservationSpecificationParameters `json:"capacityReservationSpecification,omitempty" tf:"capacity_reservation_specification"`

	CreditSpecification []CreditSpecificationParameters `json:"creditSpecification,omitempty" tf:"credit_specification"`

	DefaultVersion *int64 `json:"defaultVersion,omitempty" tf:"default_version"`

	Description *string `json:"description,omitempty" tf:"description"`

	DisableAPITermination *bool `json:"disableAPITermination,omitempty" tf:"disable_api_termination"`

	EbsOptimized *string `json:"ebsOptimized,omitempty" tf:"ebs_optimized"`

	ElasticGpuSpecifications []ElasticGpuSpecificationsParameters `json:"elasticGpuSpecifications,omitempty" tf:"elastic_gpu_specifications"`

	ElasticInferenceAccelerator []ElasticInferenceAcceleratorParameters `json:"elasticInferenceAccelerator,omitempty" tf:"elastic_inference_accelerator"`

	EnclaveOptions []EnclaveOptionsParameters `json:"enclaveOptions,omitempty" tf:"enclave_options"`

	HibernationOptions []HibernationOptionsParameters `json:"hibernationOptions,omitempty" tf:"hibernation_options"`

	IamInstanceProfile []IamInstanceProfileParameters `json:"iamInstanceProfile,omitempty" tf:"iam_instance_profile"`

	ImageID *string `json:"imageID,omitempty" tf:"image_id"`

	InstanceInitiatedShutdownBehavior *string `json:"instanceInitiatedShutdownBehavior,omitempty" tf:"instance_initiated_shutdown_behavior"`

	InstanceMarketOptions []InstanceMarketOptionsParameters `json:"instanceMarketOptions,omitempty" tf:"instance_market_options"`

	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type"`

	KernelID *string `json:"kernelID,omitempty" tf:"kernel_id"`

	KeyName *string `json:"keyName,omitempty" tf:"key_name"`

	LicenseSpecification []LicenseSpecificationParameters `json:"licenseSpecification,omitempty" tf:"license_specification"`

	MetadataOptions []LaunchTemplateMetadataOptionsParameters `json:"metadataOptions,omitempty" tf:"metadata_options"`

	Monitoring []MonitoringParameters `json:"monitoring,omitempty" tf:"monitoring"`

	Name *string `json:"name,omitempty" tf:"name"`

	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`

	NetworkInterfaces []NetworkInterfacesParameters `json:"networkInterfaces,omitempty" tf:"network_interfaces"`

	Placement []PlacementParameters `json:"placement,omitempty" tf:"placement"`

	RAMDiskID *string `json:"ramDiskID,omitempty" tf:"ram_disk_id"`

	SecurityGroupNames []string `json:"securityGroupNames,omitempty" tf:"security_group_names"`

	TagSpecifications []TagSpecificationsParameters `json:"tagSpecifications,omitempty" tf:"tag_specifications"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	UpdateDefaultVersion *bool `json:"updateDefaultVersion,omitempty" tf:"update_default_version"`

	UserData *string `json:"userData,omitempty" tf:"user_data"`

	VPCSecurityGroupIds []string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids"`
}

type LicenseSpecificationObservation struct {
}

type LicenseSpecificationParameters struct {
	LicenseConfigurationARN string `json:"licenseConfigurationARN" tf:"license_configuration_arn"`
}

type MonitoringObservation struct {
}

type MonitoringParameters struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
}

type NetworkInterfacesObservation struct {
}

type NetworkInterfacesParameters struct {
	AssociateCarrierIPAddress *string `json:"associateCarrierIPAddress,omitempty" tf:"associate_carrier_ip_address"`

	AssociatePublicIPAddress *string `json:"associatePublicIPAddress,omitempty" tf:"associate_public_ip_address"`

	DeleteOnTermination *string `json:"deleteOnTermination,omitempty" tf:"delete_on_termination"`

	Description *string `json:"description,omitempty" tf:"description"`

	DeviceIndex *int64 `json:"deviceIndex,omitempty" tf:"device_index"`

	IPv4AddressCount *int64 `json:"ipv4AddressCount,omitempty" tf:"ipv4_address_count"`

	IPv4Addresses []string `json:"ipv4Addresses,omitempty" tf:"ipv4_addresses"`

	IPv6AddressCount *int64 `json:"ipv6AddressCount,omitempty" tf:"ipv6_address_count"`

	IPv6Addresses []string `json:"ipv6Addresses,omitempty" tf:"ipv6_addresses"`

	InterfaceType *string `json:"interfaceType,omitempty" tf:"interface_type"`

	NetworkInterfaceID *string `json:"networkInterfaceID,omitempty" tf:"network_interface_id"`

	PrivateIPAddress *string `json:"privateIPAddress,omitempty" tf:"private_ip_address"`

	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups"`

	SubnetID *string `json:"subnetID,omitempty" tf:"subnet_id"`
}

type PlacementObservation struct {
}

type PlacementParameters struct {
	Affinity *string `json:"affinity,omitempty" tf:"affinity"`

	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone"`

	GroupName *string `json:"groupName,omitempty" tf:"group_name"`

	HostID *string `json:"hostID,omitempty" tf:"host_id"`

	HostResourceGroupARN *string `json:"hostResourceGroupARN,omitempty" tf:"host_resource_group_arn"`

	PartitionNumber *int64 `json:"partitionNumber,omitempty" tf:"partition_number"`

	SpreadDomain *string `json:"spreadDomain,omitempty" tf:"spread_domain"`

	Tenancy *string `json:"tenancy,omitempty" tf:"tenancy"`
}

type SpotOptionsObservation struct {
}

type SpotOptionsParameters struct {
	BlockDurationMinutes *int64 `json:"blockDurationMinutes,omitempty" tf:"block_duration_minutes"`

	InstanceInterruptionBehavior *string `json:"instanceInterruptionBehavior,omitempty" tf:"instance_interruption_behavior"`

	MaxPrice *string `json:"maxPrice,omitempty" tf:"max_price"`

	SpotInstanceType *string `json:"spotInstanceType,omitempty" tf:"spot_instance_type"`

	ValidUntil *string `json:"validUntil,omitempty" tf:"valid_until"`
}

type TagSpecificationsObservation struct {
}

type TagSpecificationsParameters struct {
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
}

// LaunchTemplateSpec defines the desired state of LaunchTemplate
type LaunchTemplateSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       LaunchTemplateParameters `json:"forProvider"`
}

// LaunchTemplateStatus defines the observed state of LaunchTemplate.
type LaunchTemplateStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          LaunchTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LaunchTemplate is the Schema for the LaunchTemplates API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type LaunchTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LaunchTemplateSpec   `json:"spec"`
	Status            LaunchTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LaunchTemplateList contains a list of LaunchTemplates
type LaunchTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LaunchTemplate `json:"items"`
}

// Repository type metadata.
var (
	LaunchTemplateKind             = "LaunchTemplate"
	LaunchTemplateGroupKind        = schema.GroupKind{Group: Group, Kind: LaunchTemplateKind}.String()
	LaunchTemplateKindAPIVersion   = LaunchTemplateKind + "." + GroupVersion.String()
	LaunchTemplateGroupVersionKind = GroupVersion.WithKind(LaunchTemplateKind)
)

func init() {
	SchemeBuilder.Register(&LaunchTemplate{}, &LaunchTemplateList{})
}
