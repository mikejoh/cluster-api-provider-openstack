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

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	internalinterfaces "sigs.k8s.io/cluster-api-provider-openstack/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// OpenStackClusters returns a OpenStackClusterInformer.
	OpenStackClusters() OpenStackClusterInformer
	// OpenStackClusterTemplates returns a OpenStackClusterTemplateInformer.
	OpenStackClusterTemplates() OpenStackClusterTemplateInformer
	// OpenStackMachines returns a OpenStackMachineInformer.
	OpenStackMachines() OpenStackMachineInformer
	// OpenStackMachineTemplates returns a OpenStackMachineTemplateInformer.
	OpenStackMachineTemplates() OpenStackMachineTemplateInformer
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

// OpenStackClusters returns a OpenStackClusterInformer.
func (v *version) OpenStackClusters() OpenStackClusterInformer {
	return &openStackClusterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// OpenStackClusterTemplates returns a OpenStackClusterTemplateInformer.
func (v *version) OpenStackClusterTemplates() OpenStackClusterTemplateInformer {
	return &openStackClusterTemplateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// OpenStackMachines returns a OpenStackMachineInformer.
func (v *version) OpenStackMachines() OpenStackMachineInformer {
	return &openStackMachineInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// OpenStackMachineTemplates returns a OpenStackMachineTemplateInformer.
func (v *version) OpenStackMachineTemplates() OpenStackMachineTemplateInformer {
	return &openStackMachineTemplateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
