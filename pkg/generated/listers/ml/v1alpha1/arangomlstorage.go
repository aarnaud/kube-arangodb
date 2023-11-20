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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/arangodb/kube-arangodb/pkg/apis/ml/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ArangoMLStorageLister helps list ArangoMLStorages.
// All objects returned here must be treated as read-only.
type ArangoMLStorageLister interface {
	// List lists all ArangoMLStorages in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ArangoMLStorage, err error)
	// ArangoMLStorages returns an object that can list and get ArangoMLStorages.
	ArangoMLStorages(namespace string) ArangoMLStorageNamespaceLister
	ArangoMLStorageListerExpansion
}

// arangoMLStorageLister implements the ArangoMLStorageLister interface.
type arangoMLStorageLister struct {
	indexer cache.Indexer
}

// NewArangoMLStorageLister returns a new ArangoMLStorageLister.
func NewArangoMLStorageLister(indexer cache.Indexer) ArangoMLStorageLister {
	return &arangoMLStorageLister{indexer: indexer}
}

// List lists all ArangoMLStorages in the indexer.
func (s *arangoMLStorageLister) List(selector labels.Selector) (ret []*v1alpha1.ArangoMLStorage, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ArangoMLStorage))
	})
	return ret, err
}

// ArangoMLStorages returns an object that can list and get ArangoMLStorages.
func (s *arangoMLStorageLister) ArangoMLStorages(namespace string) ArangoMLStorageNamespaceLister {
	return arangoMLStorageNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ArangoMLStorageNamespaceLister helps list and get ArangoMLStorages.
// All objects returned here must be treated as read-only.
type ArangoMLStorageNamespaceLister interface {
	// List lists all ArangoMLStorages in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ArangoMLStorage, err error)
	// Get retrieves the ArangoMLStorage from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ArangoMLStorage, error)
	ArangoMLStorageNamespaceListerExpansion
}

// arangoMLStorageNamespaceLister implements the ArangoMLStorageNamespaceLister
// interface.
type arangoMLStorageNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ArangoMLStorages in the indexer for a given namespace.
func (s arangoMLStorageNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ArangoMLStorage, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ArangoMLStorage))
	})
	return ret, err
}

// Get retrieves the ArangoMLStorage from the indexer for a given namespace and name.
func (s arangoMLStorageNamespaceLister) Get(name string) (*v1alpha1.ArangoMLStorage, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("arangomlstorage"), name)
	}
	return obj.(*v1alpha1.ArangoMLStorage), nil
}