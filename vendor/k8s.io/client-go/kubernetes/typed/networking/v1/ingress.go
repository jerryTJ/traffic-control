/*
Copyright The Kubernetes Authors.

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

package v1

import (
	context "context"

	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationsnetworkingv1 "k8s.io/client-go/applyconfigurations/networking/v1"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// IngressesGetter has a method to return a IngressInterface.
// A group's client should implement this interface.
type IngressesGetter interface {
	Ingresses(namespace string) IngressInterface
}

// IngressInterface has methods to work with Ingress resources.
type IngressInterface interface {
	Create(ctx context.Context, ingress *networkingv1.Ingress, opts metav1.CreateOptions) (*networkingv1.Ingress, error)
	Update(ctx context.Context, ingress *networkingv1.Ingress, opts metav1.UpdateOptions) (*networkingv1.Ingress, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, ingress *networkingv1.Ingress, opts metav1.UpdateOptions) (*networkingv1.Ingress, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*networkingv1.Ingress, error)
	List(ctx context.Context, opts metav1.ListOptions) (*networkingv1.IngressList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *networkingv1.Ingress, err error)
	Apply(ctx context.Context, ingress *applyconfigurationsnetworkingv1.IngressApplyConfiguration, opts metav1.ApplyOptions) (result *networkingv1.Ingress, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, ingress *applyconfigurationsnetworkingv1.IngressApplyConfiguration, opts metav1.ApplyOptions) (result *networkingv1.Ingress, err error)
	IngressExpansion
}

// ingresses implements IngressInterface
type ingresses struct {
	*gentype.ClientWithListAndApply[*networkingv1.Ingress, *networkingv1.IngressList, *applyconfigurationsnetworkingv1.IngressApplyConfiguration]
}

// newIngresses returns a Ingresses
func newIngresses(c *NetworkingV1Client, namespace string) *ingresses {
	return &ingresses{
		gentype.NewClientWithListAndApply[*networkingv1.Ingress, *networkingv1.IngressList, *applyconfigurationsnetworkingv1.IngressApplyConfiguration](
			"ingresses",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *networkingv1.Ingress { return &networkingv1.Ingress{} },
			func() *networkingv1.IngressList { return &networkingv1.IngressList{} },
			gentype.PrefersProtobuf[*networkingv1.Ingress](),
		),
	}
}