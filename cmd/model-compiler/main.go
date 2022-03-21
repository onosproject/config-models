// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/onosproject/config-models/pkg/compiler"
	"github.com/spf13/cobra"
	"os"
)

const (
	defaultModelPath = "/config-model"
)

func main() {
	if err := getCmd().Execute(); err != nil {
		println(err)
		os.Exit(1)
	}
}

func getCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "model-compiler",
		Short: "Compiles the specified config model",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := defaultModelPath
			if len(args) > 0 {
				path = args[0]
			}
			return compiler.NewCompiler().Compile(path)
		},
	}
	return cmd
}
