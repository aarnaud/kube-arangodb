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

package crds

import (
	_ "embed"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/util/yaml"

	"github.com/arangodb/go-driver"
)

const (
	AppsJobVersion = driver.Version("1.0.1")
)

func init() {
	if err := yaml.Unmarshal(appsJobs, &appsJobsCRD); err != nil {
		panic(err)
	}
}

func AppsJob() *apiextensions.CustomResourceDefinition {
	return appsJobsCRD.DeepCopy()
}

func AppsJobDefinition() Definition {
	return Definition{
		Version: AppsJobVersion,
		CRD:     appsJobsCRD.DeepCopy(),
	}
}

var appsJobsCRD apiextensions.CustomResourceDefinition

//go:embed apps-job.yaml
var appsJobs []byte