//
// DISCLAIMER
//
// Copyright 2024 ArangoDB GmbH, Cologne, Germany
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

package integrations

import (
	"context"

	"github.com/spf13/cobra"

	pbImplShutdownV1 "github.com/arangodb/kube-arangodb/integrations/shutdown/v1"
	"github.com/arangodb/kube-arangodb/pkg/util/shutdown"
	"github.com/arangodb/kube-arangodb/pkg/util/svc"
)

func init() {
	register(func() Integration {
		return &shutdownV1{}
	})
}

type shutdownV1 struct {
}

func (s *shutdownV1) Handler(ctx context.Context) (svc.Handler, error) {
	return shutdown.NewGlobalShutdownServer(), nil
}

func (s *shutdownV1) Name() string {
	return pbImplShutdownV1.Name
}

func (s *shutdownV1) Description() string {
	return "ShutdownV1 Handler"
}

func (s *shutdownV1) Register(cmd *cobra.Command, arg ArgGen) error {
	return nil
}
