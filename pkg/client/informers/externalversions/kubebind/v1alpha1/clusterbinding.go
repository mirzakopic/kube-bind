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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"

	kubebindv1alpha1 "github.com/kube-bind/kube-bind/pkg/apis/kubebind/v1alpha1"
	versioned "github.com/kube-bind/kube-bind/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kube-bind/kube-bind/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kube-bind/kube-bind/pkg/client/listers/kubebind/v1alpha1"
)

// ClusterBindingInformer provides access to a shared informer and lister for
// ClusterBindings.
type ClusterBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ClusterBindingLister
}

type clusterBindingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewClusterBindingInformer constructs a new informer for ClusterBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterBindingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredClusterBindingInformer constructs a new informer for ClusterBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeBindV1alpha1().ClusterBindings(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeBindV1alpha1().ClusterBindings(namespace).Watch(context.TODO(), options)
			},
		},
		&kubebindv1alpha1.ClusterBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterBindingInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterBindingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubebindv1alpha1.ClusterBinding{}, f.defaultInformer)
}

func (f *clusterBindingInformer) Lister() v1alpha1.ClusterBindingLister {
	return v1alpha1.NewClusterBindingLister(f.Informer().GetIndexer())
}
