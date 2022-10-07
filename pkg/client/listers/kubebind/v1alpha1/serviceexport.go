/*
Copyright The Kubectl Bind contributors.

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
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	v1alpha1 "github.com/kube-bind/kube-bind/pkg/apis/kubebind/v1alpha1"
)

// ServiceExportLister helps list ServiceExports.
// All objects returned here must be treated as read-only.
type ServiceExportLister interface {
	// List lists all ServiceExports in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceExport, err error)
	// ServiceExports returns an object that can list and get ServiceExports.
	ServiceExports(namespace string) ServiceExportNamespaceLister
	ServiceExportListerExpansion
}

// serviceExportLister implements the ServiceExportLister interface.
type serviceExportLister struct {
	indexer cache.Indexer
}

// NewServiceExportLister returns a new ServiceExportLister.
func NewServiceExportLister(indexer cache.Indexer) ServiceExportLister {
	return &serviceExportLister{indexer: indexer}
}

// List lists all ServiceExports in the indexer.
func (s *serviceExportLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceExport, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceExport))
	})
	return ret, err
}

// ServiceExports returns an object that can list and get ServiceExports.
func (s *serviceExportLister) ServiceExports(namespace string) ServiceExportNamespaceLister {
	return serviceExportNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServiceExportNamespaceLister helps list and get ServiceExports.
// All objects returned here must be treated as read-only.
type ServiceExportNamespaceLister interface {
	// List lists all ServiceExports in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceExport, err error)
	// Get retrieves the ServiceExport from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ServiceExport, error)
	ServiceExportNamespaceListerExpansion
}

// serviceExportNamespaceLister implements the ServiceExportNamespaceLister
// interface.
type serviceExportNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServiceExports in the indexer for a given namespace.
func (s serviceExportNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceExport, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceExport))
	})
	return ret, err
}

// Get retrieves the ServiceExport from the indexer for a given namespace and name.
func (s serviceExportNamespaceLister) Get(name string) (*v1alpha1.ServiceExport, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("serviceexport"), name)
	}
	return obj.(*v1alpha1.ServiceExport), nil
}
