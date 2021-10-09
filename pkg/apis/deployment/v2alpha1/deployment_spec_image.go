//
// DISCLAIMER
//
// Copyright 2016-2021 ArangoDB GmbH, Cologne, Germany
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
// Author Adam Janikowski
//

package v2alpha1

import "github.com/arangodb/kube-arangodb/pkg/util/errors"

type DeploymentImageDiscoveryModeSpec string

func NewDeploymentImageDiscoveryModeSpec(mode DeploymentImageDiscoveryModeSpec) *DeploymentImageDiscoveryModeSpec {
	return &mode
}

const (
	DeploymentImageDiscoveryDirectMode  = "direct"
	DeploymentImageDiscoveryKubeletMode = "kubelet"
)

func (d DeploymentImageDiscoveryModeSpec) Validate() error {
	switch d {
	case DeploymentImageDiscoveryKubeletMode:
		return nil
	case DeploymentImageDiscoveryDirectMode:
		return nil
	default:
		return errors.Newf("mode %s is not supported", d)
	}
}

func (d *DeploymentImageDiscoveryModeSpec) Get() DeploymentImageDiscoveryModeSpec {
	if d == nil {
		return DeploymentImageDiscoveryKubeletMode
	}

	return *d
}
