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
func (in *ComputeEnvironment) DeepCopyInto(out *ComputeEnvironment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeEnvironment.
func (in *ComputeEnvironment) DeepCopy() *ComputeEnvironment {
	if in == nil {
		return nil
	}
	out := new(ComputeEnvironment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComputeEnvironment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeEnvironmentList) DeepCopyInto(out *ComputeEnvironmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ComputeEnvironment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeEnvironmentList.
func (in *ComputeEnvironmentList) DeepCopy() *ComputeEnvironmentList {
	if in == nil {
		return nil
	}
	out := new(ComputeEnvironmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComputeEnvironmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeEnvironmentObservation) DeepCopyInto(out *ComputeEnvironmentObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.EcsClusterArn != nil {
		in, out := &in.EcsClusterArn, &out.EcsClusterArn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.StatusReason != nil {
		in, out := &in.StatusReason, &out.StatusReason
		*out = new(string)
		**out = **in
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeEnvironmentObservation.
func (in *ComputeEnvironmentObservation) DeepCopy() *ComputeEnvironmentObservation {
	if in == nil {
		return nil
	}
	out := new(ComputeEnvironmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeEnvironmentParameters) DeepCopyInto(out *ComputeEnvironmentParameters) {
	*out = *in
	if in.ComputeEnvironmentName != nil {
		in, out := &in.ComputeEnvironmentName, &out.ComputeEnvironmentName
		*out = new(string)
		**out = **in
	}
	if in.ComputeEnvironmentNamePrefix != nil {
		in, out := &in.ComputeEnvironmentNamePrefix, &out.ComputeEnvironmentNamePrefix
		*out = new(string)
		**out = **in
	}
	if in.ComputeResources != nil {
		in, out := &in.ComputeResources, &out.ComputeResources
		*out = make([]ComputeResourcesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ServiceRole != nil {
		in, out := &in.ServiceRole, &out.ServiceRole
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeEnvironmentParameters.
func (in *ComputeEnvironmentParameters) DeepCopy() *ComputeEnvironmentParameters {
	if in == nil {
		return nil
	}
	out := new(ComputeEnvironmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeEnvironmentSpec) DeepCopyInto(out *ComputeEnvironmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeEnvironmentSpec.
func (in *ComputeEnvironmentSpec) DeepCopy() *ComputeEnvironmentSpec {
	if in == nil {
		return nil
	}
	out := new(ComputeEnvironmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeEnvironmentStatus) DeepCopyInto(out *ComputeEnvironmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeEnvironmentStatus.
func (in *ComputeEnvironmentStatus) DeepCopy() *ComputeEnvironmentStatus {
	if in == nil {
		return nil
	}
	out := new(ComputeEnvironmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeResourcesObservation) DeepCopyInto(out *ComputeResourcesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeResourcesObservation.
func (in *ComputeResourcesObservation) DeepCopy() *ComputeResourcesObservation {
	if in == nil {
		return nil
	}
	out := new(ComputeResourcesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeResourcesParameters) DeepCopyInto(out *ComputeResourcesParameters) {
	*out = *in
	if in.AllocationStrategy != nil {
		in, out := &in.AllocationStrategy, &out.AllocationStrategy
		*out = new(string)
		**out = **in
	}
	if in.BidPercentage != nil {
		in, out := &in.BidPercentage, &out.BidPercentage
		*out = new(int64)
		**out = **in
	}
	if in.DesiredVcpus != nil {
		in, out := &in.DesiredVcpus, &out.DesiredVcpus
		*out = new(int64)
		**out = **in
	}
	if in.EC2KeyPair != nil {
		in, out := &in.EC2KeyPair, &out.EC2KeyPair
		*out = new(string)
		**out = **in
	}
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.InstanceRole != nil {
		in, out := &in.InstanceRole, &out.InstanceRole
		*out = new(string)
		**out = **in
	}
	if in.InstanceType != nil {
		in, out := &in.InstanceType, &out.InstanceType
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.LaunchTemplate != nil {
		in, out := &in.LaunchTemplate, &out.LaunchTemplate
		*out = make([]LaunchTemplateParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MaxVcpus != nil {
		in, out := &in.MaxVcpus, &out.MaxVcpus
		*out = new(int64)
		**out = **in
	}
	if in.MinVcpus != nil {
		in, out := &in.MinVcpus, &out.MinVcpus
		*out = new(int64)
		**out = **in
	}
	if in.SecurityGroupIds != nil {
		in, out := &in.SecurityGroupIds, &out.SecurityGroupIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SpotIAMFleetRole != nil {
		in, out := &in.SpotIAMFleetRole, &out.SpotIAMFleetRole
		*out = new(string)
		**out = **in
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeResourcesParameters.
func (in *ComputeResourcesParameters) DeepCopy() *ComputeResourcesParameters {
	if in == nil {
		return nil
	}
	out := new(ComputeResourcesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EvaluateOnExitObservation) DeepCopyInto(out *EvaluateOnExitObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EvaluateOnExitObservation.
func (in *EvaluateOnExitObservation) DeepCopy() *EvaluateOnExitObservation {
	if in == nil {
		return nil
	}
	out := new(EvaluateOnExitObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EvaluateOnExitParameters) DeepCopyInto(out *EvaluateOnExitParameters) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.OnExitCode != nil {
		in, out := &in.OnExitCode, &out.OnExitCode
		*out = new(string)
		**out = **in
	}
	if in.OnReason != nil {
		in, out := &in.OnReason, &out.OnReason
		*out = new(string)
		**out = **in
	}
	if in.OnStatusReason != nil {
		in, out := &in.OnStatusReason, &out.OnStatusReason
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EvaluateOnExitParameters.
func (in *EvaluateOnExitParameters) DeepCopy() *EvaluateOnExitParameters {
	if in == nil {
		return nil
	}
	out := new(EvaluateOnExitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobDefinition) DeepCopyInto(out *JobDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobDefinition.
func (in *JobDefinition) DeepCopy() *JobDefinition {
	if in == nil {
		return nil
	}
	out := new(JobDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JobDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobDefinitionList) DeepCopyInto(out *JobDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JobDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobDefinitionList.
func (in *JobDefinitionList) DeepCopy() *JobDefinitionList {
	if in == nil {
		return nil
	}
	out := new(JobDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JobDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobDefinitionObservation) DeepCopyInto(out *JobDefinitionObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(int64)
		**out = **in
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobDefinitionObservation.
func (in *JobDefinitionObservation) DeepCopy() *JobDefinitionObservation {
	if in == nil {
		return nil
	}
	out := new(JobDefinitionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobDefinitionParameters) DeepCopyInto(out *JobDefinitionParameters) {
	*out = *in
	if in.ContainerProperties != nil {
		in, out := &in.ContainerProperties, &out.ContainerProperties
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.PlatformCapabilities != nil {
		in, out := &in.PlatformCapabilities, &out.PlatformCapabilities
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PropagateTags != nil {
		in, out := &in.PropagateTags, &out.PropagateTags
		*out = new(bool)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RetryStrategy != nil {
		in, out := &in.RetryStrategy, &out.RetryStrategy
		*out = make([]RetryStrategyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = make([]TimeoutParameters, len(*in))
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobDefinitionParameters.
func (in *JobDefinitionParameters) DeepCopy() *JobDefinitionParameters {
	if in == nil {
		return nil
	}
	out := new(JobDefinitionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobDefinitionSpec) DeepCopyInto(out *JobDefinitionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobDefinitionSpec.
func (in *JobDefinitionSpec) DeepCopy() *JobDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(JobDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobDefinitionStatus) DeepCopyInto(out *JobDefinitionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobDefinitionStatus.
func (in *JobDefinitionStatus) DeepCopy() *JobDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(JobDefinitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobQueue) DeepCopyInto(out *JobQueue) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobQueue.
func (in *JobQueue) DeepCopy() *JobQueue {
	if in == nil {
		return nil
	}
	out := new(JobQueue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JobQueue) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobQueueList) DeepCopyInto(out *JobQueueList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JobQueue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobQueueList.
func (in *JobQueueList) DeepCopy() *JobQueueList {
	if in == nil {
		return nil
	}
	out := new(JobQueueList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JobQueueList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobQueueObservation) DeepCopyInto(out *JobQueueObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobQueueObservation.
func (in *JobQueueObservation) DeepCopy() *JobQueueObservation {
	if in == nil {
		return nil
	}
	out := new(JobQueueObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobQueueParameters) DeepCopyInto(out *JobQueueParameters) {
	*out = *in
	if in.ComputeEnvironments != nil {
		in, out := &in.ComputeEnvironments, &out.ComputeEnvironments
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(int64)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobQueueParameters.
func (in *JobQueueParameters) DeepCopy() *JobQueueParameters {
	if in == nil {
		return nil
	}
	out := new(JobQueueParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobQueueSpec) DeepCopyInto(out *JobQueueSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobQueueSpec.
func (in *JobQueueSpec) DeepCopy() *JobQueueSpec {
	if in == nil {
		return nil
	}
	out := new(JobQueueSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobQueueStatus) DeepCopyInto(out *JobQueueStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobQueueStatus.
func (in *JobQueueStatus) DeepCopy() *JobQueueStatus {
	if in == nil {
		return nil
	}
	out := new(JobQueueStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchTemplateObservation) DeepCopyInto(out *LaunchTemplateObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchTemplateObservation.
func (in *LaunchTemplateObservation) DeepCopy() *LaunchTemplateObservation {
	if in == nil {
		return nil
	}
	out := new(LaunchTemplateObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchTemplateParameters) DeepCopyInto(out *LaunchTemplateParameters) {
	*out = *in
	if in.LaunchTemplateID != nil {
		in, out := &in.LaunchTemplateID, &out.LaunchTemplateID
		*out = new(string)
		**out = **in
	}
	if in.LaunchTemplateName != nil {
		in, out := &in.LaunchTemplateName, &out.LaunchTemplateName
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchTemplateParameters.
func (in *LaunchTemplateParameters) DeepCopy() *LaunchTemplateParameters {
	if in == nil {
		return nil
	}
	out := new(LaunchTemplateParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetryStrategyObservation) DeepCopyInto(out *RetryStrategyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetryStrategyObservation.
func (in *RetryStrategyObservation) DeepCopy() *RetryStrategyObservation {
	if in == nil {
		return nil
	}
	out := new(RetryStrategyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetryStrategyParameters) DeepCopyInto(out *RetryStrategyParameters) {
	*out = *in
	if in.Attempts != nil {
		in, out := &in.Attempts, &out.Attempts
		*out = new(int64)
		**out = **in
	}
	if in.EvaluateOnExit != nil {
		in, out := &in.EvaluateOnExit, &out.EvaluateOnExit
		*out = make([]EvaluateOnExitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetryStrategyParameters.
func (in *RetryStrategyParameters) DeepCopy() *RetryStrategyParameters {
	if in == nil {
		return nil
	}
	out := new(RetryStrategyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeoutObservation) DeepCopyInto(out *TimeoutObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeoutObservation.
func (in *TimeoutObservation) DeepCopy() *TimeoutObservation {
	if in == nil {
		return nil
	}
	out := new(TimeoutObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeoutParameters) DeepCopyInto(out *TimeoutParameters) {
	*out = *in
	if in.AttemptDurationSeconds != nil {
		in, out := &in.AttemptDurationSeconds, &out.AttemptDurationSeconds
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeoutParameters.
func (in *TimeoutParameters) DeepCopy() *TimeoutParameters {
	if in == nil {
		return nil
	}
	out := new(TimeoutParameters)
	in.DeepCopyInto(out)
	return out
}
