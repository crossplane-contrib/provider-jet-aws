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
func (in *AccepterObservation) DeepCopyInto(out *AccepterObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccepterObservation.
func (in *AccepterObservation) DeepCopy() *AccepterObservation {
	if in == nil {
		return nil
	}
	out := new(AccepterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccepterParameters) DeepCopyInto(out *AccepterParameters) {
	*out = *in
	if in.AllowClassicLinkToRemoteVpc != nil {
		in, out := &in.AllowClassicLinkToRemoteVpc, &out.AllowClassicLinkToRemoteVpc
		*out = new(bool)
		**out = **in
	}
	if in.AllowRemoteVpcDnsResolution != nil {
		in, out := &in.AllowRemoteVpcDnsResolution, &out.AllowRemoteVpcDnsResolution
		*out = new(bool)
		**out = **in
	}
	if in.AllowVpcToRemoteClassicLink != nil {
		in, out := &in.AllowVpcToRemoteClassicLink, &out.AllowVpcToRemoteClassicLink
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccepterParameters.
func (in *AccepterParameters) DeepCopy() *AccepterParameters {
	if in == nil {
		return nil
	}
	out := new(AccepterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequesterObservation) DeepCopyInto(out *RequesterObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequesterObservation.
func (in *RequesterObservation) DeepCopy() *RequesterObservation {
	if in == nil {
		return nil
	}
	out := new(RequesterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequesterParameters) DeepCopyInto(out *RequesterParameters) {
	*out = *in
	if in.AllowClassicLinkToRemoteVpc != nil {
		in, out := &in.AllowClassicLinkToRemoteVpc, &out.AllowClassicLinkToRemoteVpc
		*out = new(bool)
		**out = **in
	}
	if in.AllowRemoteVpcDnsResolution != nil {
		in, out := &in.AllowRemoteVpcDnsResolution, &out.AllowRemoteVpcDnsResolution
		*out = new(bool)
		**out = **in
	}
	if in.AllowVpcToRemoteClassicLink != nil {
		in, out := &in.AllowVpcToRemoteClassicLink, &out.AllowVpcToRemoteClassicLink
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequesterParameters.
func (in *RequesterParameters) DeepCopy() *RequesterParameters {
	if in == nil {
		return nil
	}
	out := new(RequesterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcPeeringConnection) DeepCopyInto(out *VpcPeeringConnection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcPeeringConnection.
func (in *VpcPeeringConnection) DeepCopy() *VpcPeeringConnection {
	if in == nil {
		return nil
	}
	out := new(VpcPeeringConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpcPeeringConnection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcPeeringConnectionList) DeepCopyInto(out *VpcPeeringConnectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VpcPeeringConnection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcPeeringConnectionList.
func (in *VpcPeeringConnectionList) DeepCopy() *VpcPeeringConnectionList {
	if in == nil {
		return nil
	}
	out := new(VpcPeeringConnectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpcPeeringConnectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcPeeringConnectionObservation) DeepCopyInto(out *VpcPeeringConnectionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcPeeringConnectionObservation.
func (in *VpcPeeringConnectionObservation) DeepCopy() *VpcPeeringConnectionObservation {
	if in == nil {
		return nil
	}
	out := new(VpcPeeringConnectionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcPeeringConnectionParameters) DeepCopyInto(out *VpcPeeringConnectionParameters) {
	*out = *in
	if in.Accepter != nil {
		in, out := &in.Accepter, &out.Accepter
		*out = make([]AccepterParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AutoAccept != nil {
		in, out := &in.AutoAccept, &out.AutoAccept
		*out = new(bool)
		**out = **in
	}
	if in.PeerOwnerId != nil {
		in, out := &in.PeerOwnerId, &out.PeerOwnerId
		*out = new(string)
		**out = **in
	}
	if in.PeerRegion != nil {
		in, out := &in.PeerRegion, &out.PeerRegion
		*out = new(string)
		**out = **in
	}
	if in.Requester != nil {
		in, out := &in.Requester, &out.Requester
		*out = make([]RequesterParameters, len(*in))
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcPeeringConnectionParameters.
func (in *VpcPeeringConnectionParameters) DeepCopy() *VpcPeeringConnectionParameters {
	if in == nil {
		return nil
	}
	out := new(VpcPeeringConnectionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcPeeringConnectionSpec) DeepCopyInto(out *VpcPeeringConnectionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcPeeringConnectionSpec.
func (in *VpcPeeringConnectionSpec) DeepCopy() *VpcPeeringConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(VpcPeeringConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcPeeringConnectionStatus) DeepCopyInto(out *VpcPeeringConnectionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcPeeringConnectionStatus.
func (in *VpcPeeringConnectionStatus) DeepCopy() *VpcPeeringConnectionStatus {
	if in == nil {
		return nil
	}
	out := new(VpcPeeringConnectionStatus)
	in.DeepCopyInto(out)
	return out
}