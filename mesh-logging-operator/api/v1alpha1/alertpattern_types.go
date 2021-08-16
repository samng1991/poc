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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AlertPatternItem struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	EventId string `json:"eventId,omitempty"`
	Regex   string `json:"regex,omitempty"`
}

// AlertPatternSpec defines the desired state of AlertPattern
type AlertPatternSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	AlertPatternItems []AlertPatternItem `json:"alertPatterns,omitempty"`
}

// AlertPatternStatus defines the observed state of AlertPattern
type AlertPatternStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// AlertPattern is the Schema for the alertpatterns API
type AlertPattern struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlertPatternSpec   `json:"spec,omitempty"`
	Status AlertPatternStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AlertPatternList contains a list of AlertPattern
type AlertPatternList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlertPattern `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AlertPattern{}, &AlertPatternList{})
}
