/*
Copyright 2019 the original author or authors.

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
	"time"

	v1alpha1 "github.com/projectriff/bindings/pkg/apis/bindings/v1alpha1"
	scheme "github.com/projectriff/bindings/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BindableServicesGetter has a method to return a BindableServiceInterface.
// A group's client should implement this interface.
type BindableServicesGetter interface {
	BindableServices(namespace string) BindableServiceInterface
}

// BindableServiceInterface has methods to work with BindableService resources.
type BindableServiceInterface interface {
	Create(*v1alpha1.BindableService) (*v1alpha1.BindableService, error)
	Update(*v1alpha1.BindableService) (*v1alpha1.BindableService, error)
	UpdateStatus(*v1alpha1.BindableService) (*v1alpha1.BindableService, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.BindableService, error)
	List(opts v1.ListOptions) (*v1alpha1.BindableServiceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BindableService, err error)
	BindableServiceExpansion
}

// bindableServices implements BindableServiceInterface
type bindableServices struct {
	client rest.Interface
	ns     string
}

// newBindableServices returns a BindableServices
func newBindableServices(c *BindingsV1alpha1Client, namespace string) *bindableServices {
	return &bindableServices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the bindableService, and returns the corresponding bindableService object, and an error if there is any.
func (c *bindableServices) Get(name string, options v1.GetOptions) (result *v1alpha1.BindableService, err error) {
	result = &v1alpha1.BindableService{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bindableservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BindableServices that match those selectors.
func (c *bindableServices) List(opts v1.ListOptions) (result *v1alpha1.BindableServiceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BindableServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bindableservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested bindableServices.
func (c *bindableServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("bindableservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a bindableService and creates it.  Returns the server's representation of the bindableService, and an error, if there is any.
func (c *bindableServices) Create(bindableService *v1alpha1.BindableService) (result *v1alpha1.BindableService, err error) {
	result = &v1alpha1.BindableService{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("bindableservices").
		Body(bindableService).
		Do().
		Into(result)
	return
}

// Update takes the representation of a bindableService and updates it. Returns the server's representation of the bindableService, and an error, if there is any.
func (c *bindableServices) Update(bindableService *v1alpha1.BindableService) (result *v1alpha1.BindableService, err error) {
	result = &v1alpha1.BindableService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bindableservices").
		Name(bindableService.Name).
		Body(bindableService).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *bindableServices) UpdateStatus(bindableService *v1alpha1.BindableService) (result *v1alpha1.BindableService, err error) {
	result = &v1alpha1.BindableService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bindableservices").
		Name(bindableService.Name).
		SubResource("status").
		Body(bindableService).
		Do().
		Into(result)
	return
}

// Delete takes name of the bindableService and deletes it. Returns an error if one occurs.
func (c *bindableServices) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bindableservices").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *bindableServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bindableservices").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched bindableService.
func (c *bindableServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BindableService, err error) {
	result = &v1alpha1.BindableService{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("bindableservices").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
