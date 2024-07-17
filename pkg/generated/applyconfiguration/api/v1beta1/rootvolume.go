/*
Copyright 2024 The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

// RootVolumeApplyConfiguration represents an declarative configuration of the RootVolume type for use
// with apply.
type RootVolumeApplyConfiguration struct {
	SizeGiB                             *int `json:"sizeGiB,omitempty"`
	BlockDeviceVolumeApplyConfiguration `json:",inline"`
}

// RootVolumeApplyConfiguration constructs an declarative configuration of the RootVolume type for use with
// apply.
func RootVolume() *RootVolumeApplyConfiguration {
	return &RootVolumeApplyConfiguration{}
}

// WithSizeGiB sets the SizeGiB field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SizeGiB field is set to the value of the last call.
func (b *RootVolumeApplyConfiguration) WithSizeGiB(value int) *RootVolumeApplyConfiguration {
	b.SizeGiB = &value
	return b
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *RootVolumeApplyConfiguration) WithType(value string) *RootVolumeApplyConfiguration {
	b.Type = &value
	return b
}

// WithAvailabilityZone sets the AvailabilityZone field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AvailabilityZone field is set to the value of the last call.
func (b *RootVolumeApplyConfiguration) WithAvailabilityZone(value *VolumeAvailabilityZoneApplyConfiguration) *RootVolumeApplyConfiguration {
	b.AvailabilityZone = value
	return b
}
