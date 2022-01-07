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
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *App) DeepCopyInto(out *App) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new App.
func (in *App) DeepCopy() *App {
	if in == nil {
		return nil
	}
	out := new(App)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *App) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppList) DeepCopyInto(out *AppList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]App, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppList.
func (in *AppList) DeepCopy() *AppList {
	if in == nil {
		return nil
	}
	out := new(AppList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppObservation) DeepCopyInto(out *AppObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.DefaultDomain != nil {
		in, out := &in.DefaultDomain, &out.DefaultDomain
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ProductionBranch != nil {
		in, out := &in.ProductionBranch, &out.ProductionBranch
		*out = make([]ProductionBranchObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppObservation.
func (in *AppObservation) DeepCopy() *AppObservation {
	if in == nil {
		return nil
	}
	out := new(AppObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppParameters) DeepCopyInto(out *AppParameters) {
	*out = *in
	if in.AccessTokenSecretRef != nil {
		in, out := &in.AccessTokenSecretRef, &out.AccessTokenSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AutoBranchCreationConfig != nil {
		in, out := &in.AutoBranchCreationConfig, &out.AutoBranchCreationConfig
		*out = make([]AutoBranchCreationConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AutoBranchCreationPatterns != nil {
		in, out := &in.AutoBranchCreationPatterns, &out.AutoBranchCreationPatterns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.BasicAuthCredentialsSecretRef != nil {
		in, out := &in.BasicAuthCredentialsSecretRef, &out.BasicAuthCredentialsSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.BuildSpec != nil {
		in, out := &in.BuildSpec, &out.BuildSpec
		*out = new(string)
		**out = **in
	}
	if in.CustomRule != nil {
		in, out := &in.CustomRule, &out.CustomRule
		*out = make([]CustomRuleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EnableAutoBranchCreation != nil {
		in, out := &in.EnableAutoBranchCreation, &out.EnableAutoBranchCreation
		*out = new(bool)
		**out = **in
	}
	if in.EnableBasicAuth != nil {
		in, out := &in.EnableBasicAuth, &out.EnableBasicAuth
		*out = new(bool)
		**out = **in
	}
	if in.EnableBranchAutoBuild != nil {
		in, out := &in.EnableBranchAutoBuild, &out.EnableBranchAutoBuild
		*out = new(bool)
		**out = **in
	}
	if in.EnableBranchAutoDeletion != nil {
		in, out := &in.EnableBranchAutoDeletion, &out.EnableBranchAutoDeletion
		*out = new(bool)
		**out = **in
	}
	if in.EnvironmentVariables != nil {
		in, out := &in.EnvironmentVariables, &out.EnvironmentVariables
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
	if in.IAMServiceRoleArn != nil {
		in, out := &in.IAMServiceRoleArn, &out.IAMServiceRoleArn
		*out = new(string)
		**out = **in
	}
	if in.IAMServiceRoleArnRef != nil {
		in, out := &in.IAMServiceRoleArnRef, &out.IAMServiceRoleArnRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.IAMServiceRoleArnSelector != nil {
		in, out := &in.IAMServiceRoleArnSelector, &out.IAMServiceRoleArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OauthTokenSecretRef != nil {
		in, out := &in.OauthTokenSecretRef, &out.OauthTokenSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.Platform != nil {
		in, out := &in.Platform, &out.Platform
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Repository != nil {
		in, out := &in.Repository, &out.Repository
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppParameters.
func (in *AppParameters) DeepCopy() *AppParameters {
	if in == nil {
		return nil
	}
	out := new(AppParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSpec) DeepCopyInto(out *AppSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSpec.
func (in *AppSpec) DeepCopy() *AppSpec {
	if in == nil {
		return nil
	}
	out := new(AppSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppStatus) DeepCopyInto(out *AppStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppStatus.
func (in *AppStatus) DeepCopy() *AppStatus {
	if in == nil {
		return nil
	}
	out := new(AppStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoBranchCreationConfigObservation) DeepCopyInto(out *AutoBranchCreationConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoBranchCreationConfigObservation.
func (in *AutoBranchCreationConfigObservation) DeepCopy() *AutoBranchCreationConfigObservation {
	if in == nil {
		return nil
	}
	out := new(AutoBranchCreationConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoBranchCreationConfigParameters) DeepCopyInto(out *AutoBranchCreationConfigParameters) {
	*out = *in
	if in.BasicAuthCredentialsSecretRef != nil {
		in, out := &in.BasicAuthCredentialsSecretRef, &out.BasicAuthCredentialsSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.BuildSpec != nil {
		in, out := &in.BuildSpec, &out.BuildSpec
		*out = new(string)
		**out = **in
	}
	if in.EnableAutoBuild != nil {
		in, out := &in.EnableAutoBuild, &out.EnableAutoBuild
		*out = new(bool)
		**out = **in
	}
	if in.EnableBasicAuth != nil {
		in, out := &in.EnableBasicAuth, &out.EnableBasicAuth
		*out = new(bool)
		**out = **in
	}
	if in.EnablePerformanceMode != nil {
		in, out := &in.EnablePerformanceMode, &out.EnablePerformanceMode
		*out = new(bool)
		**out = **in
	}
	if in.EnablePullRequestPreview != nil {
		in, out := &in.EnablePullRequestPreview, &out.EnablePullRequestPreview
		*out = new(bool)
		**out = **in
	}
	if in.EnvironmentVariables != nil {
		in, out := &in.EnvironmentVariables, &out.EnvironmentVariables
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
	if in.Framework != nil {
		in, out := &in.Framework, &out.Framework
		*out = new(string)
		**out = **in
	}
	if in.PullRequestEnvironmentName != nil {
		in, out := &in.PullRequestEnvironmentName, &out.PullRequestEnvironmentName
		*out = new(string)
		**out = **in
	}
	if in.Stage != nil {
		in, out := &in.Stage, &out.Stage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoBranchCreationConfigParameters.
func (in *AutoBranchCreationConfigParameters) DeepCopy() *AutoBranchCreationConfigParameters {
	if in == nil {
		return nil
	}
	out := new(AutoBranchCreationConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendEnvironment) DeepCopyInto(out *BackendEnvironment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendEnvironment.
func (in *BackendEnvironment) DeepCopy() *BackendEnvironment {
	if in == nil {
		return nil
	}
	out := new(BackendEnvironment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackendEnvironment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendEnvironmentList) DeepCopyInto(out *BackendEnvironmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackendEnvironment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendEnvironmentList.
func (in *BackendEnvironmentList) DeepCopy() *BackendEnvironmentList {
	if in == nil {
		return nil
	}
	out := new(BackendEnvironmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackendEnvironmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendEnvironmentObservation) DeepCopyInto(out *BackendEnvironmentObservation) {
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendEnvironmentObservation.
func (in *BackendEnvironmentObservation) DeepCopy() *BackendEnvironmentObservation {
	if in == nil {
		return nil
	}
	out := new(BackendEnvironmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendEnvironmentParameters) DeepCopyInto(out *BackendEnvironmentParameters) {
	*out = *in
	if in.AppID != nil {
		in, out := &in.AppID, &out.AppID
		*out = new(string)
		**out = **in
	}
	if in.DeploymentArtifacts != nil {
		in, out := &in.DeploymentArtifacts, &out.DeploymentArtifacts
		*out = new(string)
		**out = **in
	}
	if in.EnvironmentName != nil {
		in, out := &in.EnvironmentName, &out.EnvironmentName
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.StackName != nil {
		in, out := &in.StackName, &out.StackName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendEnvironmentParameters.
func (in *BackendEnvironmentParameters) DeepCopy() *BackendEnvironmentParameters {
	if in == nil {
		return nil
	}
	out := new(BackendEnvironmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendEnvironmentSpec) DeepCopyInto(out *BackendEnvironmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendEnvironmentSpec.
func (in *BackendEnvironmentSpec) DeepCopy() *BackendEnvironmentSpec {
	if in == nil {
		return nil
	}
	out := new(BackendEnvironmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendEnvironmentStatus) DeepCopyInto(out *BackendEnvironmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendEnvironmentStatus.
func (in *BackendEnvironmentStatus) DeepCopy() *BackendEnvironmentStatus {
	if in == nil {
		return nil
	}
	out := new(BackendEnvironmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Branch) DeepCopyInto(out *Branch) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Branch.
func (in *Branch) DeepCopy() *Branch {
	if in == nil {
		return nil
	}
	out := new(Branch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Branch) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BranchList) DeepCopyInto(out *BranchList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Branch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BranchList.
func (in *BranchList) DeepCopy() *BranchList {
	if in == nil {
		return nil
	}
	out := new(BranchList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BranchList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BranchObservation) DeepCopyInto(out *BranchObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.AssociatedResources != nil {
		in, out := &in.AssociatedResources, &out.AssociatedResources
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.CustomDomains != nil {
		in, out := &in.CustomDomains, &out.CustomDomains
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.DestinationBranch != nil {
		in, out := &in.DestinationBranch, &out.DestinationBranch
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.SourceBranch != nil {
		in, out := &in.SourceBranch, &out.SourceBranch
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BranchObservation.
func (in *BranchObservation) DeepCopy() *BranchObservation {
	if in == nil {
		return nil
	}
	out := new(BranchObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BranchParameters) DeepCopyInto(out *BranchParameters) {
	*out = *in
	if in.AppID != nil {
		in, out := &in.AppID, &out.AppID
		*out = new(string)
		**out = **in
	}
	if in.BackendEnvironmentArn != nil {
		in, out := &in.BackendEnvironmentArn, &out.BackendEnvironmentArn
		*out = new(string)
		**out = **in
	}
	if in.BasicAuthCredentialsSecretRef != nil {
		in, out := &in.BasicAuthCredentialsSecretRef, &out.BasicAuthCredentialsSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.BranchName != nil {
		in, out := &in.BranchName, &out.BranchName
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.EnableAutoBuild != nil {
		in, out := &in.EnableAutoBuild, &out.EnableAutoBuild
		*out = new(bool)
		**out = **in
	}
	if in.EnableBasicAuth != nil {
		in, out := &in.EnableBasicAuth, &out.EnableBasicAuth
		*out = new(bool)
		**out = **in
	}
	if in.EnableNotification != nil {
		in, out := &in.EnableNotification, &out.EnableNotification
		*out = new(bool)
		**out = **in
	}
	if in.EnablePerformanceMode != nil {
		in, out := &in.EnablePerformanceMode, &out.EnablePerformanceMode
		*out = new(bool)
		**out = **in
	}
	if in.EnablePullRequestPreview != nil {
		in, out := &in.EnablePullRequestPreview, &out.EnablePullRequestPreview
		*out = new(bool)
		**out = **in
	}
	if in.EnvironmentVariables != nil {
		in, out := &in.EnvironmentVariables, &out.EnvironmentVariables
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
	if in.Framework != nil {
		in, out := &in.Framework, &out.Framework
		*out = new(string)
		**out = **in
	}
	if in.PullRequestEnvironmentName != nil {
		in, out := &in.PullRequestEnvironmentName, &out.PullRequestEnvironmentName
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Stage != nil {
		in, out := &in.Stage, &out.Stage
		*out = new(string)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BranchParameters.
func (in *BranchParameters) DeepCopy() *BranchParameters {
	if in == nil {
		return nil
	}
	out := new(BranchParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BranchSpec) DeepCopyInto(out *BranchSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BranchSpec.
func (in *BranchSpec) DeepCopy() *BranchSpec {
	if in == nil {
		return nil
	}
	out := new(BranchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BranchStatus) DeepCopyInto(out *BranchStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BranchStatus.
func (in *BranchStatus) DeepCopy() *BranchStatus {
	if in == nil {
		return nil
	}
	out := new(BranchStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomRuleObservation) DeepCopyInto(out *CustomRuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomRuleObservation.
func (in *CustomRuleObservation) DeepCopy() *CustomRuleObservation {
	if in == nil {
		return nil
	}
	out := new(CustomRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomRuleParameters) DeepCopyInto(out *CustomRuleParameters) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(string)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomRuleParameters.
func (in *CustomRuleParameters) DeepCopy() *CustomRuleParameters {
	if in == nil {
		return nil
	}
	out := new(CustomRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainAssociation) DeepCopyInto(out *DomainAssociation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainAssociation.
func (in *DomainAssociation) DeepCopy() *DomainAssociation {
	if in == nil {
		return nil
	}
	out := new(DomainAssociation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainAssociation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainAssociationList) DeepCopyInto(out *DomainAssociationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DomainAssociation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainAssociationList.
func (in *DomainAssociationList) DeepCopy() *DomainAssociationList {
	if in == nil {
		return nil
	}
	out := new(DomainAssociationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainAssociationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainAssociationObservation) DeepCopyInto(out *DomainAssociationObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.CertificateVerificationDNSRecord != nil {
		in, out := &in.CertificateVerificationDNSRecord, &out.CertificateVerificationDNSRecord
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainAssociationObservation.
func (in *DomainAssociationObservation) DeepCopy() *DomainAssociationObservation {
	if in == nil {
		return nil
	}
	out := new(DomainAssociationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainAssociationParameters) DeepCopyInto(out *DomainAssociationParameters) {
	*out = *in
	if in.AppID != nil {
		in, out := &in.AppID, &out.AppID
		*out = new(string)
		**out = **in
	}
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SubDomain != nil {
		in, out := &in.SubDomain, &out.SubDomain
		*out = make([]SubDomainParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WaitForVerification != nil {
		in, out := &in.WaitForVerification, &out.WaitForVerification
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainAssociationParameters.
func (in *DomainAssociationParameters) DeepCopy() *DomainAssociationParameters {
	if in == nil {
		return nil
	}
	out := new(DomainAssociationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainAssociationSpec) DeepCopyInto(out *DomainAssociationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainAssociationSpec.
func (in *DomainAssociationSpec) DeepCopy() *DomainAssociationSpec {
	if in == nil {
		return nil
	}
	out := new(DomainAssociationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainAssociationStatus) DeepCopyInto(out *DomainAssociationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainAssociationStatus.
func (in *DomainAssociationStatus) DeepCopy() *DomainAssociationStatus {
	if in == nil {
		return nil
	}
	out := new(DomainAssociationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProductionBranchObservation) DeepCopyInto(out *ProductionBranchObservation) {
	*out = *in
	if in.BranchName != nil {
		in, out := &in.BranchName, &out.BranchName
		*out = new(string)
		**out = **in
	}
	if in.LastDeployTime != nil {
		in, out := &in.LastDeployTime, &out.LastDeployTime
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.ThumbnailURL != nil {
		in, out := &in.ThumbnailURL, &out.ThumbnailURL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProductionBranchObservation.
func (in *ProductionBranchObservation) DeepCopy() *ProductionBranchObservation {
	if in == nil {
		return nil
	}
	out := new(ProductionBranchObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProductionBranchParameters) DeepCopyInto(out *ProductionBranchParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProductionBranchParameters.
func (in *ProductionBranchParameters) DeepCopy() *ProductionBranchParameters {
	if in == nil {
		return nil
	}
	out := new(ProductionBranchParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubDomainObservation) DeepCopyInto(out *SubDomainObservation) {
	*out = *in
	if in.DNSRecord != nil {
		in, out := &in.DNSRecord, &out.DNSRecord
		*out = new(string)
		**out = **in
	}
	if in.Verified != nil {
		in, out := &in.Verified, &out.Verified
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubDomainObservation.
func (in *SubDomainObservation) DeepCopy() *SubDomainObservation {
	if in == nil {
		return nil
	}
	out := new(SubDomainObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubDomainParameters) DeepCopyInto(out *SubDomainParameters) {
	*out = *in
	if in.BranchName != nil {
		in, out := &in.BranchName, &out.BranchName
		*out = new(string)
		**out = **in
	}
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubDomainParameters.
func (in *SubDomainParameters) DeepCopy() *SubDomainParameters {
	if in == nil {
		return nil
	}
	out := new(SubDomainParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Webhook) DeepCopyInto(out *Webhook) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Webhook.
func (in *Webhook) DeepCopy() *Webhook {
	if in == nil {
		return nil
	}
	out := new(Webhook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Webhook) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookList) DeepCopyInto(out *WebhookList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Webhook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookList.
func (in *WebhookList) DeepCopy() *WebhookList {
	if in == nil {
		return nil
	}
	out := new(WebhookList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WebhookList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookObservation) DeepCopyInto(out *WebhookObservation) {
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
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookObservation.
func (in *WebhookObservation) DeepCopy() *WebhookObservation {
	if in == nil {
		return nil
	}
	out := new(WebhookObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookParameters) DeepCopyInto(out *WebhookParameters) {
	*out = *in
	if in.AppID != nil {
		in, out := &in.AppID, &out.AppID
		*out = new(string)
		**out = **in
	}
	if in.BranchName != nil {
		in, out := &in.BranchName, &out.BranchName
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookParameters.
func (in *WebhookParameters) DeepCopy() *WebhookParameters {
	if in == nil {
		return nil
	}
	out := new(WebhookParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookSpec) DeepCopyInto(out *WebhookSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookSpec.
func (in *WebhookSpec) DeepCopy() *WebhookSpec {
	if in == nil {
		return nil
	}
	out := new(WebhookSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookStatus) DeepCopyInto(out *WebhookStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookStatus.
func (in *WebhookStatus) DeepCopy() *WebhookStatus {
	if in == nil {
		return nil
	}
	out := new(WebhookStatus)
	in.DeepCopyInto(out)
	return out
}
