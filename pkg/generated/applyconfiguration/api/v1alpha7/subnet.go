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

package v1alpha7

// SubnetApplyConfiguration represents an declarative configuration of the Subnet type for use
// with apply.
type SubnetApplyConfiguration struct {
	Name *string  `json:"name,omitempty"`
	ID   *string  `json:"id,omitempty"`
	CIDR *string  `json:"cidr,omitempty"`
	Tags []string `json:"tags,omitempty"`
}

// SubnetApplyConfiguration constructs an declarative configuration of the Subnet type for use with
// apply.
func Subnet() *SubnetApplyConfiguration {
	return &SubnetApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *SubnetApplyConfiguration) WithName(value string) *SubnetApplyConfiguration {
	b.Name = &value
	return b
}

// WithID sets the ID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ID field is set to the value of the last call.
func (b *SubnetApplyConfiguration) WithID(value string) *SubnetApplyConfiguration {
	b.ID = &value
	return b
}

// WithCIDR sets the CIDR field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CIDR field is set to the value of the last call.
func (b *SubnetApplyConfiguration) WithCIDR(value string) *SubnetApplyConfiguration {
	b.CIDR = &value
	return b
}

// WithTags adds the given value to the Tags field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tags field.
func (b *SubnetApplyConfiguration) WithTags(values ...string) *SubnetApplyConfiguration {
	for i := range values {
		b.Tags = append(b.Tags, values[i])
	}
	return b
}
