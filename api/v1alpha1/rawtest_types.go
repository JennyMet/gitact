/*
Copyright 2021.

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

package v1alpha1

import (
	"github.com/gardener/gardener/pkg/apis/core/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type system string

const (
	sys system = "VHS"
)

type System struct {
	Type     system           `json:"type,omitempty"`
	Provider v1beta1.Provider `json:"provider,omitempty"`
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RawtestSpec defines the desired state of Rawtest
type RawtestSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	System *System `json:"system,omitempty"`
}

// RawtestStatus defines the observed state of Rawtest
type RawtestStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Rawtest is the Schema for the rawtests API
type Rawtest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RawtestSpec   `json:"spec,omitempty"`
	Status RawtestStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RawtestList contains a list of Rawtest
type RawtestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Rawtest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Rawtest{}, &RawtestList{})
}
