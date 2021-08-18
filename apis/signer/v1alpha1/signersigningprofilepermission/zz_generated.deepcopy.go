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
func (in *SignerSigningProfilePermission) DeepCopyInto(out *SignerSigningProfilePermission) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignerSigningProfilePermission.
func (in *SignerSigningProfilePermission) DeepCopy() *SignerSigningProfilePermission {
	if in == nil {
		return nil
	}
	out := new(SignerSigningProfilePermission)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SignerSigningProfilePermission) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignerSigningProfilePermissionList) DeepCopyInto(out *SignerSigningProfilePermissionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SignerSigningProfilePermission, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignerSigningProfilePermissionList.
func (in *SignerSigningProfilePermissionList) DeepCopy() *SignerSigningProfilePermissionList {
	if in == nil {
		return nil
	}
	out := new(SignerSigningProfilePermissionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SignerSigningProfilePermissionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignerSigningProfilePermissionObservation) DeepCopyInto(out *SignerSigningProfilePermissionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignerSigningProfilePermissionObservation.
func (in *SignerSigningProfilePermissionObservation) DeepCopy() *SignerSigningProfilePermissionObservation {
	if in == nil {
		return nil
	}
	out := new(SignerSigningProfilePermissionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignerSigningProfilePermissionParameters) DeepCopyInto(out *SignerSigningProfilePermissionParameters) {
	*out = *in
	if in.ProfileVersion != nil {
		in, out := &in.ProfileVersion, &out.ProfileVersion
		*out = new(string)
		**out = **in
	}
	if in.StatementId != nil {
		in, out := &in.StatementId, &out.StatementId
		*out = new(string)
		**out = **in
	}
	if in.StatementIdPrefix != nil {
		in, out := &in.StatementIdPrefix, &out.StatementIdPrefix
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignerSigningProfilePermissionParameters.
func (in *SignerSigningProfilePermissionParameters) DeepCopy() *SignerSigningProfilePermissionParameters {
	if in == nil {
		return nil
	}
	out := new(SignerSigningProfilePermissionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignerSigningProfilePermissionSpec) DeepCopyInto(out *SignerSigningProfilePermissionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignerSigningProfilePermissionSpec.
func (in *SignerSigningProfilePermissionSpec) DeepCopy() *SignerSigningProfilePermissionSpec {
	if in == nil {
		return nil
	}
	out := new(SignerSigningProfilePermissionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignerSigningProfilePermissionStatus) DeepCopyInto(out *SignerSigningProfilePermissionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignerSigningProfilePermissionStatus.
func (in *SignerSigningProfilePermissionStatus) DeepCopy() *SignerSigningProfilePermissionStatus {
	if in == nil {
		return nil
	}
	out := new(SignerSigningProfilePermissionStatus)
	in.DeepCopyInto(out)
	return out
}