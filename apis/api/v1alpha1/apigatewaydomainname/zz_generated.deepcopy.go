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
func (in *ApiGatewayDomainName) DeepCopyInto(out *ApiGatewayDomainName) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiGatewayDomainName.
func (in *ApiGatewayDomainName) DeepCopy() *ApiGatewayDomainName {
	if in == nil {
		return nil
	}
	out := new(ApiGatewayDomainName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApiGatewayDomainName) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiGatewayDomainNameList) DeepCopyInto(out *ApiGatewayDomainNameList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApiGatewayDomainName, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiGatewayDomainNameList.
func (in *ApiGatewayDomainNameList) DeepCopy() *ApiGatewayDomainNameList {
	if in == nil {
		return nil
	}
	out := new(ApiGatewayDomainNameList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApiGatewayDomainNameList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiGatewayDomainNameObservation) DeepCopyInto(out *ApiGatewayDomainNameObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiGatewayDomainNameObservation.
func (in *ApiGatewayDomainNameObservation) DeepCopy() *ApiGatewayDomainNameObservation {
	if in == nil {
		return nil
	}
	out := new(ApiGatewayDomainNameObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiGatewayDomainNameParameters) DeepCopyInto(out *ApiGatewayDomainNameParameters) {
	*out = *in
	if in.CertificateArn != nil {
		in, out := &in.CertificateArn, &out.CertificateArn
		*out = new(string)
		**out = **in
	}
	if in.CertificateBody != nil {
		in, out := &in.CertificateBody, &out.CertificateBody
		*out = new(string)
		**out = **in
	}
	if in.CertificateChain != nil {
		in, out := &in.CertificateChain, &out.CertificateChain
		*out = new(string)
		**out = **in
	}
	if in.CertificateName != nil {
		in, out := &in.CertificateName, &out.CertificateName
		*out = new(string)
		**out = **in
	}
	if in.CertificatePrivateKey != nil {
		in, out := &in.CertificatePrivateKey, &out.CertificatePrivateKey
		*out = new(string)
		**out = **in
	}
	if in.EndpointConfiguration != nil {
		in, out := &in.EndpointConfiguration, &out.EndpointConfiguration
		*out = make([]EndpointConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MutualTlsAuthentication != nil {
		in, out := &in.MutualTlsAuthentication, &out.MutualTlsAuthentication
		*out = make([]MutualTlsAuthenticationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RegionalCertificateArn != nil {
		in, out := &in.RegionalCertificateArn, &out.RegionalCertificateArn
		*out = new(string)
		**out = **in
	}
	if in.RegionalCertificateName != nil {
		in, out := &in.RegionalCertificateName, &out.RegionalCertificateName
		*out = new(string)
		**out = **in
	}
	if in.SecurityPolicy != nil {
		in, out := &in.SecurityPolicy, &out.SecurityPolicy
		*out = new(string)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiGatewayDomainNameParameters.
func (in *ApiGatewayDomainNameParameters) DeepCopy() *ApiGatewayDomainNameParameters {
	if in == nil {
		return nil
	}
	out := new(ApiGatewayDomainNameParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiGatewayDomainNameSpec) DeepCopyInto(out *ApiGatewayDomainNameSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiGatewayDomainNameSpec.
func (in *ApiGatewayDomainNameSpec) DeepCopy() *ApiGatewayDomainNameSpec {
	if in == nil {
		return nil
	}
	out := new(ApiGatewayDomainNameSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiGatewayDomainNameStatus) DeepCopyInto(out *ApiGatewayDomainNameStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiGatewayDomainNameStatus.
func (in *ApiGatewayDomainNameStatus) DeepCopy() *ApiGatewayDomainNameStatus {
	if in == nil {
		return nil
	}
	out := new(ApiGatewayDomainNameStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointConfigurationObservation) DeepCopyInto(out *EndpointConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointConfigurationObservation.
func (in *EndpointConfigurationObservation) DeepCopy() *EndpointConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(EndpointConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointConfigurationParameters) DeepCopyInto(out *EndpointConfigurationParameters) {
	*out = *in
	if in.Types != nil {
		in, out := &in.Types, &out.Types
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointConfigurationParameters.
func (in *EndpointConfigurationParameters) DeepCopy() *EndpointConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(EndpointConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MutualTlsAuthenticationObservation) DeepCopyInto(out *MutualTlsAuthenticationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MutualTlsAuthenticationObservation.
func (in *MutualTlsAuthenticationObservation) DeepCopy() *MutualTlsAuthenticationObservation {
	if in == nil {
		return nil
	}
	out := new(MutualTlsAuthenticationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MutualTlsAuthenticationParameters) DeepCopyInto(out *MutualTlsAuthenticationParameters) {
	*out = *in
	if in.TruststoreVersion != nil {
		in, out := &in.TruststoreVersion, &out.TruststoreVersion
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MutualTlsAuthenticationParameters.
func (in *MutualTlsAuthenticationParameters) DeepCopy() *MutualTlsAuthenticationParameters {
	if in == nil {
		return nil
	}
	out := new(MutualTlsAuthenticationParameters)
	in.DeepCopyInto(out)
	return out
}