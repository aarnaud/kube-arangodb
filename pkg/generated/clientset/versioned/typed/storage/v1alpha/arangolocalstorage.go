//
// DISCLAIMER
//
// Copyright 2016-2022 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

// Code generated by client-gen. DO NOT EDIT.

package v1alpha

import (
	"context"
	"time"

	v1alpha "github.com/arangodb/kube-arangodb/pkg/apis/storage/v1alpha"
	scheme "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ArangoLocalStoragesGetter has a method to return a ArangoLocalStorageInterface.
// A group's client should implement this interface.
type ArangoLocalStoragesGetter interface {
	ArangoLocalStorages() ArangoLocalStorageInterface
}

// ArangoLocalStorageInterface has methods to work with ArangoLocalStorage resources.
type ArangoLocalStorageInterface interface {
	Create(ctx context.Context, arangoLocalStorage *v1alpha.ArangoLocalStorage, opts v1.CreateOptions) (*v1alpha.ArangoLocalStorage, error)
	Update(ctx context.Context, arangoLocalStorage *v1alpha.ArangoLocalStorage, opts v1.UpdateOptions) (*v1alpha.ArangoLocalStorage, error)
	UpdateStatus(ctx context.Context, arangoLocalStorage *v1alpha.ArangoLocalStorage, opts v1.UpdateOptions) (*v1alpha.ArangoLocalStorage, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha.ArangoLocalStorage, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha.ArangoLocalStorageList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha.ArangoLocalStorage, err error)
	ArangoLocalStorageExpansion
}

// arangoLocalStorages implements ArangoLocalStorageInterface
type arangoLocalStorages struct {
	client rest.Interface
}

// newArangoLocalStorages returns a ArangoLocalStorages
func newArangoLocalStorages(c *StorageV1alphaClient) *arangoLocalStorages {
	return &arangoLocalStorages{
		client: c.RESTClient(),
	}
}

// Get takes name of the arangoLocalStorage, and returns the corresponding arangoLocalStorage object, and an error if there is any.
func (c *arangoLocalStorages) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha.ArangoLocalStorage, err error) {
	result = &v1alpha.ArangoLocalStorage{}
	err = c.client.Get().
		Resource("arangolocalstorages").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ArangoLocalStorages that match those selectors.
func (c *arangoLocalStorages) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha.ArangoLocalStorageList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha.ArangoLocalStorageList{}
	err = c.client.Get().
		Resource("arangolocalstorages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested arangoLocalStorages.
func (c *arangoLocalStorages) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("arangolocalstorages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a arangoLocalStorage and creates it.  Returns the server's representation of the arangoLocalStorage, and an error, if there is any.
func (c *arangoLocalStorages) Create(ctx context.Context, arangoLocalStorage *v1alpha.ArangoLocalStorage, opts v1.CreateOptions) (result *v1alpha.ArangoLocalStorage, err error) {
	result = &v1alpha.ArangoLocalStorage{}
	err = c.client.Post().
		Resource("arangolocalstorages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(arangoLocalStorage).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a arangoLocalStorage and updates it. Returns the server's representation of the arangoLocalStorage, and an error, if there is any.
func (c *arangoLocalStorages) Update(ctx context.Context, arangoLocalStorage *v1alpha.ArangoLocalStorage, opts v1.UpdateOptions) (result *v1alpha.ArangoLocalStorage, err error) {
	result = &v1alpha.ArangoLocalStorage{}
	err = c.client.Put().
		Resource("arangolocalstorages").
		Name(arangoLocalStorage.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(arangoLocalStorage).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *arangoLocalStorages) UpdateStatus(ctx context.Context, arangoLocalStorage *v1alpha.ArangoLocalStorage, opts v1.UpdateOptions) (result *v1alpha.ArangoLocalStorage, err error) {
	result = &v1alpha.ArangoLocalStorage{}
	err = c.client.Put().
		Resource("arangolocalstorages").
		Name(arangoLocalStorage.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(arangoLocalStorage).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the arangoLocalStorage and deletes it. Returns an error if one occurs.
func (c *arangoLocalStorages) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("arangolocalstorages").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *arangoLocalStorages) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("arangolocalstorages").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched arangoLocalStorage.
func (c *arangoLocalStorages) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha.ArangoLocalStorage, err error) {
	result = &v1alpha.ArangoLocalStorage{}
	err = c.client.Patch(pt).
		Resource("arangolocalstorages").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
