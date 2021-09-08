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

type Route53HealthCheckObservation struct {
}

type Route53HealthCheckParameters struct {
	ChildHealthThreshold *int64 `json:"childHealthThreshold,omitempty" tf:"child_health_threshold"`

	ChildHealthchecks []string `json:"childHealthchecks,omitempty" tf:"child_healthchecks"`

	CloudwatchAlarmName *string `json:"cloudwatchAlarmName,omitempty" tf:"cloudwatch_alarm_name"`

	CloudwatchAlarmRegion *string `json:"cloudwatchAlarmRegion,omitempty" tf:"cloudwatch_alarm_region"`

	Disabled *bool `json:"disabled,omitempty" tf:"disabled"`

	EnableSni *bool `json:"enableSni,omitempty" tf:"enable_sni"`

	FailureThreshold *int64 `json:"failureThreshold,omitempty" tf:"failure_threshold"`

	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn"`

	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address"`

	InsufficientDataHealthStatus *string `json:"insufficientDataHealthStatus,omitempty" tf:"insufficient_data_health_status"`

	InvertHealthcheck *bool `json:"invertHealthcheck,omitempty" tf:"invert_healthcheck"`

	MeasureLatency *bool `json:"measureLatency,omitempty" tf:"measure_latency"`

	Port *int64 `json:"port,omitempty" tf:"port"`

	ReferenceName *string `json:"referenceName,omitempty" tf:"reference_name"`

	Regions []string `json:"regions,omitempty" tf:"regions"`

	RequestInterval *int64 `json:"requestInterval,omitempty" tf:"request_interval"`

	ResourcePath *string `json:"resourcePath,omitempty" tf:"resource_path"`

	SearchString *string `json:"searchString,omitempty" tf:"search_string"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	Type string `json:"type" tf:"type"`
}

// Route53HealthCheckSpec defines the desired state of Route53HealthCheck
type Route53HealthCheckSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Route53HealthCheckParameters `json:"forProvider"`
}

// Route53HealthCheckStatus defines the observed state of Route53HealthCheck.
type Route53HealthCheckStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Route53HealthCheckObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Route53HealthCheck is the Schema for the Route53HealthChecks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Route53HealthCheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Route53HealthCheckSpec   `json:"spec"`
	Status            Route53HealthCheckStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Route53HealthCheckList contains a list of Route53HealthChecks
type Route53HealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route53HealthCheck `json:"items"`
}

// Repository type metadata.
var (
	Route53HealthCheckKind             = "Route53HealthCheck"
	Route53HealthCheckGroupKind        = schema.GroupKind{Group: Group, Kind: Route53HealthCheckKind}.String()
	Route53HealthCheckKindAPIVersion   = Route53HealthCheckKind + "." + GroupVersion.String()
	Route53HealthCheckGroupVersionKind = GroupVersion.WithKind(Route53HealthCheckKind)
)

func init() {
	SchemeBuilder.Register(&Route53HealthCheck{}, &Route53HealthCheckList{})
}
