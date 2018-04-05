//
// DISCLAIMER
//
// Copyright 2018 ArangoDB GmbH, Cologne, Germany
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
package fake

import (
	v1alpha "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/storage/v1alpha"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeStorageV1alpha struct {
	*testing.Fake
}

func (c *FakeStorageV1alpha) ArangoLocalStorages() v1alpha.ArangoLocalStorageInterface {
	return &FakeArangoLocalStorages{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeStorageV1alpha) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}