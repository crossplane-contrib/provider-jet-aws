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

type EgressObservation struct {
}

type EgressParameters struct {
	Action string `json:"action" tf:"action"`

	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block"`

	FromPort int64 `json:"fromPort" tf:"from_port"`

	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block"`

	IcmpCode *int64 `json:"icmpCode,omitempty" tf:"icmp_code"`

	IcmpType *int64 `json:"icmpType,omitempty" tf:"icmp_type"`

	Protocol string `json:"protocol" tf:"protocol"`

	RuleNo int64 `json:"ruleNo" tf:"rule_no"`

	ToPort int64 `json:"toPort" tf:"to_port"`
}

type IngressObservation struct {
}

type IngressParameters struct {
	Action string `json:"action" tf:"action"`

	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block"`

	FromPort int64 `json:"fromPort" tf:"from_port"`

	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block"`

	IcmpCode *int64 `json:"icmpCode,omitempty" tf:"icmp_code"`

	IcmpType *int64 `json:"icmpType,omitempty" tf:"icmp_type"`

	Protocol string `json:"protocol" tf:"protocol"`

	RuleNo int64 `json:"ruleNo" tf:"rule_no"`

	ToPort int64 `json:"toPort" tf:"to_port"`
}

type NetworkAclObservation struct {
	ARN string `json:"arn" tf:"arn"`

	OwnerID string `json:"ownerID" tf:"owner_id"`
}

type NetworkAclParameters struct {
	Egress []EgressParameters `json:"egress,omitempty" tf:"egress"`

	Ingress []IngressParameters `json:"ingress,omitempty" tf:"ingress"`

	SubnetIds []string `json:"subnetIds,omitempty" tf:"subnet_ids"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	VPCID string `json:"vpcID" tf:"vpc_id"`
}

// NetworkAclSpec defines the desired state of NetworkAcl
type NetworkAclSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       NetworkAclParameters `json:"forProvider"`
}

// NetworkAclStatus defines the observed state of NetworkAcl.
type NetworkAclStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          NetworkAclObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkAcl is the Schema for the NetworkAcls API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type NetworkAcl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkAclSpec   `json:"spec"`
	Status            NetworkAclStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkAclList contains a list of NetworkAcls
type NetworkAclList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkAcl `json:"items"`
}

// Repository type metadata.
var (
	NetworkAclKind             = "NetworkAcl"
	NetworkAclGroupKind        = schema.GroupKind{Group: Group, Kind: NetworkAclKind}.String()
	NetworkAclKindAPIVersion   = NetworkAclKind + "." + GroupVersion.String()
	NetworkAclGroupVersionKind = GroupVersion.WithKind(NetworkAclKind)
)

func init() {
	SchemeBuilder.Register(&NetworkAcl{}, &NetworkAclList{})
}
