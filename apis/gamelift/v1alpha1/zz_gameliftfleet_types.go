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

type Ec2InboundPermissionObservation struct {
}

type Ec2InboundPermissionParameters struct {
	FromPort int64 `json:"fromPort" tf:"from_port"`

	IPRange string `json:"ipRange" tf:"ip_range"`

	Protocol string `json:"protocol" tf:"protocol"`

	ToPort int64 `json:"toPort" tf:"to_port"`
}

type GameliftFleetObservation struct {
	ARN string `json:"arn" tf:"arn"`

	LogPaths []string `json:"logPaths" tf:"log_paths"`

	OperatingSystem string `json:"operatingSystem" tf:"operating_system"`
}

type GameliftFleetParameters struct {
	BuildID string `json:"buildID" tf:"build_id"`

	Description *string `json:"description,omitempty" tf:"description"`

	Ec2InboundPermission []Ec2InboundPermissionParameters `json:"ec2InboundPermission,omitempty" tf:"ec2_inbound_permission"`

	Ec2InstanceType string `json:"ec2InstanceType" tf:"ec2_instance_type"`

	FleetType *string `json:"fleetType,omitempty" tf:"fleet_type"`

	InstanceRoleARN *string `json:"instanceRoleARN,omitempty" tf:"instance_role_arn"`

	MetricGroups []string `json:"metricGroups,omitempty" tf:"metric_groups"`

	Name string `json:"name" tf:"name"`

	NewGameSessionProtectionPolicy *string `json:"newGameSessionProtectionPolicy,omitempty" tf:"new_game_session_protection_policy"`

	ResourceCreationLimitPolicy []ResourceCreationLimitPolicyParameters `json:"resourceCreationLimitPolicy,omitempty" tf:"resource_creation_limit_policy"`

	RuntimeConfiguration []RuntimeConfigurationParameters `json:"runtimeConfiguration,omitempty" tf:"runtime_configuration"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type ResourceCreationLimitPolicyObservation struct {
}

type ResourceCreationLimitPolicyParameters struct {
	NewGameSessionsPerCreator *int64 `json:"newGameSessionsPerCreator,omitempty" tf:"new_game_sessions_per_creator"`

	PolicyPeriodInMinutes *int64 `json:"policyPeriodInMinutes,omitempty" tf:"policy_period_in_minutes"`
}

type RuntimeConfigurationObservation struct {
}

type RuntimeConfigurationParameters struct {
	GameSessionActivationTimeoutSeconds *int64 `json:"gameSessionActivationTimeoutSeconds,omitempty" tf:"game_session_activation_timeout_seconds"`

	MaxConcurrentGameSessionActivations *int64 `json:"maxConcurrentGameSessionActivations,omitempty" tf:"max_concurrent_game_session_activations"`

	ServerProcess []ServerProcessParameters `json:"serverProcess,omitempty" tf:"server_process"`
}

type ServerProcessObservation struct {
}

type ServerProcessParameters struct {
	ConcurrentExecutions int64 `json:"concurrentExecutions" tf:"concurrent_executions"`

	LaunchPath string `json:"launchPath" tf:"launch_path"`

	Parameters *string `json:"parameters,omitempty" tf:"parameters"`
}

// GameliftFleetSpec defines the desired state of GameliftFleet
type GameliftFleetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       GameliftFleetParameters `json:"forProvider"`
}

// GameliftFleetStatus defines the observed state of GameliftFleet.
type GameliftFleetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          GameliftFleetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GameliftFleet is the Schema for the GameliftFleets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type GameliftFleet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GameliftFleetSpec   `json:"spec"`
	Status            GameliftFleetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GameliftFleetList contains a list of GameliftFleets
type GameliftFleetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GameliftFleet `json:"items"`
}

// Repository type metadata.
var (
	GameliftFleetKind             = "GameliftFleet"
	GameliftFleetGroupKind        = schema.GroupKind{Group: Group, Kind: GameliftFleetKind}.String()
	GameliftFleetKindAPIVersion   = GameliftFleetKind + "." + GroupVersion.String()
	GameliftFleetGroupVersionKind = GroupVersion.WithKind(GameliftFleetKind)
)

func init() {
	SchemeBuilder.Register(&GameliftFleet{}, &GameliftFleetList{})
}
