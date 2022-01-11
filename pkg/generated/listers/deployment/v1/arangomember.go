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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/arangodb/kube-arangodb/pkg/apis/deployment/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ArangoMemberLister helps list ArangoMembers.
// All objects returned here must be treated as read-only.
type ArangoMemberLister interface {
	// List lists all ArangoMembers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ArangoMember, err error)
	// ArangoMembers returns an object that can list and get ArangoMembers.
	ArangoMembers(namespace string) ArangoMemberNamespaceLister
	ArangoMemberListerExpansion
}

// arangoMemberLister implements the ArangoMemberLister interface.
type arangoMemberLister struct {
	indexer cache.Indexer
}

// NewArangoMemberLister returns a new ArangoMemberLister.
func NewArangoMemberLister(indexer cache.Indexer) ArangoMemberLister {
	return &arangoMemberLister{indexer: indexer}
}

// List lists all ArangoMembers in the indexer.
func (s *arangoMemberLister) List(selector labels.Selector) (ret []*v1.ArangoMember, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ArangoMember))
	})
	return ret, err
}

// ArangoMembers returns an object that can list and get ArangoMembers.
func (s *arangoMemberLister) ArangoMembers(namespace string) ArangoMemberNamespaceLister {
	return arangoMemberNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ArangoMemberNamespaceLister helps list and get ArangoMembers.
// All objects returned here must be treated as read-only.
type ArangoMemberNamespaceLister interface {
	// List lists all ArangoMembers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ArangoMember, err error)
	// Get retrieves the ArangoMember from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ArangoMember, error)
	ArangoMemberNamespaceListerExpansion
}

// arangoMemberNamespaceLister implements the ArangoMemberNamespaceLister
// interface.
type arangoMemberNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ArangoMembers in the indexer for a given namespace.
func (s arangoMemberNamespaceLister) List(selector labels.Selector) (ret []*v1.ArangoMember, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ArangoMember))
	})
	return ret, err
}

// Get retrieves the ArangoMember from the indexer for a given namespace and name.
func (s arangoMemberNamespaceLister) Get(name string) (*v1.ArangoMember, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("arangomember"), name)
	}
	return obj.(*v1.ArangoMember), nil
}
