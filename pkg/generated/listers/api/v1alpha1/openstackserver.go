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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "sigs.k8s.io/cluster-api-provider-openstack/api/v1alpha1"
)

// OpenStackServerLister helps list OpenStackServers.
// All objects returned here must be treated as read-only.
type OpenStackServerLister interface {
	// List lists all OpenStackServers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.OpenStackServer, err error)
	// OpenStackServers returns an object that can list and get OpenStackServers.
	OpenStackServers(namespace string) OpenStackServerNamespaceLister
	OpenStackServerListerExpansion
}

// openStackServerLister implements the OpenStackServerLister interface.
type openStackServerLister struct {
	listers.ResourceIndexer[*v1alpha1.OpenStackServer]
}

// NewOpenStackServerLister returns a new OpenStackServerLister.
func NewOpenStackServerLister(indexer cache.Indexer) OpenStackServerLister {
	return &openStackServerLister{listers.New[*v1alpha1.OpenStackServer](indexer, v1alpha1.Resource("openstackserver"))}
}

// OpenStackServers returns an object that can list and get OpenStackServers.
func (s *openStackServerLister) OpenStackServers(namespace string) OpenStackServerNamespaceLister {
	return openStackServerNamespaceLister{listers.NewNamespaced[*v1alpha1.OpenStackServer](s.ResourceIndexer, namespace)}
}

// OpenStackServerNamespaceLister helps list and get OpenStackServers.
// All objects returned here must be treated as read-only.
type OpenStackServerNamespaceLister interface {
	// List lists all OpenStackServers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.OpenStackServer, err error)
	// Get retrieves the OpenStackServer from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.OpenStackServer, error)
	OpenStackServerNamespaceListerExpansion
}

// openStackServerNamespaceLister implements the OpenStackServerNamespaceLister
// interface.
type openStackServerNamespaceLister struct {
	listers.ResourceIndexer[*v1alpha1.OpenStackServer]
}
