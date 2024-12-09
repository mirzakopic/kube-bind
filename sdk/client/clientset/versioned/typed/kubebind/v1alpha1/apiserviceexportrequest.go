/*
Copyright 2024 The Kube Bind Authors.

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

package v1alpha1

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"

	v1alpha1 "github.com/kube-bind/kube-bind/sdk/apis/kubebind/v1alpha1"
	scheme "github.com/kube-bind/kube-bind/sdk/client/clientset/versioned/scheme"
)

// APIServiceExportRequestsGetter has a method to return a APIServiceExportRequestInterface.
// A group's client should implement this interface.
type APIServiceExportRequestsGetter interface {
	APIServiceExportRequests(namespace string) APIServiceExportRequestInterface
}

// APIServiceExportRequestInterface has methods to work with APIServiceExportRequest resources.
type APIServiceExportRequestInterface interface {
	Create(ctx context.Context, aPIServiceExportRequest *v1alpha1.APIServiceExportRequest, opts v1.CreateOptions) (*v1alpha1.APIServiceExportRequest, error)
	Update(ctx context.Context, aPIServiceExportRequest *v1alpha1.APIServiceExportRequest, opts v1.UpdateOptions) (*v1alpha1.APIServiceExportRequest, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, aPIServiceExportRequest *v1alpha1.APIServiceExportRequest, opts v1.UpdateOptions) (*v1alpha1.APIServiceExportRequest, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.APIServiceExportRequest, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.APIServiceExportRequestList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.APIServiceExportRequest, err error)
	APIServiceExportRequestExpansion
}

// aPIServiceExportRequests implements APIServiceExportRequestInterface
type aPIServiceExportRequests struct {
	*gentype.ClientWithList[*v1alpha1.APIServiceExportRequest, *v1alpha1.APIServiceExportRequestList]
}

// newAPIServiceExportRequests returns a APIServiceExportRequests
func newAPIServiceExportRequests(c *KubeBindV1alpha1Client, namespace string) *aPIServiceExportRequests {
	return &aPIServiceExportRequests{
		gentype.NewClientWithList[*v1alpha1.APIServiceExportRequest, *v1alpha1.APIServiceExportRequestList](
			"apiserviceexportrequests",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.APIServiceExportRequest { return &v1alpha1.APIServiceExportRequest{} },
			func() *v1alpha1.APIServiceExportRequestList { return &v1alpha1.APIServiceExportRequestList{} }),
	}
}