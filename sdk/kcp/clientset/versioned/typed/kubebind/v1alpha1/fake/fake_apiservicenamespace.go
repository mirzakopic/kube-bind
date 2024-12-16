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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	gentype "k8s.io/client-go/gentype"

	v1alpha1 "github.com/kube-bind/kube-bind/sdk/apis/kubebind/v1alpha1"
	kubebindv1alpha1 "github.com/kube-bind/kube-bind/sdk/kcp/applyconfiguration/kubebind/v1alpha1"
	typedkubebindv1alpha1 "github.com/kube-bind/kube-bind/sdk/kcp/clientset/versioned/typed/kubebind/v1alpha1"
)

// fakeAPIServiceNamespaces implements APIServiceNamespaceInterface
type fakeAPIServiceNamespaces struct {
	*gentype.FakeClientWithListAndApply[*v1alpha1.APIServiceNamespace, *v1alpha1.APIServiceNamespaceList, *kubebindv1alpha1.APIServiceNamespaceApplyConfiguration]
	Fake *FakeKubeBindV1alpha1
}

func newFakeAPIServiceNamespaces(fake *FakeKubeBindV1alpha1, namespace string) typedkubebindv1alpha1.APIServiceNamespaceInterface {
	return &fakeAPIServiceNamespaces{
		gentype.NewFakeClientWithListAndApply[*v1alpha1.APIServiceNamespace, *v1alpha1.APIServiceNamespaceList, *kubebindv1alpha1.APIServiceNamespaceApplyConfiguration](
			fake.Fake,
			namespace,
			v1alpha1.SchemeGroupVersion.WithResource("apiservicenamespaces"),
			v1alpha1.SchemeGroupVersion.WithKind("APIServiceNamespace"),
			func() *v1alpha1.APIServiceNamespace { return &v1alpha1.APIServiceNamespace{} },
			func() *v1alpha1.APIServiceNamespaceList { return &v1alpha1.APIServiceNamespaceList{} },
			func(dst, src *v1alpha1.APIServiceNamespaceList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha1.APIServiceNamespaceList) []*v1alpha1.APIServiceNamespace {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1alpha1.APIServiceNamespaceList, items []*v1alpha1.APIServiceNamespace) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}