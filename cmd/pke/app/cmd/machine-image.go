// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/banzaicloud/pke/cmd/pke/app/phases"
	"github.com/banzaicloud/pke/cmd/pke/app/phases/kubeadm/images"
	"github.com/banzaicloud/pke/cmd/pke/app/phases/kubeadm/version"
	"github.com/banzaicloud/pke/cmd/pke/app/phases/runtime/container"
	"github.com/banzaicloud/pke/cmd/pke/app/phases/runtime/kubernetes"
	"github.com/spf13/cobra"
)

// NewCmdImage .
func NewCmdImage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "machine-image",
		Short: "Machine image build helper for Banzai Cloud Pipeline Kubernetes Engine (PKE)",
		RunE:  phases.RunEAllSubcommands,
	}

	cmd.AddCommand(version.NewCommand())
	cmd.AddCommand(container.NewCommand())
	cmd.AddCommand(kubernetes.NewCommand())
	cmd.AddCommand(images.NewCommand())

	phases.MakeRunnable(cmd)

	return cmd
}
