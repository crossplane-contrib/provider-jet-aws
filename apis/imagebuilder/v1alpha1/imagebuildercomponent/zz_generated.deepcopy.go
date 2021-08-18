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
func (in *ImagebuilderComponent) DeepCopyInto(out *ImagebuilderComponent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagebuilderComponent.
func (in *ImagebuilderComponent) DeepCopy() *ImagebuilderComponent {
	if in == nil {
		return nil
	}
	out := new(ImagebuilderComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImagebuilderComponent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagebuilderComponentList) DeepCopyInto(out *ImagebuilderComponentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImagebuilderComponent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagebuilderComponentList.
func (in *ImagebuilderComponentList) DeepCopy() *ImagebuilderComponentList {
	if in == nil {
		return nil
	}
	out := new(ImagebuilderComponentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImagebuilderComponentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagebuilderComponentObservation) DeepCopyInto(out *ImagebuilderComponentObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagebuilderComponentObservation.
func (in *ImagebuilderComponentObservation) DeepCopy() *ImagebuilderComponentObservation {
	if in == nil {
		return nil
	}
	out := new(ImagebuilderComponentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagebuilderComponentParameters) DeepCopyInto(out *ImagebuilderComponentParameters) {
	*out = *in
	if in.ChangeDescription != nil {
		in, out := &in.ChangeDescription, &out.ChangeDescription
		*out = new(string)
		**out = **in
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.KmsKeyId != nil {
		in, out := &in.KmsKeyId, &out.KmsKeyId
		*out = new(string)
		**out = **in
	}
	if in.SupportedOsVersions != nil {
		in, out := &in.SupportedOsVersions, &out.SupportedOsVersions
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
	if in.Uri != nil {
		in, out := &in.Uri, &out.Uri
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagebuilderComponentParameters.
func (in *ImagebuilderComponentParameters) DeepCopy() *ImagebuilderComponentParameters {
	if in == nil {
		return nil
	}
	out := new(ImagebuilderComponentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagebuilderComponentSpec) DeepCopyInto(out *ImagebuilderComponentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagebuilderComponentSpec.
func (in *ImagebuilderComponentSpec) DeepCopy() *ImagebuilderComponentSpec {
	if in == nil {
		return nil
	}
	out := new(ImagebuilderComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagebuilderComponentStatus) DeepCopyInto(out *ImagebuilderComponentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagebuilderComponentStatus.
func (in *ImagebuilderComponentStatus) DeepCopy() *ImagebuilderComponentStatus {
	if in == nil {
		return nil
	}
	out := new(ImagebuilderComponentStatus)
	in.DeepCopyInto(out)
	return out
}