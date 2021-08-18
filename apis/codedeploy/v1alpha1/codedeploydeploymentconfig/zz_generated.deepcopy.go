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
func (in *CodedeployDeploymentConfig) DeepCopyInto(out *CodedeployDeploymentConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CodedeployDeploymentConfig.
func (in *CodedeployDeploymentConfig) DeepCopy() *CodedeployDeploymentConfig {
	if in == nil {
		return nil
	}
	out := new(CodedeployDeploymentConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CodedeployDeploymentConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CodedeployDeploymentConfigList) DeepCopyInto(out *CodedeployDeploymentConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CodedeployDeploymentConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CodedeployDeploymentConfigList.
func (in *CodedeployDeploymentConfigList) DeepCopy() *CodedeployDeploymentConfigList {
	if in == nil {
		return nil
	}
	out := new(CodedeployDeploymentConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CodedeployDeploymentConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CodedeployDeploymentConfigObservation) DeepCopyInto(out *CodedeployDeploymentConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CodedeployDeploymentConfigObservation.
func (in *CodedeployDeploymentConfigObservation) DeepCopy() *CodedeployDeploymentConfigObservation {
	if in == nil {
		return nil
	}
	out := new(CodedeployDeploymentConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CodedeployDeploymentConfigParameters) DeepCopyInto(out *CodedeployDeploymentConfigParameters) {
	*out = *in
	if in.ComputePlatform != nil {
		in, out := &in.ComputePlatform, &out.ComputePlatform
		*out = new(string)
		**out = **in
	}
	if in.MinimumHealthyHosts != nil {
		in, out := &in.MinimumHealthyHosts, &out.MinimumHealthyHosts
		*out = make([]MinimumHealthyHostsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TrafficRoutingConfig != nil {
		in, out := &in.TrafficRoutingConfig, &out.TrafficRoutingConfig
		*out = make([]TrafficRoutingConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CodedeployDeploymentConfigParameters.
func (in *CodedeployDeploymentConfigParameters) DeepCopy() *CodedeployDeploymentConfigParameters {
	if in == nil {
		return nil
	}
	out := new(CodedeployDeploymentConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CodedeployDeploymentConfigSpec) DeepCopyInto(out *CodedeployDeploymentConfigSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CodedeployDeploymentConfigSpec.
func (in *CodedeployDeploymentConfigSpec) DeepCopy() *CodedeployDeploymentConfigSpec {
	if in == nil {
		return nil
	}
	out := new(CodedeployDeploymentConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CodedeployDeploymentConfigStatus) DeepCopyInto(out *CodedeployDeploymentConfigStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CodedeployDeploymentConfigStatus.
func (in *CodedeployDeploymentConfigStatus) DeepCopy() *CodedeployDeploymentConfigStatus {
	if in == nil {
		return nil
	}
	out := new(CodedeployDeploymentConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinimumHealthyHostsObservation) DeepCopyInto(out *MinimumHealthyHostsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinimumHealthyHostsObservation.
func (in *MinimumHealthyHostsObservation) DeepCopy() *MinimumHealthyHostsObservation {
	if in == nil {
		return nil
	}
	out := new(MinimumHealthyHostsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinimumHealthyHostsParameters) DeepCopyInto(out *MinimumHealthyHostsParameters) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinimumHealthyHostsParameters.
func (in *MinimumHealthyHostsParameters) DeepCopy() *MinimumHealthyHostsParameters {
	if in == nil {
		return nil
	}
	out := new(MinimumHealthyHostsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeBasedCanaryObservation) DeepCopyInto(out *TimeBasedCanaryObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeBasedCanaryObservation.
func (in *TimeBasedCanaryObservation) DeepCopy() *TimeBasedCanaryObservation {
	if in == nil {
		return nil
	}
	out := new(TimeBasedCanaryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeBasedCanaryParameters) DeepCopyInto(out *TimeBasedCanaryParameters) {
	*out = *in
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(int64)
		**out = **in
	}
	if in.Percentage != nil {
		in, out := &in.Percentage, &out.Percentage
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeBasedCanaryParameters.
func (in *TimeBasedCanaryParameters) DeepCopy() *TimeBasedCanaryParameters {
	if in == nil {
		return nil
	}
	out := new(TimeBasedCanaryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeBasedLinearObservation) DeepCopyInto(out *TimeBasedLinearObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeBasedLinearObservation.
func (in *TimeBasedLinearObservation) DeepCopy() *TimeBasedLinearObservation {
	if in == nil {
		return nil
	}
	out := new(TimeBasedLinearObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeBasedLinearParameters) DeepCopyInto(out *TimeBasedLinearParameters) {
	*out = *in
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(int64)
		**out = **in
	}
	if in.Percentage != nil {
		in, out := &in.Percentage, &out.Percentage
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeBasedLinearParameters.
func (in *TimeBasedLinearParameters) DeepCopy() *TimeBasedLinearParameters {
	if in == nil {
		return nil
	}
	out := new(TimeBasedLinearParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrafficRoutingConfigObservation) DeepCopyInto(out *TrafficRoutingConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficRoutingConfigObservation.
func (in *TrafficRoutingConfigObservation) DeepCopy() *TrafficRoutingConfigObservation {
	if in == nil {
		return nil
	}
	out := new(TrafficRoutingConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrafficRoutingConfigParameters) DeepCopyInto(out *TrafficRoutingConfigParameters) {
	*out = *in
	if in.TimeBasedCanary != nil {
		in, out := &in.TimeBasedCanary, &out.TimeBasedCanary
		*out = make([]TimeBasedCanaryParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TimeBasedLinear != nil {
		in, out := &in.TimeBasedLinear, &out.TimeBasedLinear
		*out = make([]TimeBasedLinearParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficRoutingConfigParameters.
func (in *TrafficRoutingConfigParameters) DeepCopy() *TrafficRoutingConfigParameters {
	if in == nil {
		return nil
	}
	out := new(TrafficRoutingConfigParameters)
	in.DeepCopyInto(out)
	return out
}