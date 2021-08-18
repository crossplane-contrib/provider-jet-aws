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
func (in *VpcEndpointSubnetAssociation) DeepCopyInto(out *VpcEndpointSubnetAssociation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpointSubnetAssociation.
func (in *VpcEndpointSubnetAssociation) DeepCopy() *VpcEndpointSubnetAssociation {
	if in == nil {
		return nil
	}
	out := new(VpcEndpointSubnetAssociation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpcEndpointSubnetAssociation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcEndpointSubnetAssociationList) DeepCopyInto(out *VpcEndpointSubnetAssociationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VpcEndpointSubnetAssociation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpointSubnetAssociationList.
func (in *VpcEndpointSubnetAssociationList) DeepCopy() *VpcEndpointSubnetAssociationList {
	if in == nil {
		return nil
	}
	out := new(VpcEndpointSubnetAssociationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpcEndpointSubnetAssociationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcEndpointSubnetAssociationObservation) DeepCopyInto(out *VpcEndpointSubnetAssociationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpointSubnetAssociationObservation.
func (in *VpcEndpointSubnetAssociationObservation) DeepCopy() *VpcEndpointSubnetAssociationObservation {
	if in == nil {
		return nil
	}
	out := new(VpcEndpointSubnetAssociationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcEndpointSubnetAssociationParameters) DeepCopyInto(out *VpcEndpointSubnetAssociationParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpointSubnetAssociationParameters.
func (in *VpcEndpointSubnetAssociationParameters) DeepCopy() *VpcEndpointSubnetAssociationParameters {
	if in == nil {
		return nil
	}
	out := new(VpcEndpointSubnetAssociationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcEndpointSubnetAssociationSpec) DeepCopyInto(out *VpcEndpointSubnetAssociationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpointSubnetAssociationSpec.
func (in *VpcEndpointSubnetAssociationSpec) DeepCopy() *VpcEndpointSubnetAssociationSpec {
	if in == nil {
		return nil
	}
	out := new(VpcEndpointSubnetAssociationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcEndpointSubnetAssociationStatus) DeepCopyInto(out *VpcEndpointSubnetAssociationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpointSubnetAssociationStatus.
func (in *VpcEndpointSubnetAssociationStatus) DeepCopy() *VpcEndpointSubnetAssociationStatus {
	if in == nil {
		return nil
	}
	out := new(VpcEndpointSubnetAssociationStatus)
	in.DeepCopyInto(out)
	return out
}