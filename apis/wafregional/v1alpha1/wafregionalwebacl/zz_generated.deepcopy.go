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
func (in *ActionObservation) DeepCopyInto(out *ActionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionObservation.
func (in *ActionObservation) DeepCopy() *ActionObservation {
	if in == nil {
		return nil
	}
	out := new(ActionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionParameters) DeepCopyInto(out *ActionParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionParameters.
func (in *ActionParameters) DeepCopy() *ActionParameters {
	if in == nil {
		return nil
	}
	out := new(ActionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultActionObservation) DeepCopyInto(out *DefaultActionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultActionObservation.
func (in *DefaultActionObservation) DeepCopy() *DefaultActionObservation {
	if in == nil {
		return nil
	}
	out := new(DefaultActionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultActionParameters) DeepCopyInto(out *DefaultActionParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultActionParameters.
func (in *DefaultActionParameters) DeepCopy() *DefaultActionParameters {
	if in == nil {
		return nil
	}
	out := new(DefaultActionParameters)
	in.DeepCopyInto(out)
	return out
}

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
func (in *LoggingConfigurationObservation) DeepCopyInto(out *LoggingConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingConfigurationObservation.
func (in *LoggingConfigurationObservation) DeepCopy() *LoggingConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(LoggingConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingConfigurationParameters) DeepCopyInto(out *LoggingConfigurationParameters) {
	*out = *in
	if in.RedactedFields != nil {
		in, out := &in.RedactedFields, &out.RedactedFields
		*out = make([]RedactedFieldsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingConfigurationParameters.
func (in *LoggingConfigurationParameters) DeepCopy() *LoggingConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(LoggingConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OverrideActionObservation) DeepCopyInto(out *OverrideActionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OverrideActionObservation.
func (in *OverrideActionObservation) DeepCopy() *OverrideActionObservation {
	if in == nil {
		return nil
	}
	out := new(OverrideActionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OverrideActionParameters) DeepCopyInto(out *OverrideActionParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OverrideActionParameters.
func (in *OverrideActionParameters) DeepCopy() *OverrideActionParameters {
	if in == nil {
		return nil
	}
	out := new(OverrideActionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedactedFieldsObservation) DeepCopyInto(out *RedactedFieldsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedactedFieldsObservation.
func (in *RedactedFieldsObservation) DeepCopy() *RedactedFieldsObservation {
	if in == nil {
		return nil
	}
	out := new(RedactedFieldsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedactedFieldsParameters) DeepCopyInto(out *RedactedFieldsParameters) {
	*out = *in
	if in.FieldToMatch != nil {
		in, out := &in.FieldToMatch, &out.FieldToMatch
		*out = make([]FieldToMatchParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedactedFieldsParameters.
func (in *RedactedFieldsParameters) DeepCopy() *RedactedFieldsParameters {
	if in == nil {
		return nil
	}
	out := new(RedactedFieldsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleObservation) DeepCopyInto(out *RuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleObservation.
func (in *RuleObservation) DeepCopy() *RuleObservation {
	if in == nil {
		return nil
	}
	out := new(RuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleParameters) DeepCopyInto(out *RuleParameters) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = make([]ActionParameters, len(*in))
		copy(*out, *in)
	}
	if in.OverrideAction != nil {
		in, out := &in.OverrideAction, &out.OverrideAction
		*out = make([]OverrideActionParameters, len(*in))
		copy(*out, *in)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleParameters.
func (in *RuleParameters) DeepCopy() *RuleParameters {
	if in == nil {
		return nil
	}
	out := new(RuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafregionalWebAcl) DeepCopyInto(out *WafregionalWebAcl) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafregionalWebAcl.
func (in *WafregionalWebAcl) DeepCopy() *WafregionalWebAcl {
	if in == nil {
		return nil
	}
	out := new(WafregionalWebAcl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WafregionalWebAcl) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafregionalWebAclList) DeepCopyInto(out *WafregionalWebAclList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WafregionalWebAcl, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafregionalWebAclList.
func (in *WafregionalWebAclList) DeepCopy() *WafregionalWebAclList {
	if in == nil {
		return nil
	}
	out := new(WafregionalWebAclList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WafregionalWebAclList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafregionalWebAclObservation) DeepCopyInto(out *WafregionalWebAclObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafregionalWebAclObservation.
func (in *WafregionalWebAclObservation) DeepCopy() *WafregionalWebAclObservation {
	if in == nil {
		return nil
	}
	out := new(WafregionalWebAclObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafregionalWebAclParameters) DeepCopyInto(out *WafregionalWebAclParameters) {
	*out = *in
	if in.DefaultAction != nil {
		in, out := &in.DefaultAction, &out.DefaultAction
		*out = make([]DefaultActionParameters, len(*in))
		copy(*out, *in)
	}
	if in.LoggingConfiguration != nil {
		in, out := &in.LoggingConfiguration, &out.LoggingConfiguration
		*out = make([]LoggingConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Rule != nil {
		in, out := &in.Rule, &out.Rule
		*out = make([]RuleParameters, len(*in))
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafregionalWebAclParameters.
func (in *WafregionalWebAclParameters) DeepCopy() *WafregionalWebAclParameters {
	if in == nil {
		return nil
	}
	out := new(WafregionalWebAclParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafregionalWebAclSpec) DeepCopyInto(out *WafregionalWebAclSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafregionalWebAclSpec.
func (in *WafregionalWebAclSpec) DeepCopy() *WafregionalWebAclSpec {
	if in == nil {
		return nil
	}
	out := new(WafregionalWebAclSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafregionalWebAclStatus) DeepCopyInto(out *WafregionalWebAclStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafregionalWebAclStatus.
func (in *WafregionalWebAclStatus) DeepCopy() *WafregionalWebAclStatus {
	if in == nil {
		return nil
	}
	out := new(WafregionalWebAclStatus)
	in.DeepCopyInto(out)
	return out
}