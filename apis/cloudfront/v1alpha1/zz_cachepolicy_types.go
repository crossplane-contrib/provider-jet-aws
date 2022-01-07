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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CachePolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CachePolicyParameters struct {

	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultTTL *int64 `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// +kubebuilder:validation:Optional
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// +kubebuilder:validation:Optional
	MaxTTL *int64 `json:"maxTtl,omitempty" tf:"max_ttl,omitempty"`

	// +kubebuilder:validation:Optional
	MinTTL *int64 `json:"minTtl,omitempty" tf:"min_ttl,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	ParametersInCacheKeyAndForwardedToOrigin []ParametersInCacheKeyAndForwardedToOriginParameters `json:"parametersInCacheKeyAndForwardedToOrigin,omitempty" tf:"parameters_in_cache_key_and_forwarded_to_origin,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type CookiesConfigObservation struct {
}

type CookiesConfigParameters struct {

	// +kubebuilder:validation:Required
	CookieBehavior *string `json:"cookieBehavior" tf:"cookie_behavior,omitempty"`

	// +kubebuilder:validation:Optional
	Cookies []CookiesParameters `json:"cookies,omitempty" tf:"cookies,omitempty"`
}

type CookiesObservation struct {
}

type CookiesParameters struct {

	// +kubebuilder:validation:Optional
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type HeadersConfigObservation struct {
}

type HeadersConfigParameters struct {

	// +kubebuilder:validation:Optional
	HeaderBehavior *string `json:"headerBehavior,omitempty" tf:"header_behavior,omitempty"`

	// +kubebuilder:validation:Optional
	Headers []HeadersParameters `json:"headers,omitempty" tf:"headers,omitempty"`
}

type HeadersObservation struct {
}

type HeadersParameters struct {

	// +kubebuilder:validation:Optional
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type ParametersInCacheKeyAndForwardedToOriginObservation struct {
}

type ParametersInCacheKeyAndForwardedToOriginParameters struct {

	// +kubebuilder:validation:Required
	CookiesConfig []CookiesConfigParameters `json:"cookiesConfig" tf:"cookies_config,omitempty"`

	// +kubebuilder:validation:Optional
	EnableAcceptEncodingBrotli *bool `json:"enableAcceptEncodingBrotli,omitempty" tf:"enable_accept_encoding_brotli,omitempty"`

	// +kubebuilder:validation:Optional
	EnableAcceptEncodingGzip *bool `json:"enableAcceptEncodingGzip,omitempty" tf:"enable_accept_encoding_gzip,omitempty"`

	// +kubebuilder:validation:Required
	HeadersConfig []HeadersConfigParameters `json:"headersConfig" tf:"headers_config,omitempty"`

	// +kubebuilder:validation:Required
	QueryStringsConfig []QueryStringsConfigParameters `json:"queryStringsConfig" tf:"query_strings_config,omitempty"`
}

type QueryStringsConfigObservation struct {
}

type QueryStringsConfigParameters struct {

	// +kubebuilder:validation:Required
	QueryStringBehavior *string `json:"queryStringBehavior" tf:"query_string_behavior,omitempty"`

	// +kubebuilder:validation:Optional
	QueryStrings []QueryStringsParameters `json:"queryStrings,omitempty" tf:"query_strings,omitempty"`
}

type QueryStringsObservation struct {
}

type QueryStringsParameters struct {

	// +kubebuilder:validation:Optional
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

// CachePolicySpec defines the desired state of CachePolicy
type CachePolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CachePolicyParameters `json:"forProvider"`
}

// CachePolicyStatus defines the observed state of CachePolicy.
type CachePolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CachePolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CachePolicy is the Schema for the CachePolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type CachePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CachePolicySpec   `json:"spec"`
	Status            CachePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CachePolicyList contains a list of CachePolicys
type CachePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CachePolicy `json:"items"`
}

// Repository type metadata.
var (
	CachePolicy_Kind             = "CachePolicy"
	CachePolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CachePolicy_Kind}.String()
	CachePolicy_KindAPIVersion   = CachePolicy_Kind + "." + CRDGroupVersion.String()
	CachePolicy_GroupVersionKind = CRDGroupVersion.WithKind(CachePolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&CachePolicy{}, &CachePolicyList{})
}
