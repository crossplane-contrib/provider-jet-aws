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
func (in *FieldToMatchObservation) DeepCopyInto(out *FieldToMatchObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FieldToMatchObservation.
func (in *FieldToMatchObservation) DeepCopy() *FieldToMatchObservation {
	if in == nil {
		return nil
	}
	out := new(FieldToMatchObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FieldToMatchParameters) DeepCopyInto(out *FieldToMatchParameters) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FieldToMatchParameters.
func (in *FieldToMatchParameters) DeepCopy() *FieldToMatchParameters {
	if in == nil {
		return nil
	}
	out := new(FieldToMatchParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegexMatchTupleObservation) DeepCopyInto(out *RegexMatchTupleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegexMatchTupleObservation.
func (in *RegexMatchTupleObservation) DeepCopy() *RegexMatchTupleObservation {
	if in == nil {
		return nil
	}
	out := new(RegexMatchTupleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegexMatchTupleParameters) DeepCopyInto(out *RegexMatchTupleParameters) {
	*out = *in
	if in.FieldToMatch != nil {
		in, out := &in.FieldToMatch, &out.FieldToMatch
		*out = make([]FieldToMatchParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegexMatchTupleParameters.
func (in *RegexMatchTupleParameters) DeepCopy() *RegexMatchTupleParameters {
	if in == nil {
		return nil
	}
	out := new(RegexMatchTupleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafregionalRegexMatchSet) DeepCopyInto(out *WafregionalRegexMatchSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafregionalRegexMatchSet.
func (in *WafregionalRegexMatchSet) DeepCopy() *WafregionalRegexMatchSet {
	if in == nil {
		return nil
	}
	out := new(WafregionalRegexMatchSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WafregionalRegexMatchSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafregionalRegexMatchSetList) DeepCopyInto(out *WafregionalRegexMatchSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WafregionalRegexMatchSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafregionalRegexMatchSetList.
func (in *WafregionalRegexMatchSetList) DeepCopy() *WafregionalRegexMatchSetList {
	if in == nil {
		return nil
	}
	out := new(WafregionalRegexMatchSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WafregionalRegexMatchSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafregionalRegexMatchSetObservation) DeepCopyInto(out *WafregionalRegexMatchSetObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafregionalRegexMatchSetObservation.
func (in *WafregionalRegexMatchSetObservation) DeepCopy() *WafregionalRegexMatchSetObservation {
	if in == nil {
		return nil
	}
	out := new(WafregionalRegexMatchSetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafregionalRegexMatchSetParameters) DeepCopyInto(out *WafregionalRegexMatchSetParameters) {
	*out = *in
	if in.RegexMatchTuple != nil {
		in, out := &in.RegexMatchTuple, &out.RegexMatchTuple
		*out = make([]RegexMatchTupleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafregionalRegexMatchSetParameters.
func (in *WafregionalRegexMatchSetParameters) DeepCopy() *WafregionalRegexMatchSetParameters {
	if in == nil {
		return nil
	}
	out := new(WafregionalRegexMatchSetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafregionalRegexMatchSetSpec) DeepCopyInto(out *WafregionalRegexMatchSetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafregionalRegexMatchSetSpec.
func (in *WafregionalRegexMatchSetSpec) DeepCopy() *WafregionalRegexMatchSetSpec {
	if in == nil {
		return nil
	}
	out := new(WafregionalRegexMatchSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafregionalRegexMatchSetStatus) DeepCopyInto(out *WafregionalRegexMatchSetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafregionalRegexMatchSetStatus.
func (in *WafregionalRegexMatchSetStatus) DeepCopy() *WafregionalRegexMatchSetStatus {
	if in == nil {
		return nil
	}
	out := new(WafregionalRegexMatchSetStatus)
	in.DeepCopyInto(out)
	return out
}