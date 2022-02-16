/*
Copyright 2022 The OpenYurt Authors.

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

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=koledaemonsets,shortName=kd,categories=all
// +kubebuilder:subresource:status

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// KoleDaemonSet is
type KoleDaemonSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   *PodSpec             `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status *KoleDaemonSetStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type KoleDaemonSetStatus struct {
	CurrentNumberScheduled int `json:"currentNumberScheduled"`
	DesiredNumberScheduled int `json:"desiredNumberScheduled"`
	NumberReady            int `json:"numberReady"`
}

type PodSpec struct {
	Image        string            `json:"image,omitempty"`
	Command      []string          `json:"command,omitempty"`
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KoleDaemonSetList is
type KoleDaemonSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []KoleDaemonSet `json:"items"`
}