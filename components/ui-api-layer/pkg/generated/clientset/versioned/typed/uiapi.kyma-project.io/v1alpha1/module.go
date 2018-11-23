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

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/kyma-project/kyma/components/ui-api-layer/pkg/apis/uiapi.kyma-project.io/v1alpha1"
	scheme "github.com/kyma-project/kyma/components/ui-api-layer/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ModulesGetter has a method to return a ModuleInterface.
// A group's client should implement this interface.
type ModulesGetter interface {
	Modules(namespace string) ModuleInterface
}

// ModuleInterface has methods to work with Module resources.
type ModuleInterface interface {
	Create(*v1alpha1.Module) (*v1alpha1.Module, error)
	Update(*v1alpha1.Module) (*v1alpha1.Module, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Module, error)
	List(opts v1.ListOptions) (*v1alpha1.ModuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Module, err error)
	ModuleExpansion
}

// modules implements ModuleInterface
type modules struct {
	client rest.Interface
	ns     string
}

// newModules returns a Modules
func newModules(c *UiapiV1alpha1Client, namespace string) *modules {
	return &modules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the module, and returns the corresponding module object, and an error if there is any.
func (c *modules) Get(name string, options v1.GetOptions) (result *v1alpha1.Module, err error) {
	result = &v1alpha1.Module{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("modules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Modules that match those selectors.
func (c *modules) List(opts v1.ListOptions) (result *v1alpha1.ModuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ModuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("modules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested modules.
func (c *modules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("modules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a module and creates it.  Returns the server's representation of the module, and an error, if there is any.
func (c *modules) Create(module *v1alpha1.Module) (result *v1alpha1.Module, err error) {
	result = &v1alpha1.Module{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("modules").
		Body(module).
		Do().
		Into(result)
	return
}

// Update takes the representation of a module and updates it. Returns the server's representation of the module, and an error, if there is any.
func (c *modules) Update(module *v1alpha1.Module) (result *v1alpha1.Module, err error) {
	result = &v1alpha1.Module{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("modules").
		Name(module.Name).
		Body(module).
		Do().
		Into(result)
	return
}

// Delete takes name of the module and deletes it. Returns an error if one occurs.
func (c *modules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("modules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *modules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("modules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched module.
func (c *modules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Module, err error) {
	result = &v1alpha1.Module{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("modules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
