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

type EnclaveOptionsObservation struct {
}

type EnclaveOptionsParameters struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
}

type LaunchTemplateObservation struct {
}

type LaunchTemplateParameters struct {
	ID *string `json:"id,omitempty" tf:"id"`

	Name *string `json:"name,omitempty" tf:"name"`

	Version *string `json:"version,omitempty" tf:"version"`
}

type MetadataOptionsObservation struct {
}

type MetadataOptionsParameters struct {
	HTTPEndpoint *string `json:"httpEndpoint,omitempty" tf:"http_endpoint"`

	HTTPPutResponseHopLimit *int64 `json:"httpPutResponseHopLimit,omitempty" tf:"http_put_response_hop_limit"`

	HTTPTokens *string `json:"httpTokens,omitempty" tf:"http_tokens"`
}

type NetworkInterfaceObservation struct {
}

type NetworkInterfaceParameters struct {
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination"`

	DeviceIndex int64 `json:"deviceIndex" tf:"device_index"`

	NetworkInterfaceID string `json:"networkInterfaceID" tf:"network_interface_id"`
}

type SpotInstanceRequestEbsBlockDeviceObservation struct {
	VolumeID string `json:"volumeID" tf:"volume_id"`
}

type SpotInstanceRequestEbsBlockDeviceParameters struct {
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination"`

	DeviceName string `json:"deviceName" tf:"device_name"`

	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted"`

	Iops *int64 `json:"iops,omitempty" tf:"iops"`

	KmsKeyID *string `json:"kmsKeyID,omitempty" tf:"kms_key_id"`

	SnapshotID *string `json:"snapshotID,omitempty" tf:"snapshot_id"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	Throughput *int64 `json:"throughput,omitempty" tf:"throughput"`

	VolumeSize *int64 `json:"volumeSize,omitempty" tf:"volume_size"`

	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type"`
}

type SpotInstanceRequestEphemeralBlockDeviceObservation struct {
}

type SpotInstanceRequestEphemeralBlockDeviceParameters struct {
	DeviceName string `json:"deviceName" tf:"device_name"`

	NoDevice *bool `json:"noDevice,omitempty" tf:"no_device"`

	VirtualName *string `json:"virtualName,omitempty" tf:"virtual_name"`
}

type SpotInstanceRequestObservation struct {
	ARN string `json:"arn" tf:"arn"`

	InstanceState string `json:"instanceState" tf:"instance_state"`

	OutpostARN string `json:"outpostARN" tf:"outpost_arn"`

	PasswordData string `json:"passwordData" tf:"password_data"`

	PrimaryNetworkInterfaceID string `json:"primaryNetworkInterfaceID" tf:"primary_network_interface_id"`

	PrivateDNS string `json:"privateDNS" tf:"private_dns"`

	PublicDNS string `json:"publicDNS" tf:"public_dns"`

	PublicIP string `json:"publicIP" tf:"public_ip"`

	SpotBidStatus string `json:"spotBidStatus" tf:"spot_bid_status"`

	SpotInstanceID string `json:"spotInstanceID" tf:"spot_instance_id"`

	SpotRequestState string `json:"spotRequestState" tf:"spot_request_state"`
}

type SpotInstanceRequestParameters struct {
	Ami *string `json:"ami,omitempty" tf:"ami"`

	AssociatePublicIPAddress *bool `json:"associatePublicIPAddress,omitempty" tf:"associate_public_ip_address"`

	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone"`

	BlockDurationMinutes *int64 `json:"blockDurationMinutes,omitempty" tf:"block_duration_minutes"`

	CPUCoreCount *int64 `json:"cpuCoreCount,omitempty" tf:"cpu_core_count"`

	CPUThreadsPerCore *int64 `json:"cpuThreadsPerCore,omitempty" tf:"cpu_threads_per_core"`

	CapacityReservationSpecification []CapacityReservationSpecificationParameters `json:"capacityReservationSpecification,omitempty" tf:"capacity_reservation_specification"`

	CreditSpecification []CreditSpecificationParameters `json:"creditSpecification,omitempty" tf:"credit_specification"`

	DisableAPITermination *bool `json:"disableAPITermination,omitempty" tf:"disable_api_termination"`

	EbsBlockDevice []SpotInstanceRequestEbsBlockDeviceParameters `json:"ebsBlockDevice,omitempty" tf:"ebs_block_device"`

	EbsOptimized *bool `json:"ebsOptimized,omitempty" tf:"ebs_optimized"`

	EnclaveOptions []EnclaveOptionsParameters `json:"enclaveOptions,omitempty" tf:"enclave_options"`

	EphemeralBlockDevice []SpotInstanceRequestEphemeralBlockDeviceParameters `json:"ephemeralBlockDevice,omitempty" tf:"ephemeral_block_device"`

	GetPasswordData *bool `json:"getPasswordData,omitempty" tf:"get_password_data"`

	Hibernation *bool `json:"hibernation,omitempty" tf:"hibernation"`

	HostID *string `json:"hostID,omitempty" tf:"host_id"`

	IPv6AddressCount *int64 `json:"ipv6AddressCount,omitempty" tf:"ipv6_address_count"`

	IPv6Addresses []string `json:"ipv6Addresses,omitempty" tf:"ipv6_addresses"`

	IamInstanceProfile *string `json:"iamInstanceProfile,omitempty" tf:"iam_instance_profile"`

	InstanceInitiatedShutdownBehavior *string `json:"instanceInitiatedShutdownBehavior,omitempty" tf:"instance_initiated_shutdown_behavior"`

	InstanceInterruptionBehavior *string `json:"instanceInterruptionBehavior,omitempty" tf:"instance_interruption_behavior"`

	InstanceInterruptionBehaviour *string `json:"instanceInterruptionBehaviour,omitempty" tf:"instance_interruption_behaviour"`

	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type"`

	KeyName *string `json:"keyName,omitempty" tf:"key_name"`

	LaunchGroup *string `json:"launchGroup,omitempty" tf:"launch_group"`

	LaunchTemplate []LaunchTemplateParameters `json:"launchTemplate,omitempty" tf:"launch_template"`

	MetadataOptions []MetadataOptionsParameters `json:"metadataOptions,omitempty" tf:"metadata_options"`

	Monitoring *bool `json:"monitoring,omitempty" tf:"monitoring"`

	NetworkInterface []NetworkInterfaceParameters `json:"networkInterface,omitempty" tf:"network_interface"`

	PlacementGroup *string `json:"placementGroup,omitempty" tf:"placement_group"`

	PrivateIP *string `json:"privateIP,omitempty" tf:"private_ip"`

	RootBlockDevice []SpotInstanceRequestRootBlockDeviceParameters `json:"rootBlockDevice,omitempty" tf:"root_block_device"`

	SecondaryPrivateIps []string `json:"secondaryPrivateIps,omitempty" tf:"secondary_private_ips"`

	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups"`

	SourceDestCheck *bool `json:"sourceDestCheck,omitempty" tf:"source_dest_check"`

	SpotPrice *string `json:"spotPrice,omitempty" tf:"spot_price"`

	SpotType *string `json:"spotType,omitempty" tf:"spot_type"`

	SubnetID *string `json:"subnetID,omitempty" tf:"subnet_id"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	Tenancy *string `json:"tenancy,omitempty" tf:"tenancy"`

	UserData *string `json:"userData,omitempty" tf:"user_data"`

	UserDataBase64 *string `json:"userDataBase64,omitempty" tf:"user_data_base64"`

	VPCSecurityGroupIds []string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids"`

	ValidFrom *string `json:"validFrom,omitempty" tf:"valid_from"`

	ValidUntil *string `json:"validUntil,omitempty" tf:"valid_until"`

	VolumeTags map[string]string `json:"volumeTags,omitempty" tf:"volume_tags"`

	WaitForFulfillment *bool `json:"waitForFulfillment,omitempty" tf:"wait_for_fulfillment"`
}

type SpotInstanceRequestRootBlockDeviceObservation struct {
	DeviceName string `json:"deviceName" tf:"device_name"`

	VolumeID string `json:"volumeID" tf:"volume_id"`
}

type SpotInstanceRequestRootBlockDeviceParameters struct {
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination"`

	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted"`

	Iops *int64 `json:"iops,omitempty" tf:"iops"`

	KmsKeyID *string `json:"kmsKeyID,omitempty" tf:"kms_key_id"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	Throughput *int64 `json:"throughput,omitempty" tf:"throughput"`

	VolumeSize *int64 `json:"volumeSize,omitempty" tf:"volume_size"`

	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type"`
}

// SpotInstanceRequestSpec defines the desired state of SpotInstanceRequest
type SpotInstanceRequestSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SpotInstanceRequestParameters `json:"forProvider"`
}

// SpotInstanceRequestStatus defines the observed state of SpotInstanceRequest.
type SpotInstanceRequestStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SpotInstanceRequestObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SpotInstanceRequest is the Schema for the SpotInstanceRequests API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type SpotInstanceRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpotInstanceRequestSpec   `json:"spec"`
	Status            SpotInstanceRequestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpotInstanceRequestList contains a list of SpotInstanceRequests
type SpotInstanceRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpotInstanceRequest `json:"items"`
}

// Repository type metadata.
var (
	SpotInstanceRequestKind             = "SpotInstanceRequest"
	SpotInstanceRequestGroupKind        = schema.GroupKind{Group: Group, Kind: SpotInstanceRequestKind}.String()
	SpotInstanceRequestKindAPIVersion   = SpotInstanceRequestKind + "." + GroupVersion.String()
	SpotInstanceRequestGroupVersionKind = GroupVersion.WithKind(SpotInstanceRequestKind)
)

func init() {
	SchemeBuilder.Register(&SpotInstanceRequest{}, &SpotInstanceRequestList{})
}
