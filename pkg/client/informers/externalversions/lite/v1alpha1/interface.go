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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/openyurtio/kole/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// InfDaemonSets returns a InfDaemonSetInformer.
	InfDaemonSets() InfDaemonSetInformer
	// InfEdgeNodes returns a InfEdgeNodeInformer.
	InfEdgeNodes() InfEdgeNodeInformer
	// QueryNodes returns a QueryNodeInformer.
	QueryNodes() QueryNodeInformer
	// Summaries returns a SummaryInformer.
	Summaries() SummaryInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// InfDaemonSets returns a InfDaemonSetInformer.
func (v *version) InfDaemonSets() InfDaemonSetInformer {
	return &infDaemonSetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// InfEdgeNodes returns a InfEdgeNodeInformer.
func (v *version) InfEdgeNodes() InfEdgeNodeInformer {
	return &infEdgeNodeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// QueryNodes returns a QueryNodeInformer.
func (v *version) QueryNodes() QueryNodeInformer {
	return &queryNodeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Summaries returns a SummaryInformer.
func (v *version) Summaries() SummaryInformer {
	return &summaryInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
