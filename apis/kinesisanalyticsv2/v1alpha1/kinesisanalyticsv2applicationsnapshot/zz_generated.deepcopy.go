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
func (in *Kinesisanalyticsv2ApplicationSnapshot) DeepCopyInto(out *Kinesisanalyticsv2ApplicationSnapshot) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kinesisanalyticsv2ApplicationSnapshot.
func (in *Kinesisanalyticsv2ApplicationSnapshot) DeepCopy() *Kinesisanalyticsv2ApplicationSnapshot {
	if in == nil {
		return nil
	}
	out := new(Kinesisanalyticsv2ApplicationSnapshot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Kinesisanalyticsv2ApplicationSnapshot) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kinesisanalyticsv2ApplicationSnapshotList) DeepCopyInto(out *Kinesisanalyticsv2ApplicationSnapshotList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Kinesisanalyticsv2ApplicationSnapshot, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kinesisanalyticsv2ApplicationSnapshotList.
func (in *Kinesisanalyticsv2ApplicationSnapshotList) DeepCopy() *Kinesisanalyticsv2ApplicationSnapshotList {
	if in == nil {
		return nil
	}
	out := new(Kinesisanalyticsv2ApplicationSnapshotList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Kinesisanalyticsv2ApplicationSnapshotList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kinesisanalyticsv2ApplicationSnapshotObservation) DeepCopyInto(out *Kinesisanalyticsv2ApplicationSnapshotObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kinesisanalyticsv2ApplicationSnapshotObservation.
func (in *Kinesisanalyticsv2ApplicationSnapshotObservation) DeepCopy() *Kinesisanalyticsv2ApplicationSnapshotObservation {
	if in == nil {
		return nil
	}
	out := new(Kinesisanalyticsv2ApplicationSnapshotObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kinesisanalyticsv2ApplicationSnapshotParameters) DeepCopyInto(out *Kinesisanalyticsv2ApplicationSnapshotParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kinesisanalyticsv2ApplicationSnapshotParameters.
func (in *Kinesisanalyticsv2ApplicationSnapshotParameters) DeepCopy() *Kinesisanalyticsv2ApplicationSnapshotParameters {
	if in == nil {
		return nil
	}
	out := new(Kinesisanalyticsv2ApplicationSnapshotParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kinesisanalyticsv2ApplicationSnapshotSpec) DeepCopyInto(out *Kinesisanalyticsv2ApplicationSnapshotSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kinesisanalyticsv2ApplicationSnapshotSpec.
func (in *Kinesisanalyticsv2ApplicationSnapshotSpec) DeepCopy() *Kinesisanalyticsv2ApplicationSnapshotSpec {
	if in == nil {
		return nil
	}
	out := new(Kinesisanalyticsv2ApplicationSnapshotSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kinesisanalyticsv2ApplicationSnapshotStatus) DeepCopyInto(out *Kinesisanalyticsv2ApplicationSnapshotStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kinesisanalyticsv2ApplicationSnapshotStatus.
func (in *Kinesisanalyticsv2ApplicationSnapshotStatus) DeepCopy() *Kinesisanalyticsv2ApplicationSnapshotStatus {
	if in == nil {
		return nil
	}
	out := new(Kinesisanalyticsv2ApplicationSnapshotStatus)
	in.DeepCopyInto(out)
	return out
}