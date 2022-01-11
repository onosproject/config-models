// Copyright 2020-present Open Networking Foundation.
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
		RunE:  func(cmd *cobra.Command, args []string) error {
			path := defaultModelPath
			if len(args) > 0 {
				path = args[0]
			}
			return compiler.NewCompiler().Compile(path)
		},
	}
	return cmd
}


