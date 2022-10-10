/*
Copyright 2022 The Kubectl Bind contributors.

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

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"k8s.io/cli-runtime/pkg/genericclioptions"

	"github.com/kube-bind/kube-bind/pkg/kubectl/bind/plugin"
)

var (
	// TODO: add other examples related to permission claim commands.
	bindExampleUses = `
	# binds to the given remote API service
	%[1]s bind apiservice https://mangodb.com/exports
	`
)

func New(streams genericclioptions.IOStreams) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:              "bind",
		Short:            "Bind different remote types into the current cluster.",
		SilenceUsage:     true,
		TraverseChildren: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	bindOpts := plugin.NewBindOptions(streams)
	bindCmd := &cobra.Command{
		Use:          "apiservice [https://<url>]",
		Short:        "Bind to a remote API service",
		Example:      fmt.Sprintf(bindExampleUses, "kubectl"),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			if err := bindOpts.Complete(args); err != nil {
				return err
			}

			if err := bindOpts.Validate(); err != nil {
				return err
			}

			return bindOpts.Run(cmd.Context())
		},
	}
	bindOpts.BindFlags(bindCmd)

	cmd.AddCommand(bindCmd)
	return cmd, nil
}