/*
Copyright 2025 The etcd-operator Authors

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

package v1beta2

// OSSBackupSourceApplyConfiguration represents a declarative configuration of the OSSBackupSource type for use
// with apply.
type OSSBackupSourceApplyConfiguration struct {
	Path      *string `json:"path,omitempty"`
	OSSSecret *string `json:"ossSecret,omitempty"`
	Endpoint  *string `json:"endpoint,omitempty"`
}

// OSSBackupSourceApplyConfiguration constructs a declarative configuration of the OSSBackupSource type for use with
// apply.
func OSSBackupSource() *OSSBackupSourceApplyConfiguration {
	return &OSSBackupSourceApplyConfiguration{}
}

// WithPath sets the Path field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Path field is set to the value of the last call.
func (b *OSSBackupSourceApplyConfiguration) WithPath(value string) *OSSBackupSourceApplyConfiguration {
	b.Path = &value
	return b
}

// WithOSSSecret sets the OSSSecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OSSSecret field is set to the value of the last call.
func (b *OSSBackupSourceApplyConfiguration) WithOSSSecret(value string) *OSSBackupSourceApplyConfiguration {
	b.OSSSecret = &value
	return b
}

// WithEndpoint sets the Endpoint field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Endpoint field is set to the value of the last call.
func (b *OSSBackupSourceApplyConfiguration) WithEndpoint(value string) *OSSBackupSourceApplyConfiguration {
	b.Endpoint = &value
	return b
}
