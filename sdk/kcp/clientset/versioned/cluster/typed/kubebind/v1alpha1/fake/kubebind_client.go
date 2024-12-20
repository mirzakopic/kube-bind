//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kube Bind Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package fake

import (
	"github.com/kcp-dev/logicalcluster/v3"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	"k8s.io/client-go/rest"

	kcpkubebindv1alpha1 "github.com/kube-bind/kube-bind/sdk/kcp/clientset/versioned/cluster/typed/kubebind/v1alpha1"
	kubebindv1alpha1 "github.com/kube-bind/kube-bind/sdk/kcp/clientset/versioned/typed/kubebind/v1alpha1"
)

var _ kcpkubebindv1alpha1.KubeBindV1alpha1ClusterInterface = (*KubeBindV1alpha1ClusterClient)(nil)

type KubeBindV1alpha1ClusterClient struct {
	*kcptesting.Fake
}

func (c *KubeBindV1alpha1ClusterClient) Cluster(clusterPath logicalcluster.Path) kubebindv1alpha1.KubeBindV1alpha1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &KubeBindV1alpha1Client{Fake: c.Fake, ClusterPath: clusterPath}
}

func (c *KubeBindV1alpha1ClusterClient) APIServiceBindings() kcpkubebindv1alpha1.APIServiceBindingClusterInterface {
	return &aPIServiceBindingsClusterClient{Fake: c.Fake}
}

func (c *KubeBindV1alpha1ClusterClient) APIServiceExports() kcpkubebindv1alpha1.APIServiceExportClusterInterface {
	return &aPIServiceExportsClusterClient{Fake: c.Fake}
}

func (c *KubeBindV1alpha1ClusterClient) APIServiceExportRequests() kcpkubebindv1alpha1.APIServiceExportRequestClusterInterface {
	return &aPIServiceExportRequestsClusterClient{Fake: c.Fake}
}

func (c *KubeBindV1alpha1ClusterClient) APIServiceNamespaces() kcpkubebindv1alpha1.APIServiceNamespaceClusterInterface {
	return &aPIServiceNamespacesClusterClient{Fake: c.Fake}
}

func (c *KubeBindV1alpha1ClusterClient) ClusterBindings() kcpkubebindv1alpha1.ClusterBindingClusterInterface {
	return &clusterBindingsClusterClient{Fake: c.Fake}
}

var _ kubebindv1alpha1.KubeBindV1alpha1Interface = (*KubeBindV1alpha1Client)(nil)

type KubeBindV1alpha1Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *KubeBindV1alpha1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *KubeBindV1alpha1Client) APIServiceBindings() kubebindv1alpha1.APIServiceBindingInterface {
	return &aPIServiceBindingsClient{Fake: c.Fake, ClusterPath: c.ClusterPath}
}

func (c *KubeBindV1alpha1Client) APIServiceExports(namespace string) kubebindv1alpha1.APIServiceExportInterface {
	return &aPIServiceExportsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *KubeBindV1alpha1Client) APIServiceExportRequests(namespace string) kubebindv1alpha1.APIServiceExportRequestInterface {
	return &aPIServiceExportRequestsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *KubeBindV1alpha1Client) APIServiceNamespaces(namespace string) kubebindv1alpha1.APIServiceNamespaceInterface {
	return &aPIServiceNamespacesClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *KubeBindV1alpha1Client) ClusterBindings(namespace string) kubebindv1alpha1.ClusterBindingInterface {
	return &clusterBindingsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}
