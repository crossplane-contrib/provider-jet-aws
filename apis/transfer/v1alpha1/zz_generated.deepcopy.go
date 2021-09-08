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
func (in *EndpointDetailsObservation) DeepCopyInto(out *EndpointDetailsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointDetailsObservation.
func (in *EndpointDetailsObservation) DeepCopy() *EndpointDetailsObservation {
	if in == nil {
		return nil
	}
	out := new(EndpointDetailsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointDetailsParameters) DeepCopyInto(out *EndpointDetailsParameters) {
	*out = *in
	if in.AddressAllocationIds != nil {
		in, out := &in.AddressAllocationIds, &out.AddressAllocationIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroupIds != nil {
		in, out := &in.SecurityGroupIds, &out.SecurityGroupIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SubnetIds != nil {
		in, out := &in.SubnetIds, &out.SubnetIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VPCEndpointID != nil {
		in, out := &in.VPCEndpointID, &out.VPCEndpointID
		*out = new(string)
		**out = **in
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointDetailsParameters.
func (in *EndpointDetailsParameters) DeepCopy() *EndpointDetailsParameters {
	if in == nil {
		return nil
	}
	out := new(EndpointDetailsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HomeDirectoryMappingsObservation) DeepCopyInto(out *HomeDirectoryMappingsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HomeDirectoryMappingsObservation.
func (in *HomeDirectoryMappingsObservation) DeepCopy() *HomeDirectoryMappingsObservation {
	if in == nil {
		return nil
	}
	out := new(HomeDirectoryMappingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HomeDirectoryMappingsParameters) DeepCopyInto(out *HomeDirectoryMappingsParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HomeDirectoryMappingsParameters.
func (in *HomeDirectoryMappingsParameters) DeepCopy() *HomeDirectoryMappingsParameters {
	if in == nil {
		return nil
	}
	out := new(HomeDirectoryMappingsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PosixProfileObservation) DeepCopyInto(out *PosixProfileObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PosixProfileObservation.
func (in *PosixProfileObservation) DeepCopy() *PosixProfileObservation {
	if in == nil {
		return nil
	}
	out := new(PosixProfileObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PosixProfileParameters) DeepCopyInto(out *PosixProfileParameters) {
	*out = *in
	if in.SecondaryGids != nil {
		in, out := &in.SecondaryGids, &out.SecondaryGids
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PosixProfileParameters.
func (in *PosixProfileParameters) DeepCopy() *PosixProfileParameters {
	if in == nil {
		return nil
	}
	out := new(PosixProfileParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferServer) DeepCopyInto(out *TransferServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferServer.
func (in *TransferServer) DeepCopy() *TransferServer {
	if in == nil {
		return nil
	}
	out := new(TransferServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TransferServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferServerList) DeepCopyInto(out *TransferServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TransferServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferServerList.
func (in *TransferServerList) DeepCopy() *TransferServerList {
	if in == nil {
		return nil
	}
	out := new(TransferServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TransferServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferServerObservation) DeepCopyInto(out *TransferServerObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferServerObservation.
func (in *TransferServerObservation) DeepCopy() *TransferServerObservation {
	if in == nil {
		return nil
	}
	out := new(TransferServerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferServerParameters) DeepCopyInto(out *TransferServerParameters) {
	*out = *in
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(string)
		**out = **in
	}
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.EndpointDetails != nil {
		in, out := &in.EndpointDetails, &out.EndpointDetails
		*out = make([]EndpointDetailsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EndpointType != nil {
		in, out := &in.EndpointType, &out.EndpointType
		*out = new(string)
		**out = **in
	}
	if in.ForceDestroy != nil {
		in, out := &in.ForceDestroy, &out.ForceDestroy
		*out = new(bool)
		**out = **in
	}
	if in.HostKey != nil {
		in, out := &in.HostKey, &out.HostKey
		*out = new(string)
		**out = **in
	}
	if in.IdentityProviderType != nil {
		in, out := &in.IdentityProviderType, &out.IdentityProviderType
		*out = new(string)
		**out = **in
	}
	if in.InvocationRole != nil {
		in, out := &in.InvocationRole, &out.InvocationRole
		*out = new(string)
		**out = **in
	}
	if in.LoggingRole != nil {
		in, out := &in.LoggingRole, &out.LoggingRole
		*out = new(string)
		**out = **in
	}
	if in.Protocols != nil {
		in, out := &in.Protocols, &out.Protocols
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityPolicyName != nil {
		in, out := &in.SecurityPolicyName, &out.SecurityPolicyName
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
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferServerParameters.
func (in *TransferServerParameters) DeepCopy() *TransferServerParameters {
	if in == nil {
		return nil
	}
	out := new(TransferServerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferServerSpec) DeepCopyInto(out *TransferServerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferServerSpec.
func (in *TransferServerSpec) DeepCopy() *TransferServerSpec {
	if in == nil {
		return nil
	}
	out := new(TransferServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferServerStatus) DeepCopyInto(out *TransferServerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferServerStatus.
func (in *TransferServerStatus) DeepCopy() *TransferServerStatus {
	if in == nil {
		return nil
	}
	out := new(TransferServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferSshKey) DeepCopyInto(out *TransferSshKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferSshKey.
func (in *TransferSshKey) DeepCopy() *TransferSshKey {
	if in == nil {
		return nil
	}
	out := new(TransferSshKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TransferSshKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferSshKeyList) DeepCopyInto(out *TransferSshKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TransferSshKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferSshKeyList.
func (in *TransferSshKeyList) DeepCopy() *TransferSshKeyList {
	if in == nil {
		return nil
	}
	out := new(TransferSshKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TransferSshKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferSshKeyObservation) DeepCopyInto(out *TransferSshKeyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferSshKeyObservation.
func (in *TransferSshKeyObservation) DeepCopy() *TransferSshKeyObservation {
	if in == nil {
		return nil
	}
	out := new(TransferSshKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferSshKeyParameters) DeepCopyInto(out *TransferSshKeyParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferSshKeyParameters.
func (in *TransferSshKeyParameters) DeepCopy() *TransferSshKeyParameters {
	if in == nil {
		return nil
	}
	out := new(TransferSshKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferSshKeySpec) DeepCopyInto(out *TransferSshKeySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferSshKeySpec.
func (in *TransferSshKeySpec) DeepCopy() *TransferSshKeySpec {
	if in == nil {
		return nil
	}
	out := new(TransferSshKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferSshKeyStatus) DeepCopyInto(out *TransferSshKeyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferSshKeyStatus.
func (in *TransferSshKeyStatus) DeepCopy() *TransferSshKeyStatus {
	if in == nil {
		return nil
	}
	out := new(TransferSshKeyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferUser) DeepCopyInto(out *TransferUser) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferUser.
func (in *TransferUser) DeepCopy() *TransferUser {
	if in == nil {
		return nil
	}
	out := new(TransferUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TransferUser) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferUserList) DeepCopyInto(out *TransferUserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TransferUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferUserList.
func (in *TransferUserList) DeepCopy() *TransferUserList {
	if in == nil {
		return nil
	}
	out := new(TransferUserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TransferUserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferUserObservation) DeepCopyInto(out *TransferUserObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferUserObservation.
func (in *TransferUserObservation) DeepCopy() *TransferUserObservation {
	if in == nil {
		return nil
	}
	out := new(TransferUserObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferUserParameters) DeepCopyInto(out *TransferUserParameters) {
	*out = *in
	if in.HomeDirectory != nil {
		in, out := &in.HomeDirectory, &out.HomeDirectory
		*out = new(string)
		**out = **in
	}
	if in.HomeDirectoryMappings != nil {
		in, out := &in.HomeDirectoryMappings, &out.HomeDirectoryMappings
		*out = make([]HomeDirectoryMappingsParameters, len(*in))
		copy(*out, *in)
	}
	if in.HomeDirectoryType != nil {
		in, out := &in.HomeDirectoryType, &out.HomeDirectoryType
		*out = new(string)
		**out = **in
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
	if in.PosixProfile != nil {
		in, out := &in.PosixProfile, &out.PosixProfile
		*out = make([]PosixProfileParameters, len(*in))
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferUserParameters.
func (in *TransferUserParameters) DeepCopy() *TransferUserParameters {
	if in == nil {
		return nil
	}
	out := new(TransferUserParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferUserSpec) DeepCopyInto(out *TransferUserSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferUserSpec.
func (in *TransferUserSpec) DeepCopy() *TransferUserSpec {
	if in == nil {
		return nil
	}
	out := new(TransferUserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferUserStatus) DeepCopyInto(out *TransferUserStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferUserStatus.
func (in *TransferUserStatus) DeepCopy() *TransferUserStatus {
	if in == nil {
		return nil
	}
	out := new(TransferUserStatus)
	in.DeepCopyInto(out)
	return out
}
