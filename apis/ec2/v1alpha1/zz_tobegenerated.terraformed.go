package v1alpha1

import (
	jsoniter "github.com/json-iterator/go"
)

// GetObservation of this VPC
func (mg *VPC) GetObservation() ([]byte, error) {
	return jsoniter.Marshal(mg.Status.AtProvider)
}

// SetObservation for this VPC
func (mg *VPC) SetObservation(data []byte) error {
	return jsoniter.Unmarshal(data, mg.Status.AtProvider)
}

// GetParameters of this VPC
func (mg *VPC) GetParameters() ([]byte, error) {
	return jsoniter.Marshal(mg.Spec.ForProvider)
}

// SetParameters for this VPC
func (mg *VPC) SetParameters(data []byte) error {
	return jsoniter.Unmarshal(data, mg.Spec.ForProvider)
}
