// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultNetworkAcl) DeepCopyInto(out *DefaultNetworkAcl) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultNetworkAcl.
func (in *DefaultNetworkAcl) DeepCopy() *DefaultNetworkAcl {
	if in == nil {
		return nil
	}
	out := new(DefaultNetworkAcl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DefaultNetworkAcl) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultNetworkAclList) DeepCopyInto(out *DefaultNetworkAclList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DefaultNetworkAcl, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultNetworkAclList.
func (in *DefaultNetworkAclList) DeepCopy() *DefaultNetworkAclList {
	if in == nil {
		return nil
	}
	out := new(DefaultNetworkAclList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DefaultNetworkAclList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultNetworkAclObservation) DeepCopyInto(out *DefaultNetworkAclObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultNetworkAclObservation.
func (in *DefaultNetworkAclObservation) DeepCopy() *DefaultNetworkAclObservation {
	if in == nil {
		return nil
	}
	out := new(DefaultNetworkAclObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultNetworkAclParameters) DeepCopyInto(out *DefaultNetworkAclParameters) {
	*out = *in
	if in.Egress != nil {
		in, out := &in.Egress, &out.Egress
		*out = make([]EgressParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = make([]IngressParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SubnetIds != nil {
		in, out := &in.SubnetIds, &out.SubnetIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultNetworkAclParameters.
func (in *DefaultNetworkAclParameters) DeepCopy() *DefaultNetworkAclParameters {
	if in == nil {
		return nil
	}
	out := new(DefaultNetworkAclParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultNetworkAclSpec) DeepCopyInto(out *DefaultNetworkAclSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultNetworkAclSpec.
func (in *DefaultNetworkAclSpec) DeepCopy() *DefaultNetworkAclSpec {
	if in == nil {
		return nil
	}
	out := new(DefaultNetworkAclSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultNetworkAclStatus) DeepCopyInto(out *DefaultNetworkAclStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultNetworkAclStatus.
func (in *DefaultNetworkAclStatus) DeepCopy() *DefaultNetworkAclStatus {
	if in == nil {
		return nil
	}
	out := new(DefaultNetworkAclStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultRouteTable) DeepCopyInto(out *DefaultRouteTable) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultRouteTable.
func (in *DefaultRouteTable) DeepCopy() *DefaultRouteTable {
	if in == nil {
		return nil
	}
	out := new(DefaultRouteTable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DefaultRouteTable) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultRouteTableList) DeepCopyInto(out *DefaultRouteTableList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DefaultRouteTable, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultRouteTableList.
func (in *DefaultRouteTableList) DeepCopy() *DefaultRouteTableList {
	if in == nil {
		return nil
	}
	out := new(DefaultRouteTableList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DefaultRouteTableList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultRouteTableObservation) DeepCopyInto(out *DefaultRouteTableObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultRouteTableObservation.
func (in *DefaultRouteTableObservation) DeepCopy() *DefaultRouteTableObservation {
	if in == nil {
		return nil
	}
	out := new(DefaultRouteTableObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultRouteTableParameters) DeepCopyInto(out *DefaultRouteTableParameters) {
	*out = *in
	if in.PropagatingVgws != nil {
		in, out := &in.PropagatingVgws, &out.PropagatingVgws
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Route != nil {
		in, out := &in.Route, &out.Route
		*out = make([]RouteParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultRouteTableParameters.
func (in *DefaultRouteTableParameters) DeepCopy() *DefaultRouteTableParameters {
	if in == nil {
		return nil
	}
	out := new(DefaultRouteTableParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultRouteTableSpec) DeepCopyInto(out *DefaultRouteTableSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultRouteTableSpec.
func (in *DefaultRouteTableSpec) DeepCopy() *DefaultRouteTableSpec {
	if in == nil {
		return nil
	}
	out := new(DefaultRouteTableSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultRouteTableStatus) DeepCopyInto(out *DefaultRouteTableStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultRouteTableStatus.
func (in *DefaultRouteTableStatus) DeepCopy() *DefaultRouteTableStatus {
	if in == nil {
		return nil
	}
	out := new(DefaultRouteTableStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultSecurityGroup) DeepCopyInto(out *DefaultSecurityGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultSecurityGroup.
func (in *DefaultSecurityGroup) DeepCopy() *DefaultSecurityGroup {
	if in == nil {
		return nil
	}
	out := new(DefaultSecurityGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DefaultSecurityGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultSecurityGroupEgressObservation) DeepCopyInto(out *DefaultSecurityGroupEgressObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultSecurityGroupEgressObservation.
func (in *DefaultSecurityGroupEgressObservation) DeepCopy() *DefaultSecurityGroupEgressObservation {
	if in == nil {
		return nil
	}
	out := new(DefaultSecurityGroupEgressObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultSecurityGroupEgressParameters) DeepCopyInto(out *DefaultSecurityGroupEgressParameters) {
	*out = *in
	if in.CidrBlocks != nil {
		in, out := &in.CidrBlocks, &out.CidrBlocks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.IPv6CidrBlocks != nil {
		in, out := &in.IPv6CidrBlocks, &out.IPv6CidrBlocks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PrefixListIds != nil {
		in, out := &in.PrefixListIds, &out.PrefixListIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Self != nil {
		in, out := &in.Self, &out.Self
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultSecurityGroupEgressParameters.
func (in *DefaultSecurityGroupEgressParameters) DeepCopy() *DefaultSecurityGroupEgressParameters {
	if in == nil {
		return nil
	}
	out := new(DefaultSecurityGroupEgressParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultSecurityGroupIngressObservation) DeepCopyInto(out *DefaultSecurityGroupIngressObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultSecurityGroupIngressObservation.
func (in *DefaultSecurityGroupIngressObservation) DeepCopy() *DefaultSecurityGroupIngressObservation {
	if in == nil {
		return nil
	}
	out := new(DefaultSecurityGroupIngressObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultSecurityGroupIngressParameters) DeepCopyInto(out *DefaultSecurityGroupIngressParameters) {
	*out = *in
	if in.CidrBlocks != nil {
		in, out := &in.CidrBlocks, &out.CidrBlocks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.IPv6CidrBlocks != nil {
		in, out := &in.IPv6CidrBlocks, &out.IPv6CidrBlocks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PrefixListIds != nil {
		in, out := &in.PrefixListIds, &out.PrefixListIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Self != nil {
		in, out := &in.Self, &out.Self
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultSecurityGroupIngressParameters.
func (in *DefaultSecurityGroupIngressParameters) DeepCopy() *DefaultSecurityGroupIngressParameters {
	if in == nil {
		return nil
	}
	out := new(DefaultSecurityGroupIngressParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultSecurityGroupList) DeepCopyInto(out *DefaultSecurityGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DefaultSecurityGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultSecurityGroupList.
func (in *DefaultSecurityGroupList) DeepCopy() *DefaultSecurityGroupList {
	if in == nil {
		return nil
	}
	out := new(DefaultSecurityGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DefaultSecurityGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultSecurityGroupObservation) DeepCopyInto(out *DefaultSecurityGroupObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultSecurityGroupObservation.
func (in *DefaultSecurityGroupObservation) DeepCopy() *DefaultSecurityGroupObservation {
	if in == nil {
		return nil
	}
	out := new(DefaultSecurityGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultSecurityGroupParameters) DeepCopyInto(out *DefaultSecurityGroupParameters) {
	*out = *in
	if in.Egress != nil {
		in, out := &in.Egress, &out.Egress
		*out = make([]DefaultSecurityGroupEgressParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = make([]DefaultSecurityGroupIngressParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RevokeRulesOnDelete != nil {
		in, out := &in.RevokeRulesOnDelete, &out.RevokeRulesOnDelete
		*out = new(bool)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultSecurityGroupParameters.
func (in *DefaultSecurityGroupParameters) DeepCopy() *DefaultSecurityGroupParameters {
	if in == nil {
		return nil
	}
	out := new(DefaultSecurityGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultSecurityGroupSpec) DeepCopyInto(out *DefaultSecurityGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultSecurityGroupSpec.
func (in *DefaultSecurityGroupSpec) DeepCopy() *DefaultSecurityGroupSpec {
	if in == nil {
		return nil
	}
	out := new(DefaultSecurityGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultSecurityGroupStatus) DeepCopyInto(out *DefaultSecurityGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultSecurityGroupStatus.
func (in *DefaultSecurityGroupStatus) DeepCopy() *DefaultSecurityGroupStatus {
	if in == nil {
		return nil
	}
	out := new(DefaultSecurityGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultSubnet) DeepCopyInto(out *DefaultSubnet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultSubnet.
func (in *DefaultSubnet) DeepCopy() *DefaultSubnet {
	if in == nil {
		return nil
	}
	out := new(DefaultSubnet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DefaultSubnet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultSubnetList) DeepCopyInto(out *DefaultSubnetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DefaultSubnet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultSubnetList.
func (in *DefaultSubnetList) DeepCopy() *DefaultSubnetList {
	if in == nil {
		return nil
	}
	out := new(DefaultSubnetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DefaultSubnetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultSubnetObservation) DeepCopyInto(out *DefaultSubnetObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultSubnetObservation.
func (in *DefaultSubnetObservation) DeepCopy() *DefaultSubnetObservation {
	if in == nil {
		return nil
	}
	out := new(DefaultSubnetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultSubnetParameters) DeepCopyInto(out *DefaultSubnetParameters) {
	*out = *in
	if in.CustomerOwnedIPv4Pool != nil {
		in, out := &in.CustomerOwnedIPv4Pool, &out.CustomerOwnedIPv4Pool
		*out = new(string)
		**out = **in
	}
	if in.MapCustomerOwnedIPOnLaunch != nil {
		in, out := &in.MapCustomerOwnedIPOnLaunch, &out.MapCustomerOwnedIPOnLaunch
		*out = new(bool)
		**out = **in
	}
	if in.MapPublicIPOnLaunch != nil {
		in, out := &in.MapPublicIPOnLaunch, &out.MapPublicIPOnLaunch
		*out = new(bool)
		**out = **in
	}
	if in.OutpostARN != nil {
		in, out := &in.OutpostARN, &out.OutpostARN
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultSubnetParameters.
func (in *DefaultSubnetParameters) DeepCopy() *DefaultSubnetParameters {
	if in == nil {
		return nil
	}
	out := new(DefaultSubnetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultSubnetSpec) DeepCopyInto(out *DefaultSubnetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultSubnetSpec.
func (in *DefaultSubnetSpec) DeepCopy() *DefaultSubnetSpec {
	if in == nil {
		return nil
	}
	out := new(DefaultSubnetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultSubnetStatus) DeepCopyInto(out *DefaultSubnetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultSubnetStatus.
func (in *DefaultSubnetStatus) DeepCopy() *DefaultSubnetStatus {
	if in == nil {
		return nil
	}
	out := new(DefaultSubnetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultVpc) DeepCopyInto(out *DefaultVpc) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultVpc.
func (in *DefaultVpc) DeepCopy() *DefaultVpc {
	if in == nil {
		return nil
	}
	out := new(DefaultVpc)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DefaultVpc) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultVpcDhcpOptions) DeepCopyInto(out *DefaultVpcDhcpOptions) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultVpcDhcpOptions.
func (in *DefaultVpcDhcpOptions) DeepCopy() *DefaultVpcDhcpOptions {
	if in == nil {
		return nil
	}
	out := new(DefaultVpcDhcpOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DefaultVpcDhcpOptions) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultVpcDhcpOptionsList) DeepCopyInto(out *DefaultVpcDhcpOptionsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DefaultVpcDhcpOptions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultVpcDhcpOptionsList.
func (in *DefaultVpcDhcpOptionsList) DeepCopy() *DefaultVpcDhcpOptionsList {
	if in == nil {
		return nil
	}
	out := new(DefaultVpcDhcpOptionsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DefaultVpcDhcpOptionsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultVpcDhcpOptionsObservation) DeepCopyInto(out *DefaultVpcDhcpOptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultVpcDhcpOptionsObservation.
func (in *DefaultVpcDhcpOptionsObservation) DeepCopy() *DefaultVpcDhcpOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(DefaultVpcDhcpOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultVpcDhcpOptionsParameters) DeepCopyInto(out *DefaultVpcDhcpOptionsParameters) {
	*out = *in
	if in.NetbiosNameServers != nil {
		in, out := &in.NetbiosNameServers, &out.NetbiosNameServers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NetbiosNodeType != nil {
		in, out := &in.NetbiosNodeType, &out.NetbiosNodeType
		*out = new(string)
		**out = **in
	}
	if in.OwnerID != nil {
		in, out := &in.OwnerID, &out.OwnerID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultVpcDhcpOptionsParameters.
func (in *DefaultVpcDhcpOptionsParameters) DeepCopy() *DefaultVpcDhcpOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(DefaultVpcDhcpOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultVpcDhcpOptionsSpec) DeepCopyInto(out *DefaultVpcDhcpOptionsSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultVpcDhcpOptionsSpec.
func (in *DefaultVpcDhcpOptionsSpec) DeepCopy() *DefaultVpcDhcpOptionsSpec {
	if in == nil {
		return nil
	}
	out := new(DefaultVpcDhcpOptionsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultVpcDhcpOptionsStatus) DeepCopyInto(out *DefaultVpcDhcpOptionsStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultVpcDhcpOptionsStatus.
func (in *DefaultVpcDhcpOptionsStatus) DeepCopy() *DefaultVpcDhcpOptionsStatus {
	if in == nil {
		return nil
	}
	out := new(DefaultVpcDhcpOptionsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultVpcList) DeepCopyInto(out *DefaultVpcList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DefaultVpc, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultVpcList.
func (in *DefaultVpcList) DeepCopy() *DefaultVpcList {
	if in == nil {
		return nil
	}
	out := new(DefaultVpcList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DefaultVpcList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultVpcObservation) DeepCopyInto(out *DefaultVpcObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultVpcObservation.
func (in *DefaultVpcObservation) DeepCopy() *DefaultVpcObservation {
	if in == nil {
		return nil
	}
	out := new(DefaultVpcObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultVpcParameters) DeepCopyInto(out *DefaultVpcParameters) {
	*out = *in
	if in.EnableClassiclink != nil {
		in, out := &in.EnableClassiclink, &out.EnableClassiclink
		*out = new(bool)
		**out = **in
	}
	if in.EnableClassiclinkDNSSupport != nil {
		in, out := &in.EnableClassiclinkDNSSupport, &out.EnableClassiclinkDNSSupport
		*out = new(bool)
		**out = **in
	}
	if in.EnableDNSHostnames != nil {
		in, out := &in.EnableDNSHostnames, &out.EnableDNSHostnames
		*out = new(bool)
		**out = **in
	}
	if in.EnableDNSSupport != nil {
		in, out := &in.EnableDNSSupport, &out.EnableDNSSupport
		*out = new(bool)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultVpcParameters.
func (in *DefaultVpcParameters) DeepCopy() *DefaultVpcParameters {
	if in == nil {
		return nil
	}
	out := new(DefaultVpcParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultVpcSpec) DeepCopyInto(out *DefaultVpcSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultVpcSpec.
func (in *DefaultVpcSpec) DeepCopy() *DefaultVpcSpec {
	if in == nil {
		return nil
	}
	out := new(DefaultVpcSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultVpcStatus) DeepCopyInto(out *DefaultVpcStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultVpcStatus.
func (in *DefaultVpcStatus) DeepCopy() *DefaultVpcStatus {
	if in == nil {
		return nil
	}
	out := new(DefaultVpcStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgressObservation) DeepCopyInto(out *EgressObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgressObservation.
func (in *EgressObservation) DeepCopy() *EgressObservation {
	if in == nil {
		return nil
	}
	out := new(EgressObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgressParameters) DeepCopyInto(out *EgressParameters) {
	*out = *in
	if in.CidrBlock != nil {
		in, out := &in.CidrBlock, &out.CidrBlock
		*out = new(string)
		**out = **in
	}
	if in.IPv6CidrBlock != nil {
		in, out := &in.IPv6CidrBlock, &out.IPv6CidrBlock
		*out = new(string)
		**out = **in
	}
	if in.IcmpCode != nil {
		in, out := &in.IcmpCode, &out.IcmpCode
		*out = new(int64)
		**out = **in
	}
	if in.IcmpType != nil {
		in, out := &in.IcmpType, &out.IcmpType
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgressParameters.
func (in *EgressParameters) DeepCopy() *EgressParameters {
	if in == nil {
		return nil
	}
	out := new(EgressParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressObservation) DeepCopyInto(out *IngressObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressObservation.
func (in *IngressObservation) DeepCopy() *IngressObservation {
	if in == nil {
		return nil
	}
	out := new(IngressObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressParameters) DeepCopyInto(out *IngressParameters) {
	*out = *in
	if in.CidrBlock != nil {
		in, out := &in.CidrBlock, &out.CidrBlock
		*out = new(string)
		**out = **in
	}
	if in.IPv6CidrBlock != nil {
		in, out := &in.IPv6CidrBlock, &out.IPv6CidrBlock
		*out = new(string)
		**out = **in
	}
	if in.IcmpCode != nil {
		in, out := &in.IcmpCode, &out.IcmpCode
		*out = new(int64)
		**out = **in
	}
	if in.IcmpType != nil {
		in, out := &in.IcmpType, &out.IcmpType
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressParameters.
func (in *IngressParameters) DeepCopy() *IngressParameters {
	if in == nil {
		return nil
	}
	out := new(IngressParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteObservation) DeepCopyInto(out *RouteObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteObservation.
func (in *RouteObservation) DeepCopy() *RouteObservation {
	if in == nil {
		return nil
	}
	out := new(RouteObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteParameters) DeepCopyInto(out *RouteParameters) {
	*out = *in
	if in.CidrBlock != nil {
		in, out := &in.CidrBlock, &out.CidrBlock
		*out = new(string)
		**out = **in
	}
	if in.DestinationPrefixListID != nil {
		in, out := &in.DestinationPrefixListID, &out.DestinationPrefixListID
		*out = new(string)
		**out = **in
	}
	if in.EgressOnlyGatewayID != nil {
		in, out := &in.EgressOnlyGatewayID, &out.EgressOnlyGatewayID
		*out = new(string)
		**out = **in
	}
	if in.GatewayID != nil {
		in, out := &in.GatewayID, &out.GatewayID
		*out = new(string)
		**out = **in
	}
	if in.IPv6CidrBlock != nil {
		in, out := &in.IPv6CidrBlock, &out.IPv6CidrBlock
		*out = new(string)
		**out = **in
	}
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(string)
		**out = **in
	}
	if in.NatGatewayID != nil {
		in, out := &in.NatGatewayID, &out.NatGatewayID
		*out = new(string)
		**out = **in
	}
	if in.NetworkInterfaceID != nil {
		in, out := &in.NetworkInterfaceID, &out.NetworkInterfaceID
		*out = new(string)
		**out = **in
	}
	if in.TransitGatewayID != nil {
		in, out := &in.TransitGatewayID, &out.TransitGatewayID
		*out = new(string)
		**out = **in
	}
	if in.VPCEndpointID != nil {
		in, out := &in.VPCEndpointID, &out.VPCEndpointID
		*out = new(string)
		**out = **in
	}
	if in.VPCPeeringConnectionID != nil {
		in, out := &in.VPCPeeringConnectionID, &out.VPCPeeringConnectionID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteParameters.
func (in *RouteParameters) DeepCopy() *RouteParameters {
	if in == nil {
		return nil
	}
	out := new(RouteParameters)
	in.DeepCopyInto(out)
	return out
}
