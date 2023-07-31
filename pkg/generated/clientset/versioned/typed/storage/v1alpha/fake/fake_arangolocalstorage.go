//
// DISCLAIMER
//
// Copyright 2023 ArangoDB GmbH, Cologne, Germany
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

package fake

import (
	"context"

	v1alpha "github.com/arangodb/kube-arangodb/pkg/apis/storage/v1alpha"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeArangoLocalStorages implements ArangoLocalStorageInterface
type FakeArangoLocalStorages struct {
	Fake *FakeStorageV1alpha
}

var arangolocalstoragesResource = schema.GroupVersionResource{Group: "storage.arangodb.com", Version: "v1alpha", Resource: "arangolocalstorages"}

var arangolocalstoragesKind = schema.GroupVersionKind{Group: "storage.arangodb.com", Version: "v1alpha", Kind: "ArangoLocalStorage"}

// Get takes name of the arangoLocalStorage, and returns the corresponding arangoLocalStorage object, and an error if there is any.
func (c *FakeArangoLocalStorages) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha.ArangoLocalStorage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(arangolocalstoragesResource, name), &v1alpha.ArangoLocalStorage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.ArangoLocalStorage), err
}

// List takes label and field selectors, and returns the list of ArangoLocalStorages that match those selectors.
func (c *FakeArangoLocalStorages) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha.ArangoLocalStorageList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(arangolocalstoragesResource, arangolocalstoragesKind, opts), &v1alpha.ArangoLocalStorageList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha.ArangoLocalStorageList{ListMeta: obj.(*v1alpha.ArangoLocalStorageList).ListMeta}
	for _, item := range obj.(*v1alpha.ArangoLocalStorageList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested arangoLocalStorages.
func (c *FakeArangoLocalStorages) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(arangolocalstoragesResource, opts))
}

// Create takes the representation of a arangoLocalStorage and creates it.  Returns the server's representation of the arangoLocalStorage, and an error, if there is any.
func (c *FakeArangoLocalStorages) Create(ctx context.Context, arangoLocalStorage *v1alpha.ArangoLocalStorage, opts v1.CreateOptions) (result *v1alpha.ArangoLocalStorage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(arangolocalstoragesResource, arangoLocalStorage), &v1alpha.ArangoLocalStorage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.ArangoLocalStorage), err
}

// Update takes the representation of a arangoLocalStorage and updates it. Returns the server's representation of the arangoLocalStorage, and an error, if there is any.
func (c *FakeArangoLocalStorages) Update(ctx context.Context, arangoLocalStorage *v1alpha.ArangoLocalStorage, opts v1.UpdateOptions) (result *v1alpha.ArangoLocalStorage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(arangolocalstoragesResource, arangoLocalStorage), &v1alpha.ArangoLocalStorage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.ArangoLocalStorage), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeArangoLocalStorages) UpdateStatus(ctx context.Context, arangoLocalStorage *v1alpha.ArangoLocalStorage, opts v1.UpdateOptions) (*v1alpha.ArangoLocalStorage, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(arangolocalstoragesResource, "status", arangoLocalStorage), &v1alpha.ArangoLocalStorage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.ArangoLocalStorage), err
}

// Delete takes name of the arangoLocalStorage and deletes it. Returns an error if one occurs.
func (c *FakeArangoLocalStorages) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(arangolocalstoragesResource, name, opts), &v1alpha.ArangoLocalStorage{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeArangoLocalStorages) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(arangolocalstoragesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha.ArangoLocalStorageList{})
	return err
}

// Patch applies the patch and returns the patched arangoLocalStorage.
func (c *FakeArangoLocalStorages) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha.ArangoLocalStorage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(arangolocalstoragesResource, name, pt, data, subresources...), &v1alpha.ArangoLocalStorage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.ArangoLocalStorage), err
}
