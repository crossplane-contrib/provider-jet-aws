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
func (in *AmiFromInstance) DeepCopyInto(out *AmiFromInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AmiFromInstance.
func (in *AmiFromInstance) DeepCopy() *AmiFromInstance {
	if in == nil {
		return nil
	}
	out := new(AmiFromInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AmiFromInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AmiFromInstanceList) DeepCopyInto(out *AmiFromInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AmiFromInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AmiFromInstanceList.
func (in *AmiFromInstanceList) DeepCopy() *AmiFromInstanceList {
	if in == nil {
		return nil
	}
	out := new(AmiFromInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AmiFromInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AmiFromInstanceObservation) DeepCopyInto(out *AmiFromInstanceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AmiFromInstanceObservation.
func (in *AmiFromInstanceObservation) DeepCopy() *AmiFromInstanceObservation {
	if in == nil {
		return nil
	}
	out := new(AmiFromInstanceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AmiFromInstanceParameters) DeepCopyInto(out *AmiFromInstanceParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EbsBlockDevice != nil {
		in, out := &in.EbsBlockDevice, &out.EbsBlockDevice
		*out = make([]EbsBlockDeviceParameters, len(*in))
		copy(*out, *in)
	}
	if in.EphemeralBlockDevice != nil {
		in, out := &in.EphemeralBlockDevice, &out.EphemeralBlockDevice
		*out = make([]EphemeralBlockDeviceParameters, len(*in))
		copy(*out, *in)
	}
	if in.SnapshotWithoutReboot != nil {
		in, out := &in.SnapshotWithoutReboot, &out.SnapshotWithoutReboot
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AmiFromInstanceParameters.
func (in *AmiFromInstanceParameters) DeepCopy() *AmiFromInstanceParameters {
	if in == nil {
		return nil
	}
	out := new(AmiFromInstanceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AmiFromInstanceSpec) DeepCopyInto(out *AmiFromInstanceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AmiFromInstanceSpec.
func (in *AmiFromInstanceSpec) DeepCopy() *AmiFromInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(AmiFromInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AmiFromInstanceStatus) DeepCopyInto(out *AmiFromInstanceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AmiFromInstanceStatus.
func (in *AmiFromInstanceStatus) DeepCopy() *AmiFromInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(AmiFromInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EbsBlockDeviceObservation) DeepCopyInto(out *EbsBlockDeviceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EbsBlockDeviceObservation.
func (in *EbsBlockDeviceObservation) DeepCopy() *EbsBlockDeviceObservation {
	if in == nil {
		return nil
	}
	out := new(EbsBlockDeviceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EbsBlockDeviceParameters) DeepCopyInto(out *EbsBlockDeviceParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EbsBlockDeviceParameters.
func (in *EbsBlockDeviceParameters) DeepCopy() *EbsBlockDeviceParameters {
	if in == nil {
		return nil
	}
	out := new(EbsBlockDeviceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EphemeralBlockDeviceObservation) DeepCopyInto(out *EphemeralBlockDeviceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EphemeralBlockDeviceObservation.
func (in *EphemeralBlockDeviceObservation) DeepCopy() *EphemeralBlockDeviceObservation {
	if in == nil {
		return nil
	}
	out := new(EphemeralBlockDeviceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EphemeralBlockDeviceParameters) DeepCopyInto(out *EphemeralBlockDeviceParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EphemeralBlockDeviceParameters.
func (in *EphemeralBlockDeviceParameters) DeepCopy() *EphemeralBlockDeviceParameters {
	if in == nil {
		return nil
	}
	out := new(EphemeralBlockDeviceParameters)
	in.DeepCopyInto(out)
	return out
}