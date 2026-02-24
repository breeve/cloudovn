/*
Copyright 2026.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// UnderlayNetworkSpec defines the desired state of UnderlayNetwork
type UnderlayNetworkSpec struct {
	// DefaultInterface
	DefaultInterface string `json:"defaultInterface,omitempty"`
}

// UnderlayNetworkStatus defines the observed state of UnderlayNetwork.
type UnderlayNetworkStatus struct {
	ResourceStatus `json:",inline"`
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:printcolumn:name="State",type=string,JSONPath=`.status.state`

// UnderlayNetwork is the Schema for the underlaynetworks API
type UnderlayNetwork struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`

	// spec defines the desired state of UnderlayNetwork
	// +required
	Spec UnderlayNetworkSpec `json:"spec"`

	// status defines the observed state of UnderlayNetwork
	// +optional
	Status UnderlayNetworkStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// UnderlayNetworkList contains a list of UnderlayNetwork
type UnderlayNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []UnderlayNetwork `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UnderlayNetwork{}, &UnderlayNetworkList{})
}
