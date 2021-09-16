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
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type LbCookieStickinessPolicyObservation struct {
}

type LbCookieStickinessPolicyParameters struct {

	// +kubebuilder:validation:Optional
	CookieExpirationPeriod *int64 `json:"cookieExpirationPeriod,omitempty" tf:"cookie_expiration_period"`

	// +kubebuilder:validation:Required
	LbPort int64 `json:"lbPort" tf:"lb_port"`

	// +kubebuilder:validation:Required
	LoadBalancer string `json:"loadBalancer" tf:"load_balancer"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`
}

// LbCookieStickinessPolicySpec defines the desired state of LbCookieStickinessPolicy
type LbCookieStickinessPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       LbCookieStickinessPolicyParameters `json:"forProvider"`
}

// LbCookieStickinessPolicyStatus defines the observed state of LbCookieStickinessPolicy.
type LbCookieStickinessPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          LbCookieStickinessPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LbCookieStickinessPolicy is the Schema for the LbCookieStickinessPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type LbCookieStickinessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbCookieStickinessPolicySpec   `json:"spec"`
	Status            LbCookieStickinessPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LbCookieStickinessPolicyList contains a list of LbCookieStickinessPolicys
type LbCookieStickinessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LbCookieStickinessPolicy `json:"items"`
}

// Repository type metadata.
var (
	LbCookieStickinessPolicyKind             = "LbCookieStickinessPolicy"
	LbCookieStickinessPolicyGroupKind        = schema.GroupKind{Group: Group, Kind: LbCookieStickinessPolicyKind}.String()
	LbCookieStickinessPolicyKindAPIVersion   = LbCookieStickinessPolicyKind + "." + GroupVersion.String()
	LbCookieStickinessPolicyGroupVersionKind = GroupVersion.WithKind(LbCookieStickinessPolicyKind)
)

func init() {
	SchemeBuilder.Register(&LbCookieStickinessPolicy{}, &LbCookieStickinessPolicyList{})
}